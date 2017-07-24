---
title: Docker学习记录
date: 2017-07-20 11:02
---
# docker学习记录

## 概述
* Docker为利用Go语言开发，基于LXC（Linux Container）技术之上构建的Container容器引擎；
* 优点：屏蔽了开发/部署环境的复杂性，通过image打包，简化了多种应用的部署工作，适应了云计算的发展

### 技术概览
* Docker核心为操作系统级虚拟化（隔离性，可度量/可配额，便携性，安全性）方法：

#### 隔离性：
* namespace将container的进程(pid)、网络(net)、消息(ipc)、文件系统(mnt)、UTS("UNIX Time-sharing System")和用户空间(user)隔离开:-通过pid namespace实现不同用户的   
    * 进程相互隔离。在Docker中所有LXC进程的父进程为Docker进程，每个LXC进程有不同namespace。由于允许嵌套，可实现Docker in Docker
    * 通过net namespace实现网络隔离。每个net namespace有独立的network devices, IP addresses, IP routing tables,/proc/net目录。docker默认采用veth的方式将container中的虚拟网卡同host上的一个docker bridge: docker0连接在一起
    * - container中进程交互采用linux常见的进程间交互方法(interprocess communication - IPC，包括常见的信号量、消息队列和共享内存），实际为host上相同pid namespace的进程间交互，需要在IPC资源申请时赋予每个IPC资源有一个唯一的32位ID(标识namespace信息)
    * - mnt namespace允许不同namespace的进程看到的文件结构不同，从而将每个namespace中的进程所看到的文件目录就被隔离开，同时每个namespace中的container在/proc/mounts的信息只包含所在namespace的mount point。
    * - uts namespace允许每个container拥有独立的hostname和domain name,使其在网络上可以被视作一个独立的节点而非Host上的一个进程。
    * 8- user namespace：每个container可以有不同的user和group id,即在container内部用container内部的用户执行程序而非Host上的用户

#### 可度量/可配额
* cgroups 实现了对资源的配额和度量,提供类似文件的接口，在 /cgroup目录下新建一个文件夹即可新建一个group;新建task文件，并将pid写入该文件，即可实现对该进程的资源控制,可以限制blkio(每个块设备的输入输出控制)、cpu(使用调度程序为cgroup任务提供cpu的访问)、cpuacct(产生cgroup任务的cpu资源报告)、cpuset(对于多核CPU,为cgroup任务分配单独的cpu和内存)、devices(允许或拒绝cgroup任务对设备的访问)、freezer( 暂停和恢复cgroup任务)、memory(设置每个cgroup的内存限制以及产生内存资源报告)、net_cls(标记每个网络包以供cgroup方便使用)、ns(名称空间)九大子系统的资源

#### 便携性
* AUFS (AnotherUnionFS)支持将不同目录挂载到同一个虚拟文件系统下(unite several directories into a single virtual filesystem)的文件系统,支持为每一个成员目录(类似Git Branch)设定读写权限

* 典型的启动Linux运行需要两个FS: bootfs + rootfs:bootfs (boot file system) 主要包含 bootloader 和 kernel, bootloader主要是引导加载kernel, 当boot成功后 kernel 被加载到内存中后 bootfs就被umount了. rootfs (root file system) 包含的就是典型 Linux 系统中的 /dev, /proc,/bin, /etc 等标准目录和文件。
* 在Docker中，初始化时将rootfs以readonly方式加载并检查，然而接下来利用union mount的方式将一个readwrite文件系统挂载在readonly的rootfs之上，并且允许再次将下层的FS(file system)设定为readonly并且向上叠加, 从而一组readonly和一个writeable的结构构成一个container的运行时态, 每一个FS被称作一个FS层
* 得益于AUFS的特性, 每一个对readonly层文件/目录的修改都只会存在于上层的writeable层中。从而由于不存在竞争, 多个container可以共享readonly的FS层。 因此Docker将readonly的FS层称作"image" - 对于container而言整个rootfs都是read-write的，但事实上所有的修改都写入最上层的writeable层中, image不保存用户状态，只用于模板、新建和复制使用
* 上层的image依赖下层的image，因此Docker中把下层的image称作父image，没有父image的image称作base image。因此想要从一个image启动一个container，Docker会先加载这个image和依赖的父images以及base image，用户的进程运行在writeable的layer中。所有parent image中的数据信息以及 ID、网络和lxc管理的资源限制等具体container的配置，构成一个Docker概念上的container。

### 安全性
* 由kernel namespaces和cgroups实现的Linux系统固有的安全标准;
* Docker Deamon的安全接口
* Linux本身的安全加固解决方案,类如AppArmor, SELinux

## 使用dockerfile构建image镜像:
* 概述:dockerfile有自己特定的定义语言，构建时会被翻译为linux指令;根据作用，指令分为:
    * 构建指令:用于构建image不在docker容器中运行)；
    * 设置指令：用于设置docker容器的属性，将会在docker容器中执行)）
### 构建指令：
* FROM \<image\>\[:tag\]: 指定基础image，可通过官方远程仓库/本地仓库指定（tag可选指定版本）,一般用在文件开始位置；
* MAINTAINER \<name\>: 用来指定镜像创建者信息(可通过docker inspect命令查看)
* RUN \<command\>或RUN \["executable","param1","param2"....\] :用来安装软件，可运行基础image支持的命令
* ENV \<key\> \<value\>: 用于在image中设置环境变量(以保证RUN命令可以使用)，可通过docker inspect查看环境变量信息，此外也可以在启动container时通过docker run -- env key=value 设置或修改环境变量信息
* ADD \<src\> \<dest\>:将指定文件或目录拷贝到container的dest路径，所有拷贝到container的文件权限为0755，uid和gid为0；
    * 如果src是目录会将目录下为所有文件拷贝(不包括目录);
    * 如果src为文件,当dest为目录且带有'/'是，会将文件拷贝到该目录下,如果dest为一个文件且不带'/'时，将将源文件追加到目标文件
    * 如果src为可识别的压缩文件，在container中会自动将其解压缩；


### 设置指令：
* CMD \["executable","param1","param2"....\] 或 CMD <command> <param1> ...: 设置container启动时执行的操作，只能在文件中出现一次（若出现多次则只执行最后一条，若需要执行多条指令，可通过shell脚本定义）
> 当Dockerfile指定了ENTRYPOINT（为一个可执行脚本或程序路径），可使用CMD \[”param1“,"param2"\]形式，param作为ENTRYPOINT的执行参数；
* ENTRYPONT \["executable","param1","param2"....\] 或 ENTRYPOINT <command> <param>: 设置container启动时执行的操作，可与CMD配合使用，用来指定其执行参数，也可以单独使用，此时会与CMD相互覆盖，只有最后一个有效
* USER \<name\>: 用来设置container的用户
* EXPOSE \<port\> \[...\]: 指定容器需要映射到宿主机的端口(container port->host port),在运行container时通过-p选项加上EXPOST定义的端口实现映射(可以指定宿主机端口，也可以不指定随机选择) 
> 注:每次运行container时，其ip地址不固定，在桥接网卡的地址范围内随机生成，因此可利用端口映射，通过宿主机访问container内的端口服务，可以通过 docker port \<conttainer_port\> \<container_ID\> 查看映射到的宿主机端口
* VOLUME \["\<mountpoint\>"\]: 指定挂载点，使container中一个目录具有持久化存储数据的功能，该目录可以被container自身使用，也可以共享给其他container使用
> 注：container使用为文件系统为AUFS,不能持久化数据，当container关闭时，所有更改会丢失。
* WORKDIR \<path\>： 切换目录，相当于cd;
* ONBUILD <Dockerfile关键字>：在子镜像中执行
