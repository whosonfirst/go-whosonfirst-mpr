package mpr

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

func TestJSONMPR(t *testing.T) {

	raw := `{"id": 102527513, "uri": "102/527/513/102527513.geojson", "repo": "whosonfirst-data-admin-us"}`

	r := strings.NewReader(raw)

	m, err := UnmarshalJsonMPR(r)

	if err != nil {
		t.Fatalf("Failed to unmarshal MPR, %v", err)
	}

	if m.Id() != 102527513 {
		t.Fatalf("Invalid ID field")
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err = MarshalJsonMPR(m, wr)

	if err != nil {
		t.Fatalf("Failed to marshal MPR, %v", err)
	}

	if bytes.Equal(buf.Bytes(), []byte(raw)) {
		t.Fatalf("Output does not match input")
	}
}
