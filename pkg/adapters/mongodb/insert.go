package mongodb

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
	"github.com/lbernardo/go-database-odm/pkg/utils/models"
)

type InsertMongodb struct {
	adapter   *MongodbAdapter
	model     any
	tableName string
}

func NewInsertMongodb(adapter *MongodbAdapter) *InsertMongodb {
	return &InsertMongodb{
		adapter: adapter,
	}
}

func (i *InsertMongodb) Model(model any) adapters.Insert {
	i.model = model
	i.tableName = models.GetTableNameByModel(model)
	return i
}

func (i *InsertMongodb) Exec(ctx context.Context) (*adapters.InsertResult, error) {
	if i.model == nil {
		return nil, fmt.Errorf(".Model(model any) is required")
	}
	res, err := i.adapter.database.Collection(i.tableName).InsertOne(ctx, i.model)
	if err != nil {
		return nil, err
	}
	return &adapters.InsertResult{
		Result: res.InsertedID,
	}, nil
}
