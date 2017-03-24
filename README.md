# go-weather

## prerequisites
1. Have Golang installed on your system, and in your `$PATH`
2. Get API keys for [Open Weather Maps][https://openweathermap.org/] and [Weather Underground][https://www.wunderground.com/]

## building
1. Once you have set your respective API keys
2. Build the binary with `go build`
3. Run the program with `./go-weather`

## endpoint
* Try hitting these endpoints
* `localhost:8080/weather/Rochester`
* `curl localhost:8080/weather/Tokyo`
* `curl localhost:8080/weather/London`
* `curl localhost:8080/weather/Paris`
* `curl localhost:8080/weather/Quebec`
