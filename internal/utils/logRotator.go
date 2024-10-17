package utils

import (
	// "fmt"
	"os"
	"sync"
	"time"
	"strings"
)

type RotateWriter struct {
	Lock		sync.Mutex
	FilePath	string
	FileName	string
	File		*os.File
}

func NewLogRotator(filePath, fileName string) *RotateWriter {
	rw := RotateWriter{
		FilePath:		filePath,
		FileName:		fileName,
	}
	if err := rw.Rotate(); err != nil {
		return nil
	}
	return &rw
}

func (rw *RotateWriter) Write(output []byte) (int, error) {
	rw.Lock.Lock()
	defer rw.Lock.Unlock()
	return rw.File.Write(output)
}

func (rw *RotateWriter) Rotate() error {
	rw.Lock.Lock()
	defer rw.Lock.Unlock()
	if rw.File != nil {
		if err := rw.File.Close(); err != nil {
			return err
		}
		rw.File = nil
	}
	_, err := os.Stat(rw.FilePath)
	if err == nil {
		fn := strings.Split(rw.FilePath,"/")
		fileName := fn[len(fn) - 1]
		fnNoExt := strings.Split(fileName, ".")
		newFn := fnNoExt[0] + "_" + time.Now().Format("2006-01-02") + ".log"
		newFn = strings.Join(fn[:len(fn) - 1], "/") + "/" + newFn
		if err := os.Rename(rw.FilePath, newFn); err != nil {
			return err
		}
	}
	rw.File, err = os.OpenFile(rw.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, DefaultFileMod)
	if err != nil {
		return err
	}
	return nil
}