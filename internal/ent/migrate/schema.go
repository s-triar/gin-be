// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeUUID, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_by", Type: field.TypeUUID, Nullable: true},
		{Name: "fullname", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "phone", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "is_email_confirmed", Type: field.TypeBool, Default: false},
		{Name: "is_phone_confirmed", Type: field.TypeBool, Default: false},
		{Name: "token_auth", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0]},
			},
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[8]},
			},
			{
				Name:    "user_phone",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[9]},
			},
			{
				Name:    "user_token_auth",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[13]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
	UsersTable.Annotation = &entsql.Annotation{
		Table: "users",
	}
}