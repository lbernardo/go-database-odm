package mongodb

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbAdapter struct {
	cfg      *Config
	client   *mongo.Client
	database *mongo.Database
}

func NewMongodbAdapter(cfg *Config) (*MongodbAdapter, error) {
	a := &MongodbAdapter{
		cfg: cfg,
	}
	if err := a.connect(); err != nil {
		return nil, err
	}
	return a, nil
}

func (m *MongodbAdapter) connect() error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.cfg.DatabaseUri))
	if err != nil {
		return fmt.Errorf("error to connect mongodb: %v", err)
	}
	db := client.Database(m.cfg.DatabaseName)
	m.client = client
	m.database = db
	return nil
}

func (m *MongodbAdapter) Disconnect() error {
	return m.client.Disconnect(context.TODO())
}

func (m *MongodbAdapter) NewInsert() adapters.Insert {
	return NewInsertMongodb(m)
}

func (m *MongodbAdapter) NewUpdate() adapters.Update {
	return NewUpdateMongodb(m)
}

func (m *MongodbAdapter) NewDelete() adapters.Delete {
	return NewDeleteMongodb(m)
}

func (m *MongodbAdapter) NewSelect() adapters.Select {
	return NewSelect(m)
}

func (m *MongodbAdapter) NewFind() adapters.Find {
	return NewFindMongodb(m)
}

func (m *MongodbAdapter) GetInstance() any {
	return m.database
}
