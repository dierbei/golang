package logxlt

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	defaultLogFormat        = "2006-01-02-15-04"
	defaultLogOutPutDir     = "./logs"
	defaultLogDuration      = time.Second
	defaultLogOldFileMaxDay = 3
)

func TestLoggerFile() {
	logger := Logger{}
	if err := logger.Setup(); err != nil {
		log.Printf("xiao la tiao [Setup] failed, err: %v\n", err.Error())
		return
	}

	i := 1
	for {
		if _, err := logger.Write(fmt.Sprintf("xiao la tiao %d\n", i)); err != nil {
			log.Printf("xiao la tiao [Write] failed, err: %v\n", err.Error())
			return
		}
		fmt.Printf("xiao la tiao %d\n", i)
		i++

		time.Sleep(time.Second)

		if i%10 == 0 {
			fmt.Println(logger.oldFileList)
		}
	}
}

type Logger struct {
	lastFile      *os.File      //最后创建的日志文件
	lastFileDate  *time.Time    //最后创建日志文件的时间
	FileFormat    string        //文件格式
	LoopDuration  time.Duration //循环检查的时间间隔
	oldFileList   []string      //保存旧文件名
	oldFileMaxDay int           //最多保存多少天旧文件
}

// Write 向日志文件中输入文本
func (l *Logger) Write(content string) (int, error) {
	return l.lastFile.Write([]byte(content))
}

// autoRemoveOldFile 自动删除旧的日志文件
func (l *Logger) autoRemoveOldFile() (err error) {
	if len(l.oldFileList) > l.oldFileMaxDay {
		if err := l.remove(l.oldFileList[0]); err != nil {
			return err
		}
		l.oldFileList = l.oldFileList[1:]
	}

	return nil
}

// remove 删除指定的日志文件
func (l *Logger) remove(name string) error {
	runningDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// example: antl-super-automation/logs/2021-07-31-17-08.log
	return os.Remove(runningDir + defaultLogOutPutDir[1:] + "/" + name)
}

// Setup 初始化Logger
// LoopDuration 默认间隔1秒
// FileFormat 默认文件名格式 2006-01-02-15-04
// oldFileMaxSize 默认最多保存3天旧日志文件
func (l *Logger) Setup() error {

	if l.LoopDuration.Seconds() == 0 {
		l.LoopDuration = defaultLogDuration
	}

	if l.FileFormat == "" {
		l.FileFormat = defaultLogFormat
	}

	if l.oldFileMaxDay == 0 {
		l.oldFileMaxDay = defaultLogOldFileMaxDay
	}

	// example: C:\Users\16286\Desktop\new-dir\new-dir\antl-super-automation
	runningDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// example: antl-super-automation/logs
	logDir := runningDir + defaultLogOutPutDir[1:]

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			return err
		}
	}

	if err := l.setLogFile(); err != nil {
		return err
	}

	go func() {
		for {
			l.setLogFile()
			time.Sleep(l.LoopDuration)
		}
	}()

	return nil
}

// setLogFile 创建日志文件
func (l *Logger) setLogFile() error {
	if l.needNewFile() {
		// example: 2021-07-31-10-56.log
		fileName := fmt.Sprintf("%s.log", time.Now().Format(l.FileFormat))

		// example: antl-super-automation/logs/2021-07-31-10-56.log
		path := fmt.Sprintf("%s/%s", defaultLogOutPutDir, fileName)

		newFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}

		// 关闭旧的文件
		if l.lastFile != nil {
			if err := l.lastFile.Close(); err != nil {
				return err
			}
		}

		// 设置新文件
		l.lastFile = newFile
		now := time.Now()
		l.lastFileDate = &now
		l.oldFileList = append(l.oldFileList, fileName)

		return l.autoRemoveOldFile()
	}

	return nil
}

// needNewFile 是否需要新建日志文件
func (l *Logger) needNewFile() bool {
	// 程序第一次运行需要创建日志文件
	if l.lastFileDate == nil || l.lastFile == nil {
		return true
	}

	now := time.Now().Format(l.FileFormat)
	last := l.lastFileDate.Format(l.FileFormat)

	// example: 2021-07-31-10-56.log  2021-07-31-10-57.log
	// 2006-01-02-15-04
	// 年----月-日-时-分
	// 默认格式是每分钟生成一个日志文件
	if now != last {
		return true
	}

	return false
}
