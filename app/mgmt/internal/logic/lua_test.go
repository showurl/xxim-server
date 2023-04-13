package logic

import (
	"github.com/cherish-chat/xxim-server/common/utils"
	_ "github.com/yuin/gopher-lua"
	lua "github.com/yuin/gopher-lua"
	"testing"
	"time"
)

const script = `
-- 姓
local lastNamePool = {
	"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "朱", "秦", "尤", "许", "何", "吕", "施", "张", "孔", "曹", "严", "华", "金", "魏", "陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章", "云", "苏", "潘", "葛", "奚", "范", "彭", "郎", "鲁", "韦", "昌", "马", "苗", "凤", "花", "方", "俞", "任", "袁", "柳", "酆", "鲍", "史", "唐", "费", "廉", "岑", "薛", "雷", "贺", "倪", "汤", "滕", "殷", "罗", "毕", "郝", "邬", "安", "常", "乐", "于", "时", "傅", "皮", "卞", "齐", "康", "伍", "余", "元", "卜", "顾", "孟", "平", "黄", "和", "穆", "萧", "尹", "姚", "邵", "湛", "汪", "祁", "毛", "禹", "狄", "米", "贝", "明",
}
-- 名
local firstNamePool = {
	"伟", "刚", "勇", "毅", "俊", "峰", "强", "军", "平", "保", "东", "文", "辉", "力", "明", "永", "健", "世", "广", "志", "义", "兴", "良", "海", "山", "仁", "波", "宁", "贵", "福", "生", "龙", "元", "全", "国", "胜", "学", "祥", "才", "发", "武", "新", "利", "清", "飞", "彬", "富", "顺", "信", "子", "杰", "涛", "昌", "成", "康", "星", "光", "天", "达", "安", "岩", "中", "茂", "进", "林", "有", "坚", "和", "彪", "博", "诚", "先", "敬", "震", "振", "壮", "会", "思", "群", "豪", "心", "邦", "承", "乐", "绍", "功", "松", "善", "厚", "庆", "磊", "民", "友", "裕", "河", "哲", "江", "超", "浩", "亮", "政", "谦", "亨", "翔", "辉", "远", "旭", "鸿", "朋", "斌", "梁", "栋", "维", "启", "克", "伦", "翔", "旭", "秀英", "俊杰", "子墨", "家乐", "佳乐"
}
-- 定义颜色数组
local colors = {
	-- 红色类
	"FFC0CB", -- 粉红
	"FFB6C1", -- 淡粉红
	"FF69B4", -- 热情的粉红
	"FF1493", -- 深粉红
	"DB7093", -- 苍白的紫罗兰红色
	"C71585", -- 适中的紫罗兰红色
	"DA70D6", -- 兰花的紫色
	"D8BFD8", -- 蓟
	"DDA0DD", -- 李子
	"EE82EE", -- 紫罗兰
	"FF00FF", -- 纯红
	-- 绿色类
	"7CFC00", -- 草坪绿
	"7FFF00", -- 调色板中的绿色
	"ADFF2F", -- 绿黄色
	"7FFFD4", -- 美团色
	"00FF7F", -- 春天的绿色
	"00FA9A", -- 中间的春天的绿色
	"90EE90", -- 淡绿色
}

-- table 转 json
function table2json(result) 
	local json_result = "{"
	local is_first = true
	for k, v in pairs(result) do
		if not is_first then
			json_result = json_result .. ","
		end
		is_first = false
		json_result = json_result .. '"' .. k .. '":' .. '"' .. tostring(v) .. '"' 
	end
	json_result = json_result .. "}"
	
	-- 返回 JSON 字符串
	return json_result
end 

-- 随机姓
function rand_last_name(timestamp)
	-- 生成随机姓
	math.randomseed(timestamp)
	local random_last_name = lastNamePool[math.random(1, #lastNamePool)]
	return random_last_name
end

-- 随机名
function rand_first_name(timestamp)
	-- 生成随机名
	math.randomseed(timestamp)
	local random_first_name = firstNamePool[math.random(1, #firstNamePool)]
	return random_first_name
end

-- 随机生成头像
function rand_avatar(timestamp, last_name)
	-- 通过昵称得到一个唯一的数字
	math.randomseed(timestamp)
	local color = colors[math.random(#colors)]
	return string.format("https://fakeimg.pl/500x500/%s/FFF/?font=noto&text=%s", color, last_name)
end

-- lua 入口函数 输入参数为时间戳 返回值为 JSON 字符串 格式为 {name: "张三", avatar: "https://fakeimg.pl/500x500/FFC0CB/FFF/?font=noto&text=张"}
function main(timestamp)
	-- 生成随机姓
	local random_last_name = rand_last_name(timestamp)

	-- 生成随机名
	local random_first_name = rand_first_name(timestamp)

	-- 生成随机头像
	local random_avatar = rand_avatar(timestamp, random_last_name)
	local result = {name = random_last_name .. random_first_name, avatar = random_avatar}
	return result
end
`

func TestLua(t *testing.T) {
	l := lua.NewState()
	defer l.Close()
	results, err := utils.Lua.ExecLua(l, script, "main", lua.LNumber(time.Now().UnixNano()))
	if err != nil {
		t.Fatalf("exec lua failed. %s", err)
	}
	if results.Type() != lua.LTTable {
		t.Fatalf("results type is not table")
	}
	table := results.(*lua.LTable)
	name := table.RawGetString("name")
	avatar := table.RawGetString("avatar")
	t.Logf("name: %s, avatar: %s", name, avatar)
}
