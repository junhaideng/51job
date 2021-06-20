### 51Job前程无忧工作岗位爬取

爬取[前程无忧](https://www.51job.com/)上的招聘岗位

### 使用方式
```
go get github.com/junhaideng/51job
```
```go
// 创建一个新的Job，指定log，数据保存地方以及关键字
j := job.NewJob(zap.NewExample(), f, keyword)
// 直接运行
j.Run()
```