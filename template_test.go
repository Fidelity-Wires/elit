package elit

import (
	"encoding/json"
	"log"
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

func TestTemplateMarshalJSON(t *testing.T) {
	table := []struct {
		template Template
		out      string
	}{
		{
			template: Template{
				Template: "template-*",
				Settings: Settings{
					NumberOfReplicas: 4,
					NumberOfShards:   1,
				},
			},
			out: `{"template":"template-*","settings":{"number_of_shards":1,"number_of_replicas":4},"mappings":{}}`,
		},
	}

	for _, row := range table {
		b, err := json.Marshal(row.template)
		if err != nil {
			t.Fatalf("json.Marshal got error: %s", err)
		}

		log.Println(string(b))
		if string(b) != row.out {
			t.Errorf("out expected (%s) but (%s)", string(b), row.out)
		}
	}
}
