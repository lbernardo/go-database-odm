package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/lbernardo/go-database-odm/pkg/adapters"
	"github.com/lbernardo/go-database-odm/pkg/utils/models"
)

type DynamoAdapter struct {
	config *Config
	client *dynamodb.Client
}

func NewDynamoAdapter(config *Config) (*DynamoAdapter, error) {
	d := &DynamoAdapter{
		config: config,
	}
	if err := d.connect(); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DynamoAdapter) connect() error {
	region := d.config.AwsRegion
	if region == "" {
		region = "us-east-1"
	}
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion(region))
	if err != nil {
		return fmt.Errorf("error to start aws config %v", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	d.client = client
	return nil
}

func (d *DynamoAdapter) TableName(model any) string {
	table := models.GetTableNameByModel(model)
	if d.config.TablePrefix != "" {
		table = fmt.Sprintf("%v%v", d.config.TablePrefix, table)
	}
	return table
}

func (d *DynamoAdapter) NewInsert() adapters.Insert {
	return NewInsert(d)
}

func (d *DynamoAdapter) NewUpdate() adapters.Update {
	return NewUpdateDynamodb(d)
}

func (d *DynamoAdapter) NewDelete() adapters.Delete {
	return NewDeleteDynamodb(d)
}

func (d *DynamoAdapter) NewSelect() adapters.Select {
	return NewSelect(d)
}

func (d *DynamoAdapter) NewFind() adapters.Find {
	return NewFindDynamodb(d)
}

func (d *DynamoAdapter) GetInstance() any {
	return d.client
}
