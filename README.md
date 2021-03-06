# gin-hydra
[Hydra](https://github.com/ory-am/hydra) middleware for [Gin](https://gin-gonic.github.io/gin/) framework.
It uses Hydra's API to extract and validate auth token.

## Install
``` bash
go get github.com/janekolszak/gin-hydra
```

## Example

``` go
import (
    "github.com/gin-gonic/gin"
    "github.com/ory-am/hydra/firewall"
  	hydra "github.com/ory-am/hydra/sdk"

    gh "go get github.com/janekolszak/gin-hydra"

)

func handler(c *gin.Context) {
	ctx := c.Get("hydra").(*firewall.Context)
	// Now you can access ctx.Subject etc.
}

func main(){
	// Initialize Hydra and gin-hydra
	hc, err := hydra.Connect(
		hydra.ClientID("..."),
		hydra.ClientSecret("..."),
		hydra.ClusterURL("..."),
	)

	if err != nil {
		panic(err)
	}

	gh.Init(hc)

	// Use the middleware
 	router := gin.Default()
	router.GET("/", gh.ScopesRequired("scope1", "scope2"), handler)
	router.Run()
}
```
