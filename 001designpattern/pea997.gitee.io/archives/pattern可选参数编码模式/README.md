李文周 https://www.liwenzhou.com/posts/Go/functional-options-pattern/



## 不使用功能选项模式
​ 在设计 API 时，可能会出现一个问题：我们如何处理可选配置？ 有效地解决这个问题可以提高我们的 API 的便利性。 本节将通过一个具体示例介绍处理可选配置的不同方法。

​ 假设我们必须设计一个库来公开一个函数用于创建HTTP服务器。此函数将接受不同的输入：地址和端口，下面显示了函数的骨架
```go
func NewServer(addr string, port int) (*http.Server, error) { 
  // ...
}
```
​ 我们库的客户已经开始使用这个功能了，大家都很高兴。 但在某些时候，我们的客户开始抱怨这个功能有些受限并且缺少其他参数（例如，写入超时和连接上下文）。 但是，我们注意到添加新的函数参数会破坏兼容性，迫使客户端修改调用 NewServer 的方式。 同时，我们希望通过这种方式丰富端口管理相关的逻辑
```
如果端口没有被设置，我们将使用默认端口
如果这个端口无效，那么我们会返回错误
如果这个端口为0，那么我们会使用随机端口
否则，我们将使用客户提供的端口
```

​ 我们如何以 API 友好的方式实现此功能？ 让我们看看不同的选项。

### Config Struct
​ 因为 Go 不支持函数签名中的可选参数，第一种可能的方法是使用配置结构来传达什么是必需的，什么是可选的。 例如，强制参数可以作为函数参数存在，而可选参数可以在 Config 结构中处理:
```go
type Config struct {
   Port        int
}

func NewServer(addr string, cfg Config) {
  
}
```
​ 此解决方案修复了兼容性问题。 事实上，如果我们添加新的选项，它不会在客户端中断。 但是，这种方法并没有解决我们与端口管理相关的需求。 事实上，我们应该记住，如果没有提供结构字段，它会被初始化为零值
```
int类型的零值是0
float类型的零值是0.0
string类型的零值是””
Nil会作为，slice，map，channel，指针,接口和function的零值
```
​ 因此，在下面的例子中，两个结构是相等的：
```go
c1 := httplib.Config {
  Port: 0,
}

c2 := http.lib.Config {
  
}
```
​ 在我们的示例中，我们需要找到一个方法来区分是用户设置为0还是没有进行端口的设置。有一种选择是将配置结构中的所有参数作为指针处理
```
type Config struct {
  Port *int
}
```
​ 使用整数指针，在语义上，我们可以突出值 0 和缺失值（nil 指针）之间的区别。

​ 这个选项可行，但它有几个缺点。 首先，客户端提供整数指针并不方便。 客户端必须创建一个变量，然后以这种方式传递一个指针
```go
port := 0
config := httplib.Config{
  Port: &port, // 使用integer指针
}
```
​ 它本身并不是一个亮点，但整个 API 使用起来不太方便。 此外，我们添加的选项越多，代码就会变得越复杂。
第二个缺点是客户端使用我们的默认配置库将需要以这种方式传递一个空结构
```
httplib.NewServer("localhost", httplib.Config{})
```
​ 这段代码看起来不太好。 读者将不得不理解这个神奇结构的含义

​ 另一种选择是使用经典的构建器模式，如下一节所述

## 建造者模式
​ 构建器模式最初是四人组设计模式的一部分，为各种对象创建问题提供了灵活的解决方案。 Config 的构造与结构本身是分开的。 它需要一个额外的结构体 ConfigBuilder，它接收配置和构建 Config 的方法

​ 让我们看一个具体的例子，看看它如何帮助我们设计一个友好的 API 来满足我们所有的需求，包括端口管理
```go
type Config struct {
  Port int
}

type ConfigBuilder struct{
  port *int
}

func (b *ConfigBuilder) Port(port int) *ConfigBuilder{
  b.port = &port
  return b 
}

func (b *ConfigBuilder) Build() (Config, error) {
   cfg := Config{}
  
  if b.port == nil {
     cfg.Port = defaultHTTPPort
  }else {
    if *b.port == 0 {
        cfg.Port = randomPort()
    } else if *b.port < 0 {
        return Config{}, errors.New("port should be positive")
    } else {
        cfg.Port = *b.port
    }
  }
  return cfg,nil
}


func NewServer(addr string, config Config) (*http.Server, error) { 
  // ...
}
```
​ ConfigBuilder 结构保存客户端配置。 它公开了一个 Port 方法来设置端口。 通常，这样的配置方法会返回构建器本身，以便我们可以使用方法链接（例如，builder.Foo(“foo”).Bar(“bar”)）。 它还公开了一个 Build 方法，该方法包含初始化端口值的逻辑（指针是否为 nil 等），并在创建后返回一个 Config 结构

​ 然后，客户端将按以下方式使用我们基于构建器的 API（我们假设我们已将代码放在 httplib 包中
```go
builder :=httplib.ConfigBuilder{}
builder.Port(8080)
cfg, err := builder.Build()
if err != nil{
  return err 
}

server, err := httplib.NewServer("localhost", cfg)
if err != nil {
	return err
}
```
​ 首先，客户端创建一个 ConfigBuilder 并使用它来设置一个可选字段，例如端口。 然后，它调用 Build 方法并检查错误。 如果成功，配置将传递给 NewServer

​ 这种方法使端口管理更加方便。 不需要传递整数指针，因为 Port 方法接受整数。 但是，如果客户端想要使用默认配置，我们仍然需要传递一个可以为空的配置结构
```go
server, err := httplib.NewServer("localhost", nil)
```
​ 在某些情况下，另一个缺点与错误管理有关。 在会抛出异常的编程语言中，如果输入无效，构建器方法（例如 Port）会引发异常。 如果我们想保持链接调用的能力，函数就不能返回错误。

​ 因此，我们必须延迟Build方法中的验证，如果客户端可以传递多个选项，但我们想要精确处理端口无效的情况，这会使错误处理更加复杂

​ 现在让我们看看另一种称为函数选项模式的方法，它依赖于可变参数

## 功能选项模式
​ 我们要讨论的最后一种方法是功能选项模式。 尽管有不同的实现方式，变化很小，但主要思想如下
```
未导出的结构包含配置：options
每个options都是一个返回相同类型的函数：type Option func(options *options) error 。例如：WithPort 接受一个int类型的参数，并返回一个Option的结构类型
```
​ 这是options结构、options类型和 WithPort 选项的 Go 实现
```go
type Options struct {
  port *int
}

type Option func (options *options) error 

func WithPort(port int) Option {
  return func (options *options) error {
    if port < 0 {
      return errors.New("port should be positive")
    }
    options.port = &port
    return nil
  }
}
```
​ 这里，WithPort 返回一个闭包。 闭包是一个匿名函数，它从其主体外部引用变量； 在这种情况下，端口变量。 闭包遵循 Option 类型并实现端口验证逻辑。 每个配置字段都需要创建一个包含类似逻辑的公共函数（按照惯例以 With 前缀开头）：在需要时验证输入并更新配置结构。

让我们看一下提供者端的最后一部分：NewServer 实现。 我们会将选项作为可变参数传递。 因此，我们必须迭代这些选项来改变选项配置结构：
```go
func NewServer(addr string ，opts ...Options) (*http.Server,error){
  var options Options
  for _ , opt :=range opts {
    err := opt(&options)
    if err!=nil{
      return nil,err
    }
  }
 	
  var port int 
  if options.port == nil {
    port = defaultHTTPPort
  } else {
    if *options.port == 0 {
      port = randomPort()
    }else{
      port = *options.port
    } 
  }
  // ...
}
```
​ 我们首先创建一个空的Options选项结构。然后，我们便利每一个Option参数并执行他们以改变Option结构（option类型是一个函数)。一旦构建了选项结构，我们就可以实现有关端口管理的最终逻辑

​ 因为 NewServer 接受可变选项参数，所以客户端现在可以通过在强制地址参数后传递多个选项来调用此 API，例如：
```go
server, err := httplib.NewServer("localhost",
        httplib.WithPort(8080),
        httplib.WithTimeout(time.Second))
```        
​ 但是，如果客户端需要默认配置，则不必提供参数（例如，一个空结构，正如我们在前面的方法中看到的那样)。
```
server, err := httplib.NewServer("localhost")
```
​ 此模式是功能选项模式。 它提供了一种方便且 API 友好的方式来处理选项。 尽管构建器模式可以是一个有效的选项，但它有一些小的缺点，这些缺点往往使功能选项模式成为 Go 中处理此问题的惯用方法。 我们还要注意，这种模式用于不同的 Go 库，例如 gRPC