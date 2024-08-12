package main

import (
	"fmt"

	"github.com/unidoc/unioffice/common/license"
	"github.com/unidoc/unioffice/document"
)

func init() {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	//err := license.SetMeteredKey(os.Getenv(`UNIDOC_LICENSE_API_KEY`))
	err := license.SetMeteredKey("70ce5e6ed4723181a4faa9e71abec79e16c7c603a4cfb120c34270dfec227a1e")
	if err != nil {
		panic(err)
	}
}

func main() {
	document.New()
	//doc, err := document.Open("../../test/demo.docx")
	doc, err := document.Open("/Users/Golden/Documents/GoProjects/lixin-excel/test/demo.docx")
	if err != nil {
		fmt.Println("打开文档错误:", err)
		return
	}
	defer doc.Close()

	// 修改文档中的一些内容
	para := doc.Paragraphs()[0]
	para.AddRun().AddText("修改后的内容111")

	err = doc.SaveToFile("new_example.docx")
	if err != nil {
		fmt.Println("保存文档错误:", err)
	}
}
