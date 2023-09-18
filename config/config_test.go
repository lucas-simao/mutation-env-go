package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	t.Run("should return field custom with a value of any", func(t *testing.T) {
		env := NewConfig[any]()

		if env.Custom != nil {
			t.Errorf("expect <nil> but got %v", env.Custom)
		}
	})

	t.Run("should search values created inside of struct custom", func(t *testing.T) {
		type CustomEnv struct {
			SearchThis string
		}

		c := CustomEnv{
			SearchThis: "SEARCH_THIS",
		}

		searchThisValue := "value of ambient SEARCH_THIS"

		os.Setenv("SEARCH_THIS", searchThisValue)

		env := NewConfig[CustomEnv](c)

		if _, ok := env.Custom.(CustomEnv); !ok {
			t.Errorf("expect CustomEnv but got %v", env.Custom)
		}

		if env.Custom.(CustomEnv).SearchThis != searchThisValue {
			t.Errorf("expect %s but got %s", searchThisValue, env.Custom.(CustomEnv).SearchThis)
		}

	})
}
