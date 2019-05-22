package elit

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

func ExampleGenerate() {
	t := Template{
		Template: "sample_template_*",
		Settings: Settings{
			NumberOfShards:   5,
			NumberOfReplicas: 1,
		},
		Mappings: Mappings{
			Properties: map[string]Property{},
		},
	}

	opts := NewGenerateOption()
	p, err := Generate(Human{}, opts)
	if err != nil {
		panic(err)
	}

	t.Mappings.Properties = p

	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	// Output:
	// {"template":"sample_template_*","settings":{"number_of_shards":5,"number_of_replicas":1},"mappings":{"properties":{"Age":{"type":"integer"},"Name":{"type":"text"}}}}
}
