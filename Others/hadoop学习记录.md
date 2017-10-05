---
title: Hadoop学习记录
---


#　Hadoop学习记录


## Hadoop安装配置(基于ubuntu:14.04)

### Hadoop单机版安装配置
* 创建hadoop用户：
    * sudo useradd -m hadoop -s /bin/bash    : 用户名:hadoop;   执行shell:bin/bash
    * sudo passwd hadoop      :设置密码
    * sudo adduser hadoop sudo  :添加管理员权限
    * su - hadoop 切换为hadoop用户    
* 安装其他必要软件
    * sudo apt-get update                                  #更新软件源(可安装软件列表)
    * sudp apt-get -y install vim                          #安装vim编辑器
    * sudo apt-get -y install openssh-server               #安装ssh服务
* 配置ssh免密码(私钥／公钥)登录：
    * sudo /etc/init.d/ssh start                            ＃启动ssh服务(否则ssh localhost测试可能会出现port 22 Connection refused !)
    * ssh-keygen -t rsa                                     #生成本机的公钥／私钥
    * cat ~/.ssh/id_rsa.pub >> ~/.ssh／authorized_keys      #将公钥加入授权列表(可通过ssh localhost测试)
* 安装java环境(也可通过压缩包安装)
    * sudo apt-get -y install openjdk-7-jre openjdk-7-jdk    ＃安装jdk和jre
    * dpkg -L openjdk-7-jdk |grep '/bin/javac'               ＃查看java安装目录
    * 在~/.basnrc文件中添加以下两行信息：                         #设置JAVA相关环境变量
    ```
    export JAVA_HOME=${上一步中查找的java安装目录}
    export PATH=$PATH:$JAVA_HOME/bin
    ```
    * source ~/.bashrc                                        ＃使环境变量生效
* 安装HADOOP环境[http://112.17.13.243/files/42440000077C9234/apache.fayea.com/hadoop/common/hadoop-2.7.1/hadoop-2.7.1.tar.gz](hadoop2.7.1下载地址)
    * tar -zxf hadoop-2.7.1.tar.gz -C /usr/local          ＃解压缩到/usr/local目录下
    * cd /usr/local && mv ./hadoop-2.7.1 hadoop/           #将HADOOP_HOME目录重命名为hadoop
    * ./hadoop/bin/hadoop version                         ＃测试hadoop是否可用