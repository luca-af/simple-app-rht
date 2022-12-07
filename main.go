package main

import (
	"flag"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func setEnv(key, def string) string {
	value, found := os.LookupEnv(key)
	if !found {
		value = def
	}
	return value
}

func IntensiveCalculation(num int) {
	var result int
	result = 1
	for i := 1; i < num; i++ {
		result = i * result
		println(result)
	}

}

func main() {
	r := gin.Default()
	cat_emoji := "🐈"
	catFlag := flag.Bool("cat", true, "a bool")
	flag.Parse()
	enable_msg := setEnv("FLAG_MSG", "false")
	mode := setEnv("GIN_MODE", "debug")
	msg := setEnv("MSG_CUSTOM", "Hello, World!")
	host := setEnv("HOSTNAME", "ERROR")
	static_dir := setEnv("GIN_STATIC_DIR", "./static")

	r.Static("/static", static_dir)          // serve static files from public folder
	r.GET("/healthz", func(c *gin.Context) { // healthz endpoint
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy!",
		})
	})

	r.GET("/cat", func(c *gin.Context) { // cat endpoint
		if *catFlag {
			c.JSON(http.StatusOK, gin.H{
				"message": "You found the cat! " + cat_emoji,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "No cat here!",
			})
		}
	})
	r.GET("/cats/:number", func(c *gin.Context) { // cats endpoint
		number, err := strconv.Atoi(c.Param("number"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error!",
			})
		} else {
			if *catFlag {
				c.JSON(http.StatusOK, gin.H{
					"message": "You found a lot of cats! " + strings.Repeat(cat_emoji, number),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message": "No cat here!",
				})
			}
		}
	})

	r.GET("/hello", func(c *gin.Context) { // hello endpoint
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World! Workgin in " + mode,
			"host":    host,
		})
	})

	r.GET("/help", func(c *gin.Context) { // hello endpoint
		c.JSON(http.StatusOK, gin.H{
			"/cat":          "Show a cat (if enabled)",
			"/cats/:number": "Show a lot of cats (if enabled)",
			"/hello":        "Show a hello message and the hostname",
			"/healthz":      "Show a healthy message",
			"/msg":          "Show a custom message (if enabled)",
			"/help":         "Show this message",
		})
	})

	r.GET("/msg", func(c *gin.Context) { // msg endpoint
		if enable_msg == "true" {
			c.JSON(http.StatusOK, gin.H{
				"message": msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Message disabled",
			})
		}

	})

	r.Run() // listen and serve on port 8080
}
