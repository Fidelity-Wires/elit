package elit_test

import (
	"encoding/json"
	"fmt"

	"github.com/fidelitywires/elit"
)

type Human struct {
	Name string
	Age  int
}

func ExampleGenerate() {
	t := elit.Template{
		Template: "sample_template_*",
		Settings: elit.Settings{
			NumberOfShards:   5,
			NumberOfReplicas: 1,
		},
		Mappings: map[string]elit.Type{},
	}

	opts := elit.NewGenerateOption()
	p, err := elit.Generate(Human{}, opts)
	if err != nil {
		panic(err)
	}

	t.Mappings["sample"] = elit.Type{
		Properties: p,
	}

	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	// Output:
	// {"template":"sample_template_*","settings":{"number_of_shards":5,"number_of_replicas":1},"mappings":{"sample":{"properties":{"Age":{"type":"integer"},"Name":{"type":"text"}}}}}
}
