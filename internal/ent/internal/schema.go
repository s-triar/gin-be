// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = "{\"Schema\":\"gin-be/internal/ent/schema\",\"Package\":\"gin-be/internal/ent\",\"Schemas\":[{\"name\":\"User\",\"config\":{\"Table\":\"\"},\"fields\":[{\"name\":\"created_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_kind\":19,\"immutable\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"created_by\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"default\":true,\"default_kind\":19,\"update_default\":true,\"position\":{\"Index\":2,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"updated_by\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":3,\"MixedIn\":true,\"MixinIndex\":0}},{\"name\":\"deleted_at\",\"type\":{\"Type\":2,\"Ident\":\"\",\"PkgPath\":\"time\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"deleted_by\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":1,\"MixedIn\":true,\"MixinIndex\":1}},{\"name\":\"id\",\"type\":{\"Type\":4,\"Ident\":\"uuid.UUID\",\"PkgPath\":\"github.com/google/uuid\",\"PkgName\":\"uuid\",\"Nillable\":false,\"RType\":{\"Name\":\"UUID\",\"Ident\":\"uuid.UUID\",\"Kind\":17,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":{\"ClockSequence\":{\"In\":[],\"Out\":[{\"Name\":\"int\",\"Ident\":\"int\",\"Kind\":2,\"PkgPath\":\"\",\"Methods\":null}]},\"Domain\":{\"In\":[],\"Out\":[{\"Name\":\"Domain\",\"Ident\":\"uuid.Domain\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"ID\":{\"In\":[],\"Out\":[{\"Name\":\"uint32\",\"Ident\":\"uint32\",\"Kind\":10,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalBinary\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"MarshalText\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"NodeID\":{\"In\":[],\"Out\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}]},\"Scan\":{\"In\":[{\"Name\":\"\",\"Ident\":\"interface {}\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"String\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"Time\":{\"In\":[],\"Out\":[{\"Name\":\"Time\",\"Ident\":\"uuid.Time\",\"Kind\":6,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"URN\":{\"In\":[],\"Out\":[{\"Name\":\"string\",\"Ident\":\"string\",\"Kind\":24,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalBinary\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"UnmarshalText\":{\"In\":[{\"Name\":\"\",\"Ident\":\"[]uint8\",\"Kind\":23,\"PkgPath\":\"\",\"Methods\":null}],\"Out\":[{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Value\":{\"In\":[],\"Out\":[{\"Name\":\"Value\",\"Ident\":\"driver.Value\",\"Kind\":20,\"PkgPath\":\"database/sql/driver\",\"Methods\":null},{\"Name\":\"error\",\"Ident\":\"error\",\"Kind\":20,\"PkgPath\":\"\",\"Methods\":null}]},\"Variant\":{\"In\":[],\"Out\":[{\"Name\":\"Variant\",\"Ident\":\"uuid.Variant\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]},\"Version\":{\"In\":[],\"Out\":[{\"Name\":\"Version\",\"Ident\":\"uuid.Version\",\"Kind\":8,\"PkgPath\":\"github.com/google/uuid\",\"Methods\":null}]}}}},\"default\":true,\"default_kind\":19,\"position\":{\"Index\":0,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"fullname\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":1,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"email\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"validators\":1,\"position\":{\"Index\":2,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"phone\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"unique\":true,\"validators\":1,\"position\":{\"Index\":3,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"password\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"position\":{\"Index\":4,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"is_email_confirmed\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":5,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"is_phone_confirmed\",\"type\":{\"Type\":1,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"default\":true,\"default_value\":false,\"default_kind\":1,\"position\":{\"Index\":6,\"MixedIn\":false,\"MixinIndex\":0}},{\"name\":\"token_auth\",\"type\":{\"Type\":7,\"Ident\":\"\",\"PkgPath\":\"\",\"PkgName\":\"\",\"Nillable\":false,\"RType\":null},\"nillable\":true,\"optional\":true,\"position\":{\"Index\":7,\"MixedIn\":false,\"MixinIndex\":0}}],\"indexes\":[{\"unique\":true,\"fields\":[\"id\"]},{\"unique\":true,\"fields\":[\"email\"]},{\"unique\":true,\"fields\":[\"phone\"]},{\"fields\":[\"token_auth\"]}],\"hooks\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":0},{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"interceptors\":[{\"Index\":0,\"MixedIn\":true,\"MixinIndex\":1}],\"annotations\":{\"EntSQL\":{\"table\":\"users\"}}}],\"Features\":[\"intercept\",\"schema/snapshot\"]}"
