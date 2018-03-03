package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"strings"
	"unicode"

	"github.com/BurntSushi/toml"
)

// Updater represents updater
type Updater struct {
	Name   string
	URL    string
	Params map[string]string
}

// Update executes update and returns result
func (updater *Updater) Update() (string, error) {
	url, err := updater.parse()
	if err != nil {
		return "", err
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// LoadUpdater loads updater from file
func LoadUpdater(fileName string) (*Updater, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var updater Updater
	err = toml.Unmarshal(b, &updater)
	if err != nil {
		return nil, err
	}

	return &updater, nil
}

func (updater *Updater) parse() (string, error) {
	r := strings.NewReader(updater.URL)
	reader := bufio.NewReader(r)

	result := ""
	vari := ""
	readingVar := false
	for {
		r, _, err := reader.ReadRune()
		if err == io.EOF {
			if readingVar && len(vari) != 0 {
				val, ok := updater.Params[vari]
				if !ok {
					return "", wrapError(errVarNotFound, vari)
				}
				result += val
			}
			break
		}
		if err != nil {
			return "", err
		}
		if r == '$' {
			if !readingVar {
				readingVar = true
				continue
			}

			val, ok := updater.Params[vari]
			if !ok {
				return "", wrapError(errVarNotFound, vari)
			}
			result += val
			vari = ""
			continue
		}
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && readingVar {
			readingVar = false
			val, ok := updater.Params[vari]
			if !ok {
				return "", wrapError(errVarNotFound, vari)
			}
			result += val
			vari = ""
		}
		if readingVar {
			vari += string(r)
			continue
		}
		result += string(r)
	}
	return result, nil
}
