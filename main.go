package main

import (
	"youzai/active"
	"youzai/report"

	"github.com/gookit/color"
)

func banner() {
	banner := []string{`
	  ▓██   ██▓ ▒█████   █    ██ ▒███████▒ ▄▄▄       ██▓
	   ▒██  ██▒▒██▒  ██▒ ██  ▓██▒▒ ▒ ▒ ▄▀░▒████▄    ▓██▒
	    ▒██ ██░▒██░  ██▒▓██  ▒██░░ ▒ ▄▀▒░ ▒██  ▀█▄  ▒██▒
	    ░ ▐██▓░▒██   ██░▓▓█  ░██░  ▄▀▒   ░░██▄▄▄▄██ ░██░
	    ░ ██▒▓░░ ████▓▒░▒▒█████▓ ▒███████▒ ▓█   ▓██▒░██░
	     ██▒▒▒ ░ ▒░▒░▒░ ░▒▓▒ ▒ ▒ ░▒▒ ▓░▒░▒ ▒▒   ▓▒█░░▓  
	   ▓██ ░▒░   ░ ▒ ▒░ ░░▒░ ░ ░ ░░▒ ▒ ░ ▒  ▒   ▒▒ ░ ▒ ░
	   ▒ ▒ ░░  ░ ░ ░ ▒   ░░░ ░ ░ ░ ░ ░ ░ ░  ░   ▒    ▒ ░
	   ░ ░         ░ ░     ░       ░ ░          ░  ░ ░  
	   ░ ░                       ░                      
	`,
		`
	  ██╗   ██╗ ██████╗ ██╗   ██╗███████╗ █████╗ ██╗
	  ╚██╗ ██╔╝██╔═══██╗██║   ██║╚══███╔╝██╔══██╗██║
	   ╚████╔╝ ██║   ██║██║   ██║  ███╔╝ ███████║██║
	    ╚██╔╝  ██║   ██║██║   ██║ ███╔╝  ██╔══██║██║
	     ██║   ╚██████╔╝╚██████╔╝███████╗██║  ██║██║
	     ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝	                                                
	`}
	color.Blueln(banner[0])
	color.Greenp("Version：v 1.0")
	color.Cyanln("\t\t\t\t\t\t", "By youzai")
	color.Redln("Scan Results Info")
}

// 生成目标信息
func target_Info() {
	url := "http://localhost/wordpress"
	userAgent := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1"
	active.Target.Target_url = url
	active.Target.User_agent = userAgent
}

// 执行扫描
func active_Check() {
	active.PocInit()
	active.XSS_Check_Http(1, true, "http://127.0.0.1:8888")
}

// 扫描器入口
func main() {
	banner()
	target_Info()
	active_Check()
	report.OutTable()
}
