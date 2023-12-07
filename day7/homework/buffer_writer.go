package main

import (
	"fmt"
	"os"
)

type BufferedFileWriter struct {
	buffer      [1024]byte
	endPos      int
	fileHandler *os.File
}

func NewBufferedFileWriter(fd *os.File) *BufferedFileWriter {
	return &BufferedFileWriter{
		fileHandler: fd,
	}
}
func (writer *BufferedFileWriter) Flush() {
	if writer.endPos > 0 {
		writer.fileHandler.Write(writer.buffer[:writer.endPos])
		fmt.Println("触发真正的写磁盘")
		writer.endPos = 0
	}
}

func (writer *BufferedFileWriter) Write(content []byte) {

	if len(content) >= 1024 {
		writer.Flush()
		writer.fileHandler.Write(content)
	} else {
		if writer.endPos+len(content) >= 1024 {
			writer.Flush()
			writer.Write(content)
		} else {
			copy(writer.buffer[writer.endPos:], content)
			writer.endPos += len(content)
		}
	}
}

func (writer *BufferedFileWriter) WriteString(content string) {
	writer.Write([]byte(content))
}

func testBufferWriter() {
	fout, err := os.OpenFile("zcy.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()

	writer := NewBufferedFileWriter(fout)
	for i := 0; i < 500; i++ {
		writer.WriteString("034o314o3u1oi43\n")
	}
	writer.Flush()
}
