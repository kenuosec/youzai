package poc

func (Info *PocInfo) Git_INFO_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "web"
	poc.Info.Type = "INFO"
	poc.Info.Name = "Git INFO"
	poc.Info.Level = 3
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/.git/config"}
	poc.Poc.Param = []string{""}
	poc.Poc.Header = nil
	poc.Poc.Data = nil
	poc.Poc.Word = []string{"repositoryformatversion"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
