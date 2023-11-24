package main

import (
	"cchw2/pkg/conf"
	"cchw2/pkg/ninja"
	"cchw2/pkg/redis"
	"fmt"
	"os"

	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

var (
	redisSession *redis.Redis
	ninjaSession *ninja.Ninja
	cfg          conf.Config
)

func weatherHandler(c echo.Context) error {
	hostName, _ := os.Hostname()
	city := c.Param("city")
	value, err := redisSession.Read(city)
	if err == nil {
		fmt.Println(value, "city found in redis.")
		split := strings.Split(value, "*")
		return c.JSON(http.StatusOK, hostName+": [form redis] city min temprature is "+split[0]+" and max temprature is "+split[1])
	} else {
		fmt.Println("not found in redis")
		min, max, err := ninjaSession.GetWeather(city)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "error connecting to ninja.")
		}
		err = redisSession.Write(city, fmt.Sprintf("%d*%d", min, max))
		if err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusInternalServerError, "error writing to redis")
		}

		return c.JSON(http.StatusOK, hostName+": [form ninjas] city min temprature is "+fmt.Sprintf("%d", min)+" and max temprature is "+fmt.Sprintf("%d", max))
	}
}

func main() {
	cfg = conf.Load()
	redisSession = redis.New(cfg.Redis)
	ninjaSession = &ninja.Ninja{APIkey: cfg.Ninja.APIkey}
	e := echo.New()
	e.GET("/weather/:city", weatherHandler)
	e.Start("0.0.0.0:8080")
}
