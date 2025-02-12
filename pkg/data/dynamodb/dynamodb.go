package dynamodb

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"

	dynamodbModel "github.com/jhonquirama/hexa-scaffolding-ms/pkg/data/dynamodb/model"
)

//go:generate mockery --name Dynamodb
//go:generate mockery --name QueryOptions
type (
	Config interface {
		MaxRetries() int
		MaxBackoffDelaySecond() int
	}

	QueryOptions                     func(input *dynamodb.QueryInput)
	ExpressionAttributeValuesOptions func(input *dynamodb.PutItemInput)

	Dynamodb interface {
		Query(ctx context.Context, tableName string, keyCondition, filter ConditionalValues, result interface{},
			pagination interface{}, options ...QueryOptions) error
		QueryAll(ctx context.Context, tableName string, keyCondition, filter ConditionalValues, result interface{},
			pagination interface{}, limitTotal int32, options ...QueryOptions) error
		GetItem(ctx context.Context, tableName string, keyCondition ConditionalValues, result interface{}) error
		BatchGetItemAll(ctx context.Context, tableName string, keyConditions []ConditionalValues, result interface{}) error
		DeleteItem(ctx context.Context, tableName string, filter ConditionalValues) error
		PutItem(ctx context.Context, tableName string, item interface{}) error

		TransactWriteItems(ctx context.Context, data TransactWriteItems) error

		WithPagination(pagination map[string]interface{}) QueryOptions
		WithQueryIndex(indexName *string) QueryOptions
		WithLimit(limit *int32) QueryOptions

		TransactWriteItemsCall(ctx context.Context, input dynamodbModel.TransactWriteItemsInputCall) error

		QueryAllCall(
			ctx context.Context, tableName string, filter dynamodbModel.QueryAllInputCall,
			result interface{}, pagination interface{}, iterations int32, options ...QueryOptions) error
		GetItemCall(
			ctx context.Context, tableName string, filter dynamodbModel.GetItemInputCall, result interface{}) error
		BatchGetItemAllCall(
			ctx context.Context, filter dynamodbModel.BatchGetItemInputCall, result interface{}) error
		UpdateItem(ctx context.Context, tableName string, input dynamodbModel.UpdateItemInput) error
	}

	dynamo struct {
		client *dynamodb.Client
	}
)

func NewDynamoDB(ctx context.Context, conf Config) (Dynamodb, error) {
	var (
		c   aws.Config
		err error
	)
	if c, err = config.LoadDefaultConfig(ctx, config.WithRetryer(func() aws.Retryer {
		maxBackoffDelay := time.Duration(conf.MaxBackoffDelaySecond()) * time.Second
		newRetry := retry.AddWithMaxBackoffDelay(retry.NewStandard(), maxBackoffDelay)
		return retry.AddWithMaxAttempts(newRetry, conf.MaxRetries())
	})); err != nil {
		return nil, err
	}
	return &dynamo{
		client: dynamodb.NewFromConfig(c),
	}, nil
}

func (d *dynamo) Query(
	ctx context.Context, tableName string, keyCondition, filter ConditionalValues,
	result interface{}, pagination interface{}, options ...QueryOptions) error {
	var (
		queryOutput *dynamodb.QueryOutput
		builder     expression.Builder
		expr        expression.Expression
		err         error
	)
	if builder, err = d.getBuilderWithKeyConditions(keyCondition); err != nil {
		return err
	}
	builder = d.getBuilderWithFilter(builder, filter)
	if expr, err = builder.Build(); err != nil {
		return err
	}

	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 &tableName,
	}
	for _, option := range options {
		option(params)
	}
	if queryOutput, err = d.client.Query(ctx, params); err != nil {
		return err
	}
	if queryOutput == nil || len(queryOutput.Items) == 0 {
		return nil
	}
	if err = attributevalue.UnmarshalListOfMaps(queryOutput.Items, &result); err != nil {
		return err
	}
	if pagination != nil {
		if err = attributevalue.UnmarshalMap(queryOutput.LastEvaluatedKey, &pagination); err != nil {
			return err
		}
	}
	return nil
}

func (d *dynamo) QueryAll(
	ctx context.Context, tableName string, keyCondition, filter ConditionalValues,
	result interface{}, pagination interface{}, iterations int32, options ...QueryOptions) error {
	var (
		queryPaginator *dynamodb.QueryPaginator
		builder        expression.Builder
		expr           expression.Expression
		err            error
	)
	if builder, err = d.getBuilderWithKeyConditions(keyCondition); err != nil {
		return err
	}
	builder = d.getBuilderWithFilter(builder, filter)
	if expr, err = builder.Build(); err != nil {
		return err
	}

	params := &dynamodb.QueryInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 &tableName,
	}
	for _, option := range options {
		option(params)
	}
	// se elimina el limit si se pasa en las options.
	// IMPORTANTE: no eliminar esta linea
	params.Limit = nil

	var (
		allItems []map[string]types.AttributeValue
		i        = int32(1)
	)
	queryPaginator = dynamodb.NewQueryPaginator(d.client, params)

	for queryPaginator.HasMorePages() {
		var (
			queryOutput *dynamodb.QueryOutput
		)
		if queryOutput, err = queryPaginator.NextPage(ctx); err != nil {
			// *** OJO: REVISAR *** . Error lanzo dynamodb por hacer varias petciones seguidas.
			// *** The level of configured provisioned throughput for one or more global secondary indexes
			// *** of the table was exceeded. Consider increasing your provisioning level for the
			// *** under-provisioned global secondary indexes with the UpdateTable API error
			return err
		}
		allItems = append(allItems, queryOutput.Items...)
		// cantidad de iteraciones a realizar el ciclo de paginacion
		// solo se evaluara si se pasa el parametro iterations
		if iterations > 0 && iterations <= i {
			if pagination != nil {
				if err = attributevalue.UnmarshalMap(queryOutput.LastEvaluatedKey, &pagination); err != nil {
					return err
				}
			}
			break
		}
		i++
	}
	if err = attributevalue.UnmarshalListOfMaps(allItems, &result); err != nil {
		return err
	}
	return nil
}

func (d *dynamo) DeleteItem(ctx context.Context, tableName string, filter ConditionalValues) error {
	var (
		attributeValue map[string]types.AttributeValue
		err            error
	)
	if attributeValue, err = d.getAttributeKey(filter); err != nil {
		return err
	}
	params := &dynamodb.DeleteItemInput{
		Key:       attributeValue,
		TableName: &tableName,
	}
	if _, err = d.client.DeleteItem(ctx, params); err != nil {
		return err
	}
	return nil
}

func (d *dynamo) GetItem(
	ctx context.Context, tableName string, keyCondition ConditionalValues, result interface{}) error {
	var (
		attributesValueAll map[string]types.AttributeValue
		itemOutput         *dynamodb.GetItemOutput
		err                error
	)
	if attributesValueAll, err = d.getAttributeKey(keyCondition); err != nil {
		return err
	}

	params := &dynamodb.GetItemInput{
		Key:       attributesValueAll,
		TableName: &tableName,
	}
	if itemOutput, err = d.client.GetItem(ctx, params); err != nil {
		return err
	}

	if err = attributevalue.UnmarshalMap(itemOutput.Item, &result); err != nil {
		return err
	}

	return nil
}

func (d *dynamo) BatchGetItemAll(
	ctx context.Context, tableName string, keyCondition []ConditionalValues, result interface{}) error {
	var (
		attributesValueAll      []map[string]types.AttributeValue
		tableItems              []map[string]types.AttributeValue
		totalattributesValueAll int
		err                     error
		size                    = 100 // batch de 100(restriction de dynamodb)
	)
	if attributesValueAll, err = d.getAttributeKeys(keyCondition); err != nil {
		return err
	}
	totalattributesValueAll = len(attributesValueAll)
	for start := 0; start < totalattributesValueAll; start += size {
		end := start + size

		if end > totalattributesValueAll {
			end = totalattributesValueAll
		}
		var (
			batchGetItemOutput *dynamodb.BatchGetItemOutput
			attributesValue    = attributesValueAll[start:end]
		)
		params := &dynamodb.BatchGetItemInput{
			RequestItems: map[string]types.KeysAndAttributes{
				tableName: {
					Keys: attributesValue,
				},
			},
		}
		if batchGetItemOutput, err = d.client.BatchGetItem(ctx, params); err != nil {
			return err
		}
		if items, ok := batchGetItemOutput.Responses[tableName]; ok {
			tableItems = append(tableItems, items...)
		}
	}
	if len(tableItems) > 0 {
		if err = attributevalue.UnmarshalListOfMaps(tableItems, &result); err != nil {
			return err
		}
	}

	return nil
}

func (d *dynamo) PutItem(ctx context.Context, tableName string, item interface{}) error {
	var (
		attributeValue map[string]types.AttributeValue
		err            error
	)
	if attributeValue, err = attributevalue.MarshalMap(item); err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      attributeValue,
		TableName: &tableName,
	}
	if _, err = d.client.PutItem(ctx, input); err != nil {
		return err
	}
	return nil
}

func (d *dynamo) TransactWriteItems(ctx context.Context, input TransactWriteItems) error {
	var (
		putItemData   map[string]types.AttributeValue
		err           error
		transactItems = []types.TransactWriteItem{}
	)
	for _, put := range input.Put {
		if len(put.TableName) > 0 && put.Item != nil {
			if putItemData, err = attributevalue.MarshalMap(put.Item); err != nil {
				return err
			}
			transactItems = append(transactItems, types.TransactWriteItem{
				Put: &types.Put{
					Item:      putItemData,
					TableName: aws.String(put.TableName),
				},
			})
		}
	}
	for _, update := range input.Update {
		if len(update.TableName) > 0 && update.Data != nil && len(update.Data) > 0 {
			var (
				builder   expression.Builder
				expr      expression.Expression
				keyUpdate map[string]types.AttributeValue
			)
			if builder, err = d.getBuilderWithUpdate(update.Data); err != nil {
				return err
			}
			if expr, err = builder.Build(); err != nil {
				return err
			}
			if keyUpdate, err = d.getAttributeKey(update.KeyCondition); err != nil {
				return err
			}
			transactItems = append(transactItems, types.TransactWriteItem{
				Update: &types.Update{
					ConditionExpression:       expr.Condition(),
					ExpressionAttributeNames:  expr.Names(),
					ExpressionAttributeValues: expr.Values(),
					Key:                       keyUpdate,
					TableName:                 aws.String(update.TableName),
					UpdateExpression:          expr.Update(),
				},
			})
		}
	}

	if _, err = d.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: transactItems,
	}); err != nil {
		return err
	}
	return nil
}

// --- nueva forma de implementar acceso a dynamodb --- //

func (d *dynamo) TransactWriteItemsCall(ctx context.Context, input dynamodbModel.TransactWriteItemsInputCall) error {
	var (
		putItemData map[string]types.AttributeValue
		err         error

		transactItems = []types.TransactWriteItem{}
		size          = 100 // restriction de dynamodb 100 items
	)
	for _, put := range input.Put {
		if len(put.TableName) > 0 && put.Item != nil {
			if putItemData, err = attributevalue.MarshalMap(put.Item); err != nil {
				return err
			}
			transactItems = append(transactItems, types.TransactWriteItem{
				Put: &types.Put{
					Item:      putItemData,
					TableName: aws.String(put.TableName),
				},
			})
		}
	}
	for _, update := range input.Update {
		if len(update.TableName) > 0 && update.Key != nil && len(update.Key) > 0 {
			var (
				keyValue       map[string]types.AttributeValue
				attributeValue map[string]types.AttributeValue
				err            error
			)
			if keyValue, err = attributevalue.MarshalMap(update.Key); err != nil {
				return err
			}
			if attributeValue, err = attributevalue.MarshalMap(update.AttributeValues); err != nil {
				return err
			}
			transactItems = append(transactItems, types.TransactWriteItem{
				Update: &types.Update{
					TableName:                 aws.String(update.TableName),
					Key:                       keyValue,
					UpdateExpression:          update.UpdateExpression,
					ConditionExpression:       update.ConditionalExpression,
					ExpressionAttributeNames:  update.AttributeNames,
					ExpressionAttributeValues: attributeValue,
				},
			})
		}
	}
	lenTransactItems := len(transactItems)
	for start := 0; start < lenTransactItems; start += size {
		end := start + size

		if end > lenTransactItems {
			end = lenTransactItems
		}
		if _, err = d.client.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
			TransactItems: transactItems[start:end],
		}); err != nil {
			return err
		}
	}

	return nil
}

func (d *dynamo) QueryAllCall(
	ctx context.Context, tableName string, filter dynamodbModel.QueryAllInputCall,
	result interface{}, pagination interface{}, iterations int32, options ...QueryOptions) error {
	var (
		attributeValues map[string]types.AttributeValue
		queryPaginator  *dynamodb.QueryPaginator
		err             error
	)

	if attributeValues, err = attributevalue.MarshalMap(filter.AttributeValues); err != nil {
		return err
	}

	params := &dynamodb.QueryInput{
		TableName:                 &tableName,
		KeyConditionExpression:    filter.KeyExpression,
		FilterExpression:          filter.FilterExpression,
		ExpressionAttributeNames:  filter.AttributeNames,
		ExpressionAttributeValues: attributeValues,
	}
	for _, option := range options {
		option(params)
	}
	// se elimina el limit si se pasa en las options.
	// IMPORTANTE: no eliminar esta linea
	params.Limit = nil

	var (
		allItems []map[string]types.AttributeValue
		i        = int32(1)
	)
	queryPaginator = dynamodb.NewQueryPaginator(d.client, params)

	for queryPaginator.HasMorePages() {
		var (
			queryOutput *dynamodb.QueryOutput
		)
		if queryOutput, err = queryPaginator.NextPage(ctx); err != nil {
			// *** OJO: REVISAR *** . Error lanzo dynamodb por hacer varias petciones seguidas.
			// *** The level of configured provisioned throughput for one or more global secondary indexes
			// *** of the table was exceeded. Consider increasing your provisioning level for the
			// *** under-provisioned global secondary indexes with the UpdateTable API error
			return err
		}
		allItems = append(allItems, queryOutput.Items...)
		// cantidad de iteraciones a realizar el ciclo de paginacion
		// solo se evaluara si se pasa el parametro iterations
		if iterations > 0 && iterations <= i {
			if pagination != nil {
				if err = attributevalue.UnmarshalMap(queryOutput.LastEvaluatedKey, &pagination); err != nil {
					return err
				}
			}
			break
		}
		i++
	}
	if err = attributevalue.UnmarshalListOfMaps(allItems, &result); err != nil {
		return err
	}
	return nil
}

func (d *dynamo) GetItemCall(
	ctx context.Context, tableName string, filter dynamodbModel.GetItemInputCall, result interface{}) error {
	var (
		attributeValues map[string]types.AttributeValue
		itemOutput      *dynamodb.GetItemOutput
		err             error
	)
	if attributeValues, err = attributevalue.MarshalMap(filter.Key); err != nil {
		return err
	}

	params := &dynamodb.GetItemInput{
		Key:       attributeValues,
		TableName: &tableName,
	}
	if itemOutput, err = d.client.GetItem(ctx, params); err != nil {
		return err
	}

	if err = attributevalue.UnmarshalMap(itemOutput.Item, &result); err != nil {
		return err
	}

	return nil
}

func (d *dynamo) BatchGetItemAllCall(
	ctx context.Context, filter dynamodbModel.BatchGetItemInputCall, result interface{}) error {
	var (
		tableItems []map[string]types.AttributeValue
		err        error
		size       = 100 // batch de 100(restriction de dynamodb)
	)
	var (
		areqItems = map[string]types.KeysAndAttributes{}
		tables    = []string{}
	)
	for table, itemsRawKeys := range filter.TableKeys {
		var keys = []map[string]types.AttributeValue{}

		for _, itemRawKey := range itemsRawKeys {
			var key map[string]types.AttributeValue

			if key, err = attributevalue.MarshalMap(itemRawKey); err != nil {
				return err
			}
			keys = append(keys, key)
		}
		if len(keys) > 0 {
			areqItems[table] = types.KeysAndAttributes{
				Keys: keys,
			}
			tables = append(tables, table)
		}
	}
	var (
		newReqItems = []map[string]types.KeysAndAttributes{}
	)
	for _, table := range tables {
		// TODO: optimizar el codigo cuando son tablas diferentes y no se llega a los 100 items
		if keys, ok := areqItems[table]; ok {
			if len(keys.Keys) <= size {
				item := map[string]types.KeysAndAttributes{
					table: keys,
				}
				newReqItems = append(newReqItems, item)
			} else {
				var (
					keysAll   = keys.Keys
					totalKeys = len(keysAll)
				)
				for start := 0; start < totalKeys; start += size {
					end := start + size

					if end > totalKeys {
						end = totalKeys
					}
					var (
						keysPart = keysAll[start:end]
					)
					item := map[string]types.KeysAndAttributes{
						table: {
							Keys: keysPart,
						},
					}
					newReqItems = append(newReqItems, item)
				}
			}
		}
	}
	for _, reqItems := range newReqItems {
		params := &dynamodb.BatchGetItemInput{
			RequestItems: reqItems,
		}
		var (
			batchGetItemOutput *dynamodb.BatchGetItemOutput
		)
		if batchGetItemOutput, err = d.client.BatchGetItem(ctx, params); err != nil {
			return err
		}
		for _, items := range batchGetItemOutput.Responses {
			tableItems = append(tableItems, items...)
		}
	}
	if len(tableItems) > 0 {
		if err = attributevalue.UnmarshalListOfMaps(tableItems, &result); err != nil {
			return err
		}
	}
	return nil
}

func (d *dynamo) UpdateItem(ctx context.Context, tableName string, input dynamodbModel.UpdateItemInput) error {
	var (
		keyValue       map[string]types.AttributeValue
		attributeValue map[string]types.AttributeValue
		err            error
	)
	if keyValue, err = attributevalue.MarshalMap(input.Key); err != nil {
		return err
	}
	if attributeValue, err = attributevalue.MarshalMap(input.AttributeValues); err != nil {
		return err
	}
	in := &dynamodb.UpdateItemInput{
		TableName:                 &tableName,
		Key:                       keyValue,
		UpdateExpression:          input.UpdateExpression,
		ConditionExpression:       input.ConditionalExpression,
		ExpressionAttributeValues: attributeValue,
		ExpressionAttributeNames:  input.AttributeNames,
	}
	if _, err = d.client.UpdateItem(ctx, in); err != nil {
		return err
	}
	return nil
}

// --- helpers --- //

func (d *dynamo) WithQueryIndex(indexName *string) QueryOptions {
	return func(input *dynamodb.QueryInput) {
		input.IndexName = indexName
	}
}
func (d *dynamo) WithLimit(limit *int32) QueryOptions {
	return func(input *dynamodb.QueryInput) {
		if limit == nil || *limit <= 0 {
			input.Limit = nil
		} else {
			input.Limit = limit
		}
	}
}
func (d *dynamo) WithPagination(pagination map[string]interface{}) QueryOptions {
	return func(input *dynamodb.QueryInput) {
		attributeValue, _ := attributevalue.MarshalMap(pagination)
		if len(attributeValue) > 0 {
			input.ExclusiveStartKey = attributeValue
		}
	}
}
