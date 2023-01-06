# 应急响应日志收集工具

目前收集的内容如下：

Account

Arp

DnsCatch

Drive

EventLog（只收集Security.evtx以及System.evtx）

FileList（只收集C盘两层目录以及用户temp目录，要是想要收集全目录，可以自己简单改下）

FileSensitiveDir（只收集了recent和后面的recent完全一致，没想好要收集哪些目录）

History

Hosts

Kb

Networks

Others（minidump了一些东西）

Pipe

Process

Program

Regedit（注册表只收集了启动部分）

Route

Schtasks

Services

Share

StartUp

SysStartUp

UserTemp

WMIObject

SystemInfo

收集内容为json格式，启用了AES加密，加密函数再util里，可取消。

希望大家多多支持，点点赞

有bug、需求或想法随时提issue，感谢
