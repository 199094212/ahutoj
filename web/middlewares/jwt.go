package middlewares

import (
	"ahutoj/web/utils"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

var sign []byte
var ExpTime time.Duration

// MyClaims 自定义声明结构体并内嵌jwt.StandardClaims
// jwt包自带的jwt.StandardClaims只包含了官方字段
// 我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type MyClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func InitJwt() {
	config := utils.GetInstance()
	sign = []byte(config.Sign)
	ExpTime = 24 * time.Hour
}

func GetToken(userID string) (string, error) {
	// 创建一个我们自己的声明的数据
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpTime).Unix(), // 过期时间
			Issuer:    "ahutoj",                       // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(sign)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	var myclaims = new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, myclaims, func(token *jwt.Token) (i interface{}, err error) {
		return sign, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return myclaims, nil
	}
	return nil, errors.New("invalid token")
}
