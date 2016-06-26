package main // import "github.com/buildkite/polyglot-co-demo-backend"

import (
  "io/ioutil"
  "math/rand"
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
  weatherServiceUrl := os.Getenv("WEATHER_SERVICE_URL")

  cities := []gin.H{
    {"name": "Auckland",   "lat": "...", "lng": "..."},
    {"name": "Wellington", "lat": "...", "lng": "..."},
    {"name": "Milan",      "lat": "...", "lng": "..."},
    {"name": "Tokyo",      "lat": "...", "lng": "..."},
  }

  if weatherServiceUrl == "" {
    // Generate some dummy data
    return gin.H{
      "forecasts": []gin.H{
        {"name": cities[0]["name"], "high": rand.Intn(40), "forecast": "Sunny, clear skies with a chance of rain"},
        {"name": cities[1]["name"], "high": rand.Intn(40), "forecast": "Morning rain clearing to a sunny afternoon"},
        {"name": cities[2]["name"], "high": rand.Intn(40), "forecast": "Hot afternoon with chance of showers"},
        {"name": cities[3]["name"], "high": rand.Intn(40), "forecast": "Cloudy periods with heavy rain in the evening"},
      },
      "build": "42",
    }
  } else {
    // TODO:
    //
    // POST ${weatherServiceUrl}
    // {
    //   "locations": [
    //     { "lat": "...", "lng": "..." },
    //     { "lat": "...", "lng": "..." }
    //   ]
    // }
    //
    // which returns:
    //
    // {
    //   "forecasts": [
    //     { "high": 1.2, "summary": "Sunny with a chance..."},
    //     { "high": 1.2, "summary": "Sunny with a chance..."}
    //   ]
    // }
    //
    // and then we transform into:
    //
    // {
    //   "forecasts": [
    //     {"name": "A City", "high":}
    //   ]
    // }
    return gin.H{}
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
