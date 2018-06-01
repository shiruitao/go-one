package filter

var MapFilter map[string]interface{}

func init() {
	MapFilter = make(map[string]interface{})
	MapFilter["/bbs/user/register"] = struct{}{}
	MapFilter["/bbs/user/login"] = struct{}{}
	MapFilter["/bbs/art/get"] = struct{}{}
	MapFilter["/readTime"] = struct{}{}
}
