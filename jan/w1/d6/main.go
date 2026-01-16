package d6

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type Plugin interface {
	Name() string
	Execute(input string) (string, error)
}

type ReversePlugin struct{}

func (rp ReversePlugin) Name() string {
	return "ReversePlugin"
}
func (rp ReversePlugin) Execute(input string) (string, error) {
	// take the input string and reverse it
	stringRunes := []rune(input)
	var reversedStringRunes []rune
	for i := len(stringRunes) - 1; i >= 0; i-- {
		reversedStringRunes = append(reversedStringRunes, stringRunes[i])
	}
	return string(reversedStringRunes), nil
}

type Base64Plugin struct{}

func (bp Base64Plugin) Name() string {
	return "Base64Plugin"
}
func (bp Base64Plugin) Execute(input string) (string, error) {
	return base64.StdEncoding.EncodeToString([]byte(input)), nil
}

type HashPlugin struct{}

func (bp HashPlugin) Name() string {
	return "HashPlugin"
}
func (bp HashPlugin) Execute(input string) (string, error) {
	inputBytes := []byte(input)
	h := sha256.New()
	h.Write(inputBytes)
	sha := h.Sum(nil) // to get just the current hash

	return string(sha), nil
}

type PluginManager struct {
	Plugins         map[string]Plugin
	PluginExecOrder []Plugin //we use this to enable us run the plugins sequentially
}

func NewPluginManager() *PluginManager {
	return &PluginManager{
		Plugins:         map[string]Plugin{},
		PluginExecOrder: []Plugin{},
	}
}

func (pm *PluginManager) RegisterPlugin(p Plugin) {
	pm.Plugins[p.Name()] = p
	pm.PluginExecOrder = append(pm.PluginExecOrder, p)
}

func (pm *PluginManager) Run(name, input string) (string, error) {
	plugin, exists := pm.Plugins[name]
	if !exists {
		return "", fmt.Errorf("This plugin does not exist")
	}
	val, err := plugin.Execute(input)
	if err != nil {
		return "", fmt.Errorf("This plugin failed to execute")
	}
	return val,nil
}

func (pm PluginManager) RunAll(input string) (string, error) {
	var startInput string = input
	for _, v := range pm.PluginExecOrder {
		val, err := v.Execute(startInput)
		if err != nil {
			return "", fmt.Errorf("This plugin failed to execute")
		}
		startInput = val
	}
	return startInput,nil
}
