package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"gogin/pkg/errno"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

var (
	Secret     = "testxxx" // 加盐
	ExpireTime = 3600      // token有效期

)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const (
	ErrorReason_ServerBusy = "服务器繁忙"
	ErrorReason_ReLogin    = "请重新登陆"
)

type JWTClaims struct {
	// token里面添加用户信息，验证token后可能会用到用户信息
	jwt.StandardClaims
	Password    string   `json:"password"`
	Username    string   `json:"username"`
	Permissions []string `json:"permissions"`
}

func Register(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, message)
}

func Login(c *gin.Context) {
	user := User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
	}

	claims := &JWTClaims{
		Username:    user.Username,
		Password:    user.Password,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()

	signedToken, err := getToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)
}

func getToken(claims *JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", errors.New(ErrorReason_ServerBusy)
	}
	return signedToken, nil

}

func Verify(c *gin.Context) {
	strToken := c.Request.Header.Get("AuthToken")
	claim, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "verify,", claim.Username)
}

func Refresh(c *gin.Context) {
	strToken := c.Request.Header.Get("AuthToken")
	claims, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	signedToken, err := getToken(claims)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, signedToken)

}

func verifyAction(strToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return nil, errors.New("token failure.")
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ErrorReason_ReLogin)
	}
	fmt.Println("verify")
	return claims, nil

}

func SayHello(c *gin.Context) {
	strToken := c.Request.Header.Get("AuthToken")

	claim, err := verifyAction(strToken)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.String(http.StatusOK, "hello,", claim.Username)

}
