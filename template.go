package elit

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Template is root object
type Template struct {
	Template string   `json:"template"`
	Settings Settings `json:"settings"`
	Mappings Mappings `json:"mappings"`
}

// Settings node settings
type Settings struct {
	NumberOfShards   uint `json:"number_of_shards"`
	NumberOfReplicas uint `json:"number_of_replicas"`
}

// Mappings has pair of key properties
type Mappings struct {
	Type map[string]interface{}
}

// Properties .
type Properties struct {
	Properties map[string]Property
}

// Property .
type Property struct {
	Type      string           `json:"type"`
	Format    string           `json:"format"`
	FieldData string           `json:"fielddata"`
	Fields    map[string]Field `json:"fields"`
	Properies map[string]Property
}

// Field .
type Field struct {
	Type        string `json:"type"`
	IgnoreAbove int    `json:"ignore_above"`
}

// MarshalJSON for json marshaler
func (m Mappings) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}

	if _, err := b.WriteString("{"); err != nil {
		return nil, err
	}

	l := len(m.Type)
	c := 0
	for k, i := range m.Type {
		if _, err := b.WriteString(fmt.Sprintf("\"%s\":", k)); err != nil {
			return nil, err
		}

		jb, err := json.Marshal(i)
		if err != nil {
			return nil, err
		}

		if _, err := b.Write(jb); err != nil {
			return nil, err
		}

		c = c + 1
		if c != l {
			b.WriteString(",")
		}
	}

	if _, err := b.WriteString("}"); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
