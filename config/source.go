package config

type KeyValue struct {
	Key    string
	Value  []byte
	Format string
}

type Source interface {
	Load() ([]*KeyValue, error)
	Watch() (Watcher, error)
}

type Watcher interface {
	Next() ([]*KeyValue, error)
	Stop() error
}
