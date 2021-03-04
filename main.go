package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
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

	r.GET("/:urlID", func(c *gin.Context) {
		urlID := c.Param("urlID")
		c.JSON(http.StatusOK, gin.H{"urlID": urlID})
	})

	r.POST("/create", func(c *gin.Context) {
		var form create

		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !isURL(form.URL) {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Incorrect url"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"url": form.URL})
	})

	r.Run()
}

func isURL(urlString string) bool {
	u, err := url.Parse(urlString)
	return err == nil && u.Host != "" && u.Scheme != ""
}

func main() {
	// start := time.Now()

	encoded := encode(765130)
	fmt.Printf("%s encoded\n", encoded)
	// fmt.Printf("%d decoded\n", decode(encoded))

	// elapsed := time.Since(start)
	// fmt.Printf("Encoding took %s\n", elapsed)
	setup()
}
