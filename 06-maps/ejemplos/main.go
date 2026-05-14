package main

import "fmt"

func main() {
	telefonos := make(map[string]int)

	telefonos["Juan"] = 1533697845
	telefonos["Ana"] = 1547896547
	telefonos["Mauro"] = 1525873214
	telefonos["Pedro"] = 1533557788

	for clave, valor := range telefonos {
		fmt.Println(clave, "\t", valor)
	}

	agenda := map[string]string{
		"Juan":  "El Payador 4457",
		"Ana":   "Sarmiento 2544",
		"Mauro": "Santa Fe 2588",
		"Pedro": "San Martin 1278",
	}

	for clave, _ := range telefonos {
		fmt.Println(clave, "\t", agenda[clave], "\t", telefonos[clave])
	}

}
