package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mperezguendulain/gophercises/urlshort"
)

var yamlFile string

func init() {

	flag.StringVar(&yamlFile, "yaml", "db_urls.yml", "Nombre del archivo yaml donde se encuentran las urls.")
	flag.Parse()

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", hello)

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlPaths, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	yamlHandler := urlshort.YamlHandler(yamlPaths, mapHandler)

	fmt.Println("Iniciando el servidor en el puerto: 8080")
	http.ListenAndServe(":8080", yamlHandler)

}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "Hello, world!")

}
