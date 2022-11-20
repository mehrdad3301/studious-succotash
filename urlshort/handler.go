package urlshort

import (
	"net/http"
	"github.com/go-yaml/yaml"
)

func MapHandler(
	pathsToUrls map[string]string, 
	fallback http.Handler, 
) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) { 
		path := r.URL.Path 
		if dest, ok := pathsToUrls[path]; ok { 
			http.Redirect(w, r, dest, http.StatusFound)	
		}
		fallback.ServeHTTP(w, r)
	}
}

type pathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"` 
}

func YAMLHandler(
	yamlBytes []byte, 
	fallback http.Handler,
) (http.HandlerFunc, error) {
	
	pathUrls, err := parseYAML(yamlBytes)
	if err != nil { 
		return nil, err 
	}
	pathMap := buildMap(pathUrls)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(
	data []byte,
) ([]pathUrl, error) { 

	pathUrls := make([]pathUrl, 0)
	err := yaml.Unmarshal(data, &pathUrls)	
	if err != nil { 
		return nil, err 
	}
	return pathUrls, nil
}

func buildMap(
	pathUrls []pathUrl,
) (map[string]string) { 

	m := make(map[string]string)
	for _, pu := range pathUrls { 
		m[pu.Path] = pu.Url	
	}
	return m
}














