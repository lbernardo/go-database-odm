# go-database-odm

Adapters ODM (Object Document Mappers) database clients

## Install
```shell
go get github.com/lbernardo/go-database-odm
```

## Usage


`MongoDB adapter`
```go
db, err := mongodb.NewMongodbAdapter(&mongodb.Config{
		DatabaseName: "database",
		DatabaseUri:  "mongodb://localhost:27017",
})
```

`DynamoDB adapter`
```go
db, err := mongodb.NewMongodbAdapter(&dynamodb.Config{
    AwsRegion:   "us-east-1",
    TablePrefix: "table-dev-", // is optional
})
```


`Usage functions`

```go
package main

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters/mongodb"
	"log"
)

type Model struct {
	// Set table name with tag database, if not defined tableName
	// is model struct name in plural
	_     any    `database:"table:users"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	db, err := mongodb.NewMongodbAdapter(&mongodb.Config{
		DatabaseName: "database",
		DatabaseUri:  "mongodb://localhost:27017",
	})
	defer db.Disconnect()
	if err != nil {
		log.Fatalf("error to connect: %v", err)
	}
	
	// Insert model in table
	if _, err := db.NewInsert().Model(Model{
		Name:  "Lucas Bernardo",
		Email: "example@gmail.com",
	}).Exec(context.Background()); err != nil {
		log.Fatalf("error to execute: %v", err)
	}
	var model Model
	// Find and return in model
	if err := db.NewFind().Model(&model).Condition("name", "Lucas Bernardo").Exec(context.Background()); err != nil {
		log.Fatalf("error to find %v", err)
	}
	fmt.Println(model)

	var modelList = []Model{}
	// Find list in model
	if err := db.NewSelect().Model(&modelList).Condition("name", "Lucas Bernardo").Exec(context.Background()); err != nil {
		log.Fatalf("error in select %v", err)
	}
	fmt.Println(modelList)

	// Update value in table, '.Set(name string, value string)' set values updated 
	if err := db.NewUpdate().Model(&Model{}).Set("name", "Bernardinho").Condition("name", "Lucas Bernardo").Exec(context.Background()); err != nil {
		log.Fatalf("error in update %v", err)
	}

	// Delete value in table
	if err := db.NewDelete().Model(&Model{}).Condition("name", "Bernardinho").Exec(context.Background()); err != nil {
		log.Fatalf("error decode %v", err)
	}
}
```