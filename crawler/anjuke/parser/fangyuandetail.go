package parser

import (
	"learn/crawler/engine"
	"learn/crawler/model"
	"regexp"
	"strconv"
)

var photoRe = regexp.MustCompile(`<div class="img_wrap">[^<]+<img data-src="(https://pic1.ajkimg.com/display/hj/[^.jpg]+.jpg[?]t=1)"[^>]+>[^<]+</div>`)

var priceRe = regexp.MustCompile(`<span class="price"><em>([0-9]+)</em>元/月</span>`)

var titleRe = regexp.MustCompile(`<h3 class="house-title">([^<]+)</h3>`)

var typeRe = regexp.MustCompile(`<span class="type">([^<]+)</span>`)

var areaRe = regexp.MustCompile(`<span class="info">([0-9]+)平方米</span>`)

var doorRe = regexp.MustCompile(`<li class="house-info-item l-width">[^<]+<span class="type">户型：</span>[^<]+<span class="info">([^<]+)</span>[^<]+</li>`)

var idUrlRe = regexp.MustCompile(`https://[a-z]+.zu.anjuke.com/fangyuan/[0-9]+`)

func ParseFangYuanDetail(
	contents []byte, title string, url string) engine.ParseResult {
	profile := model.Profile{}

	allMatches := photoRe.FindAllSubmatch(contents, -1)

	for _, m := range allMatches {
		profile.Photos = append(profile.Photos, string(m[1]))
	}

	matches := priceRe.FindSubmatch(contents)

	if matches != nil {
		f, err := strconv.Atoi(string(matches[1]))
		if err == nil {
			profile.Price = f
		}
	}

	profile.Title = title //extractString(contents, titleRe)

	profile.Type = extractString(contents, typeRe)

	area, err := strconv.Atoi(extractString(contents, areaRe))
	if err == nil {
		profile.Area = area
	}

	profile.Door = extractString(contents, doorRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "anjuke",
				Id:      extractString([]byte(url), idUrlRe),
				PayLoad: profile,
			},
		},
	}
	return result
}

func extractString(
	contents []byte, re *regexp.Regexp) string {
	matche := re.FindSubmatch(contents)

	if len(matche) >= 2 {
		return string(matche[1])
	} else {
		return ""
	}
}
