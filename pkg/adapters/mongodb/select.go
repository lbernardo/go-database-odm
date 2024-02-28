package mongodb

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
	"github.com/lbernardo/go-database-odm/pkg/utils/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type Select struct {
	adapter     *MongodbAdapter
	model       any
	tableName   string
	conditions  bson.M
	findOptions *options.FindOptions
}

func NewSelect(adapter *MongodbAdapter) *Select {
	return &Select{adapter: adapter, conditions: bson.M{}}
}

func (s *Select) Model(model any) adapters.Select {
	s.model = model
	s.tableName = models.GetTableNameByModel(model)
	return s
}

func (s *Select) Condition(name string, value any) adapters.Select {
	s.conditions[name] = value
	return s
}

func (s *Select) GreaterThan(name string, value any) adapters.Select {
	s.Condition(name, bson.M{"$gt": value})
	return s
}

func (s *Select) GreaterThanEqual(name string, value any) adapters.Select {
	s.Condition(name, bson.M{"$gte": value})
	return s
}

func (s *Select) LessThan(name string, value any) adapters.Select {
	s.Condition(name, bson.M{"$lt": value})
	return s
}

func (s *Select) LessThanEqual(name string, value any) adapters.Select {
	s.Condition(name, bson.M{"lte": value})
	return s
}

func (s *Select) In(name string, value any) adapters.Select {
	s.Condition(name, bson.M{"in": value})
	return s
}

func (s *Select) NotEqual(name string, value any) adapters.Select {
	s.Condition(name, bson.M{"ne": value})
	return s
}

func (s *Select) Count(ctx context.Context) (int64, error) {
	if s.model == nil {
		return -1, fmt.Errorf(".Model(model any) is required")
	}
	return s.adapter.database.Collection(s.tableName).CountDocuments(ctx, s.conditions)
}

func (s *Select) OrderBy(name string, order string) adapters.Select {
	orderNum := 1
	if strings.ToLower(order) == "desc" {
		orderNum = -1
	}
	s.findOptions = options.Find().SetSort(bson.M{name: orderNum})
	return s
}

func (s *Select) Exec(ctx context.Context) error {
	if s.model == nil {
		return fmt.Errorf(".Model(model any) is required")
	}
	cursor, err := s.adapter.database.Collection(s.tableName).Find(ctx, s.conditions, s.findOptions)
	if err != nil {
		return err
	}
	if err := cursor.All(ctx, s.model); err != nil {
		return fmt.Errorf("error to decode .Select() %v", err)
	}
	return nil
}
