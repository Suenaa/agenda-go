# Agenda

## 需求描述
- 业务需求：见附件
- 功能需求：设计一组命令完成agenda的管理，例如：
    + agenda help ：列出命令说明
    + agenda register -uUserName --password pass -email=a@xxx.com：注册用户
    + agenda help register ：列出 register 命令的描述
    + agenda cm ... : 创建一个会议
    + 原则上一个命令对应一个业务功能

- 持久化要求
    + 使用 json 存储 User 和 Meeting 实体
    + 当前用户信息存储在 curUser.txt 中

- 开发需求
    + 团队：2-3人，一人作为 master 创建程序框架，其他人 fork 该项目，所有人同时开发。团队 不能少于 2 人
    + 时间：两周完成

-项目目录
    + cmd ：存放命令实现代码
    + entity ：存放 User 和 Meeting 对象读写与处理逻辑
    + 其他目录 ： 自由添加

- 日志服务
    + 使用 log 包记录命令执行情况

## 第一周任务
    1.按 3.3 安装 cobra 并完成小案例
    2.按需求设计 agenda 的命令与参数（制品 cmd-design.md）
    3.master 创建项目，提交到 github， 其他人 fork 该项目
    4.每人分别创建属于自己的命令（命令实现 Print 读取的参数即可），提交并归并。确保不同人管理不同文件，以便于协作
    5.如时间富余，请完成 User 和 Meeting 实体 json 文件读写

### 思考

    假设你要建立项目 testcobra，你必须建立在 $GOPATH\src\$your_git_account\$repo\testcobra 目录下，才能正常实现协作。采用 golang 术语，$GOPATH\src\$your_git_account\$repo\ 是你当前的 工作区 

    问题：如果你的项目不能公开代码，甚至不能使用 github，如何实现多人协作呢？


## 第二周任务
    1.在项目中添加 .travis.yml 文件，并添加测试程序。让你的项目“持续集成” -- “CI” 了！
    2.添加 log 服务，记录用户的操作过程，以及关键的输出
    3.约定 entity 和 cmd 之间的接口服务，实现 agenda，并在 README.md 文件中给出简要使用说明和测试结果
    4.（如果你有兴趣），使用 pflag 包，自己实现一个简版 Command.go 取代 cobra。必须使用组合设计模式

### 思考
    为什么 cobra 的设计，让我们构建命令行程序变的如此简单？如何持续提升我们的设计能力。


## 附件
### 用户注册

    1. 注册新用户时，用户需设置一个唯一的用户名和一个密码。另外，还需登记邮箱及电话信息。
    2.如果注册时提供的用户名已由其他用户使用，应反馈一个适当的出错信息；成功注册后，亦应反馈一个成功注册的信息。

### 用户登录

    1.用户使用用户名和密码登录 Agenda 系统。
    2.用户名和密码同时正确则登录成功并反馈一个成功登录的信息。否则，登录失败并反馈一个失败登录的信息。

### 用户登出

    1. 已登录的用户登出系统后，只能使用用户注册和用户登录功能。

### 用户查询

    1.已登录的用户可以查看已注册的所有用户的用户名、邮箱及电话信息。

### 用户删除

    1.已登录的用户可以删除本用户账户（即销号）。
    2.操作成功，需反馈一个成功注销的信息；否则，反馈一个失败注销的信息。
    3.删除成功则退出系统登录状态。删除后，该用户账户不再存在。
    4.用户账户删除以后：
        + 以该用户为 发起者 的会议将被删除
        + 以该用户为 参与者 的会议将从 参与者 列表中移除该用户。若因此造成会议 参与者 人数为0，则会议也将被删除。

### 创建会议

    1.已登录的用户可以添加一个新会议到其议程安排中。会议可以在多个已注册
用户间举行，不允许包含未注册用户。添加会议时提供的信息应包括：
        + 会议主题(title)（在会议列表中具有唯一性）
        + 会议参与者(participator)
        + 会议起始时间(start time)
        + 会议结束时间(end time)
    2.注意，任何用户都无法分身参加多个会议。如果用户已有的会议安排（作为发起者或参与者）与将要创建的会议在时间上重叠 （允许仅有端点重叠的情况），则无法创建该会议。
    3.用户应获得适当的反馈信息，以便得知是成功地创建了新会议，还是在创建过程中出现了某些错误。


### 增删会议参与者

    1.已登录的用户可以向 自己发起的某一会议增加/删除 参与者 。
    2.增加参与者时需要做 时间重叠 判断（允许仅有端点重叠的情况）。
    3.删除会议参与者后，若因此造成会议 参与者 人数为0，则会议也将被删除。

### 查询会议

    1.已登录的用户可以查询自己的议程在某一时间段(time interval)内的所有会议安排。
    2.用户给出所关注时间段的起始时间和终止时间，返回该用户议程中在指定时间范围内找到的所有会议安排的列表。
    3.在列表中给出每一会议的起始时间、终止时间、主题、以及发起者和参与者。
    4.注意，查询会议的结果应包括用户作为 发起者或参与者 的会议。

### 取消会议

    1.已登录的用户可以取消 自己发起 的某一会议安排。
    2.取消会议时，需提供唯一标识：会议主题（title）。

### 退出会议

    1.已登录的用户可以退出 自己参与 的某一会议安排。
    2.退出会议时，需提供一个唯一标识：会议主题（title）。若因此造成会议 参与者 人数为0，则会议也将被删除。

### 清空会议

    1.已登录的用户可以清空 自己发起 的所有会议安排。