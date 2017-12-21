package clipboard

import (
	"io/ioutil"
	"os"
)

func readAll() (string, error) {
	f, err := os.Open("/dev/snarf")
	if err != nil {
		return "", err
	}
	defer f.Close()
	v, err := ioutil.ReadAll(f)
	return string(v), err
}

func writeAll(text string) error {
	return ioutil.WriteFile("/dev/snarf", []byte(text), 0755)
}
