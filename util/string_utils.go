package util

import "strings"

var punctuationMap = map[rune]rune{

	8216: 39, // '

	8217: 39, // '
	//这里注释掉的中文双引号不进行转换，由于翻译软件会将单引号'也翻译为"为了降低后续调整文件的成本，这里直接使用中文双引号，
	//这样在poyal进行读取的时候不会报错，不然会有大面积的内嵌双引号需要修改例如 "name is 'Ok'" = " 名字是"Ok"" 你需要将"OK"改为'OK'
	//8220: 34, // "
	//
	//8221: 34, // "

	12290: 46, // .

	12304: 91, // [

	12305: 93, // ]

	65281: 33, // !

	65288: 40, // (

	65289: 41, // )

	65292: 44, // ,

	65306: 58, // :

	65307: 59, // ;

	65311: 63, // ?

	65371: 123, //{

	65373: 125, //}

	0xff06: 0x0026, //&

	0xff05: 0x0025, //%

	0xff04: 0x0024, //$

	0xff5e: 0x007e, //~
}

// PunctuationToEn 全角符号转换为半角符号
func PunctuationToEn(text string) string {

	text = strings.Map(func(r rune) rune {

		if v, ok := punctuationMap[r]; ok {

			return v

		}

		return r

	}, text)

	return text

}
