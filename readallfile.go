package apiiget

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/// ReadAllFileAndGetComment 可读取所有的返回的json格式中的评论文本 若是出现错误请先查询文本是否符合json格式
func ReadAllFileAndGetComment(path string) {

	fileresults := Dfs(path)
	writefile, err := os.OpenFile("./finaltext/final.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Printf("打开写入文件出现错误:%s", err.Error())
		return
	}
	fmt.Printf("打开写入文件成功\n")
	defer writefile.Close()

	for i := range fileresults {
		var allStruct = AllStruct{}
		fmt.Printf("文件名:%s\n", fileresults[i])
		file, err := os.OpenFile(fileresults[i], os.O_RDONLY, os.ModePerm)

		if err != nil {
			fmt.Printf("打开文件失败:%s", err.Error())
			return
		}
		defer file.Close()

		fmt.Printf("打开文件:%s成功\n", fileresults[i])

		reader := bufio.NewReader(file)
		var jsonstring = ""
		for {
			line, _, err := reader.ReadLine()
			if err == io.EOF {
				break
			}
			jsonstring += string(line)
		}
		if jsonstring == "" {
			continue
		}

		if err := json.Unmarshal([]byte(jsonstring), &allStruct); err != nil {
			fmt.Printf("解析到json格式失败:%s", err.Error())
			return
		}
		for i := range allStruct.Responses.ReviewList {
			var content = allStruct.Responses.ReviewList[i].Reviewdetails.ReviewContent
			if _, err := writefile.WriteString(content); err != nil {
				fmt.Printf("文件写入失败:%s", err.Error())
				return
			}

			if _, err := writefile.WriteString("\n"); err != nil {
				fmt.Printf("文件写入失败:%s", err.Error())
				return
			}

		}
		fmt.Printf("文件:%s读取完毕。。。\n", fileresults[i])

	}
}

func Dfs(basepath string) []string {
	dirs, err := ioutil.ReadDir(basepath)
	if err != nil {
		fmt.Printf("读取文件夹出现错误:%s", err.Error())
	}
	var resultlist = make([]string, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			filepath := basepath + "/" + dir.Name()
			results := Dfs(filepath)

			resultlist = append(resultlist, results...)

		} else {
			resultlist = append(resultlist, basepath+"/"+dir.Name())
		}

	}
	return resultlist

}
