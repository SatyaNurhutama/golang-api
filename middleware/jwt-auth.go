package middleware

import (
	"golang_api/helper"
	"golang_api/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//validates the token user given, return 401 if not valid
func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]: ", claims["user_id"])
			log.Println("Claim[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			respose := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respose)
		}
	}
}
