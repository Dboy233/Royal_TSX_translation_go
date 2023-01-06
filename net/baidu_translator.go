package net

import (
	"Royal_TSX_translation_go/util"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
)

// 百度api地址
const questUrl = "https://fanyi-api.baidu.com/api/trans/vip/translate"

// 百度appId
var appId = os.Getenv("BAIDU_APP_ID")

// 百度appKey
var key = os.Getenv("BAIDU_KEY")

func newBaiDuTranslator(frequency, overflowType int) Translator {
	return &privateBaiDuTranslator{
		frequencyUtil: util.FrequencyUtil{
			Frequency:    frequency,
			OverflowType: overflowType,
		},
	}
}

type (
	// privateBaiDuTranslator 百度翻译
	privateBaiDuTranslator struct {
		frequencyUtil util.FrequencyUtil
	}

	//{
	//    "from": "en",
	//    "to": "zh",
	//    "trans_result": [
	//        {
	//            "src": "apple",
	//            "dst": "苹果"
	//        }
	//    ]
	//}
	baiduResult struct {
		From        string `json:"from"`
		To          string `json:"to"`
		TransResult []struct {
			Src string `json:"src"`
			Dst string `json:"dst"`
		} `json:"trans_result"`
	}
)

func (b *privateBaiDuTranslator) Translation(query, srcLan, toLan string) (result string, err error) {

	if len(appId) == 0 {
		return "", errors.New("appId = nil 请配置百度的BAIDU_APP_ID环境变量")
	}

	if len(key) == 0 {
		return "", errors.New("key = nil 请配置百度的BAIDU_KEY环境变量")
	}

	//请求频率控制
	if err := b.frequencyUtil.CheckFrequency(); err != nil {
		return "", err
	}

	//随机数
	salt := fmt.Sprintf("%d", rand.Uint32())
	//签名
	sign := fmt.Sprintf("%x", md5.Sum([]byte(appId+query+salt+key)))

	//组合请求内容
	form := url.Values{
		"q":     {query},
		"from":  {srcLan},
		"to":    {toLan},
		"appid": {appId},
		"salt":  {salt},
		"sign":  {sign},
	}

	//发起翻译请求
	resp, err := http.Get(questUrl + "?" + form.Encode())

	if err != nil {
		return "", err
	}

	defer func() {
		closeError := resp.Body.Close()
		if closeError != nil {
			fmt.Println("关闭请求发生错误：", closeError.Error())
		}
	}()

	r := new(baiduResult)

	buffer := new(bytes.Buffer)

	if _, err := buffer.ReadFrom(resp.Body); err != nil {
		return "", err
	}

	err = json.Unmarshal([]byte(buffer.String()), r)

	if err != nil {
		return "", err
	}

	if len(r.TransResult) == 0 {
		return "", errors.New("没有翻译结果")
	}

	//只返回查询到的第一个
	for _, s := range r.TransResult {
		return s.Dst, nil
	}

	return "", errors.New("没有翻译结果")
}
