package controllers

import (
	"main/forms"
	"main/models"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// utils - move elsewhere

var base = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var baseLength = len(base)

func isURL(urlString string) bool {
	u, err := url.Parse(urlString)

	return err == nil && u.Host != "" && u.Scheme != ""
}

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

// LinkController ...
type LinkController struct{}

var linkModel = new(models.LinkModel)

// Create ...
func (controller LinkController) Create(c *gin.Context) {
	var form forms.Create

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}

	if !isURL(form.URL) {
		c.HTML(http.StatusNotAcceptable, "error.html", gin.H{"error": "Incorrect url"})
		return
	}

	link, err := linkModel.Create(form.URL)

	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "url.html", gin.H{"url": encode(link.NumID)})
}

// Match ..
func (controller LinkController) Match(c *gin.Context) {
	numID := decode(c.Param("urlID"))

	link, err := linkModel.Find(numID)

	if err != nil {
		c.HTML(http.StatusNotFound, "404.html", nil)
		return
	}

	c.HTML(http.StatusOK, "redirect.tmpl", gin.H{"redirect_to": link.URL})
}
