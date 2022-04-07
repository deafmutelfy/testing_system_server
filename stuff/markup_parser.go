package stuff

import (
	"github.com/goccy/go-yaml"
	"github.com/vokestd/TestingSystem-server/logger"
	"os"
)

type MarkupParser interface {
	Load([]byte, interface{})
}

type MarkupParserYaml struct{}

func (m *MarkupParserYaml) Load(data []byte, v interface{}) {
	if err := yaml.Unmarshal(data, v); err != nil {
		logger.GetInstance().Errorf("Could not parse object: %s\n", err)

		os.Exit(1)
	}
}
