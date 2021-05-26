package apiiget

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// IsExists 判断所给路径文件/文件夹是否存在
func IsExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil && !os.IsExist(err) {
		return false
	}
	return true
}

// GetApi 根据酒店的id获取酒店相关的评论 目前设置为 30页 每页 50条数据，经测试50条为最大的每页能获取的数据
// filename 为存储获取到的评论数据的位置 可自行修改
func GetApi(hotelId string) {

	for i := 1; i < 30; i++ {
		var pageno = strconv.Itoa(i)
		var filename = "./hotelcomment/" + hotelId + "page" + pageno + ".json"

		// 新建或打开文件
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		if err != nil {
			fmt.Printf("文件打开失败:%s\n", err.Error())
		}
		defer file.Close()
		var postform = "{\"PageNo\":" + pageno + ",\"PageSize\":50,\"MasterHotelId\":" + hotelId + ",\"NeedFilter\":true,\"UnUsefulPageNo\":1,\"UnUsefulPageSize\":5,\"isHasFold\":false,\"head\":{\"Locale\":\"zh-CN\",\"Currency\":\"CNY\",\"Device\":\"PC\",\"UserIP\":\"117.136.81.150\",\"Group\":\"\",\"ReferenceID\":\"\",\"UserRegion\":\"CN\",\"AID\":\"4897\",\"SID\":\"155952\",\"Ticket\":\"\",\"UID\":\"\",\"IsQuickBooking\":\"\",\"ClientID\":\"1618626143655.3ono9g\",\"OUID\":\"index\",\"TimeZone\":\"8\",\"P\":\"58056018605\",\"PageID\":\"102003\",\"Version\":\"\",\"HotelExtension\":{\"WebpSupport\":true,\"group\":\"CTRIP\",\"Qid\":\"623738380460\",\"hasAidInUrl\":false},\"Frontend\":{\"vid\":\"1618626143655.3ono9g\",\"sessionID\":15,\"pvid\":231}},\"ServerData\":\"\"}"

		response, err := http.Post("https://m.ctrip.com/restapi/soa2/16709/json/GetReviewList", "application/json", strings.NewReader(postform))
		if err != nil {
			fmt.Printf("错误信息为:%s", err.Error())
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		if len(body) < 1000 {
			fmt.Println("当前文件已无更多评论，退出循环...")
			break
		}
		count, err := file.WriteString(string(body))
		if err != nil {
			fmt.Printf("文件写入失败：%s\n", err.Error())
			return
		}
		fmt.Printf("文件写入成功，共:%v个字节！\n", count)
	}
	fmt.Printf("酒店id为:%s的酒店评论数据读取完成\n", hotelId)

}
