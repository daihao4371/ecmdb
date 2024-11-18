# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased

## [v1.6.2](https://github.com/Duke1616/ecmdb/releases/tag/v1.6.2) - 2024-10-11

- [`b009298`](https://github.com/Duke1616/ecmdb/commit/b0092980d2b82ed8b2732c0f41854e8a49598b2d) fix: 替换新的菜单、权限部署SQL文件

## [v1.6.1](https://github.com/Duke1616/ecmdb/releases/tag/v1.6.1) - 2024-10-11

- [`8b4591a`](https://github.com/Duke1616/ecmdb/commit/8b4591a3dd0e5c283a3a0d9c7883c1c86176b273) fix: 调整任务运行顺序，修复 err.Error nil 的情况

## [v1.6.0](https://github.com/Duke1616/ecmdb/releases/tag/v1.6.0) - 2024-10-11

- [`81426bb`](https://github.com/Duke1616/ecmdb/commit/81426bb0245b053a55da510ba868602ca736ec91) chore: 任务状态新增调度失败情况
- [`9b6fd8b`](https://github.com/Duke1616/ecmdb/commit/9b6fd8b16a70dad0ffc70cee0caeb7f8402bf5a1) chore: 优化部分接口性能
- [`7c2a13a`](https://github.com/Duke1616/ecmdb/commit/7c2a13a7452d285baa71beee5858556248bccd24) chore: 删除部分弃用代码，优化函数名称
- [`3558f50`](https://github.com/Duke1616/ecmdb/commit/3558f50018e1b4fa31fbfb713aeea58d078122cd) chore: 模型关联 查询bug、etcd连接异常panic
- [`2e4d81e`](https://github.com/Duke1616/ecmdb/commit/2e4d81e20ae9aa0acff3363a71aec219f82e30d1) chore: 优化企业微信来源审批，消息通知
- [`fba3b6d`](https://github.com/Duke1616/ecmdb/commit/fba3b6d32ea0237574edead28b122210a81429bd) chore: 完善审批通过、驳回状态变更
- [`26a2c3b`](https://github.com/Duke1616/ecmdb/commit/26a2c3b8b123e4f27afbb0ce2ec5387285e3fb72) chore: 工单撤销
- [`488349a`](https://github.com/Duke1616/ecmdb/commit/488349a398e5a9fafc4836e0430d63ef2bb7ea49) chore: 修改查询工单状态数组
- [`995611b`](https://github.com/Duke1616/ecmdb/commit/995611b88aaecf9455fc00bef0636a20124f8369) chore: 工单历史
- [`ce2bdbd`](https://github.com/Duke1616/ecmdb/commit/ce2bdbddf9fa3c5a2218db6d6786e04172184140) chore: 优化工单列表提单人、处理人前端展示名称
- [`f33f3ce`](https://github.com/Duke1616/ecmdb/commit/f33f3ce1103523a67b69ee061eb26cc835bfa7e8) refactor: 重构工单消息通知，新增验证流程控制是否开启消息通知，封装 NotifierIntegration 接口支持多消息接收源
- [`d8c7414`](https://github.com/Duke1616/ecmdb/commit/d8c74147906a6471317f0c4b689382eae55650c4) chore: 前端提供接口支持
- [`c74b6e9`](https://github.com/Duke1616/ecmdb/commit/c74b6e9e5f37f317d61b868659d8bd5d08abceef) feat: 新增属性 link 字段, 支持跳转情况
- [`83303d8`](https://github.com/Duke1616/ecmdb/commit/83303d89302e39e269e70fe006493b9f930fbe78) feat: 完善用户管理、部门管理模版，前端联调
- [`98daea0`](https://github.com/Duke1616/ecmdb/commit/98daea0741364423d9b6debe855c092a55befab9) feat: 新增部分管理模块
- [`5df4ebf`](https://github.com/Duke1616/ecmdb/commit/5df4ebf225d80e9562ae93f0f5609531a287cff5) refactor: 改写 notify 消息通知，使用 notify.NotifierWrap 数组
- [`4b4b810`](https://github.com/Duke1616/ecmdb/commit/4b4b81017b0384b1b81574d1126dcaebff955dd5) feat: 添加CMDB相关，资产、属性修改接口
- [`5b48e4b`](https://github.com/Duke1616/ecmdb/commit/5b48e4b582f9c8d2e09f31ae0fe7f75f4f6b4d85) fix: 升级 golang 版本，dockerfile 打包镜像
- [`4df159c`](https://github.com/Duke1616/ecmdb/commit/4df159cf3c010a2c0e031bf0ea044ff73fe66720) fix: 参数传递错误
- [`cd0f1b5`](https://github.com/Duke1616/ecmdb/commit/cd0f1b55ef7f573fb9eb151d66c65fc04d9bdfcc) fix: 修复 wire 依赖注入
- [`29d485f`](https://github.com/Duke1616/ecmdb/commit/29d485f49cf20b5f17ab7df8fdcee49cee084d0a) feat: 当在飞书点击审批后，撤回消息, 防止显示过乱
- [`7018251`](https://github.com/Duke1616/ecmdb/commit/7018251c00b26e777410f1b6e855c7100c5f68ab) feat: 接入飞书回调，审批通过、拒绝
- [`fd58b79`](https://github.com/Duke1616/ecmdb/commit/fd58b790c07f67891c72d90125479a7cf33ac616) feat: 接入 enotify 消息通知
- [`2d3c5ae`](https://github.com/Duke1616/ecmdb/commit/2d3c5ae1933a383c53f3643282afcc74a1ec1d85) chore: 服务启动美化
- [`5ee4241`](https://github.com/Duke1616/ecmdb/commit/5ee424134223fe55eead2ef38cb7ab2f8bdbd52c) fix: getVersion 逻辑返回错误，修正
- [`11f168c`](https://github.com/Duke1616/ecmdb/commit/11f168ca476f31721e2c61f9198341df02782cb9) fix: Dockerfile 打包镜像传递版本信息
- [`d137736`](https://github.com/Duke1616/ecmdb/commit/d1377361bf9de9695c40933080cc0ad7b7c31714) fix: 版本号 compare 比较

## [v1.5.0](https://github.com/Duke1616/ecmdb/releases/tag/v1.5.0) - 2024-08-26

- [`148048b`](https://github.com/Duke1616/ecmdb/commit/148048bb6d39c3f22307030b8583257b9d53bfe6) fix: 修复第一次查询版本为空的情况
- [`3c7f2bf`](https://github.com/Duke1616/ecmdb/commit/3c7f2bf8d97f9576d78a877ae49e3acc181464dc) feat: 新增 initial 全量增量模式初始化数据方式
- [`95d2919`](https://github.com/Duke1616/ecmdb/commit/95d291978db52b142c013a368085c27a2dc04d34) fix: 添加github action

## [v1.4.1](https://github.com/Duke1616/ecmdb/releases/tag/v1.4.1) - 2024-08-23

- [`4b1d2fb`](https://github.com/Duke1616/ecmdb/commit/4b1d2fbaa7e046f1edd764cb9ea6ace6e4bf8543) fix: 删除bumps
- [`15c63e7`](https://github.com/Duke1616/ecmdb/commit/15c63e7147ab6ca31cb87c0742937fa07b84b57a) fix: upleft release
- [`5496a52`](https://github.com/Duke1616/ecmdb/commit/5496a529cdacf37064f5710efad9401bb3f9e4cd) action
- [`7b37e54`](https://github.com/Duke1616/ecmdb/commit/7b37e545c57a36c2db2d4cc6318e7c60452efc9b) action
- [`0f256f5`](https://github.com/Duke1616/ecmdb/commit/0f256f51a97cedd4babbe383c060c0a892ce8048) chorm: 修改 uplift.yaml

## [v1.4.0](https://github.com/Duke1616/ecmdb/releases/tag/v1.4.0) - 2024-08-23

- [`7031e93`](https://github.com/Duke1616/ecmdb/commit/7031e93d3de57c1132604dca38cf34f47cb293c5) 简单完善一下文档
- [`9c92a13`](https://github.com/Duke1616/ecmdb/commit/9c92a138eeab403b2a8a0bd4aae9310fd161730a) 完善用户登录逻辑
- [`0121aa1`](https://github.com/Duke1616/ecmdb/commit/0121aa1364a69b09c774549299d03b663e76cac2) 部署 compose 编写
- [`d520b62`](https://github.com/Duke1616/ecmdb/commit/d520b62e0e8543076363da17838eac231221c382) fixbug: 创建用户逻辑错误
- [`1861db0`](https://github.com/Duke1616/ecmdb/commit/1861db0daaaeb38b2f0c7f32e2d4dc0d38e0335f) fixbug：修复没有数据的情况，返回错误
- [`6e64c05`](https://github.com/Duke1616/ecmdb/commit/6e64c0550e4bd2aea961ffa4fae6b06bad5075b8) 完善部署文档
- [`93f5550`](https://github.com/Duke1616/ecmdb/commit/93f555046f4e73bf1d90c92cd843f27ff980a6c5) fix: 修复readme.md中安装步骤；增加了部署文档和安装脚本
- [`682c4a2`](https://github.com/Duke1616/ecmdb/commit/682c4a2f23c6e7d1ab7f224ac95b7b1f35df985d) fix: 修复readme.md文件中的IP地址
- [`8a72717`](https://github.com/Duke1616/ecmdb/commit/8a7271775c5fb8582d311b0bce2f225374a36c0a) Merge pull request #3 from cplinux98/fix_install
- [`5d55b3c`](https://github.com/Duke1616/ecmdb/commit/5d55b3c57e63d3863590acc6040a1bdd724516de) cobra 封装命令行启动
- [`eb8c7e5`](https://github.com/Duke1616/ecmdb/commit/eb8c7e5bfead407344b9cf250523ad191e6991af) GOPROXY 代理
- [`1c91695`](https://github.com/Duke1616/ecmdb/commit/1c9169560667af16e369ed6060ec795e0eb5ffbb) fix: 添加 changelog
- [`f944ebb`](https://github.com/Duke1616/ecmdb/commit/f944ebb277998ed2bdccf06eec07065db4a19294) feat: 简单实现，init 初始化数据，版本增量数据，但不支持降级操作
- [`fc914ca`](https://github.com/Duke1616/ecmdb/commit/fc914ca690721ad163011e71709c3866fc68a794) 添加临时

## [v1.3.0](https://github.com/Duke1616/ecmdb/releases/tag/v1.3.0) - 2024-08-15

- [`20e9cf6`](https://github.com/Duke1616/ecmdb/commit/20e9cf640b40bed2f2472a743ec619a3f68f1c2e) readme
- [`6997cfa`](https://github.com/Duke1616/ecmdb/commit/6997cfaaff00ebe59466971328852c74cc4718d5) 增强runner，对敏感变量脱敏
- [`7702d38`](https://github.com/Duke1616/ecmdb/commit/7702d3842fe2da447a822cb7c76dde798d753cbb) 任务变量脱敏
- [`f9708e3`](https://github.com/Duke1616/ecmdb/commit/f9708e33181c41b97a8087aa27ee55f1b64225e2) 变量数据库层面通过 AES 加密存储
- [`af1c40b`](https://github.com/Duke1616/ecmdb/commit/af1c40b51bd91985588a1b8e3d4c641a15446cb5) 配置文件 example 同步
- [`d18391b`](https://github.com/Duke1616/ecmdb/commit/d18391b65ba0947d0113eb6381b0a93b6a2971be) 新增重试状态
- [`18497a3`](https://github.com/Duke1616/ecmdb/commit/18497a3978f305870427f1aafab389bd65587122) fixbug: 逻辑问题，当触发修改时, 加密变量会变更为None
- [`b102824`](https://github.com/Duke1616/ecmdb/commit/b102824a8b99ce2d7e950c0d40f65aab8696cf9b) 权限设计
- [`1481baa`](https://github.com/Duke1616/ecmdb/commit/1481baa22ac8995a2f6f1a0708ce963a531e63ff) 去除casbin配置文件
- [`e127be6`](https://github.com/Duke1616/ecmdb/commit/e127be6ed96edf0240c0a02be17258428c3dae32) 集成 casbin 策略功能
- [`e9bff7e`](https://github.com/Duke1616/ecmdb/commit/e9bff7e1c84ec111ab12ff9aea0bae00338b284c) menu 菜单
- [`c4d29f0`](https://github.com/Duke1616/ecmdb/commit/c4d29f034a6125fc0775ff5232db00e97245d83d) api 注册
- [`8fc2866`](https://github.com/Duke1616/ecmdb/commit/8fc28664f84584b9b874e0e6411103c9c956b7da) 角色模块
- [`7f37e77`](https://github.com/Duke1616/ecmdb/commit/7f37e779ccd7027cd663e78020cdfced87d1ff34) 权限模块联调
- [`3560592`](https://github.com/Duke1616/ecmdb/commit/35605920aa58929ec3c4f8662576b86e8f3c4a36) 语法错误
- [`62e74ff`](https://github.com/Duke1616/ecmdb/commit/62e74ff9b43d44f70586e4a5b89dbc90f000caad) 角色完成 50%，待开发，用户模块进行对接联调
- [`09c59dc`](https://github.com/Duke1616/ecmdb/commit/09c59dc4aebc6fb29b8b4b0a9d0b1f7d56ed00bb) 用户模块
- [`477056d`](https://github.com/Duke1616/ecmdb/commit/477056d5517bd84b8fa4cdc9fe34003cb411478f) 补充部署逻辑，录入用户角色信息到casbin库中
- [`fdec76b`](https://github.com/Duke1616/ecmdb/commit/fdec76bb1a874d5e52a9ff7dae50b9fdaa35af3a) fixbug: casbin 没有 filter gourping的方法
- [`497cded`](https://github.com/Duke1616/ecmdb/commit/497cdedace21b814646a1d5a0123502b0c996b62) 动态权限返回给前端
- [`34ec849`](https://github.com/Duke1616/ecmdb/commit/34ec84917c5ae6b7e03be0b93bc07fc723f4f77c) 抽出 permission 模块，优化菜单与角色之间映射关系
- [`068e0aa`](https://github.com/Duke1616/ecmdb/commit/068e0aa5a2105ac864425f4849411069f2d80ed2) 修改路由名称
- [`67998aa`](https://github.com/Duke1616/ecmdb/commit/67998aa3e53cada125f5e400732d97cd788de226) 优化部分逻辑
- [`3c8fabb`](https://github.com/Duke1616/ecmdb/commit/3c8fabbc66dafc002c2cd2a0ef9ed909e2d1384e) 优化逻辑
- [`7786ba6`](https://github.com/Duke1616/ecmdb/commit/7786ba6186aad6e0801c0f6763dcaca6700e0909) 调整 gin 路由注册顺序
- [`bf519cb`](https://github.com/Duke1616/ecmdb/commit/bf519cbbfefc8bff028b87b1553bf23b55776aa7) fixbug: 获取用户菜单权限，通过sess获取用户id
- [`99f1704`](https://github.com/Duke1616/ecmdb/commit/99f17046fa40a33285b1d1cde0521bb515d3bab2) github action 替换阿里云镜像
- [`46016b6`](https://github.com/Duke1616/ecmdb/commit/46016b606260b1aac6e57ff89a0db847c449ec34) 去除缓存，报错

## [v1.2.0](https://github.com/Duke1616/ecmdb/releases/tag/v1.2.0) - 2024-07-29

- [`bfe93e4`](https://github.com/Duke1616/ecmdb/commit/bfe93e48a336ed64443af0369b14399eaea4a31b) 新增MQ：对接 wechat 回调信息
- [`1ae3e93`](https://github.com/Duke1616/ecmdb/commit/1ae3e9305b16f1f0c0e2e54507ecee5f08ea19cb) 同步wechat OA 模版信息
- [`cbc6b06`](https://github.com/Duke1616/ecmdb/commit/cbc6b06504f7eda6fff5dcaf52b4b786fb3ddbab) 优化处理逻辑，把调用wechat移动到service中
- [`d89bd77`](https://github.com/Duke1616/ecmdb/commit/d89bd77d9373aa42636429d80ede02fcfec0d8c1) 适配前端 form-create 数据录入数据库
- [`481e2e0`](https://github.com/Duke1616/ecmdb/commit/481e2e015a7d0ba852712d3bd638b67b194cf2a7) 工单模版，增加查询接口
- [`7bd7a5c`](https://github.com/Duke1616/ecmdb/commit/7bd7a5c949a5d735a41c901b6de24d9f6e8848a0) 新增 codebook 模块
- [`ea129e1`](https://github.com/Duke1616/ecmdb/commit/ea129e1ec104b1a9f97c0cf6461fd406120284e3) 去除 output
- [`66a028b`](https://github.com/Duke1616/ecmdb/commit/66a028bfaaa5b2e075ceeb4ccc349926c8aa550a) 完善codebook模块CRUD
- [`9e62597`](https://github.com/Duke1616/ecmdb/commit/9e625979ad084aeda4a4c842940a165ec4915024) 新增 worker 工作节点模块
- [`999275f`](https://github.com/Duke1616/ecmdb/commit/999275f9069da54d9b157eb1633d093418af4299) 对接 runner 执行器
- [`6797753`](https://github.com/Duke1616/ecmdb/commit/6797753f3ff498e8f2ac29ebfdd4f49cca195bd8) mongodb 配置 ioc引用
- [`85031d6`](https://github.com/Duke1616/ecmdb/commit/85031d6270bd58f6f25e94d57b7b1412344f4b7f) 修改节点状态
- [`65df10f`](https://github.com/Duke1616/ecmdb/commit/65df10f71cd416791bb24ea1c0491b62b14ba0da) 修改节点状态
- [`f454000`](https://github.com/Duke1616/ecmdb/commit/f454000e113bd3f252ccc396358584cc5fc0552e) README 文档
- [`dc51b6e`](https://github.com/Duke1616/ecmdb/commit/dc51b6e030cb0c51af5d19ec2acbd2fdd0d0c95c) 换行
- [`fdbffdb`](https://github.com/Duke1616/ecmdb/commit/fdbffdb87b2e76195d71f6c3d3ccee99f0dd5e81) 新增 list 接口
- [`53ab63f`](https://github.com/Duke1616/ecmdb/commit/53ab63f5acee00aa85a69ae05ff3bfb1cce2bd27) 节点注册，替换成 ETCD
- [`cc97b9e`](https://github.com/Duke1616/ecmdb/commit/cc97b9e687c2583ee1cac08022579fb6d456f13e) 服务启动，自动开启topic
- [`faf5780`](https://github.com/Duke1616/ecmdb/commit/faf578011c14a37747a40bbcea4861bccd1f9f0c) 优化完成 worker 注册
- [`b8e8afd`](https://github.com/Duke1616/ecmdb/commit/b8e8afd2c9b9c6714e4dd49d0f17dc96d8033737) 启动补偿机制，etcd + mongodb 数据库数据校对
- [`a741029`](https://github.com/Duke1616/ecmdb/commit/a7410295d7c6535e7e5c6b9737b732da8572fa42) runner 注册
- [`9b2bc2a`](https://github.com/Duke1616/ecmdb/commit/9b2bc2a544f96e4e03f52342973906dacb4b556d) 消息队列注册验证
- [`0ee24eb`](https://github.com/Duke1616/ecmdb/commit/0ee24ebfe039c3627358ece228c53a81c5448fea) 工单、策略模块
- [`6f7ebef`](https://github.com/Duke1616/ecmdb/commit/6f7ebef8b42e653d410f1a9016a1ad2acbf273d0) 添加 order 工单信息
- [`9dc0690`](https://github.com/Duke1616/ecmdb/commit/9dc069098d84695ca9123186810600b7687f1c62) 集成 easyflow
- [`2eec347`](https://github.com/Duke1616/ecmdb/commit/2eec347eaa178b8fb61c4c5f42ad95d56412b705) 适配 logic flow
- [`cc75201`](https://github.com/Duke1616/ecmdb/commit/cc75201b72a517b36b2f1b57814693c752cdc1cf) workflow crud
- [`9cbe0f5`](https://github.com/Duke1616/ecmdb/commit/9cbe0f5767d616ba7b0e914f8cc578dea2f5f356) 添加 LICENSE
- [`af79840`](https://github.com/Duke1616/ecmdb/commit/af798405902e0dfe074fac99952415f04689436e) easy-workflow 简单测试
- [`cc56c47`](https://github.com/Duke1616/ecmdb/commit/cc56c47216c736ef375b0a4e0f1b06d2e0bde974) 模版分组
- [`dc9f437`](https://github.com/Duke1616/ecmdb/commit/dc9f4376129d2f890e6222242dc5fab6deeb611e) 新增字段
- [`78c4ef0`](https://github.com/Duke1616/ecmdb/commit/78c4ef0ccca19cefd3889ca460d57769647bf67a) 聚合组数据查询
- [`f113f48`](https://github.com/Duke1616/ecmdb/commit/f113f487067afea3c1ad68b63e031a965b0bead4) order 创建工单，同步本地数据库，发送事件到Kafka
- [`3ff0dae`](https://github.com/Duke1616/ecmdb/commit/3ff0daeb252bda9c8641cf232d0bb7224d3228ca) order 模块
- [`2cbaaa9`](https://github.com/Duke1616/ecmdb/commit/2cbaaa906a4c2ad91f7e2792fe11d452c2a44cfa) 联调前端，创建工单，对接后端流程引擎
- [`fd61893`](https://github.com/Duke1616/ecmdb/commit/fd618938fabc738bbe857e5f38d8f5f718be36a6) elog 日志打印问题
- [`022f91a`](https://github.com/Duke1616/ecmdb/commit/022f91aa8f92261fc3dd3e74805f7791b8a36d18) todo order
- [`3e362c0`](https://github.com/Duke1616/ecmdb/commit/3e362c0ded4cdac2358dc7f7870280903f503fd4) 调整目录结构，流程引擎 engine
- [`a7f0a54`](https://github.com/Duke1616/ecmdb/commit/a7f0a5433f0697eef131a50ff9dc95ae314f4da9) 调整目录组织，解决循环引用，通过Kafka解偶 easyflow event调用order修改状态
- [`d90fbf7`](https://github.com/Duke1616/ecmdb/commit/d90fbf778dfd060a8a7eef867d639118168ceeee) 流程引擎代码，抽象 Instance 统一展示
- [`0506730`](https://github.com/Duke1616/ecmdb/commit/0506730493b8e324f543dfd09571c1e0bad0b452) pass 流程
- [`0d32bb4`](https://github.com/Duke1616/ecmdb/commit/0d32bb48a574234e41027cc3f06f7b3282dca792) 并行网关
- [`13f2fbd`](https://github.com/Duke1616/ecmdb/commit/13f2fbd91eebaa2a59a8a1931bb4d3c4f9e89dce) 包容网关、并行网关
- [`bb370fa`](https://github.com/Duke1616/ecmdb/commit/bb370faf548a11995056ad7f44f89e0a00797bf9) 审批记录
- [`28c9793`](https://github.com/Duke1616/ecmdb/commit/28c979309eff53c996b922b89867e916519eb0b9) 改写获取我的工单列表
- [`16ff3b7`](https://github.com/Duke1616/ecmdb/commit/16ff3b74ed9fd78264f03c0f8bb126fedc51cdbd) 前端 el-table 动态 合并单元格
- [`3eb0465`](https://github.com/Duke1616/ecmdb/commit/3eb0465486d8364c374cc585f5b078718d679641) engine 拆分 event 为独立模块
- [`fad25e8`](https://github.com/Duke1616/ecmdb/commit/fad25e8318d11e12da0ab46f3e5b4aac9f36af7e) wire 注解
- [`07b5e29`](https://github.com/Duke1616/ecmdb/commit/07b5e2971b55be921a4377248965e166007f9fcc) 联动任务模块
- [`b0a582d`](https://github.com/Duke1616/ecmdb/commit/b0a582d69deede20cd1433b3ee89169406623adc) 自动化执行
- [`81bf74d`](https://github.com/Duke1616/ecmdb/commit/81bf74d816f25de7129590bec884fff2e7629f03) 处理任务执行结果
- [`66f2fc3`](https://github.com/Duke1616/ecmdb/commit/66f2fc3def11676a37a22fc46f00eb02cfb4dc00) 流程图展示
- [`fbfdf68`](https://github.com/Duke1616/ecmdb/commit/fbfdf68f60343c472af6218fc777b2f8531d47c3) 任务历史
- [`d434b18`](https://github.com/Duke1616/ecmdb/commit/d434b18b4cf7792d0289a9509564e0bf1c333ab2) Args 传递参数
- [`9a50f16`](https://github.com/Duke1616/ecmdb/commit/9a50f1626a04f12834d43cd2e1d0c3c976f6efd8) 增强任务模块，支持重试、修改参数，runner模块新增环境变量
- [`ac2cbc2`](https://github.com/Duke1616/ecmdb/commit/ac2cbc216d8b7b88fe1f275c0310564cbd2efa9b) 完成基本自动化功能、支持变量
- [`dab8e85`](https://github.com/Duke1616/ecmdb/commit/dab8e856a90fb566c3f5da7729d175f160b5e5d2) task 定时任务
- [`72bc3f2`](https://github.com/Duke1616/ecmdb/commit/72bc3f233759b1eea4ec739cc00b30523afe7170) 定时任务改为 goroutine 启用
- [`3c327f4`](https://github.com/Duke1616/ecmdb/commit/3c327f4543bac3d686cdd19c981dfee18c4db6ab) 新增关闭自动化流程任务定时任务，任务运行，重试机制
- [`84fc9ca`](https://github.com/Duke1616/ecmdb/commit/84fc9ca8082bf0763d9864aec0c3a5aa18110736) 定时任务启动配置
- [`973df3a`](https://github.com/Duke1616/ecmdb/commit/973df3ad06fe0f8a9e3a44063092d8345e34ba2d) 验证兼容 wechat 审批 OA
- [`172bb4e`](https://github.com/Duke1616/ecmdb/commit/172bb4e4369eb55e2da11a79d57a8193794dc4e2) 工单系统流程图

## [v1.1.0](https://github.com/Duke1616/ecmdb/releases/tag/v1.1.0) - 2024-06-07

- [`687089c`](https://github.com/Duke1616/ecmdb/commit/687089c552029dd54c8bf7b52a65b9262b35592a) fixbug: 左侧伸展方向无法展示
- [`8606786`](https://github.com/Duke1616/ecmdb/commit/8606786b4b511f75f744969b7b63f569ac416fc5) add: 新增属性安全模型
- [`20a6713`](https://github.com/Duke1616/ecmdb/commit/20a6713a3d998290387deea407b36e5698bd47ab) 全局搜索 secure 类型展示

## [v1.0.0](https://github.com/Duke1616/ecmdb/releases/tag/v1.0.0) - 2024-05-28

- [`3b2876c`](https://github.com/Duke1616/ecmdb/commit/3b2876ccb3bf3a9292aa5d155d96589faa46a439) Initial commit
- [`31e4b18`](https://github.com/Duke1616/ecmdb/commit/31e4b1804102fb246c9928bed61e5e2e10fb82aa) 初始化buf、taskfile
- [`cb47651`](https://github.com/Duke1616/ecmdb/commit/cb47651c02d1f983283e16262f2ced0c8335e329) 目录结构设计
- [`2e76ae6`](https://github.com/Duke1616/ecmdb/commit/2e76ae62490db46af19f294a61cbb3d80b42aefb) update
- [`a20dc46`](https://github.com/Duke1616/ecmdb/commit/a20dc46627960e4a4b757c5d94fd6aec31a7b794) 项目初始化
- [`ec614f6`](https://github.com/Duke1616/ecmdb/commit/ec614f6464c15aa89df98619c714162ac1249a7b) 项目初始化
- [`2e82c99`](https://github.com/Duke1616/ecmdb/commit/2e82c992b7bb620ed1665461354258f13e3739e7) CMDB 初始化
- [`280e59c`](https://github.com/Duke1616/ecmdb/commit/280e59cae37f4e7a4d818d29816193a15a2cf932) mongo 连接 探测
- [`cdfbbb7`](https://github.com/Duke1616/ecmdb/commit/cdfbbb7356351fe3690dd4428e416a390d6d7b57) mongo 自增ID
- [`9d0c4f6`](https://github.com/Duke1616/ecmdb/commit/9d0c4f68a6197f21ee4a0c0bfac53ffb4d78f42f) 创建模型逻辑
- [`3d242c7`](https://github.com/Duke1616/ecmdb/commit/3d242c7dfaeda0255364843966ecc2c664cfafdd) 初步设计 attribute
- [`0ef3b92`](https://github.com/Duke1616/ecmdb/commit/0ef3b9229eb7f9de400205d6ca3aedde4c625f42) Resource 基本设计
- [`c06120c`](https://github.com/Duke1616/ecmdb/commit/c06120ca3584867e5f076fefdfc011b6c7685b3d) Resouce 数据录入
- [`623426c`](https://github.com/Duke1616/ecmdb/commit/623426ca233b4865cfa35917ea2839e71e7323b3) resource
- [`3f71a3d`](https://github.com/Duke1616/ecmdb/commit/3f71a3d5863be68071bcbddf2d0464a12f08f45b) model list detail 逻辑
- [`4cd67bd`](https://github.com/Duke1616/ecmdb/commit/4cd67bd9ffda38946b7574e6be1b4a9c87c9f0cb) relation 关联关系
- [`44cc007`](https://github.com/Duke1616/ecmdb/commit/44cc007ed0fe475c4ce5c7fadc3cdf9f38865a1b) wire 依赖
- [`c865c74`](https://github.com/Duke1616/ecmdb/commit/c865c7430c1ec696083916587e5c5157cfe9fb62) 修改数据库存储结构
- [`50943fb`](https://github.com/Duke1616/ecmdb/commit/50943fb9e1558ae223bc67102fd606f42712798b) 修改数据库存储结构
- [`30b0093`](https://github.com/Duke1616/ecmdb/commit/30b00938e904b08a90b103b9607ffcec3d80887b) 待完成字段映射，查询Mongo
- [`0e9ef4c`](https://github.com/Duke1616/ecmdb/commit/0e9ef4c8d717565cf13c7e29478416cc8a594f65) resource 查询逻辑
- [`fc28510`](https://github.com/Duke1616/ecmdb/commit/fc285108dab7145a44a009672d572075c7364ba6) 优化resource 和 attribute 关联处理逻辑
- [`4e01c57`](https://github.com/Duke1616/ecmdb/commit/4e01c57dcbb2defbbdc3685134f84077a695a7f9) 封装 mongox 自增ID
- [`a8ec82f`](https://github.com/Duke1616/ecmdb/commit/a8ec82f2b3cc8a87d63c33ad798df70822304210) 资产关联关系
- [`e787bab`](https://github.com/Duke1616/ecmdb/commit/e787babcc429eca812b37495727c66b665aec4a7) gin context 封装
- [`eef90c0`](https://github.com/Duke1616/ecmdb/commit/eef90c0bdcc4b58972de116b47cf8114024a98ea) 条件查询
- [`f205018`](https://github.com/Duke1616/ecmdb/commit/f20501879eab99aab2c9da08da75ec2547349609) UniqueIdentifier => uid 修改统一命名，代表唯一标识
- [`8f892a7`](https://github.com/Duke1616/ecmdb/commit/8f892a70655d6f553ddb45e445b8347d742a5c02) 前端传递
- [`f8a52dc`](https://github.com/Duke1616/ecmdb/commit/f8a52dc7d221fe0782e58a0e76c68ad5b1b88a01) 获取关联resource数据
- [`e9ddd9a`](https://github.com/Duke1616/ecmdb/commit/e9ddd9a539a4952d67f3cd589089d305cebd52f1) 完善 ioc
- [`51fa4b2`](https://github.com/Duke1616/ecmdb/commit/51fa4b2f876a1bbae4c9489f2b3b898209964194) 新增å通过关联类型和模型UID，查询数资源数据
- [`3df6020`](https://github.com/Duke1616/ecmdb/commit/3df6020c714966bebd51e53c5a280c1436ef0929) 拆分relation为多个文件
- [`8bd5acd`](https://github.com/Duke1616/ecmdb/commit/8bd5acd7731bb34037dd4e30426ed2b02b542a20) realtion type
- [`268df1a`](https://github.com/Duke1616/ecmdb/commit/268df1a0a4bcf5293a733885f7fa00407894185c) 关联类型
- [`f622306`](https://github.com/Duke1616/ecmdb/commit/f622306907fa87c637fb41533fd75db4a961eea7) 模型拓补图
- [`e3915c1`](https://github.com/Duke1616/ecmdb/commit/e3915c151c3e03fe7b26a0f7842e2a3155d5c269) 模型拓补图
- [`14b0df5`](https://github.com/Duke1616/ecmdb/commit/14b0df5324297591ca45ffd9f7062953bb573e65) 封装LDAPX
- [`4fcb08b`](https://github.com/Duke1616/ecmdb/commit/4fcb08bafddd938ec1e0654a84540211c2666561) 用户登录逻辑
- [`d0c22d9`](https://github.com/Duke1616/ecmdb/commit/d0c22d9af3ef33582ad491a28d3a22adb696941b) 用户LDAP登录逻辑
- [`1bc168d`](https://github.com/Duke1616/ecmdb/commit/1bc168dca21bd516cfd7e910800c66469e2163f7) Session + Jwt 登录认证
- [`60878de`](https://github.com/Duke1616/ecmdb/commit/60878dedc89415b3235ef533f23f41e440950674) 继续完善 关联关系模块
- [`4e18851`](https://github.com/Duke1616/ecmdb/commit/4e18851783d7bcaa21b5a8d137360b800c05c712) 通过聚合，处理资源列表
- [`0a8c492`](https://github.com/Duke1616/ecmdb/commit/0a8c4924918e291e803cbcf9d055f4581fd2eecb) 修复聚合列表
- [`88963d6`](https://github.com/Duke1616/ecmdb/commit/88963d6fe0e8f7986d3832ba1013abca5c5ca215) 模型重构
- [`b831eac`](https://github.com/Duke1616/ecmdb/commit/b831eac0f02e960958821208775a1e00cb26e4a2) 资产列表
- [`526f58d`](https://github.com/Duke1616/ecmdb/commit/526f58dc5acb01115d8eafe7ecd110cede2ca49e) 字段添加 detail 方法
- [`45c6ccf`](https://github.com/Duke1616/ecmdb/commit/45c6ccf06001ec2bded9ce4c75288c7172f46ef5) 模型属性 列表
- [`cb30b9e`](https://github.com/Duke1616/ecmdb/commit/cb30b9e511c85026fb1adbc9433e86e12772ca39) 去除没必要的指针
- [`d7d94af`](https://github.com/Duke1616/ecmdb/commit/d7d94af14dc7aa4ef5969bb814f1350698e48500) 真的令人头大
- [`4cf0412`](https://github.com/Duke1616/ecmdb/commit/4cf041221519496e4be93330f762032803c11bec) relation resource 拓补图 资产标记
- [`3183352`](https://github.com/Duke1616/ecmdb/commit/3183352255739d423e6a54f9eb6d6148e3b4924c) update
- [`4a50b32`](https://github.com/Duke1616/ecmdb/commit/4a50b32d7bce903393dbdb175452c56ec788b80c) 重构 relation 创建逻辑
- [`86c7ba0`](https://github.com/Duke1616/ecmdb/commit/86c7ba08145a34cfa871c8d114b0f0ff72edfa5a) 查询可以æ关联的模型数据
- [`7ee2180`](https://github.com/Duke1616/ecmdb/commit/7ee2180f2218924edbb06e408062a090cba100f9) 解决循环引用
- [`1482273`](https://github.com/Duke1616/ecmdb/commit/14822737764265038cded2afc175535824d4fcb1) 测试完成 查询å以关联的节点
- [`3538cc6`](https://github.com/Duke1616/ecmdb/commit/3538cc6763ef426d756bc79087ef5314df3d663f) 计划封装 mongox
- [`e45e02c`](https://github.com/Duke1616/ecmdb/commit/e45e02c53225a75c31a28eb69bdec26a217b5a73) 计划封装 mongox
- [`103d44d`](https://github.com/Duke1616/ecmdb/commit/103d44d89520cbcd29d0e90f41ed667cb66dc349) 修改mongox 作为入参，方便编写测试
- [`ed77274`](https://github.com/Duke1616/ecmdb/commit/ed772747f4b5151aef89041eac22ec8a3eb6551a) e2e 测试
- [`19d7974`](https://github.com/Duke1616/ecmdb/commit/19d7974ff6e8fe543c96fca805c327b6eff06fb3) 创建资源，e2e测试
- [`382179d`](https://github.com/Duke1616/ecmdb/commit/382179df17dcdb6aacbdec828c225133f3454a37) 新增 e2e测试
- [`163c2a3`](https://github.com/Duke1616/ecmdb/commit/163c2a3639e43941905b6afb35877013237af875) 重构 searchattrubute 返回信息
- [`7d37d09`](https://github.com/Duke1616/ecmdb/commit/7d37d095213d462c0490a2bbdaef54e1c795e845) attribute 添加联合唯一索引
- [`2347767`](https://github.com/Duke1616/ecmdb/commit/2347767632b58255ea30d063ef9cfd81222f89da) 优化 attribute 模块，完善e2e测试
- [`97e93cf`](https://github.com/Duke1616/ecmdb/commit/97e93cf06dd03672167db922061a256548d7b9cb) 创建 开启mongo事务
- [`50fd743`](https://github.com/Duke1616/ecmdb/commit/50fd7433c243cce03b44bfd675d218e9db1804f2) 删除多余方法
- [`921d9a8`](https://github.com/Duke1616/ecmdb/commit/921d9a857dbf1ca2bd5ec334cd5e5e7e65a780b5) model 重构
- [`f3c0599`](https://github.com/Duke1616/ecmdb/commit/f3c0599c8656c5fa28a672940b62451f89726433) 循环引用，我吐了
- [`5cb21fd`](https://github.com/Duke1616/ecmdb/commit/5cb21fde53c12a035e6c594ef3200d171d650834) 解决循环引用
- [`64f4418`](https://github.com/Duke1616/ecmdb/commit/64f4418d11f118b8ed76d0bc25b5232de215ae73) 优化 模型关联关系
- [`f04e434`](https://github.com/Duke1616/ecmdb/commit/f04e4345577c58f8bf1ae7c4136eccd9d090c480) 关联类型 优化
- [`634b519`](https://github.com/Duke1616/ecmdb/commit/634b519755eb896e9d1f8fbf511a370df9b2b872) 优化rsource 模块
- [`11c22e7`](https://github.com/Duke1616/ecmdb/commit/11c22e7e5cd79b99b26f6af06039860bc944cdb6) 优化 resource
- [`b019d20`](https://github.com/Duke1616/ecmdb/commit/b019d20d709f20bfc1c9b11db71a2fc39095ea34) 优化
- [`2bf03ba`](https://github.com/Duke1616/ecmdb/commit/2bf03ba4ebe252935a5ddb9d52e1fae0f303eba2) realtion resource 优化完成
- [`de7705c`](https://github.com/Duke1616/ecmdb/commit/de7705c6508a2da65f0ae66060f7a4fc6b1a77ac) user 模块
- [`a723e24`](https://github.com/Duke1616/ecmdb/commit/a723e24d43590e8470dcacb555f0732a001627ad) relation e2e测试模版
- [`96ce2e3`](https://github.com/Duke1616/ecmdb/commit/96ce2e35fdc7e89a1dca17c64fa8bee2b0b35694) 结构体修改，启动错误
- [`79d3721`](https://github.com/Duke1616/ecmdb/commit/79d3721b13c89262aebdca75984823a3d945260b) 模型分组返回
- [`70cd366`](https://github.com/Duke1616/ecmdb/commit/70cd366776547032df934a4ad2cc6f42396d10d0) 去除事务操作
- [`3747bd5`](https://github.com/Duke1616/ecmdb/commit/3747bd5d0646104ec9d56029b163a99a072717f7) 自定义模型列
- [`6d0c93e`](https://github.com/Duke1616/ecmdb/commit/6d0c93e401eb9ab8f05b2e50a10c6de71ef231c0) 前后端联调，模块
- [`00d59f5`](https://github.com/Duke1616/ecmdb/commit/00d59f523e6eb9d7e63b82d4032f9d75d2f21af9) 模型模块、前端联调
- [`4e9a727`](https://github.com/Duke1616/ecmdb/commit/4e9a7279b4a5cd128440fe80a24b64c724ffaa5e) 对接前端
- [`7797167`](https://github.com/Duke1616/ecmdb/commit/77971676c318483dd84f11f3db09cc221e97041e) 新建关联
- [`7017757`](https://github.com/Duke1616/ecmdb/commit/70177571b8a260d93938a23079f5c7622a7a6139) 新增关联，SRC方向过滤查询bug
- [`479592d`](https://github.com/Duke1616/ecmdb/commit/479592d5675cc99ff886a6f0b024681ac3f454ff) 资产 关联信息展示
- [`09c6886`](https://github.com/Duke1616/ecmdb/commit/09c6886195cc1fc605dcbd46f58fcba260836cf6) 取消关联
- [`1501fc0`](https://github.com/Duke1616/ecmdb/commit/1501fc0411a4b4d683107cfc3b2a1600db27cf4c) 取消关联
- [`629e623`](https://github.com/Duke1616/ecmdb/commit/629e623401f273d3424e134b645eaf9f7290433b) 资产拓扑图
- [`06d9261`](https://github.com/Duke1616/ecmdb/commit/06d92617009caa0e8a13d78dfa16202bd7aa2056) 新增模型，初始化创建属性分组、属性字段
- [`e4216aa`](https://github.com/Duke1616/ecmdb/commit/e4216aaae7dd881eb93f5220e8e6f0ae5daa71f8) 模型属性分组
- [`fb4da5e`](https://github.com/Duke1616/ecmdb/commit/fb4da5e74ec4144bbf85463bcf95ea3df5ac2225) 准备测试整体功能性是否完善
- [`f111bed`](https://github.com/Duke1616/ecmdb/commit/f111bed476c053e58c659ae2abf5330a3e4b44bd) 添加 api 前缀
- [`1cea9e6`](https://github.com/Duke1616/ecmdb/commit/1cea9e645fd3a1326b478d8e47eb10e14e6ab23c) 修改 Dockerfile
- [`7aa20ea`](https://github.com/Duke1616/ecmdb/commit/7aa20ea1b1cc74e22efa47bf8dba08206d472368) 部署文件修改，挂载配置文件
- [`84a1171`](https://github.com/Duke1616/ecmdb/commit/84a1171a814dd01d399ad36c2c12fbe0671e4828) github active
- [`445be77`](https://github.com/Duke1616/ecmdb/commit/445be7725a7154682643b7809bc3daaad410f08c) 去除缓存
- [`53b4b80`](https://github.com/Duke1616/ecmdb/commit/53b4b80cea6e3b5201facfc40980ff09827249a8) update
- [`9a007fe`](https://github.com/Duke1616/ecmdb/commit/9a007fe1dda378efb32cd85f4b338a1938673735) 模型删除验证
- [`cbcf21e`](https://github.com/Duke1616/ecmdb/commit/cbcf21e489e9bdaa09c649e4b967396c4cba1a4f) 全局搜索功能
- [`e12424f`](https://github.com/Duke1616/ecmdb/commit/e12424fd4b8dacfa4ddb72e89a942f1d8511fb79) 新增关联，新增过滤条件
- [`253c2e2`](https://github.com/Duke1616/ecmdb/commit/253c2e28b18fcc41713f13e92520479ba812d763) fixbug： 计算count 传递filter输入错误
- [`13af8d4`](https://github.com/Duke1616/ecmdb/commit/13af8d4f438758e3d0427e36c7806d769f45be8c) 资产关联拖布图 left right 方向扩展