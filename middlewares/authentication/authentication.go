package authentication

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

/* --- JWT ---*/
/* This file contains configuration of JWT purposed for API autorization */
type JwtCustomClaim struct {
	Identity string
	jwt.StandardClaims
}

var (
	Jwtsecretkey = []byte("mysecretkey")
	JwtClaim     = JwtCustomClaim{
		"customidentity",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "test",
		},
	}
)

func Authentication(c *gin.Context) {
	fmt.Println("... Authentication middleware hit")
	// This is just DUMMY token
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYmYiOjE0NDQ0Nzg0MDB9.u1riaD1rW97opCoAuRCTy4w58Br-Zk-bh7vLiRIsrpU"

	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return Jwtsecretkey, nil
	})

	c.Next()

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
		c.Next()
	} else {
		if err != nil {
			//panic(err)
		}
		//c.JSON(404, err)
		//c.Abort()
		c.Next()
	}
}
