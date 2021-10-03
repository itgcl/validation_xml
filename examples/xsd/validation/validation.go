package validation

import (
	"os"
	"path/filepath"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
)

func ValidationByXSDFile(xmlString string) error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	xsdFilePath := filepath.Join(path, "examples", "xsd", "resource", "34", "release-notification.xsd")
	s, err := xsd.ParseFromFile(xsdFilePath)
	if err != nil {
		return err
	}
	defer s.Free()
	input, err := libxml2.ParseString(xmlString)
	if err != nil {
		return  err
	}
	defer input.Free()
	err = s.Validate(input)
	if err != nil {
		return err
	}
	return nil
}

