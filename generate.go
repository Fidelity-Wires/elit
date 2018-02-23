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
	fields := Fields(rt)

	for _, field := range fields {
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

// Fields get all fields
func Fields(t reflect.Type) []reflect.StructField {
	fields := []reflect.StructField{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			fields = Fields(field.Type)
		} else {
			fields = append(fields, field)
		}
	}

	return fields
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

	return selectFromKind(field.Type.Kind(), opts)
}

func selectFromKind(k reflect.Kind, opts *GenerateOption) (PropertyEncoderFunc, error) {
	e, ok := opts.Encoders[k]
	if ok {
		return e, nil
	}

	switch k {
	case reflect.Bool:
		return boolEncoder, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return integer32Encoder, nil
	case reflect.Int64, reflect.Uint64:
		return integer64Encoder, nil
	case reflect.Float32:
		return float32Encoder, nil
	case reflect.Float64:
		return float64Encoder, nil
	case reflect.String:
		return stringEncoder, nil
	case reflect.Struct:
		return structEncoder, nil
	case reflect.Array, reflect.Slice:
		return arrayEncoder, nil
	case reflect.Ptr:
		return ptrEncoder, nil
	case reflect.Map, reflect.Interface:
		return nil, fmt.Errorf("unsupported type: %v", k)
	}

	return nil, fmt.Errorf("unsupported type: %v", k)
}

func structEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	child := map[string]Property{}
	m[key] = Property{
		Type:      PropertyTypeNested,
		Properies: child,
	}

	fields := Fields(rt)

	for _, field := range fields {
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

func ptrEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	encoder, err := selectFromKind(rt.Elem().Kind(), opts)
	if err != nil {
		return fmt.Errorf("selectFromKind got error: %s", err)
	}

	if err := encoder(key, rt.Elem(), m, opts); err != nil {
		return fmt.Errorf("encoder got error: %s", err)
	}

	return nil
}

func arrayEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	encoder, err := selectFromKind(rt.Elem().Kind(), opts)
	if err != nil {
		return fmt.Errorf("selectFromKind got error: %s", err)
	}

	if err := encoder(key, rt.Elem(), m, opts); err != nil {
		return fmt.Errorf("encoder got error: %s", err)
	}

	return nil
}

func boolEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeBoolean,
	}

	return nil
}

func stringEncoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeText,
	}

	return nil
}

func integer32Encoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeInteger32,
	}

	return nil
}

func integer64Encoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeInteger64,
	}

	return nil
}

func float32Encoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeFloat32,
	}

	return nil
}

func float64Encoder(key string, rt reflect.Type, m map[string]Property, opts *GenerateOption) error {
	m[key] = Property{
		Type: PropertyTypeFloat64,
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
