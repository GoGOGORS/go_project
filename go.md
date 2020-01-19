

# GoLang

### 1.安装：

**GOROOT	指向go目录下的bin目录**

**GOPATH  	指向以后存放go项目的目录**

### 2.go的执行流程分析：

go env 查看go相关的环境变量

go run   打包编译直接运行，不会生成可执行文件

go buid   打包编译生成可执行文件，再运行可执行文件

go build -o test.exe  指定生成的可执行文件名

**区别：go run指令依赖go语言的开发环境运行，迁移性不高；go build生成的可执行文件是在打包编译的时候把所需依赖加载进了可执行文件中，所以可执行文件会大于编译之前的，可在没有go语言环境的其他机器上面运行。**

### 3.golang的注释：

```go
	// fmt.Println("hello go")	单行注释

	/*
	fmt.Println("hello go")		多行注释/块注释
	fmt.Println("hello go")	
	fmt.Println("hello go")	
    */
```

