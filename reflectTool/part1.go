package reflectTool

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func GetFunName() string {
	pc, file, line, ok := runtime.Caller(1)
	_ = file
	_ = line
	_ = ok
	f := runtime.FuncForPC(pc)
	names := strings.Split(f.Name(), ".")
	funcName := names[len(names)-1]
	return funcName
}

//动态调用方法n次并统计时长
func InvokeMethod(times int, testObj interface{}, methodName string, args []reflect.Value) []reflect.Value {
	t1 := time.Now().UnixNano()
	var result []reflect.Value
	for i := 0; i < times; i++ {
		values := reflect.ValueOf(testObj).MethodByName(methodName).Call(args)
		if i == 0 {
			result = values
		}
	}
	t2 := time.Now().UnixNano()
	runTime := (t2 - t1) / 1e6 //time.Since(t1)//纳秒转成毫秒 并省略小数点后面
	fmt.Println("run time : ", runTime, "毫秒")
	return result
}
