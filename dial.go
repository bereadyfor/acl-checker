package main

import (
	"io/ioutil"
	"log"
	"net"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Type   string
	Dst    string
	Port   string
	Notify string
}

type Configs struct {
	Cfgs []Config `tasks`
}

func main() {
	confFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Panicf("Read config. err #%v ", err)
	}

	var configs Configs
	err = yaml.Unmarshal(confFile, &configs)
	if err != nil {
		log.Panicf("Parse config. err #%v ", err)
	}

	for _, config := range configs.Cfgs {
		log.Println(config)
		if config.Type == "ip" {
			health, err := checkPort("tcp", config.Dst, config.Port)
			log.Println(health, err)
		}
	}
}

func checkPort(protocol string, ip string, port string) (bool, error) {
	conn, err := net.Dial(protocol, ip+":"+port)
	if err != nil {
		return false, err
	} else {
		defer conn.Close()
		return true, nil
	}
}

func checkPing() {
	/* TODO implement */
}

func checkDNSLookup() {
	/* TODO implement */
}

func notify() {
	/* TODO implement */
}
