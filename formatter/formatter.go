package formatter

import (
	"fmt"
	"encoding/json"

	"github.com/ghodss/yaml"
)

const (
	JSON = "json"
	YAML = "yaml"
)

type Formatter struct {
	Type string
}

// NewFormatter ...
func NewFormatter (formatType string) *Formatter {
	return &Formatter{
		Type: formatType,
	}
}

// Format formats an object into either
func (f *Formatter) Format(subject interface{}) (*Buffer, error) {
	buffer := NewBuffer()
	var err error

	switch f.Type {
	case JSON:
		err = f.json(buffer, subject)
	case YAML:
		err = f.yaml(buffer, subject)
	default:
		err = fmt.Errorf("Unknown formatting type: %s", f.Type)
	}

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func (f *Formatter) json(buffer *Buffer, subject interface{}) error {
	enc := json.NewEncoder(buffer.GetBytesBuffer())
	enc.SetIndent("", "  ")
	if err := enc.Encode(&subject); err != nil {
		return err
	}

	return nil
}

func (f *Formatter) yaml(buffer *Buffer, subject interface{}) error {
	yml, err := yaml.Marshal(subject)
	if err != nil {
		return err
	}

	buffer.WriteBytes(yml)

	return nil
}
