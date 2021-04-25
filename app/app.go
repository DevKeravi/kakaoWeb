package app

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
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

func indexHandler(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/login.html")
	c.Abort()
}
func loginHandler(c *gin.Context) {
	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
func Init() {
	rand.Seed(time.Now().UnixNano())
}

func Run(addr string) {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public/", false)))

	r.GET("/", indexHandler)
	r.GET("/login", loginHandler)

	r.Run(addr)
}
