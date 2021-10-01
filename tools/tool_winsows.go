package tools

import "strings"

// JoinPath windows环境下拼接路径的工具函数
// @params srcPath   string    源路径
// @params itemPath  string    子路径
// @params itemPaths ...string 其他子路径
// @return           string    拼接完成的windows系统文件（夹）路径
func JoinPath(srcPath, itemPath string, itemPaths ...string) (dstPath string) {
	// 拼接第一位子路径
	dstPath = srcPath + "\\" + itemPath
	// 拼接其余路径
	for _, item := range itemPaths {
		dstPath += "\\" + item
	}
	// 替换路径中可能存在的linux路径规则
	dstPath = strings.ReplaceAll(dstPath, "/", "\\")
	return dstPath
}
