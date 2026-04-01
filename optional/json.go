package optional

import (
	"bytes"
	"encoding/json"
)

// MarshalJSON encodes marshals the value if present or 'null' if empty.
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.present {
		return []byte("null"), nil
	}

	return json.Marshal(o.value)
}

// UnmarshalJSON creates a present Optional if the value is not 'null',
// otherwise an empty.
func (o *Optional[T]) UnmarshalJSON(d []byte) error {
	if bytes.Equal(d, []byte("null")) {
		var t T
		o.value = t
		o.present = false

		return nil
	}

	err := json.Unmarshal(d, &o.value)
	if err != nil {
		return err
	}

	o.present = true

	return nil
}
