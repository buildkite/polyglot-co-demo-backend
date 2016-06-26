package main // import "github.com/buildkite/polyglot-co-demo-backend"

import (
  "io/ioutil"
	"net/http"
	"os"
  "strings"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("templates/*.tmpl.html")
	r.Static("/static", "static")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
      "script": scriptPath(),
    })
	})

	r.GET("/build", func(c *gin.Context) {
    c.JSON(200, gin.H{"build": buildNumber()})
  })

	r.GET("/forecasts", func(c *gin.Context) {
    c.JSON(200, fetchForecasts())
  })

	r.Run(":" + port)
}

func scriptPath() string {
  if os.Getenv("FRONTEND_DEV") == "true" {
    return "http://localhost:8000/polyglot-co.js"
  } else {
    return "/static/polyglot-co.js"
  }
}

func fetchForecasts() gin.H {
  return gin.H{
    "forecasts": []gin.H{
      {"name": "Auckland",   "high": 18, "forecast": "Sunny, clear skies with a chance of rain"},
      {"name": "Wellington", "high": 19, "forecast": "Morning rain clearing to a sunny afternoon"},
      {"name": "Milan",      "high": 32, "forecast": "Hot afternoon with chance of showers"},
      {"name": "Tokyo",      "high": 27, "forecast": "Cloudy periods with heavy rain in the evening"},
    },
    "build": "42",
  }
}

func buildNumber() string {
  number, err := ioutil.ReadFile("static/build-number")

  if err != nil {
    return "42"
  } else {
    return strings.TrimSpace(string(number))
  }
}
