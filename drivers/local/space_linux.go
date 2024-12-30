//go:build linux
// +build linux

package local

import (
	"syscall"

	"github.com/alist-org/alist/v3/internal/model"
)

func getDiskSpace(path string) model.Space {
	var stat syscall.Statfs_t
	err := syscall.Statfs(path, &stat)
	if err != nil {
		return model.Space{}
	}

	total := int64(stat.Blocks) * int64(stat.Bsize)
	free := int64(stat.Bfree) * int64(stat.Bsize)
	usage := total - free

	return model.Space{Usage: usage, Total: total}
}
