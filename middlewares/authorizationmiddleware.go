package middlewares

import (
	"github.com/draco121/common/clients"
	"github.com/draco121/common/constants"
	"github.com/draco121/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthMiddleware(actions ...constants.Action) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get user ID from the request (you may get it from the authentication process)
		authorizationServiceApiClient := clients.NewAuthorizationServiceApiClient(os.Getenv("AUTHORIZATION_SERVICE_BASEURL"))
		requestData := models.AuthorizationInput{
			Token:   c.GetHeader("Authorization"),
			Actions: actions,
		}
		authResponse, err := authorizationServiceApiClient.Authorize(requestData)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("UserId", authResponse.UserId)
		// If the user has the required role, proceed to the next middleware/handler
		c.Next()
	}
}
