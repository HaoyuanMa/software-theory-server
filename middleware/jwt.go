package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
)

type JWT struct {
	JwtKey []byte
}

func NewJWT() *JWT {
	return &JWT{JwtKey: []byte("jsldkjfeljsflskjf")}
}

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//fmt.Println("in auth")
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "request without token",
			})
			c.Abort()
			return
		}

		checkToken := strings.Split(tokenHeader, " ")

		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "token error",
			})
			c.Abort()
			return
		}

		var j = NewJWT()
		var claims, ok = j.ParserToken(checkToken[1])
		if !ok {
			c.JSON(401, gin.H{
				"status":  401,
				"message": "token error",
			})
			c.Abort()
		} else {
			c.Set("username", claims)
			c.Next()
		}

	}
}

func (j *JWT) ParserToken(tokenString string) (*MyClaims, bool) {
	var token, err = jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.JwtKey, nil
	})
	if err != nil || token == nil {
		return nil, false
	}
	var claims, ok = token.Claims.(*MyClaims)
	if ok && token.Valid {
		return claims, true
	}
	return nil, false
}

func (j *JWT) CreateToken(claims MyClaims) (string, bool) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var tokenString, err = token.SignedString(j.JwtKey)
	if err != nil {
		return "", false
	}
	return tokenString, true
}
