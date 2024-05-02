package examples

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	properties map[string]string
	mu         sync.RWMutex
)

func loadPropertiesOnce(filePath string) {
	props := make(map[string]string)
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(nil, "Erreur lors de l'ouverture du fichier", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			props[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		http.Error(nil, "Erreur lors de la lecture du fichier", http.StatusInternalServerError)
	}

	mu.Lock()
	properties = props
	mu.Unlock()
}

func ExampleProperites2() {

	println("req / ")
	if properties == nil {
		loadPropertiesOnce("./local/messages_en.properties")
	}

	// Utilisez les clés pour accéder aux valeurs
	greeting := properties["welcome.message"]
	println(">", greeting) // Affiche "Hello, world!" si la clé "greeting" est définie dans le fichier

}

func ExampleProperites2Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Chargez les propriétés une seule fois
		println("req / ")
		if properties == nil {
			loadPropertiesOnce("messages_en.properties")
		}

		// Utilisez les clés pour accéder aux valeurs
		greeting := properties["welcome.message"]
		fmt.Fprintf(w, "%s", greeting) // Affiche "Hello, world!" si la clé "greeting" est définie dans le fichier
	})

	http.ListenAndServe(":8080", nil)
}
