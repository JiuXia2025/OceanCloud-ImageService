## OceanCloud - Image

随手写的QEMU镜像管理服务，与Dashboard、主控制器、计算节点配套

目前还是未完工的状态仅供参考，刚学golang写的有点烂，开发中，后期可能会考虑重构

暂不提供已构建成品，需要自行关闭CGO然后编译得出程序本体

喜欢的话可以点个Star

### 使用说明

#### 初次安装

config.yaml中填入数据库信息（暂时只支持MySQL不支持SQLite）

然后命令行模式下执行初始化数据库服务：./ocimage setup

#### WebAPI模式

执行命令启动主服务：./ocimage start

关闭WebAPI只需Ctrl+C就可以优雅地关闭主服务

#### 命令行模式

```
查看帮助：./ocimage --help

启动主服务：./ocimage start
初始化数据库：./ocimage setup
镜像列表：./ocimage listImage
创建镜像：./ocimage createImage <镜像类型:{centos,ubuntu,debian,windows}> <镜像路径> <镜像显示名称>
删除镜像：./ocimage deleteImage <镜像ID>
查询镜像：./ocimage infoImage <镜像ID>
```

#### API文档

文档地址：[OceanCloud](https://apifox.com/apidoc/shared-dd3f0669-966a-4a18-8109-cc87189cbc71)

开发中，后续可能会有变动