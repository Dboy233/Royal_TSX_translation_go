package main

import (
	"Royal_TSX_translation_go/net"
	"Royal_TSX_translation_go/util"
	"fmt"
	"regexp"
)

//可调参数
const (
	//输入文件、需要翻译的源文件
	inputFile = "Localizable.strings"

	//输出文件、翻译后保存的文件
	outputFile = "zh_Hans.lproj/Localizable.strings"

	//目标语言
	srcLan = "en"

	//翻译语言
	toLan = "zh"

	//正则提取当行中需要翻译的文本 "hello,you" = "***" regexp=>  hello,you
	regex = "\"(.+)\" = \"(?:.+)\";"

	//文件输出格式 "***" = "***"
	outputFormat = "\"%s\" = \"%s\";"
)

func main() {
	fmt.Println("===正在翻译文件请稍等...===")

	//正则解析
	compile := regexp.MustCompile(regex)

	//翻译后的文本保存组
	var lines []string

	//翻译工具
	translator := net.NewTranslator()

	util.ReadFileLines(inputFile, func(line string, err error) {

		//当文件行读取发生异常的时候继续将当前行写入行的文件，以保证程序的正常执行。
		if err != nil {
			lines = append(lines, line)
			fmt.Println("文件当行读取错误", err)
			return
		}
		//提取需要翻译的内容
		matchGroup := compile.FindStringSubmatch(line)

		if len(matchGroup) == 0 {
			//提取失败，原数据写入新文件
			fmt.Println("正则识别错误：", line)
			lines = append(lines, line)
			return
		}

		//获取需要翻译的文本
		query := matchGroup[1]

		//翻译结果保存地址
		var result string

		//执行翻译
		result, err = translator.Translation(query, srcLan, toLan)

		if err != nil {
			//翻译失败的写入新的文件
			fmt.Println("翻译错误：", query, err)
			lines = append(lines, line)
			return
		}

		//标点符号调整
		result = util.PunctuationToEn(result)

		//格式化写入的文本
		wLine := fmt.Sprintf(outputFormat, query, result)

		fmt.Println("翻译结果 : ", wLine)

		lines = append(lines, wLine)

	})

	err := util.WriteFileLines(outputFile, lines)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("===翻译任务结束===")

}
