package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
)

type Select struct {
	adapter    *DynamoAdapter
	model      any
	tableName  string
	conditions []expression.ConditionBuilder
}

func NewSelect(adapter *DynamoAdapter) *Select {
	return &Select{adapter: adapter, conditions: []expression.ConditionBuilder{}}
}

func (s *Select) Model(model any) adapters.Select {
	s.tableName = s.adapter.TableName(model)
	s.model = model
	return s
}

func (s *Select) Condition(name string, value any) adapters.Select {
	filter := expression.Name(name).Equal(expression.Value(value))
	s.conditions = append(s.conditions, filter)
	return s
}

func (s *Select) GreaterThan(name string, value any) adapters.Select {
	filter := expression.Name(name).GreaterThan(expression.Value(value))
	s.conditions = append(s.conditions, filter)
	return s
}

func (s *Select) GreaterThanEqual(name string, value any) adapters.Select {
	filter := expression.Name(name).GreaterThanEqual(expression.Value(value))
	s.conditions = append(s.conditions, filter)
	return s
}

func (s *Select) LessThan(name string, value any) adapters.Select {
	filter := expression.Name(name).LessThan(expression.Value(value))
	s.conditions = append(s.conditions, filter)
	return s
}

func (s *Select) LessThanEqual(name string, value any) adapters.Select {
	filter := expression.Name(name).LessThanEqual(expression.Value(value))
	s.conditions = append(s.conditions, filter)
	return s
}

func (s *Select) In(name string, value any) adapters.Select {
	filter := expression.Name(name).In(expression.Value(value))
	s.conditions = append(s.conditions, filter)
	return s
}

func (s *Select) NotEqual(name string, value any) adapters.Select {
	filter := expression.Name(name).NotEqual(expression.Value(value))
	s.conditions = append(s.conditions, filter)
	return s
}

func (s *Select) Count(ctx context.Context) (int64, error) {
	if s.model == nil {
		return -1, fmt.Errorf(".Model(model any) is required")
	}
	output, err := s.scan(ctx)
	if err != nil {
		return -1, fmt.Errorf("error in scan: %v", err)
	}
	return int64(output.Count), nil
}

func (s *Select) OrderBy(name string, order string) adapters.Select {
	fmt.Println("\033[1;33m[WARN]\033[0m DynamoDB not implements OrderBy")
	return s
}

func (s *Select) Exec(ctx context.Context) error {
	if s.model == nil {
		return fmt.Errorf(".Model(model any) is required")
	}
	output, err := s.scan(ctx)

	if err != nil {
		return fmt.Errorf("error in Scan table %v , %v", s.tableName, err)
	}
	if err := attributevalue.UnmarshalListOfMaps(output.Items, s.model); err != nil {
		return fmt.Errorf("error to UnmarshlListOfMaps %v", err)
	}
	return nil
}

func (s *Select) scan(ctx context.Context) (*dynamodb.ScanOutput, error) {
	var filter expression.ConditionBuilder
	for i, condition := range s.conditions {
		if i == 0 {
			filter = condition
			continue
		}
		filter.And(condition)
	}

	expr, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return nil, fmt.Errorf("error to create conditions %v", err)
	}

	output, err := s.adapter.client.Scan(ctx, &dynamodb.ScanInput{
		TableName:                 aws.String(s.tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	})
	return output, err
}
