package godjan

import (
	"fmt"
	"reflect"
	"strings"
)

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
	Create(arg interface{}) (object interface{}, created bool, err error)
	Save()
	Go()
	Delete() (deleted bool, err error)
}

type Model struct {
	ModelManager
	ID               uint64
	statement        string
	description      map[string]FieldType
	tableDescription map[string]map[string]PostgresType
	oldValues        map[string]string
	tableName        string
}

type Deletion int

const (
	Protect Deletion = iota
	SetNull
	Cascade
	Ignore
)

func (m *Model) describe() {

}

func (m *Model) Get(query interface{}) (object interface{}, err error) {
	if reflect.TypeOf(query).Kind() == reflect.Pointer {
		fmt.Println("Found a pointer")
	}

	if reflect.TypeOf(query).Kind() == reflect.Struct {
		fmt.Println("Found a struct")
		vf := reflect.VisibleFields(reflect.TypeOf(query))
		for field, val := range vf {

			fmt.Println(field, val, val.Name, val.Type)
		}

		o := reflect.Zero(reflect.TypeOf(query))
		oInterface := o.Interface()
		return oInterface, nil
	}

	rv := reflect.ValueOf(query)
	m.tableName = strings.ToLower(reflect.Indirect(rv).Type().Name()) + "s"
	fmt.Println("m.tableName", m.tableName)
	m.describe()
	return nil, nil
}

func (m *Model) UnsafeGet(query interface{}) (object interface{}) {
	m.describe()
	return nil
}

func (m *Model) GetOrCreate(query interface{}) (object interface{}, created bool, err error) {
	m.describe()
	panic("implement me")
}

func (m *Model) Filter(query interface{}, limit uint64, offset uint64) ([]interface{}, error) {
	m.describe()
	//TODO implement me
	panic("implement me")
}

func (m *Model) UpdateWhere(query interface{}, update interface{}) (err error) {
	m.describe()
	//TODO implement me
	panic("implement me")
}

func (m *Model) All() (returns []interface{}) {
	m.describe()
	//TODO implement me
	panic("implement me")
}

func (m *Model) Save() {
	m.describe()

	panic("implement me")
}

func (m *Model) Go() {
	//TODO implement me
	panic("implement me")
}

func (i *Model) Create(arg interface{}) (object interface{}, created bool, err error) {
	// TODO Implement me
	panic("implement me")
}

func (m *Model) Delete() (deleted bool, err error) {
	//TODO implement me
	panic("implement me")
}
