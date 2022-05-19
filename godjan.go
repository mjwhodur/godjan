package godjan

import "github.com/google/uuid"

type FieldManager interface {
	Set(value interface{}) (err error)
	Get() (value interface{})
}

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
	ID          uint64
	statement   string
	description map[string]string
}

type Deletion int

const (
	Protect Deletion = iota
	SetNull
	Cascade
	Ignore
)

type InfiniteModel struct {
	ID          uuid.UUID
	statement   string
	description map[string]string
}

func (m *Model) describe() {

}

func (m *Model) Get(query interface{}) (object interface{}, err error) {
	m.describe()
	panic("implement me")
}

func (m *Model) GetOrCreate(query interface{}) (object interface{}, created bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *Model) Filter(query interface{}, limit uint64, offset uint64) ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Model) UpdateWhere(query interface{}, update interface{}) (err error) {
	//TODO implement me
	panic("implement me")
}

func (m *Model) All() (returns []interface{}) {
	//TODO implement me
	panic("implement me")
}

func (m *Model) Save() {
	//TODO implement me
	panic("implement me")
}

func (m *Model) Go() {
	//TODO implement me
	panic("implement me")
}

func (m *Model) Delete() (deleted bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfiniteModel) Get(query interface{}) (object interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfiniteModel) GetOrCreate(query interface{}) (object interface{}, created bool,
	err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfiniteModel) Filter(query interface{}, limit uint64, offset uint64) ([]interface{},
	error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfiniteModel) UpdateWhere(query interface{}, update interface{}) (err error) {
	//TODO implement me
	panic("implement me")
}

func (i *InfiniteModel) All() (returns []interface{}) {
	//TODO implement me
	panic("implement me")
}

func (i *InfiniteModel) Save() {
	//TODO implement me
	panic("implement me")
}

func (i *InfiniteModel) Go() {
	//TODO implement me
	panic("implement me")
}

func (i *InfiniteModel) Delete() (deleted bool, err error) {
	//TODO implement me
	panic("implement me")
}
