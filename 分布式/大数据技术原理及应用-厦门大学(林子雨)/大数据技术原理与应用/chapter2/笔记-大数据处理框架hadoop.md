
优点
* 屏蔽底层，易于使用
* Java实现跨平台,多语言支持
* 核心：M/R处理,HDFS存储

特性:
高可靠／高容错(冗余副本)，高效／成本低／可扩展（分布式）

hadoop家族:
Tez:任务流程ＤＡＧ优化
spark类似MR但在内存中执行，
Hive支持SQL,底层转化为一系列的M/R,
Pig，类SQL,轻量级查询
Oozie:工作流管理
Zookeeper:协调调度，集群管理，分布式锁
Hbase:面向列，支持随机读写
Flume:日志收集
Sqoop:关系数据库导入Hadoop平台，或导出到外部关系数据库　
Ambari:快速部署管理

Hadoop主要节点:
NameNode:协调数据存储
DataNode:数据被拆分为数据块
JObTracker:管理M/R作业，会把大的作业拆分为小作业，分配给TaskTracker.管理TaskTracker
TaskTracker：部署在不同的节点上，执行作业
SecondaryNameNode:冷备份，NameNOde失败备份不能马上生效，可加快启动
(在集群中根据不同节点特性选择不同的硬件配置)

分布式文件系统HDFS
