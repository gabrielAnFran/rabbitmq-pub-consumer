package main

import (
	"consume/rabbit"
	"context"
	"fmt"
)

func main() {

	var ctx context.Context
	fmt.Println("Só de olho nessa fila...")
	err := rabbit.ConsumeMessage(ctx)
	if err != nil {
		panic(err)
	}

}
