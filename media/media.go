package media

import (
	"github.com/Dasongzi1366/AutoGo/utils"
	"regexp"
	"strings"
)

// ScanFile 扫描路径path的媒体文件，将它加入媒体库中
func ScanFile(path string) {
	utils.Shell("am broadcast -a android.intent.action.MEDIA_SCANNER_SCAN_FILE -d \"file://" + path + "\"")
	mediaPath := strings.Replace(path, "/sdcard", "/storage/emulated/0", 1)
	result := utils.Shell(`content query --uri content://media/external/images/media`)
	lines := strings.Split(result, "\n")
	var mediaID string
	for _, line := range lines {
		if strings.Contains(line, "_data="+mediaPath) {
			re := regexp.MustCompile(`\b_id=([0-9]+)\b`)
			match := re.FindStringSubmatch(line)
			if len(match) == 2 {
				mediaID = match[1]
			}
			break
		}
	}
	if mediaID != "" {
		utils.Shell("content update --uri content://media/external/images/media/" + mediaID + " --bind is_pending:i:0")
	}
}
