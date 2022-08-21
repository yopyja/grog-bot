package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Token   string
	Prefix  string
	General string
	ItemURL string
	OwnerID string
	RoleID string
	config  *configStruct
)

type configStruct struct {
	Token   string `json:"token"`
	Prefix  string `json:"prefix"`
	General string `json:"general"`
	ItemURL string `json:"itemURL"`
	OwnerID string `json:"ownerID"`
	RoleID string `json;"roleID"`
}

func ReadConfig() error {
	file, err := os.ReadFile("./json/config.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Token = config.Token
	Prefix = config.Prefix
	General = config.General
	ItemURL = config.ItemURL
	OwnerID = config.OwnerID
	RoleID = config.RoleID

	return nil
}
