#### 2021/11/25 更新

有好几个人加我问了些问题，基本问的都一样

* 关于QQ登录：如果没法使用只能改一下代码，自己做一个登录注册

* 关于存储：文件上传后是保存到阿里云OSS，如果要保存本地就修改一下逻辑

这个项目是我还在学校的时候做的，做的可能不是很好，当时是在B站看到的这个项目，那个UP主是用Java做的，我把前端部分拿了过来用Golang做了后端，因为是照着模板做的，所以登录是使用了QQ登录，当时想着正好学习一下接入QQ登录，这样也能很方便的显示头像，如果你不想用QQ登录需要自己实现一个登录注册~



___________




### 运行

首先需要在conf文件夹中配置好ini文件

```
go mod tidy
```

```
go run main.go
```

访问localhost:8080



### 项目截图

1. 使用情况

   ![status](https://git-xg.oss-cn-shanghai.aliyuncs.com/status.png)

2. 全部文件

   ![all](https://git-xg.oss-cn-shanghai.aliyuncs.com/all.png)

3. 上传文件

   ![upload](https://git-xg.oss-cn-shanghai.aliyuncs.com/upload.png)

4. 分类文件

   ![classify](https://git-xg.oss-cn-shanghai.aliyuncs.com/classify.png)

5. 帮助

   ![help](https://git-xg.oss-cn-shanghai.aliyuncs.com/help.png)
