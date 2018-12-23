package properties

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

func Read(file string, encoder *encoding.Encoder) (*Properties, error) {
	if encoder == nil {
		encoder = unicode.UTF8.NewEncoder()
	}
	fileProps, err := envsFromFile(file, encoder)
	if err != nil {
		return nil, err
	}
	envProps, err := envsFromProperties()
	if err != nil {
		return nil, err
	}

	return &Properties{mergeProperties(fileProps, envProps)}, nil
}

func mergeProperties(someProp, otherProp map[string]string) map[string]string {
	for k, v := range otherProp {
		someProp[k] = v
	}

	return someProp
}
