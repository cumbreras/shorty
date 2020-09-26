// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/cumbreras/shortener/ent/shortenurl"
	"github.com/google/uuid"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeShortenURL = "ShortenURL"
)

// ShortenURLMutation represents an operation that mutate the ShortenURLs
// nodes in the graph.
type ShortenURLMutation struct {
	config
	op            Op
	typ           string
	id            *int
	_URL          *string
	_Code         *uuid.UUID
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ShortenURL, error)
}

var _ ent.Mutation = (*ShortenURLMutation)(nil)

// shortenurlOption allows to manage the mutation configuration using functional options.
type shortenurlOption func(*ShortenURLMutation)

// newShortenURLMutation creates new mutation for $n.Name.
func newShortenURLMutation(c config, op Op, opts ...shortenurlOption) *ShortenURLMutation {
	m := &ShortenURLMutation{
		config:        c,
		op:            op,
		typ:           TypeShortenURL,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withShortenURLID sets the id field of the mutation.
func withShortenURLID(id int) shortenurlOption {
	return func(m *ShortenURLMutation) {
		var (
			err   error
			once  sync.Once
			value *ShortenURL
		)
		m.oldValue = func(ctx context.Context) (*ShortenURL, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ShortenURL.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withShortenURL sets the old ShortenURL of the mutation.
func withShortenURL(node *ShortenURL) shortenurlOption {
	return func(m *ShortenURLMutation) {
		m.oldValue = func(context.Context) (*ShortenURL, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ShortenURLMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ShortenURLMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *ShortenURLMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetURL sets the URL field.
func (m *ShortenURLMutation) SetURL(s string) {
	m._URL = &s
}

// URL returns the URL value in the mutation.
func (m *ShortenURLMutation) URL() (r string, exists bool) {
	v := m._URL
	if v == nil {
		return
	}
	return *v, true
}

// OldURL returns the old URL value of the ShortenURL.
// If the ShortenURL object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *ShortenURLMutation) OldURL(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldURL is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURL: %w", err)
	}
	return oldValue.URL, nil
}

// ResetURL reset all changes of the "URL" field.
func (m *ShortenURLMutation) ResetURL() {
	m._URL = nil
}

// SetCode sets the Code field.
func (m *ShortenURLMutation) SetCode(u uuid.UUID) {
	m._Code = &u
}

// Code returns the Code value in the mutation.
func (m *ShortenURLMutation) Code() (r uuid.UUID, exists bool) {
	v := m._Code
	if v == nil {
		return
	}
	return *v, true
}

// OldCode returns the old Code value of the ShortenURL.
// If the ShortenURL object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *ShortenURLMutation) OldCode(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCode is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCode: %w", err)
	}
	return oldValue.Code, nil
}

// ResetCode reset all changes of the "Code" field.
func (m *ShortenURLMutation) ResetCode() {
	m._Code = nil
}

// Op returns the operation name.
func (m *ShortenURLMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (ShortenURL).
func (m *ShortenURLMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *ShortenURLMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m._URL != nil {
		fields = append(fields, shortenurl.FieldURL)
	}
	if m._Code != nil {
		fields = append(fields, shortenurl.FieldCode)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *ShortenURLMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case shortenurl.FieldURL:
		return m.URL()
	case shortenurl.FieldCode:
		return m.Code()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *ShortenURLMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case shortenurl.FieldURL:
		return m.OldURL(ctx)
	case shortenurl.FieldCode:
		return m.OldCode(ctx)
	}
	return nil, fmt.Errorf("unknown ShortenURL field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *ShortenURLMutation) SetField(name string, value ent.Value) error {
	switch name {
	case shortenurl.FieldURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURL(v)
		return nil
	case shortenurl.FieldCode:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCode(v)
		return nil
	}
	return fmt.Errorf("unknown ShortenURL field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *ShortenURLMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *ShortenURLMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *ShortenURLMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ShortenURL numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *ShortenURLMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *ShortenURLMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *ShortenURLMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ShortenURL nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *ShortenURLMutation) ResetField(name string) error {
	switch name {
	case shortenurl.FieldURL:
		m.ResetURL()
		return nil
	case shortenurl.FieldCode:
		m.ResetCode()
		return nil
	}
	return fmt.Errorf("unknown ShortenURL field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *ShortenURLMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *ShortenURLMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *ShortenURLMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *ShortenURLMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *ShortenURLMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *ShortenURLMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *ShortenURLMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ShortenURL unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *ShortenURLMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ShortenURL edge %s", name)
}
