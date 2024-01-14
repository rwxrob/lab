package simple

import (
	"bytes"
	"errors"
	"io"

	"gopkg.in/yaml.v3"
)

type yamlMap struct {
	O map[string]any `yaml:",inline"`
}

// Format just returns a string with the YAML formatted by default
// yaml.v3.
func FormatYAML(in []byte) ([]byte, error) {
	maps, err := ParseYAMLMaps(in)
	if err != nil {
		return nil, err
	}
	var out string
	for _, m := range maps {
		out += "---\n" + string(ToYAML(m))
	}
	return []byte(out), nil
}

func ParseYAMLMaps(in []byte) ([]map[string]any, error) {
	maps := []map[string]any{}
	d := yaml.NewDecoder(bytes.NewReader(in))
	for {
		ymap := new(yamlMap)
		err := d.Decode(ymap)
		if ymap == nil {
			continue
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		maps = append(maps, ymap.O)
	}
	return maps, nil
}

func ToYAML(in any) []byte {
	byt, err := yaml.Marshal(in)
	if err != nil {
		byt = []byte(`null`)
	}
	return byt
}
