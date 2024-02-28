package adapters

import "context"

type Find interface {
	Model(model any) Find
	Condition(name string, value any) Find
	Exec(ctx context.Context) error
}
