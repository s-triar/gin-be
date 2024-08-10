package schema

import (
	"context"
	"fmt"
	gen "gin-be/internal/ent"
	"gin-be/internal/ent/hook"
	"gin-be/internal/ent/intercept"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	// "entgo.io/ent/entc/integration/hooks/ent/intercept"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type UserStruct struct { // User logged In
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

// MixinSoftDelete implements the soft delete pattern for schemas.
type MixinSoftDelete struct {
	mixin.Schema
}

// Fields of the MixinSoftDelete.
func (MixinSoftDelete) Fields() []ent.Field {
	return []ent.Field{
		field.Time("deleted_at").Optional().Nillable().Default(nil),
		field.UUID("deleted_by", uuid.UUID{}).Optional().Nillable(),
	}
}

type softDeleteKey struct{}

// SkipSoftDelete returns a new context that skips the soft-delete interceptor/mutators.
func SkipSoftDelete(parent context.Context) context.Context {
	return context.WithValue(parent, softDeleteKey{}, true)
}

// Interceptors of the MixinSoftDelete.
func (d MixinSoftDelete) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		intercept.TraverseFunc(func(ctx context.Context, q intercept.Query) error {
			// Skip soft-delete, means include soft-deleted entities.
			if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
				return nil
			}
			d.P(q)
			return nil
		}),
	}
}

// Hooks of the MixinSoftDelete.
func (d MixinSoftDelete) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					// Skip soft-delete, means delete the entity permanently.
					if skip, _ := ctx.Value(softDeleteKey{}).(bool); skip {
						return next.Mutate(ctx, m)
					}
					mx, ok := m.(interface {
						SetOp(ent.Op)
						Client() *gen.Client
						SetDeletedAt(time.Time)
						SetDeletedBy(uuid.UUID)
						DeletedBy() (id uuid.UUID, exists bool)
						WhereP(...func(*sql.Selector))
					})
					if !ok {
						return nil, fmt.Errorf("unexpected mutation type %T", m)
					}
					// usr, err := viewer.UserFromContext(ctx)
					usr, ok := ctx.Value("user").(*UserStruct)
					if !ok {
						return nil, fmt.Errorf("user is not logged in")
					}
					d.P(mx)
					mx.SetOp(ent.OpUpdate)
					mx.SetDeletedAt(time.Now())
					if _, exists := mx.DeletedBy(); !exists {
						mx.SetDeletedBy(usr.ID)
					}
					return mx.Client().Mutate(ctx, m)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
	}
}

// P adds a storage-level predicate to the queries and mutations.
func (d MixinSoftDelete) P(w interface{ WhereP(...func(*sql.Selector)) }) {
	w.WhereP(
		sql.FieldIsNull(d.Fields()[0].Descriptor().Name),
	)
}
