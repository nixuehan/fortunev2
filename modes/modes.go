package modes

import (
	. "fortunev2/lib"
)

type Modes interface {
	Init(map[string]interface{}, Notify, *Repository)
	Iterator() []interface{}
	Run(map[interface{}]interface{})
	On()
	Off(interface{})
	Suspend(interface{}) bool
}
