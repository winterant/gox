package xlog

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"
)

func redirectOutputFile(path string, fd int) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	if err := syscall.Dup2(int(file.Fd()), fd); err != nil {
		panic(err)
	}
	return file
}

func rotateOutputFile(ctx context.Context, filePathPrefix string, fd int) {
	genFilePath := func() string {
		return fmt.Sprintf("%s%s.log", filePathPrefix, time.Now().Format("20060102"))
	}

	// init
	currentFile := redirectOutputFile(genFilePath(), fd)

	// rotate
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			filePath := genFilePath()
			if currentFile.Name() != filePath {
				nextFile := redirectOutputFile(filePath, fd)
				fmt.Println(currentFile)
				if currentFile != nil {
					currentFile.Close()
					// Remove if file is empty.
					if info, err := os.Stat(currentFile.Name()); err == nil && info.Size() == 0 {
						_ = os.Remove(currentFile.Name())
					}
				}
				currentFile = nextFile
			}

			time.Sleep(10 * time.Minute)
		}
	}()
}

func HookStdout(ctx context.Context, filePathPrefix string) {
	rotateOutputFile(ctx, filePathPrefix, int(os.Stdout.Fd()))
}

func HookStderr(ctx context.Context, filePathPrefix string) {
	rotateOutputFile(ctx, filePathPrefix, int(os.Stderr.Fd()))
}
