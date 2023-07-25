package auth

import (
	"fmt"
	"log"
	"time"

	_ "crypto/hmac"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type JWTClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

type Authorizor struct {
	secret []byte
	logger *log.Logger
}

func New(secret []byte, logger *log.Logger) *Authorizor {
	return &Authorizor{
		secret: secret,
		logger: logger,
	}
}

func newClaims(username, role string) jwt.Claims {
	return JWTClaims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			Id:        uuid.NewString(),
			Issuer:    "http://localhost:8080",
			Audience:  "http://localhost:8080/api/v1/",
			Subject:   username,
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
		},
	}
}

func (a *Authorizor) NewToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(username, role))
	return token.SignedString(a.secret)
}

func (a *Authorizor) verifyToken(token *jwt.Token) (signature any, err error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return a.secret, nil
}

func (a *Authorizor) VerifyTokenAndGetClaims(tokenString string) (claims *JWTClaims, err error) {
	var token *jwt.Token

	token, err = jwt.Parse(tokenString, a.verifyToken)
	if err != nil {
		a.logger.Printf("jwt.Parse(tokenString): %v\n", err)
		return
	}
	var mapClaims jwt.MapClaims
	var ok bool
	if mapClaims, ok = token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		err = fmt.Errorf("token not valid")
		a.logger.Println(err)
		return
	}
	return verifyClaims(mapClaims)
}

func verifyClaims(mapClaims jwt.MapClaims) (claims *JWTClaims, err error) {
	if err = mapClaims.Valid(); err != nil {
		return
	}
	claims = &JWTClaims{}
	for k, v := range mapClaims {
		switch k {
		case "role":
			claims.Role = v.(string)
		case "aud":
			claims.Audience = v.(string)
		case "exp":
			if i, ok := v.(int64); ok {
				claims.ExpiresAt = i
			} else if i, ok := v.(float64); ok {
				claims.ExpiresAt = int64(i)
			} else {
				err = fmt.Errorf("invalid 'exp' field'")
				return
			}
		case "jti":
			claims.Id = v.(string)
		case "iat":
			if i, ok := v.(int64); ok {
				claims.IssuedAt = i
			} else if i, ok := v.(float64); ok {
				claims.IssuedAt = int64(i)
			} else {
				err = fmt.Errorf("invalid 'iat' field'")
				return
			}
		case "iss":
			claims.Issuer = v.(string)
		case "nbf":
			if i, ok := v.(int64); ok {
				claims.NotBefore = i
			} else if i, ok := v.(float64); ok {
				claims.NotBefore = int64(i)
			} else {
				err = fmt.Errorf("invalid 'nbf' field'")
				return
			}
		case "sub":
			claims.Subject = v.(string)
		}
	}
	if len(claims.Role) == 0 {
		err = fmt.Errorf("no valid role associated with claims")
	}
	return
}
