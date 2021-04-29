package app

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"os"

	"golang.org/x/oauth2"
)

var (
	state = createHash()
	conf  = &oauth2.Config{
		ClientID:     os.Getenv("KAKAO_RESTAPI_KEY"),
		ClientSecret: os.Getenv("KAKAO_CLIENT_SECRET"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://kauth.kakao.com/oauth/authorize",
			TokenURL: "https://kauth.kakao.com/oauth/token",
		},
		RedirectURL: os.Getenv("KAKAO_REDIRECT_URL"),
	}
)

func RandomString(n int) string {
	var letters = []rune("absdjfkaQWEASSDFDSCVRUULLsjaeiflcmdfkl")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)

}

func createHash() string {

	data := RandomString(rand.Intn(10) + 5)
	hash := sha256.New()

	hash.Write([]byte(data))

	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)

	return mdStr

}
