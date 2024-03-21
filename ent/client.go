// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Encedeus/panel/ent/migrate"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Encedeus/panel/ent/apikey"
	"github.com/Encedeus/panel/ent/game"
	"github.com/Encedeus/panel/ent/node"
	"github.com/Encedeus/panel/ent/role"
	"github.com/Encedeus/panel/ent/server"
	"github.com/Encedeus/panel/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ApiKey is the client for interacting with the ApiKey builders.
	ApiKey *ApiKeyClient
	// Game is the client for interacting with the Game builders.
	Game *GameClient
	// Node is the client for interacting with the Node builders.
	Node *NodeClient
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// Server is the client for interacting with the Server builders.
	Server *ServerClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ApiKey = NewApiKeyClient(c.config)
	c.Game = NewGameClient(c.config)
	c.Node = NewNodeClient(c.config)
	c.Role = NewRoleClient(c.config)
	c.Server = NewServerClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		ApiKey: NewApiKeyClient(cfg),
		Game:   NewGameClient(cfg),
		Node:   NewNodeClient(cfg),
		Role:   NewRoleClient(cfg),
		Server: NewServerClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		ApiKey: NewApiKeyClient(cfg),
		Game:   NewGameClient(cfg),
		Node:   NewNodeClient(cfg),
		Role:   NewRoleClient(cfg),
		Server: NewServerClient(cfg),
		User:   NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ApiKey.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	for _, n := range []interface{ Use(...Hook) }{
		c.ApiKey, c.Game, c.Node, c.Role, c.Server, c.User,
	} {
		n.Use(hooks...)
	}
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	for _, n := range []interface{ Intercept(...Interceptor) }{
		c.ApiKey, c.Game, c.Node, c.Role, c.Server, c.User,
	} {
		n.Intercept(interceptors...)
	}
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ApiKeyMutation:
		return c.ApiKey.mutate(ctx, m)
	case *GameMutation:
		return c.Game.mutate(ctx, m)
	case *NodeMutation:
		return c.Node.mutate(ctx, m)
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *ServerMutation:
		return c.Server.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ApiKeyClient is a client for the ApiKey schema.
type ApiKeyClient struct {
	config
}

// NewApiKeyClient returns a client for the ApiKey from the given config.
func NewApiKeyClient(c config) *ApiKeyClient {
	return &ApiKeyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `apikey.Hooks(f(g(h())))`.
func (c *ApiKeyClient) Use(hooks ...Hook) {
	c.hooks.ApiKey = append(c.hooks.ApiKey, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `apikey.Intercept(f(g(h())))`.
func (c *ApiKeyClient) Intercept(interceptors ...Interceptor) {
	c.inters.ApiKey = append(c.inters.ApiKey, interceptors...)
}

// Create returns a builder for creating a ApiKey entity.
func (c *ApiKeyClient) Create() *ApiKeyCreate {
	mutation := newApiKeyMutation(c.config, OpCreate)
	return &ApiKeyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ApiKey entities.
func (c *ApiKeyClient) CreateBulk(builders ...*ApiKeyCreate) *ApiKeyCreateBulk {
	return &ApiKeyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ApiKey.
func (c *ApiKeyClient) Update() *ApiKeyUpdate {
	mutation := newApiKeyMutation(c.config, OpUpdate)
	return &ApiKeyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ApiKeyClient) UpdateOne(ak *ApiKey) *ApiKeyUpdateOne {
	mutation := newApiKeyMutation(c.config, OpUpdateOne, withApiKey(ak))
	return &ApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ApiKeyClient) UpdateOneID(id uuid.UUID) *ApiKeyUpdateOne {
	mutation := newApiKeyMutation(c.config, OpUpdateOne, withApiKeyID(id))
	return &ApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ApiKey.
func (c *ApiKeyClient) Delete() *ApiKeyDelete {
	mutation := newApiKeyMutation(c.config, OpDelete)
	return &ApiKeyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ApiKeyClient) DeleteOne(ak *ApiKey) *ApiKeyDeleteOne {
	return c.DeleteOneID(ak.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ApiKeyClient) DeleteOneID(id uuid.UUID) *ApiKeyDeleteOne {
	builder := c.Delete().Where(apikey.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ApiKeyDeleteOne{builder}
}

// Query returns a query builder for ApiKey.
func (c *ApiKeyClient) Query() *ApiKeyQuery {
	return &ApiKeyQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeApiKey},
		inters: c.Interceptors(),
	}
}

// Get returns a ApiKey entity by its id.
func (c *ApiKeyClient) Get(ctx context.Context, id uuid.UUID) (*ApiKey, error) {
	return c.Query().Where(apikey.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ApiKeyClient) GetX(ctx context.Context, id uuid.UUID) *ApiKey {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a ApiKey.
func (c *ApiKeyClient) QueryUser(ak *ApiKey) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ak.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(apikey.Table, apikey.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, apikey.UserTable, apikey.UserColumn),
		)
		fromV = sqlgraph.Neighbors(ak.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ApiKeyClient) Hooks() []Hook {
	return c.hooks.ApiKey
}

// Interceptors returns the client interceptors.
func (c *ApiKeyClient) Interceptors() []Interceptor {
	return c.inters.ApiKey
}

func (c *ApiKeyClient) mutate(ctx context.Context, m *ApiKeyMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ApiKeyCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ApiKeyUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ApiKeyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ApiKeyDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ApiKey mutation op: %q", m.Op())
	}
}

// GameClient is a client for the Game schema.
type GameClient struct {
	config
}

// NewGameClient returns a client for the Game from the given config.
func NewGameClient(c config) *GameClient {
	return &GameClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `game.Hooks(f(g(h())))`.
func (c *GameClient) Use(hooks ...Hook) {
	c.hooks.Game = append(c.hooks.Game, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `game.Intercept(f(g(h())))`.
func (c *GameClient) Intercept(interceptors ...Interceptor) {
	c.inters.Game = append(c.inters.Game, interceptors...)
}

// Create returns a builder for creating a Game entity.
func (c *GameClient) Create() *GameCreate {
	mutation := newGameMutation(c.config, OpCreate)
	return &GameCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Game entities.
func (c *GameClient) CreateBulk(builders ...*GameCreate) *GameCreateBulk {
	return &GameCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Game.
func (c *GameClient) Update() *GameUpdate {
	mutation := newGameMutation(c.config, OpUpdate)
	return &GameUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GameClient) UpdateOne(ga *Game) *GameUpdateOne {
	mutation := newGameMutation(c.config, OpUpdateOne, withGame(ga))
	return &GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GameClient) UpdateOneID(id uuid.UUID) *GameUpdateOne {
	mutation := newGameMutation(c.config, OpUpdateOne, withGameID(id))
	return &GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Game.
func (c *GameClient) Delete() *GameDelete {
	mutation := newGameMutation(c.config, OpDelete)
	return &GameDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *GameClient) DeleteOne(ga *Game) *GameDeleteOne {
	return c.DeleteOneID(ga.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *GameClient) DeleteOneID(id uuid.UUID) *GameDeleteOne {
	builder := c.Delete().Where(game.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GameDeleteOne{builder}
}

// Query returns a query builder for Game.
func (c *GameClient) Query() *GameQuery {
	return &GameQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeGame},
		inters: c.Interceptors(),
	}
}

// Get returns a Game entity by its id.
func (c *GameClient) Get(ctx context.Context, id uuid.UUID) (*Game, error) {
	return c.Query().Where(game.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GameClient) GetX(ctx context.Context, id uuid.UUID) *Game {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGames queries the games edge of a Game.
func (c *GameClient) QueryGames(ga *Game) *ServerQuery {
	query := (&ServerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ga.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(game.Table, game.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, game.GamesTable, game.GamesColumn),
		)
		fromV = sqlgraph.Neighbors(ga.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *GameClient) Hooks() []Hook {
	return c.hooks.Game
}

// Interceptors returns the client interceptors.
func (c *GameClient) Interceptors() []Interceptor {
	return c.inters.Game
}

func (c *GameClient) mutate(ctx context.Context, m *GameMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&GameCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&GameUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&GameUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&GameDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Game mutation op: %q", m.Op())
	}
}

// NodeClient is a client for the Node schema.
type NodeClient struct {
	config
}

// NewNodeClient returns a client for the Node from the given config.
func NewNodeClient(c config) *NodeClient {
	return &NodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `node.Hooks(f(g(h())))`.
func (c *NodeClient) Use(hooks ...Hook) {
	c.hooks.Node = append(c.hooks.Node, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `node.Intercept(f(g(h())))`.
func (c *NodeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Node = append(c.inters.Node, interceptors...)
}

// Create returns a builder for creating a Node entity.
func (c *NodeClient) Create() *NodeCreate {
	mutation := newNodeMutation(c.config, OpCreate)
	return &NodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Node entities.
func (c *NodeClient) CreateBulk(builders ...*NodeCreate) *NodeCreateBulk {
	return &NodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Node.
func (c *NodeClient) Update() *NodeUpdate {
	mutation := newNodeMutation(c.config, OpUpdate)
	return &NodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NodeClient) UpdateOne(n *Node) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNode(n))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NodeClient) UpdateOneID(id uuid.UUID) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNodeID(id))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Node.
func (c *NodeClient) Delete() *NodeDelete {
	mutation := newNodeMutation(c.config, OpDelete)
	return &NodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NodeClient) DeleteOne(n *Node) *NodeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *NodeClient) DeleteOneID(id uuid.UUID) *NodeDeleteOne {
	builder := c.Delete().Where(node.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NodeDeleteOne{builder}
}

// Query returns a query builder for Node.
func (c *NodeClient) Query() *NodeQuery {
	return &NodeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeNode},
		inters: c.Interceptors(),
	}
}

// Get returns a Node entity by its id.
func (c *NodeClient) Get(ctx context.Context, id uuid.UUID) (*Node, error) {
	return c.Query().Where(node.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NodeClient) GetX(ctx context.Context, id uuid.UUID) *Node {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryNodes queries the nodes edge of a Node.
func (c *NodeClient) QueryNodes(n *Node) *ServerQuery {
	query := (&ServerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, node.NodesTable, node.NodesColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NodeClient) Hooks() []Hook {
	return c.hooks.Node
}

// Interceptors returns the client interceptors.
func (c *NodeClient) Interceptors() []Interceptor {
	return c.inters.Node
}

func (c *NodeClient) mutate(ctx context.Context, m *NodeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&NodeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&NodeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&NodeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Node mutation op: %q", m.Op())
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id uuid.UUID) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id uuid.UUID) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id uuid.UUID) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id uuid.UUID) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
}

// ServerClient is a client for the Server schema.
type ServerClient struct {
	config
}

// NewServerClient returns a client for the Server from the given config.
func NewServerClient(c config) *ServerClient {
	return &ServerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `server.Hooks(f(g(h())))`.
func (c *ServerClient) Use(hooks ...Hook) {
	c.hooks.Server = append(c.hooks.Server, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `server.Intercept(f(g(h())))`.
func (c *ServerClient) Intercept(interceptors ...Interceptor) {
	c.inters.Server = append(c.inters.Server, interceptors...)
}

// Create returns a builder for creating a Server entity.
func (c *ServerClient) Create() *ServerCreate {
	mutation := newServerMutation(c.config, OpCreate)
	return &ServerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Server entities.
func (c *ServerClient) CreateBulk(builders ...*ServerCreate) *ServerCreateBulk {
	return &ServerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Server.
func (c *ServerClient) Update() *ServerUpdate {
	mutation := newServerMutation(c.config, OpUpdate)
	return &ServerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ServerClient) UpdateOne(s *Server) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServer(s))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ServerClient) UpdateOneID(id uuid.UUID) *ServerUpdateOne {
	mutation := newServerMutation(c.config, OpUpdateOne, withServerID(id))
	return &ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Server.
func (c *ServerClient) Delete() *ServerDelete {
	mutation := newServerMutation(c.config, OpDelete)
	return &ServerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ServerClient) DeleteOne(s *Server) *ServerDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ServerClient) DeleteOneID(id uuid.UUID) *ServerDeleteOne {
	builder := c.Delete().Where(server.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ServerDeleteOne{builder}
}

// Query returns a query builder for Server.
func (c *ServerClient) Query() *ServerQuery {
	return &ServerQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeServer},
		inters: c.Interceptors(),
	}
}

// Get returns a Server entity by its id.
func (c *ServerClient) Get(ctx context.Context, id uuid.UUID) (*Server, error) {
	return c.Query().Where(server.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ServerClient) GetX(ctx context.Context, id uuid.UUID) *Server {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryNode queries the node edge of a Server.
func (c *ServerClient) QueryNode(s *Server) *NodeQuery {
	query := (&NodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, server.NodeTable, server.NodeColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwner queries the owner edge of a Server.
func (c *ServerClient) QueryOwner(s *Server) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(server.Table, server.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, server.OwnerTable, server.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ServerClient) Hooks() []Hook {
	return c.hooks.Server
}

// Interceptors returns the client interceptors.
func (c *ServerClient) Interceptors() []Interceptor {
	return c.inters.Server
}

func (c *ServerClient) mutate(ctx context.Context, m *ServerMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ServerCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ServerUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ServerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ServerDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Server mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRole queries the role edge of a User.
func (c *UserClient) QueryRole(u *User) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, user.RoleTable, user.RoleColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryOwners queries the owners edge of a User.
func (c *UserClient) QueryOwners(u *User) *ServerQuery {
	query := (&ServerClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.OwnersTable, user.OwnersColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ApiKey, Game, Node, Role, Server, User []ent.Hook
	}
	inters struct {
		ApiKey, Game, Node, Role, Server, User []ent.Interceptor
	}
)
