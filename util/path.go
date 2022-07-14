package util

import (
	"fmt"
	"github.com/ian-kent/go-log/log"
	"os"
	"path"
)

// GetRootPath 当前项目的根目录；统一testcase和打包后包获取conf目录的方式
// 假设项目目录中存在子目录: conf
func GetRootPath() string {
	// 存在，且可以访问
	dirExists := func(path string) (bool, error) {
		stat, err := os.Stat(path)
		if err == nil && stat.IsDir() {
			return true, nil
		}

		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	filename, err := os.Getwd()
	fmt.Println(fmt.Sprintf("fileName: %v", filename))
	if err == nil {
		dir := path.Dir(filename)
		for i := 0; i < 3; i++ {

			p := path.Join(dir, "conf") // root目录下必须有conf目录

			exists, err := dirExists(p)
			if exists {
				return dir
			}
			if err != nil {
				log.Fatalf("GetRootPath error: %v", err)
			}
			if len(dir) <= 1 {
				break
			}
			dir = path.Dir(dir)
		}

	} else {
		log.Fatalf("os.Getwd() err: %v", err)
		fmt.Println(fmt.Sprintf("os.Getwd() err: %v", err))
	}

	if err != nil {
		panic(err)
	}

	// 默认认为当前目录是根目录
	return filename
}
