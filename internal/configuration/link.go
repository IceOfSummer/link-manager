// 设置环境变量的实现.
package configuration

import (
	"os"
	"path"
)

const appDirectory = "app"

func UseEnv(link *Link) {
	holder := path.Join(DENV_HOME, appDirectory)
	_, err := os.Stat(holder)
	if os.IsNotExist(err) {
		err := os.Mkdir(holder, 0b111_101_101)
		if err != nil {
			panic(err)
		}
	}

	target := path.Join(holder, link.Name)
	_, err = os.Stat(target)
	if !os.IsNotExist(err) {
		os.Remove(target)
	}
	err = createLink(link.Path, target)
	if err != nil {
		panic(err)
	}

}
