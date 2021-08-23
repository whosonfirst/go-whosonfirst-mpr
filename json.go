package mpr

import (
	"encoding/json"
	"io"
)

type JsonMPR struct {
	MPR
	JsonId   int64  `json:"id"`
	JsonURI  string `json:"uri"`
	JsonRepo string `json:"repo"`
}

func (m *JsonMPR) Id() int64 {
	return m.JsonId
}

func (m *JsonMPR) URI() string {
	return m.JsonURI
}

func (m *JsonMPR) Repo() string {
	return m.JsonRepo
}

func MarshalJsonMPR(m *JsonMPR, wr io.Writer) error {
	enc := json.NewEncoder(wr)
	return enc.Encode(m)
}

func UnmarshalJsonMPR(r io.Reader) (*JsonMPR, error) {

	var m *JsonMPR
	dec := json.NewDecoder(r)
	err := dec.Decode(&m)

	if err != nil {
		return nil, err
	}

	return m, nil
}
