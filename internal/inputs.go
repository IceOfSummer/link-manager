package internal

import (
	"errors"
	"strings"
)

func SplitVersion(versionLike string) (string, string, error) {
	sp := strings.Split(versionLike, ":")
	if len(sp) != 2 {
		return "", "", errors.New("` " + versionLike + "` 存在多个 `:`，无法解析别名!")
	}
	return sp[0], sp[1], nil
}
