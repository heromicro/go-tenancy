package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Express = new(express)

type express struct{}

var expresses = []model.Express{
	{Code: "LIMINWL", Name: "利民物流", Sort: 1, Status: 2},
	{Code: "XINTIAN", Name: "鑫天顺物流", Sort: 1, Status: 2},
	{Code: "henglu", Name: "恒路物流", Sort: 1, Status: 2},
	{Code: "klwl", Name: "康力物流", Sort: 1, Status: 2},
	{Code: "meiguo", Name: "美国快递", Sort: 1, Status: 2},
	{Code: "a2u", Name: "A2U速递", Sort: 1, Status: 2},
	{Code: "benteng", Name: "奔腾物流", Sort: 1, Status: 2},
	{Code: "ahdf", Name: "德方物流", Sort: 1, Status: 2},
	{Code: "timedg", Name: "万家通", Sort: 1, Status: 2},
	{Code: "ztong", Name: "智通物流", Sort: 1, Status: 2},
	{Code: "xindan", Name: "新蛋物流", Sort: 1, Status: 2},
	{Code: "bgpyghx", Name: "挂号信", Sort: 1, Status: 2},
	{Code: "XFHONG", Name: "鑫飞鸿物流快递", Sort: 1, Status: 2},
	{Code: "ALP", Name: "阿里物流", Sort: 1, Status: 2},
	{Code: "BFWL", Name: "滨发物流", Sort: 1, Status: 2},
	{Code: "SJWL", Name: "宋军物流", Sort: 1, Status: 2},
	{Code: "SHUNFAWL", Name: "顺发物流", Sort: 1, Status: 2},
	{Code: "TIANHEWL", Name: "天河物流", Sort: 1, Status: 2},
	{Code: "YBWL", Name: "邮联物流", Sort: 1, Status: 2},
	{Code: "SWHY", Name: "盛旺货运", Sort: 1, Status: 2},
	{Code: "TSWL", Name: "汤氏物流", Sort: 1, Status: 2},
	{Code: "YUANYUANWL", Name: "圆圆物流", Sort: 1, Status: 2},
	{Code: "BALIANGWL", Name: "八梁物流", Sort: 1, Status: 2},
	{Code: "ZGWL", Name: "振刚物流", Sort: 1, Status: 2},
	{Code: "JIAYU", Name: "佳宇物流", Sort: 1, Status: 2},
	{Code: "SHHX", Name: "昊昕物流", Sort: 1, Status: 2},
	{Code: "ande", Name: "安得物流", Sort: 1, Status: 2},
	{Code: "ppbyb", Name: "贝邮宝", Sort: 1, Status: 2},
	{Code: "dida", Name: "递达快递", Sort: 1, Status: 2},
	{Code: "jppost", Name: "日本邮政", Sort: 1, Status: 2},
	{Code: "intmail", Name: "中国邮政", Sort: 96, Status: 2},
	{Code: "HENGCHENGWL", Name: "恒诚物流", Sort: 1, Status: 2},
	{Code: "HENGFENGWL", Name: "恒丰物流", Sort: 1, Status: 2},
	{Code: "gdems", Name: "广东ems快递", Sort: 1, Status: 2},
	{Code: "xlyt", Name: "祥龙运通", Sort: 1, Status: 2},
	{Code: "gjbg", Name: "国际包裹", Sort: 1, Status: 2},
	{Code: "uex", Name: "UEX", Sort: 1, Status: 2},
	{Code: "singpost", Name: "新加坡邮政", Sort: 1, Status: 2},
	{Code: "guangdongyouzhengwuliu", Name: "广东邮政", Sort: 1, Status: 2},
	{Code: "bht", Name: "BHT", Sort: 1, Status: 2},
	// 41	cces	CCES快递	,Sort:	1,Status: 2 },
	// 42	cloudexpress	CE易欧通国际速递	,Sort:	1,Status: 2 },
	// 43	dasu	达速物流	,Sort:	1,Status: 2 },
	// 44	pfcexpress	皇家物流	,Sort:	1,Status: 2 },
	// 45	hjs	猴急送	,Sort:	1,Status: 2 },
	// 46	huilian	辉联物流	,Sort:	1,Status: 2 },
	// 47	huanqiu	环球速运	,Sort:	1,Status: 2 },
	// 48	huada	华达快运	,Sort:	1,Status: 2 },
	// 49	htwd	华通务达物流	,Sort:	1,Status: 2 },
	// 50	hipito	海派通	,Sort:	1,Status: 2 },
	// 51	hqtd	环球通达	,Sort:	1,Status: 2 },
	// 52	airgtc	航空快递	,Sort:	1,Status: 2 },
	// 53	haoyoukuai	好又快物流	,Sort:	1,Status: 2 },
	// 54	hanrun	韩润物流	,Sort:	1,Status: 2 },
	// 55	ccd	河南次晨达	,Sort:	1,Status: 2 },
	// 56	hfwuxi	和丰同城	,Sort:	1,Status: 2 },
	// 57	Sky	荷兰	,Sort:	1,Status: 2 },
	// 58	hongxun	鸿讯物流	,Sort:	1,Status: 2 },
	// 59	hongjie	宏捷国际物流	,Sort:	1,Status: 2 },
	// 60	httx56	汇通天下物流	,Sort:	1,Status: 2 },
	// 61	lqht	恒通快递	,Sort:	1,Status: 2 },
	// 62	jinguangsudikuaijian	京广速递快件	,Sort:	1,Status: 2 },
	// 63	junfengguoji	骏丰国际速递	,Sort:	1,Status: 2 },
	// 64	jiajiatong56	佳家通	,Sort:	1,Status: 2 },
	// 65	jrypex	吉日优派	,Sort:	1,Status: 2 },
	// 66	jinchengwuliu	锦程国际物流	,Sort:	1,Status: 2 },
	// 67	jgwl	景光物流	,Sort:	1,Status: 2 },
	// 68	pzhjst	急顺通	,Sort:	1,Status: 2 },
	// 69	ruexp	捷网俄全通	,Sort:	1,Status: 2 },
	// 70	jmjss	金马甲	,Sort:	1,Status: 2 },
	// 71	lanhu	蓝弧快递	,Sort:	1,Status: 2 },
	// 72	ltexp	乐天速递	,Sort:	1,Status: 2 },
	// 73	lutong	鲁通快运	,Sort:	1,Status: 2 },
	// 74	ledii	乐递供应链	,Sort:	1,Status: 2 },
	// 75	lundao	论道国际物流	,Sort:	1,Status: 2 },
	// 76	mailikuaidi	麦力快递	,Sort:	1,Status: 2 },
	// 77	mchy	木春货运	,Sort:	1,Status: 2 },
	// 78	meiquick	美快国际物流	,Sort:	1,Status: 2 },
	// 79	valueway	美通快递	,Sort:	1,Status: 2 },
	// 80	nuoyaao	偌亚奥国际	,Sort:	1,Status: 2 },
	// 81	euasia	欧亚专线	,Sort:	1,Status: 2 },
	// 82	pca	澳大利亚PCA快递	,Sort:	1,Status: 2 },
	// 83	pingandatengfei	平安达腾飞	,Sort:	1,Status: 2 },
	// 84	pjbest	品骏快递	,Sort:	1,Status: 2 },
	// 85	qbexpress	秦邦快运	,Sort:	1,Status: 2 },
	// 86	quanxintong	全信通快递	,Sort:	1,Status: 2 },
	// 87	quansutong	全速通国际快递	,Sort:	1,Status: 2 },
	// 88	qinyuan	秦远物流	,Sort:	1,Status: 2 },
	// 89	qichen	启辰国际物流	,Sort:	1,Status: 2 },
	// 90	quansu	全速快运	,Sort:	1,Status: 2 },
	// 91	qzx56	全之鑫物流	,Sort:	1,Status: 2 },
	// 92	qskdyxgs	千顺快递	,Sort:	1,Status: 2 },
	// 93	runhengfeng	全时速运	,Sort:	1,Status: 2 },
	// 94	rytsd	日益通速递	,Sort:	1,Status: 2 },
	// 95	ruidaex	瑞达国际速递	,Sort:	1,Status: 2 },
	// 96	shiyun	世运快递	,Sort:	1,Status: 2 },
	// 97	sfift	十方通物流	,Sort:	1,Status: 2 },
	// 98	stkd	顺通快递	,Sort:	1,Status: 2 },
	// 99	bgn	布谷鸟快递	,Sort:	1,Status: 2 },
	// 100	jiahuier	佳惠尔快递	,Sort:	1,Status: 2 },
	// 101	pingyou	小包	,Sort:	1,Status: 2 },
	// 102	yumeijie	誉美捷快递	,Sort:	1,Status: 2 },
	// 103	meilong	美龙快递	,Sort:	1,Status: 2 },
	// 104	guangtong	广通速递	,Sort:	1,Status: 2 },
	// 105	STARS	星晨急便	,Sort:	1,Status: 2 },
	// 106	NANHANG	中国南方航空股份有限公司	,Sort:	1,Status: 2 },
	// 107	lanbiao	蓝镖快递	,Sort:	1,Status: 2 },
	// 109	baotongda	宝通达物流	,Sort:	1,Status: 2 },
	// 110	dashun	大顺物流	,Sort:	1,Status: 2 },
	// 111	dada	大达物流	,Sort:	1,Status: 2 },
	// 112	fangfangda	方方达物流	,Sort:	1,Status: 2 },
	// 113	hebeijianhua	河北建华物流	,Sort:	1,Status: 2 },
	// 114	haolaiyun	好来运快递	,Sort:	1,Status: 2 },
	// 115	jinyue	晋越快递	,Sort:	1,Status: 2 },
	// 116	kuaitao	快淘快递	,Sort:	1,Status: 2 },
	// 117	peixing	陪行物流	,Sort:	1,Status: 2 },
	// 118	hkpost	香港邮政	,Sort:	1,Status: 2 },
	// 119	ytfh	一统飞鸿快递	,Sort:	1,Status: 2 },
	// 120	zhongxinda	中信达快递	,Sort:	1,Status: 2 },
	// 121	zhongtian	中天快运	,Sort:	1,Status: 2 },
	// 122	zuochuan	佐川急便	,Sort:	1,Status: 2 },
	// 123	chengguang	程光快递	,Sort:	1,Status: 2 },
	// 124	cszx	城市之星	,Sort:	1,Status: 2 },
	// 125	chuanzhi	传志快递	,Sort:	1,Status: 2 },
	// 126	feibao	飞豹快递	,Sort:	1,Status: 2 },
	// 127	huiqiang	汇强快递	,Sort:	1,Status: 2 },
	// 128	lejiedi	乐捷递	,Sort:	1,Status: 2 },
	// 129	lijisong	成都立即送快递	,Sort:	1,Status: 2 },
	// 130	minbang	民邦速递	,Sort:	1,Status: 2 },
	// 131	ocs	OCS国际快递	,Sort:	1,Status: 2 },
	// 132	santai	三态速递	,Sort:	1,Status: 2 },
	// 133	saiaodi	赛澳递	,Sort:	1,Status: 2 },
	// 134	jd	京东快递	,Sort:	1,Status: 2 },
	// 135	zengyi	增益快递	,Sort:	1,Status: 2 },
	// 136	fanyu	凡宇速递	,Sort:	1,Status: 2 },
	// 137	fengda	丰达快递	,Sort:	1,Status: 2 },
	// 138	coe	东方快递	,Sort:	1,Status: 2 },
	// 139	ees	百福东方快递	,Sort:	1,Status: 2 },
	// 140	disifang	递四方速递	,Sort:	1,Status: 2 },
	// 141	rufeng	如风达快递	,Sort:	1,Status: 2 },
	// 142	changtong	长通物流	,Sort:	1,Status: 2 },
	// 143	chengshi100	城市100快递	,Sort:	1,Status: 2 },
	// 144	feibang	飞邦物流	,Sort:	1,Status: 2 },
	// 145	haosheng	昊盛物流	,Sort:	1,Status: 2 },
	// 146	yinsu	音速速运	,Sort:	1,Status: 2 },
	// 147	kuanrong	宽容物流	,Sort:	1,Status: 2 },
	// 148	tongcheng	通成物流	,Sort:	1,Status: 2 },
	// 149	tonghe	通和天下物流	,Sort:	1,Status: 2 },
	// 150	zhima	芝麻开门	,Sort:	1,Status: 2 },
	// 151	ririshun	日日顺物流	,Sort:	1,Status: 2 },
	// 152	anxun	安迅物流	,Sort:	1,Status: 2 },
	// 153	baiqian	百千诚国际物流	,Sort:	1,Status: 2 },
	// 154	chukouyi	出口易	,Sort:	1,Status: 2 },
	// 155	diantong	店通快递	,Sort:	1,Status: 2 },
	// 156	dajin	大金物流	,Sort:	1,Status: 2 },
	// 157	feite	飞特物流	,Sort:	1,Status: 2 },
	// 159	gnxb	国内小包	,Sort:	1,Status: 2 },
	// 160	huacheng	华诚物流	,Sort:	1,Status: 2 },
	// 161	huahan	华翰物流	,Sort:	1,Status: 2 },
	// 162	hengyu	恒宇运通	,Sort:	1,Status: 2 },
	// 163	huahang	华航快递	,Sort:	1,Status: 2 },
	// 164	jiuyi	久易快递	,Sort:	1,Status: 2 },
	// 165	jiete	捷特快递	,Sort:	1,Status: 2 },
	// 166	jingshi	京世物流	,Sort:	1,Status: 2 },
	// 167	kuayue	跨越快递	,Sort:	1,Status: 2 },
	// 168	mengsu	蒙速快递	,Sort:	1,Status: 2 },
	// 169	nanbei	南北快递	,Sort:	1,Status: 2 },
	// 171	pinganda	平安达快递	,Sort:	1,Status: 2 },
	// 172	ruifeng	瑞丰速递	,Sort:	1,Status: 2 },
	// 173	rongqing	荣庆物流	,Sort:	1,Status: 2 },
	// 174	suijia	穗佳物流	,Sort:	1,Status: 2 },
	// 175	simai	思迈快递	,Sort:	1,Status: 2 },
	// 176	suteng	速腾快递	,Sort:	1,Status: 2 },
	// 177	shengbang	晟邦物流	,Sort:	1,Status: 2 },
	// 178	suchengzhaipei	速呈宅配	,Sort:	1,Status: 2 },
	// 179	wuhuan	五环速递	,Sort:	1,Status: 2 },
	// 180	xingchengzhaipei	星程宅配	,Sort:	1,Status: 2 },
	// 181	yinjie	顺捷丰达	,Sort:	1,Status: 2 },
	// 183	yanwen	燕文物流	,Sort:	1,Status: 2 },
	// 184	zongxing	纵行物流	,Sort:	1,Status: 2 },
	// 185	aae	AAE快递	,Sort:	1,Status: 2 },
	// 186	dhl	DHL快递	,Sort:	1,Status: 2 },
	// 187	feihu	飞狐快递	,Sort:	1,Status: 2 },
	// 188	shunfeng	顺丰速运	92	1
	// 189	spring	春风物流	,Sort:	1,Status: 2 },
	// 190	yidatong	易达通快递	,Sort:	1,Status: 2 },
	// 191	PEWKEE	彪记快递	,Sort:	1,Status: 2 },
	// 192	PHOENIXEXP	凤凰快递	,Sort:	1,Status: 2 },
	// 193	CNGLS	GLS快递	,Sort:	1,Status: 2 },
	// 194	BHTEXP	华慧快递	,Sort:	1,Status: 2 },
	// 195	B2B	卡行天下	,Sort:	1,Status: 2 },
	// 196	PEISI	配思货运	,Sort:	1,Status: 2 },
	// 197	SUNDAPOST	上大物流	,Sort:	1,Status: 2 },
	// 198	SUYUE	苏粤货运	,Sort:	1,Status: 2 },
	// 199	F5XM	伍圆速递	,Sort:	1,Status: 2 },
	// 200	GZWENJIE	文捷航空速递	,Sort:	1,Status: 2 },
	// 201	yuancheng	远成物流	,Sort:	1,Status: 2 },
	// 202	dpex	DPEX快递	,Sort:	1,Status: 2 },
	// 203	anjie	安捷快递	,Sort:	1,Status: 2 },
	// 204	jldt	嘉里大通	,Sort:	1,Status: 2 },
	// 205	yousu	优速快递	,Sort:	1,Status: 2 },
	// 206	wanbo	万博快递	,Sort:	1,Status: 2 },
	// 207	sure	速尔物流	,Sort:	1,Status: 2 },
	// 208	sutong	速通物流	,Sort:	1,Status: 2 },
	// 209	JUNCHUANWL	骏川物流	,Sort:	1,Status: 2 },
	// 210	guada	冠达快递	,Sort:	1,Status: 2 },
	// 211	dsu	D速快递	,Sort:	1,Status: 2 },
	// 212	LONGSHENWL	龙胜物流	,Sort:	1,Status: 2 },
	// 213	abc	爱彼西快递	,Sort:	1,Status: 2 },
	// 214	eyoubao	E邮宝	,Sort:	1,Status: 2 },
	// 215	aol	AOL快递	,Sort:	1,Status: 2 },
	// 216	jixianda	急先达物流	,Sort:	1,Status: 2 },
	// 217	haihong	山东海红快递	,Sort:	1,Status: 2 },
	// 218	feiyang	飞洋快递	,Sort:	1,Status: 2 },
	// 219	rpx	RPX保时达	,Sort:	1,Status: 2 },
	// 220	zhaijisong	宅急送	,Sort:	1,Status: 2 },
	// 221	tiantian	天天快递	99	0
	// 222	yunwuliu	云物流	,Sort:	1,Status: 2 },
	// 223	jiuye	九曳供应链	,Sort:	1,Status: 2 },
	// 224	bsky	百世快运	,Sort:	1,Status: 2 },
	// 225	higo	黑狗物流	,Sort:	1,Status: 2 },
	// 226	arke	方舟速递	,Sort:	1,Status: 2 },
	// 227	zwsy	中外速运	,Sort:	1,Status: 2 },
	// 228	jxy	吉祥邮	,Sort:	1,Status: 2 },
	// 229	aramex	Aramex	,Sort:	1,Status: 2 },
	// 230	guotong	国通快递	,Sort:	1,Status: 2 },
	// 231	jiayi	佳怡物流	,Sort:	1,Status: 2 },
	// 232	longbang	龙邦快运	,Sort:	1,Status: 2 },
	// 233	minhang	民航快递	,Sort:	1,Status: 2 },
	// 234	quanyi	全一快递	,Sort:	1,Status: 2 },
	// 235	quanchen	全晨快递	,Sort:	1,Status: 2 },
	// 236	usps	USPS快递	,Sort:	1,Status: 2 },
	// 237	xinbang	新邦物流	,Sort:	1,Status: 2 },
	// 238	yuanzhi	元智捷诚快递	,Sort:	1,Status: 2 },
	// 239	zhongyou	中邮物流	,Sort:	1,Status: 2 },
	// 240	yuxin	宇鑫物流	,Sort:	1,Status: 2 },
	// 241	cnpex	中环快递	,Sort:	1,Status: 2 },
	// 242	shengfeng	盛丰物流	,Sort:	1,Status: 2 },
	// 243	yuantong	圆通速递	97	1
	// 244	jiayunmei	加运美物流	,Sort:	1,Status: 2 },
	// 245	ywfex	源伟丰快递	,Sort:	1,Status: 2 },
	// 246	xinfeng	信丰物流	,Sort:	1,Status: 2 },
	// 247	wanxiang	万象物流	,Sort:	1,Status: 2 },
	// 248	menduimen	门对门	,Sort:	1,Status: 2 },
	// 249	mingliang	明亮物流	,Sort:	1,Status: 2 },
	// 250	fengxingtianxia	风行天下	,Sort:	1,Status: 2 },
	// 251	gongsuda	共速达物流	,Sort:	1,Status: 2 },
	// 252	zhongtong	中通快递	100	1
	// 253	quanritong	全日通快递	,Sort:	1,Status: 2 },
	// 254	ems	EMS	1	1
	// 255	wanjia	万家物流	,Sort:	1,Status: 2 },
	// 256	yuntong	运通快递	,Sort:	1,Status: 2 },
	// 257	feikuaida	飞快达物流	,Sort:	1,Status: 2 },
	// 258	haimeng	海盟速递	,Sort:	1,Status: 2 },
	// 259	zhongsukuaidi	中速快件	,Sort:	1,Status: 2 },
	// 260	yuefeng	越丰快递	,Sort:	1,Status: 2 },
	// 261	shenghui	盛辉物流	,Sort:	1,Status: 2 },
	// 262	datian	大田物流	,Sort:	1,Status: 2 },
	// 263	quanjitong	全际通快递	,Sort:	1,Status: 2 },
	// 264	longlangkuaidi	隆浪快递	,Sort:	1,Status: 2 },
	// 265	neweggozzo	新蛋奥硕物流	,Sort:	1,Status: 2 },
	// 266	shentong	申通快递	95	1
	// 267	haiwaihuanqiu	海外环球	,Sort:	1,Status: 2 },
	// 268	yad	源安达快递	,Sort:	1,Status: 2 },
	// 269	jindawuliu	金大物流	,Sort:	1,Status: 2 },
	// 270	sevendays	七天连锁	,Sort:	1,Status: 2 },
	// 271	tnt	TNT快递	,Sort:	1,Status: 2 },
	// 272	huayu	天地华宇物流	,Sort:	1,Status: 2 },
	// 273	lianhaotong	联昊通快递	,Sort:	1,Status: 2 },
	// 274	nengda	港中能达快递	,Sort:	1,Status: 2 },
	// 275	LBWL	联邦物流	,Sort:	1,Status: 2 },
	// 276	ontrac	onTrac	,Sort:	1,Status: 2 },
	// 277	feihang	原飞航快递	,Sort:	1,Status: 2 },
	// 278	bangsongwuliu	邦送物流	,Sort:	1,Status: 2 },
	// 279	huaxialong	华夏龙物流	,Sort:	1,Status: 2 },
	// 280	ztwy	中天万运快递	,Sort:	1,Status: 2 },
	// 281	fkd	飞康达物流	,Sort:	1,Status: 2 },
	// 282	anxinda	安信达快递	,Sort:	1,Status: 2 },
	// 283	quanfeng	全峰快递	,Sort:	1,Status: 2 },
	// 284	shengan	圣安物流	,Sort:	1,Status: 2 },
	// 285	jiaji	佳吉物流	,Sort:	1,Status: 2 },
	// 286	yunda	韵达快运	94	0
	// 287	ups	UPS快递	,Sort:	1,Status: 2 },
	// 288	debang	德邦物流	,Sort:	1,Status: 2 },
	// 289	yafeng	亚风速递	,Sort:	1,Status: 2 },
	// 290	kuaijie	快捷速递	98	0
	// 291	huitong	百世快递	93	0
	// 293	aolau	AOL澳通速递	,Sort:	1,Status: 2 },
	// 294	anneng	安能物流	,Sort:	1,Status: 2 },
	// 295	auexpress	澳邮中国快运	,Sort:	1,Status: 2 },
	// 296	exfresh	安鲜达	,Sort:	1,Status: 2 },
	// 297	bcwelt	BCWELT	,Sort:	1,Status: 2 },
	// 298	youzhengguonei	挂号信	,Sort:	1,Status: 2 },
	// 299	xiaohongmao	北青小红帽	,Sort:	1,Status: 2 },
	// 300	lbbk	宝凯物流	,Sort:	1,Status: 2 },
	// 301	byht	博源恒通	,Sort:	1,Status: 2 },
	// 302	idada	百成大达物流	,Sort:	1,Status: 2 },
	// 303	baitengwuliu	百腾物流	,Sort:	1,Status: 2 },
	// 304	birdex	笨鸟海淘	,Sort:	1,Status: 2 },
	// 305	bsht	百事亨通	,Sort:	1,Status: 2 },
	// 306	dayang	大洋物流快递	,Sort:	1,Status: 2 },
	// 307	dechuangwuliu	德创物流	,Sort:	1,Status: 2 },
	// 308	donghanwl	东瀚物流	,Sort:	1,Status: 2 },
	// 309	dfpost	达方物流	,Sort:	1,Status: 2 },
	// 310	dongjun	东骏快捷物流	,Sort:	1,Status: 2 },
	// 311	dindon	叮咚澳洲转运	,Sort:	1,Status: 2 },
	// 312	dazhong	大众佐川急便	,Sort:	1,Status: 2 },
	// 313	decnlh	德中快递	,Sort:	1,Status: 2 },
	// 314	dekuncn	德坤供应链	,Sort:	1,Status: 2 },
	// 315	eshunda	俄顺达	,Sort:	1,Status: 2 },
	// 316	ewe	EWE全球快递	,Sort:	1,Status: 2 },
	// 317	fedexuk	FedEx英国	,Sort:	1,Status: 2 },
	// 318	fox	FOX国际速递	,Sort:	1,Status: 2 },
	// 319	rufengda	凡客如风达	,Sort:	1,Status: 2 },
	// 320	fandaguoji	颿达国际快递	,Sort:	1,Status: 2 },
	// 321	hnfy	飞鹰物流	,Sort:	1,Status: 2 },
	// 322	flysman	飞力士物流	,Sort:	1,Status: 2 },
	// 323	sccod	丰程物流	,Sort:	1,Status: 2 },
	// 324	farlogistis	泛远国际物流	,Sort:	1,Status: 2 },
	// 325	gsm	GSM	,Sort:	1,Status: 2 },
	// 326	gaticn	GATI快递	,Sort:	1,Status: 2 },
	// 327	gts	GTS快递	,Sort:	1,Status: 2 },
	// 328	gangkuai	港快速递	,Sort:	1,Status: 2 },
	// 329	gtsd	高铁速递	,Sort:	1,Status: 2 },
	// 330	tiandihuayu	华宇物流	,Sort:	1,Status: 2 },
	// 331	huangmajia	黄马甲快递	,Sort:	1,Status: 2 },
	// 332	ucs	合众速递	,Sort:	1,Status: 2 },
	// 333	huoban	伙伴物流	,Sort:	1,Status: 2 },
	// 334	nedahm	红马速递	,Sort:	1,Status: 2 },
	// 335	huiwen	汇文配送	,Sort:	1,Status: 2 },
	// 336	nmhuahe	华赫物流	,Sort:	1,Status: 2 },
	// 337	hangyu	航宇快递	,Sort:	1,Status: 2 },
	// 338	minsheng	闽盛物流	,Sort:	1,Status: 2 },
	// 339	riyu	日昱物流	,Sort:	1,Status: 2 },
	// 340	sxhongmajia	山西红马甲	,Sort:	1,Status: 2 },
	// 341	syjiahuier	沈阳佳惠尔	,Sort:	1,Status: 2 },
	// 342	shlindao	上海林道货运	,Sort:	1,Status: 2 },
	// 343	shunjiefengda	顺捷丰达	,Sort:	1,Status: 2 },
	// 344	subida	速必达物流	,Sort:	1,Status: 2 },
	// 345	bphchina	速方国际物流	,Sort:	1,Status: 2 },
	// 346	sendtochina	速递中国	,Sort:	1,Status: 2 },
	// 347	suning	苏宁快递	,Sort:	1,Status: 2 },
	// 348	sihaiet	四海快递	,Sort:	1,Status: 2 },
	// 349	tianzong	天纵物流	,Sort:	1,Status: 2 },
	// 350	chinatzx	同舟行物流	,Sort:	1,Status: 2 },
	// 351	nntengda	腾达速递	,Sort:	1,Status: 2 },
	// 352	sd138	泰国138	,Sort:	1,Status: 2 },
	// 353	tongdaxing	通达兴物流	,Sort:	1,Status: 2 },
	// 354	tlky	天联快运	,Sort:	1,Status: 2 },
	// 355	youshuwuliu	UC优速快递	,Sort:	1,Status: 2 },
	// 356	ueq	UEQ快递	,Sort:	1,Status: 2 },
	// 357	weitepai	微特派快递	,Sort:	1,Status: 2 },
	// 358	wtdchina	威时沛运	,Sort:	1,Status: 2 },
	// 359	wzhaunyun	微转运	,Sort:	1,Status: 2 },
	// 360	gswtkd	万通快递	,Sort:	1,Status: 2 },
	// 361	wotu	渥途国际速运	,Sort:	1,Status: 2 },
	// 362	xiyoute	希优特快递	,Sort:	1,Status: 2 },
	// 363	xilaikd	喜来快递	,Sort:	1,Status: 2 },
	// 364	xsrd	鑫世锐达	,Sort:	1,Status: 2 },
	// 365	xtb	鑫通宝物流	,Sort:	1,Status: 2 },
	// 366	xintianjie	信天捷快递	,Sort:	1,Status: 2 },
	// 367	xaetc	西安胜峰	,Sort:	1,Status: 2 },
	// 368	xianfeng	先锋快递	,Sort:	1,Status: 2 },
	// 369	sunspeedy	新速航	,Sort:	1,Status: 2 },
	// 370	xipost	西邮寄	,Sort:	1,Status: 2 },
	// 371	sinatone	信联通	,Sort:	1,Status: 2 },
	// 372	sunjex	新杰物流	,Sort:	1,Status: 2 },
	// 373	yundaexus	韵达美国件	,Sort:	1,Status: 2 },
	// 374	yxwl	宇鑫物流	,Sort:	1,Status: 2 },
	// 375	yitongda	易通达	,Sort:	1,Status: 2 },
	// 376	yiqiguojiwuliu	一柒物流	,Sort:	1,Status: 2 },
	// 377	yilingsuyun	亿领速运	,Sort:	1,Status: 2 },
	// 378	yujiawuliu	煜嘉物流	,Sort:	1,Status: 2 },
	// 379	gml	英脉物流	,Sort:	1,Status: 2 },
	// 380	leopard	云豹国际货运	,Sort:	1,Status: 2 },
	// 381	czwlyn	云南中诚	,Sort:	1,Status: 2 },
	// 382	sdyoupei	优配速运	,Sort:	1,Status: 2 },
	// 383	yongchang	永昌物流	,Sort:	1,Status: 2 },
	// 384	yufeng	御风速运	,Sort:	1,Status: 2 },
	// 385	yamaxunwuliu	亚马逊物流	,Sort:	1,Status: 2 },
	// 386	yousutongda	优速通达	,Sort:	1,Status: 2 },
	// 387	yishunhang	亿顺航	,Sort:	1,Status: 2 },
	// 388	yongwangda	永旺达快递	,Sort:	1,Status: 2 },
	// 389	ecmscn	易满客	,Sort:	1,Status: 2 },
	// 390	yingchao	英超物流	,Sort:	1,Status: 2 },
	// 391	edlogistics	益递物流	,Sort:	1,Status: 2 },
	// 392	yyexpress	远洋国际	,Sort:	1,Status: 2 },
	// 393	onehcang	一号仓	,Sort:	1,Status: 2 },
	// 394	ycgky	远成快运	,Sort:	1,Status: 2 },
	// 395	lineone	一号线	,Sort:	1,Status: 2 },
	// 396	ypsd	壹品速递	,Sort:	1,Status: 2 },
	// 397	vipexpress	鹰运国际速递	,Sort:	1,Status: 2 },
	// 398	el56	易联通达物流	,Sort:	1,Status: 2 },
	// 399	yyqc56	一运全成物流	,Sort:	1,Status: 2 },
	// 400	zhongtie	中铁快运	,Sort:	1,Status: 2 },
	// 401	ZTKY	中铁物流	,Sort:	1,Status: 2 },
	// 402	zzjh	郑州建华快递	,Sort:	1,Status: 2 },
	// 403	zhongruisudi	中睿速递	,Sort:	1,Status: 2 },
	// 404	zhongwaiyun	中外运速递	,Sort:	1,Status: 2 },
	// 405	zengyisudi	增益速递	,Sort:	1,Status: 2 },
	{Code: "sujievip", Name: "郑州速捷", Sort: 1, Status: 2},
	{Code: "zhichengtongda", Name: "至诚通达快递", Sort: 1, Status: 2},
	{Code: "zhdwl", Name: "众辉达物流", Sort: 1, Status: 2},
	{Code: "kuachangwuliu", Name: "直邮易", Sort: 1, Status: 2},
	{Code: "topspeedex", Name: "中运全速", Sort: 1, Status: 2},
	{Code: "otobv", Name: "中欧快运", Sort: 1, Status: 2},
	{Code: "zsky123", Name: "准实快运", Sort: 1, Status: 2},
	{Code: "donghong", Name: "东红物流", Sort: 1, Status: 2},
	{Code: "kuaiyouda", Name: "快优达速递", Sort: 1, Status: 2},
	{Code: "balunzhi", Name: "巴伦支快递", Sort: 1, Status: 2},
	{Code: "hutongwuliu", Name: "户通物流", Sort: 1, Status: 2},
	{Code: "xianchenglian", Name: "西安城联速递", Sort: 1, Status: 2},
	{Code: "youbijia", Name: "邮必佳", Sort: 1, Status: 2},
	{Code: "feiyuan", Name: "飞远物流", Sort: 1, Status: 2},
	{Code: "chengji", Name: "城际速递", Sort: 1, Status: 2},
	{Code: "huaqi", Name: "华企快运", Sort: 1, Status: 2},
	{Code: "yibang", Name: "一邦快递", Sort: 1, Status: 2},
	{Code: "citylink", Name: "CityLink快递", Sort: 1, Status: 2},
	{Code: "meixi", Name: "美西快递", Sort: 1, Status: 2},
	{Code: "zbkd", Name: "重磅快递", Sort: 100, Status: 2},
	{Code: "123456789", Name: "众邦快递", Sort: 1, Status: 1},
}

//@description: sys_expresses 表数据初始化
func (a *express) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.Express{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_expresses 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&expresses).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_expresses 表初始数据成功!")
		return nil
	})
}
