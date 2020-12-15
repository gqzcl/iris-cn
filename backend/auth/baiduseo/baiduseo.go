package baiduseo

import (
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"

	"iris-cn/conf"
)

func PushUrl(url string) {
	PushUrls([]string{url})
}

// 百度链接推送
func PushUrls(urls []string) {
	if urls == nil || len(urls) == 0 {
		return
	}
	if len(conf.Instance.BaiduSEO.Site) == 0 || len(conf.Instance.BaiduSEO.Token) == 0 {
		return
	}
	api := "http://data.zz.baidu.com/urls?site=" + conf.Instance.BaiduSEO.Site + "&token=" +
		conf.Instance.BaiduSEO.Token
	body := strings.Join(urls, "\n")
	if response, err := resty.New().R().SetBody(body).Post(api); err != nil {
		logrus.Error(err)
	} else {
		logrus.Info("百度链接提交完成：", string(response.Body()))
	}
}
