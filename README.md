# go编码规范

# 项目结构

以下为一般项目结构，根据不同的项目类型和需求，可自由增删某些结构：

```                 
- cmd                        # 命令行程序命令
  - etc                      # 配置文件
  - data                     # 应用生成数据文件
  - log                      # 应用生成日志文件
- global                     # 通用文件
- dg                         # 数据采集层
- services                   # 接口服务层
- models                     # 数据逻辑层
- utils                      # 工具类
```

## 命令行应用

当应用类型为命令行应用时，需要将命令相关文件存放于 `/cmd` 目录下，并为每个命令创建一个单独的源文件：

```
/cmd
	main.go
	serve.go
	web.go
```


# 导入标准库、第三方或其它包
符合 ```go fmt``` 规范

除标准库外，Go 语言的导入路径基本上依赖代码托管平台上的 URL 路径，因此一个源文件需要导入的包有 4 种分类：标准库、第三方包、组织内其它包和当前包的子包。

基本规则：

- 如果同时存在 2 种及以上，则需要使用分区来导入。每个分类使用一个分区，采用空行作为分区之间的分割。
- 在非测试文件（`*_test.go`）中，禁止使用 `.` 来简化导入包的对象调用。
- 禁止使用相对路径导入（`./subpackage`），所有导入路径必须符合 `go get` 标准。

下面是一个完整的示例：

```Go
import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/go-openapi/spec"
	"github.com/robfig/cron"

	"etstone.cn/db/redis"
	"etstone.cn/logger"

	"etstone.cn/ic/dg"
	"etstone.cn/ic/global"
	"etstone.cn/ic/model"
	"etstone.cn/ic/service"
)
```

# 注释规范
符合 ```github.com/golang/lint/golint``` 规范

- 所有导出对象都需要注释说明其用途；非导出对象根据情况进行注释。
- 如果对象可数且无明确指定数量的情况下，一律使用单数形式和一般进行时描述；否则使用复数形式。
- 包、函数、方法和类型的注释说明都是一个完整的句子。
- 句子类型的注释首字母均需大写；短语类型的注释首字母需小写。
- 注释的单行长度不能超过 80 个字符。

### 包级别

- 包级别的注释就是对包的介绍，只需在同个包的任一源文件中说明即可有效。
- 对于 `main` 包，一般只有一行简短的注释用以说明包的用途，且以项目名称开头：

	```Go
	// Gogs (Go Git Service) is a painless self-hosted Git Service.
	package main
	```

- 对于一个复杂项目的子包，一般情况下不需要包级别注释，除非是代表某个特定功能的模块。
- 对于简单的非 `main` 包，也可用一行注释概括。
- 对于相对功能复杂的非 `main` 包，一般都会增加一些使用示例或基本说明，且以 `Package <name>` 开头：

	```Go
	/*
	Package regexp implements a simple library for regular expressions.

	The syntax of the regular expressions accepted is:

	    regexp:
	        concatenation { '|' concatenation }
	    concatenation:
	        { closure }
	    closure:
	        term [ '*' | '+' | '?' ]
	    term:
	        '^'
	        '$'
	        '.'
	        character
	        '[' [ '^' ] character-ranges ']'
	        '(' regexp ')'
	*/
	package regexp
	```

- 特别复杂的包说明，可单独创建 [`doc.go`](https://github.com/robfig/cron/blob/master/doc.go) 文件来加以说明。

### 结构、接口及其它类型

- 类型的定义一般都以单数形式描述：

	```Go
	// Request represents a request to run a command.
	type Request struct { ...
	```

- 如果为接口，以 `er` 作为后缀，接口的实现则去掉 `er`，则一般以以下形式描述：

	```Go
	// FileInfo is the interface that describes a file and is returned by Stat and Lstat.
	type Reader interface { ...
	```


### 函数与方法

- 函数与方法的注释需以函数或方法的名称作为开头：

	```Go
	// Post returns *BeegoHttpRequest with POST method.
	```

- 如果一句话不足以说明全部问题，则可换行继续进行更加细致的描述：

	```Go
	// Copy copies file from source to target path.
	// It returns false and error when error occurs in underlying function calls.
	```

- 若函数或方法为判断类型（返回值主要为 `bool` 类型），则以 `<name> returns true if` 开头：

	```Go
	// HasPrefix returns true if name has any string in given slice as prefix.
	func HasPrefix(name string, prefixes []string) bool { ...
	```

### 其它说明


- 当某个部分等待完成时，可用 `TODO:` 开头的注释来提醒维护人员。
- 当某个部分存在已知问题进行需要修复或改进时，可用 `FIXME:` 开头的注释来提醒维护人员。
- 当需要特别说明某个问题时，可用 `NOTE:` 开头的注释：

	```Go
	// NOTE: os.Chmod and os.Chtimes don't recognize symbolic link,
	// which will lead "no such file or directory" error.
	return os.Symlink(target, dest)
	```

# 命名规则

### 文件名

- 整个应用或包的主入口文件应当是 `main.go` 或与应用名称简写相同。例如：`Gogs` 的主入口文件名为 `gogs.go`。普通文件命名应当是全部小写 `context.go` 如复杂文件名应当是使用下划线分词 `file_reader.go`

### 函数或方法

- 若函数或方法为判断类型（返回值主要为 `bool` 类型），则名称应以 `Has`, `Is` 等判断性动词开头：

	```go
	func HasPrefix(name string, prefixes []string) bool { ... }
	func IsEntry(name string, entries []string) bool { ... }
	```

### 常量

- 常量均需使用全部大写字母组成，并使用下划线分词：

	```go
	const APP_VER = "0.7.0.1110 Beta"
	```

- 如果是枚举类型的常量，需要先创建相应类型：

	```go
	type Scheme string

	const (
		HTTP  Scheme = "http"
		HTTPS Scheme = "https"
	)
	```

- 如果模块的功能较为复杂、常量名称容易混淆的情况下，为了更好地区分枚举类型，可以使用完整的前缀：

	```go
	type PullRequestStatus int

	const (
		PULL_REQUEST_STATUS_CONFLICT PullRequestStatus = iota
		PULL_REQUEST_STATUS_CHECKING
		PULL_REQUEST_STATUS_MERGEABLE
	)
	```

### 变量

- 变量命名基本上遵循相应的英文表达或简写。
- 在相对简单的环境（对象数量少、针对性强）中，可以将一些名称由完整单词简写为单个字母，例如：
	- `user` 可以简写为 `u`
	- `userID` 可以简写 `uid`
- 若变量类型为 `bool` 类型，则名称应以 `Has`, `Is`  开头：

	```go
	var isExist bool
	var hasConflict bool
	```
	
- 上条规则也适用于结构定义：

	```go
	// Webhook represents a web hook object.
	type Webhook struct {
		ID           int64 `gorm:"id"`
		RepoID       int64
		OrgID        int64
		URL          string `gorm:"url TEXT"`
		ContentType  HookContentType
		Secret       string `gorm:"TEXT"`
		Events       string `gorm:"TEXT"`
		*HookEvent   `gorm:"-"`
		IsSSL        bool `gorm:"is_ssl"`
		IsActive     bool
		HookTaskType HookTaskType
		Meta         string     `gorm:"TEXT"` // store hook-specific attributes
		LastStatus   HookStatus // Last delivery status
		Created      time.Time  `gorm:"CREATED"`
		Updated      time.Time  `gorm:"UPDATED"`
	}
	```

#### 变量命名惯例

变量名称一般遵循驼峰法，但遇到特有名词时，需要遵循以下规则：

- 如果变量为私有，且特有名词为首个单词，则使用小写，如 `apiClient`。
- 其它情况都应当使用该名词原有的写法，如 `APIClient`、`repoID`、`UserID`。

下面列举了一些常见的特有名词：

```go
// A GonicMapper that contains a list of common initialisms taken from golang/lint
var LintGonicMapper = GonicMapper{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
	"XSRF":  true,
	"XSS":   true,
}
```

# 声明语句

### 函数或方法

函数或方法的参数排列顺序遵循以下几点原则（从左到右）：

1. 参数的重要程度与逻辑顺序
2. 简单类型优先于复杂类型
3. 尽可能将同种类型的参数放在相邻位置，则只需写一次类型

#### 示例

以下声明语句，`User` 类型要复杂于 `string` 类型，但由于 `Repository` 是 `User` 的附属品，首先确定 `User` 才能继而确定 `Repository`。因此，`User` 的顺序要优先于 `repoName`。

```Go
func IsRepositoryExist(user *User, repoName string) (bool, error) { ...
```

# 代码指导

### 基本约束

- 所有应用的 `main` 包需要有 `APP_VER` 常量表示版本，格式为 `X.Y.Z.Date [Status]`，例如：`0.7.6.1112 Beta`。
- 单独的库需要有函数 `Version` 返回库版本号的字符串，格式为 `X.Y.Z[.Date]`。
- 当单行代码超过 80 个字符时，就要考虑分行。分行的规则是以参数为单位将从较长的参数开始换行，以此类推直到每行长度合适：

	```go
	So(z.ExtractTo(
		path.Join(os.TempDir(), "testdata/test2"),
		"dir/", "dir/bar", "readonly"), ShouldBeNil)
	```

- 当单行声明语句超过 80 个字符时，就要考虑分行。分行的规则是将参数按类型分组，紧接着的声明语句的是一个空行，以便和函数体区别：

	```go
	// NewNode initializes and returns a new Node representation.
	func NewNode(
		importPath, downloadUrl string,
		tp RevisionType, val string,
		isGetDeps bool) *Node {

		n := &Node{
			Pkg: Pkg{
				ImportPath: importPath,
				RootPath:   GetRootPath(importPath),
				Type:       tp,
				Value:      val,
			},
			DownloadURL: downloadUrl,
			IsGetDeps:   isGetDeps,
		}
		n.InstallPath = path.Join(setting.InstallRepoPath, n.RootPath) + n.ValSuffix()
		return n
	}
	```

- 定义对象函数时,指针用p：
	```go
	func (p *Node)NewNode(){}
	```

- 分组声明一般需要按照功能来区分，而不是将所有类型都分在一组：

	```go
	const (
		// Default section name.
		DEFAULT_SECTION = "DEFAULT"
		// Maximum allowed depth when recursively substituing variable names.
		_DEPTH_VALUES = 200
	)

	type ParseError int

	const (
		ERR_SECTION_NOT_FOUND ParseError = iota + 1
		ERR_KEY_NOT_FOUND
		ERR_BLANK_SECTION_NAME
		ERR_COULD_NOT_PARSE
	)
	```

- 当一个源文件中存在多个相对独立的部分时，为方便区分，需使用由 [ASCII Generator](http://www.network-science.de/ascii/) 提供的句型字符标注（示例：`Comment`）：

	```go
	// _________                                       __
	// \_   ___ \  ____   _____   _____   ____   _____/  |_
	// /    \  \/ /  _ \ /     \ /     \_/ __ \ /    \   __\
	// \     \___(  <_> )  Y Y  \  Y Y  \  ___/|   |  \  |
	//  \______  /\____/|__|_|  /__|_|  /\___  >___|  /__|
	//         \/             \/      \/     \/     \/
	```

- 函数或方法的顺序一般需要按照依赖关系由浅入深由上至下排序，即最底层的函数出现在最前面。例如，下方的代码，函数 `ExecCmdDirBytes` 属于最底层的函数，它被 `ExecCmdDir` 函数调用，而 `ExecCmdDir` 又被 `ExecCmd` 调用：

	```go
	// ExecCmdDirBytes executes system command in given directory
	// and return stdout, stderr in bytes type, along with possible error.
	func ExecCmdDirBytes(dir, cmdName string, args ...string) ([]byte, []byte, error) {
		...
	}

	// ExecCmdDir executes system command in given directory
	// and return stdout, stderr in string type, along with possible error.
	func ExecCmdDir(dir, cmdName string, args ...string) (string, string, error) {
		bufOut, bufErr, err := ExecCmdDirBytes(dir, cmdName, args...)
		return string(bufOut), string(bufErr), err
	}

	// ExecCmd executes system command
	// and return stdout, stderr in string type, along with possible error.
	func ExecCmd(cmdName string, args ...string) (string, string, error) {
		return ExecCmdDir("", cmdName, args...)
	}
	```

- 结构附带的方法应置于结构定义之后，按照所对应操作的字段顺序摆放方法：

	```go
	type Webhook struct { ... }
	func (w *Webhook) GetEvent() { ... }
	func (w *Webhook) SaveEvent() error { ... }
	func (w *Webhook) HasPushEvent() bool { ... }
	```

- 如果一个结构拥有对应操作函数，大体上按照 `CRUD` 的顺序放置结构定义之后：

	```go
	func CreateWebhook(w *Webhook) error { ... }
	func GetWebhookById(hookId int64) (*Webhook, error) { ... }
	func UpdateWebhook(w *Webhook) error { ... }
	func DeleteWebhook(hookId int64) error { ... }
	```

- 如果一个结构拥有以 `Has`、`Is` 开头的函数或方法，则应将它们至于所有其它函数及方法之前；这些函数或方法以 `Has`、`Is` 的顺序排序。
- 变量的定义要放置在相关函数之前：

	```go
	var CmdDump = cli.Command{
		Name:  "dump",
		...
		Action: runDump,
		Flags:  []cli.Flag{},
	}

	func runDump(*cli.Context) { ...
	```

- 在初始化结构时，尽可能使用一一对应方式：

	```go
	AddHookTask(&HookTask{
		Type:        HTT_WEBHOOK,
		Url:         w.Url,
		Payload:     p,
		ContentType: w.ContentType,
		IsSsl:       w.IsSsl,
	})
	```

# 测试用例

- 单元测试都必须使用 [GoConvey](http://goconvey.co/) 编写，且辅助包覆盖率必须在 80% 以上。

### 使用示例

- 为辅助包书写使用示例的时，文件名均命名为 `example_test.go`。
- 测试用例的函数名称必须以 `Test_` 开头，例如：`Test_Logger`。
- 如果为方法书写测试用例，则需要以 `Text_<Struct>_<Method>` 的形式命名，例如：`Test_Macaron_Run`。