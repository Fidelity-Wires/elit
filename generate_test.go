package elit

import (
	"encoding/json"
	"fmt"
	"testing"
)

type SampleModel struct {
	ID        int     `json:"id"`
	Second    float64 `json:"second"`
	Empty     string
	Hyphen    string   `json:"-"`
	OmitEmpty string   `json:"omit_empty,omitempty"`
	Normal    string   `json:"normal"`
	Sub       SubModel `json:"sub"`
}

type SubModel struct {
	None  int    `json:"-"`
	Body  string `json:"body"`
	Child ChildModel
}

type ChildModel struct {
	Name string `json:"name"`
}

func TestGenerate(t *testing.T) {
	table := []struct {
		input interface{}
	}{
		{
			input: SampleModel{},
		},
	}

	for _, row := range table {
		tpl, err := Generate(row.input, nil)
		if err != nil {
			t.Fatalf("Generate got error: %s", err)

		}

		b, err := json.Marshal(tpl)
		if err != nil {
			t.Fatalf("json.Marshal: %s", err)
		}

		fmt.Println(string(b))
	}
}
