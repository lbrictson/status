// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/lbrictson/status/ent/operator"
)

// Operator is the model entity for the Operator schema.
type Operator struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
	// HashedPassword holds the value of the "HashedPassword" field.
	HashedPassword string `json:"HashedPassword,omitempty"`
	// Role holds the value of the "Role" field.
	Role string `json:"Role,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Operator) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case operator.FieldID:
			values[i] = new(sql.NullInt64)
		case operator.FieldEmail, operator.FieldHashedPassword, operator.FieldRole:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Operator", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Operator fields.
func (o *Operator) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case operator.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case operator.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Email", values[i])
			} else if value.Valid {
				o.Email = value.String
			}
		case operator.FieldHashedPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field HashedPassword", values[i])
			} else if value.Valid {
				o.HashedPassword = value.String
			}
		case operator.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Role", values[i])
			} else if value.Valid {
				o.Role = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Operator.
// Note that you need to call Operator.Unwrap() before calling this method if this Operator
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Operator) Update() *OperatorUpdateOne {
	return (&OperatorClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Operator entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Operator) Unwrap() *Operator {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Operator is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Operator) String() string {
	var builder strings.Builder
	builder.WriteString("Operator(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", Email=")
	builder.WriteString(o.Email)
	builder.WriteString(", HashedPassword=")
	builder.WriteString(o.HashedPassword)
	builder.WriteString(", Role=")
	builder.WriteString(o.Role)
	builder.WriteByte(')')
	return builder.String()
}

// Operators is a parsable slice of Operator.
type Operators []*Operator

func (o Operators) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
