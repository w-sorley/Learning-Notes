# Linux内核设计与实现

## 第二章：进程管理
* 进程（任务）：是正在执行的程序以及它所占有的资源的总称
    * 拥有进程描述符
    * 存在内核的任务队列（双向循环链表）
* 线程: 具体的执行单元 
    * 拥有独立程序计数器，进程栈和进程寄存器，
    * 相互之前共享虚拟内存,拥有独立的虚拟处理器
    * 是内核调度的对象
### 2.1进程描述符以及任务队列
* 进程描述符(task_struct)
    * 包含了内核管理一个进程所需的全部信息，
    * 通过 slab分配器进行分配，以实现对象复用(避免动态分配和释放)和缓存着色
    * 2.6版本之前每个进程的task_struct存放于内核栈的尾端(可利用栈指针访问，避免占用额外寄存器),现在在栈底(或栈顶)存放thread_info的结构体（拥有指向进程描述符的指针）
    * 内核通过唯一的进程标识值(或PID)来标识每个进程，最大为32767(如不考虑后向兼容，可修改/proc/sys/kernel/pid_max)，存放在task_struct中
    * 通过current   
### 注:
* 缓存着色技术:
    * 整个内存被分为n个cache page，每个cache page包含数量固定的cache line。 
    * 整个cache被分为m个cache way,每个cache way有相同数量的cache line。
    * 内存中的数据依据自己在内存中的cache line索引[getCacheLineIndex(addr)]只能被放入某一个cache way中相对应的cache line里面。
    * 假设已经从地址中提取出cache line的索引i，那么硬件会同时访问所有cache way的第i块cache line，找出一个拥有空闲行i的cache way放入。找不到则启动淘汰策略，淘出一个空行((set-association）)
    * color则将不同slab中的同样的数据结构的地址进行一个内存页偏移,因此这些数据结构的cache line索引就错开了(避免频繁淘汰)从而能更好的利用cache
