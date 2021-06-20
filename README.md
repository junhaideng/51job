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

### 效果
```csv
job_name, job_href, company_name, company_href, salary, workarea, issuedate, jobwelf
运行维护工程师,https://jobs.51job.com/beijing/129507148.html?s=sou_sou_soulb&t=0_0,中国互联网络信息中心（CNNIC）,https://jobs.51job.com/all/co2738.html,1.5-2万/月,北京,2021-06-19 04:01:44,五险一金 年终奖金 补充医疗保险 弹性工作 定期体检
边缘计算高级研发工程师,https://jobs.51job.com/beijing-cpq/131901519.html?s=sou_sou_soulb&t=0_0,慧博云通科技股份有限公司,https://jobs.51job.com/all/co5097835.html,1-1.5万/月,北京-昌平区,2021-06-18 22:09:16,
区块链钱包技术总监,https://jobs.51job.com/beijing-hdq/130423137.html?s=sou_sou_soulb&t=0_0,神话科技传媒（深圳）有限公司上海分公司,https://jobs.51job.com/all/co5605982.html,5-7万/月,北京-海淀区,2021-06-18 18:33:40,弹性工作 绩效奖金 节日福利
Golang开发工程师,https://jobs.51job.com/beijing-xcq/131523305.html?s=sou_sou_soulb&t=0_0,北京信安世纪科技股份有限公司,https://jobs.51job.com/all/co107634.html,1-2.5万/月,北京-西城区,2021-06-18 18:12:13,五险一金 餐饮补贴 带薪年假 通讯补贴 周末双休 节日福利 专业培训 补充医疗保险 定期体检
云平台测试工程师,https://jobs.51job.com/beijing-hdq/130980710.html?s=sou_sou_soulb&t=0_0,上海理想信息产业（集团）有限公司,https://jobs.51job.com/all/co244754.html,1.5-2万/月,北京-海淀区,2021-06-18 17:51:57,
openstack运维工程师,https://jobs.51job.com/beijing/129746172.html?s=sou_sou_soulb&t=0_0,上海微创软件股份有限公司,https://jobs.51job.com/all/co210498.html,1.2-1.5万/月,北京,2021-06-18 17:46:11,五险一金 补充医疗保险 免费班车 弹性工作 定期体检 绩效奖金 年终奖金 专业培训 员工旅游 节日福利
全栈开发工程师,https://jobs.51job.com/beijing-dcq/131640482.html?s=sou_sou_soulb&t=0_0,上海本原网络科技有限公司,https://jobs.51job.com/all/co6033389.html,1.2-2万/月,北京-东城区,2021-06-18 17:45:48,做五休二 周末双休 弹性工作
运维工程师,https://jobs.51job.com/beijing-syq/132549045.html?s=sou_sou_soulb&t=0_0,雅昌文化集团-互联网及雅昌平台,https://jobs.51job.com/all/co3998012.html,1.5-2万/月,北京-顺义区,2021-06-18 16:16:32,
知名银行金融服务公司 测试工程师,https://jobs.51job.com/beijing-cyq/130403146.html?s=sou_sou_soulb&t=0_0,万宝盛华企业管理咨询（上海）有限公司,https://jobs.51job.com/all/co5761023.html,1-1.5万/月,北京-朝阳区,2021-06-18 15:32:16,
客户经理/Account Manager,https://jobs.51job.com/beijing/131964477.html?s=sou_sou_soulb&t=0_0,PTL Group,https://jobs.51job.com/all/co2091769.html,2-2.5万/月,北京,2021-06-18 15:21:00,五险一金 补充医疗保险 定期体检
Business development manager,https://jobs.51job.com/beijing-cyq/132099920.html?s=sou_sou_soulb&t=0_0,世泓（上海）企业管理咨询服务有限公司,https://jobs.51job.com/all/co5751452.html,2-2.5万/月,北京-朝阳区,2021-06-18 14:54:57,
```