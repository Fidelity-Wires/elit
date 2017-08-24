package elit

import (
	"encoding/json"
	"reflect"
	"testing"
)

// func TestMappingsMarshalJSON(t *testing.T) {
// 	table := []struct {
// 		mappings Mappings
// 		out      string
// 	}{
// 		{
// 			mappings: Mappings{
// 				Type: map[string]Type{
// 					"type1": Type{
// 						Source: Source{
// 							Enabled: false,
// 						},
// 						Properties: map[string]Property{
// 							"page": Property{
// 								Type: "integer",
// 							},
// 						},
// 					},
// 				},
// 			},
// 			out: "{}",
// 		},
// 	}
//
// 	for _, row := range table {
// 		_, err := json.Marshal(row.mappings)
// 		if err != nil {
// 			t.Fatalf("json.Marshal got error: %s", err)
// 		}
//
// 		m := Mappings{}
// 		if err := json.Unmarshal([]byte(row.out), &m); err != nil {
// 			t.Fatalf("json.Unmarshal got error: %s", err)
// 		}
//
// 		if !reflect.DeepEqual(row.mappings, m) {
// 			t.Errorf("deep equal missed. expected(%v) but (%b)", row.mappings, m)
// 		}
// 	}
// }

func TestTemplateMarshalJSON(t *testing.T) {
	table := []struct {
		template Template
		out      string
	}{
		{
			template: Template{
				Template: "te*",
				Settings: Settings{
					NumberOfShards: 1,
				},
				Mappings: map[string]Type{
					"type1": Type{
						Source: &Source{
							Enabled: false,
						},
						Properties: map[string]Property{
							"host_name": {
								Type: "keyword",
							},
							"created_at": {
								Type:   "date",
								Format: "EEE MMM dd HH:mm:ss Z YYYY",
							},
						},
					},
				},
			},
			out: `{"template":"te*","settings":{"number_of_shards":1},"mappings":{"type1":{"_source":{"enabled":false},"properties":{"created_at":{"type":"date","format":"EEE MMM dd HH:mm:ss Z YYYY"},"host_name":{"type":"keyword"}}}}}`,
		},
	}

	for _, row := range table {
		_, err := json.Marshal(row.template)
		if err != nil {
			t.Fatalf("json.Marshal got error: %s", err)
		}

		tpl := Template{}
		if err := json.Unmarshal([]byte(row.out), &tpl); err != nil {
			t.Fatalf("json.Unmarshal got error: %s", err)
		}

		// log.Println("out:", string(b))
		if !reflect.DeepEqual(row.template, tpl) {
			t.Errorf("deep equal missed. expected(%v) but (%v)", row.template, tpl)
		}
	}
}
