package adapters

import "context"

type Update interface {
	Model(model any) Update
	Condition(name string, value any) Update
	Set(name string, value any) Update
	Exec(ctx context.Context) error
}
