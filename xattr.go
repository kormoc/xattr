package xattr

import "github.com/kormoc/xattr/xattrsyscall"

func Remove(filePath string, xattrName string) error {
    return syscallErrorToXAttrError(xattrsyscall.Removexattr(filePath, xattrName))
}
