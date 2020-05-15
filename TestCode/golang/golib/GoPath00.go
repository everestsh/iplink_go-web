//http://www.microembedded.cn/images/3_product.png

//path
package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	//path操作
	fmt.Println("Path操作--------------------")
	fmt.Println("http://www.baidu.com/file/aa.jpg")
	//func Base(path string) string // 返回路径的最后一个元素
	fmt.Println(path.Base("http://www.baidu.com/file/aa.jpg"))
	fmt.Println(path.Base("http://www.microembedded.cn/images/3_product.png")) //3_product.png
	fmt.Println(path.Clean("c:\\file//abc///aa.jpg"))
	fmt.Println(os.Getwd()) ///Users/raylea/Workspace/go/src/test/golang/golib <nil>
	//func Dir(path string) string // 返回路径中目录部分
	fmt.Println(path.Dir("http://www.baidu.com/aa/aaa.jpg")) //
	fmt.Println(path.Dir("/Users/raylea/Workspace/go/src/test/golang/golib/test.txt"))
	//func Ext(path string) string // 返回路径中扩展部分
	fmt.Println(path.Ext("/Users/raylea/Workspace/go/src/test/golang/golib/test.txt"))
	//func IsAbs(path string) bool // 判断是否是一个绝对路径
	fmt.Println(path.IsAbs("./src/test/golang/test.txt"))
	fmt.Println(path.Clean("c:\\file//abc///aa.jpg"))
	//func Join(elem ...string) string // 将多个字符串合并为一个路径
	fmt.Println(path.Join("c:", "aa", "bb", "cc.txt")) //c:/aa/bb/cc.txt
	//func Match(pattern, name string) (matched bool, err error) // 正则是否匹配路径（shell 文件名匹配）
	isMatch, err := path.Match("*.xml", "a.xml")
	fmt.Println(isMatch, err)

	//FilePath操作
	fmt.Println("FilePath操作-----------------")
	fmt.Println(filepath.IsAbs("./raylea/Workspace/go/src/test/golang/golib/test.txt")) //true
	//func Abs(path string) (string, error) // 返回path 相对当前路径的绝对路径
	fmt.Println(filepath.Abs("."))
	fmt.Println(filepath.Base("c:\\aa\\baa.exe"))

}
