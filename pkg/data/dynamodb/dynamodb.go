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

	dynamodbModel "github.com/jhonquirama/my-portfolio/pkg/data/dynamodb/model"
)

type (
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
		PutItemIfNotExist(ctx context.Context, tableName string,
			ConditionExpression string, item interface{}, filter ConditionalValues) error
		UpdateItem(ctx context.Context,
			tableName string, keyCondition ConditionalValues, data map[string]any) error
		WithPagination(pagination map[string]interface{}) QueryOptions
		WithQueryIndex(indexName *string) QueryOptions
		WithLimit(limit *int32) QueryOptions

		// --- nueva forma de implementar acceso a dynamodb --- //

		UpdateItemCall(ctx context.Context, tableName string, input dynamodbModel.UpdateItemInput) error
		QueryCall(ctx context.Context, tableName string,
			filter dynamodbModel.QueryCallInput, result interface{}, pagination interface{}, options ...QueryOptions) error
	}

	dynamo struct {
		client *dynamodb.Client
	}

	Config interface {
		MaxRetries() int
		MaxBackoffDelaySecond() int
	}
)

func NewDynamoDB(ctx context.Context, conf Config) (Dynamodb, error) {
	awsConfig, err := config.LoadDefaultConfig(ctx,
		config.WithRetryer(func() aws.Retryer {
			maxBackoffDelay := time.Duration(conf.MaxBackoffDelaySecond()) * time.Second
			newRetry := retry.AddWithMaxBackoffDelay(retry.NewStandard(), maxBackoffDelay)
			return retry.AddWithMaxAttempts(newRetry, conf.MaxRetries())
		}))
	if err != nil {
		return nil, err
	}
	return &dynamo{
		client: dynamodb.NewFromConfig(awsConfig),
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

func (d *dynamo) PutItemIfNotExist(ctx context.Context, tableName string,
	conditionExpression string, item interface{}, filter ConditionalValues) error {
	var (
		attributeValue  map[string]types.AttributeValue
		attributeValues map[string]types.AttributeValue
		err             error
	)

	if attributeValue, err = attributevalue.MarshalMap(item); err != nil {
		return err
	}
	if attributeValues, err = d.getAttributeKey(filter); err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:                      attributeValue,
		ExpressionAttributeValues: attributeValues,
		ConditionExpression:       aws.String(conditionExpression),
		TableName:                 &tableName,
	}

	if _, err = d.client.PutItem(ctx, input); err != nil {
		return err
	}
	return nil
}

func (d *dynamo) UpdateItem(ctx context.Context,
	tableName string, keyCondition ConditionalValues, data map[string]any) error {
	var (
		builder   expression.Builder
		expr      expression.Expression
		keyUpdate map[string]types.AttributeValue
		err       error
	)
	if builder, err = d.getBuilderWithUpdate(data); err != nil {
		return err
	}
	if expr, err = builder.Build(); err != nil {
		return err
	}
	if keyUpdate, err = d.getAttributeKey(keyCondition); err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		Key:                       keyUpdate,
		TableName:                 &tableName,
		ConditionExpression:       expr.Condition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	if _, err = d.client.UpdateItem(ctx, input); err != nil {
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
func (*dynamo) getBuilderWithUpdate(data map[string]any) (expression.Builder, error) {
	var (
		firstValue    = true
		updateBuilder expression.UpdateBuilder
	)

	for key, value := range data {
		if firstValue {
			updateBuilder = expression.Set(expression.Name(key), expression.Value(value))
			firstValue = false
		} else {
			updateBuilder = updateBuilder.Set(expression.Name(key), expression.Value(value))
		}
	}
	return expression.NewBuilder().WithUpdate(updateBuilder), nil
}

// --- nueva forma de implementar acceso a dynamodb --- //

func (d *dynamo) UpdateItemCall(ctx context.Context, tableName string, input dynamodbModel.UpdateItemInput) error {
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

func (d *dynamo) QueryCall(ctx context.Context, tableName string,
	filter dynamodbModel.QueryCallInput, result interface{}, pagination interface{}, options ...QueryOptions) error {
	var (
		attributeValues map[string]types.AttributeValue
		queryOutput     *dynamodb.QueryOutput
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
