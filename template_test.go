package elit

import (
	"encoding/json"
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
					"integer": 1,
					"string":  "string",
					"boolean": true,
				},
			},
			out: `{"integer":1,"string":"string","boolean":true}`,
		},
	}

	for _, row := range table {
		b, err := json.Marshal(row.mappings)
		if err != nil {
			t.Fatalf("json.Marshal got error: %s", err)
		}

		if string(b) != row.out {
			t.Errorf("output expected(%s) but (%s)", row.out, string(b))
		}
	}
}
