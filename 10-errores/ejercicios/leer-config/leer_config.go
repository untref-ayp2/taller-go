package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Puerto  int
	Host    string
	Debug   bool
}

func main() {
	c, err := LeerConfig("config.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", c)
	}
}

func LeerConfig(archivo string) (Config, error) {
	var cfg Config
	datos, err := os.ReadFile(archivo)
	if err != nil {
		return cfg, err
	}
	for _, linea := range strings.Split(string(datos), "\n") {
		linea = strings.TrimSpace(linea)
		if linea == "" || strings.HasPrefix(linea, "#") {
			continue
		}
		partes := strings.SplitN(linea, "=", 2)
		if len(partes) != 2 {
			continue
		}
		clave, valor := strings.TrimSpace(partes[0]), strings.TrimSpace(partes[1])
		switch clave {
		case "puerto":
			cfg.Puerto, _ = strconv.Atoi(valor)
		case "host":
			cfg.Host = valor
		case "debug":
			cfg.Debug = valor == "true"
		}
	}
	return cfg, nil
}
