package asciibuilder

import (
	"localmods/asciigenerator"
	"localmods/asciimodel"
)

func Default() (string, error) {
	return ",.-*'*-.,.-*'*-.~Default~,.-*'*-.,.-*'*-.", nil
}

func Text(key string, path string) (string, error) {
	model := new(asciimodel.Model)

	model.Load(path)
	value, err := model.LookUp(key)
	if err != nil {
		return "", err
	}

	tmptxt, err := asciigenerator.RandomLine()
	if err != nil {
		return "", err
	}
	value += tmptxt

	return value, nil
}

func Add(key string, value string, path string) error {
	model := new(asciimodel.Model)

	model.Load(path)
	value, err := model.LookUp(key)
	if err != nil {
		return err
	}

	model.Add(key, value)

	return nil
}
