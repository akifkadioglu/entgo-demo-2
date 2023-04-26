package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// TableExample holds the schema definition for the TableExample entity.
type TableExample struct {
	ent.Schema
}

// Fields of the Table.
func (TableExample) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("field1", []FieldStruct{}),
		field.JSON("field2", FieldStruct{}),
	}
}
type FieldStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Edges of the TableExample.
func (TableExample) Edges() []ent.Edge {
	return nil
}
