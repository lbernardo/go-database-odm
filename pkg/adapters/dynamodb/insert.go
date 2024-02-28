package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
)

type Insert struct {
	model     any
	tableName string
	adapter   *DynamoAdapter
}

func NewInsert(adapter *DynamoAdapter) *Insert {
	return &Insert{adapter: adapter}
}

func (i *Insert) Model(model any) adapters.Insert {
	i.model = model
	i.tableName = i.adapter.TableName(model)
	return i
}

func (i *Insert) Exec(ctx context.Context) (*adapters.InsertResult, error) {
	if i.model == nil {
		return nil, fmt.Errorf(".Model(model any) is required")
	}
	item, err := attributevalue.MarshalMap(i.model)
	if err != nil {
		return nil, fmt.Errorf("error to decode marshal %v", err)
	}
	output, err := i.adapter.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(i.tableName),
		Item:      item,
	})
	if err != nil {
		return nil, fmt.Errorf("error to putItem %v", err)
	}
	return &adapters.InsertResult{
		Result: output.ResultMetadata,
	}, nil
}
