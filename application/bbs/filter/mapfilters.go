package filter

var MapFilter map[string]interface{}

func init() {
	MapFilter = make(map[string]interface{})
	MapFilter["/bbs/user/register"] = struct{}{}
	MapFilter["/shop/user/create"] = struct{}{}
	MapFilter["/shop/ware/create"] = struct{}{}
	MapFilter["/shop/ware/getall"] = struct{}{}
	MapFilter["/shop/ware/recommend"] = struct{}{}
	MapFilter["/readTime"] = struct{}{}
}