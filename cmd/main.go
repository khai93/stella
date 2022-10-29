package main

import (
	"strconv"

	"github.com/khai93/stella/config"
	"github.com/khai93/stella/routes"
)

func main() {
	c, err := config.Get()
	if err != nil {
		panic(err)
	}

	r := routes.InitRoutes(*c)
	r.Run(":" + strconv.Itoa(c.Port))
}
