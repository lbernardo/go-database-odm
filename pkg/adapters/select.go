package adapters

import "context"

type Select interface {
	Model(model any) Select
	Condition(name string, value any) Select
	Count(ctx context.Context) (int64, error)
	OrderBy(name string, order string) Select
	Exec(ctx context.Context) error

	// GreaterThan Operator where the value of the field is greater than (>)
	GreaterThan(name string, value any) Select
	// GreaterThanEqual Operator where the value of the field is greater than or equal (>=)
	GreaterThanEqual(name string, value any) Select
	// LessThan Operator where the value of the field is less than (<)
	LessThan(name string, value any) Select
	// LessThanEqual Operator where the value of the field is less than or equal (<=)
	LessThanEqual(name string, value any) Select
	// In Operator where the value of the field equals any value the specified list
	In(name string, value any) Select
	// NotEqual Operator where the value of the field is not equal
	NotEqual(name string, value any) Select
}
