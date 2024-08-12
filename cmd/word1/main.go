package main

import (
	"github.com/nguyenthenguyen/docx"
)

func main() {
	// Read from docx file
	r, err := docx.ReadDocxFile("/Users/Golden/Documents/GoProjects/lixin-excel/test/testDocument.docx")
	// Or read from memory
	// r, err := docx.ReadDocxFromMemory(data io.ReaderAt, size int64)

	// Or read from a filesystem object:
	// r, err := docx.ReadDocxFromFS(file string, fs fs.FS)

	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	// Replace like https://golang.org/pkg/strings/#Replace
	docx1.Replace("old_1_1", "new_1_1", -1)
	docx1.ReplaceLink("http://example.com/", "https://github.com/nguyenthenguyen/docx", 1)
	docx1.ReplaceHeader("out with the old", "in with the new")
	docx1.ReplaceFooter("Change This Footer", "new footer")
	docx1.WriteToFile("./new_result_1.docx")

	// Or write to ioWriter
	// docx2.Write(ioWriter io.Writer)

	//docx3 := r.Editable()
	////Currently only swaps apples for apples i.e. png to png, and not png to jpeg etc.
	//docx3.ReplaceImage("word/media/image1.png", "./new.png")
	//
	//// replace the last image
	//imageIndex := docx3.ImagesLen()
	//docx3.ReplaceImage("word/media/image"+strconv.Itoa(imageIndex)+".png", "./new.png")
	//docx3.WriteToFile("./new_result_3.docx")

	r.Close()
}
