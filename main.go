package main

import (
	. "fortunev2/lib"
	. "fortunev2/modes"
	"github.com/robfig/cron"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"time"

	//"time"
)

type Fortune struct {
	modes Modes
}

func NewFortune() *Fortune{
	return &Fortune{}
}

func (fortune *Fortune) run(modes ...Modes)  {
	configuration := make(map[string]interface{})
	config,err := ioutil.ReadFile("./fortune.yaml")
	E(err)

	err = yaml.Unmarshal(config,configuration)
	E(err)

	alarm := &Alarm{}
	stockSource := &Repository{}

	c := cron.New()
	c.AddFunc("*/9 * * * * ?", func() {
		cstSh, _ := time.LoadLocation("Asia/Shanghai")
		t := time.Now().In(cstSh)
		hour,minute := t.Hour(),t.Minute()
		if  hour >= 9 && minute >= 20 && hour <= 15 && minute >= 1  {
			for _,mode := range modes {
				mode.Init(configuration,alarm,stockSource)
				for _,stock := range mode.Iterator() {
					go mode.Run(stock.(map[interface{}]interface{}))
				}
			}
		}

		log.Println("====> 大叔ヾ(◍°∇°◍)ﾉﾞ乘风破浪吧")
	})

	c.AddFunc("* * */23 * * ?", func() {
		for _,mode := range modes {
			mode.Init(configuration,alarm,stockSource)
			mode.On()
		}
	})

	c.Start()

	select{}
}

func main() {
	fortune := NewFortune()
	fortune.run(NewGreater(),NewLess())
}
