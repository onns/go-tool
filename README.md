# 一些有的没的

## work-day

> 计算当前距离年底还需要工作的时间

## parse-holiday

> 把国务院的通知转换为日历的中间步骤
>
> 通知地址：http://www.gov.cn/zhengce/content/2021-10/25/content_5644835.htm
>
> 结果json的使用方式：https://github.com/onns/memories
>
>

### 使用方式

```bash
cd parse-hoilday
# 到`parse-hoilday`目录
pwd
# /Users/onns/Downloads/github/go-tool/parse-holiday
go build
# 构建
./parse-holiday
# 运行
ls 
# 会有个结果：china-public-holiday.json
```
