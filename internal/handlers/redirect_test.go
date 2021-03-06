package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetId(t *testing.T) {
	ctx := gin.Context{
		Params: gin.Params{gin.Param{Key: "id", Value: "boba"}},
	}
	id := getID(&ctx)
	emptyString := getID(&gin.Context{})
	require.Equal(t, "boba", id)
	require.Equal(t, "", emptyString)
}
