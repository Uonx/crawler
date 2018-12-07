package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(80 * time.Millisecond)

func Fetche(url string) ([]byte, error) {
	<-rateLimiter
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	request.Header.Add("Cookie", "aQQ_ajkguid=9CF3C7E6-0514-1387-BA12-0BB03F507304; twe=2; sessid=76E1AC74-9D13-F464-317D-48F778998B43; _ga=GA1.2.1389409829.1543729425; _gid=GA1.2.2121268512.1543729425; 58tj_uuid=f7f0bde3-0e96-4a16-8e34-5039f95f6961; wmda_uuid=b9dc37c2697c8bd7ce6a36632b0494e2; wmda_new_uuid=1; wmda_visited_projects=%3B6289197098934; als=0; isp=true; Hm_lvt_c5899c8768ebee272710c9c5f365a6d8=1543729937; Hm_lpvt_c5899c8768ebee272710c9c5f365a6d8=1543729999; ajk_member_captcha=b5d199f22294b2d330eca863a3623b58; ajk_bfp=1; lps=http%3A%2F%2Fbj.zu.anjuke.com%2F%7C; ctid=72; __xsptplusUT_8=1; init_refer=; new_uv=6; propertys=kizp6t-pj58cv_; __xsptplus8=8.6.1543811207.1543811216.2%232%7Csp0.baidu.com%7C%7C%7C%25E5%25AE%2589%25E5%25B1%2585%25E5%25AE%25A2%7C%23%23Z1HzBJLG9GHr2Im9H3pE6nzahXSbZwkI%23; new_session=0")
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong url: %s; but status code: %d", url, resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}
