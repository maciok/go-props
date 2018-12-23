package properties

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var expectedProperties = []struct {
	key, value string
}{
	{"foo", "bar"},
	{"some_foo", "some_bar"},
	{"other_foo", "other_bar"},
	{"empty", ""},
}

func TestReadingProperties(t *testing.T) {
	properties, err := Read("test.properties", nil)
	if err != nil {
		t.Errorf("shouldn't emmit error: %s", err)
	}

	for _, ep := range expectedProperties {
		val, err := properties.GetString(ep.key)
		if err != nil {
			t.Errorf("shouldn't emmit error: %s", err)
		}

		assert.Equal(t, val, ep.value)
	}

}
