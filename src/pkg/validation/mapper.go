package validation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Map(attribute, rule string, vars ...[]interface{}) string {

	wdPath, err := os.Getwd()
	if err != nil {
		panic("cant get current working directory !")
	}

	path :=wdPath + string(os.PathSeparator) +
		"pkg/lang" +
		string(os.PathSeparator) +
		os.Getenv("APP_LOCAL") +
		string(os.PathSeparator) +
		"validation.json"

	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		panic("cant load validation file !")
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]string

	json.Unmarshal([]byte(byteValue), &result)

	return strings.Replace(result[rule], ":attribute", attribute, -1)
}
