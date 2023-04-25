// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/akifkadioglu/entgo-demo-2/ent/predicate"
	"github.com/akifkadioglu/entgo-demo-2/ent/schema"
	"github.com/akifkadioglu/entgo-demo-2/ent/tableexample"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTableExample = "TableExample"
)

// TableExampleMutation represents an operation that mutates the TableExample nodes in the graph.
type TableExampleMutation struct {
	config
	op            Op
	typ           string
	id            *int
	field1        *schema.FieldStruct
	field2        *[]schema.FieldStruct
	appendfield2  []schema.FieldStruct
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*TableExample, error)
	predicates    []predicate.TableExample
}

var _ ent.Mutation = (*TableExampleMutation)(nil)

// tableexampleOption allows management of the mutation configuration using functional options.
type tableexampleOption func(*TableExampleMutation)

// newTableExampleMutation creates new mutation for the TableExample entity.
func newTableExampleMutation(c config, op Op, opts ...tableexampleOption) *TableExampleMutation {
	m := &TableExampleMutation{
		config:        c,
		op:            op,
		typ:           TypeTableExample,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTableExampleID sets the ID field of the mutation.
func withTableExampleID(id int) tableexampleOption {
	return func(m *TableExampleMutation) {
		var (
			err   error
			once  sync.Once
			value *TableExample
		)
		m.oldValue = func(ctx context.Context) (*TableExample, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().TableExample.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTableExample sets the old TableExample of the mutation.
func withTableExample(node *TableExample) tableexampleOption {
	return func(m *TableExampleMutation) {
		m.oldValue = func(context.Context) (*TableExample, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TableExampleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TableExampleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TableExampleMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TableExampleMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().TableExample.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetField1 sets the "field1" field.
func (m *TableExampleMutation) SetField1(ss schema.FieldStruct) {
	m.field1 = &ss
}

// Field1 returns the value of the "field1" field in the mutation.
func (m *TableExampleMutation) Field1() (r schema.FieldStruct, exists bool) {
	v := m.field1
	if v == nil {
		return
	}
	return *v, true
}

// OldField1 returns the old "field1" field's value of the TableExample entity.
// If the TableExample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TableExampleMutation) OldField1(ctx context.Context) (v schema.FieldStruct, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldField1 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldField1 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldField1: %w", err)
	}
	return oldValue.Field1, nil
}

// ResetField1 resets all changes to the "field1" field.
func (m *TableExampleMutation) ResetField1() {
	m.field1 = nil
}

// SetField2 sets the "field2" field.
func (m *TableExampleMutation) SetField2(ss []schema.FieldStruct) {
	m.field2 = &ss
	m.appendfield2 = nil
}

// Field2 returns the value of the "field2" field in the mutation.
func (m *TableExampleMutation) Field2() (r []schema.FieldStruct, exists bool) {
	v := m.field2
	if v == nil {
		return
	}
	return *v, true
}

// OldField2 returns the old "field2" field's value of the TableExample entity.
// If the TableExample object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TableExampleMutation) OldField2(ctx context.Context) (v []schema.FieldStruct, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldField2 is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldField2 requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldField2: %w", err)
	}
	return oldValue.Field2, nil
}

// AppendField2 adds ss to the "field2" field.
func (m *TableExampleMutation) AppendField2(ss []schema.FieldStruct) {
	m.appendfield2 = append(m.appendfield2, ss...)
}

// AppendedField2 returns the list of values that were appended to the "field2" field in this mutation.
func (m *TableExampleMutation) AppendedField2() ([]schema.FieldStruct, bool) {
	if len(m.appendfield2) == 0 {
		return nil, false
	}
	return m.appendfield2, true
}

// ResetField2 resets all changes to the "field2" field.
func (m *TableExampleMutation) ResetField2() {
	m.field2 = nil
	m.appendfield2 = nil
}

// Where appends a list predicates to the TableExampleMutation builder.
func (m *TableExampleMutation) Where(ps ...predicate.TableExample) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TableExampleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TableExampleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.TableExample, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TableExampleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TableExampleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (TableExample).
func (m *TableExampleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TableExampleMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.field1 != nil {
		fields = append(fields, tableexample.FieldField1)
	}
	if m.field2 != nil {
		fields = append(fields, tableexample.FieldField2)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TableExampleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case tableexample.FieldField1:
		return m.Field1()
	case tableexample.FieldField2:
		return m.Field2()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TableExampleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case tableexample.FieldField1:
		return m.OldField1(ctx)
	case tableexample.FieldField2:
		return m.OldField2(ctx)
	}
	return nil, fmt.Errorf("unknown TableExample field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TableExampleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case tableexample.FieldField1:
		v, ok := value.(schema.FieldStruct)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetField1(v)
		return nil
	case tableexample.FieldField2:
		v, ok := value.([]schema.FieldStruct)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetField2(v)
		return nil
	}
	return fmt.Errorf("unknown TableExample field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TableExampleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TableExampleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TableExampleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown TableExample numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TableExampleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TableExampleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TableExampleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown TableExample nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TableExampleMutation) ResetField(name string) error {
	switch name {
	case tableexample.FieldField1:
		m.ResetField1()
		return nil
	case tableexample.FieldField2:
		m.ResetField2()
		return nil
	}
	return fmt.Errorf("unknown TableExample field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TableExampleMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TableExampleMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TableExampleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TableExampleMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TableExampleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TableExampleMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TableExampleMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown TableExample unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TableExampleMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown TableExample edge %s", name)
}
