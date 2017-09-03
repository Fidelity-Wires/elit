package elit

import (
	"encoding/json"
	"reflect"
	"testing"
)

type SampleModel struct {
	BaseSampleModel
	ID          int     `json:"id"`
	Point       uint    `json:"point"`
	Second      float64 `json:"second"`
	SmallSecond float32 `json:"small_second"`
	OK          bool    `json:"ok"`
	Empty       string
	Hyphen      string    `json:"-"`
	OmitEmpty   string    `json:"omit_empty,omitempty"`
	Normal      string    `json:"normal"`
	StringList  []string  `json:"string_list"`
	IntList     []int     `json:"int_list"`
	Geo         GeoModel  `json:"geo" elit:"geo"`
	Sub         SubModel  `json:"sub"`
	PSub        *SubModel `json:"psub"`
}

type BaseSampleModel struct {
	BaseBaseSampleMode
	Name string `json:"base_sample_model_name"`
}

type BaseBaseSampleMode struct {
	Live uint8 `json:"base_base_sample_model_live"`
}

type GeoModel struct {
	Lat int
	Lon int
}

type SubModel struct {
	BaseSubModel
	None  int    `json:"-"`
	Body  string `json:"body"`
	Child ChildModel
}

type BaseSubModel struct {
	BaseSubModelTitle string `json:"base_sub_model_title"`
}

type ChildModel struct {
	Name string `json:"name"`
}

func TestGenerate(t *testing.T) {
	table := []struct {
		input  interface{}
		result string
	}{
		{
			input:  SampleModel{},
			result: `{"Empty":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}},"base_base_sample_model_live":{"type":"integer"},"base_sample_model_name":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}},"geo":{"type":"geo_point"},"id":{"type":"integer"},"int_list":{"type":"integer"},"normal":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}},"ok":{"type":"boolean"},"omit_empty":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}},"point":{"type":"integer"},"psub":{"type":"boolean"},"second":{"type":"float"},"small_second":{"type":"float"},"string_list":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}},"sub":{"type":"nested","properties":{"":{"type":"integer"},"Child":{"type":"nested","properties":{"name":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}}}},"base_sub_model_title":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}},"body":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}}}}}`,
		},
	}

	for _, row := range table {
		propertyMap, err := Generate(row.input, NewGenerateOption())
		if err != nil {
			t.Fatalf("Generate got error: %s", err)
		}

		res := map[string]Property{}
		if err := json.Unmarshal([]byte(row.result), &res); err != nil {
			t.Errorf("json.Unmarshal got error: %s", err)
		}

		// b, _ := json.Marshal(propertyMap)
		// fmt.Println(string(b))

		if !reflect.DeepEqual(res, propertyMap) {
			t.Errorf("result map expected (%v) but (%v)", res, propertyMap)
		}
	}
}
