package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("fullname"),
		field.String("email").Unique().Match(regexp.MustCompile("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$")),
		field.String("phone").Unique().Match(regexp.MustCompile("^\\d{10,12}$")),
		field.String("password"),
		field.Bool("is_email_confirmed").Default(false),
		field.Bool("is_phone_confirmed").Default(false),
		field.String("token_auth").Nillable().Optional(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("email").Unique(),
		index.Fields("phone").Unique(),
		index.Fields("token_auth"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.To("enterprises", Enterprise.Type),
		// edge.To("buy_histories", ShopSell.Type),
		// edge.To("warehouses", Warehouse.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		MixinAuditLogger{},
		MixinSoftDelete{},
		// mixin_proj.SoftDeleteMixin{},
		// mixin_proj.AuditLoggerMixin{},
	}
}
