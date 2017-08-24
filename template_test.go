package elit

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMappingsMarshalJSON(t *testing.T) {
	table := []struct {
		mappings Mappings
		out      string
	}{
		{
			mappings: Mappings{
				Type: map[string]interface{}{
					"integer": float64(1),
					"string":  "string",
					"boolean": true,
				},
			},
		},
	}

	for _, row := range table {
		b, err := json.Marshal(row.mappings)
		if err != nil {
			t.Fatalf("json.Marshal got error: %s", err)
		}

		o := map[string]interface{}{}
		if err := json.Unmarshal(b, &o); err != nil {
			t.Fatalf("json.Unmarshal got error: %s", err)
		}

		for k := range o {
			if !reflect.DeepEqual(o[k], row.mappings.Type[k]) {
				t.Errorf("type key (%s) expected (%v) but (%v)", k, row.mappings.Type[k], o[k])
			}
		}
	}
}
