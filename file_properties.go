package properties

import (
	"golang.org/x/text/encoding"
)

var fs fileSystem = osFS{}

func envsFromFile(path string, encoder *encoding.Encoder) (map[string]string, error) {
	lines, err := loadLines(path, encoder, fs)
	if err != nil {
		return nil, err
	}

	props, err := toMap(lines)
	if err != nil {
		return nil, err
	}

	return props, nil
}
