package ginhydra

import (
	hydra "github.com/ory-am/hydra/sdk"

	"github.com/gin-gonic/gin"
)

var (
	hc *hydra.Client
)

func Init(hydraClient *hydra.Client) {
	hc = hydraClient
}

func ScopesRequired(scopes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, err := hc.Warden.TokenValid(c, hc.Warden.TokenFromRequest(c.Request), scopes...)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		// All required scopes are found
		c.Set("hydra", ctx)
		c.Next()
	}
}
