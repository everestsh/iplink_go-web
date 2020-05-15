package mecmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/tabwriter"
)

//命令行参数
var (
	flagSet  = flag.NewFlagSet("mecmd", flag.PanicOnError)
	mkdfFlag = flagSet.Bool("mkdf", false, "mkdir os.Args[2] and create os.Args[2]/os.Args[2].go")
	lsdflag  = flagSet.Bool("lsd", false, "list dir")
	rmdFlag  = flagSet.Bool("rmd", false, "rm os.Args[2] （删除目录")
	rmdoFlag = flagSet.Bool("rmdo", false, "rm os.Args[2]/os.Args[2] (删除生成的二进制文件)")
	rmdaFlag = flagSet.Bool("rmda", false, "rm os.Args[2]/os.Args[2] (删除所有目录生成的二进制文件)")
	helpFlag = flagSet.Bool("help", false, "Show this help")
	out      = tabwriter.NewWriter(os.Stdout, 4, 4, 1, ' ', 0)
)

//MeCmd Option
func Mecmd(call []string) error {
	err := flagSet.Parse(call[1:])
	if err != nil {
		return err
	}
	args := flagSet.Args()
	//fmt.Println("11111", flagSet.Args())
	//NArg()获得non-flag个数
	//-help
	if flagSet.NArg() < 0 || *helpFlag {
		println("`mecmd` [options] [dirs...]")
		flagSet.PrintDefaults()
		return nil
	} else if flagSet.NArg() == 0 {
		dirs, e := getDirList()
		if e != nil {
			return e
		}

		for _, dir := range dirs {
			fmt.Println("1111: " + dir)
			if *lsdflag {
				e := list(dir, "")
				if e != nil {
					return e
				}
			} else if *rmdaFlag {
				e := list(dir, "")
				if e != nil {
					return e
				}
			}
		}

	}

	//fmt.Println("hhhhhh", flagSet.Args())
	//创建文件夹及文件
	for _, arg := range args {
		fileBool, _ := pathExists(arg)
		if *mkdfFlag {
			//fmt.Println("22222", arg)
			//fmt.Println(*mkdfFlag)

			if fileBool {
				fmt.Printf("%s文件已经存在！\n", arg)
			} else {
				err := os.Mkdir(arg, os.ModePerm)
				if err != nil {
					log.Fatal(err)
				}
				_, err = os.Create(arg + "/" + arg + ".go")
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s/%s.go 已经创建。\n", arg, arg)
			}
		}

		if *rmdFlag {
			//fmt.Println("22222", arg)
			//fmt.Println(*mkdfFlag)
			err := os.RemoveAll("./" + arg)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s/目录 和 %s/%s.go 文件已经删除。\n", arg, arg, arg)
		}

		fileBoolo, _ := pathExists(arg + "/" + arg)
		if *rmdoFlag {
			if fileBoolo {
				fmt.Printf("%s/%s文件存在！\n", arg, arg)
				err := os.Remove("./" + arg + "/" + arg)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s/%s.go 已经删除。\n", arg, arg)
			} else {
				fmt.Printf("%s/%s文件不存在！\n", arg, arg)
			}
		}

	}
	return nil
}

//golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断:
//如果返回的错误为nil,说明文件或文件夹存在
//如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
//如果返回的错误为其它类型,则不确定是否在存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getDirList() ([]string, error) {
	if flagSet.NArg() <= 0 { //获得non-flag个数
		cwd, e := os.Getwd()
		return []string{cwd}, e
	}
	return flagSet.Args(), nil
}
func list(dir, prefix string) error {
	// ReadDir 读取目录 dirmane 中的所有目录和文件（不包括子目录）
	entries, e := ioutil.ReadDir(dir)
	if e != nil {
		return e
	}

	for _, entry := range entries {
		//printEntry(entry)
		//entry.IsDir()

		if entry.IsDir() {
			//fmt.Fprintf(out, "%s:\n", entry.Name())
			folder := entry.Name()
			//folder := prefix + "/" + entry.Name()
			if *lsdflag {
				fmt.Fprintf(out, "%s\n", folder)
			}

			entries1, e := ioutil.ReadDir(folder)
			if e != nil {
				return e
			}
			for _, entry := range entries1 {
				folder1 := entry.Name()
				if *lsdflag {
					fmt.Fprintf(out, "%s/%s\n", folder, folder1)
				}

				if folder1 == folder {
					//fmt.Fprintf(out, "#####%s/%s#####\n", folder, folder1)
					if *rmdaFlag {
						//fmt.Println("$$$$$$$$$$$$$")
						err := os.Remove("./" + folder + "/" + folder1)
						//err := os.Remove(folder1)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Printf("%s/%s 已经删除。\n", folder, folder1)
					}
				}
			}
			//e := list(dir+"/"+entry.Name(), folder)
			//if e != nil {
			//	return e
			//}
		}

	}
	return nil
}
func printEntry(e os.FileInfo) {
	fmt.Fprintf(out, "%s%s\t", e.Name(), getEntryTypeString(e))
	fmt.Fprintln(out, "")
}

func getEntryTypeString(e os.FileInfo) string {
	if e.IsDir() {
		return "/"
	} else if e.Mode()&os.ModeDevice != 0 {
		return "<>"
	} else if e.Mode()&os.ModeNamedPipe != 0 {
		return ">>"
	} else if e.Mode()&os.ModeSymlink != 0 {
		return "@"
	} else if e.Mode()&os.ModeSocket != 0 {
		return "&"
	} else if e.Mode().IsRegular() && (e.Mode()&0001 == 0001) {
		return "*"
	}
	return ""
}
