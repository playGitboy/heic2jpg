## heic2jpg
将从ios中导出的HEIC格式图片批量转化为常见的JPG格式(保留EXIF信息)  

生成一个可以用rundll32运行调用的dll文件——配合注册表项修改实现在.heic文件上右键新增功能菜单调用转换功能的效果

## Windows下编译  
`go build -gcflags=-trimpath=$GOPATH -asmflags=-trimpath=$GOPATH -ldflags "-w -s" -buildmode=c-shared -o Heic2Jpg.dll`  
编译成功生成Heic2Jpg.dll文件  

## 在heic文件新增右键菜单项  
将以下内容保存为a.reg双击打开导入注册表(注意对应修改Heic2Jpg.dll路径)，在任意*.heic文件右键选择"Heic2Jpg"菜单即可自动调用转换功能——转换后会在heic文件同目录生成同名jpg文件（支持批量选择处理）  
`reg
Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\SystemFileAssociations\.heic\Shell\Heic2Jpg\command]
@="rundll32.exe D:\\Tool\\Heic2Jpg.dll,Heic2Jpg %1"
`

## Reference:
https://www.cnblogs.com/gatling/p/17218785.html
