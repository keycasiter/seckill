package util

import "log"

// 日志打印
func PrintLog(obj interface{}) {
	log.Println(Marshal(obj))
}
