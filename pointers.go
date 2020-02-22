package gojson

import (
	"fmt"
	"strings"
)

type PointerizerMode int

var Pointerization PointerizerMode

const (
	NoPointers PointerizerMode = iota
	ValuePointers
	StructPointers
	Everything
)

func ParsePointerizerMode(s string) (PointerizerMode, error) {
	switch strings.ToLower(s) {
	case "", "none", "nopointers", "no-pointers":
		return NoPointers, nil
	case "values", "value-pointers":
		return ValuePointers, nil
	case "structs", "struct-pointers":
		return StructPointers, nil
	case "all", "everything":
		return Everything, nil
	default:
		return 0, fmt.Errorf("unknown pointerizer mode \"%s\"", s)
	}

}
