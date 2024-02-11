// Code generated by ent, DO NOT EDIT.

package server

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the server type in the database.
	Label = "server"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldRAM holds the string denoting the ram field in the database.
	FieldRAM = "ram"
	// FieldStorage holds the string denoting the storage field in the database.
	FieldStorage = "storage"
	// FieldLogicalCores holds the string denoting the logical_cores field in the database.
	FieldLogicalCores = "logical_cores"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// EdgeNode holds the string denoting the node edge name in mutations.
	EdgeNode = "node"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// Table holds the table name of the server in the database.
	Table = "servers"
	// NodeTable is the table that holds the node relation/edge.
	NodeTable = "servers"
	// NodeInverseTable is the table name for the Node entity.
	// It exists in this package in order to avoid circular dependency with the "node" package.
	NodeInverseTable = "nodes"
	// NodeColumn is the table column denoting the node relation/edge.
	NodeColumn = "node_nodes"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "servers"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_owners"
	// GameTable is the table that holds the game relation/edge.
	GameTable = "servers"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// GameColumn is the table column denoting the game relation/edge.
	GameColumn = "game_games"
)

// Columns holds all SQL columns for server fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldRAM,
	FieldStorage,
	FieldLogicalCores,
	FieldPort,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "servers"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"game_games",
	"node_nodes",
	"user_owners",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Server queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByRAM orders the results by the ram field.
func ByRAM(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRAM, opts...).ToFunc()
}

// ByStorage orders the results by the storage field.
func ByStorage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStorage, opts...).ToFunc()
}

// ByLogicalCores orders the results by the logical_cores field.
func ByLogicalCores(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLogicalCores, opts...).ToFunc()
}

// ByPort orders the results by the port field.
func ByPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPort, opts...).ToFunc()
}

// ByNodeField orders the results by node field.
func ByNodeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNodeStep(), sql.OrderByField(field, opts...))
	}
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByGameField orders the results by game field.
func ByGameField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGameStep(), sql.OrderByField(field, opts...))
	}
}
func newNodeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NodeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
	)
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newGameStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GameInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GameTable, GameColumn),
	)
}