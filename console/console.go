package console

/*
#include "imgui.h"
#include <stdlib.h>
#cgo arm64 LDFLAGS: -L../../resources/libs/arm64-v8a -limgui
#cgo amd64 LDFLAGS: -L../../resources/libs/x86_64 -limgui
#cgo 386 LDFLAGS: -L../../resources/libs/x86 -limgui
*/
import "C"
import (
	"fmt"
	"github.com/Dasongzi1366/AutoGo/device"
	"github.com/Dasongzi1366/AutoGo/https"
	"github.com/Dasongzi1366/AutoGo/utils"
	"os"
	"runtime"
	"strings"
	"unsafe"
)

var isInit = false

func Init(noCaptureMode bool) {
	isInit = true
	if int(C.CheckInit()) == 0 {
		if runtime.GOARCH != "arm64" {
			noCaptureMode = true
			_, err := os.Stat("/system/fonts/NotoSansCJK-Regular.ttc")
			if err != nil {
				_, err = os.Stat("/system/fonts/NotoSerifCJK-Regular.ttc")
				if err != nil {
					_, err = os.Stat("/data/local/tmp/NotoSansCJK-Regular.ttc")
					if err != nil {
						fmt.Println("imgui初始化中..")
						code, data := https.Get("https://vip.123pan.cn/1823847070/AutoGo/NotoSansCJK-Regular.ttc", 20000)
						if code == 200 {
							fmt.Println("imgui初始化完毕")
							os.WriteFile("/data/local/tmp/NotoSansCJK-Regular.ttc", data, 0644)
						} else {
							fmt.Println("imgui初始失败,下载依赖文件超时,中文可能显示乱码")
						}
					}
				}
			}
		}
		go C.Init(C.int(b2i(noCaptureMode)))
		success := false
		for i := 0; i < 100; i++ {
			utils.Sleep(10)
			if int(C.CheckInit()) == 1 {
				success = true
				break
			}
		}
		if !success {
			fmt.Fprintf(os.Stderr, "[AutoGo] imgui初始化失败")
			os.Exit(1)
		}
		w := intMin(device.Width, device.Height)
		h := intMax(device.Width, device.Height)
		if w < 1080 {
			C.Toast_setTextSize(35)
		} else if h > 1920 {
			C.Toast_setTextSize(50)
		}
		scale := float32(1080) / float32(w)
		C.Console_setPosition(25, C.int(int(float32(50)/scale)))
		if device.Width > device.Height {
			C.Console_setSize(C.int(int(float32(800)/scale)), C.int(int(float32(520)/scale)))
		} else {
			C.Console_setSize(C.int(int(float32(520)/scale)), C.int(int(float32(800)/scale)))
		}
	}
}

func SetWindowSize(width, height int) {
	C.Console_setSize(C.int(width), C.int(height))
}

func SetWindowPosition(x, y int) {
	C.Console_setPosition(C.int(x), C.int(y))
}

func SetWindowColor(color string) {
	cColor := C.CString(color)
	defer C.free(unsafe.Pointer(cColor))
	C.Console_setWindowColor(cColor)
}

func SetTextColor(color string) {
	cColor := C.CString(color)
	defer C.free(unsafe.Pointer(cColor))
	C.Console_setTextColor(cColor)
}

// SetTextSize 设置字体大小
func SetTextSize(size int) {
	C.Console_setTextSize(C.int(size))
}

func Println(a ...any) {
	str := fmt.Sprint(a...)
	arr := strings.Split(str, "\n")
	for _, line := range arr {
		cLine := C.CString(line)
		defer C.free(unsafe.Pointer(cLine))
		C.Console_println(cLine)
	}
}

func Clear() {
	C.Console_clear()
}

func Show() {
	if !isInit {
		Init(false)
	}
	C.Console_show()
}

func Hide() {
	C.Console_hide()
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func intMax(a, b int) int {
	if a < b {
		return b
	}
	return a
}
