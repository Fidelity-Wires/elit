package elit

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// PropertyEncoderFunc function for generate
type PropertyEncoderFunc func(reflect.StructField, map[string]Property) error

type GenerateOption struct {
	Presets map[string]PropertyEncoderFunc
}

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

// TypePropertyEncoder .
func TypePropertyEncoder(field reflect.StructField, opt *GenerateOption) (PropertyEncoderFunc, error) {
	elit := elitPropertyName(field)
	if elit != "" {
		return opt.Presets[elit], nil
	}

	switch field.Type.Kind() {
	case reflect.Bool:
		return boolEncoder, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return integerEncoder, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return integerEncoder, nil
	case reflect.Float32:
		return floatEncoder, nil
	case reflect.Float64:
		return floatEncoder, nil
	case reflect.String:
		return stringEncoder, nil
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

func boolEncoder(f reflect.StructField, m map[string]Property) error {
	return nil
}

func stringEncoder(f reflect.StructField, m map[string]Property) error {
	key := jsonAttributeName(f)
	m[key] = Property{
		Type:      PropertyTypeText,
		FieldData: true,
		Fields: map[string]Property{
			"keyword": Property{
				Type:        PropertyTypeKeyword,
				IgnoreAbove: 256,
			},
		},
	}

	return nil
}

func integerEncoder(f reflect.StructField, m map[string]Property) error {
	key := jsonAttributeName(f)
	m[key] = Property{
		Type: PropertyTypeInteger,
	}

	return nil
}

func floatEncoder(f reflect.StructField, m map[string]Property) error {
	key := jsonAttributeName(f)
	m[key] = Property{
		Type: PropertyTypeFloat,
	}

	return nil
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

func elitPropertyName(f reflect.StructField) string {
	elit := f.Tag.Get("elit")

	return elit
}
