package elit

import (
	"fmt"
	"reflect"
	"strings"
)

// PropertyEncoderFunc function for generate
type PropertyEncoderFunc func(key string, v interface{}, properties map[string]Property, opts *GenerateOption) error

// GenerateOption is elit generate options
type GenerateOption struct {
	Presets map[string]PropertyEncoderFunc
}

// Generate .
func Generate(v interface{}, opts *GenerateOption) (map[string]Property, error) {
	m := map[string]Property{}
	rt := reflect.TypeOf(v)
	// rv := reflect.ValueOf(v)

	switch rt.Kind() {
	case reflect.Struct:
		structEncoder("", v, m, opts)
	default:
		return m, fmt.Errorf("v must be struct")
	}

	return m, nil
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
		return structEncoder, nil
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

func boolEncoder(key string, v interface{}, m map[string]Property, opts *GenerateOption) error {
	return nil
}

func stringEncoder(key string, v interface{}, m map[string]Property, opts *GenerateOption) error {
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

func integerEncoder(key string, v interface{}, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeInteger,
	}

	return nil
}

func floatEncoder(key string, v interface{}, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeFloat,
	}

	return nil
}

func structEncoder(key string, v interface{}, m map[string]Property, opts *GenerateOption) error {
	rt := reflect.TypeOf(v)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		encoder, err := TypePropertyEncoder(field, opts)
		if err != nil {
			return fmt.Errorf("TypePropertyEncoder got error: %s", err)
		}

		k := jsonAttributeName(field)
		if err := encoder(k, v, m, opts); err != nil {
			return fmt.Errorf("encoder got error: %s", err)
		}

	}
	return nil
}

func geoPointEncoder(key string, v interface{}, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeGeoPoint,
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
