package cmd

import (
	"log"
	"strings"

	"github.com/iforgot2021/golang/tour/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUStoUC
	ModeUStoLC
	ModeCCtoUS
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下",
	"1：全部单词转化为大写",
	"2：全部单词转换为小写",
	"3：下划线单词转换为大写驼峰",
	"4：下划线单词转换为小写驼峰单词",
	"5：驼峰单词转换为下划线单词",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeLower:
			content = word.ToUpper(str)
		case ModeUpper:
			content = word.ToLower(str)
		case ModeUStoUC:
			content = word.UStoUC(str)
		case ModeUStoLC:
			content = word.UStoLC(str)
		case ModeCCtoUS:
			content = word.CCtoUS(str)
		}
		log.Printf("转换结果：%s \n", content)
	},
}
var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
