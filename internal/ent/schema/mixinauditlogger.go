package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type MixinAuditLogger struct {
	mixin.Schema
}

// Fields of the MixinAuditLogger.
func (MixinAuditLogger) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.UUID("created_by", uuid.UUID{}).
			Optional().Nillable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).Nillable(),
		field.UUID("updated_by", uuid.UUID{}).
			Optional().Nillable(),
	}
}

// Hooks of the MixinAuditLogger.
func (MixinAuditLogger) Hooks() []ent.Hook {
	return []ent.Hook{
		AuditHook,
	}
}

// A AuditHook is an example for audit-log hook.
func AuditHook(next ent.Mutator) ent.Mutator {
	// AuditLogger wraps the methods that are shared between all mutations of
	// schemas that embed the AuditLog mixin. The variable "exists" is true, if
	// the field already exists in the mutation (e.g. was set by a different hook).
	type AuditLogger interface {
		SetCreatedAt(time.Time)
		CreatedAt() (value time.Time, exists bool)
		SetCreatedBy(uuid.UUID)
		CreatedBy() (id uuid.UUID, exists bool)
		SetUpdatedAt(time.Time)
		UpdatedAt() (value time.Time, exists bool)
		SetUpdatedBy(uuid.UUID)
		UpdatedBy() (id uuid.UUID, exists bool)
	}
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		ml, ok := m.(AuditLogger)
		if !ok {
			return nil, fmt.Errorf("unexpected audit-log call from mutation type %T", m)
		}

		user_id := ctx.Value("user_id")
		var uuidUser uuid.UUID
		if user_id != nil {
			uuidUsertemp, errr := uuid.Parse(user_id.(string))
			if errr != nil {
				return nil, errr
			}
			uuidUser = uuidUsertemp
		}

		if ok {
			// return nil, fmt.Errorf("user is not logged in") // if not
			switch op := m.Op(); {
			case op.Is(ent.OpCreate):
				ml.SetCreatedAt(time.Now())
				if _, exists := ml.CreatedBy(); !exists {
					if uuidUser != uuid.Nil {
						ml.SetCreatedBy(uuidUser)
					}
				}
			case op.Is(ent.OpUpdateOne | ent.OpUpdate):
				ml.SetUpdatedAt(time.Now())
				if _, exists := ml.UpdatedBy(); !exists {
					if uuidUser != uuid.Nil {
						ml.SetUpdatedBy(uuidUser)
					}
				}
			}
		}

		return next.Mutate(ctx, m)
	})
}
