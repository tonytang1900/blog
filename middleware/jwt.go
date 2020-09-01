package middleware

import (
	"blog/utils"
	"blog/utils/errmsg"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var jwtkey = []byte(utils.Secretkey)

type MyClaims struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	jwt.StandardClaims
}

func GenToken(username, password string) (tokenstr string, code int) {
	expiretime := jwt.Now().Add(7 * time.Hour * 24)

	claims := MyClaims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(expiretime),
			Issuer:    "tonytang",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstr, err := token.SignedString([]byte(utils.Secretkey))
	if err != nil {
		return "", errmsg.ERROR
	}
	return tokenstr, errmsg.Success
}

func ValidateToken(tokenstr string) (claims *MyClaims, code int) {
	token, _ := jwt.ParseWithClaims(tokenstr, &MyClaims{}, func(*jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})

	if token.Valid {
		claims = token.Claims.(*MyClaims)
		code = errmsg.Success
		return
	}
	return nil, errmsg.ERROR
}

//gin中间件，验证jwt
func VerifyJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerValue := c.GetHeader("authorization")
		if headerValue == "" {
			c.JSON(http.StatusOK, gin.H{
				"code" : errmsg.ERROR_TOKEN_NOT_EXIST,
				"msg" : errmsg.Errmsg(errmsg.ERROR_TOKEN_NOT_EXIST),
			})
			c.Abort()
		}

		sects := strings.SplitN(headerValue, " ", 2)
		if len(sects) != 2 && sects[0] != "Bearer" {
			c.JSON(http.StatusOK, gin.H{
				"code" : errmsg.ERROR_TOKEN_WRONG,
				"msg" : errmsg.Errmsg(errmsg.ERROR_TOKEN_WRONG),
			})
			c.Abort()
		}

		cs, _ := ValidateToken(sects[1])

		if cs.ExpiresAt.Unix() < jwt.Now().Unix() {
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
		}


		c.Set("username", cs.Username)
		c.Set("password", cs.Password)
		c.Next()
	}
}