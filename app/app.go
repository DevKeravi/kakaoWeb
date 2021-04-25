package app

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/garyburd/go-oauth/oauth"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func indexHandler(c *gin.Context) {
	c.Redirect(http.StatusTemporaryRedirect, "/login.html")
	c.Abort()
}
func loginHandler(c *gin.Context) {
	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}
func callbackhandler(c *gin.Context) {
	str := c.Request.FormValue("state")
	if str != state {
		log.Printf("invaild oauth state, expected '%s', got '%s'\n", state, str)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		c.Abort()
		return
	}

	code := c.Request.FormValue("code")

	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx := context.WithValue(oauth2.NoContext, oauth.HTTPClient, httpClient)

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Printf("conf.Exchange() failed with %s \n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		c.Abort()
		return
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://kapi.kakao.com/v1/api/talk/profile", nil)
	req.Header.Set("Host", "kapi.kakao.com")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-type", "application/x-www-form-urlencoded;charset=utf-8")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ReadAll:", err)
		return
	}
	log.Println("parseResponseBody: ")
	log.Println(string(body))

	c.Redirect(http.StatusTemporaryRedirect, "/")
	c.Abort()
}
func Init() {
	rand.Seed(time.Now().UnixNano())
}

func Run(addr string) {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./public/", false)))

	r.GET("/", indexHandler)
	r.GET("/login", loginHandler)
	r.GET("/auth", callbackhandler)

	r.Run(addr)
}
