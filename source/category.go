package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Category = new(category)

type category struct{}

var categories = []model.ProductCategory{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 173}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "品牌服饰", Path: "/", Sort: 2, Status: g.StatusTrue, Level: 0}, SysTenancyID: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 179}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "美容美发", Path: "/", Sort: 10, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 180}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "生鲜食品", Path: "/", Sort: 1, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 181}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "美容彩妆", Path: "/", Sort: 4, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 182}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "母婴专区", Path: "/", Sort: 2, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 183}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "食品饮料", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 184}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "数码家电", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 185}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "营养保健", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 186}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "精品服装", Path: "/", Sort: 20, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 187}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "鲜花预定", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 188}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "教育培训", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 162}, BaseProductCategory: model.BaseProductCategory{Pid: 186, CateName: "男士上衣", Path: "/186/", Sort: 9, Pic: "http://qmplusimg.henrongyi.top/head.png", Level: 1, Status: g.StatusTrue}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 165}, BaseProductCategory: model.BaseProductCategory{Pid: 171, CateName: "祛斑祛痘", Path: "/181/171/", Sort: 0, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 171}, BaseProductCategory: model.BaseProductCategory{Pid: 181, CateName: "普通化妆", Path: "/181/", Sort: 0, Status: g.StatusTrue, Level: 1}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 172}, BaseProductCategory: model.BaseProductCategory{Pid: 162, CateName: "T恤/POLO", Path: "/186/162/", Sort: 9, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 174}, BaseProductCategory: model.BaseProductCategory{Pid: 173, CateName: "时尚女装", Path: "/173/", Sort: 0, Status: g.StatusTrue, Level: 1}, SysTenancyID: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 175}, BaseProductCategory: model.BaseProductCategory{Pid: 162, CateName: "商务衬衫", Path: "/186/162/", Sort: 6, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 176}, BaseProductCategory: model.BaseProductCategory{Pid: 162, CateName: "休闲衬衫", Path: "/186/162/", Sort: 6, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 177}, BaseProductCategory: model.BaseProductCategory{Pid: 171, CateName: "补水保湿", Path: "	/181/171/", Sort: 0, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
}

func (m *category) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.ProductCategory{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> categories 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&categories).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> categories 表初始数据成功!")
		return nil
	})
}
