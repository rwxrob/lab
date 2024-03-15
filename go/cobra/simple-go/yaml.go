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
func FormatYAML(r io.Reader) ([]byte, error) {
	maps, err := ParseYAMLMaps(r)
	if err != nil {
		return nil, err
	}
	var out string
	for _, m := range maps {
		byt, err := yaml.Marshal(m)
		if err != nil {
			return byt, err
		}
		out += "---\n" + string(byt)
	}
	return []byte(out), nil
}

func ParseYAMLMaps(r io.Reader) ([]map[string]any, error) {
	maps := []map[string]any{}
	d := yaml.NewDecoder(r)
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

func ToYAML(it any) ([]byte, error) {
	switch v := it.(type) {
	case []*yaml.Node:
		buf := new(bytes.Buffer)
		for _, node := range v {
			buf.WriteString("---\n")
			byt, err := yaml.Marshal(node)
			if err != nil {
				return nil, err
			}
			buf.Write(byt)
		}
		return buf.Bytes(), nil
	}
	return yaml.Marshal(it)
}
