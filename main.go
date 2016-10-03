package main // import "github.com/buildkite/polyglot-co-demo-backend"

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{"script": scriptPath()})
	})

	r.GET("/build", func(c *gin.Context) {
		c.JSON(200, gin.H{"build": buildNumber()})
	})

	r.GET("/forecasts", func(c *gin.Context) {
		forecastReq := ForecastRequest{
			Locations: []ForecastRequestLocation{
				// {"Auckland", "-36.8485", "174.7633"},
				{"Melbourne", "-37.8163", "144.9642"},
				// {"Wellington", "-41.2865", "174.7762"},
				// {"Milan", "45.4654", "9.1859"},
				// {"Shanghai", "31.2304", "121.4737"},
				{"Kraków", "50.0647", "19.9450"},
				// {"Tokyo", "35.6895", "139.6917"},
				// {"London", "51.5074", "0.1278"},
				// {"Tel Aviv", "32.0853", "34.7818"},
				{"Dubai", "25.2048", "55.2708"},
				// {"Berlin", "52.5200", "13.4050"},
        {"North Pole", "90", "0"},
			},
		}

		forecastResp, err := fetchForecasts(forecastReq)

		if err == nil {
			c.JSON(200, forecastResp)
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
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

type ForecastRequest struct {
	Locations []ForecastRequestLocation `json:"locations"`
}

type ForecastRequestLocation struct {
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Lng  string `json:"lng"`
}

type ForecastResponse struct {
	Forecasts []Forecast `json:"forecasts"`
	Build     string     `json:"build"`
}

type Forecast struct {
	Name    string  `json:"name"`
	Lat     string  `json:"lat"`
	Lng     string  `json:"lng"`
	High    float64 `json:"high"`
	Summary string  `json:"summary"`
}

func fetchForecasts(req ForecastRequest) (ForecastResponse, error) {
	weatherServiceUrl := os.Getenv("WEATHER_SERVICE_URL")

	if weatherServiceUrl == "" {
		return dummyResponse(req), nil
	}

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(req)

	fmt.Println("Request:\n%s", b)

	res, err := http.Post(weatherServiceUrl, "application/json; charset=utf-8", b)
	if err != nil {
		return ForecastResponse{}, err
	}

	var forecastResp ForecastResponse
	json.NewDecoder(res.Body).Decode(&forecastResp)

	fmt.Println("Response:\n%s", forecastResp)

	return forecastResp, nil
}

func dummyResponse(req ForecastRequest) ForecastResponse {
	return ForecastResponse{
		Forecasts: []Forecast{
			{req.Locations[0].Name, req.Locations[0].Lat, req.Locations[0].Lng, float64(rand.Intn(40)), "Sunny, clear skies with a chance of rain"},
			{req.Locations[1].Name, req.Locations[1].Lat, req.Locations[1].Lng, float64(rand.Intn(40)), "Morning rain clearing to a sunny afternoon"},
			{req.Locations[2].Name, req.Locations[2].Lat, req.Locations[2].Lng, float64(rand.Intn(40)), "Hot afternoon with chance of showers"},
			{req.Locations[3].Name, req.Locations[3].Lat, req.Locations[3].Lng, float64(rand.Intn(40)), "Cloudy periods with heavy rain in the evening"},
		},
		Build: "42",
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
