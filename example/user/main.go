package main

import (
	"errors"
	"fmt"
	"os"

	userSdk "github.com/SanGameDev/gocourse_sdk/user"
)

func main() {
	userTrans := userSdk.NewHttpClient("http://localhost:8081", "")

	user, err := userTrans.Get("aad1b18c-9526-4d8f-a74a-fdc038cdacfb")
	if err != nil {
		if errors.As(err, &userSdk.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
