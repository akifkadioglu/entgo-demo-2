// Code generated by ent, DO NOT EDIT.

package tableexample

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the tableexample type in the database.
	Label = "table_example"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldField1 holds the string denoting the field1 field in the database.
	FieldField1 = "field1"
	// FieldField2 holds the string denoting the field2 field in the database.
	FieldField2 = "field2"
	// Table holds the table name of the tableexample in the database.
	Table = "table_examples"
)

// Columns holds all SQL columns for tableexample fields.
var Columns = []string{
	FieldID,
	FieldField1,
	FieldField2,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the TableExample queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}
