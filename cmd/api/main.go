package main

import "log"

// "github.com/zekielmp/Agent/internal/server"

func main() {
	if err := Server(&Application{}); err != nil {
		log.Fatal(err)
	}

}
