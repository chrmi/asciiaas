package asciibuilder

import (
	"localmods/asciimodel"
)

func Default() (string, error) {
	return ",.-*'*-.,.-*'*-.~Default~,.-*'*-.,.-*'*-.", nil
}

func Text(key string) (string, error) {
	model := new(asciimodel.Model)

	model.Load()
	text, err := model.LookUp(key)
	if err != nil {
		return "", err
	}

	return text, nil
}
