package rhino

import (
	"github.com/Dasongzi1366/AutoGo/utils"
)

func Eval(contextId, script string) string {
	if script == "" {
		return ""
	}
	if contextId == "" {
		contextId = "__TEMP__"
	}
	return utils.CallJavaMethod("js", contextId+"|"+script)
}
