package models

import "testing"

func TestGetTableNameByModel(t *testing.T) {
	t.Run("struct", func(t *testing.T) {
		type user struct {
		}
		u := user{}
		name := GetTableNameByModel(u)
		if name != "users" {
			t.Errorf("was expected users and return %v", name)
		}
	})
	t.Run("struct with table name", func(t *testing.T) {
		type user struct {
			_ any `database:"table:clients"`
		}
		u := user{}
		name := GetTableNameByModel(u)
		if name != "clients" {
			t.Errorf("was expected users and return %v", name)
		}
	})
	t.Run("ref", func(t *testing.T) {
		type user struct {
		}
		u := &user{}
		name := GetTableNameByModel(u)
		if name != "users" {
			t.Errorf("was expected users and return %v", name)
		}
	})
	t.Run("slice", func(t *testing.T) {
		type user struct {
		}
		u := []user{}
		name := GetTableNameByModel(u)
		if name != "users" {
			t.Errorf("was expected users and return %v", name)
		}
	})
	t.Run("pointer", func(t *testing.T) {
		type user struct {
		}
		var u *user
		u = &user{}
		name := GetTableNameByModel(&u)
		if name != "users" {
			t.Errorf("was expected users and return %v", name)
		}
	})
	t.Run("slice pointer", func(t *testing.T) {
		type user struct {
		}
		name := GetTableNameByModel(&[]user{})
		if name != "users" {
			t.Errorf("was expected users and return %v", name)
		}
	})
}
