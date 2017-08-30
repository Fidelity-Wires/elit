package elit

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type propertyEncoderFunc func(v interface{}) (Property, error)

// Generate .
func Generate(v interface{}) {
	rt := reflect.TypeOf(v)
	// rv := reflect.ValueOf(v)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		a := jsonAttributeName(field)
		json.Marshal(a)

	}
}

func TypePropertyEncoder(field reflect.StructField) (propertyEncoderFunc, error) {
	switch field.Type.Kind() {
	case reflect.Bool:
		return boolEncoder, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return boolEncoder, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return boolEncoder, nil
	case reflect.Float32:
		return boolEncoder, nil
	case reflect.Float64:
		return boolEncoder, nil
	case reflect.String:
		return boolEncoder, nil
	case reflect.Interface:
		return boolEncoder, nil
	case reflect.Struct:
		return boolEncoder, nil
	case reflect.Map:
		return boolEncoder, nil
	case reflect.Slice:
		return boolEncoder, nil
	case reflect.Array:
		return boolEncoder, nil
	case reflect.Ptr:
		return boolEncoder, nil
	default:
		return nil, fmt.Errorf("unsupported type")
	}
}

func boolEncoder(v interface{}) (Property, error) {
	return Property{}, nil
}

func jsonAttributeName(f reflect.StructField) string {
	j := f.Tag.Get("json")

	if j == "" {
		return f.Name
	}

	l := strings.Split(j, ",")
	if len(l) == 0 {
		return f.Name
	}

	first := l[0]
	switch first {
	case "-":
		return ""
	case "":
		return f.Name
	}

	return first
}
