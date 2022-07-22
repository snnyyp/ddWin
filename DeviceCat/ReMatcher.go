package DeviceCat

import "regexp"

/*
几种路径类型
简单来说，就是有Path的前面有路径前缀
HarddiskVolume -> HarddiskVolume1
HarddiskVolumePath -> \\.\HarddiskVolume1
DriveLetter -> C:
DriveLetterPath -> \\.\C:
*/
var (
	HarddiskVolumeMatcher  = regexp.MustCompile("HarddiskVolume[0-9]+")
	DriveLetterPathMatcher = regexp.MustCompile(`\\\\.\\[a-zA-Z]:`)
)

func IsHarddiskVolume(path string) bool {
	return HarddiskVolumeMatcher.MatchString(path)
}

func IsDriveLetterPath(path string) bool {
	return DriveLetterPathMatcher.MatchString(path)
}
