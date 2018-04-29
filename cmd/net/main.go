package main

import (
	"os"
	"strings"
	"reflect"
)

var modules = Modules{map[string]string{}}

func init() {
	clazz := reflect.TypeOf(&Modules{})
	for i := 0; i < clazz.NumMethod(); i++ {
		name := clazz.Method(i).Name
		short := string(name[0])
		modules.childCmds[short] = name
	}
}

func main() {
	childCmd := strings.ToUpper(os.Args[1])

	if len(childCmd) == 0 {
		println("No child command specified")
		os.Exit(11)
	}

	if len(childCmd) == 1 {
		value, ok := modules.childCmds[childCmd]
		if ok {
			childCmd = value
		} else {
			println("Unknown child command", childCmd)
			os.Exit(13)
		}
	} else {
		exist := false
		for _, v := range modules.childCmds {
			if v == childCmd {
				exist = true
				break
			}
		}
		if !exist {
			println("Unknown child command", childCmd)
			os.Exit(15)
		}
	}

	for i, v := range os.Args {
		if i <= 1 {
			continue
		}
		os.Args[i-1] = v
	}

	clazz := reflect.ValueOf(&modules)
	clazz.MethodByName(childCmd).Call(nil)
}