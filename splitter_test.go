package properties

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

func TestSplittingProperty(t *testing.T) {

	testCases := []struct {
		property, expectedKey, expectedValue string
	}{
		{"KEY=VALUE", "KEY", "VALUE"},
		{"someKey = someValue", "someKey", "someValue"},
	}

	for _, tc := range testCases {
		k, v, e := splitProperty(tc.property)

		if e != nil {
			t.Errorf("shouldn't emmit error: %s", e)
		}

		assert.Equal(t, k, tc.expectedKey)
		assert.Equal(t, v, tc.expectedValue)
	}
}
