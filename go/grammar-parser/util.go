package lib

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

func toJSON(i interface{}) string {
	byt, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf(`{"ERROR":"%v"}`, err)
	}
	return string(byt)
}

func toYAML(i interface{}) string {
	byt, err := yaml.Marshal(i)
	if err != nil {
		return fmt.Sprintf(`ERROR:"%v`, err)
	}
	return string(byt)
}

func trapPanic() {
	p := recover()
	if p != nil {
		log.Println(p)
	}
}
