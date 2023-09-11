# sc-richard-exec
sc1.0 richard的练习项目
> ## ❗️❗️❗️ 开发前务必设置好私有仓库配置 ❗️❗️❗️
> ```bash
> go env -w GOPROXY="https://goproxy.duowan.com,direct"
> go env -w GOSUMDB="sum.golang.org gosum.duowan.com"
> go env -w GOPRIVATE="gopkg.inshopline.com,git.duowan.com"
> ```

## 基础结构

```
./                          # 项目根，同样是业务模块
│
├── .idea/                  # GoLand 代码配置
│   ├── codeStyles/         # 代码风格，在 Go 中基本只有缩进字符的配置，应该提交到 Git 中，以确保同一个项目中的缩进一致
│   ├── inspectionProfiles/ # 代码检查配置，例如命名检查等，应该提交到 Git 中
│   └── ...                 # 其它配置不建议提交到 Git 中
│
├── proto/                  # 协议模块：idl 及 idl 生成代码模块目录
│   ├── gen.sh              # 生成pb相关文件（示例的foo目录下的文件）的脚本。进入proto目录, 执行./gen.sh即可
│   └── go.mod              # 模块描述文件，此模块需要发布，必须使用完整模块路径
│
├── .../                    # 其它业务相关划分的子包（无强制要求）
│   └── ...
│
├── main.go                 # main 入口函数
│
├── .gitignore              # git提交时忽略的文件列表声明
├── build.sh                # cicd构建所需的脚本
├── go.mod                  # 业务模块描述文件
├── go.sum                  # 业务模块依赖校验文件，确保已引用的依赖不会被篡改，应该提交到 Git 中
├── go.work                 # Go 工作空间描述文件，方便引用项目内的 proto 模块
└── go.work.sum             # 工作空间依赖校验文件，确保已引用的依赖不会被篡改，应该提交到 Git 中
```

详细规范见：https://shopline.yuque.com/sczv0i/go/project-spec


## 生成代码

进入./proto目录, 执行 [`./gen.sh`](./proto/gen.sh)（详细逻辑见脚本实现）。注：必须先安装protoc-gen-go、protoc-gen-go-grpc
