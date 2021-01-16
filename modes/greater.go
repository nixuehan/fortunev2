package modes

import (
	. "fortunev2/lib"
	"log"
	"strconv"
	"sync"
)

type Greater struct {
	stocks []interface{}
	notify Notify
	greaterMap sync.Map
	stockSource *Repository
}

func NewGreater() *Greater{
	return &Greater{}
}

func (greater *Greater) Init(stocks map[string]interface{}, notify Notify, stockSource *Repository) {
	greater.stocks = stocks["greater"].([]interface{})
	greater.notify = notify
	greater.stockSource = stockSource
}

func (greater *Greater) Iterator() []interface{} {
	return greater.stocks
}

func (greater *Greater) Run(stock map[interface{}]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("OMG！%s",err)
		}
	}()

	newestStock,_ := greater.stockSource.Fetch(stock["name"].(string))
	for _,v := range stock["price"].([]interface{}) {
		price := v.(float64)
		if newestStock.CurrentPrice > price {
			if  !greater.Suspend(stock["name"]) {
				greater.notify.Ring(newestStock.Name + ": " + strconv.FormatFloat(newestStock.CurrentPrice,'f',2,64))
			}
		}
	}
}

func (greater *Greater) On()  {
	for _,v1 := range greater.stocks {
		stock := v1.(map[interface{}]interface{})
		greater.greaterMap.Delete(stock["name"])
	}
}

func (greater *Greater) Off(k interface{}) {
	greater.greaterMap.Store(k, 1)
}

func (greater *Greater) Suspend(k interface{}) bool{
	_,ok := greater.greaterMap.Load(k)
	if ok {
		return true
	}else {
		greater.Off(k)
		return false
	}
}
