package main

import (
	v1 "store/adaptor/handle/api/v1"
)

func main() {

	r := v1.NewRouter()
	r.Logger.Fatal(r.Start("localhost" + ":" + "8080"))
}
