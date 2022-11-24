package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Id_ente    string `json:"en"`
	Id_prod    string `json:"pr"`
	Id_orga    string `json:"og"`
	Id_user    string `json:"us"`
	Id_other   string `json:"ot"`
	Modules    string `json:"mod"`
	Cargo      int64  `json:"cargo"`
	Date       string `json:"date"`
	Month      int64  `json:"month"`
	Year       int64  `json:"year"`
	Branch     string `json:"branch"`
	StoreHouse string `json:"storehouse"`
	jwt.StandardClaims
}

func ValidateToken(signedToken string) (body JWTClaim, err error) {
	mySigningKey := GetKey_PrivateJwt()
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("Error en el token")
		return
	}
	body = *claims
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token a expirado")
		return
	}
	return
}

func GetValToken(r *http.Request, key string) interface{} {
	token := r.Header.Get("access-token")
	token_verify, _ := ValidateToken(token)
	var myMap map[string]interface{}
	data, _ := json.Marshal(token_verify)

	json.Unmarshal(data, &myMap)
	// fmt.Println(myMap)
	return myMap[key]

}
