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
* 构建时需要将DockerFile放在一个特定的文件夹，称此环境为构建上下文，在此目录执行docker build(-t指定仓库和名称)则Docker构建镜像时，会将构建上下文和此环境的下的目录和文件上传到Docker的守护进程，从而Docker守护进程可直接访问想在镜像中存储的文件数据(可利用.dockerignore文件忽略指定文件或目录)；
* DockerFile中的指令会依照顺序执行，每条指令都会创建一个新的镜像层并对镜像进行提交(类似docker commit),然后在　基于刚提交的镜像运行一个新容器
* 使用构建缓存可以实现简单的DockerFile模板，基于模板简化可以简化DockerFile的编写
### 构建指令：
* FROM \<image\>\[:tag\]: 指定基础image，可通过官方远程仓库/本地仓库指定（tag可选指定版本）,一般用在文件开始位置；
* MAINTAINER \<name\>: 用来指定镜像创建者信息(可通过docker inspect命令查看)
* RUN \<command\>或RUN \["executable","param1","param2"....\] :用来安装软件，可运行基础image支持的命令（前一个默认使用/bin/sh -c执行）
* ENV \<key\> \<value\>: 用于在image中设置环境变量(以保证RUN命令可以使用)，可通过docker inspect查看环境变量信息，此外也可以在启动container时通过docker run -- env key=value 设置或修改环境变量信息
* ADD \<src\> \<dest\>:将指定文件或目录(只能是位于构建上下文中)拷贝到container的dest路径(目录路径必须以“／”结尾)，所有拷贝到container的文件权限为0755，uid和gid为0；
    * 如果src是目录会将目录下为所有文件拷贝(不包括目录);
    * 如果src为文件,当dest为目录且带有'/'是，会将文件拷贝到该目录下,如果dest为一个文件且不带'/'时，将将源文件追加到目标文件
    * 如果src为可识别的压缩文件，在container中会自动将其解压缩；
* COPY:类似ADD指令，但不做文件提取和解压缩

### 设置指令：
* CMD \["executable","param1","param2"....\] 或 CMD <command> <param1> ...: 设置container启动时执行的操作，只能在文件中出现一次（若出现多次则只执行最后一条，若需要执行多条指令，可通过shell脚本定义），在执行docker run启动容器时指定的命令，会覆盖CMD定义的命令
> 当Dockerfile指定了ENTRYPOINT（为一个可执行脚本或程序路径），可使用CMD \[”param1“,"param2"\]形式，param作为ENTRYPOINT的执行参数；
* ENTRYPONT \["executable","param1","param2"....\] 或 ENTRYPOINT <command> <param>: 设置container启动时执行的操作，可与CMD配合使用，用来指定其执行参数，也可以单独使用，此时会与CMD相互覆盖，只有最后一个有效，执行docker run命令时附带的任何参数都会作为参数传递给ENTRYPOINT指令的命令(也可以在docker run --entrypoint覆盖ENTRYPOINT指定的指令)
* USER \<name\>: 用来设置container的用户
* EXPOSE \<port\> \[...\]: 指定容器需要映射到宿主机的端口(container port->host port),在运行container时通过-p选项加上EXPOST定义的端口实现映射(可以指定宿主机端口，也可以不指定随机选择) 
> 注:每次运行container时，其ip地址不固定，在桥接网卡的地址范围内随机生成，因此可利用端口映射，通过宿主机访问container内的端口服务，可以通过 docker port \<conttainer_port\> \<container_ID\> 查看映射到的宿主机端口
* VOLUME \["\<mountpoint\>"\]: 指定挂载点，使container中一个目录具有持久化存储数据的功能，该目录可以被container自身使用，也可以共享给其他container使用
> 注：container使用为文件系统为AUFS,不能持久化数据，当container关闭时，所有更改会丢失。
* WORKDIR \<path\>： 切换目录，相当于cd，可用来设定命令执行的工作目录;
* ONBUILD <Dockerfile关键字>：在子镜像中执行，为镜像添加触发器(当一个镜像被用作其他镜像的基础镜像时，该镜像中的触发器将会被执行)，触发器会在构建过程中在FROM之后插入新指令(为下面准备环境等)，触发器只能继承一次(即只能在子镜像中执行，不会在孙子镜像中执行)




## 联合文件系统(Union File System):
* 联合文件系统内部允许多个文件系统堆叠，目录中可能包含来自多个文件系统的文件(若两个文件具有相同的路径，则最后挂载的文件会隐藏之前的文件)，但在用户看来只是一个单独的文件系统
* docker支持多种不同的UFS实现，包括AUFS,Overlay,devicemapper,BTRFS,ZFS等，根据系统需要选择(可通过docker info查看)


## Linux namespace
* 概述:namespace是linux对计算机PID,IPC，network等资源进行隔离的一种方案，只需在调用clone创建进程时指定不同的flag,使得属于不同namespace的资源资源之间相互隔离，彼此透明，LXC(linux conatainer)的是实现就是利用了这一技术，使得属于不同container的资源位于不同namespace;
### PID隔离
* 当调用clone()创建进程时，指定flag为CLONE_NEWPID，将会创建一个新的PID Namespace,提供一个独立的PID环境，创建出的进程将作为该namespce的第一个进程(PID=1，)，在该namespace内调用fork,vfork,clone产生的进程都将从属该namespace(产生一个在namespace中独立的PID)
* 在同一namespace下的孤儿进程都以初始时创建的PID=1的进程为父进程，当该进程(PID=1)结束时该namespace下的所有进程都会随之结束
* PID namespace均有层次性，在一个namespace下创建的新的namespace将作为该namespace的子namespace，同时子namespace中的进程同时对上层namespace可见，从而一个进程不止拥有一个PID(在其所属的上层namespace中均有一个PID)；当linux系统启动时将创建一个默认的PID namespace作为以后所有PID namespace的祖先(因此系统所有进程在该祖先PID namespace中可见)

### IPC隔离
* 当调用clone()创建进程时，指定flag为CLONE_NEWIPC，将会创建一个新的IPC Namespace,创建出的进程将作为该namespce的第一个进程;
* 一个IPC namespace有一组System V IPC objects标识符构成，在该IPC namespace下创建的IPC object仅仅对此namespace下的进程可见,从而保证只有同一个namespace下的进程才可以直接通信，当IPC namespace被销毁时,该namespace下的所有IPC object随之销毁
* 可在调用clone()时同时指定flag为CLONE_NEWPID和CLONE_NEWIPC同时实现PID和IPC隔离，保证不同namespce下进程彼此不可见且不能相互通信，从而实现进程间的隔离

### 文件系统隔离
* 当调用clone()时，指定flag为CLONE_NEWNS,将会创建一个新的mount namespace,为进程提供一个文件层次视图(否则子进程和父进程将共享一个mount space)，从而子进程调用mount或umount将会创建一份新的文件层次视图(否则将会影响该namespace下的其他进程)，
* 结合pivot_root,可以为进程创建一个独立的目录空间
### 网络资源隔离
* 当调用clone()时，指定flag为CLONE_NEWNET,将会创建一个新的network namespace,为进程提供一个完全独立的网络协议栈视图(包括网络设备接口，IPv4/IPv6协议栈，IP路由表，防火墙规则，socket等)，类似一个独立的系统。
* 一个物理设备只能存在于一个network namespace(可以移动)，
* 虚拟网络设备(virtual network device)提供了类似管道的抽象，可以在不同的namespace之间建立隧道，从而建立其到其他namespace下物理设备的桥接。
* 当一个network namespace被销毁时，其中的物理设备将会转移到系统初始init network namespace
### 主机host隔离
* 当调用clone()时，指定flag为CLONE_NEWUTS，将会创建一个新的UTS namespace(即一组被uname返回的标识符)，新的UTS namespace将会复制其所属的上层UTS namespace的标识符来初始化，
* 同时clone出的进程可以通过相关系统调用改变这些标识符(如sethostname()改变该namespace的HOSTNAME)改变对同一namespace下的所有进程可见
* 结合CLONE_NEWNET和CLONE_NEWUTS可以虚拟出一个具有独立主机名的网络空间环境  

### 小结
* 以上所有flag均可以结合一起使用