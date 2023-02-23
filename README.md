# 第五届字节跳动青训营“抖声”项目

#### 技术选型

框架：Gin、Gorm

中间件：JWT、FFmpeg、MinIO、RabbitMQ、Redis

数据库：MySQL



#### 项目配置

- config.go中MinIO配置Endpoint变量应改为服务器（即本机）IP+port9000
- FFmpeg: 安装下载地址：https://www.gyan.dev/ffmpeg/builds/#release-builds [ffmpeg-5.0.1-essentials_build.zip ] (https://www.gyan.dev/ffmpeg/builds/packages/ffmpeg-5.0.1-essentials_build.zip)77 MB [.sha256](https://www.gyan.dev/ffmpeg/builds/packages/ffmpeg-5.0.1-essentials_build.zip.sha256)
- MinIO:windows 服务端下载地址： https://dl.minio.io/server/minio/release/windows-amd64/minio.exe    执行命令：` .\minio.exe server [存储目录]`
  - MinIO需设置永久访问链接

- 安装rabbitMQ
- 安装redis，启动服务端：`redis-server.exe redis.windows.conf`
  - 设置密码 tiktok：config set requirepass root 
  - 设置持久化： config set stop-writes-on-bgsave-error no



#### 代码结构

```undefined
├─.idea
├─config  配置
├─controller 控制器
├─dao  数据库
├─data  结构体
├─gateway  路由
├─middleware  中间件
│  ├─FFmpeg  视频截图
│  ├─jwt     鉴权
│  ├─MinIO   对象存储
│  ├─rabbitMQ 消息队列
│  └─redis    缓存
├─public      
└─service     服务层
    └─relation
```



#### 功能测试

功能测试设计：

- jjk上传两个视频（测试开始前）
- 视频流接口  -> tjt注册 -> tjt登录 -> tjt投稿视频 -> tjt发布列表及用户信息
- jjk登录 -> jjk点赞 -> jjk喜欢列表 -> jjk评论 -> 查看评论列表
- jjk关注 -> 关注列表 -> 好友列表
- tjt粉丝列表 -> 关注 ->好友列表

测试结果：[https://www.bilibili.com/video/BV1rM41177ka/?spm_id_from=333.999.0.0](https://juqnqytxyi.feishu.cn/docx/S92Ydhk4FoIVNkxClpicQO8inl4)



#### 团队成员

| **团队成员** |    **主要贡献**    |
| :----------: | :----------------: |
|    冀锦康    |      视频模块      |
|    王明贤    |      用户模块      |
|    王重人    | 关注模块、消息模块 |
|    谭欣妍    |      评论模块      |
|     周珂     |      点赞模块      |



未完待续。。。
