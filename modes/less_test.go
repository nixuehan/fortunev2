package modes

import (
	. "fortunev2/lib"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"testing"
)

func TestLess_Run(t *testing.T) {

	configuration := make(map[string]interface{})
	config,_ := ioutil.ReadFile("../fortune.yaml")

	yaml.Unmarshal(config,configuration)

	alarm := &Alarm{}
	stockSource := &Repository{}

	less := NewLess()
	less.Init(configuration,alarm,stockSource)

	for _,stock := range less.Iterator() {
		less.Run(stock.(map[interface{}]interface{}))
	}
}
