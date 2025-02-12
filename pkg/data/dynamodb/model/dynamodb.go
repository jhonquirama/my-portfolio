package model

type (
	QueryAllInputCall struct {
		AttributeValues  map[string]any
		AttributeNames   map[string]string
		KeyExpression    *string
		FilterExpression *string
	}
	GetItemInputCall struct {
		Key map[string]any
	}
	BatchGetItemInputCall struct {
		// TableKeys is a map of table name to a list of keys
		// {
		// tablaName1: [
		// 		{keyPK: keyValue1},
		// 		{keyPK: keyValue2},
		// ],
		// tablaName2: [
		// 		{keyPK: keyValue1},
		// 		{keyPK: keyValue2},
		// ]
		// }
		TableKeys map[string][]map[string]any
	}
	PutCall struct {
		TableName string
		Item      any
	}
	UpdateCall struct {
		TableName             string
		Key                   map[string]any
		UpdateExpression      *string
		ConditionalExpression *string
		AttributeNames        map[string]string
		AttributeValues       map[string]any
	}
	TransactWriteItemsInputCall struct {
		Put    []PutCall
		Update []UpdateCall
	}

	UpdateItemInput struct {
		Key                   map[string]any
		AttributeValues       map[string]any
		AttributeNames        map[string]string
		UpdateExpression      *string
		ConditionalExpression *string
	}
)
