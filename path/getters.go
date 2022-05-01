package path

import (
	"errors"
	
	"github.com/TheAlyxGreen/Wrench/data"
)

// GetKeys fetches the list of Data in the path
func GetKeys(p Path) []data.Data {
	return p.keys
}

// PeekKey returns the last Data in the path without removing it
func PeekKey(p Path) (data.Data, error) {
	numberOfKeys := len(p.keys)
	if numberOfKeys > 0 {
		return p.keys[numberOfKeys-1], nil
	} else {
		return data.Data{}, errors.New(ErrorPathHasNoKeys)
	}
}

// PopKey pops the last Data out of the path and returns it
func PopKey(p *Path) (data.Data, error) {
	numberOfKeys := len(p.keys)
	if numberOfKeys > 0 {
		out := p.keys[numberOfKeys-1]
		p.keys = p.keys[0 : numberOfKeys-2]
		return out, nil
	} else {
		return data.Data{}, errors.New(ErrorPathHasNoKeys)
	}
}
