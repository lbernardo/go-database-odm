package main

import (
	"context"
	"fmt"
	"github.com/lbernardo/go-database-odm/pkg/adapters/mongodb"
	"log"
)

type Model struct {
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
	if _, err := db.NewInsert().Model(Model{
		Name:  "Lucas Bernardo",
		Email: "example@gmail.com",
	}).Exec(context.Background()); err != nil {
		log.Fatalf("error to execute: %v", err)
	}
	var model Model
	if err := db.NewFind().Model(&model).Condition("name", "Lucas Bernardo").Exec(context.Background()); err != nil {
		log.Fatalf("error to find %v", err)
	}
	fmt.Println(model)

	var modelList = []Model{}
	if err := db.NewSelect().Model(&modelList).Condition("name", "Lucas Bernardo").Exec(context.Background()); err != nil {
		log.Fatalf("error in select %v", err)
	}
	fmt.Println(modelList)

	if err := db.NewUpdate().Model(&Model{}).Set("name", "Bernardinho").Condition("name", "Lucas Bernardo").Exec(context.Background()); err != nil {
		log.Fatalf("error in update %v", err)
	}

	if err := db.NewDelete().Model(&Model{}).Condition("name", "Bernardinho").Exec(context.Background()); err != nil {
		log.Fatalf("error decode %v", err)
	}

	db.NewInsert().Model(Model{
		Name:  "Lucas Bernardo",
		Email: "example@gmail.com",
	}).Exec(context.Background())
	db.NewInsert().Model(Model{
		Name:  "Zacarias Bernardo",
		Email: "example@gmail.com",
	}).Exec(context.Background())
	db.NewInsert().Model(Model{
		Name:  "Andre Bernardo",
		Email: "example@gmail.com",
	}).Exec(context.Background())
	if err := db.NewSelect().Model(&modelList).OrderBy("name", "asc").Exec(context.Background()); err != nil {
		log.Fatalf("error order by")
	}
	fmt.Println(modelList)
}
