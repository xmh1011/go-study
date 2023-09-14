package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 定义PDF文件和输出Word文件的路径
	pdfPath := "/path/to/pdf/file.pdf"
	wordPath := "/path/to/word/file.docx"
	
	// 使用Unoconv将PDF转换为Word文档
	cmd := exec.Command("unoconv", "-f", "docx", "-o", wordPath, pdfPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error converting PDF to Word:", err)
	} else {
		fmt.Println("PDF converted to Word successfully!")
	}
}
