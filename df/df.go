//使用http状态码表示业务状态，注意避开已被浏览器使用的
package df

const (
	HTTP_STATUS_PARAM_ERROR  int = 601 //参数错误
	HTTP_STATUS_AUTH_ERROR   int = 602 //未登录
	HTTP_STATUS_SERVER_ERROR int = 603 //服务器错误，比如数据库执行失败等
)

var HTTP_STATUS_TEXT map[int]string = map[int]string{
	HTTP_STATUS_PARAM_ERROR: "参数错误",
}
