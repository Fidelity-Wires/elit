package elit

import (
	"fmt"
	"reflect"
	"strings"
)

// PropertyEncoderFunc function for generate
type PropertyEncoderFunc func(key string, rt reflect.Type, properties map[string]Property, opts *GenerateOption) error

// Generate .
func Generate(v interface{}, opts *GenerateOption) (map[string]Property, error) {
	m := map[string]Property{}

	if err := generate(v, m, opts); err != nil {
		return m, fmt.Errorf("E got error: %s", err)
	}

	return m, nil
}

func generate(v interface{}, m map[string]Property, opts *GenerateOption) error {
	rt := reflect.TypeOf(v)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		key := jsonAttributeName(field)
		if key != "" {
			encoder, err := TypePropertyEncoder(field, opts)
			if err != nil {
				return fmt.Errorf("TypePropertyEncoder got error: %s", err)
			}
			if err := encoder(key, field.Type, m, opts); err != nil {
				return fmt.Errorf("encoder got error: %s", err)
			}
		}
	}

	return nil
}

// TypePropertyEncoder .
func TypePropertyEncoder(field reflect.StructField, opts *GenerateOption) (PropertyEncoderFunc, error) {
	if opts != nil {
		elit := elitPropertyName(field)
		if elit != "" {
			e, ok := opts.Presets[elit]
			if !ok {
				return nil, fmt.Errorf("encoder not found in presets")
			}
			return e, nil
		}
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
	case reflect.Struct:
		return structEncoder, nil
	case reflect.Ptr, reflect.Slice:
		return boolEncoder, nil
	case reflect.Array, reflect.Map, reflect.Interface:
		return nil, fmt.Errorf("unsupported type: %v", field.Type.Kind())
	default:
		return nil, fmt.Errorf("unsupported type: %v", field.Type.Kind())
	}
}

func structEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	child := map[string]Property{}
	m[key] = Property{
		Type:      PropertyTypeNested,
		Properies: child,
	}

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)

		k := jsonAttributeName(field)
		if key != "" {
			encoder, err := TypePropertyEncoder(field, opts)
			if err != nil {
				return fmt.Errorf("TypePropertyEncoder got error: %s", err)
			}
			if err := encoder(k, field.Type, child, opts); err != nil {
				return fmt.Errorf("encoder got error: %s", err)
			}
		}
	}

	return nil
}

func boolEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	return nil
}

func stringEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
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

func integerEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeInteger,
	}

	return nil
}

func floatEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeFloat,
	}

	return nil
}

func geoPointEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
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
