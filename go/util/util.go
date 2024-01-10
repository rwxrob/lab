package util

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

type YAMLMapDocs []InlineYAMLMap

func (m YAMLMapDocs) String() string {
	if m == nil {
		return `null`
	}
	b := strings.Builder{}
	for _, v := range m {
		b.WriteString("---\n")
		b.WriteString(v.String())
	}
	return b.String()
}

type InlineYAMLMap map[string]any

func (m InlineYAMLMap) String() string {
	if m == nil {
		return `null`
	}
	buf, err := yaml.Marshal(m)
	if err != nil {
		log.Print(err)
		return `null`
	}
	return string(buf)
}

type mapbuf struct {
	Map InlineYAMLMap `yaml:",inline" json:""`
}

// ParseYAMLMaps uses the 'inline' tag from yaml.v3 to parse a multidoc
// YAML stream into its separate documents by iteratively decoding the
// input bytes until io.EOF into InlineYAMLMap types that are returned
// as a single slice. Note that this is the only valid method of
// extracting map documents from a YAML stream because of YAML
// directives. In other words, splitting on `---` is not enough.
func ParseYAMLMaps(in []byte) (YAMLMapDocs, error) {
	maps := YAMLMapDocs{}
	d := yaml.NewDecoder(bytes.NewReader(in))
	for {
		o := new(mapbuf)
		err := d.Decode(o)
		if o == nil {
			continue
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return maps, err
		}
		maps = append(maps, o.Map)
	}
	return maps, nil
}
