package logxlt

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	// 默认日志文件格式
	defaultLogFormat          = "2006-01-02-15-04"
	// 默认日志文件输出文件夹
	defaultLogOutPutDir       = "./logs"
	// 默认循环检查时间间隔
	defaultLogDuration        = time.Second
	// 默认保存多少天的日志文件
	defaultLogOldFileMaxDay   = 7
	// 默认几天压缩一次日志文件
	defaultLogFileCompressDay = 7
	// 默认日志文件压缩之后的存放文件夹
	defaultLogFileCompressDir = "compresslog/"
)

func TestLoggerFile() {
	i := 1
	for {
		if _, err := XltLog.Write(fmt.Sprintf("xiao la tiao %d\n", i)); err != nil {
			log.Printf("xiao la tiao [Write] failed, err: %v\n", err.Error())
			return
		}
		fmt.Printf("xiao la tiao %d\n", i)
		i++

		time.Sleep(time.Second)

		if i%10 == 0 {
			fmt.Println(XltLog.oldFileList)
		}
	}
}

var XltLog = &Logger{}

func init() {
	if err := XltLog.Setup(); err != nil {
		log.Printf("xiao la tiao [Setup] failed, err: %v\n", err.Error())
		return
	}
}

type Logger struct {
	FileFormat   string        //文件格式
	LoopDuration time.Duration //循环检查的时间间隔
	NoticeChan   chan struct{} //重启V2Ray通知管道

	lastFile      *os.File   //最后创建的日志文件
	lastFileDate  *time.Time //最后创建日志文件的时间
	oldFileList   []string   //保存旧文件名
	oldFileMaxDay int        //最多保存多少天旧文件
	isRestart     bool       //是否需要重启
	isCompress    int        //几天进行压缩
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
	runningDir, err := l.getRunningDirPath()
	if err != nil {
		return err
	}

	// example: antl-super-automation/logs/2021-07-31-17-08.log
	return os.Remove(runningDir + defaultLogOutPutDir[1:] + "/" + name)
}

// Setup 初始化Logger
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

	if l.NoticeChan == nil {
		l.NoticeChan = make(chan struct{}, 1)
	}

	if err := l.createLogFileDir(); err != nil {
		return err
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

// createLogFileDir 创建保存日志文件的文件夹
func (l *Logger) createLogFileDir() error {
	// example: C:\Users\16286\Desktop\new-dir\new-dir\antl-super-automation
	runningDirPath, err := l.getRunningDirPath()
	if err != nil {
		return err
	}

	// example: antl-super-automation/logs
	logDir := runningDirPath + defaultLogOutPutDir[1:]

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.Mkdir(logDir, 0755); err != nil {
			return err
		}
	}

	return nil
}

// GetLogFile 获取日志文件
func (l *Logger) GetLogFile() *os.File {
	return l.lastFile
}

// setLogFile 尝试获取新的日志文件
func (l *Logger) setLogFile() error {
	if l.needNewFile() {

		if err := l.needCompress(); err != nil {
			return err
		}

		_, err := l.getNewLogFile()
		if err != nil {
			return err
		}

		// 是否需要重启V2Ray
		if l.isRestart {
			l.NoticeChan <- struct{}{}
		} else {
			l.isRestart = true
		}

		return l.autoRemoveOldFile()
	}

	return nil
}

// getNewLogFile 新建日志文件 关闭旧的日志 保存文件信息
func (l *Logger) getNewLogFile() (*os.File, error) {
	// example: 2021-07-31-10-56.log
	fileName := fmt.Sprintf("%s.log", time.Now().Format(l.FileFormat))
	// example: antl-super-automation/logs/2021-07-31-10-56.log
	path := fmt.Sprintf("%s/%s", defaultLogOutPutDir, fileName)

	newFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	if l.lastFile != nil {
		if err := l.lastFile.Close(); err != nil {
			return nil, err
		}
	}

	l.lastFile = newFile
	now := time.Now()
	l.lastFileDate = &now
	l.oldFileList = append(l.oldFileList, fileName)

	return newFile, nil
}

// needNewFile 是否需要新建日志文件
func (l *Logger) needNewFile() bool {
	// 程序第一次运行需要创建日志文件
	if l.lastFileDate == nil || l.lastFile == nil {
		return true
	}

	now := time.Now().Format(l.FileFormat)
	last := l.lastFileDate.Format(l.FileFormat)

	// example:
	//		last:2021-07-31-10-56.log  now:2021-07-31-10-57.log
	// 格式: 2006-01-02-15-04
	// 		年----月-日-时-分
	if now != last {
		return true
	}

	return false
}

func (l *Logger) needCompress() error {
	l.isCompress++
	if l.isCompress%(defaultLogFileCompressDay+1) == 0 {

		runningDirPath, err := l.getRunningDirPath()
		if err != nil {
			return err
		}

		// example: dst=dst=antl-super-automation/compresslog/2021-08-02-11-27~2021-08-02-11-33.tar.gz
		if err := l.compressLogFileToTgz(defaultLogOutPutDir[2:],
			fmt.Sprintf("%s/%s%s.tar.gz", runningDirPath, defaultLogFileCompressDir, l.getCompressLogFileName())); err != nil {
			return err
		}
		l.isCompress = 1
	}

	return nil
}

// Tar 压缩文件
func (l *Logger) compressLogFileToTgz(src, dst string) (err error) {
	// 创建文件
	fw, err := os.Create(dst)
	if err != nil {
		return
	}
	defer fw.Close()

	// 将 tar 包使用 gzip 压缩，其实添加压缩功能很简单，
	// 只需要在 fw 和 tw 之前加上一层压缩就行了，和 Linux 的管道的感觉类似
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// 创建 Tar.Writer 结构
	tw := tar.NewWriter(gw)
	// 如果需要启用 gzip 将上面代码注释，换成下面的

	defer tw.Close()

	// 下面就该开始处理数据了，这里的思路就是递归处理目录及目录下的所有文件和目录
	// 这里可以自己写个递归来处理，不过 Golang 提供了 filepath.Walk 函数，可以很方便的做这个事情
	// 直接将这个函数的处理结果返回就行，需要传给它一个源文件或目录，它就可以自己去处理
	// 我们就只需要去实现我们自己的 打包逻辑即可，不需要再去路径相关的事情
	return filepath.Walk(src, func(fileName string, fi os.FileInfo, err error) error {
		// 因为这个闭包会返回个 error ，所以先要处理一下这个
		if err != nil {
			return err
		}

		// 这里就不需要我们自己再 os.Stat 了，它已经做好了，我们直接使用 fi 即可
		hdr, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		// 这里需要处理下 hdr 中的 Name，因为默认文件的名字是不带路径的，
		// 打包之后所有文件就会堆在一起，这样就破坏了原本的目录结果
		// 例如： 将原本 hdr.Name 的 syslog 替换程 log/syslog
		// 这个其实也很简单，回调函数的 fileName 字段给我们返回来的就是完整路径的 log/syslog
		// strings.TrimPrefix 将 fileName 的最左侧的 / 去掉，
		// 熟悉 Linux 的都知道为什么要去掉这个
		hdr.Name = strings.TrimPrefix(fileName, string(filepath.Separator))

		// 写入文件信息
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}

		// 判断下文件是否是标准文件，如果不是就不处理了，
		// 如： 目录，这里就只记录了文件信息，不会执行下面的 copy
		if !fi.Mode().IsRegular() {
			return nil
		}

		// 打开文件
		fr, err := os.Open(fileName)
		defer fr.Close()
		if err != nil {
			return err
		}

		// copy 文件数据到 tw
		n, err := io.Copy(tw, fr)
		if err != nil {
			return err
		}

		// 记录下过程，这个可以不记录，这个看需要，这样可以看到打包的过程
		XltLog.Write(fmt.Sprintf("成功打包 %s ，共写入了 %d 字节的数据\n", fileName, n))

		return nil
	})
}

// getCompressLogFileName 获取压缩日志文件名
func (l *Logger) getCompressLogFileName() string {

	// example: 2021-08-02-11-27
	prefix := l.oldFileList[0][:strings.Index(l.oldFileList[0], ".")]

	// example: 2021-08-02-11-33
	suffix := l.oldFileList[len(l.oldFileList)-1][:strings.Index(l.oldFileList[len(l.oldFileList)-1], ".")]

	// example: 2021-08-02-11-27~2021-08-02-11-33
	return prefix + "~" + suffix
}

// getRunningDirPath 获取项目运行路径
func (l *Logger) getRunningDirPath() (string, error) {
	runningDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return runningDir, nil
}