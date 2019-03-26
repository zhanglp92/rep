package base

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// desc: 路径相关工具

// GetPathByGopath 通过GOPATH 匹配获取path的绝对路径
func GetPathByGopath(path string) string {
	for _, root := range strings.Split(os.Getenv("GOPATH"), ":") {
		absp := filepath.Join(root, path)
		if _, err := os.Stat(absp); err == nil || os.IsExist(err) {
			return absp
		}
	}
	return ""
}

// CurrentFuncName 获取当前函数名
func CurrentFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	segs := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	if len(segs) <= 0 {
		return ""
	}
	return segs[len(segs)-1]
}

// CurrentBaseName 获取当前文件名(不包含后缀)
func CurrentBaseName() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
}

// CurrentDirBaseName 获取当前目录名
func CurrentDirBaseName() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filepath.Base(filepath.Dir(filename))
}

// CurrentDir 获取当前目录
func CurrentDir() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return ""
	}
	return filepath.Dir(filename)
}
