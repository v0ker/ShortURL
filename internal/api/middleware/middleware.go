package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRecovery,
)

func abortRequest(c *gin.Context, err string) {
	c.Redirect(302, "/error?msg="+err)
}
