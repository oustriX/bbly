package pg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDBConnectionConfig(t *testing.T) {
	_, err := getDBConnectionConfig()
	require.NoError(t, err)
}
