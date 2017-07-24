## 任务定义： 
* 在系统中每个learning task 用一个有向无环图(DAG)进行描述：每个节点代表一个单机学习算法或分布式学习算法，每条边代表从一个节点流向它对应子节点的的数据流，后台根据提交的task DAG调度执行相应的任务；
> 注：学习任务可以通过图形界面手动自定义，也可以从一个模板克隆得到。同时可以通过图像界面的变化观察到后台任务的执行情况；

## 系统组成：
### 分布式学习算法库：
* 主要基于Spark实现，包括现在主流的机器学习算法，数据预/后处理算法，数据格式转换，特征工程，性能测试等
### 图形化的任务定义系统：
* 用户可以利用拖拽的形式定义自己的任务，可以自定义模型参数，可以自己上传自己的算法和数据集；
### 后台分布式服务器
* 主要基于hadoop和spark的任务执行云服务器，根据接收到的task DAG,在依赖数据集准备就绪的情况下，根据DAG node自动调度执行。对应的算法会根据自身的定义在Linux,MapReduce 或spark上运行。

## 数据库设计：
* 账户表：存储用户的个人信息和权限信息，包括用户名，密码，有效时间等字段：
```
CREATE TABLE `account` (
  `email` varchar(255) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varbinary(41) NOT NULL DEFAULT '',
  `verified` varchar(255) DEFAULT NULL,
  `createtime` varchar(255) NOT NULL,
  `serial` varchar(255) DEFAULT NULL,
  `token` varchar(255) DEFAULT NULL,
  `activetime` varchar(255) DEFAULT NULL,
  `company` varchar(255) DEFAULT NULL,
  `position` varchar(255) DEFAULT NULL,
  `verifylink` varchar(255) DEFAULT NULL,
  `power` varchar(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
* 任务表：主要用来描述用户定义/提交的任务信息，包括任务DAG的XML定义，关联账户，执行情况等字段：
```
CREATE TABLE `bdajob` (
  `job_id` varchar(255) CHARACTER SET utf8 NOT NULL,
  `job_name` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `graphxml` text CHARACTER SET utf8,
  `account` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 DEFAULT NULL,
  `last_submit_time` varchar(30) CHARACTER SET utf8 DEFAULT NULL,
  `endtime` varchar(30) CHARACTER SET utf8 DEFAULT NULL,
  `is_example` tinyint(5) DEFAULT '0',
  `oozie_job` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`job_id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;
```
* 类目表：主要用来保存数据集以及节点算法的分类信息，包括存储路径，父/子节点，创建时间等字段：
```
CREATE TABLE `category` (
  `id` varchar(100) NOT NULL,
  `name` varchar(60) NOT NULL,
  `level` varchar(60) NOT NULL,
  `type` varchar(10) NOT NULL,
  `path` longtext,
  `fatherid` varchar(100) DEFAULT NULL,
  `haschild` tinyint(1) DEFAULT NULL,
  `createtime` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```


* 数据集表：描述需要处理的数据集信息，包括分类，属主，存储路径，数据类型，存储类型等字段
```
CREATE TABLE `dataset` (
  `id` varchar(100) NOT NULL,
  `name` varchar(60) NOT NULL,
  `category` varchar(100) DEFAULT NULL,
  `owner` varchar(100) DEFAULT NULL,
  `path` varchar(255) NOT NULL,
  `deprecated` tinyint(1) DEFAULT NULL,
  `contenttype` varchar(30) DEFAULT NULL,
  `version` varchar(30) DEFAULT NULL,
  `createdate` varchar(100) DEFAULT NULL,
  `description` mediumtext,
  `storetype` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```


