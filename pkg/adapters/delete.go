package adapters

import "context"

type Delete interface {
	Model(model any) Delete
	Condition(name string, value any) Delete
	Exec(ctx context.Context) error
}
