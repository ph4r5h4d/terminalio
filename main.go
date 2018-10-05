package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

type ResponseStandard struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    interface{}
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, ResponseStandard{Status: 200,
			Message: "We are only as strong as we are united, as weak as we are divided."})
	})

	router.GET("/help", func(c *gin.Context) {
		commands := [9]string{"help", "me", "self", "phone", "email", "movies", "tvseries", "resume", "ip"}
		c.JSON(200, ResponseStandard{Status: 200, Message: "Here are the commands I may answer:",
			Data: commands})
	})

	router.GET("/me", func(c *gin.Context) {
		message := `
A dreamer actually, loves to read fantasy and sci-fi books, watches movies and animes, play games and like other normal organisms on earth uses O2 as an input and output CO2.(there are rumors that I feed upon caffeine, it’s true)

Loves learning new technologies and also, prefer group work.

Former Ninja who has resigned after Bigbang, currently a member of Kuiper belt security team, gaurding solar system…
`
		c.JSON(200, ResponseStandard{Status: 200, Message: strings.Replace(message, "\n", "<br/>", -1)})

	})

	router.GET("/self", func(c *gin.Context) {
		c.JSON(200,ResponseStandard{Status:200,Message:"Well there is a saying that just look into your eyes and says nothing"})
	})

	router.GET("/phone", func(c *gin.Context) {
		c.JSON(200,ResponseStandard{Status:200,Message:"You may call me while i'm not sleep: +98.902.7799.902"})
	})

	router.GET("/email", func(c *gin.Context) {
		c.JSON(200,ResponseStandard{Status:200,Message:"Well send me an email: farshad [at] nematdoust.com"})
	})

	router.GET("/movies", func(c *gin.Context) {
		c.JSON(200,ResponseStandard{Status:200,Message:"hmm, i'll let you know"})
	})

	router.GET("/tvseries", func(c *gin.Context) {
		c.JSON(200,ResponseStandard{Status:200,Message:"hmm, i'll let you know"})
	})

	router.GET("/resume", func(c *gin.Context) {
		c.JSON(200,ResponseStandard{Status:200,Message:"just take a look at my linkedin profile..."})
	})

	router.GET("/ip", func(c *gin.Context) {
		c.String(200, c.ClientIP())
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, ResponseStandard{Status: 404, Message: "I don't wanna answer you :-D"})
	})

	router.Run(":" + port)
}
