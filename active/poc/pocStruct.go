package poc

var PocStruct = []PocInfo{} // 存储所有poc信息的结构体，调用PocInit()函数后，所有的操作将使用该结构体数组中的poc

var PocMap = make(map[string][]PocInfo) // 按照漏洞类型，存储所有poc信息的结构体分类

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
		Method []string          `json:"method"` // 请求方法
		Path   []string          `json:"path"`   // 请求的路径
		Param  []string          `json:"param"`  // 请求的参数信息
		Header map[string]string `json:"header"` // http头部信息
		Data   []string          `json:"data"`   // 请求体数据，一般用于post请求
		Code   []int             `json:"code"`   // 请求后的状态码
		Word   []string          `json:"word"`   // 请求后需要匹配的字段
	} `json:"poc"`
	// poc配置
	Config struct {
		Retry        bool     `json:"retry"`        // 是否需要多次请求
		Retry_catch  bool     `json:"retry_catch"`  // 是否需要获取响应包中的数据
		Retry_Proto  string   `json:"retry_proto"`  // 请求的协议
		Retry_Method []string `json:"retry_method"` // 请求的方法
		Retry_Path   []string `json:"retry_path"`   // 请求的路径
		Retry_Data   []string `json:"retry_data"`   // 请求的数据
		Retry_Code   []int    `json:"retry_code"`   // 请求后的状态码
		Retry_Word   []string `json:"retry_word"`   // 请求后需要匹配的字段
	}
}
