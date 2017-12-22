package elit

import (
	"encoding/json"
	"testing"
)

func TestSynonymMarshalJSON(t *testing.T) {
	table := []struct {
		key     string
		aliases []string
		result  string
		err     bool
	}{
		{
			key:     "",
			aliases: []string{"go", "golang", "ゴー"},
			result:  "\"go,golang,ゴー\"",
			err:     false,
		},
	}

	for _, row := range table {
		var s Synonym
		if row.key == "" {
			s = NewListSynonym(row.aliases)
		} else {
			s = NewMapSynonym(row.key, row.aliases)
		}

		b, err := json.Marshal(&s)
		if !row.err && err != nil {
			t.Fatalf("json.Marshal got error: %s", err)
		}
		if row.err && err == nil {
			t.Fatalf("json.Marshal should be error but not")
		}

		if string(b) != row.result {
			t.Errorf("result expected (%s) but (%s)", row.result, b)
		}
	}
}
