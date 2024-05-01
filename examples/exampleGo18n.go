package examples

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func ExampleGo18n() {

	fmt.Println(">> exampleGo18n <<")

	bundle := i18n.NewBundle(language.Spanish)

	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// No need to load active.en.toml since we are providing default translations.
	// bundle.MustLoadMessageFile("active.en.toml")
	bundle.MustLoadMessageFile("./local/active.es.toml")

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// lang := r.FormValue("lang")
	// accept := r.Header.Get("Accept-Language")
	localizer := i18n.NewLocalizer(bundle, "es", "es")

	fmt.Println(">>  " + trad(localizer, "HelloPerson"))

	trad(localizer, "toto")
	// name := "billy!"

	// helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{
	// 	DefaultMessage: &i18n.Message{
	// 		ID:    "HelloPerson",
	// 		Other: "Hello {{.Name}}",
	// 	},
	// 	TemplateData: map[string]string{
	// 		"Name": name,
	// 	},
	// })
	//fmt.Println(">>  " + helloPerson)

	//})

	//fmt.Println("Listening on http://localhost:8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))

}

func trad(localizer *i18n.Localizer, key string) string {
	key, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: key,
		},
	})
	if err != nil {
		fmt.Println("aze :" + err.Error())
	}
	return key
}
