package elit

import "testing"

type SampleModel struct {
	ID        int `json:"id"`
	Empty     string
	Hyphen    string `json:"-"`
	OmitEmpty string `json:"omit_empty,omitempty"`
	Normal    string `json:"normal"`
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
		Generate(row.input)
	}
}
