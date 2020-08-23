package main

import (
	"github.com/hunzo/go-gin-handler-v2/routers"
)

func main() {

	s := routers.Routers()
	s.Run()

}
