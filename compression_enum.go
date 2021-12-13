// Code generated by "enumer -transform snake_upper -type Compression -trimprefix Compression -output compression_enum.go"; DO NOT EDIT.

package ch

import (
	"fmt"
	"strings"
)

const _CompressionName = "DISABLEDLZ4"

var _CompressionIndex = [...]uint8{0, 8, 11}

const _CompressionLowerName = "disabledlz4"

func (i Compression) String() string {
	if i >= Compression(len(_CompressionIndex)-1) {
		return fmt.Sprintf("Compression(%d)", i)
	}
	return _CompressionName[_CompressionIndex[i]:_CompressionIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _CompressionNoOp() {
	var x [1]struct{}
	_ = x[CompressionDisabled-(0)]
	_ = x[CompressionLZ4-(1)]
}

var _CompressionValues = []Compression{CompressionDisabled, CompressionLZ4}

var _CompressionNameToValueMap = map[string]Compression{
	_CompressionName[0:8]:       CompressionDisabled,
	_CompressionLowerName[0:8]:  CompressionDisabled,
	_CompressionName[8:11]:      CompressionLZ4,
	_CompressionLowerName[8:11]: CompressionLZ4,
}

var _CompressionNames = []string{
	_CompressionName[0:8],
	_CompressionName[8:11],
}

// CompressionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CompressionString(s string) (Compression, error) {
	if val, ok := _CompressionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _CompressionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Compression values", s)
}

// CompressionValues returns all values of the enum
func CompressionValues() []Compression {
	return _CompressionValues
}

// CompressionStrings returns a slice of all String values of the enum
func CompressionStrings() []string {
	strs := make([]string, len(_CompressionNames))
	copy(strs, _CompressionNames)
	return strs
}

// IsACompression returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Compression) IsACompression() bool {
	for _, v := range _CompressionValues {
		if i == v {
			return true
		}
	}
	return false
}
