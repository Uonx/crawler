package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	//contents, err := fetcher.Fetche("https://www.anjuke.com/sy-city.html")
	contents, err := ioutil.ReadFile(
		"citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s/n", contents)
	result := ParseCityList(contents, "")

	const resultSize int = 791

	expectedUrls := []string{
		"https://beijing.anjuke.com", "https://shanghai.anjuke.com", "https://guangzhou.anjuke.com",
	}

	expectedCities := []string{
		"北京", "上海", "广州",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expcted url #%d: %s; but"+
				"was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d "+
			"requests; but had %d",
			resultSize, len(result.Items))
	}

}
