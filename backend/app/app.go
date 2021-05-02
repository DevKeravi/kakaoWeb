package app

import (
	"io/ioutil"
	"kakaoWeb/backend/model"
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

func loginHandler(c *gin.Context) {
	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func callbackhandler(c *gin.Context) {
	str := c.Request.FormValue("state")
	if str != state {
		log.Printf("invaild oauth state, expected '%s', got '%s'\n", state, str)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	code := c.Request.FormValue("code")

	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx := context.WithValue(oauth2.NoContext, oauth.HTTPClient, httpClient)

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Printf("conf.Exchange() failed with %s \n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
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

	model.Create(body)

	c.SetCookie("kakaoAuth", string(body), 3600, "/", "", false, false)
	c.Redirect(http.StatusTemporaryRedirect, "/")

}
func Init() {
	rand.Seed(time.Now().UnixNano())
}

func indexHandler(c *gin.Context) {
	cookie, err := c.Cookie("kakaoAuth")
	if err != nil {
		c.JSON(http.StatusOK, "")
		return
	}

	data := model.NewData(cookie)
	c.JSON(http.StatusOK, model.Get(data.Name))
}

func Run(addr string) {
	r := gin.Default()
	model.Init()
	r.LoadHTMLGlob("../frontend/kakaoweb/public/*")
	r.Use(static.Serve("/", static.LocalFile("../frontend/kakaoweb/public/", false)))

	api := r.Group("api")
	{
		api.GET("/login", loginHandler)
		api.GET("/auth", callbackhandler)
		api.GET("/index", indexHandler)
	}

	r.Run(addr)
}
