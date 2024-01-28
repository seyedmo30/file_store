package main

import (
	v1 "store/adaptor/handle/api/v1"
	"store/pkg/envs"
)

func main() {

	envs.Setup()
	r := v1.NewRouter()
	r.Logger.Fatal(r.Start("localhost" + ":" + "8080"))
}
