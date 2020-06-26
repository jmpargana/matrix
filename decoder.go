package matrix

import (
	"bytes"
	"encoding/gob"
)

// MarshalBinary is the method needed to implement the BinaryMarshaler interface.
// With it one can call gob.Encode on the Matrix struct.
func (m *Matrix) MarshalBinary() ([]byte, error) {
	w := wrapMatrix{m.NumRows, m.NumCols, m.data}

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(w); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalBinary is the method needed to implement the BinaryUnmarshaler interface.
// With it one can call gob.Decode on the Matrix struct.
func (m *Matrix) UnmarshalBinary(data []byte) error {
	w := wrapMatrix{}

	reader := bytes.NewReader(data)
	dec := gob.NewDecoder(reader)
	if err := dec.Decode(&w); err != nil {
		return err
	}

	m.NumRows = w.NumRows
	m.NumCols = w.NumCols
	m.data = w.Data

	return nil
}
