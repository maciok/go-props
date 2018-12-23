package properties

import "os"

func envsFromProperties() (map[string]string, error) {
	props := make(map[string]string)
	envs := os.Environ()

	for _, line := range envs {
		k, v, err := splitProperty(line)
		if err != nil {
			return nil, err
		}

		props[k] = v
	}

	return props, nil
}
