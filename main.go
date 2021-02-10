package main

import (
	"fmt"
)

func main() {

	fmt.Println("Texto a traducir")
	var text string
	fmt.Scanln(&text)
	fmt.Println("Origen")
	var or string
	fmt.Scanln(&or)
	fmt.Println("Destino")
	var des string
	fmt.Scanln(&des)
	translate(text, or, des)

}

type Dictionary struct {
	tex   string
	bina  string
	morse string
}

func translate(textoATraducir string, formatoOrigen string, formatoDestino string) {
	valid := []string{"TEXT", "BINARY", "MORSE"}
	var o = false
	var d = false
	for _, v := range valid {
		if v == formatoOrigen {
			o = true
		}
		if v == formatoDestino {
			d = true
		}
	}
	if !o || !d {
		fmt.Println("No soportado")
	} else {
		seleccion(textoATraducir, formatoOrigen, formatoDestino)
	}
}

func seleccion(textoATraducir string, formatoOrigen string, formatoDestino string) {

	var posOr = 0
	var posDe = 0
	switch formatoOrigen {
	case "TEXT":
		posOr = 0
	case "BINARY":
		posOr = 1
	case "MORSE":
		posOr = 2
	default:
		fmt.Println("It's a weekday")
	}

	switch formatoOrigen {
	case "TEXT":
		posDe = 0
	case "BINARY":
		posDe = 1
	case "MORSE":
		posDe = 2
	default:
		fmt.Println("It's a weekday")
	}
	findDictionary(textoATraducir, posOr, posDe)
}

func findDictionary(textoATraducir string, formatoOrigen int, formatoDestino int) {
	val := []Dictionary{
		{"a", "._", ""},
		{"a", "._", ""},
	}

	//for para recorrer cada letra
	for _, letra := range textoATraducir {
		for _, v := range val {
			if v[formatoOrigen].Name == letra {
				fmt.Println(v[0])
			}
		}
	}
}
