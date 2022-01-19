package report

import (
	"strings"
	"youzai/active"

	"github.com/gookit/color"
)

// 用于表格输出的函数，调用前需判断漏洞信息是否为空
func OutTable() {
	signs := []string{`+`, `-`, `|`} // 表格输出符号

	// 计算漏洞url最大的长度
	target_url_length := func() int {
		length := 0
		if len(active.Target.Target_Url) <= 10 {
			length = 10
		} else {
			length = len(active.Target.Target_Url) + 2
		}
		return length
	}()

	// 计算漏洞名最大长度
	target_vuln_name_length := func() int {
		length_all := []int{}
		var maxVal int
		for _, poc := range active.Target.Vulns {
			length_all = append(length_all, len(poc.Info.Name))
		}
		func(arr []int) {
			maxVal = arr[0]
			for i := 1; i < len(arr); i++ {
				if maxVal < arr[i] {
					maxVal = arr[i]
				}
			}
		}(length_all)
		return maxVal + 2
	}()

	target_vuln_id_length := func() int { // 计算漏id最大长度
		length_all := []int{}
		var maxVal int
		for _, poc := range active.Target.Vulns {
			length_all = append(length_all, len(poc.Info.ID))
		}
		func(arr []int) {
			maxVal = arr[0]
			for i := 1; i < len(arr); i++ {
				if maxVal < arr[i] {
					maxVal = arr[i]
				}
			}
		}(length_all)
		return maxVal + 2
	}()

	target_vuln_level_length := 14 // 漏洞等级最大长度

	// 打印边界线
	line := func() {
		color.Yellowln(signs[0], strings.Repeat(signs[1], target_url_length), signs[0], strings.Repeat(signs[1], target_vuln_name_length), signs[0], strings.Repeat(signs[1], target_vuln_id_length), signs[0], strings.Repeat(signs[1], target_vuln_level_length), signs[0])
	}

	line()

	color.Yellowln(signs[2], "Target_Url", strings.Repeat(" ", target_url_length-11), signs[2], "Vuln_Name", strings.Repeat(" ", target_vuln_name_length-10), signs[2], "Vuln_ID", strings.Repeat(" ", target_vuln_id_length-8), signs[2], "Vuln_Level", strings.Repeat(" ", target_vuln_level_length-11), signs[2])

	line()

	for _, poc := range active.Target.Vulns {
		// 打印漏洞信息
		func() {
			color.Yellowln(signs[2], active.Target.Target_Url, strings.Repeat(" ", target_url_length-(len(active.Target.Target_Url)+1)), signs[2], poc.Info.Name, strings.Repeat(" ", target_vuln_name_length-(len(poc.Info.Name)+1)), signs[2], poc.Info.ID, strings.Repeat(" ", target_vuln_id_length-(len(poc.Info.ID)+1)), signs[2], poc.Info.Level, strings.Repeat(" ", target_vuln_level_length-2), signs[2])
		}()
		line()
	}
}
