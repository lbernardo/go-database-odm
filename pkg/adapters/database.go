package adapters

type Database interface {
	NewInsert() Insert
	NewUpdate() Update
	NewDelete() Delete
	NewSelect() Select
	NewFind() Find
	GetInstance() any
}
