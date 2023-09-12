package main

import "net/http"

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/current.json?key=b1ff2516113d47aea5a215637231009&q=LUCKNOW&aqi=no")
	if err != nil {
		panic(err) // stops the execution of the program and prints out the error.
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic("Weather API not available")
	}
}
