# 获取携程网站的评论数据

# 运行方式

切换到当前目录
在 getapi_test.go中的test函数中进行修改需要运行的内容
getapi函数为 根据 酒店id获取对应的返回结果并写入文件
ReadAllFileAndGetComment 函数则是读取返回结果文件并提取出评论数据
datastruct 为 返回的结果的json数据格式 用于解析数据
