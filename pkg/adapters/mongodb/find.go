package mongodb

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
	"github.com/lbernardo/go-database-odm/pkg/utils/models"
	"go.mongodb.org/mongo-driver/bson"
)

type FindMongodb struct {
	model      any
	tableName  string
	adapter    *MongodbAdapter
	conditions bson.M
}

func NewFindMongodb(a *MongodbAdapter) *FindMongodb {
	return &FindMongodb{
		adapter:    a,
		conditions: bson.M{},
	}
}

func (f *FindMongodb) Model(model any) adapters.Find {
	f.model = model
	f.tableName = models.GetTableNameByModel(model)
	return f
}

func (f *FindMongodb) Condition(name string, value any) adapters.Find {
	f.conditions[name] = value
	return f
}

func (f *FindMongodb) Exec(ctx context.Context) error {
	if f.model == nil {
		return fmt.Errorf(".Model(model any) is required")
	}
	result := f.adapter.database.Collection(f.tableName).FindOne(ctx, f.conditions)
	if err := result.Decode(f.model); err != nil {
		return err
	}
	return nil
}
