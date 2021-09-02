package mpr

import (
	"encoding/json"
	"io"
)

// JsonMPR implements the MPR interface for a JSON-encoded representation of a "minimal places response" (MPR).
type JsonMPR struct {
	MPR
	// A valid Who's On First ID for a record contained in a MRP response.	
	JsonId   int64  `json:"id"`
	// A valid Who's On First style URL for a record contained in a MRP response.	
	JsonURI  string `json:"uri"`
	// The name of the (Git) respository that the record is stored in.	
	JsonRepo string `json:"repo"`
}

// Return the Who's On First ID for 'm'.
func (m *JsonMPR) Id() int64 {
	return m.JsonId
}

// Return the Who's On First URI for 'm'
func (m *JsonMPR) URI() string {
	return m.JsonURI
}

// Return the name of the (Git) respository that the ID associated with 'm' is stored in.
func (m *JsonMPR) Repo() string {
	return m.JsonRepo
}

// MarshalJsonMPR will write a JSON-encoded serialization of 'm' to 'wr'.
func MarshalJsonMPR(m *JsonMPR, wr io.Writer) error {
	enc := json.NewEncoder(wr)
	return enc.Encode(m)
}

// UnmarshalJsonMPR will parse the body of 'r' in to a JsonMPR instance.
func UnmarshalJsonMPR(r io.Reader) (*JsonMPR, error) {

	var m *JsonMPR
	dec := json.NewDecoder(r)
	err := dec.Decode(&m)

	if err != nil {
		return nil, err
	}

	return m, nil
}
