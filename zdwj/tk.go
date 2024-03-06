package zdwj

import (
	"fmt"
	"github.com/Lofanmi/pinyin-golang/pinyin"
	gopy "github.com/mozillazg/go-pinyin"
	"strings"
)

func init() {
	return
	pinyin.NewDict()
	a := gopy.NewArgs()
	fmt.Println(a)
	str := `持才傲物，才士，裸衣骂曹操—祢衡
	杨修 持才傲物，才士，曹操主薄一一
	曹植 七步成诗，才士，才高八斗一一
	孔融 刚正不阿，才士，北海太守一一
	郭嘉 官渡之战，曹操，谋士一一
	马良 马氏五常，谋士，夷陵之战一一
	马谡 马氏五常，谋士，街亭之战一一
	文鸯 叛魏投吴，七进七出，复投魏国——
	诸葛诞 叛魏投吴，寿春，举兵一一。
	周瑜 智力很高，苦肉计，赤壁之战一一
	黄盖 武力很高，苦肉计，赤壁之战一一
	司马师 三马同槽，擅行废立，东兴之战——
	司马懿 三马同槽，狼顾之相，高平陵事变——
	司马懿 被围之人，上方谷，火烧一一
	诸葛亮 设计之人，上方谷，火烧——
	魏延 诱敌之人，上方谷，火烧一一
	孙策 江东统一战，主帅，小霸王—一
	黄忠 定军山之战神射手，射杀夏侯渊一一
	司马炎 以晋代魏，文明皇后之子，统一三国一
	诸葛亮 主帅，街亭之战，一出祁山一一
	太史慈 东莱之人，神射手，酣斗小霸王——
	诸葛恪 主帅，官至太傅，东兴之战一一
	张郃 武将，街亭之战，五子良将一一
	张辽 逍遥津之战，五子良将一一
	甘宁 逍遥津之战，吴国，统率很高一一
	凌统 逍遥津之战，吴国，武力很高一—
	孙权 逍遥津之战，吴国，政治很高一一
	徐晃 樊城之战，步步为营，武力很高一一
	董卓 西凉铁骑，凉州军阀，武力很高—一
	曹操 华容道，败走，统率很高一一
	关羽 华容道，赤兔，武力很高—一
	吕布 白门楼，赤兔，武力很高一一
	关羽 麦城，败走，赤兔一—
	丁奉 东兴之战，江表之虎臣，雪中奋短兵一一
	甘宁 锦帆贼，江表之虎臣，官至西陵太守—一
	周泰 肤如刻画，江表之虎臣，英勇救主一一
	程普 第一位，江表之虎臣，饮醇自醉一一
	韩当 开国老将，江表之虎臣，长于弓箭一一
	关羽 挂印封金，刘备之弟，五虎上将一一
	张飞 挑灯夜战，刘备之弟，五虎上将一一
	马超 西凉铁骑，挑灯夜战，五虎上将——
	孙尚香 武力很高，绑架阿斗的人，截江夺斗—一
	赵云 武力很高，救回阿斗的人，截江夺斗一一
	典韦 死亡，宛城，曹操宿卫猛将一一
	邓艾 偷渡阴平，灭蜀元勋，先入成都一一
	姜维 死于兵变，抗魏先锋，谋害邓艾一一
	钟会 死于兵变，灭蜀元勋，谋害邓艾一一
	陆抗 东吴重臣，寿春，救援一一
	孙坚 被黄祖所杀，长沙太守，讨董主力一一
	袁绍 提议之人，招揽董卓入京，十常侍之乱一一
	何进 招揽董卓入京，下令之人，十常侍之乱一一
	孙权 劝吕蒙学习，碧眼儿，称帝一一
	鲁肃 东吴四杰，赤壁，联刘抗曹——
	陆逊 东吴四杰，火烧，夷陵一一
	吕蒙 东吴四杰，白衣渡江，荆州一一
	沮授 袁绍，谋士，官渡之战一一
	吕伯奢 收留曹操，曹操逃离洛阳，被曹操所杀一一
	陈宫 曹操逃离洛阳，反目离去，救下曹操一一
	孙权出生于182年   182
	孙策出生于公元175年   175
	公元221年孙权封吴王   221
	孙权女儿孙鲁班，先嫁周循(周瑜之子)，再嫁全琮。
	孙坚是孙权、孙策的父亲，出生于公元155年
	孙坚被当作是孙武的后裔
	袁绍被众人推选为讨伐军盟主，孙坚请缨当先锋
	孙坚 虎牢关战役登场时，头戴着的是红色的头巾
	周瑜 羽扇纶巾，谈笑间，樯橹灰飞烟灭，指的是
	张纮和张昭 孙策借兵走扬州，周瑜向其推荐“二张”分别为
	孙坚 《三国演义》之初众诸侯中违约藏匿玉玺的是
	“百骑劫魏营”的甘宁杀了凌统的父亲。
	徐庶 曹仁布八门金锁阵被 识破
	吕蒙 白衣渡江典故中的主角是
	184 张角，死亡，起义一一公元184年
	200 官渡之战，孙策死亡一一公元200年
	208 赤壁之战，曹操封汉丞相，刘表死亡一一公元208年
	220 曹丕，称帝，曹魏一一公元220年
	210 铜雀台，曹操，建成一一公元210年
	221 称帝，刘备，蜀汉一一公元221年
	222 孙权，称王，东吴一一公元222年
	229 孙权，称帝，东吴一一公元229年
	青缸剑 被赵云所夺，曹操兵器，夏侯恩配之一一
	青龙偃月刀 所用兵器，演义，过五关斩六将一一
	溯 所用兵器，正史，过五关斩六将一一
	丈八蛇矛 长坂坡，蜀国武将，长兵一一
	古锭刀 汜水关，讨伐董卓，孙坚兵器—一
	七星刀 刺杀董卓，失败献刀，曹操兵器一一
	太平清领道 著作，道学，于吉一—《》
	遁甲天书 左慈，著作，天书——《》
	孟德新书 曹操，著作，兵法——《》
	赤壁之战 蜀吴，战役，曹操一一
	夷陵之战	火烧连营，蜀吴，战役一一
	易京之战 公孙瓒，袁绍，战役一一
	讨董之战 袁绍，孙坚，曹操一一
	虎豹骑 骑兵，魏国，精锐士兵——
	虎卫 步兵，魏国，精锐士兵—一
	西凉铁骑 骑兵，蜀国，精锐士兵一一
	白毦兵 步兵，蜀国，精锐士兵—一
	绝影 死亡，宛城，曹操爱驹—一
	小沛 刘备留驻之地，三让徐州，第二次一一
	徐州 刘备留驻之地，三让徐州，第三次一一
	关羽  刘备部下，三让徐州，武将一一
	曹操生于155年，在207年著《观沧海》
	曹腾 曹操的祖父是
	曹操 宁教我负天下人休教天下人负我——
	曹操 铜雀台是曹操下令修建的   
	曹休 被曹操称为“吾家千里驹”的人是   
	‘吉利’是曹操的小名  曹操  吉利
	爪黄飞电 是曹操的坐骑，武器是倚天剑
	建安二十四年，曹操因受关羽的威胁，牵许都以避其锐
	马超 曹操运用离间计打败了
	曹植 煮豆燃豆萁，豆在釜中泣，此诗是所作
	曹真 不是曹操的儿子
	许劭 直言曹操是“治世之能臣，乱世之奸雄”的是
	曹昂 曹安民 典韦  宛城之战，曹操长子曹昂、侄子曹安民、大将典韦被杀。
	袁谭 南皮之战 曹操于中击败了
	尹楷 毛城之战 曹操在中击败了
	张让 曹操曾带刀闯入的府邸
	荀或 向曹操献“二虎竞食”及“驱虎吞狼”之计的人是
	荀或 三国志中四胜四败论是提出的
	郭嘉 三国演义中十胜十败论是提出的
	五子良将是  张辽、乐进、于禁、张郃、徐晃。
	曹奂 魏国的末代皇帝是
	威震逍遥津的是张辽，张辽字文远
	夏侯敦 三国中，在战场上拔矢啖睛的是 
	徐晃 被曹操赞誉有周亚夫之风
	许褚 拥有虎痴称号的魏国武将是
	河北双雄是颜良、文丑
	夏侯杰 长板桥被张飞吓到跌落下马的是
	司马懿诛杀曹爽之后，掌握了魏国的军政大权
	双铁戟 是典韦的武器
	桃园三结义，在桃园中痛饮有三百多人
	桃园三结义，除了用马做祭祀物品，还用了黑牛
	桃园三英中，最早逝世的是关羽
	刘关张出道后杀死的第一个副将是邓茂
	刘备生于161年，是中山靖王刘胜之子，武器是双股剑
	刘备家住楼桑村，死于白帝城，享年63岁
	吴懿 刘备入蜀后娶了的妹妹
	刘焉 发榜招军时，刘备 二十八岁
	张宝被刘备射中左臂
	卢植 刘备十五岁时，其母就让他外出游学，曾拜在门下“
	刘备 人情势利古犹今，谁识英雄是白身”这句话指的是
	张飞 被张达、范强所杀。
	张飞 三国中喝断长板桥的主角是
	张飞曾痛骂吕布为三姓家奴
	三国的常胜将军是赵云，在长坂坡大战时获得了青釭剑
	历史上，赵云最初从事于公孙瓒之手
	《三国演义》中，赵云最初从事于袁绍之手
	刘禅是甘氏之子
	临泪战败，关羽、关平一同被斩
	民间传说中，赵云自创的枪术叫七探盘蛇枪
	关羽的美髯有二尺长，他还有丹凤眼
	“冷艳锯”是关羽的兵器
	关羽是河东郡解县人
	关羽原来的字叫长生
	关羽离开曹操的时候，唯一带走的是赤兔马
	建年二十四年，于禁战败，向关羽投降
	百步穿杨指的是黄忠
	姜维和钟会密谋反叛被告发，最终被卫灌率兵剿灭
	蜀国五虎将中，马超最后一一个拜入刘备麾下
	马超的坐骑是里飞沙
	张松 《三国演义》中有云“一览无遗世所稀，谁知书信泄天机。未观玄德兴王业，先向成都血染衣。”这段话说的是
	诸葛亮躬耕之时，经常吟唱《梁甫吟》，以管仲、乐毅自比
	诸葛亮与黄月英是夫妻。
	刘琦受诸葛亮指点，效仿公子重耳出逃。
	三国演义中，诸葛亮舌战群儒并获得胜利
	诸葛亮挥泪斩马谡的原因是失守街亭
	三国中，出师未捷身先死，说的是诸葛亮.最后一次北伐病逝于五丈原
	诸葛亮在出师表中称赞向宠将军“性行淑均，晓畅军事”
	诸葛亮和诸葛恪是 叔侄 关系
	刘禅的皇后是 张星彩
	刘禅的母亲是 甘夫人
	袁绍 邀请 董卓 于 189 年入洛阳
	刘备 和袁术 相持不下时，吕布 乘其后防空虚袭击了 下邳
	七星宝刀 原为 王允 所有
	汉献帝的皇后是 曹节
	因迎奉汉灵帝之功，被封为长安乡候的宦官是  曹节
	汉灵帝继位后，宦官 曹节 肆意弄权
	南蛮第一智者是 朵思大王
	三力是‘武力、智力、天子之力’
	陈寿《三国志》作者 
	甄宓《洛神赋》中的洛神意指曹丕的妻子 
	张角 天公将军是。
	张宝 地公将军是。
	张梁 人公将军是。
	于吉 《太平清领道》第一发现者是
	马腾 汉灵帝末年，在西州与边章、韩遂等共同举兵起事的是
	马腾 曾被朝廷授予征西大将军之职
	虎豹骑  不是三国演义提过的特殊兵种
	仲颖 董卓字
	刘协 汉献帝的名讳是
	司马炎 最终统一了三国
	古今多少事，都付笑谈中，与青山依旧在，几度夕阳红呼应
	太平要术是南华老仙送给张角的三卷天书
	最早预言曹操“将来能安天下”的是何颛
	贩马 苏双、张世平两人从事的职业是
	季然 程幾的字
	280 公元280年三国归晋`
	py := `chicaiaowu，caishi，luoyimacaocao—miheng
	chicaiaowu，caishi，caocaozhubaoyiyiyangxiu
	qibuchengshi，caishi，caigaobadouyiyicaozhi
	gangzhengbua，caishi，beihaitaishouyiyikongrong
	guanduzhizhan，caocao，moushiyiyiguojia
	mashiwuchang，moushi，yilingzhizhanyiyimaliang
	mashiwuchang，moushi，jietingzhizhanyiyimasu
	panweitouwu，qijinqichu，futouweiguo——wenyang
	panweitouwu，shouchun，jubingyiyizhugedan。
	zhilihengao，kurouji，chibizhizhanyiyizhouyu
	wulihengao，kurouji，chibizhizhanyiyihuanggai
	sanmatongcao，shanhangfeili，dongxingzhizhan——simashi
	sanmatongcao，langguzhixiang，gaopinglingshibian——simayi
	beiweizhiren，shangfanggu，huoshaoyiyisimayi
	shejizhiren，shangfanggu，huoshao——zhugeliang
	youdizhiren，shangfanggu，huoshaoyiyiweiyan
	jiangdongtongyizhan，zhushuai，xiaobawang—yisunce
	dingjunshanzhizhanshensheshou，sheshaxiahouyuanyiyihuangzhong
	yijindaiwei，wenminghuanghouzhizi，tongyisanguoyisimayan
	zhushuai，jietingzhizhan，yichuqishanyiyizhugeliang
	donglaizhiren，shensheshou，handouxiaobawang——taishici
	zhushuai，guanzhitaifu，dongxingzhizhanyiyizhugeke
	wujiang，jietingzhizhan，wuziliangjiangyiyizhanghe
	xiaoyaojinzhizhan，wuziliangjiangyiyizhangliao
	xiaoyaojinzhizhan，wuguo，tongshuaihengaoyiyiganning
	xiaoyaojinzhizhan，wuguo，wulihengaoyi—lingtong
	xiaoyaojinzhizhan，wuguo，zhengzhihengaoyiyisunquan
	fanchengzhizhan，bubuweiying，wulihengaoyiyixuhuang
	xiliangtieqi，liangzhoujunfa，wulihengao—yidongzhuo
	huarongdao，baizou，tongshuaihengaoyiyicaocao
	huarongdao，chitu，wulihengao—yiguanyu
	baimenlou，chitu，wulihengaoyiyilvbu
	maicheng，baizou，chituyi—guanyu
	dongxingzhizhan，jiangbiaozhihuchen，xuezhongfenduanbingyiyidingfeng
	jinfanzei，jiangbiaozhihuchen，guanzhixilingtaishou—yiganning
	furukehua，jiangbiaozhihuchen，yingyongjiuzhuyiyizhoutai
	diyiwei，jiangbiaozhihuchen，yinchunzizuiyiyichengpu
	kaiguolaojiang，jiangbiaozhihuchen，changyugongjianyiyihandang
	guayinfengjin，liubeizhidi，wuhushangjiangyiyiguanyu
	tiaodengyezhan，liubeizhidi，wuhushangjiangyiyizhangfei
	xiliangtieqi，tiaodengyezhan，wuhushangjiang——machao
	wulihengao，bangjiaadouderen，jiejiangduodou—yisunshangxiang
	wulihengao，jiuhuiadouderen，jiejiangduodouyiyizhaoyun
	siwang，wancheng，caocaosuweimengjiangyiyidianwei
	touduyinping，mieshuyuanxun，xianruchengduyiyidengai
	siyubingbian，kangweixianfeng，mouhaidengaiyiyijiangwei
	siyubingbian，mieshuyuanxun，mouhaidengaiyiyizhonghui
	dongwuzhongchen，shouchun，jiuyuanyiyilukang
	beihuangzusuosha，changshataishou，taodongzhuliyiyisunjian
	tiyizhiren，zhaolandongzhuorujing，shichangshizhiluanyiyiyuanshao
	zhaolandongzhuorujing，xialingzhiren，shichangshizhiluanyiyihejin
	quanlvmengxuexi，biyaner，chengdiyiyisunquan
	dongwusijie，chibi，lianliukangcao——lusu
	dongwusijie，huoshao，yilingyiyiluxun
	dongwusijie，baiyidujiang，jingzhouyiyilvmeng
	yuanshao，moushi，guanduzhizhanyiyijushou
	shouliucaocao，caocaotaoliluoyang，beicaocaosuoshayiyilvboshe
	caocaotaoliluoyang，fanmuliqu，jiuxiacaocaoyiyichengong
	sunquanchushengyu182nian
	suncechushengyugongyuan175nian
	gongyuan221niansunquanfengwuwang
	sunquannversunluban，xianjiazhouxun(zhouyuzhizi)，zaijiaquancong。
	sunjianshisunquan、suncedefuqin，chushengyugongyuan155nian
	sunjianbeidangzuoshisunwudehouyi
	yuanshaobeizhongrentuixuanweitaofajunmengzhu，sunjianqingyingdangxianfeng
	hulaoguanzhanyidengchangshi，sunjiantoudaizhuodeshihongsedetoujin
	yushanlunjin，tanxiaojian，qiangluhuifeiyanmie，zhideshizhouyu
	suncejiebingzouyangzhou，zhouyuxiangqituijian“erzhang”fenbieweizhanghonghezhangzhao
	《sanguoyanyi》zhichuzhongzhuhouzhongweiyuecangniyuxideshisunjian
	“baiqijieweiying”deganningshalelingtongdefuqin。
	caorenbubamenjinsuozhenbeixushushipo
	baiyidujiangdianguzhongdezhujiaoshilvmeng
	zhangjiao，siwang，qiyiyiyigongyuan184nian
	guanduzhizhan，suncesiwangyiyigongyuan200nian
	chibizhizhan，caocaofenghanchengxiang，liubiaosiwangyiyigongyuan208nian
	caopi，chengdi，caoweiyiyigongyuan220nian
	tongquetai，caocao，jianchengyiyigongyuan210nian
	chengdi，liubei，shuhanyiyigongyuan221nian
	sunquan，chengwang，dongwuyiyigongyuan222nian
	sunquan，chengdi，dongwuyiyigongyuan229nian
	beizhaoyunsuoduo，caocaobingqi，xiahouenpeizhiyiyiqinggangjian
	suoyongbingqi，yanyi，guowuguanzhanliujiangyiyiqinglongyanyuedao
	suoyongbingqi，zhengshi，guowuguanzhanliujiangyiyisu
	changbanpo，shuguowujiang，changbingyiyizhangbashemao
	sishuiguan，taofadongzhuo，sunjianbingqi—yigudingdao
	cishadongzhuo，shibaixiandao，caocaobingqiyiyiqixingdao
	zhuzuo，daoxue，yujiyi—《taipingqinglingdao》
	zuoci，zhuzuo，tianshu——《dunjiatianshu》
	caocao，zhuzuo，bingfa——《mengdexinshu》
	shuwu，zhanyi，caocaoyiyichibizhizhan
	huoshaolianying，shuwu，zhanyiyiyiyilingzhizhan
	gongsunzan，yuanshao，zhanyiyiyiyijingzhizhan
	yuanshao，sunjian，caocaoyiyitaodongzhizhan
	qibing，weiguo，jingruishibing——hubaoqi
	bubing，weiguo，jingruishibing—yihuwei
	qibing，shuguo，jingruishibingyiyixiliangtieqi
	bubing，shuguo，jingruishibing—yibaierbing
	siwang，wancheng，caocaoaiju—yijueying
	liubeiliuzhuzhidi，sanrangxuzhou，dierciyiyixiaopei
	liubeiliuzhuzhidi，sanrangxuzhou，disanciyiyixuzhou
	liubeibuxia，sanrangxuzhou，wujiangyiyiguanyu
	caocaoshengyu155nian，zai207nianzhu《guancanghai》
	caocaodezufushicaoteng
	ningjiaowofutianxiarenxiujiaotianxiarenfuwo——shicaocaoshuode
	tongquetaishicaocaoxialingxiujiande
	beicaocaochengwei“wujiaqianliju”derenshicaoxiu
	‘jili’shicaocaodexiaoming
	zhaohuangfeidianshicaocaodezuoqi，wuqishiyitianjian
	jiananershisinian，caocaoyinshouguanyudeweixie，qianxuduyibiqirui
	caocaoyunyonglijianjidabailemachao
	zhudourandouqi，douzaifuzhongqi，cishishicaozhisuozuo
	caozhenbushicaocaodeerzi
	zhiyancaocaoshi“zhishizhinengchen，luanshizhijianxiong”deshixushao
	wanchengzhizhan，caocaochangzicaoang、zhizicaoanmin、dajiangdianweibeisha。
	caocaoyunanpizhizhanzhongjibaileyuantan
	caocaozaimaochengzhizhanzhongjibaileyinkai
	caocaozengdaidaochuangruzhangrangdefudi
	xiangcaocaoxian“erhujingshi”ji“quhutunlang”zhijiderenshixunhuo
	sanguozhizhongsishengsibailunshixunhuotichude
	sanguoyanyizhongshishengshibailunshiguojiatichude
	wuziliangjiangshizhangliao、lejin、yujin、zhanghe、xuhuang。
	weiguodemodaihuangdishicaohuan
	weizhenxiaoyaojindeshizhangliao，zhangliaoziwenyuan
	sanguozhong，zaizhanchangshangbashidanjingdeshixiahoudun
	xuhuangbeicaocaozanyuyouzhouyafuzhifeng
	yongyouhuchichenghaodeweiguowujiangshixuzhu
	hebeishuangxiongshiyanliang、wenchou
	changbanqiaobeizhangfeixiadaodielaxiamadeshixiahoujie
	simayizhushacaoshuangzhihou，zhangwoleweiguodejunzhengdaquan
	shuangtiejishidianweidewuqi
	taoyuansanjieyi，zaitaoyuanzhongtongyinyousanbaiduoren
	taoyuansanjieyi，chuleyongmazuojisiwupin，huanyongleheiniu
	taoyuansanyingzhong，zuizaoshishideshiguanyu
	liuguanzhangchudaohoushasidediyigefujiangshidengmao
	liubeishengyu161nian，shizhongshanjingwangliushengzhizi，wuqishishuanggujian
	liubeijiazhulousangcun，siyubaidicheng，xiangnian63sui
	liubeirushuhouqulewuyidemeimei
	liuyanfabangzhaojunshi，liubeiershibasui
	zhangbaobeiliubeishezhongzuobi
	liubeishiwusuishi，qimujiurangtawaichuyouxue，zengbaizailuzhimenxia“
	renqingshiliguyoujin，shuishiyingxiongshibaishen”zhejuhuazhideshiliubei
	zhangfeibeizhangda、fanqiangsuosha。
	sanguozhongheduanchangbanqiaodezhujiaoshizhangfei
	zhangfeizengtongmalvbuweisanxingjianu
	sanguodechangshengjiangjunshizhaoyun，zaichangbanpodazhanshihuodeleqinggangjian
	lishishang，zhaoyunzuichucongshiyugongsunzanzhishou
	《sanguoyanyi》zhong，zhaoyunzuichucongshiyuyuanshaozhishou
	liuchanshiganshizhizi
	linleizhanbai，guanyu、guanpingyitongbeizhan
	minjianchuanshuozhong，zhaoyunzichuangdeqiangshujiaoqitanpansheqiang
	guanyudemeiranyouerchichang，tahuanyoudanfengyan
	“lengyanju”shiguanyudebingqi
	guanyushihedongjunjiexianren
	guanyuyuanlaidezijiaochangsheng
	guanyulikaicaocaodeshihou，weiyidaizoudeshichituma
	jiannianershisinian，yujinzhanbai，xiangguanyutoujiang
	baibuchuanyangzhideshihuangzhong
	jiangweihezhonghuimimoufanpanbeigaofa，zuizhongbeiweiguanshuaibingjiaomie
	shuguowuhujiangzhong，machaozuihouyiyigebairuliubeihuixia
	machaodezuoqishilifeisha
	《sanguoyanyi》zhongyouyun“yilanwuyishisuoxi，shuizhishuxinxietianji。weiguanxuandexingwangye，xianxiangchengduxieranyi。”zheduanhuashuodeshizhangsong
	zhugelianggonggengzhishi，jingchangyinchang《liangfuyin》，yiguanzhong、leyizibi
	zhugeliangyuhuangyueyingshifuqi。
	liuqishouzhugeliangzhidian，xiaofanggongzizhongerchutao。
	sanguoyanyizhong，zhugeliangshezhanqunrubinghuodeshengli
	zhugelianghuileizhanmasudeyuanyinshishishoujieting
	sanguozhong，chushiweijieshenxiansi，shuodeshizhugeliang.zuihouyicibeifabingshiyuwuzhangyuan
	zhugeliangzaichushibiaozhongchengzanxiangchongjiangjun“xinghangshujun，xiaochangjunshi”
	zhugelianghezhugekeshishuzhiguanxi
	liuchandehuanghoushizhangxingcai
	liuchandemuqinshiganfuren
	yuanshaoyaoqingdongzhuoyu189nianruluoyang
	liubeiheyuanshuxiangchibuxiashi，lvbuchengqihoufangkongxuxijilexiapi
	qixingbaodaoyuanweiwangyunsuoyou
	hanxiandidehuanghoushicaojie
	yinyingfenghanlingdizhigong，beifengweichanganxianghoudehuanguanshicaojie
	hanlingdijiweihou，huanguancaojiesiyinongquan
	nanmandiyizhizheshiduosidawang
	sanlishi‘wuli、zhili、tianzizhili’
	《sanguozhi》zuozhechenshou
	《luoshenfu》zhongdeluoshenyizhicaopideqizizhenmi
	tiangongjiangjunshizhangjiao。
	digongjiangjunshizhangbao。
	rengongjiangjunshizhangliang。
	《taipingqinglingdao》diyifaxianzheshiyuji
	hanlingdimonian，zaixizhouyubianzhang、hansuidenggongtongjubingqishideshimateng
	matengzengbeizhaotingshouyuzhengxidajiangjunzhizhi
	hubaoqibushisanguoyanyitiguodeteshubingzhong
	dongzhuozizhongying
	hanxiandideminghuishiliuxie
	zuizhongsimayantongyilesanguo
	gujinduoshaoshi，dufuxiaotanzhong，yuqingshanyijiuzai，jiduxiyanghonghuying
	taipingyaoshushinanhualaoxiansonggeizhangjiaodesanjuantianshu
	zuizaoyuyancaocao“jianglainengantianxia”deshihezhuan
	sushuang、zhangshipingliangrencongshidezhiyeshifanma
	chengjidezijiran
	gongyuan280niansanguoguijin`
	strs := strings.Split(str, "\n")
	pys := strings.Split(py, "\n")
	//pys := make([]string, len(string()))
	//for _, v := range strs {
	//	zh := dict.Convert(v, "/").None()
	//
	//}
	fmt.Println(len(pys), len(strs))
	lenP := len(strs)
	for {
		var cx string
		fmt.Scanf("%s", &cx)
		//fmt.Print("\033[2j")
		for i := 0; i < lenP; i++ {
			strp := pys[i]
			lenSr := len(cx)
			lenS := len(strp)
			isOk := false
			for i := 0; i < lenS-lenSr+1; i++ {
				tt := true
				for j := 0; j < lenSr; j++ {
					if strp[i+j] != cx[j] {
						tt = false
						break
					}
				}
				if tt {
					isOk = true
				}
			}
			if isOk {
				fmt.Println(strs[i])
			}
		}
		fmt.Println("结束")
	}
}
