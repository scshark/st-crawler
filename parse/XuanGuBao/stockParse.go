package XuanGuBao

import (
	"regexp"
	"st-crawler/engine"
)
// <div class="meta_1SSLH"><div class="title_1yCTf"><a href="/stock/603106.SS" target="_blank"><header> 恒银金融 </header> <span>603106.SS</span></a></div> <div class="price_1b2oR"><span class="number-color-up DIN-Medium">7.88</span></div> <div><span class="number-color-up DIN-Medium">+6.92%</span></div></div>
var stockParseCompile = regexp.MustCompile(
	`<div class="title_[a-zA-Z0-9]+"><a href="/stock/.+".+><header>([^>]+)</header>.*<span>([0-9]+.[A-Z]*)</span>.*<div class="price_[a-zA-Z0-9]+"><span.+>([^<]+)</span>.+>([^<]+%)</span>`)

func StockParse(content []byte) engine.ParseResult{
	header := stockParseCompile.FindAllSubmatch(content,-1)

	parseResult := engine.ParseResult{}
	for _,h := range header{
		if len(h) == 0 {
			continue
		}
		content := string(h[1]) + string(h[2]) +"  "+ string(h[3]) +"  "+ string(h[4])+"  "
		// itemData := make(map[string]string)
		// itemData["stockName"] = string(h[1])
		// itemData["code"] = string(h[2])
		// itemData["price"] = string(h[3])
		// itemData["float"] = string(h[4])

		parseResult.Item = append(parseResult.Item,content)
	}
	return parseResult
}
