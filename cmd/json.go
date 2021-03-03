package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/go-programming-tour-book/tour/internal/json2struct"

	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "json转换和处理",
	Long:  "json转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var jsonFilePath string

var json2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "json转换",
	Long:  "json转换",
	Run: func(cmd *cobra.Command, args []string) {
		if jsonFilePath != "" {
			var err error
			str, err = readStr(jsonFilePath)
			if err != nil {
				log.Fatalf("read file error: %v", err)
			}
		}

		parser, err := json2struct.NewParser(str)
		if err != nil {
			log.Fatalf("json2struct.NewParser err: %v", err)
		}
		content := parser.Json2Struct()
		log.Printf("输出结果: \n%s", content)
	},
}

func init() {
	jsonCmd.AddCommand(json2structCmd)
	json2structCmd.Flags().StringVarP(&str, "str", "s", "", "请输入json字符串")
	json2structCmd.Flags().StringVarP(&jsonFilePath, "file", "f", "", "请输入json文件路径")
}

func readStr(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
