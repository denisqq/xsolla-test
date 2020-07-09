package security

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"strings"
)

type BasicAuth struct {
	Authenticate Authenticate
}

func NewBasicAuth(authenticate Authenticate) BasicAuth {
	return BasicAuth{Authenticate: authenticate}
}

func (ba BasicAuth) BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.Request.Header.Get("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			respondWithError(401, "Unauthorized", c)
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		username := pair[0]
		password := pair[1]

		if len(pair) != 2 || !ba.Authenticate.authenticateUser(username, password) {
			respondWithError(401, "Unauthorized", c)
			return
		}
		c.Set(gin.AuthUserKey, username)
	}
}

type Authenticate struct {
	BasicAuthUserDetailsService BasicAuthUserDetailsService
}

func NewAuthenticate(basicAuthUserDetailsService BasicAuthUserDetailsService) Authenticate {
	return Authenticate{BasicAuthUserDetailsService: basicAuthUserDetailsService}
}

func (a Authenticate) authenticateUser(username, password string) bool {
	user, err := a.BasicAuthUserDetailsService.LoadUserDetails(username)
	if err != nil {
		return false
	}

	return user.Password == password
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}
