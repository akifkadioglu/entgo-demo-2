// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TableExamplesColumns holds the columns for the "table_examples" table.
	TableExamplesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "field1", Type: field.TypeJSON},
		{Name: "field2", Type: field.TypeJSON},
	}
	// TableExamplesTable holds the schema information for the "table_examples" table.
	TableExamplesTable = &schema.Table{
		Name:       "table_examples",
		Columns:    TableExamplesColumns,
		PrimaryKey: []*schema.Column{TableExamplesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TableExamplesTable,
	}
)

func init() {
}
