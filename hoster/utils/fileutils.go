package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/looplab/fsm"
)

type IFileUtils interface {
	InjectRule(envRulesPath string) error
	ClearRule(envRulesPat string) error
	RefreshRules() error
	ClearAll()
}

type fileUtils struct {
	envFile      string
	enabledRules EnabledRules
}

type enabledRulesImplementation struct {
	enabledRulesFile string
}

func (e enabledRulesImplementation) GetAll() ([]string, error) {
	panic("implement me")
}

func (e enabledRulesImplementation) Add(ruleName string) error {
	panic("implement me")
}

func (e enabledRulesImplementation) Remove(ruleName string) error {
	panic("implement me")
}

func (e enabledRulesImplementation) ClearAll() error {
	panic("implement me")
}

func (e enabledRulesImplementation) ReadRule(ruleName string) (string, error) {
	panic("implement me")
}

type EnabledRules interface {
	GetAll() ([]string, error)
	Add(ruleName string) error
	Remove(ruleName string) error
	ClearAll() error
	ReadRule(ruleName string) (string, error)
}

type fsmWrapper struct {
	fsm *fsm.FSM
}

func (fsw *fsmWrapper) StateAwareAppendLine(outputLines []string, line string) []string {
	if fsw.fsm.Current() == "base" {
		return append(outputLines, line)
	}

	return outputLines
}

func (fu *fileUtils) RefreshRules() error {
	rules, err := fu.enabledRules.GetAll()
	if err != nil {
		panic(err)
	}

	input, err := ioutil.ReadFile(fu.envFile)
	if err != nil {
		panic(err)
	}

	fsmInstance := fsmWrapper{fsm.NewFSM(
		"base",
		fsm.Events{
			{Name: "start", Src: []string{"base"}, Dst: "injecting"},
			{Name: "close", Src: []string{"injecting"}, Dst: "base"},
		},
		fsm.Callbacks{},
	)}

	lines := strings.Split(string(input), "\n")
	var outputLines []string

	for _, line := range lines {
		if fu.beginHosterBlockMatched([]byte(line)) {
			err := fsmInstance.fsm.Event("start")
			if err != nil {
				panic(err)
			}
			outputLines = append(outputLines, hosterBlockHeader)
			for _, rule := range rules {
				ruleContent, err := fu.enabledRules.ReadRule(rule)
				if err != nil {
					log.Printf("could not read %s, reason: %s", rule, err)
					continue
				}
				ruleLines := strings.Split(ruleContent, "\n")
				outputLines = append(outputLines, ruleLines...)
			}
		}
		outputLines = fsmInstance.StateAwareAppendLine(outputLines, line)
		if fu.endHosterBlockMatched([]byte(line)) && fsmInstance.fsm.Can("close") {
			err := fsmInstance.fsm.Event("close")
			if err != nil {
				panic(err)
			}
			outputLines = append(outputLines, hosterBlockFooter)
		}
	}

	output := strings.Join(outputLines, "\n")

	err = ioutil.WriteFile(fu.envFile, []byte(output), 0644)
	if err != nil {
		panic(err)
	}
	return nil
}

const hosterHostsBlockPrefix string = "^\\# Added by Hoster"
const endingHosterHostsBlockPrefix string = "^\\# End of section"

const hosterBlockHeader = "# Added by Hoster"
const hosterBlockFooter = "# End of section"

var hosterBlock = fmt.Sprintf(`
%s
%s
`, hosterBlockHeader, hosterBlockFooter)

func (fu *fileUtils) getEnvFile() (file *os.File) {
	file, err := os.OpenFile(fu.envFile, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	return
}

func (fu *fileUtils) beginHosterBlockMatched(lineBytes []byte) bool {
	matched, err := regexp.Match(hosterHostsBlockPrefix, lineBytes)
	if err != nil {
		return false
	}
	return matched
}

func (fu *fileUtils) endHosterBlockMatched(lineBytes []byte) bool {
	matched, err := regexp.Match(endingHosterHostsBlockPrefix, lineBytes)
	if err != nil {
		return false
	}
	return matched
}

func (fu *fileUtils) EnsureHosterBlock() (err error) {
	file := fu.getEnvFile()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matched := fu.beginHosterBlockMatched(scanner.Bytes())
		if matched {
			return nil
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
		envFile:      filePath,
		enabledRules: enabledRulesImplementation{enabledRulesFile: ""},
	}
}
