package config

import (
	"strings"
)

type Option func(*options)

type Decoder func(*KeyValue, map[string]any) error

type options struct {
	source  Source
	decoder Decoder
}

func WithSource(s Source) Option {
	return func(o *options) {
		o.source = s
	}
}

func WithDecoder(d Decoder) Option {
	return func(o *options) {
		o.decoder = d
	}
}

func defaultDecoder(src *KeyValue, target map[string]any) error {
	if src.Format == "" {
		// expand key "aa.bb" into map[aa][bb]any
		keys := strings.Split(src.Key, ".")
		for i, k := range keys {
			if i == len(keys)-1 {
				target[k] = src.Value
			} else {
				sub := make(map[string]any)
				target[k] = sub
				target = sub
			}
		}

		return nil
	}

	return nil
}
