package main

import (
	"fmt"

	"github.com/lucas-simao/mutation-env-go/config"
)

type CustomEnv struct {
	Redis string
	S3    string
}

func main() {
	c := CustomEnv{
		Redis: "REDIS",
		S3:    "S3",
	}

	env := config.NewConfig[CustomEnv](c)

	fmt.Println(env.Custom.(CustomEnv).Redis)
	fmt.Println(env.Custom.(CustomEnv).S3)
}
