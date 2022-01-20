package poc

func (Info *PocInfo) WordPress_Example_JQuery_XSS_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "wordpress"
	poc.Info.Type = "XSS"
	poc.Info.Name = "WordPress example.html jQuery 1.7.2 reflect XSS"
	poc.Info.Level = 3
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/wp-content/themes/twentyfifteen/genericons/example.html"}
	poc.Poc.Param = []string{""}
	poc.Poc.Header = func() map[string]string {
		mapTemp := make(map[string]string)
		mapTemp["Content-Type"] = "application/x-www-form-urlencoded"
		return mapTemp
	}()
	poc.Poc.Data = []string{""}
	poc.Poc.Word = []string{"jquery/1.7.2/jquery.min.js"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
