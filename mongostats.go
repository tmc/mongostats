package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	if err := mongostats(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func mongostats() error {
	url := os.Getenv("MONGO_URL")
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	var serverStatus = &bson.M{}
	if err := session.Run("serverStatus", serverStatus); err != nil {
		return err
	}
	out, err := json.MarshalIndent(serverStatus, "", "  ")
	if err != nil {
		return err
	}
	var ss ServerStatus

	if err := json.Unmarshal(out, &ss); err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
