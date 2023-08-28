package filex

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// IsPathExist 路径是否存在，无论路径是文件还是文件夹，只要存在，则放回true，否则返回false
func IsPathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// IsFileExist 文件是否存在，仅仅验证文件，如果路径是个文件夹,返回false
func IsFileExist(path string) bool {
	f, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return !f.IsDir()
}

// IsDirExist 文件夹是否存在
func IsDirExist(path string) bool {
	f, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return f.IsDir()
}

// IsLink 检查路径是否为链接
func IsLink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeSymlink != 0
}

// IsZipFile 是否为zip文件
func IsZipFile(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// CreateFile 创建文件
func CreateFile(path string) bool {
	file, err := os.Create(path)
	if err != nil {
		return false
	}

	defer file.Close()
	return true
}

// CreateDir 创建文件夹
func CreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// RemoveDir 删除文件夹
func RemoveDir(dir string) error {
	return os.RemoveAll(dir)
}

// Copy 拷贝文件夹或者文件
func Copy(src, dst string) error {
	srcInfo, srcErr := os.Lstat(src)
	if srcErr != nil {
		return srcErr
	}
	_, dstErr := os.Lstat(dst)
	if dstErr == nil {
		// TODO(rog) add a flag to permit overwriting?
		return fmt.Errorf("will not overwrite %q", dst)
	}
	if !os.IsNotExist(dstErr) {
		return dstErr
	}
	switch mode := srcInfo.Mode(); mode & os.ModeType {
	case os.ModeSymlink:
		return copySymLink(src, dst)
	case os.ModeDir:
		return copyDir(src, dst, mode)
	case 0:
		return copyFile(src, dst, mode)
	default:
		return fmt.Errorf("cannot copy file with mode %v", mode)
	}
}

func copySymLink(src, dst string) error {
	target, err := os.Readlink(src)
	if err != nil {
		return err
	}
	return os.Symlink(target, dst)
}

func copyFile(src, dst string, mode os.FileMode) error {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()
	dstf, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, mode.Perm())
	if err != nil {
		return err
	}
	defer dstf.Close()
	// Make the actual permissions match the source permissions
	// even in the presence of umask.
	if err := os.Chmod(dstf.Name(), mode.Perm()); err != nil {
		return err
	}
	if _, err := io.Copy(dstf, srcf); err != nil {
		return fmt.Errorf("cannot copy %q to %q: %v", src, dst, err)
	}
	return nil
}

func copyDir(src, dst string, mode os.FileMode) error {
	srcf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcf.Close()
	if mode&0500 == 0 {
		// The source directory doesn't have write permission,
		// so give the new directory write permission anyway
		// so that we have permission to create its contents.
		// We'll make the permissions match at the end.
		mode |= 0500
	}
	if err := os.Mkdir(dst, mode.Perm()); err != nil {
		return err
	}
	for {
		names, err := srcf.Readdirnames(100)
		for _, name := range names {
			if err := Copy(filepath.Join(src, name), filepath.Join(dst, name)); err != nil {
				return err
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("error reading directory %q: %v", src, err)
		}
	}
	if err := os.Chmod(dst, mode.Perm()); err != nil {
		return err
	}
	return nil
}

// FileMode 获取文件权限和模式
func FileMode(path string) (fs.FileMode, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}
	return fi.Mode(), nil
}

// CurrentPath 当前路径
func CurrentPath() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		absPath = path.Dir(filename)
	}

	return absPath
}

// FileSize 文件大小，返回大小单位为byte
func FileSize(path string) (int64, error) {
	f, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return f.Size(), nil
}

// MTime 文件修改时间戳
func MTime(filepath string) (int64, error) {
	f, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return f.ModTime().Unix(), nil
}

// Sha 返回文件的sha值, param `shaType` should be 1, 256 or 512.
func Sha(filepath string, shaType ...int) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha1.New()
	if len(shaType) > 0 {
		if shaType[0] == 1 {
			h = sha1.New()
		} else if shaType[0] == 256 {
			h = sha256.New()
		} else if shaType[0] == 512 {
			h = sha512.New()
		} else {
			return "", errors.New("param `shaType` should be 1, 256 or 512.")
		}
	}

	_, err = io.Copy(h, file)

	if err != nil {
		return "", err
	}

	sha := fmt.Sprintf("%x", h.Sum(nil))

	return sha, nil
}

// MiMeType return file mime type
// param `file` should be string(file path) or *os.File.
func MiMeType(file any) string {
	var mediatype string

	readBuffer := func(f *os.File) ([]byte, error) {
		buffer := make([]byte, 512)
		_, err := f.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}

	if filePath, ok := file.(string); ok {
		f, err := os.Open(filePath)
		if err != nil {
			return mediatype
		}
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}

	if f, ok := file.(*os.File); ok {
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}
	return mediatype
}
