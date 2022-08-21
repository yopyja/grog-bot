package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Token  string
	Prefix string

	TestSKU string

	General string
	OwnerID string
	RoleID  string

	ItemURL string

	URL     string
	Host    string
	Origin  string
	Referer string
	Payload []string
	config  *configStruct
)

type configStruct struct {
	Token  string `json:"token"`
	Prefix string `json:"prefix"`

	TestSKU string `json:"testSKU"`

	General string `json:"general"`
	OwnerID string `json:"ownerID"`
	RoleID  string `json:"roleID"`

	ItemURL string `json:"itemURL"`

	URL     string   `json:"url"`
	Host    string   `json:"host"`
	Origin  string   `json:"origin"`
	Referer string   `json:"referer"`
	Payload []string `json:"payload"`
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

	TestSKU = config.TestSKU

	General = config.General
	OwnerID = config.OwnerID
	RoleID = config.RoleID

	ItemURL = config.ItemURL

	URL = config.URL
	Host = config.Host
	Origin = config.Origin
	Referer = config.Referer
	Payload = config.Payload
	return nil
}
