package elit

import (
	"bytes"
	"strings"
)

// Synonym .
type Synonym struct {
	key     string
	aliases []string
}

// NewMapSynonym .
func NewMapSynonym(origin string, aliases []string) Synonym {
	return Synonym{
		key:     origin,
		aliases: aliases,
	}
}

// NewListSynonym .
func NewListSynonym(aliases []string) Synonym {
	return Synonym{
		aliases: aliases,
	}
}

// MarshalJSON synonym
func (s Synonym) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer

	b.WriteString("\"")
	b.WriteString(strings.Join(s.aliases, ","))
	if s.key != "" {
		b.WriteString(" => ")
		b.WriteString(s.key)
	}
	b.WriteString("\"")

	return b.Bytes(), nil
}
