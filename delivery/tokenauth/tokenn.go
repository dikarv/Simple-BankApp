package tokenauth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var ApplicationName = "BANK"
var JwtSigningMethod = jwt.SigningMethodHS256
var JwtSignatureKey = []byte("STAUFFENBERG")

type MyClaims struct {
	jwt.StandardClaims
	AccountNumber string `json:"account_number"`
	Password      string `json:"user_password"`
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type Credential struct {
	AccountNumber string `json:"account_number"`
	Password      string `json:"user_password"`
}

func GenerateToken(accountNumber, password string) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: ApplicationName,
		},
		AccountNumber: accountNumber,
		Password:      password,
	}
	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)
	return token.SignedString(JwtSignatureKey)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != JwtSigningMethod {
			return nil, fmt.Errorf("Signing method invalid")
		}
		return JwtSignatureKey, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

func AuthTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/bank/login" {
			c.Next()
		} else {
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(401, gin.H{
					"message": "unauthorized",
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			fmt.Println(tokenString)
			if tokenString == "" {
				c.JSON(401, gin.H{
					"message": "unauthorized",
				})
				c.Abort()
				return
			}
			token, err := ParseToken(tokenString)
			if err != nil {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			fmt.Println(token)
			if token["iss"] == ApplicationName {
				c.Next()
			} else {
				c.JSON(401, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return

			}
		}
	}
}
