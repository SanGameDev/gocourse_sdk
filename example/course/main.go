package main

import (
	"errors"
	"fmt"
	"os"

	courseSdk "github.com/SanGameDev/gocourse_sdk/course"
)

func main() {
	userTrans := courseSdk.NewHttpClient("http://localhost:8082", "")

	user, err := userTrans.Get("7619dda0-9a75-4fbc-ad5a-5ac0fe46e039")
	if err != nil {
		if errors.As(err, &courseSdk.ErrNotFound{}) {
			fmt.Println("Not found: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Internal Server Error: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(user)
}
