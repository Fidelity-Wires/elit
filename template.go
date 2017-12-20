package elit

// Template is root object
type Template struct {
	Template string          `json:"template"`
	Settings Settings        `json:"settings"`
	Mappings map[string]Type `json:"mappings"`
}

// Settings node settings
type Settings struct {
	NumberOfShards   uint     `json:"number_of_shards"`
	NumberOfReplicas uint     `json:"number_of_replicas,omitempty"`
	Analysis         Analysis `json:"analysis,omitempty"`
}

// Analysis settings
type Analysis struct {
	Analyzer map[string]interface{} `json:"analyzer,omitempty"`
	Filter   map[string]Filter      `json:"filter,omitempty"`
}

// Filter for analysis
type Filter struct {
	Type        FilterType `json:"type"`
	Format      string     `json:"format"`
	SynonymPath string     `json:"synonym_path,omitempty"`
	Synonyms    []string   `json:"sysnonyms,omitempty"`
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
	All         *All                `json:"_all,omitempty"`
	Type        PropertyType        `json:"type,omitempty"`
	Format      string              `json:"format,omitempty"`
	FieldData   bool                `json:"fielddata,omitempty"`
	Fields      map[string]Property `json:"fields,omitempty"`
	Properies   map[string]Property `json:"properties,omitempty"`
	IgnoreAbove int                 `json:"ignore_above,omitempty"`
}

// PropertyType .
type PropertyType string

// FilterType .
type FilterType string

const (
	PropertyTypeDate     PropertyType = "date"
	PropertyTypeInteger  PropertyType = "integer"
	PropertyTypeFloat    PropertyType = "float"
	PropertyTypeText     PropertyType = "text"
	PropertyTypeGeoPoint PropertyType = "geo_point"
	PropertyTypeNested   PropertyType = "nested"
	PropertyTypeKeyword  PropertyType = "keyword"
	PropertyTypeBoolean  PropertyType = "boolean"

	FilterTypeSynonym FilterType = "synonym"
)
