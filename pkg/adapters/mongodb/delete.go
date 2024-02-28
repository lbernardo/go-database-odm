package mongodb

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
	"github.com/lbernardo/go-database-odm/pkg/utils/models"
	"go.mongodb.org/mongo-driver/bson"
)

type DeleteMongodb struct {
	adapter    *MongodbAdapter
	conditions bson.M
	tableName  string
	model      any
}

func NewDeleteMongodb(adapter *MongodbAdapter) *DeleteMongodb {
	return &DeleteMongodb{adapter: adapter, conditions: bson.M{}}
}

func (d *DeleteMongodb) Model(model any) adapters.Delete {
	d.model = model
	d.tableName = models.GetTableNameByModel(model)
	return d
}

func (d *DeleteMongodb) Condition(name string, value any) adapters.Delete {
	d.conditions[name] = value
	return d
}

func (d *DeleteMongodb) Exec(ctx context.Context) error {
	if d.model == nil {
		return fmt.Errorf(".Model(model any) is required")
	}
	_, err := d.adapter.database.Collection(d.tableName).DeleteOne(ctx, d.conditions)
	if err != nil {
		return fmt.Errorf("error to deleteOne: %v", err)
	}
	return nil
}
