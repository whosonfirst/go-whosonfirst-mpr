package mpr

type MPR interface {
	Id() int64
	URI() string
	Repo() string
}
