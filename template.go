package elit

// Template is root object
type Template struct {
	Template string          `json:"template"`
	Settings Settings        `json:"settings"`
	Mappings map[string]Type `json:"mappings"`
}

// Settings node settings
type Settings struct {
	NumberOfShards   uint `json:"number_of_shards"`
	NumberOfReplicas uint `json:"number_of_replicas,omitempty"`
}

// Type .
type Type struct {
	All        *All                `json:"_all,omitempty"`
	Source     *Source             `json:"_source,omitempty"`
	Properties map[string]Property `json:"properties,omitempty"`
}

// Source .
type Source struct {
	Enabled bool `json:"enabled"`
}

// All .
type All struct {
	Enabled bool `json:"enabled"`
}

// Property .
type Property struct {
	All       *All                `json:"_all,omitempty"`
	Type      string              `json:"type,omitempty"`
	Format    string              `json:"format,omitempty"`
	FieldData string              `json:"fielddata,omitempty"`
	Fields    map[string]Field    `json:"fields,omitempty"`
	Properies map[string]Property `json:"properties,omitempty"`
}

// Field .
type Field struct {
	Type        string `json:"type"`
	IgnoreAbove int    `json:"ignore_above"`
}

// // MarshalJSON for json marshaler
// func (m Mappings) MarshalJSON() ([]byte, error) {
// 	b := bytes.Buffer{}
//
// 	if _, err := b.WriteString("{"); err != nil {
// 		return nil, err
// 	}
//
// 	l := len(m.Type)
// 	c := 0
// 	for k, i := range m.Type {
// 		if _, err := b.WriteString(fmt.Sprintf("\"%s\":", k)); err != nil {
// 			return nil, err
// 		}
//
// 		jb, err := json.Marshal(i)
// 		if err != nil {
// 			return nil, err
// 		}
//
// 		if _, err := b.Write(jb); err != nil {
// 			return nil, err
// 		}
//
// 		c = c + 1
// 		if c != l {
// 			b.WriteString(",")
// 		}
// 	}
//
// 	if _, err := b.WriteString("}"); err != nil {
// 		return nil, err
// 	}
//
// 	return b.Bytes(), nil
// }
