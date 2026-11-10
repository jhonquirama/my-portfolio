package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	OperatorNone OperatorType = iota
	OperatorEqual
	OperatorLessThan
	OperatorLessThanEqual
	OperatorGreaterThan
	OperatorGreaterThanEqual
	OperatorLogicAnd
)

type (
	OperatorType uint

	ConditionalValue struct {
		Key                       string
		Value                     interface{}
		Operator                  OperatorType
		LogicOperator             OperatorType
		ConditionExpression       string
		ExpressionAttributeValues map[string]types.AttributeValue
	}
	ConditionalValues []ConditionalValue
)

// --- helpers --- //

func (*dynamo) getBuilderWithKeyConditions(keyCondition ConditionalValues) (expression.Builder, error) {
	var (
		firstValue          = true
		keyConditionBuilder expression.KeyConditionBuilder
		builder             expression.Builder
	)
	for _, item := range keyCondition {
		key := item.Key
		value := item.Value
		if firstValue {
			keyConditionBuilder = expression.Key(key).Equal(expression.Value(value))
			firstValue = false
		} else {
			keyBuilder := expression.Key(key)
			valueBuilder := expression.Value(value)
			keyConditionBuilderResult := keyBuilder.Equal(valueBuilder)

			switch item.Operator {
			case OperatorLessThan:
				keyConditionBuilderResult = keyBuilder.LessThan(valueBuilder)
			case OperatorLessThanEqual:
				keyConditionBuilderResult = keyBuilder.LessThanEqual(valueBuilder)
			}
			// siempre se aplica AND porque es condicion de clave primaria
			keyConditionBuilder = keyConditionBuilder.And(keyConditionBuilderResult)
		}
	}
	builder = expression.NewBuilder().WithKeyCondition(keyConditionBuilder)
	return builder, nil
}

func (*dynamo) getBuilderWithFilter(builder expression.Builder, filter ConditionalValues) expression.Builder {
	var (
		firstValue    = true
		filterBuilder expression.ConditionBuilder
	)
	if len(filter) == 0 {
		return builder
	}
	for _, item := range filter {
		key := item.Key
		value := item.Value
		if firstValue {
			filterBuilder = expression.Name(key).Equal(expression.Value(value))
			firstValue = false
		} else {
			keyBuilder := expression.Name(key)
			valueBuilder := expression.Value(value)
			filterBuilderResult := keyBuilder.Equal(valueBuilder)

			switch item.Operator {
			case OperatorLessThan:
				filterBuilderResult = keyBuilder.LessThan(valueBuilder)
			case OperatorLessThanEqual:
				filterBuilderResult = keyBuilder.LessThanEqual(valueBuilder)
			}
			switch item.LogicOperator {
			case OperatorLogicAnd:
				filterBuilder = filterBuilder.And(filterBuilderResult)
			default:
				filterBuilder = filterBuilder.And(filterBuilderResult)
			}
		}
	}
	builder = builder.WithFilter(filterBuilder)
	return builder
}

func (*dynamo) getAttributeKey(keyCondition ConditionalValues) (map[string]types.AttributeValue, error) {
	var (
		attributeValue map[string]types.AttributeValue
		err            error
		keyValues      = map[string]interface{}{}
	)
	for _, item := range keyCondition {
		keyValues[item.Key] = item.Value
	}
	if attributeValue, err = attributevalue.MarshalMap(keyValues); err != nil {
		return attributeValue, err
	}
	return attributeValue, nil
}
func (*dynamo) getAttributeKeys(keyConditions []ConditionalValues) ([]map[string]types.AttributeValue, error) {
	var (
		attributeValues []map[string]types.AttributeValue
		err             error
	)
	for _, kc := range keyConditions {
		for _, item := range kc {
			var (
				attributeValue map[string]types.AttributeValue
				keyValues      = map[string]interface{}{}
			)
			keyValues[item.Key] = item.Value
			if attributeValue, err = attributevalue.MarshalMap(keyValues); err != nil {
				return attributeValues, err
			}
			attributeValues = append(attributeValues, attributeValue)
		}
	}

	return attributeValues, nil
}
