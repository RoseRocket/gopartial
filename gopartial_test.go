package gopartial

import (
	"reflect"
	"testing"
	"time"

	"github.com/guregu/null"
)

func TestPartialUpdate(t *testing.T) {
	type args struct {
		dest           interface{}
		partial        map[string]interface{}
		tagName        string
		updaters       []func(reflect.Value, reflect.Value) bool
		skipConditions []func(reflect.StructField) bool
	}
	type test struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}

	type sub struct {
		FieldA string `json:"fielda"`
		FieldB string `json:"fieldb"`
	}

	type destination struct {
		Field0   string      `json:"field0" props:"readonly"`
		Field1   string      `json:"field1"`
		Field1p  *string     `json:"field1p"`
		Field2   null.String `json:"field2"`
		Field3   float64     `json:"field3"`
		Field3p  *float64    `json:"field3p"`
		Field4   null.Float  `json:"field4"`
		Field5   int         `json:"field5"`
		Field5p  *int        `json:"field5p"`
		Field6   null.Int    `json:"field6"`
		Field7   bool        `json:"field7"`
		Field7p  *bool       `json:"field7p"`
		Field8   null.Bool   `json:"field8"`
		Field9   time.Time   `json:"field9"`
		Field9p  *time.Time  `json:"field9p"`
		Field10  null.Time   `json:"field10"`
		Field11  sub         `json:"field11"`
		Field11p *sub        `json:"field11p"`
	}

	var str = "foo"
	var dateStr = "2017-11-22T20:30:26.716Z"
	var i = 1
	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1
	var f32 float32 = 1.1
	var f64 = 1.1
	var check = true
	tests := []test{
		test{
			name: "Dest is a non pointer to struct",
			args: args{
				dest:           destination{},
				partial:        map[string]interface{}{},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    nil,
			wantErr: true,
		},
		test{
			name: "Dest is a pointer to string",
			args: args{
				dest:           &str,
				partial:        map[string]interface{}{},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    nil,
			wantErr: true,
		},

		// readonly
		test{
			name: "Update field0 (string field, readonly)",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field0": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// string
		test{
			name: "Update field1 (string) with strinng",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field1": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field1"},
			wantErr: false,
		},
		test{
			name: "Update field1 (string) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field1": 1,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// null.String
		test{
			name: "Update field2 (null.String) with string",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field2": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field2"},
			wantErr: false,
		},
		test{
			name: "Update field2 (null.String) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field2": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field2"},
			wantErr: false,
		},
		test{
			name: "Update field2 (null.String) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field2": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// float64
		test{
			name: "Update field3 (float64) with float32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3": f32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3"},
			wantErr: false,
		},
		test{
			name: "Update field3 (float64) with float64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3": f64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3"},
			wantErr: false,
		},
		test{
			name: "Update field3 (float64) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3"},
			wantErr: false,
		},
		test{
			name: "Update field3 (float64) with int8",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3": i8,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3"},
			wantErr: false,
		},
		test{
			name: "Update field3 (float64) with int16",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3": i16,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3"},
			wantErr: false,
		},
		test{
			name: "Update field3 (float64) with int32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3": i32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3"},
			wantErr: false,
		},
		test{
			name: "Update field3 (float64) with int64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3": i64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3"},
			wantErr: false,
		},
		test{
			name: "Update field3 (float64) to string",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// *float64
		test{
			name: "Update field3p (*float64) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3p"},
			wantErr: false,
		},
		test{
			name: "Update field3p (*float64) with float32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": f32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3p"},
			wantErr: false,
		},
		test{
			name: "Update field3p (*float64) with float64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": f64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3p"},
			wantErr: false,
		},
		test{
			name: "Update field3p (*float64) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3p"},
			wantErr: false,
		},
		test{
			name: "Update field3p (*float64) with int8",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": i8,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3p"},
			wantErr: false,
		},
		test{
			name: "Update field3p (*float64) with int16",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": i16,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3p"},
			wantErr: false,
		},
		test{
			name: "Update field3 (*float64) with int32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": i32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3p"},
			wantErr: false,
		},
		test{
			name: "Update field3 (*float64) with int64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": i64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field3p"},
			wantErr: false,
		},
		test{
			name: "Update field3 (*float64) to string",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field3p": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// null.Float
		test{
			name: "Update field4 (null.Float) with float32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": f32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field4"},
			wantErr: false,
		},
		test{
			name: "Update field4 (null.Float) with float64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": f64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field4"},
			wantErr: false,
		},
		test{
			name: "Update field4 (null.Float) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field4"},
			wantErr: false,
		},
		test{
			name: "Update field4 (null.Float) with int8",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": i8,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field4"},
			wantErr: false,
		},
		test{
			name: "Update field4 (null.Float) with int16",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": i16,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field4"},
			wantErr: false,
		},
		test{
			name: "Update field4 (null.Float) with int32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": i32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field4"},
			wantErr: false,
		},
		test{
			name: "Update field4 (null.Float) with int64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": i64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field4"},
			wantErr: false,
		},
		test{
			name: "Update field4 (null.Float) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field4"},
			wantErr: false,
		},
		test{
			name: "Update field4 (null.Int) with string",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field4": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// int
		test{
			name: "Update field5 (int) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5"},
			wantErr: false,
		},
		test{
			name: "Update field5 (int) with int8",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5": i8,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5"},
			wantErr: false,
		},
		test{
			name: "Update field5 (int) with int16",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5": i16,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5"},
			wantErr: false,
		},
		test{
			name: "Update field5 (int) with int32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5": i32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5"},
			wantErr: false,
		},
		test{
			name: "Update field5 (int) with int64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5": i64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5"},
			wantErr: false,
		},
		test{
			name: "Update field5 (int) with float32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5": f32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5"},
			wantErr: false,
		},
		test{
			name: "Update field5 (int) with float64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5": f64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5"},
			wantErr: false,
		},
		test{
			name: "Update field5 (int) to string",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// *int
		test{
			name: "Update field5p (*int) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5p"},
			wantErr: false,
		},
		test{
			name: "Update field5p (*int) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5p"},
			wantErr: false,
		},
		test{
			name: "Update field5p (*int) with int8",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": i8,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5p"},
			wantErr: false,
		},
		test{
			name: "Update field5p (*int) with int16",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": i16,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5p"},
			wantErr: false,
		},
		test{
			name: "Update field5p (*int) with int32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": i32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5p"},
			wantErr: false,
		},
		test{
			name: "Update field5p (*int) with int64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": i64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5p"},
			wantErr: false,
		},
		test{
			name: "Update field5p (*int) with float32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": f32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5p"},
			wantErr: false,
		},
		test{
			name: "Update field5p (*int) with float64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": f64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field5p"},
			wantErr: false,
		},
		test{
			name: "Update field5p (*int) to string",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field5p": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// null.Int
		test{
			name: "Update field6 (null.Int) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field6"},
			wantErr: false,
		},
		test{
			name: "Update field6 (null.Int) with int8",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": i8,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field6"},
			wantErr: false,
		},
		test{
			name: "Update field6 (null.Int) with int16",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": i16,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field6"},
			wantErr: false,
		},
		test{
			name: "Update field6 (null.Int) with int32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": i32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field6"},
			wantErr: false,
		},
		test{
			name: "Update field6 (null.Int) with int64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": i64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field6"},
			wantErr: false,
		},
		test{
			name: "Update field6 (null.Int) with float32",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": f32,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field6"},
			wantErr: false,
		},
		test{
			name: "Update field6 (null.Int) with float64",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": f64,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field6"},
			wantErr: false,
		},
		test{
			name: "Update field6 (null.Int) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field6"},
			wantErr: false,
		},
		test{
			name: "Update field6 (null.Int) with string",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field6": str,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// bool
		test{
			name: "Update field7 (bool) with bool",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field7": check,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field7"},
			wantErr: false,
		},
		test{
			name: "Update field7 (bool) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field7": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},
		test{
			name: "Update field7 (bool) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field7": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// *bool
		test{
			name: "Update field7p (*bool) with bool",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field7p": check,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field7p"},
			wantErr: false,
		},
		test{
			name: "Update field7p (*bool) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field7p": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field7p"},
			wantErr: false,
		},
		test{
			name: "Update field7p (*bool) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field7p": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// time.Time
		test{
			name: "Update field9 (time.Time) with datestring",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field9": dateStr,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field9"},
			wantErr: false,
		},
		test{
			name: "Update field9 (time.Time) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field9": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},
		test{
			name: "Update field9 (time.Time) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field9": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// *time.Time
		test{
			name: "Update field9p (*time.Time) with datestring",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field9p": dateStr,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field9p"},
			wantErr: false,
		},
		test{
			name: "Update field9p (*time.Time) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field9p": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field9p"},
			wantErr: false,
		},
		test{
			name: "Update field9p (*time.Time) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field9p": i,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},

		// null.Time
		test{
			name: "Update field10 (null.Time) with datestring",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field10": dateStr,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field10"},
			wantErr: false,
		},
		test{
			name: "Update field10 (null.Time) with null",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field10": nil,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{"Field10"},
			wantErr: false,
		},
		test{
			name: "Update field10 (null.Time) with int",
			args: args{
				dest: &destination{},
				partial: map[string]interface{}{
					"field10": 1,
				},
				tagName:        "json",
				updaters:       Updaters,
				skipConditions: SkipConditions,
			},
			want:    []string{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PartialUpdate(tt.args.dest, tt.args.partial, tt.args.tagName, tt.args.skipConditions, tt.args.updaters)
			if (err != nil) != tt.wantErr {
				t.Errorf("PartialUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PartialUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}
