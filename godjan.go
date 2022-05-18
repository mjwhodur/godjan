package godjan

import "github.com/google/uuid"

type ModelManager interface {
	Get(query interface{}) (object interface{}, err error)
	GetOrCreate(query interface{}) (object interface{}, created bool, err error)
	Filter(query interface{}, limit uint64, offset uint64) ([]interface{}, error)
	UpdateWhere(query interface{}, update interface{}) (err error)
	All() (returns []interface{})
	Save()
	Go()
	Delete() (deleted bool, err error)
}

type Model struct {
	ModelManager
	ID        uint64
	statement string
}

type Deletion int

const (
	Protect Deletion = iota
	SetNull
	Cascade
	Ignore
)

type InfiniteModel struct {
	ID        uuid.UUID
	statement string
}

type CharField struct {
	Null      bool
	Blank     bool
	MaxLength uint64
	Default   string
	Value     string
}

type NullBooleanField struct {
	Default bool
	Value   bool
}

type ForeignKey struct {
	Null     bool
	Blank    bool
	OnDelete Deletion
	To       ModelManager
}
