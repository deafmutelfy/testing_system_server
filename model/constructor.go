package model

import (
	"github.com/vokestd/TestingSystem-server/logger"
	"github.com/vokestd/TestingSystem-server/stuff"
	"io/ioutil"
	"os"
)

func construct_object(filename string, name string, parser stuff.MarkupParser, v interface{}) {
	lg := logger.GetInstance()
	lg.Infof("Loading %s...", name)

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		lg.Errorf("Could not read file: %s\n", err)

		os.Exit(1)
	}

	parser.Load(content, v)
}
