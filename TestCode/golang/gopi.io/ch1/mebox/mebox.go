package main

import (
	"flag"
)

const (
	VERSION = "0.00.06"
)

var (
	//NewFlagSet的第一个参数是可以任意定的.但第二个参数,则决定了参数解析出错时错误处理方式
	//flag.ExitOnError------ Call os.Exit(2)
	flagSet = flag.NewFlagSet("mebox", flag.ExitOnError)
	// Bool定义一个具有指定名称，默认值和用法字符串的布尔标志。
	//返回值是存储标志值的bool变量的地址。
	listFlag = flagSet.Bool("list", false, "List applets")
	//String定义一个具有指定名称，缺省值和用法字符串的字符串标志。
	//返回值是存储标志值的字符串变量的地址。
	installFlag = flagSet.String("install", "", "Create symlinks for applets in given path")
	helpFlag    = flagSet.Bool("help", false, "Show help")
)

func Mebox(call []string) (e error) {
	// Parse解析参数列表中的标志定义，不应该这样做
	//包含命令名称。必须在FlagSet中的所有标志之后被调用
	//是在程序访问标志之前定义的。
	//如果设置了-help或者-h但是没有定义，返回值将会是ErrHelp。
	e = flagSet.Parse(call[1:]) //解析
	if e != nil {
		return
	}

	if *listFlag {
		list()
	} else if *installFlag != "" {
		//e = install(*installFlag)
	} else {
		help()
	}
	return
}

func help() {
	println("`mebox` [options]")
	flagSet.PrintDefaults()
	println()
	println("Version", VERSION)
	list()
}

func list() {
	println("List of compiled applets:\n")
	for name := range Applets {
		print(name, ", ")
	}
	println("")
}

/*
func install(path string) error {
	goboxpath, e := common.GetGoboxBinaryPath()
	if e != nil {
		return e
	}
	for name := range Applets {
		// Don't overwrite the executable
		if name == "mebox" {
			continue
		}
		newpath := filepath.Join(path, name)
		e = common.ForcedSymlink(goboxpath, newpath)
		if e != nil {
			common.DumpError(e)
		}
	}
	return nil
}
*/
