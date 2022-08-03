package yamlx

import (
	"io/ioutil"
	"testing"
)

func TestIsYamlValid(t *testing.T) {
	data, _ := ioutil.ReadFile("test.yaml")
	err := IsYamlValid(data)
	t.Log(err)
}
