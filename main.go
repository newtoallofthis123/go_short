package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Go Short Server Version 0.1")

	store := NewDBInstance()
	err := store.Init()
	if err != nil {
		log.Fatal(err)
	}

	env := GetEnv()

	api := NewAPIServer(env.ListenAddr, store)
	err = api.Start()
	if err != nil {
		log.Fatal(err)
	}
}
