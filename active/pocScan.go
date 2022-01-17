package active

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
	"youzai/active/poc"
)

// 用于存取目标扫描的信息
type Target_Info struct {
	Target_url string        // 目标的url
	User_agent string        // 存放agent
	Vulns      []poc.PocInfo // 存放漏洞信息
}

var Target = &Target_Info{} // 实例化用于存储扫描结果的对象

// type PocInfoAll struct {
// 	pocInfo []struct{} // 存储poc所有信息
// 	pocName []string   // 存储poc名字
// }

// 此函数适用于安全性
// 此函数用于生成所有poc
// func PocInit() {
// 	methodName := []string{}                // 用于保存方法名
// 	pocStruct := &poc.PocInfo{}             // 实例化一个poc结构体，主要用于通过反射调用poc结构体内的方法
// 	pocReflect := reflect.TypeOf(pocStruct) // 用于获取方法的数量、方法名

// 	// 将方法名添加到数组中
// 	for i := 0; i < pocReflect.NumMethod(); i++ {
// 		method := pocReflect.Method(i) // 获取方法名
// 		methodName = append(methodName, method.Name)
// 	}

// 	// 调用方法，生成poc
// 	for _, pocName := range methodName {
// 		if fun, bl := pocReflect.MethodByName(pocName); bl {
// 			fun.Func.Call([]reflect.Value{reflect.ValueOf(pocStruct)}) // 调用方法生成poc
// 		}
// 	}
// }

// 此函数适用于快捷性
// 此函数用于生成所有poc
func PocInit() { // 用于保存方法名
	pocStruct := &poc.PocInfo{} // 实例化一个poc结构体，主要用于通过反射调用poc结构体内的方法
	pocReflect := reflect.ValueOf(pocStruct)

	for i := 0; i < pocReflect.NumMethod(); i++ {
		method := pocReflect.Method(i)
		method.Call(make([]reflect.Value, 0)) // 调用方法，生成poc
	}

	for _, pocStruct := range poc.PocStruct { // 将poc按照类型分类
		var pocType = pocStruct.Info.Type
		switch pocType {
		case "XSS":
			poc.PocMap["XSS"] = append(poc.PocMap["XSS"], pocStruct)

		case "SQLI":
			poc.PocMap["SQLI"] = append(poc.PocMap["SQLI"], pocStruct)

		case "RCE":
			poc.PocMap["RCE"] = append(poc.PocMap["RCE"], pocStruct)

		case "SSRF":
			poc.PocMap["SSRF"] = append(poc.PocMap["SSRF"], pocStruct)

		case "LFR":
			poc.PocMap["LFT"] = append(poc.PocMap["LFT"], pocStruct)

		case "UNAUTH":
			poc.PocMap["UNAUTH"] = append(poc.PocMap["UNAUTH"], pocStruct)

		case "INFO":
			poc.PocMap["INFO"] = append(poc.PocMap["INFO"], pocStruct)

		case "XXE":
			poc.PocMap["XXE"] = append(poc.PocMap["XXE"], pocStruct)

		default:
			poc.PocMap["OTHER"] = append(poc.PocMap["OTHER"], pocStruct)
		}
	}
}

// func Http_Request(method string, target_url string, timeout int, proxy bool, proxy_url string) { //第一个参数表示请求方法，第二个参数设置url，第三个参数设置超时
// 	// 设置请求属性
// 	transport := &http.Transport{
// 		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, // 取消证书认证
// 		ResponseHeaderTimeout: time.Second * time.Duration(timeout),  // 设置超时时间
// 	}

// 	// 检查是否使用代理
// 	if proxy != false && proxy_url != "" {
// 		urli := url.URL{}
// 		urlProxy, _ := urli.Parse(proxy_url)
// 		transport.Proxy = http.ProxyURL(urlProxy) // 设置代理
// 	}

// 	// 生成http客户端
// 	cli := &http.Client{
// 		Transport: transport,
// 	}
// }

func XSS_Check_Http(timeout int, proxy bool, proxy_url string) { // 第一个参数设置请求超时时间，第二个参数设置是否使用代理，第三个参数设置代理的url
	// 设置请求属性
	transport := &http.Transport{
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}, // 取消证书认证
		ResponseHeaderTimeout: time.Second * time.Duration(timeout),  // 设置超时时间
	}

	// 检查是否使用代理
	if proxy && proxy_url != "" {
		urli := url.URL{}
		urlProxy, _ := urli.Parse(proxy_url)
		transport.Proxy = http.ProxyURL(urlProxy) // 设置代理
	}

	// 生成http客户端
	cli := &http.Client{
		Transport: transport,
	}

	// 获取所有类型为XSS的poc
	if xss_poc_all, ok := poc.PocMap["XSS"]; ok {
		for _, xss_poc := range xss_poc_all {
			request := func(method_all []string, path_all []string, data_all []string, code_all []int, word_all []string, is_retry bool) {
				for i := range method_all {
					if method_all[i] == "GET" { // GET方法
						request, err := http.NewRequest(method_all[i], Target.Target_url+path_all[i], nil)
						if err != nil {
							return
						}
						request.Header.Add("User-Agent", Target.User_agent) // 设置User-Agent
						if len(xss_poc.Poc.Header) != 0 {                   // 获取poc中的header
							for header, value := range xss_poc.Poc.Header {
								request.Header.Add(header, value)
							}
						}
						if response, err := cli.Do(request); err != nil { // 发起http请求
							return
						} else {
							defer response.Body.Close()
							body, _ := ioutil.ReadAll(response.Body)
							if response.StatusCode == code_all[i] && strings.Contains(string(body), word_all[i]) {
								if is_retry {
									continue
								} else {
									Target.Vulns = append(Target.Vulns, xss_poc)
								}
							} else {
								continue
							}
						}
					} else if method_all[i] == "POST" { // POST方法
						request, err := http.NewRequest(method_all[i], Target.Target_url+path_all[i], strings.NewReader(data_all[i]))
						if err != nil {
							return
						}
						request.Header.Add("User-Agent", Target.User_agent) // 设置User-Agent
						if len(xss_poc.Poc.Header) != 0 {                   // 获取poc中的header
							for header, value := range xss_poc.Poc.Header {
								request.Header.Add(header, value)
							}
						}
						if response, err := cli.Do(request); err != nil { // 发起http请求
							return
						} else {
							defer response.Body.Close()
							body, _ := ioutil.ReadAll(response.Body)
							if response.StatusCode == code_all[i] && strings.Contains(string(body), word_all[i]) {
								if is_retry {
									continue
								} else {
									Target.Vulns = append(Target.Vulns, xss_poc)
								}
							} else {
								continue
							}
						}
					} else {
						return
					}
				}

			}

			// 判断是否需要多次请求
			if xss_poc.Config.Retry {
				request(xss_poc.Config.Retry_Method, xss_poc.Config.Retry_Path, xss_poc.Config.Retry_Data, xss_poc.Config.Retry_Code, xss_poc.Config.Retry_Word, true)
			}

			// 发起请求
			request(xss_poc.Poc.Method, xss_poc.Poc.Path, xss_poc.Poc.Data, xss_poc.Poc.Code, xss_poc.Poc.Word, false)

			// if xss_poc.Poc.Method == "GET" { // GET方法
			// 	for _, path := range xss_poc.Poc.Path {
			// 		request, err := http.NewRequest(xss_poc.Poc.Method, Target.Target_url+path, nil)
			// 		if err != nil {
			// 			return
			// 		}
			// 		request.Header.Add("User-Agent", Target.User_agent) // 设置User-Agent
			// 		if len(xss_poc.Poc.Header) != 0 {                   // 获取poc中的header
			// 			for header, value := range xss_poc.Poc.Header {
			// 				request.Header.Add(header, value)
			// 			}
			// 		}
			// 		if response, err := cli.Do(request); err != nil { // 发起http请求
			// 			return
			// 		} else {
			// 			defer response.Body.Close()
			// 			body, _ := ioutil.ReadAll(response.Body)
			// 			if strings.Contains(string(body), xss_poc.Poc.Word) {
			// 				Target.Vulns = append(Target.Vulns, xss_poc)
			// 			}
			// 		}
			// 	}
			// } else if xss_poc.Poc.Method == "POST" { // POST方法
			// 	for _, path := range xss_poc.Poc.Path {
			// 		request, err := http.NewRequest(xss_poc.Poc.Method, Target.Target_url+path, strings.NewReader(xss_poc.Poc.Data))
			// 		if err != nil {
			// 			return
			// 		}
			// 		request.Header.Add("User-Agent", Target.User_agent) // 设置User-Agent
			// 		if len(xss_poc.Poc.Header) != 0 {                   // 获取poc中的header
			// 			for header, value := range xss_poc.Poc.Header {
			// 				request.Header.Add(header, value)
			// 			}
			// 		}
			// 		if response, err := cli.Do(request); err != nil { // 发起http请求
			// 			return
			// 		} else {
			// 			defer response.Body.Close()
			// 			body, _ := ioutil.ReadAll(response.Body)
			// 			if strings.Contains(string(body), xss_poc.Poc.Word) {
			// 				Target.Vulns = append(Target.Vulns, xss_poc)
			// 			}
			// 		}
			// 	}
			// } else {
			// 	return
			// }
		}
	}
}
