package yamlx

import (
	"gopkg.in/yaml.v3"
)

func IsYamlValid(data []byte) error {
	var t struct{}
	return yaml.Unmarshal(data, &t)
}
