package main

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func readFile(inPath string, writer *bufio.Writer, writer2 *zlib.Writer) {
	if fin, err := os.Open(inPath); err != nil {
		fmt.Println("err")
		return
	} else {
		defer fin.Close()

		reader := bufio.NewReader(fin)
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				if err == io.EOF {
					if len(line) > 0 {
						writer.WriteString(line)
						writer.WriteString("\n")
						writer2.Write([]byte(line))
						writer2.Write([]byte{'\n'})
					}
				}
				break
			} else {
				writer.WriteString(line)
				writer2.Write([]byte(line))
			}
		}
	}

}

func mergeFile(dir string) {
	fout, err := os.OpenFile("big.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout.Close()

	writer := bufio.NewWriter(fout)

	fout2, err := os.OpenFile("big.zlib", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fout2.Close()

	writer2 := zlib.NewWriter(fout2)

	if files, err := ioutil.ReadDir(dir); err != nil {
		fmt.Println(err)
		return
	} else {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			baseName := file.Name()
			if strings.HasSuffix(baseName, ".txt") {
				inPath := filepath.Join(dir, baseName)
				readFile(inPath, writer, writer2)
			}
		}
	}

	writer.Flush()
	writer2.Flush()
}
