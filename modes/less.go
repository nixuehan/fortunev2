package modes

import (
	. "fortunev2/lib"
	"log"
	"strconv"
	"sync"
)

type Less struct {
	stocks []interface{}
	notify Notify
	lessMap sync.Map
	stockSource *Repository
}

func NewLess() *Less{
	return &Less{}
}

func (less *Less) Init(stocks map[string]interface{}, notify Notify, stockSource *Repository) {
	less.stocks = stocks["less"].([]interface{})
	less.notify = notify
	less.stockSource = stockSource
}

func (less *Less) Iterator() []interface{} {
	return less.stocks
}

func (less *Less) Run(stock map[interface{}]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("%s",err)
		}
	}()

	newestStock,_ := less.stockSource.Fetch(stock["name"].(string))
	for _,v := range stock["price"].([]interface{}) {
		price := v.(float64)
		if newestStock.CurrentPrice < price {
			if  !less.Suspend(stock["name"]) {
				less.notify.Ring(newestStock.Name + ": " + strconv.FormatFloat(newestStock.CurrentPrice,'f',2,64))
			}
		}
	}
}

func (less *Less) On()  {
	for _,v1 := range less.stocks {
		stock := v1.(map[interface{}]interface{})
		less.lessMap.Delete(stock["name"])
	}
}

func (less *Less) Off(k interface{}) {
	less.lessMap.Store(k, 1)
}

func (less *Less) Suspend(k interface{}) bool{
	_,ok := less.lessMap.Load(k)
	if ok {
		return true
	}else {
		less.Off(k)
		return false
	}
}
