package profile

import (
	"io/ioutil"
	"os"
)

func readFile(path string) ([]byte, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func buildArrayMaps(params map[string]string) ([]map[string]string, error) {
	var attrs [](map[string]string)
	for key, v := range params {
		j := map[string]string{
			key: v,
		}

		attrs = append(attrs, j)
	}

	return attrs, nil
}