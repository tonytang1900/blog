package middleware

import (
	"fmt"
	"testing"
)

func TestGenTokn(t *testing.T) {
	token, _ := GenToken("tony", "123123")
	fmt.Println(token)

	claims, _ := ValidateToken(token)

	fmt.Println("claims:", claims)
	//fmt.Println("claims:", claims.(*MyClaims).Username)

}

