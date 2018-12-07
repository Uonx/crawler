package parser

import (
	"fmt"
	"learn/crawler/fetcher"
	"testing"
)

func TestParseFangYuanDetail(t *testing.T) {
	contents, err := fetcher.Fetche("https://bj.zu.anjuke.com/fangyuan/1237268666")
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s ", contents)
	result := ParseFangYuanDetail(contents, "双周保洁 物业齐全 智能门锁双桥", "https://bj.zu.anjuke.com/fangyuan/1237268666")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 "+
			"element; but was %v", result.Items)
	}
	//actual := result.Items[0]
	fmt.Printf("%s", result.Items)

}
