package properties

import (
	"strings"
	"errors"
	"fmt"
)

const separator = "="

func toMap(lines []string) (map[string]string, error) {
	props := make(map[string]string)

	for _, line := range lines {
		if shouldSkip(line) {
			continue
		}
		k, v, err := splitProperty(line)
		if err != nil {
			return nil, err
		}

		props[k] = v
	}

	return props, nil
}

func splitProperty(property string) (string, string, error) {
	p := strings.SplitN(property, separator, 2)
	if len(p) != 2 {
		return "", "", errors.New(fmt.Sprintf("cannot split value '%s' by '%s' separator", property, separator))
	}

	return strings.TrimSpace(p[0]), strings.TrimSpace(p[1]), nil
}
