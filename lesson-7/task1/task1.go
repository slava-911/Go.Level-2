package task1

import (
	"errors"
	"reflect"
)

var ErrInputNil = errors.New("input parameter is nil")
var ErrInputNotStruct = errors.New("input parameter is not a struct")
var ErrDifferentType = errors.New("structure field type and map value type are different")
var ErrCantSet = errors.New("can't set value")

type ValuesMap map[string]interface{}
type InputStruct struct {
	FieldString string
	FieldInt    int
	Slice       []int
	Object      struct{ NestedField int }
}

func ChangeStruct(in interface{}, values ValuesMap) error {
	if in == nil {
		return ErrInputNil
	}
	inVal := reflect.ValueOf(in)
	if inVal.Kind() == reflect.Ptr {
		inVal = inVal.Elem()
	}
	if inVal.Kind() != reflect.Struct {
		return ErrInputNotStruct
	}
	for key, v := range values {
		if inVal.FieldByName(key).IsValid() {
			fVal := inVal.FieldByName(key)
			mVal := reflect.ValueOf(v)
			if fVal.Type() != mVal.Type() {
				return ErrDifferentType
			}
			if !fVal.CanSet() {
				return ErrCantSet
			}
			fVal.Set(mVal)
		}
	}
	// for i := 0; i < rval.NumField(); i++ {
	// 	typeField := rval.Type().Field(i)
	// 	if mval, ok := values[typeField.Name]; ok {
	// 		rval.Field(i).Set(reflect.ValueOf(mval))
	// 	}
	return nil
}
