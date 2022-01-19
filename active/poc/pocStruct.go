package poc

var PocCustomize = PocInfo{} // 存储自定义poc的Config信息

var PocStruct = []PocInfo{} // 存储所有poc信息的结构体，调用PocInit()函数后，所有的操作将使用该结构体数组中的poc

var PocMap = make(map[string][]PocInfo) // 按照漏洞类型，存储所有poc信息的结构体分类

// 模板poc扫描结构体
type PocInfo struct {
	// poc描述信息
	Info struct {
		ID     string `json:"id"`     // 漏洞的CVE、或CNVD编号
		Target string `json:"target"` // 漏洞指纹
		Type   string `json:"type"`   // 漏洞类型
		Name   string `json:"name"`   // 漏洞名称
		Level  int    `json:"level"`  // 漏洞等级
		Author string `json:"author"` // poc编写的作者
	} `json:"info"`
	// poc规则
	Poc struct {
		Proto  string            `json:"proto"`  // 请求的协议
		Method string            `json:"method"` // 请求方法
		Path   []string          `json:"path"`   // 请求的路径
		Param  []string          `json:"param"`  // 请求的参数信息
		Header map[string]string `json:"header"` // http头部信息
		Data   []string          `json:"data"`   // 请求体数据，一般用于post请求
		Code   []int             `json:"code"`   // 请求后的状态码
		Word   []string          `json:"word"`   // 请求后需要匹配的字段
	} `json:"poc"`
	// poc配置
	Config struct {
		Customize  bool        `json:"customize"`  // 是否是自定义poc
		Url        string      `json:"url"`        // 请求的url
		Timeout    int         `json:"timeout"`    // 请求的超时时间
		User_Agent string      `json:"user_agent"` // 请求的User_Agent信息
		Proxy      bool        `json:"proxy"`      // 是否使用代理
		Proxy_Url  string      `json:"proxy_url"`  // 代理的url
		Check      func() bool `json:"check"`      // poc函数
	}
}
