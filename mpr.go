package mpr

// MPR is an interface for a "minimal places response" for Who's On First related queries.
type MPR interface {
	// A valid Who's On First ID for a record contained in a MRP response.
	Id() int64
	// A valid Who's On First style URL for a record contained in a MRP response.
	URI() string
	// The name of the (Git) respository that the record is stored in.
	Repo() string
}
