package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type IFileUtils interface {
	InjectRule(envRulesPath string) error
	ClearRule(envRulesPat string) error
	ClearAll()
}

type fileUtils struct {
	envFile string
}

const hosterHostsBlockPrefix string = "^\\# Added by Hoster"

const hosterBlock string = `
# Added by Hoster
# End of section
`

func (fu *fileUtils) EnsureHosterBlock() (err error) {
	file, err := os.OpenFile(fu.envFile, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matched, err := regexp.Match(hosterHostsBlockPrefix, scanner.Bytes())
		if err != nil {
			panic(err)
		}
		if matched {
			return err
		}
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	_, err = file.WriteString(hosterBlock)
	if err != nil {
		panic(err)
	}
	return
}

func (*fileUtils) InjectRule(envRulesPath string) (err error) {
	data, err := ioutil.ReadFile(envRulesPath)
	if err != nil {
		return
	}
	fmt.Print(data)
	return
}

func (*fileUtils) ClearRule(envRulesPath string) (err error) {
	return
}

func (*fileUtils) ClearAll() {
	panic("implement me")
}

func NewFileUtils(filePath string) IFileUtils {
	return &fileUtils{
		envFile: filePath,
	}
}
