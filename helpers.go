package main

import (
	"fmt"
	"strconv"
	"time"
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
