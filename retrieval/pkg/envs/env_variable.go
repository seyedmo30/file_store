package envs

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Setup() {
	var err error

	if err = godotenv.Load(); err != nil {
		print(".env not found \nplease create .env\n")
		if err = godotenv.Load(".env.dev"); err != nil {
			findDotENVNested()
		}

	}
	validRequirementEnv()
}

// if not found , iterate in nested directories
func findDotENVNested() {
	var err error
	path := ".env"
	pathDev := ".env.dev"
	for i := 0; i < 10; i++ {

		if err = godotenv.Load(path, pathDev); err != nil {
			if strings.Contains(err.Error(), "no such file or directory") {
				path = "../" + path
				pathDev = "../" + pathDev
				continue
			}
			panic("panic in config parser : " + err.Error())
		} else {
			return
		}
	}

	panic("!!! .env not found \n")

}

func validRequirementEnv() {
	listEnvs := []string{
		"MAX_FILE_SIZE",
		"MAX_FILE_NAME_SIZE",
		"POSTGRES_HOST",
		"POSTGRES_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
		"SERVER_API_HOST",
		"SERVER_API_PORT",
		"KEY_ENCRYPTION",
	}
	isComplete := true
	for _, val := range listEnvs {
		if _, exists := os.LookupEnv(val); !exists {
			print("key :" + val + " not found in env\n")
			isComplete = false
		}
	}
	if !isComplete {
		panic("env is incomplete")
	}
}
