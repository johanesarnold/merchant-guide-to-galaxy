package helpers

import "io/ioutil"

func Read(path string) string {
	data, err := ioutil.ReadFile(path)
	FailOnError(err, "Failed to read file")
	return string(data)
}
