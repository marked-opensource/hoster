package tmocks

import "github.com/stretchr/testify/mock"

type EnabledRulesMockImplementation struct {
	mock.Mock
}

func (e EnabledRulesMockImplementation) GetAll() ([]string, error) {
	res := e.Called()
	return res.Get(0).([]string), res.Error(1)
}

func (e EnabledRulesMockImplementation) Add(ruleName string) error {
	panic("implement me")
}

func (e EnabledRulesMockImplementation) Remove(ruleName string) error {
	panic("implement me")
}

func (e EnabledRulesMockImplementation) ClearAll() error {
	panic("implement me")
}

func (e EnabledRulesMockImplementation) ReadRule(ruleName string) (string, error) {
	res := e.Called(ruleName)
	return res.String(0), res.Error(1)
}
