package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"encoding/json"
    //"fmt"
    "io/ioutil"
    "log"
    "net/url"
)
type WeatherResult struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		SeaLevel  int     `json:"sea_level"`
		GrndLevel int     `json:"grnd_level"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		OneH float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.GET("/todays_menu", show)
    e.Logger.Fatal(e.Start(":1323"))
}

func show(c echo.Context) error {
    //u := new(User)
    lat := c.QueryParam("lat")
    lon := c.QueryParam("lon")
    result := GetWeather(lat,lon)
    return c.String(http.StatusOK, "weather:"+result)
}

// APIを叩く関数
func GetWeather(lat string ,lon string) string {
    values := url.Values{}
    baseUrl := "http://api.openweathermap.org/data/2.5/weather?"

    // query
    values.Add("appid", "")    // OpenWeatherのAPIKey
    values.Add("lat", lat)
    values.Add("lon", lon)

    weather := ParseJson(baseUrl + values.Encode())
    return weather
}

// ここでパースします。
func ParseJson(url string) string {
    weather := ""

    response, err := http.Get(url)
    if err != nil {   // エラーハンドリング
        log.Fatalf("Connection Error: %v", err)
        return "取得できませんでした"
    }

   // 遅延
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalf("Connection Error: %v", err)
        return "取得できませんでした"
    }

    jsonBytes := ([]byte)(body)
    data := new(WeatherResult)
    if err := json.Unmarshal(jsonBytes, data); err != nil {
        log.Fatalf("Connection Error: %v", err)
    }

    if data.Weather != nil {
        weather = data.Weather[0].Main
    }
    return weather
}