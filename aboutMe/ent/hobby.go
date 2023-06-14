// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aboutMe/ent/hobby"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Hobby is the model entity for the Hobby schema.
type Hobby struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HobbyQuery when eager-loading is set.
	Edges        HobbyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// HobbyEdges holds the relations/edges for other nodes in the graph.
type HobbyEdges struct {
	// Owners holds the value of the owners edge.
	Owners []*User `json:"owners,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedOwners map[string][]*User
}

// OwnersOrErr returns the Owners value or an error if the edge
// was not loaded in eager-loading.
func (e HobbyEdges) OwnersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Owners, nil
	}
	return nil, &NotLoadedError{edge: "owners"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Hobby) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hobby.FieldID:
			values[i] = new(sql.NullInt64)
		case hobby.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Hobby fields.
func (h *Hobby) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hobby.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			h.ID = int(value.Int64)
		case hobby.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				h.Name = value.String
			}
		default:
			h.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Hobby.
// This includes values selected through modifiers, order, etc.
func (h *Hobby) Value(name string) (ent.Value, error) {
	return h.selectValues.Get(name)
}

// QueryOwners queries the "owners" edge of the Hobby entity.
func (h *Hobby) QueryOwners() *UserQuery {
	return NewHobbyClient(h.config).QueryOwners(h)
}

// Update returns a builder for updating this Hobby.
// Note that you need to call Hobby.Unwrap() before calling this method if this Hobby
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *Hobby) Update() *HobbyUpdateOne {
	return NewHobbyClient(h.config).UpdateOne(h)
}

// Unwrap unwraps the Hobby entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (h *Hobby) Unwrap() *Hobby {
	_tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: Hobby is not a transactional entity")
	}
	h.config.driver = _tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *Hobby) String() string {
	var builder strings.Builder
	builder.WriteString("Hobby(")
	builder.WriteString(fmt.Sprintf("id=%v, ", h.ID))
	builder.WriteString("name=")
	builder.WriteString(h.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedOwners returns the Owners named value or an error if the edge was not
// loaded in eager-loading with this name.
func (h *Hobby) NamedOwners(name string) ([]*User, error) {
	if h.Edges.namedOwners == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := h.Edges.namedOwners[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (h *Hobby) appendNamedOwners(name string, edges ...*User) {
	if h.Edges.namedOwners == nil {
		h.Edges.namedOwners = make(map[string][]*User)
	}
	if len(edges) == 0 {
		h.Edges.namedOwners[name] = []*User{}
	} else {
		h.Edges.namedOwners[name] = append(h.Edges.namedOwners[name], edges...)
	}
}

// Hobbies is a parsable slice of Hobby.
type Hobbies []*Hobby
