## heic2jpg
将HEIC格式图片批量转化为常见的JPG格式(保留EXIF信息)  

Golang实现，支持多系统平台，单文件轻量无依赖  

## 使用方法(windows为例)
* 直接将任意数量*.heic文件或者heic文件所在目录拖放到程序图标上即可  
* CLI下支持目录参数(扫描处理指定目录下所有heic格式文件)如`heic2jpg.exe ./heicImgs`，也支持多文件参数如`heic2jpg.exe 2.heic 43.heic` 

## 编译  
注意如果带`-H=windowsgui`参数编译，程序直接运行不显示终端窗口  
`go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -ldflags -H=windowsgui`

## Reference:
https://gophercoding.com/convert-heic-to-jpeg-go/
