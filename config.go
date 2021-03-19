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
			// server.Group = "myGroup"
			// server.Name = "myName"
			// server.IP = "192.168.0.3"
			// server.Data = "21.08.2020"
			// server.Comment = "my template server"
			// config.Servers = append(config.Servers, server)
			config.Servers = append(config.Servers, serverRDP{Group: "myGroup", Name: "myName", IP: "192.168.0.3", Port: 3389, Data: "21.08.2020", Comment: "my template server"})
			config.Servers = append(config.Servers, serverRDP{Group: "youGroup", Name: "youName", IP: "172.16.0.3", Port: 3389, Data: "21.08.2020", Comment: "you template server"})
			// server.Group = "youGroup"
			// server.Name = "youName"
			// server.IP = "172.16.0.3"
			// server.Data = "21.08.2020"
			// server.Comment = "you template server"
			// config.Servers = append(config.Servers, server)
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

func setConfig() {
	f, err := os.Create("uirdp.toml")
	if err != nil {
		// failed to create/open the file
		log.Fatal(err)
	}
	if err := toml.NewEncoder(f).Encode(config); err != nil {
		// failed to encode
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		// failed to close the file
		log.Fatal(err)

	}
}
