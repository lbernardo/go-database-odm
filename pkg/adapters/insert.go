package adapters

import "context"

type InsertResult struct {
	Result any `json:"result"`
}

type Insert interface {
	Model(model any) Insert
	Exec(ctx context.Context) (*InsertResult, error)
}
