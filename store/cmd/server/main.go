package main

import (
	"os"
	v1 "store/adaptor/handle/api/v1"
	"store/pkg/envs"
)

func main() {

	envs.Setup()
	r := v1.NewRouter()
	r.Logger.Fatal(r.Start(os.Getenv("SERVER_API_HOST") + ":" + os.Getenv("SERVER_API_PORT") ))
}
