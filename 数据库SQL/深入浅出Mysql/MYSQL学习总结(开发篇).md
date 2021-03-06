---
title: MYSQL学习总结(开发)
date: 2017-05-22
tags: [SQL,MYSQL,Web,开发]
categories: DataBase
---
# MYSQL开发学习总结
> －－总结自＜深入浅出MYSQL:数据开发，优化与管理维护．第２版＞（开发篇）


# 表类型(存储引擎)的选择
* Mysql只用插件式存储引擎，可以针对不同的存储需求选择最优的存储引擎。
* 可以使用命令：‘SHOW ENGINE \G’查看当前数据库支持的存储引擎
* 在创建新表时可以利用‘ENGINE=?’关键字指定存储引擎(若不指定,在mysql5.5版本后默认为InnoDB),此外可以使用alter更改已存在表的存储引擎
## 存储引擎特性：
｜　特点　　   　｜　MyISAM　　　｜　　　InnoDB　　｜　　MEMORY　　　｜　　MERGE　　｜　　　　　NDB　　　|                  
｜-------------|-------------|---------------|----------------|-------------|----------------|------------------|
| 　 存储限制　　｜　有　      　｜　64TB　　　　｜　　　　有　       ｜　　　没有　　　|  有            |                  
｜　　事物安全　　｜　　　      　｜　　　支持　　｜　　　　      　｜　　　　　　　　｜　　　　　　　　　|                  
｜　　锁机制　　　｜　　表锁　 　｜　　行锁　　　｜　　　表锁　　    ｜　　表锁　　　　　　｜　　　行锁　　　　　　|                  
｜　B树索引　　　｜　支持　 　　｜　　支持　　　｜　　支持　　    　｜　支持　　　　　　　｜　　　支持　　　　　　|                  
｜　　　哈希索引　｜　　　　　　｜　　　　　    ｜　　支持　　  　｜　　　     　　　　｜　　　　支持　　　　　|                  
｜　　全文索引　　｜　　支持　　｜　　　　    　｜　　　　　      ｜　　　　　　   　　｜　　　　　　　　　|                  
｜　集群索引　　　｜　　　   　｜　　　　支持　｜　　　　　        ｜　　　　　　　  　｜　　　　　　　　　|                  
｜　　数据缓存　　｜　　　    　｜　　支持　　　｜支持　　　　    　｜　　　　　　　  　｜　　　支持　　　　　　|                  
｜　　索引缓存　　｜　支持　　　｜　支持　　　　｜　　　支持　     　｜　　支持　　　　　　｜　　　　　支持　　　　|                  
｜数据可压缩　　　｜　　支持　　｜　　　　   　｜　　　　          　｜　　　　　　　　｜　　　　　　　　　|                  
｜　空间使用　　　｜　　低　 　｜　　高　　　 ｜　　N/A       　　　｜　　低　　　　　　｜　　　低　　　　　　|                  
｜　　内存使用　　｜　　低　　｜　　高　　　   ｜　　中等　　     　｜　　低　　　　　　｜　　　低　　　　　　|                  
｜　　批量插入速度｜　　高　　｜　低　　　　   ｜　　　高　       　｜　　高　　　　　　｜　　　　　高　　　　|                  
｜　支持外键　　　｜　　　 　｜　支持　　　  　｜　　　　　         ｜　　　　　　　　｜　　　　　　　　　|                  


## MyISAM:
* 优点：访问速度快；缺点：不支持事物，外键；
* 适用于对事物完整性没有要求或以select或insert为主的应用
* 每个MyISAM在磁盘上存储为和表名相同的三个文件，扩展名分别为.frm(存储表定义).MYD(MYData，存储表数据).MYI(MYIndex，存储表索引)
> (数据文件和索引文件可以放在不同的目录中，以平均分布IO获得更快速度，可以在创建表时通过DATA DIRECTORY和INDEX DIRECTORY指定)
* 可以通过CHECK TABLE命令检查MyISAM表是否被损坏／通过REPAIR TABLE命令修复损坏的MyISAM表
* MyISAM支持三种不同的存储格式：
<br/>－1.静态表（默认）：优点：存储快，易缓存，易恢复／缺点：固定长度，空间占用相对较大，存储时会按照列的定义宽度补足空格(应用访问时会去掉，注：若数据本身尾部带有空格，访问时也会被去掉)；
<br/>－2.动态表：优点：变长字段，记录长度不固定，占用空间相对较少／缺点:频繁更新删除记录会产生碎片(定期执行OPTIMIZE TABLE语句或myisamchk-r命令改善性能)，且不易恢复；
<br/>－3.亚索表：由myisampack工具创建，空间占用小，每个记录单独压缩，访问开支小
## InnoDB
* 优点：提供具有提交，回滚和崩溃恢复能力的事物安全；缺点：写的处理效率相对较差，需要更大的空间保存数据和索引
### Innodb特点：
#### 1.自动增长列：
* 可手动插入，若插入的值为null或0则实际设置为增长后的值
* 可以通过”ALTER TABLE *** AUTO_INCREMENT=n“设置增长的初始值，默认为１(设置的初始值保存在内存中，重启数据库失效)
* 可以通过"SELECT　LAST_INSERT_ID()"查询当前线层最后插入的增长值(若一次插入多条则返回第一条)
> 注：对于InnoDB引擎，自动增长列必须为索引或组合索引的第一列(对于MyISAM可以为组合索引的其他列)
#### 2.外键约束：
* MYSQL存储引擎中只有InnoDB支持外键，在创建外键时，父表必须有对应索引，子表在创建外键时也会自动创建对应索引
* 在创建索引时，可以指定在删除，更新父表时，对子表进行的相应操作：
* RESTRICT/NO ACTION:指限制在子表有关联记录的情况下，父表不能更新
* CASCADE:表示父表在更新或者删除时，更新或删除子表对应记录(可能导致数据丢失)
* SET NULL:表示父表在更新或者删除时，子表对应的字段置为空
> 注：在执行LOAD DATA或ALTER TABLE操作时，可以暂时关闭外键约束加快处理速度：SET FOREIGN_KEY_CHAECK=0(关闭)／１(打开)
#### 3.存储方式：
* InnoDB有两种方式存储表和索引：
* 使用共享表空间存储：表结构保存在.frm文件中，数据和索引保存在innodb_data_home_dir和innodb_data_file_path定义的表空间中(可以是多个文件)
* 使用多表空间存储：表结构保存在.frm文件中，每个表的数据和索引单独保存在.ibd中(如果是分区表，则每个分区对应单独.ibd文件，文件名＝表名＋分区名，同时可以在创建分区时指定数据文件位置)，多表空间的数据文件没有大小限制，方便单表备份和恢复(使用ALTER TABLE tab_name DISCARD TABLESPACE;/ALTER TABLE tab_name IMPORT TABLESPACE;将备份恢复到原表所在数据库，此外亦可以使用mysqldump/mysqlimport将单表恢复到目标数据库)
> 注：要使用多表空间的存储方式，需设置innodb_file_per_table参数，重启服务对新建表生效；即使在多表空间的存储方式下，共享表空间仍是必须的(Innodb把内部数据词典和在线重做日志放在放在这个文件中)

## MEMORY
* 使用存在于内存中的内容来创建表，每个MEMORY表只实际对应一个.frm磁盘文件；
* 优点：访问速度快(数据存放在内存中，且默认使用HASH索引，创建时还可指定为BTREE索引)／缺点：一旦服务关闭，表中数据就会丢失
* 在MYSQL启动时，通过--init-file选项，可以把INSERT INTO...SELECT或LOAD DATA INFILE等语句，放在指定文件中，从而服务启动时从持久稳固的数据源装载表
* 每个MEMORY表中可以存放的数据量大小，受到max_heap_table_size系统变量的约束，同时在创建时可以通过MAX_ROWS子句指定表的最大行数
* 适用于内容变化不频繁的代码表或统计操作的中间结果表
## MERGE
* 为一组结构完全相同的MyISAM表的组合，MERGE表本身没有数据，对其进行查询，更新，删除等操作时，实际操作内部的MyISAM表；
* 对其进行插入操作时，通过INSERT_METHOD子句定义插入的表（定义为FIRST/LAST使得插入操作作用在第一/最后一个表上，不定义或定义为NO表示禁止对此表进行插入操作）
* 对MERGE表的DROP操作，只是删除MERGE的定义，对内部表无影响
* MERGE在磁盘上保存为以表名开始的两个文件，.frm文件存储表定义；.MRG文件包含组合表的信息（如有哪些表组成，插入新数据的依据等）可以通过修改.MRG文件来修改MERGE表(需通过FLUSH TABLE刷新生效)
* MERGE表不能智能的将记录写到对应的表中，而分区表(5.1以后)可以

## TokuDB
* 为第三方存储引擎,是一个高性能支持事物处理的MYSQL和MariaDB存储引擎，具有高扩展性，高压缩率，高效的写入性能，支持大多数DDL操作。
> （关于第三方存储引擎，多适用于某些特定应用，如列式存储引擎Infobright）
TokuDB特性：
* 使用Fractal树索引保证高效的插入性能
<br/>　-优秀的压缩特性，比InnoDB高近10倍
<br/>　-Hot Schema Changes特性支持在线创建索引和添加／删除属性列等DDL操作
<br/>　-使用Bulk Loader达到快速加载大量数据
<br/>　-提供主从延迟消除技术
<br/>　-支持ACID和MVCC
* 适用于日志数据(插入频繁，存储量大)历史数据(不会再写，高压缩)和其他在线DDL较频繁的场景(增加系统可用性)

## 存储引擎选择
* MyISAM:适用于以读和插入操作为主，更新删除操作较少，并对事物完整性，并发性要求不高的场景。如Web，数据仓库等；
* InnoDB:事物处理，支持外键，适用于对事物完整性要求较高，在并发下要求数据的一致性，除插入查询操作外，还包括很多更新删除操作的场景。InnoDB除了有效降低了由于删除和更新导致的锁定，还可以确保事物的完整提交和回滚。如计费系统财务系统等；
* MEMORY:数据存储在内存，访问速度快，但对表的大小有限制，需保证数据库异常终止后数据可恢复。通常适用于更新不太频繁需要快速访问的小表；
* MERGE:可以突破对单个MyISAM表大小的限制，通过将不同表分布在多个磁盘，改善访问效率。适用于数据仓库等VLDB环境

# 选择合适的数据类型
## CHAR与VARCHAR
* CHAR:固定长度字符类型，空间占用较大，需要处理行尾空格(存储时空格补足，现实输出时行尾空格全部删除)，同时，处理速度相对较快，适用于长度变化不大，查询速度要求较高的场景；
* VARCHAR：可变长度字符类型，,在非严格模式下超过规定字符长度可以存储，在严格模式下超过规定字符长度不能存储，并会出现错误提示；
* 在MYSQL中，不同的存储引擎对CHAR和VARCHAR的使用原则不同，字符类型应根据存储引擎进行相应的选择：
<br/>　－MYISAM:建议使用固定长度数据列
<br/>　－MEMORY:都使用固定长度数据列存储，因此使用时两者均作为CHAR处理
<br/>　－InnoDB:建议使用VARCHAR,对于InnoDB数据表，内部行存储格式没有区分固定长度和可变长度(所有数据行都使用指向数据列值的头指针)，本质上都是固定长度，因此两者性能相当，但VARCHAR平均占用的空间相对较少。
* CHAR与VARCHAR适合保存少量字符串,当保存较大文本时通常选择TEXT或BLOB

### TEXT与BLOB
* BLOB:可用来保存二进制数据(如图片)，根据存储字节长度又分为BLOB,MEDIUMBLOB,LONGBLOB
* TEXT:只能保存字符数据，根据存储文本长度又分为TEXT,MEDIUMTEXT,LONGTEXT
* 对于BLOB和TEXT值的删除会在数据表中留下很大空洞，在插入记录填入空洞时性能会受影响(建议定期使用OPTIMIZE TABLE功能对表进行碎片整理)
* 可以使用合成的(Synthetic)索引提高大文本字段(TEXT/BLOB)的查询性能。
> 合成索引即根据大文本字段的内容建立一个散列值(可以使用MD5()/SHA1()/CRC32(),也可以自定义逻辑计算散列值)，并保存在单独的数据列中，以后可通过检索散列值找到数据行，只能用于精确匹配的查询(不能用于"<"">="等范围搜索)
* 可以利用前缀索引(即只对前n列建立索引)，对TEXT/BLOB字段进行模糊查询
* 应尽量避免检索大型的BLOB／TEXT值，可以利用索引列决定需要的数据行，然后从符合条件的数据行中检索BLOB/TEXT值。
* 在某些环境中，将TEXT/BLOB移到单独的数据表中，将原始表的数据列转为固定长度，可以减少主表中的碎片，充分利用固定长度列的性能优势，同时避免主表检索时过大的网络传输；

### 浮点数与定点数
* 浮点数：插入数据的精度超过定义的实际精度时，插入值会被四舍五入到定义的实际精度，在MYSQL中包括float,double(或real)
* 定点数:实际以字符串形式存放，可精确保存数据，在MYSQL默认SQLMode中插入数据精度大于实际定义精度会发生警告，数据按照实际精度四舍五入插入，在MYSQL中包括decimal(或numberic)
> 注：浮点数存在误差，应尽量避免比较，对于货币等一些精度要求较高的场景，应尽量使用定点数表示存储

###　日期类型选择
* 日期类型主要包括DATE,TIME,DATETIME,TIMESTAMP：应根据实际选择能够满足需求的最小存储类型，以节约存储，提高表的操作效率；
* 注意TIMESTAMP表示的日期范围比DATETIME要短得多，如果时间表示涉及不同的时区，应使用TIMESTAMP
