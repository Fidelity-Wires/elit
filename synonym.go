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

// NewSynonym .
func NewSynonym(origin string, aliases []string) Synonym {
	return Synonym{
		key:     origin,
		aliases: aliases,
	}
}

// MarshalJSON synonym
func (s Synonym) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer

	b.WriteString(strings.Join(s.aliases, ","))
	b.WriteString(" => ")
	b.WriteString(s.key)

	return b.Bytes(), nil
}
