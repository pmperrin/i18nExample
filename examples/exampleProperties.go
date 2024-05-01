package examples

import (
	"fmt"
	"log"

	"github.com/go-ini/ini"
)

func ExampleProperties() {

	fmt.Println(">> example properties <<")

	// Charger le fichier de propriétés français
	cfg, err := ini.Load("./local/messages_en.properties")
	if err != nil {
		log.Fatal("Fail to read file: ", err)
	}

	// Accéder aux valeurs par clé
	// welcomeMessage := cfg.Section("").Key("welcome.message").String()
	// exitMessage := cfg.Section("").Key("exit.message").String()

	// fmt.Println("French:")
	// fmt.Println("Welcome Message: ", welcomeMessage)
	// fmt.Println("Exit Message: ", exitMessage)

	// // Charger le fichier de propriétés anglais
	// cfg, err = ini.Load("config_en.properties")
	// if err != nil {
	// 	log.Fatal("Fail to read file: ", err)
	// }

	// Accéder aux valeurs par clé
	welcomeMessage := cfg.Section("").Key("welcome.message").String()
	exitMessage := cfg.Section("").Key("exit.message").String()

	fmt.Println("English:")
	fmt.Println("Welcome Message: ", welcomeMessage)
	fmt.Println("Exit Message: ", exitMessage)

	//fmt.Println("Listening on http://localhost:8080")
	//log.Fatal(http.ListenAndServe(":8080", nil))

}
