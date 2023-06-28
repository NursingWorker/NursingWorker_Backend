package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

var CustomSecret1 = "big_dust"

type CustomClaims1 struct {
	P string
	jwt.StandardClaims
}

var CustomSecret = "big_dust"

type CustomClaims struct {
	UserId string
	jwt.StandardClaims
}

// GenToken 生成JWT
func CreateToken(openID string) (string, error) {
	// 创建一个我们自己的声明
	expireTime := time.Now().Add(30 * 24 * time.Hour)
	claims := CustomClaims{
		openID, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "43.138.61.49", // 签名颁发者 //改
			Subject:   "nursing_work", //签名主题人 //改
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString([]byte(CustomSecret))
}

// 解析token
func GetToken(c *gin.Context) string {
	tokenString := c.GetHeader("Authorization")
	//vcalidate token formate
	if tokenString == "" {
		return ""
	}
	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		return ""
	}
	return claims.UserId
}

func ParseToken(tokenString string) (*jwt.Token, *CustomClaims, error) {
	Claims := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(CustomSecret), nil
	})
	return token, Claims, err
}
