package mpr

type MinimalPlacesResponse interface {
	Id() int64
	URI() string
	Repo() string
}
