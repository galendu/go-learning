package main

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/bytedance/sonic"
)

type Student struct {
	Name string
	Age  int
}

func jsonDemo() {
	stu := Student{"zcy", 18}
	bs, _ := json.Marshal(stu)
	fmt.Println(string(bs))

	bs, _ = sonic.Marshal(stu)
	fmt.Println(string(bs))
}

func base64Demo() {
	bs := []byte{1, 4, 2, 6, 9, 4}
	str := base64.StdEncoding.EncodeToString(bs)
	fmt.Println(str)

	if cont, err := base64.StdEncoding.DecodeString(str); err == nil {
		fmt.Println(cont)
	} else {
		fmt.Printf(err.Error())
	}
}

func compressDemo() error {
	fin, err := os.Open("main.go")
	if err != nil {
		return err
	}
	defer fin.Close()

	stat, _ := fin.Stat()
	fmt.Printf("压缩前文件的大小 %d\n", stat.Size())

	fout, err := os.OpenFile("main.zlib", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}

	defer fout.Close()

	bs := make([]byte, 1024)
	writer := zlib.NewWriter(fout)
	for {
		n, err := fin.Read(bs)
		if err == nil {
			writer.Write(bs[:n])
		} else {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				break
			}
		}
	}

	writer.Close()

	fin, err = os.Open("main.zlib")
	if err != nil {
		return err
	}
	defer fin.Close()

	stat, _ = fin.Stat()
	fmt.Printf("压缩后文件的大小 %d\n", stat.Size())

	reader, err := zlib.NewReader(fin)
	if err != nil {
		return err
	}
	io.Copy(os.Stdout, reader)
	reader.Close()
	fin.Close()

	return nil
}

func main() {
	// base64Demo()
	compressDemo()

}
