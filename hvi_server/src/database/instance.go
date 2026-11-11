package database

type Instance interface {
	Open() error
	Close() error
}
