package net

import "Royal_TSX_translation_go/util"

// Translator 翻译接口
type Translator interface {
	//Translation
	//@src 		要翻译的文本
	//@srcLan 	文本的语言码
	//@toLan	翻译成语言码
	Translation(query, srcLan, toLan string) (result string, err error)
}

func NewTranslator() Translator {
	return &TranslatorProxy{}
}

// TranslatorProxy 翻译代理由谁翻译它说了算
type TranslatorProxy struct {
}

// 百度翻译
var baiduTranslator = newBaiDuTranslator(10, util.OVER_FLOWTYPE_AWAIT)

func (t TranslatorProxy) Translation(query, srcLan, toLan string) (result string, err error) {
	translator := baiduTranslator
	translation, err := translator.Translation(query, srcLan, toLan)
	//todo 如果有其他翻译平台，进行异常,负载均衡处理
	//if err != nil {
	//	//切换其他翻译平台
	//}
	return translation, err
}
