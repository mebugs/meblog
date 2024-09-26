package config

import (
	"flag"
	"fmt"
	"net"
	"os"
	"runtime"
	"strconv"
	"strings"
)

const (
	CSysEnv  = "ENV"
	CSysNode = "NODE"

	Product = "prod"
	Staging = "stag"
	Testing = "test"
)

var (
	// Hostname 主机名
	Hostname string

	// SysEnv 系统运行环境，须配置环境变量SYS_ENV = testing/staging/product
	SysEnv string

	// 系统运行环境，应用启动时由命令行参数传入（运行环境优先使用该值）
	envFlag string

	// SysNode 系统节点，须配置环境变量SYS_NODE = app01/app02/app03
	SysNode string

	// SysNodeNo refer to the node number (node last 2 character)
	SysNodeNo string

	// LocalIP 本机 IP
	LocalIP string

	// WorkDir 程序启动目录
	WorkDir string

	// Pkg GOPATH包名，例如：siteOl.com/stone/server
	Pkg string
)

func SetEnv(pkg string) {
	Pkg = pkg
	sysInit()
}
func sysInit() {
	hostname()
	localIP()
	sysEnv()
	systemNode()
	workDir()
}

func sysEnv() {
	if len(os.Args) > 0 {
		for _, arg := range os.Args {
			if arg == "-env" { // 传递env参数时赋值
				flag.StringVar(&envFlag, "env", "", "sys environment")
				flag.Parse()
			}
		}
	}
	if envFlag != "" {
		SysEnv = envFlag
	} else {
		SysEnv = os.Getenv(CSysEnv)
		if SysEnv == "" {
			fmt.Printf("System environment variable `%s` not set, must set to `testing` or `staging` or `product`\n", CSysEnv)
			os.Exit(1)
		}
	}
	fmt.Printf("Environment:\t %s\n", SysEnv)
}

func systemNode() {
	SysNode = os.Getenv(CSysNode)
	if SysNode == "" {
		fmt.Printf("System node variable `%s` not set, must set to app01/app02/app03\n", CSysNode)
		os.Exit(1)
	}
	fmt.Printf("Sytem Node:\t %s\n", SysNode)
	nodeNo := SysNode[len(SysNode)-2:]
	if _, err := strconv.Atoi(nodeNo); err != nil {
		SysNodeNo = SysNode
	} else {
		SysNodeNo = nodeNo
	}

}

// 本机IP
func localIP() {
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err)
		LocalIP = "127.0.0.1"
	} else {
		LocalIP = strings.Split(conn.LocalAddr().String(), ":")[0]
		conn.Close()
	}
	fmt.Printf("Local IP:\t %s\n", LocalIP)
}

// 获取程序启动目录
func workDir() {
	var err error
	WorkDir, err = os.Getwd()
	if err != nil {
		fmt.Printf("Can not get work directory: %s\n", err)
		os.Exit(1)
	}
	//兼容 windows 文件分隔符 /
	if runtime.GOOS == "windows" {
		WorkDir = strings.Replace(WorkDir, "\\", "/", -1)
	}
	if Pkg != "" {
		if pos := strings.Index(WorkDir, Pkg); pos >= 0 {
			WorkDir = WorkDir[:(pos + len(Pkg))]
		}
	}
	fmt.Printf("Work Dir:\t %s\n", WorkDir)
}

// 取主机名，如果没取到，返回 `unknown`
func hostname() {
	h, err := os.Hostname()
	if err != nil {
		fmt.Errorf("get hostname error: %s", h)
		Hostname = "unknown"
	}
	Hostname = strings.Replace(h, ".", "_", -1)
	fmt.Printf("Hostname:\t %s\n", Hostname)
}
