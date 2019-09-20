package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func outputInfo(file *os.File, err error)  {
	if err == nil{
		fmt.Println("文件打开成功")
		status,_ := file.Stat()
		fmt.Println("文件名：",status.Name(), " 文件权限",status.Mode(), "文件大小： ",status.Size(), status.IsDir())
	}else {
		fmt.Println("文件打开失败,原因：", err)
		return
	}
}

func closeFile(file os.File)  {
	_ = file.Close()
	fmt.Println("文件已关闭")
}

func main()  {
	// open打开是以权限0打开
	file, err := os.Open("/usr/local/var/go/src/golearn/point/testPoint.go")
	outputInfo(file, err)
	defer closeFile(*file)

	// 只读权限
	// perm: 0+自己+组+其他（4r2w1x 读写执行）
	file2, err2 := os.OpenFile("/usr/local/var/go/src/golearn/fileOperate/fileOperate.go", os.O_RDONLY, 0666)
	outputInfo(file2, err2)
	defer closeFile(*file2)

	// bufio
	// 1.创建读取器
	reader := bufio.NewReader(file2)
	for{
		// 2.遍历去读取
		str, err := reader.ReadString('\n')
		if err == nil{
			fmt.Println(str)
		}else {
			if err == io.EOF{
				fmt.Println("已经读取到末尾")
				break
			}else {
				fmt.Println("打开错误，原因：" , err)
				return
			}
		}
	}

	// 开始飞起，使用ioutil进行读取
	bytes, err3 := ioutil.ReadFile("/usr/local/var/go/src/golearn/fileOperate/fileOperate.go")
	if err3 == nil{
		content := string(bytes)
		fmt.Println("工具读取", content)
	}else {
		fmt.Println("打开失败，原因：", err3)
	}

	// 创建文件，|O_APPEND追加写入, |os.O_TRUNC追加模式
	file4, err4 :=os.OpenFile("/usr/local/var/go/src/golearn/fileOperate/aaa.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err4 == nil{
		// 开始写入缓冲区
		writer := bufio.NewWriter(file4)
		inter,err5 := writer.WriteString("忍一时越想越气\n退一步越想越亏\n不服就干")
		if err5 == nil{
			fmt.Println("写入成功", inter)
		}else {
			fmt.Println("无法写入", err5)
		}
		// 缓冲区放入文件
		_ = writer.Flush()
	}else {
		fmt.Println("打开写入失败， 原因：", err4)
	}
	defer closeFile(*file4)

	// 又飞起，使用ioutil开始写入,写入的data是字节切片
	data := `😠😠`
	dataBytes := []byte(data)
	_ = ioutil.WriteFile("/usr/local/var/go/src/golearn/fileOperate/bbb.txt", dataBytes, 0666)

	// 文件拷贝
	bytes2, _ := ioutil.ReadFile("/usr/local/var/go/src/golearn/fileOperate/bbb.txt")
	_ = ioutil.WriteFile("/usr/local/var/go/src/golearn/fileOperate/copy.txt",  bytes2, 0666)

	// io.copy()拷贝
	dstFile,_ := os.OpenFile("/usr/local/var/go/src/golearn/fileOperate/copy2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	srcFile,_ := os.OpenFile("/usr/local/var/go/src/golearn/fileOperate/aaa.txt", os.O_RDONLY,0666)

	written, err8 := io.Copy(dstFile, srcFile)
	if  err8 == nil{
		fmt.Println("复制成功，写入字节数：", written)
	}else {
		fmt.Println("复制失败，原因：", err8)
	}

	// 大文件，缓冲区
	// 1.创建缓冲读入器读取器
	readFile, _ := os.OpenFile("/usr/local/var/go/src/golearn/fileOperate/image.png", os.O_RDONLY,0666)
	writeFile,_ := os.OpenFile("/usr/local/var/go/src/golearn/fileOperate/image2.png", os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)

	readers := bufio.NewReader(readFile)
	writerr := bufio.NewWriter(writeFile)

	// 2.创建大小桶的缓冲区
	buffer := make([]byte, 1024)
	
	// 3.开始一桶一桶写入
	var err9 error
	for err9 != io.EOF && err9 == nil {
		_,err9 = readers.Read(buffer)
		// 开始写入
		writtenByte,_ := writerr.Write(buffer)
		fmt.Println("写入成功，写入字节数：", writtenByte)
	}
	
}