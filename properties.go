package properties

import (
	"errors"
	"fmt"
	"strconv"
)

type Properties struct {
	props map[string]string
}

func (p *Properties) GetString(name string) (string, error) {
	return p.get(name)
}

func (p *Properties) GetInt(name string) (int, error) {
	v, err := p.get(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(v)
}

func (p *Properties) GetBool(name string) (bool, error) {
	v, err := p.get(name)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(v)
}

func (p *Properties) GetFloat(name string) (float64, error) {
	v, err := p.get(name)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(v, 64)
}

func (p *Properties) Contains(name string) bool {
	_, ok := p.props[name]

	return ok
}

func (p *Properties) get(name string) (string, error) {
	v, ok := p.props[name]
	if !ok {
		return "", errors.New(fmt.Sprintf("cannot find value for %s", name))
	}

	return v, nil
}
