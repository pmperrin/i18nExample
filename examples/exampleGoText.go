package examples

import (
	"fmt"

	"github.com/leonelquinteros/gotext"
)

func ExampleGoText() {

	fmt.Println(">> exampleGetText <<")
	// Create Locale with library path and language code
	l := gotext.NewLocale("./local", "fr")

	// Load domain '/path/to/locales/root/dir/es_UY/default.po'
	l.AddDomain("default")

	// Translate text from default domain
	//fmt.Println(l.Get("Translate this"))

	// Translate text with variables
	fmt.Println(l.Get("test"))
}
