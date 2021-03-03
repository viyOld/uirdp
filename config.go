package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type uirdpConfig struct {
	Title   string
	Servers []serverRDP
}

type serverRDP struct {
	Group   string `toml:"group"`
	Name    string `toml:"name"`
	IP      string `toml:"ip"`
	Port    int    `toml:"port"`
	Data    string `toml:"data"`
	Comment string `toml:"coment"`
}

var (
	config uirdpConfig
	server serverRDP
)

func getConfig() {
	// Load config file or creat it

	if _, err := toml.DecodeFile("uirdp.toml", &config); err != nil {
		if strings.Contains(err.Error(), "no such file or directory") == true {
			file, err := os.Create("uirdp.toml")
			if err != nil {
				log.Fatal(err)
			}
			config.Title = "stringTitle"
			server.Group = "myGroup"
			server.Name = "myName"
			server.IP = "192.168.0.3"
			server.Data = "21.08.2020"
			server.Comment = "my template server"
			config.Servers = append(config.Servers, server)
			server.Group = "youGroup"
			server.Name = "youName"
			server.IP = "172.16.0.3"
			server.Data = "21.08.2020"
			server.Comment = "you template server"
			config.Servers = append(config.Servers, server)
			server.Group = "themGroup"
			server.Name = "themName"
			server.IP = "10.0.0.3"
			server.Data = "21.08.2020"
			server.Comment = "them template server"
			config.Servers = append(config.Servers, server)
			if err := toml.NewEncoder(file).Encode(config); err != nil {
				// failed to encode
				log.Fatal(err)
			}

			//fmt.Println("worked: ", err)

			defer file.Close()
			//return
		} else {
			fmt.Println("file uirdp.toml is wrong", err.Error())
			return
		}
	}
	// ------------------------------------------------------------- end load config file
}