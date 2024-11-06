package configuration

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/axgle/mahonia"
)

func createLink(src, des string) error {
	src = strings.ReplaceAll(src, "/", "\\")
	des = strings.ReplaceAll(des, "/", "\\")
	// bypass admin permission requirement.
	cmd := exec.Command("cmd", "/c", "mklink", "/J", des, src)

	fmt.Println(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		decoder := mahonia.NewDecoder("gbk")
		return errors.New(string(decoder.ConvertString(string(out))))
	}
	return nil
}
