package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
)

type DeleteDynamodb struct {
	adapter    *DynamoAdapter
	model      any
	tableName  string
	conditions map[string]types.AttributeValue
}

func NewDeleteDynamodb(adapter *DynamoAdapter) *DeleteDynamodb {
	return &DeleteDynamodb{adapter: adapter, conditions: map[string]types.AttributeValue{}}
}

func (d *DeleteDynamodb) Model(model any) adapters.Delete {
	d.model = model
	d.tableName = d.adapter.TableName(model)
	return d
}

func (d *DeleteDynamodb) Condition(name string, value any) adapters.Delete {
	v, err := attributevalue.Marshal(value)
	if err != nil {
		fmt.Println("error to create attribute condition", err.Error())
		return d
	}
	d.conditions[name] = v
	return d
}

func (d *DeleteDynamodb) Exec(ctx context.Context) error {
	if d.model == nil {
		return fmt.Errorf(".Model(model any) is required")
	}
	if _, err := d.adapter.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(d.tableName),
		Key:       d.conditions,
	}); err != nil {
		return fmt.Errorf("error to deleteItem: %v", err)
	}
	return nil
}
