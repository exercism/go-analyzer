package sugg

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComment_MarshalJSON(t *testing.T) {
	var c = NewComment("string")
	bytes, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	var c2 = NewComment("")
	err = json.Unmarshal(bytes, c2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, c, c2)
}

func TestPlaceholderComment_MarshalJSON(t *testing.T) {
	var c = &placeholderComment{
		Comment: "testComment",
		Params: map[string]string{
			"param1": "foo",
			"param2": "bar",
		},
	}
	bytes, err := json.Marshal(c)
	if err != nil {
		t.Fatal(err)
	}

	var c2 = &placeholderComment{}
	err = json.Unmarshal(bytes, c2)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, c, c2)
}
