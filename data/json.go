package data

import (
	"encoding/json"
	"io"
)

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(i)
}

func FromJSON(i interface{}, w io.Reader) error {
	d := json.NewDecoder(w)

	return d.Decode(i)
}
