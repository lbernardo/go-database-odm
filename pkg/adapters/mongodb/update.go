package mongodb

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
	"github.com/lbernardo/go-database-odm/pkg/utils/models"
	"go.mongodb.org/mongo-driver/bson"
)

type UpdateMongodb struct {
	adapter    *MongodbAdapter
	model      any
	tableName  string
	conditions bson.M
	setters    bson.M
}

func NewUpdateMongodb(adapter *MongodbAdapter) *UpdateMongodb {
	return &UpdateMongodb{adapter: adapter, conditions: bson.M{}}
}

func (u *UpdateMongodb) Model(model any) adapters.Update {
	u.model = model
	u.tableName = models.GetTableNameByModel(model)
	return u
}

func (u *UpdateMongodb) Condition(name string, value any) adapters.Update {
	u.conditions[name] = value
	return u
}

func (u *UpdateMongodb) Set(name string, value any) adapters.Update {
	if u.setters == nil {
		u.setters = bson.M{}
	}
	u.setters[name] = value
	return u
}

func (u *UpdateMongodb) Exec(ctx context.Context) error {
	if u.model == nil {
		return fmt.Errorf(".Model(model any) is required")
	}

	var setters any = u.setters
	if setters == nil {
		setters = u.model
	}

	_, err := u.adapter.database.Collection(u.tableName).UpdateOne(ctx, u.conditions, bson.M{"$set": setters})
	return err
}
