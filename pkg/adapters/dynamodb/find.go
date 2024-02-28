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

type FindDynamodb struct {
	adapter    *DynamoAdapter
	model      any
	tableName  string
	conditions map[string]types.AttributeValue
}

func NewFindDynamodb(adapter *DynamoAdapter) *FindDynamodb {
	return &FindDynamodb{adapter: adapter, conditions: map[string]types.AttributeValue{}}
}

func (f *FindDynamodb) Model(model any) adapters.Find {
	f.model = model
	f.tableName = f.adapter.TableName(model)
	return f
}

func (f *FindDynamodb) Condition(name string, value any) adapters.Find {
	v, err := attributevalue.Marshal(value)
	if err != nil {
		fmt.Println("error to create attribute condition", err.Error())
		return f
	}
	f.conditions[name] = v
	return f
}

func (f *FindDynamodb) Exec(ctx context.Context) error {
	if f.model == nil {
		return fmt.Errorf(".Model(model any) is required")
	}
	output, err := f.adapter.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(f.tableName),
		Key:       f.conditions,
	})
	if err != nil {
		return fmt.Errorf("error to get item: %v", err)
	}
	if err := attributevalue.UnmarshalMap(output.Item, f.model); err != nil {
		return fmt.Errorf("error to unmarshalmap: %v", err)
	}
	return nil
}
