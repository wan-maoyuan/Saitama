package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultDecoder(t *testing.T) {
	src := &KeyValue{
		Key:    "aa.bb.cc",
		Value:  []byte("hello world"),
		Format: "",
	}

	target := make(map[string]any)
	require.Equal(t, nil, defaultDecoder(src, target))
}
