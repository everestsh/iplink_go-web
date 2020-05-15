package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

//
var (
	h    bool
	mkdf string
	rmd  string
	rmdo string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&mkdf, "mkdf", "", "mkdir os.Args[2] and create os.Args[2]/os.Args[2].go")
	flag.StringVar(&rmd, "rmd", "", "rm os.Args[2] （删除目录）")
	flag.StringVar(&rmd, "rmdo", "", "rm os.Args[2]/os.Args[2] (删除生成的二进制文件)")
	//
	flag.Usage = usage
}

//Option Option
func Option() {
	flag.Parse()
	//if h {
	//	flag.Usage()
	//}
	args, n := getArgsP()
	if n != 0 {
		args = strings.Replace(args, "-", "", -1) //去除字符-
		if n == 1 {
			n = n + 1
		} else {
			n = n - 1
		}

	} else {
		args = "h"
	}
	fmt.Println("n = ", n)
	fmt.Println("hhhhhh", args)

	switch args {
	case "h":
		flag.Usage()
	case "mkdf":
		{
			err := os.Mkdir(os.Args[n], os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			_, err = os.Create(os.Args[n] + "/" + os.Args[n] + ".go")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s/%s.go 已经创建。\n", os.Args[n], os.Args[n])
		}
	case "rmd":
		{

			err := os.RemoveAll("./" + os.Args[n])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s/%s.go 已经删除。\n", os.Args[n], os.Args[n])
		}
	case "rmdo":
		{
			args2 := strings.Replace(os.Args[n], "/", "", -1) //去除字符"/"
			err := os.Remove("./" + args2 + "/" + args2)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `mecmd /0.00.03
Usage: mecmd [-h] [-mkdf string] [-rmd string] [-rmdo string]

Options:
`)
	flag.PrintDefaults()
}

//getArgsP 获取参数列表的带有-的
func getArgsP() (string, int) {
	var v, _ = regexp.Compile("^-")
	i := 0
	for i, n := range os.Args[0:] {
		o := v.Find([]byte(os.Args[i]))
		if string(o) == "-" {
			//fmt.Println(n)
			//fmt.Println(i)
			return n, i
		}
	}
	return "", i
}
