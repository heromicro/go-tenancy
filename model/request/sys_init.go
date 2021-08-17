package request

type InitDB struct {
	Sql       Sql    `json:"sql"`
	SqlType   string `json:"sqlType"`
	Cache     Cache  `json:"cache"`
	CacheType string `json:"cacheType"`
}

type Sql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"userName" binding:"required"`
	Password string `json:"password"`
	DBName   string `json:"dbName" binding:"required"`
}
type Cache struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}
