package json

import "github.com/goravel/framework/foundation"

// Marshal serializes the given value to a JSON-encoded byte slice.
func Marshal(v any) ([]byte, error) {
	return foundation.App.GetJson().Marshal(v)
}

// Unmarshal deserializes the given JSON-encoded byte slice into the provided value.
func Unmarshal(data []byte, v any) error {
	return foundation.App.GetJson().Unmarshal(data, v)
}

// MarshalString serializes the given value to a JSON-encoded string.
func MarshalString(v any) (string, error) {
	s, err := Marshal(v)
	return string(s), err
}

// UnmarshalString deserializes the given JSON-encoded string into the provided value.
func UnmarshalString(data string, v any) error {
	return Unmarshal([]byte(data), v)
}
