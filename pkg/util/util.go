package util

import (
	"fmt"
	"os"
)

func CreateDirIfNotExists(path string) error {
	// 判断目录是否存在
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		// 如果目录不存在，则创建它（包括所有必要的父目录）
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return fmt.Errorf("无法创建目录: %w", err)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("检查目录失败: %w", err)
	}
	return nil
}
