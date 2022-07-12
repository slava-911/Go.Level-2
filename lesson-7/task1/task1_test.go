package task1

import (
	"errors"
	"reflect"
	"testing"
)

var testCases = []struct {
	name           string
	input          InputStruct
	inputInt       int
	values         ValuesMap
	expectedOutput InputStruct
	expectedError  error
}{
	{
		name: "case for struct1",
		input: InputStruct{
			FieldString: "abc",
			FieldInt:    123,
			Slice:       []int{1, 2, 3},
			Object:      struct{ NestedField int }{NestedField: 11},
		},
		values: ValuesMap{
			"FieldString": "xyz",
			"FieldInt":    456789,
			"Slice":       []int{4, 5, 6, 7, 8, 9},
			"Object":      struct{ NestedField int }{NestedField: 222},
		},
		expectedOutput: InputStruct{
			FieldString: "xyz",
			FieldInt:    456789,
			Slice:       []int{4, 5, 6, 7, 8, 9},
			Object:      struct{ NestedField int }{NestedField: 222},
		},
	},
	{
		name:          "case for different type",
		input:         InputStruct{FieldString: "abc"},
		values:        ValuesMap{"FieldString": 1},
		expectedError: ErrDifferentType,
	},
	{
		name:          "case for in is nil",
		input:         InputStruct{},
		expectedError: ErrInputNil,
	},
	{
		name:          "case for in is not struct",
		inputInt:      1,
		expectedError: ErrInputNotStruct,
	},
}

func TestChangeStruct(t *testing.T) {
	var err error
	for _, cs := range testCases {
		cs := cs
		t.Run(cs.name, func(t *testing.T) {
			if cs.inputInt != 0 {
				err = ChangeStruct(&cs.inputInt, cs.values)
			} else {
				err = ChangeStruct(&cs.input, cs.values)
			}
			if err != nil {
				if !errors.Is(err, cs.expectedError) {
					t.Fatalf("unexpected error: %s", err.Error())
				}
				return
			}
			if !reflect.DeepEqual(cs.input, cs.expectedOutput) {
				t.Fatalf("wrong result, got: %v, expected: %v", cs.input, cs.expectedOutput)
			}
		})
	}
}

func BenchmarkChangeStructCase1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChangeStruct(&testCases[0].input, testCases[0].values)
	}
}

func BenchmarkChangeStructCase2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChangeStruct(&testCases[1].input, testCases[1].values)
	}
}

func BenchmarkChangeStructCase3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChangeStruct(&testCases[2].input, testCases[2].values)
	}
}

func BenchmarkChangeStructCase4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChangeStruct(1, ValuesMap{})
	}
}
