package main

import (
	"fmt"
	"net/http"

	"github.com/pmperrin/simpleI18n"
)

/**
 * main function is the entry point of the program.
 * It initializes and runs the examples.
 */
func main() {
	// /**
	//  * ExampleGoText test the usage of Go text.
	//  * See: https://github.com/leonelquinteros/gotext
	//  * result : the PO file is to long key and value are on different line
	//  */
	// examples.ExampleGoText()

	// /**
	//  * ExampleGo18n demonstrates the usage of 18n in Go.
	//  * See: "github.com/nicksnyder/go-i18n/v2/i18n"
	//  * result : the TAML or YAML ...  file is to long key and value are on different line
	//  */
	// examples.ExampleGo18n()

	// /**
	//  * ExampleProperties reads from properties files.
	//  * See: "github.com/go-ini/ini"
	//  * result : this lib is build for init file and not for i18n
	//  */
	// examples.ExampleProperties()

	// /**
	//  * ExampleProperites2 basic read of properties files whitout lib
	//  */
	// examples.ExampleProperites2()
	// useLib()
	ExampleProperites2Server()
}

func useLib() {
	fmt.Println(">> Starting lib")
	//i18n := simpleI18n.SimpleI18n{Path: "./local", Filename: "messages"}
	i18n, err := simpleI18n.InitI18n("./local", "messages", "fr")
	if err != nil {
		fmt.Println(err.Error())
	}
	lang, err := i18n.GetLang("ru")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(">" + lang.GetText("signup.title") + "<br/>")
	fmt.Println(">" + lang.GetTextWithParam("signup.button", map[string]string{"name": "Toto"}))

}

func ExampleProperites2Server() {

	i18n, err := simpleI18n.InitI18n("./local", "messages", "fr")
	if err != nil {
		fmt.Println(err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Chargez les propriétés une seule fois
		println("req / ")
		lang, err := i18n.GetLang("ru")
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Fprintf(w, "%s", lang.GetText("signup.title")+" - ")
		fmt.Fprintf(w, "%s", lang.GetText("signup.button"))

	})

	http.ListenAndServe(":8080", nil)
}
