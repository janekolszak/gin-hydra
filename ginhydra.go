package ginhydra

import (
	hydra "github.com/ory-am/hydra/sdk"

	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	hc *hydra.Client
)

func Init(hydraClient *hydra.Client) {
	hc = hydraClient
}

// TODO: Copied from fosite
func AccessTokenFromRequest(req *http.Request) string {
	auth := req.Header.Get("Authorization")
	split := strings.SplitN(auth, " ", 2)
	if len(split) != 2 || !strings.EqualFold(split[0], "bearer") {
		// Empty string returned if there's no such parameter
		err := req.ParseForm()
		if err != nil {
			return ""
		}
		fmt.Println(req.Form.Get("access_token"))
		return req.Form.Get("access_token")
	}
	return split[1]
}

func ScopesRequired(scopes ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, err := hc.Introspection.IntrospectToken(c, AccessTokenFromRequest(c.Request), scopes...)
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
