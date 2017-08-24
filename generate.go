package elit

import (
	"fmt"
	"reflect"
	"strings"
)

// Generate .
func Generate(v interface{}) {
	rt := reflect.TypeOf(v)
	// rv := reflect.ValueOf(v)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		a := jsonAttributeName(field)

		fmt.Println(a)
	}
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
