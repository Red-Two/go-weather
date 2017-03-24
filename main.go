package main

import (
  "encoding/json"
  //"flag"
  //"log"
  "net/http"
  //"strings"
  //"time"
  "fmt"
)

func main() {
  fmt.Println(query("Rochester"))
}

func query(city string) (weatherData,error) {
  apiKey := "63c43a8187390586d5dbb89fe9567a66"
  resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiKey + "&q=" + city)
  if err != nil {
    return weatherData{}, err
  }
  defer resp.Body.Close()
  var d weatherData

  if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
    return weatherData{}, err
  }
  return d, nil
}

type weatherData struct {
  Name string `json:"name"`
  Main struct {
    Kelvin float64 `json:"temp"`
  } `json:"main"`
}
