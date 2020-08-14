package env

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type field struct {
	name     string
	value    string
	envVar   string
	kind     reflect.Kind
	required bool
}

func (f *field) String() string {
	return fmt.Sprintf("<Field name: %s, value: '%s', envVar: %s, type: %s, required: %t />", f.name, f.value, f.envVar, f.kind.String(), f.required)
}

type Parser interface {
	Parse(s interface{}) (interface{}, error)
}

func parseField(f reflect.StructField) (*field, error) {
	tagVal := f.Tag.Get("env")
	defaultTagVal := fmt.Sprintf("%s", strings.Join(camelToSlice(f.Name), "_"))

	if tagVal == "" {
		tagVal = defaultTagVal
	}

	tag, opts := parseTag(tagVal)
	envVar := strings.ToUpper(tag)
	name := f.Name
	required := !opts.Contains("omitempty")
	val, ok := os.LookupEnv(envVar)

	fmt.Println(tag, opts)

	if required && !ok {
		return nil, fmt.Errorf("Could not parse env variable '%s' for field '%s'", tag, name)
	}

	return &field{
		name:     name,
		value:    val,
		envVar:   envVar,
		required: required,
		kind:     f.Type.Kind(),
	}, nil

}

// Parse parses each field and sets struct fields accordingly
func Parse(s interface{}) error {
	v := reflect.TypeOf(s).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		fmt.Printf("f: %+v\n", f)
		fmt.Printf("v: %+v\n", reflect.ValueOf(s).Elem().Field(i).Interface())

		if f.Type.Kind() == reflect.Struct {

			structVal := reflect.ValueOf(s).Elem().Field(i).Interface()
			err := Parse(structVal)
			if err != nil {
				return err
			}

			// TODO: Set field interface
			// else {
			// reflect.ValueOf(s).Elem().Field(i).Set(structVal)
			// }

			// reflect.ValueOf(s).Elem().Field(i).SetPointer(&structVal)
		}

		field, err := parseField(f)
		if err != nil {
			return err
		}

		fmt.Println(field)

		fx := mappers[field.kind]
		if fx == nil {
			return fmt.Errorf("Unsupported field kind: %s, field: %s", field.kind, field.name)
		}

		if fx != nil {
			refVal := reflect.ValueOf(s).Elem().Field(i)
			if !refVal.CanSet() {
				return fmt.Errorf("Could not set value for field: %s", field.name)
			}

			err := fx(reflect.ValueOf(s).Elem().Field(i), field.value)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
