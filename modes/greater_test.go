package modes

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
	. "fortunev2/lib"
)

func TestGreater_Run(t *testing.T) {

	configuration := make(map[string]interface{})
	config,_ := ioutil.ReadFile("../fortune.yaml")

	yaml.Unmarshal(config,configuration)

	alarm := &Alarm{}
	stockSource := &Repository{}

	greater := NewGreater()
	greater.Init(configuration,alarm,stockSource)

	for _,stock := range greater.Iterator() {
		greater.Run(stock.(map[interface{}]interface{}))
	}
}
