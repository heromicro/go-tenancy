package model

import (
	"database/sql/driver"
	"strings"
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

type CUser struct {
	g.TENANCY_MODEL

	BaseGeneralInfo

	Username string `json:"userName" gorm:"not null;type:varchar(32);comment:用户登录名"`
	Password string `json:"-"  gorm:"not null;type:varchar(128);comment:用户登录密码"`
	Status   int    `gorm:"column:status;type:tinyint(1);not null;default:1" json:"status"`   // 账号冻结 1为正常，2为禁止
	IsShow   int    `gorm:"column:is_show;type:tinyint(1);not null;default:1" json:"is_show"` // 是否显示 1为正常，2为禁止

	Authority    SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	AuthorityId  string       `json:"authorityId" gorm:"not null;type:varchar(90)"`
	SysTenancyId uint         `json:"sysTenancyId" form:"sysTenancyId" gorm:"column:sys_tenancy_id;comment:关联标记"`
	GroupId      uint         `gorm:"column:group_id;type:int unsigned;not null;default:0" json:"groupId"` // 用户分组id
}

const (
	UnknownSex int = iota
	Male
	Female
)

const (
	Unknown         int = iota
	LoginTypeLite       // 小程序用户
	LoginTypeDevice     // 床旁设备用户
)

// BaseGeneralInfo
// - C端用户如果是小程序登录，根据UnionId判断用户唯一性
// - 如果是床旁设备登录，根据 SysTenancyId 商户医院id和 UserMerchant 表的 HospitalNO 住院号判断用户唯一性
// - 小程序端和用户端用户根据手机号码关联数据
type BaseGeneralInfo struct {
	Email     string   `json:"email" gorm:"default:'';comment:员工邮箱"`
	Phone     string   `json:"phone" gorm:"type:char(15);default:'';comment:手机号"`
	NickName  string   `json:"nickName" gorm:"type:varchar(16);comment:昵称"`
	AvatarUrl string   `json:"avatarUrl" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
	Sex       int      `json:"sex" form:"sex" gorm:"column:sex;comment:性别 1:男，2：女"`
	Age       int      `json:"age" form:"age" gorm:"column:age;comment:年龄"`
	Subscribe int      `json:"subscribe" form:"subscribe" gorm:"column:subscribe;comment:是否订阅"`
	OpenId    string   `json:"openId" form:"openId" gorm:"type:varchar(30);column:open_id;comment:openid"`
	UnionId   string   `json:"unionId" form:"unionId" gorm:"type:varchar(30);column:union_id;comment:unionId"`
	Country   string   `json:"country" form:"country" gorm:"type:varchar(32);column:country;comment:国家"`
	Province  string   `json:"province" form:"province" gorm:"type:varchar(32);column:province;comment:省份"`
	City      string   `json:"city" form:"city" gorm:"type:varchar(32);column:city;comment:城市"`
	IdCard    string   `json:"idCard" form:"idCard" gorm:"type:varchar(20);column:id_card;comment:身份证号"`
	IsAuth    int      `json:"isAuth" form:"isAuth" gorm:"column:is_auth;comment:是否实名认证"`
	RealName  string   `json:"realName" form:"realName" gorm:"type:varchar(64);column:real_name;comment:真实IP"`
	Birthday  Birthday `json:"birthday" form:"birthday" gorm:"column:birthday;comment:生日"`

	Mark      string    `gorm:"column:mark;type:varchar(255);not null;default:''" json:"mark"`                      // 用户备注
	Address   string    `gorm:"column:address;type:varchar(128)" json:"address"`                                    // 地址
	LastTime  time.Time `gorm:"column:last_time;type:timestamp" json:"lastTime"`                                    // 最后一次登录时间
	LastIP    string    `gorm:"column:last_ip;type:varchar(16);not null" json:"lastIp"`                             // 最后一次登录ip
	NowMoney  float64   `gorm:"column:now_money;type:decimal(8,2) unsigned;not null;default:0.00" json:"nowMoney"`  // 用户余额
	LoginType int       `gorm:"column:login_type;type:varchar(32);not null" json:"loginType"`                       // 用户登录类型
	PayCount  int       `gorm:"column:pay_count;type:int unsigned;not null;default:0" json:"payCount"`              // 用户购买次数
	PayPrice  float64   `gorm:"column:pay_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"payPrice"` // 用户消费金额
}

// 自定义时间格式
const timelayout = "2006-01-02"

type Birthday time.Time

func (dt *Birthday) UnmarshalJSON(data []byte) (err error) {
	value := strings.Trim(string(data), "\"")
	now, err := time.ParseInLocation(timelayout, value, time.Local)
	*dt = Birthday(now)
	return
}

func (dt Birthday) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timelayout)+2)
	b = append(b, '"')
	b = time.Time(dt).AppendFormat(b, timelayout)
	b = append(b, '"')
	return b, nil
}

func (dt Birthday) String() string {
	return time.Time(dt).Format(timelayout)
}

func (dt Birthday) Value() (driver.Value, error) {
	var zeroTime time.Time
	ti := time.Time(dt)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (dt *Birthday) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*dt = Birthday(value)
		return nil
	}
	return nil
}

// 设置生日
func SetBirthday() Birthday {
	var birthday Birthday
	birthday.UnmarshalJSON([]byte("1994-11-28"))
	return birthday
}
