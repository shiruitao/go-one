package routers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/community/controllers"
)

func init() {
	// weather
	beego.Router("/community/weather", &controllers.UserController{}, "get:Weather")

	// notice
	beego.Router("/community/notice/create", &controllers.NoticeController{}, "post:Create")
	beego.Router("/community/notice/get", &controllers.NoticeController{}, "get:GetAll")

	// user
	beego.Router("/community/user/create", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/community/user/modify", &controllers.UserController{}, "post:UserUpdate")
	beego.Router("/community/user/getuser", &controllers.UserController{}, "get:GetUser")
	beego.Router("/community/user/getusername", &controllers.UserController{}, "post:GetUserByName")
	beego.Router("/community/user/getuserarea", &controllers.UserController{}, "post:GetUserByArea")
	beego.Router("/community/user/getuserage", &controllers.UserController{}, "post:GetUserByAge")

	// economic
	beego.Router("/community/economic/get", &controllers.EconomicController{}, "post:GetEconomic")
	beego.Router("/community/economic/add", &controllers.EconomicController{}, "post:CreateEconomic")

	// industry
	beego.Router("/community/industry/get", &controllers.IndustryController{}, "post:GetIndustry")
	beego.Router("/community/industry/add", &controllers.IndustryController{}, "post:CreateIndustry")

	// company
	beego.Router("/community/company/create", &controllers.CompanyController{}, "post:CreateCompany")
	beego.Router("/community/company/getall", &controllers.CompanyController{}, "get:GetCompanyAll")
	beego.Router("/community/company/getbyarea", &controllers.CompanyController{}, "post:GetCompanyByArea")
	beego.Router("/community/company/delete", &controllers.CompanyController{}, "post:DeleteCompany")

	// region
	beego.Router("/community/region/create", &controllers.RegionController{}, "post:CreateRegion")
	beego.Router("/community/region/getall", &controllers.RegionController{}, "get:GetAll")
	beego.Router("/community/region/getname", &controllers.RegionController{}, "post:GetRegion")
	beego.Router("/community/region/modify", &controllers.RegionController{}, "post:UpdateRegion")
	beego.Router("/community/region/delete", &controllers.RegionController{}, "post:DeleteRegion")

	// house
	beego.Router("/community/house/create", &controllers.HouseController{}, "post:CreateHouse")
	beego.Router("/community/house/getall", &controllers.HouseController{}, "get:GetAll")
	beego.Router("/community/house/getname", &controllers.HouseController{}, "post:GetHouseByName")
	beego.Router("/community/house/modify", &controllers.HouseController{}, "post:UpdateHouse")
	beego.Router("/community/house/delete", &controllers.HouseController{}, "post:DeleteHouse")
}
