package main

import (
	"fmt"
	"strings"
)

func main() {
	//LECTURA DE ARCHIVOS
	fmt.Println("Texto a traducir")
	var text string
	fmt.Scanln(&text)
	fmt.Println("Origen")
	var or string
	fmt.Scanln(&or)
	fmt.Println("Destino")
	var des string
	fmt.Scanln(&des)
	translate(text, strings.ToUpper(or), strings.ToUpper(des))

}

type Dictionary struct {
	tex   string
	bina  string
	morse string
}

func translate(textoATraducir string, formatoOrigen string, formatoDestino string) {
	//VALIDA QUE LOS PARAMETROS SEAN CORRECTOS
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
	//PARA SABER QUE CAMPOS SE TOMARAN EN CUENTA
	var posOr = 0
	var posDe = 0
	switch formatoOrigen {
	case "TEXT":
		posOr = 0
	case "MORSE":
		posOr = 1
	case "BINARY":
		posOr = 2
	default:
		fmt.Println("It's a weekday")
	}

	switch formatoDestino {
	case "TEXT":
		posDe = 0
	case "MORSE":
		posDe = 1
	case "BINARY":
		posDe = 2
	default:
		fmt.Println("NO DATA")
	}
	findDictionary(textoATraducir, posOr, posDe)
}

//DICCIONARIO DE DATOS
func findDictionary(textoATraducir string, formatoOrigen int, formatoDestino int) {
	val := [][]string{
		{" ", "   ", "   "},
		{"A", "._ ", "01000001"},
		{"B", "_... ", "01000010"},
		{"C", "_._. ", "01000011"},
		{"D", "_.. ", "01000100"},
		{"E", ". ", "01000101"},
		{"F", ".._. ", "01000110"},
		{"G", "__. ", "01000111"},
		{"H", ".... ", "01001000"},
		{"I", ".. ", "01001001"},
		{"J", ".___ ", "01001010"},
		{"K", "_._ ", "01001011"},
		{"L", "._.. ", "01001100"},
		{"M", "__ ", "01001101"},
		{"N", "_. ", "01001110"},
		{"O", "____ ", "01001111"},
		{"P", ".__. ", "01010001"},
		{"Q", "__._ ", "010100101"},
		{"R", "._. ", "010100110"},
		{"S", "... ", "01010011"},
		{"T", "_ ", "01010100"},
		{"U", ".._ ", "01010101"},
		{"V", "..._ ", "01010110"},
		{"W", ".__ ", "01010111"},
		{"X", "_.._ ", "01011000"},
		{"Y", "_.__ ", "01011001"},
		{"Z", "__.. ", "01011010"},
	}

	var result = ""
	//for para recorrer cada letra
	for _, letra := range textoATraducir {
		var cont = 0
		for _, v := range val {
			cont++
			if v[formatoOrigen] == string(letra) {
				result += v[formatoDestino]
			}
		}
	}
	fmt.Println("EL RESULTADO ES:")
	fmt.Println(result)
}

/*func parse(origen int, str string) []string {
	var sp []string
	//convertir a case
	switch origen {
	case 1:
		sp := strings.Split(str, " ")
		return sp
	case 2:
		var cont = 0
		for _, l := range str {
			cont += 8
			fmt.Println(l)
		}
		return nil
	default:
		return nil
	}

}*/
