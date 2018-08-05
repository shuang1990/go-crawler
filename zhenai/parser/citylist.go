package parser

import (
	"crawler/engine"
	"regexp"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 10
	for _, match := range matches {
		result.Items = append(result.Items, "City " + string(match[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParserFunc: ParseCity,
		})
		limit --
		if limit <= 0 {
			break
		}
	}
	return result
}
