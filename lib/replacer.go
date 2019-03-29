package lib

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var matcher = regexp.MustCompile(`(?mU)^[^*\n\r]*omitempty.*$`)

func ReplaceOmits(file string) error {

	// Open the source file
	f, err := os.OpenFile(file, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Create a temp file for saving stripped text
	temp, err := ioutil.TempFile("", "replaced")
	if err != nil {
		return err
	}

	defer cleanup(temp)

	reader := bufio.NewReader(f)
	writer := bufio.NewWriter(temp)

	// Going through lines, match the regex and replace omitempty on primitives and enums
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		cleaned := matcher.ReplaceAllFunc(line, func(old []byte) []byte {
			return []byte(strings.Replace(string(old), ",omitempty", "", -1))
		})

		_, err = writer.Write(cleaned)
		if err != nil {
			return err
		}

		err = writer.Flush()
		if err != nil {
			return err
		}
		err = temp.Sync()
		if err != nil {
			return err
		}

	}

	// overwrite file

	path := f.Name()
	err = f.Close()
	if err != nil {
		return err
	}

	temp.Close()

	return os.Rename(temp.Name(), path)
}

func cleanup(f *os.File) {
	f.Close()
	os.Remove(f.Name())
}
