package main

import (
	"github.com/pmperrin/examples"
	"github.com/pmperrin/simpleI18n"
)

/**
 * main function is the entry point of the program.
 * It initializes and runs the examples.
 */
func main() {
	/**
	 * ExampleGoText test the usage of Go text.
	 * See: https://github.com/leonelquinteros/gotext
	 * result : the PO file is to long key and value are on different line
	 */
	examples.ExampleGoText()

	/**
	 * ExampleGo18n demonstrates the usage of 18n in Go.
	 * See: "github.com/nicksnyder/go-i18n/v2/i18n"
	 * result : the TAML or YAML ...  file is to long key and value are on different line
	 */
	examples.ExampleGo18n()

	/**
	 * ExampleProperties reads from properties files.
	 * See: "github.com/go-ini/ini"
	 * result : this lib is build for init file and not for i18n
	 */
	examples.ExampleProperties()

	/**
	 * ExampleProperites2 basic read of properties files whitout lib
	 */
	examples.ExampleProperites2()
	useLib()
}

func useLib() {
	i18n := simpleI18n.SimpleI18n{Path: "./local", Filename: "messages"}
	i18n.Init()

}
