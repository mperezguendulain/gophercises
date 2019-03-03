package urlshort

import (
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type yamlPathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		endpoint := r.URL.Path
		if dest, ok := pathsToUrls[endpoint]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}

}

func YamlHandler(yamlPaths []byte, fallback http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		yamlObjs, err := parseYaml(yamlPaths)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		pathsToUrls := getMapPathURLs(yamlObjs)

		endpoint := r.URL.Path
		if dest, ok := pathsToUrls[endpoint]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}

}

func parseYaml(data []byte) ([]yamlPathURL, error) {

	var yamlObjs []yamlPathURL
	err := yaml.Unmarshal(data, &yamlObjs)
	return yamlObjs, err

}

func getMapPathURLs(yamlObjs []yamlPathURL) map[string]string {

	pathsToUrls := make(map[string]string)

	for _, yamlObj := range yamlObjs {
		pathsToUrls[yamlObj.Path] = yamlObj.URL
	}

	return pathsToUrls

}
