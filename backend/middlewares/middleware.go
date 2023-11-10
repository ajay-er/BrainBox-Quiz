package middlewares

import (
	"backend/helpers"
	"backend/response"
	"fmt"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the JWT token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		tokenString := helpers.GetTokenFromHeader(authHeader)

		// Validate the token and extract the user ID
		if tokenString == "" {
			var err error
			tokenString, err = c.Cookie("Authorization")
			if err != nil {

				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}
		userID, userEmail, err := helpers.ExtractUserIDFromToken(tokenString)
		if err != nil {
			fmt.Println("error is ", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// Add the user ID to the Gin context
		c.Set("user_id", userID)
		c.Set("user_email", userEmail)

		// Call the next handler
		c.Next()
	}
}

func AuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")
		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return

		}
		tokenpart := splitted[1]
		tokenClaims, err := helpers.ValidateToken(tokenpart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token ", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return

		}
		// All good, set the Claims in the context
		c.Set("tokenClaims", tokenClaims)

		c.Next()

	}

}
