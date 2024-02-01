package jwtt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kanelinweihao/lwhFrameGo/app/conf"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
	"time"
)

type EntityJwtClaims struct {
	Data string `json:"data"`
	*jwt.RegisteredClaims
}

func (self *EntityJwtClaims) SetData(jsonAuth string) *EntityJwtClaims {
	self.Data = jsonAuth
	return self
}

func (self *EntityJwtClaims) GetData() (jsonAuth string) {
	jsonAuth = self.Data
	return jsonAuth
}

/*
Encode
*/

func GetJwtToken(jsonAuth string) (jwtToken string) {
	entityJwtClaims := getEntityJwtClaimsByJsonAuth(jsonAuth)
	entityJwtToken := getEntityJwtTokenByEntityJwtClaims(entityJwtClaims)
	jwtSecret := conf.GetJwtSecret()
	jwtToken, errToken := entityJwtToken.SignedString(jwtSecret)
	err.ErrCheck(errToken)
	return jwtToken
}

func getEntityJwtClaimsByJsonAuth(jsonAuth string) (entityJwtClaims *EntityJwtClaims) {
	entityJwtClaims = new(EntityJwtClaims)
	entityJwtClaims.SetData(jsonAuth)
	c := getRegisteredClaims()
	entityJwtClaims.RegisteredClaims = c
	return entityJwtClaims
}

func getRegisteredClaims() (c *jwt.RegisteredClaims) {
	iss := conf.GetJwtIssuer()
	sub := conf.GetJwtSubject()
	h := conf.GetJwtHoursExpires()
	exp := jwt.NewNumericDate(time.Now().Add(h * time.Hour))
	npf := jwt.NewNumericDate(time.Now())
	iat := jwt.NewNumericDate(time.Now())
	jti := conf.GetJwtID()
	c = new(jwt.RegisteredClaims)
	c.Issuer = iss
	c.Subject = sub
	c.ExpiresAt = exp
	c.NotBefore = npf
	c.IssuedAt = iat
	c.ID = jti
	return c
}

func getEntityJwtTokenByEntityJwtClaims(entityJwtClaims *EntityJwtClaims) (entityJwtToken *jwt.Token) {
	entityJwtToken = jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		entityJwtClaims)
	return entityJwtToken
}

/*
Decode
*/

func GetJsonAuth(jwtToken string) (jsonAuth string) {
	entityJwtClaims := new(EntityJwtClaims)
	entityJwtToken, errParse := jwt.ParseWithClaims(
		jwtToken,
		entityJwtClaims,
		funcSecret())
	err.ErrCheck(errParse)
	entityJwtClaims, ok := entityJwtToken.Claims.(*EntityJwtClaims)
	if !ok {
		msgError := fmt.Sprintf(
			"The type of |entityJwtClaims| is invalid")
		err.ErrPanic(msgError)
	}
	isValid := entityJwtToken.Valid
	if !isValid {
		msgError := fmt.Sprintf(
			"The jwtToken |%s| is invalid",
			jwtToken)
		err.ErrPanic(msgError)
	}
	jsonAuth = entityJwtClaims.GetData()
	return jsonAuth
}

func funcSecret() jwt.Keyfunc {
	return func(entityToken *jwt.Token) (interface{}, error) {
		jwtSecret := conf.GetJwtSecret()
		return jwtSecret, nil
	}
}
