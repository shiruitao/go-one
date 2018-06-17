package filter

var MapFilter map[string]interface{}

func init() {
	MapFilter = make(map[string]interface{})
	MapFilter["/health/user/login"] = struct{}{}
	MapFilter["/graduation/student/get"] = struct{}{}
	MapFilter["/health/question/add "] = struct{}{}
	//MapFilter["/shop/ware/recommend"] = struct{}{}
	//MapFilter["/readTime"] = struct{}{}
}
