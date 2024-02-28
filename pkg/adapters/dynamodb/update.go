package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
)

type UpdateDynamodb struct {
	adapter    *DynamoAdapter
	model      any
	tableName  string
	conditions map[string]types.AttributeValue
	setters    map[string]any
}

func NewUpdateDynamodb(adapter *DynamoAdapter) *UpdateDynamodb {
	return &UpdateDynamodb{adapter: adapter, conditions: map[string]types.AttributeValue{}}
}

func (u *UpdateDynamodb) Model(model any) adapters.Update {
	u.model = model
	u.tableName = u.adapter.TableName(model)
	return u
}

func (u *UpdateDynamodb) Condition(name string, value any) adapters.Update {
	v, err := attributevalue.Marshal(value)
	if err != nil {
		fmt.Println("error to create attribute condition", err.Error())
		return u
	}
	u.conditions[name] = v
	return u
}

func (u *UpdateDynamodb) Set(name string, value any) adapters.Update {
	if u.setters == nil {
		u.setters = map[string]any{}
	}
	u.setters[name] = value
	return u
}

func (u *UpdateDynamodb) Exec(ctx context.Context) error {
	if u.model == nil {
		return fmt.Errorf(".Model(model any) is required")
	}
	if u.setters == nil {
		return fmt.Errorf(".Set is required")
	}
	var setted = false
	var update expression.UpdateBuilder
	for name, value := range u.setters {
		if !setted {
			setted = true
			update = expression.Set(expression.Name(name), expression.Value(value))
			continue
		}
		update = update.Set(expression.Name(name), expression.Value(value))
	}
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		return fmt.Errorf("error to create conditions with expression.NewBuilder (conditions: %v) err: %v", u.conditions, err)
	}
	if _, err := u.adapter.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 aws.String(u.tableName),
		Key:                       u.conditions,
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}); err != nil {
		return fmt.Errorf("error to updateItem %v", err)
	}
	return nil
}
