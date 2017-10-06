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
* 安装HADOOP环境[hadoop2.7.1下载地址](http://112.17.13.243/files/42440000077C9234/apache.fayea.com/hadoop/common/hadoop-2.7.1/hadoop-2.7.1.tar.gz)
    * tar -zxf hadoop-2.7.1.tar.gz -C /usr/local          ＃解压缩到/usr/local目录下
    * cd /usr/local && mv ./hadoop-2.7.1 hadoop/           #将HADOOP_HOME目录重命名为hadoop
    * ./hadoop/bin/hadoop version                         ＃测试hadoop是否可用
    * hadop默认为本地模式(无需配置),可通过$HADOOP_HOME/bin/hadoop jar ./share/hadoop/mapreduce/hadoop-mapreduce-examples-2.6.0.jar查看运行hadoop本身自带的程序示例

* Hadoop伪分布式配置：
    * 修改配置文件core-site.xml和hdfs-site.xml为一下内容(配置文件位于 /usr/local/hadoop/etc/hadoop/目录下)
    ```
    //core-site.xml 
    <configuration>
        <property>
            <name>hadoop.tmp.dir</name>
            <value>file:/usr/local/hadoop/tmp</value>
            <description>Abase for other temporary directories.</description>
        </property>
        <property>
            <name>fs.defaultFS</name>
            <value>hdfs://localhost:9000</value>
        </property>
    </configuration>

    //hdfs-site.xml   
    <configuration>
        <property>
            <name>dfs.replication</name>
            <value>1</value>
        </property>
        <property>
            <name>dfs.namenode.name.dir</name>
            <value>file:/usr/local/hadoop/tmp/dfs/name</value>
        </property>
        <property>
            <name>dfs.datanode.data.dir</name>
            <value>file:/usr/local/hadoop/tmp/dfs/data</value>
        </property>
    </configuration>`
    ```
    * $HADOOP_HOME/bin/hdfs namenode -format                  #执行 NameNode 的格式化
    * $HADOOP_HOME/sbin/start-dfs.sh                          #启动DFS,开启NameNode和DataNode守护进程
    * 通过命令jps查看是否启动成功,同时通过浏览器访问localhost:50070可查看NameNode和Datanode信息和HDFS中的文件
    * 伪分布式模式下运行程序示例
        * $HADOOP_HOME/bin/hdfs dfs -mkdir -p /user/hadoop              #创建用户目录
        * $HADOOP_HOME/bin/hdfs dfs -mkdir input                        #创建输入数据目录
        * $HADOOP_HOME/bin/hdfs dfs -put ./etc/hadoop/*.xml input       #将输入数据文件放在HDFS上
        * $HADOOP_HOME/bin/hadoop jar ./share/hadoop/mapreduce/hadoop-mapreduce-examples-*.jar grep input output 'dfs[a-z.]+' #执行示例程序
        * $HADOOP_HOME/bin/hdfs dfs -cat output/*                       #查看运行结果
    > 出现JAVA_HOME NOT SET错误，可直接在$HADOOP_HOME/etc/hadoop/hadoop-env.sh修改JAVA_HOME
    > iptables -t nat -A  DOCKER -p tcp --dport \<HOST_PORT\> -j DNAT --to-destination \<container_ip\>:\<DOCKER_PORT\> 可暴露运行时的容器指定端口
* 配置启动YARN：
    * 更改mapred配置文件和yarn配置文件：mv ./etc/hadoop/mapred-site.xml.template ./etc/hadoop/mapred-site.xml
    ```
    //mapred-site.xml
    <configuration>
        <property>
            <name>mapreduce.framework.name</name>
            <value>yarn</value>
        </property>
    </configuration>
    //yarn-site.xml
    <configuration>
        <property>
            <name>yarn.nodemanager.aux-services</name>
            <value>mapreduce_shuffle</value>
            </property>
    </configuration>
    ```
    * $HADOOP_HOME/sbin/start-yarn.sh      # 启动YARN服务
    * $HADOOP_HOME/sbin/mr-jobhistory-daemon.sh start historyserver  #开启历史服务器，用于在Web中查看任务运行情况(默认端口8088)
<<<<<<< HEAD
=======

* 配置hadoop分布式集群:
    * 前期准备:
        * sudo vim /etc/hostname     #修改主机名(统一修改为hadoop-master 或 hadoop-slave$N)
        * sudo vim /etc/hosts    ＃配置hosts文件，使得主机名与ip对应
    * 配置ssh免密码登录
        * (master) ssh-keygen -t rsa   && cat ~/.ssh/id_rsa.pub >> ~/.ssh/authorized_keys    ＃(hadoop-master)生成私钥／公钥,将公钥添加到授权列表
        * (master) scp ~/.ssh/id_rsa.pub hadoop@hadoop-slave$N:/home/hadoop/     ＃将公钥复制到从机slave中
        * (slave) cat ~/id_rsa.pub >> ~/.ssh/authorized_keys                     ＃将主服务器公钥添加到所有slave丛机的授权列表
    * 修改hadoop配置文件，进行分布式集群配置:
        * slaves定义DataNode的主机名，如hadoop-slave(伪分布式一般定义为localhost)
        ```
        //core-site.xml
        <configuration>
            <property>
                    <name>fs.defaultFS</name>
                    <value>hdfs://Master:9000</value>
            </property>
            <property>
                    <name>hadoop.tmp.dir</name>
                    <value>file:/usr/local/hadoop/tmp</value>
                    <description>Abase for other temporary directories.</description>
            </property>
        </configuration>

        //hdfs-site.xml
        <configuration>
            <property>
                    <name>dfs.namenode.secondary.http-address</name>
                    <value>Master:50090</value>
            </property>
            <property>
                    <name>dfs.replication</name>
                    <value>1</value>
            </property>
            <property>
                    <name>dfs.namenode.name.dir</name>
                    <value>file:/usr/local/hadoop/tmp/dfs/name</value>
            </property>
            <property>
                    <name>dfs.datanode.data.dir</name>
                    <value>file:/usr/local/hadoop/tmp/dfs/data</value>
            </property>
        </configuration>

        //mapred-site.xml
        <configuration>
            <property>
                    <name>mapreduce.framework.name</name>
                    <value>yarn</value>
            </property>
            <property>
                    <name>mapreduce.jobhistory.address</name>
                    <value>Master:10020</value>
            </property>
            <property>
                    <name>mapreduce.jobhistory.webapp.address</name>
                    <value>Master:19888</value>
            </property>
        </configuration>

        //yarn-site.xml
        <configuration>
            <property>
                    <name>yarn.resourcemanager.hostname</name>
                    <value>Master</value>
            </property>
            <property>
                    <name>yarn.nodemanager.aux-services</name>
                    <value>mapreduce_shuffle</value>
            </property>
        </configuration>
        ```
        * 配置完成后，将Master上的 /usr/local/Hadoop 文件夹复制到各个节点上
        * hdfs namenode -format       # 首次运行需要执行初始化
        * start-dfs.sh　&& start-yarn.sh  && mr-jobhistory-daemon.sh start historyserve  ＃启动服务
        >注意：１．启动服务可能出现JAVA_HOME环境变量配置问题，可在hadoop-env.sh中直接指定JAVA_HOME目录(sudo执行shell脚本可能会重置环境变量)
        </br> 2.ssh连接失败时,检查是否是防火墙问题 
>>>>>>> af6fef73c4df13264e1087bc44048c35b6f15eed
