# mutation-env-go

This is a simple example how to mutation a func and returning new values inside of response.

It uses the generics to do this process.

```go
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
```
