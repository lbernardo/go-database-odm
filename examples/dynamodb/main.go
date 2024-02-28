package main

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters/dynamodb"
	"log"
	"time"
)

type Client struct {
	_       any `database:"table:clients"`
	Id      string
	Cluster string
	Name    string
}

func main() {

	db, err := dynamodb.NewDynamoAdapter(&dynamodb.Config{
		AwsRegion:   "us-east-1",
		TablePrefix: "backstage-cluster-",
	})
	if err != nil {
		log.Fatalf("error to dynamodbadapter %v", err)
	}
	if _, err := db.NewInsert().Model(&Client{Id: fmt.Sprintf("%v", time.Now().Unix()), Cluster: "teste", Name: "Lucas Bernardo"}).Exec(context.Background()); err != nil {
		log.Fatalf("error to insert %v", err)
	}
	var client Client
	if err := db.NewFind().Model(&client).Condition("Id", "1695644768").Condition("Cluster", "teste").Exec(context.Background()); err != nil {
		log.Fatalf("error to find %v", err)
	}
	fmt.Println(client)

	var listClient []Client
	if err := db.NewSelect().Model(&listClient).Condition("Cluster", "teste").Exec(context.Background()); err != nil {
		log.Fatalf("error to find %v", err)
	}
	fmt.Println(listClient)

	if err := db.NewUpdate().Model(&Client{}).Condition("Cluster", "teste").Condition("Id", "1695644768").Set("Name", "Lucas B S Bernardo").Exec(context.Background()); err != nil {
		log.Fatalf("error to update %v", err)
	}

	if err := db.NewDelete().Model(&Client{}).Condition("Cluster", "teste").Condition("Id", "1695644768").Exec(context.Background()); err != nil {
		log.Fatalf("error to update %v", err)
	}

}
