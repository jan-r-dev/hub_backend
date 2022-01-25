package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func createTimestamp(stringTime string) time.Time {
	i, err := strconv.ParseInt(stringTime, 10, 64)

	if err != nil {
		fmt.Print("Error conversion: ", err)
	}
	tm := time.Unix(i, 0)
	//fmt.Println(tm)

	return tm
}

func importEnv() map[string]string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return myEnv
}
