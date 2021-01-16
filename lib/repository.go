package lib

import (
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const StockServiceUrl = "http://hq.sinajs.cn/list="

type StockMessage struct {
	Name string
	CurrentPrice float64
}

type Repository struct {

}

func (repository *Repository) Fetch(stockCode string) (StockMessage,error) {
	resp,err := http.Get(StockServiceUrl + stockCode)
	E(err)
	reader := simplifiedchinese.GB18030.NewDecoder().Reader(resp.Body)

	defer resp.Body.Close()
	body,err := ioutil.ReadAll(reader)
	E(err)

	data := string(body)

	r := regexp.MustCompile(`var hq_str_\w+?="(.+)";`)
	rawData := r.FindStringSubmatch(data)
	if rawData == nil {
		E(errors.New("getting failure"))
	}

	fields := strings.Split(rawData[1],",")

	d,_ := strconv.ParseFloat(fields[3],64)
	stockMessage := StockMessage{fields[0],d}

	return stockMessage,nil
}
