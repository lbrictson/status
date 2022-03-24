// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OperatorsColumns holds the columns for the "operators" table.
	OperatorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "hashed_password", Type: field.TypeString},
		{Name: "role", Type: field.TypeString},
	}
	// OperatorsTable holds the schema information for the "operators" table.
	OperatorsTable = &schema.Table{
		Name:       "operators",
		Columns:    OperatorsColumns,
		PrimaryKey: []*schema.Column{OperatorsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OperatorsTable,
	}
)

func init() {
}
