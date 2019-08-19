## update-rc.d: startup manager
* 用来更新系统启动项的脚本，脚本的链接位于/etc/rcN.d/目录，对应脚本位于/etc/init.d/目录
### Linux启动步骤: 
    * (1)读取 MBR 的信息，启动 Boot Manager
    * (2)加载系统内核，启动init根进程: 读取/etc/inittab进入预设的运行级别;
    * (3)init进程先后运行/etc/rcS.d/ 目录和/etc/rcN.d/目录下的启动脚本;
    * (4)启动登录管理器，等待用户登录;
### 安装
* init [X]命令(0-关闭,1-单用户root, 2~5-多用户，6-重启);
* sudo install sysv-rc-conf :安装
* sudo sysv-rc-conf :启动
### 常用命令
* update-rc.d -f ＜basename＞ remove :从所有的运行级别中删除指定启动项
* update-rc.d ＜basename＞ start|stop ＜order＞ ＜runlevels＞ :按指定顺序、在指定运行级别中启动或关闭


## systemd 
* systemctl list-unit-files --type=service :查看自启动服务
* journalctl -b -N :命令可以重现你倒数第N次启动时候的信息
* systemd-analyze blame  :可以显示进程耗时
