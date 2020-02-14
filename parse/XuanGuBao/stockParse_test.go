package XuanGuBao

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestStockParse(t *testing.T) {

	bytes, err := ioutil.ReadFile("stockParse.html")
	if err != nil {
		panic("read file err")
	}
	parse := StockParse(bytes)

	fmt.Printf("%s",parse.Item)
}