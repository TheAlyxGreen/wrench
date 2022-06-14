package paths

import (
	"github.com/TheAlyxGreen/wrench/values"
)

// Keys fetches the list of Value in the paths
func (p Path) Keys() []values.Value {
	return p.keys
}

// PeekKey returns the last Value in the paths without removing it
func (p Path) PeekKey() (values.Value, error) {
	numberOfKeys := len(p.keys)
	if numberOfKeys > 0 {
		return p.keys[numberOfKeys-1], nil
	} else {
		return values.Value{}, ErrorPathHasNoKeys
	}
}

// PopKey pops the last Value out of the paths and returns it
func (p *Path) PopKey() (values.Value, error) {
	if len(p.keys) == 0 {
		return values.Value{}, ErrorPathHasNoKeys
	}
	out := p.keys[len(p.keys)-1]
	p.keys = p.keys[0 : len(p.keys)-1]
	return out, nil
}

func (p Path) FirstKey() (values.Value, error) {
	if len(p.keys) == 0 {
		return values.Value{}, ErrorPathHasNoKeys
	}
	return p.keys[0], nil
}
