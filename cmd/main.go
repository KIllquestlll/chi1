package main

import (
	"chi/cmd/app"
	"context"
	"fmt"
)

func main() {
	app := app.New()

	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println(err)
	}
}
