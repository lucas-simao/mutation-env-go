package config

import (
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Env struct {
	DBUrl  string
	Custom any
}

func NewConfig[T any](customs ...T) Env {
	err := godotenv.Load("config/.env")
	if err != nil {
		err := godotenv.Load()
		if err != nil {
			log.Panic(err)
		}
	}

	dbUrl, ok := os.LookupEnv("DB_URL")
	if !ok {
		log.Panic("error to load DB_URL")
	}

	var custom T

	if len(customs) > 0 {
		custom = loadCustomEnv[T](customs[0])
	}

	return Env{
		DBUrl:  dbUrl,
		Custom: custom,
	}
}

func loadCustomEnv[T any](custom T) T {
	rTypeOf := reflect.TypeOf(&custom).Elem()
	rValueOf := reflect.ValueOf(&custom).Elem()

	for i := 0; i < rTypeOf.NumField(); i++ {
		varEnv, ok := os.LookupEnv(rValueOf.Field(i).String())
		if !ok {
			log.Panicf("error to load %s", rTypeOf.Field(i).Name)
		}

		rValueOf.Field(i).Set(reflect.ValueOf(varEnv))
	}

	return custom
}
