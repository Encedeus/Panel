package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Server holds the schema definition for the Server entity.
type Server struct {
	ent.Schema
}

// Fields of the Server.
func (Server) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").UpdateDefault(time.Now).Default(time.Now),
		// In KB
		field.Uint("ram"),
		// In KB
		field.Uint("storage"),
		field.Uint("logical_cores"),
		field.Uint16("port"),
	}
}

// Edges of the Server.
func (Server) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("node", Node.Type).
			Ref("nodes").
			Unique(),
		edge.From("owner", User.Type).
			Ref("owners").
			Unique(),
		edge.From("game", Game.Type).
			Ref("games").
			Unique(),
	}
}
