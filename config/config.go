package config

type Config interface {
	Load() error
	Scan(v any) error
	Value(key string) Value
	Close() error
}
