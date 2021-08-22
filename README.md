# paladin-poker 策划配置

#### 介绍
策划根据需求修改xlsx中的数值并上传，自动编译生成对应的配置文件  
* 只生成数据 paladin.exe
* 生成数据和go桩文件 paladin.exe -go
* 生成数据和前后端桩文件 paladin.exe -go -cs

#### 使用方法
1. [新增需求]策划与程序沟通，确定xlsx中表文件结构与配置
1. [数值变化]直接修改对应xlsx中数值
1. 双击paladin.exe，生成新的json配置数据于data目录下

#### 解析规则
编写于config.toml， 使用toml解析格式
* workbook指定需要解析的xlsx表
* sheet指定解析的子表（列表形式）
* horizontal指定水平解析还是垂直解析（通常是false，只有全局配置会用到水平解析）
* enums枚举替换规则，将表中的字符串，按照规则，替换成指定表中的内容（一般是id/数值）
更详细的配置规则参见paladin说明：https://github.com/assalyn/paladin

### 策划表 ID 分配
* 房间配置:    200001-299999
* 积分经验:    300001-399999
* 商城配置:    400001-499999
* 礼物配置:    500001-599999
* 任务配置:    600001-699999

其他配置：全局配置表(列解析)
