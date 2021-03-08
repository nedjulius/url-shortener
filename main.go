package main

import (
	"net/http"
	"net/url"
	"strings"

	"./db"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	_ "github.com/joho/godotenv/autoload"
)

var base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var baseLength = len(base)

func reverse(target string) string {
	r := []rune(target)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func decode(encoding string) int {
	numID := 0

	for i := 0; i < len(encoding); i++ {
		numID = numID*baseLength + strings.Index(base, string(encoding[i]))
	}

	return numID
}

func encode(numID int) string {
	var encodedValueBuffer strings.Builder

	for numID > 0 {
		encodedValueBuffer.WriteString(string(base[numID%baseLength]))
		numID /= baseLength
	}

	return reverse(encodedValueBuffer.String())
}

type create struct {
	URL string `form:"url" json:"url" xml:"url"  binding:"required"`
}

func setup() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", nil)
	})

	r.GET("/:urlID", func(c *gin.Context) {
		urlID := c.Param("urlID")

		if isURL(urlID) {
			c.HTML(http.StatusOK, "redirect.tmpl", gin.H{"redirect_to": urlID})
			return
		}

		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	r.POST("/create", func(c *gin.Context) {
		var form create

		if err := c.ShouldBind(&form); err != nil {
			c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
			return
		}

		if !isURL(form.URL) {
			c.HTML(http.StatusNotAcceptable, "error.html", gin.H{"error": "Incorrect url"})
			return
		}

		c.HTML(http.StatusOK, "url.html", gin.H{"url": form.URL})
	})

	r.Run()
}

func isURL(urlString string) bool {
	u, err := url.Parse(urlString)

	return err == nil && u.Host != "" && u.Scheme != ""
}

func main() {
	db.Init()

	setup()
}
