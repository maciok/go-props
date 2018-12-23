package properties

import (
	"golang.org/x/text/transform"
	"bufio"
	"golang.org/x/text/encoding"
)

const defaultCapacity = 10

func loadLines(path string, encoder  *encoding.Encoder, system fileSystem) ([]string, error) {
	f, err := system.Open(path)
	if err != nil {
		return nil, err
	}

	lines := make([]string, 0, defaultCapacity)
	r := transform.NewReader(f, encoder)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	if err = sc.Err(); err != nil {
		return nil, err
	}
	if err = f.Close(); err != nil {
		return nil, err
	}
	return lines, nil
}
