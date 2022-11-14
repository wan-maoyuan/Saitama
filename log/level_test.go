package log

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLevelString(t *testing.T) {
	require.Equal(t, "DEBUG", LevelDebug.String())
	require.Equal(t, "INFO", LevelInfo.String())
	require.Equal(t, "WARN", LevelWarn.String())
	require.Equal(t, "ERROR", LevelError.String())
	require.Equal(t, "FATAL", LevelFatal.String())
}

func TestParseLevel(t *testing.T) {
	require.Equal(t, LevelDebug, ParseLevel("DEBUG"))
	require.Equal(t, LevelInfo, ParseLevel("INFO"))
	require.Equal(t, LevelWarn, ParseLevel("WARN"))
	require.Equal(t, LevelError, ParseLevel("ERROR"))
	require.Equal(t, LevelFatal, ParseLevel("FATAL"))
}
