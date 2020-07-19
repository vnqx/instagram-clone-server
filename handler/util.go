package handler

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("secret_key")

type Claims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func handleTokenResponse(w http.ResponseWriter, userId string) {
	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	// Create the JWT claims, which include the email and expiry time.
	claims := &Claims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed in unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Sign the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string.
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		Secure:   false,
		HttpOnly: true,
		SameSite: 0,
	})
}

func verifyToken(w http.ResponseWriter, r *http.Request) (*Claims, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	// Clear the cookie.
	if !tkn.Valid {
		http.SetCookie(w, nil)
	}

	// When token is invalid, ParseWithClaims returns an error.
	if err != nil {
		return nil, err
	}


	return claims, nil
}