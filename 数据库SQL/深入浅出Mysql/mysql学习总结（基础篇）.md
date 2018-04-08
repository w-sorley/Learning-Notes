---
title: MYSQL学习总结(基础)
date: 2017-05-22
tags: [SQL,MYSQL,Web,基础]
categories: DataBase
---
# MYSQL基础学习总结
> －－总结自＜深入浅出MYSQL:数据开发，优化与管理维护．第２版＞（基础篇）


## Linux(ubuntu16.04)下Mysql配置：
* 在Linux下通过my.cnf配置文件配置mysql(全局配置文件/etc/mysql/my.cnf; 用户配置文件～/.my.cnf)，具体配置参数可参考mysql官方给出的参考配置文件；

## Mysql常用操作命令
* 启动服务: service mysql start；
* 关闭服务: service mysql stop；
* 重启服务: service mysql restart；
> (此外还可以通过执行/usr/bin/mysqld_safe启动，通过命令mysqladmin -uroot shutdown关闭)
* 登录数据库: mysql -u username -p


## SQL语言基础:
### 数据库定义语言DDL:
* 创建数据库：CREATE DATABASE database_name;
* 切换/使用数据库: USE database_name;
* 删除数据库：DROP DATABASE database_name;

* 创建数据表:CREATE TABLE table_name(column_name_1 data_type constraints, ...., ...,)
* 删除数据表：DROP TABLE `table_name`;
* 修改表名：ALTER TABLE ｀table_name｀ RENAME new_table_name;
* 修改表的字段属性定义：ALTER TABLE `table_name` MODIFY [COLUMN] `column_name` column_definition [FIRST | AFTER `column_name`];
* 在表中添加字段：　ALTER TABLE `table_name` ADD [COLUMN] `column_name` column_definition [FIRST | AFTER `column_name`];
* 删除表中的字段；　ALTER TABLE `table_name` DROP [COLUMN] `column_name` [FIRST | AFTER `column_name`];
* 重新定义表中字段：ALTER TABLE `table_name` CHANGE [COLUMN] `old_column_name` new_column_definition [FIRST | AFTER `column_name`];(属于MYSQL在SQL上的扩展)
> ([FIRST | AFTER column_name]用来确定修改后字段在表中的位置，ADD默认在最后，FIRST表示在开始，属于MYSQL在SQL上的扩展)

* 查看已存在表的定义/结构:DESC table_name;
* 查看已存在表的创建语句:SHOW CREATE TABLE `table_name` \G;(信息更加详细，＼G表示使记录按照字段竖向排列)
* 查看已存在数据库:SHOW DATABASES;
> 其中系统默认数据库:
  *  information_schema:存储数据库对象信息，如:用户表信息，列信息，权限信息，字符集信息，分区信息等；
  *  cluster:存储系统集群信息；
  *  mysql:存储系统的用户权限信息；
  *  test:测试数据库，任何用户可用；
* 查看当前数据库已存在的表：SHOW TABLES;


### 数据库操作语言DML:
* 插入记录：INSERT INTO `table_name`(field1,field2,.....fieldn) VALUES(value1,value2,...,valuen),(....),....; 
> (INSERR:可一次性插入多组记录；没有插入值的字段可能设置为NULL,默认值或自增值；可以不指定字段名称，按字段定义顺序插入)
* 更新记录：UPDATE `table_name｀　SET `field1`=vlaue1,....fieldn=valuen [WHERE condition];
> (可同时更新多个表中记录) UPDATE table1,table2....tablen SET tablen.fieldn=value........[WHERE condition] )
* 删除记录: DELETE FROM `table_name` [WHERE condition]
> ((可同时删除多个表中记录) DELETE table1,table2....tablen FROM table1,table2....tablen [WHERE condition] ))
####＃ 查询记录：SELECT `filed1`,...`fieldn` FROM `table_name` [WHERE condition];
> (*表示所有字段；DISTINCT表示去重后输出)
*  条件查询：WHERE后的约束条件可以利用＝，＞，\<，>=,\<=,!=等比较运算符，也可以利用or,and等逻辑运算符进行多条件联合查询;
*  排序：ORDER　BY `field１` [DESC|ASC]，`field２` [DESC|ASC],...:指定按某个字段排序，DESC降序，默认为ASC升序；可跟多个排序字段，每个字段可有不同排序顺序，前一个字段相同，则按照下个字段排序
*  限制：LIMIT offset_start,rows:offser_start为起始偏移量，默认为０，rows为要显示的行数；（经常和排序配合用于分页显示，数据MYSQL扩展）
*  聚合：SELECT [field1...] func_name FROM `table_name` [WHERE condition] [GROUP BY field1,..fieldn]　[WITH ROLLUP] [HAVING condition]:</br>1.func_name:表示聚合函数：sum求和，count(*)计数，max最大值，min最小值；</br>2.GROUP BY:制定分类聚合字段；</br> 3.WITH ROLLUP:可选，表示是否对分类聚合后的字段再进行汇总;</br>4.HAVING　condition:对分类后的结果进行条件过滤
*  表连接：分为内连接FROM table1,table2...（仅选出两张表中相互匹配的记录）和外连接（外连接又分为左连接LEFT JOIN［包含所有左边表中记录］和右连接RIGHT JOIN[包含所有右边表中记录]）
*  子查询：子查询语句放在()中，具体关键字包括IN,NOT IN,=,!=,EXISTS,NOT EXISTS等
*  记录联合:UNION(去重后显示)或UNION ALL,将多个结果合并在一起显示；


### 数据库控制语言
* 创建用户：CREATE USER 'username'@'host' IDENTIFIED BY 'password';
* 删除用户：DROP USER 'username'@'host';
* 授权：GRANT privileges ON ｀databasename.tablename｀ TO 'username'@'host' ；（可与创建用户合并）
* 设置更改密码：　SET PASSWORD FOR 'username'@'host' = PASSWORD('newpassword');
* 回收授权：REVOKE privilege ON databasename.tablename FROM 'username'@'host';

#####　查阅MYSQL帮助：
* ？　contents:显示可供查询分类，可按照层次查看帮组信息；
* ?  command key:可快速查看某项命令的帮助语法

##### 查看元数据信息：
* information_schema系统数据库中存储了MYSQL的元数据信息(指数据的数据，如表名，列名，列类型，索引名等)，此表为一个虚拟的表，并不存在物理实体，数据表全部为视图，其中：
</br>1.SCHEMATA:提供当前MYSQL实例中所有数据库信息，
</br>2.TABLES:提供了关于数据库中的表信息（包括视图）
</br>3.COLUMNS:提供了表中的详细列信息
</br>4.STATISTICS:提供了关于表索引的信息

## MYSQL当中的数据类型

### 数值类型
* SQL基本数值类型：严格数据类型（INTEGER（缩写INT,四字节）,SMALLINT（两字节）,DECIMAL(M,D)/DEC(M,D)(定点浮点数M+2字节，最大取值范围与DOUBLE相同),NUMERIC）；近似数值数据类型（FLOAT(四字节),REAL,DOUBLE（八字节）,PRECISION）
* MYSQL扩展数值类型:TINYINT(一字节),MEDIUMINT（三字节）,BIGINT(八字节))，BIT(M)(1~8字节)
* MYSQL支持在类型名称后利用小括号指定显示宽度，一般配合zerofill使用（自动为UNSIGNED），如果数值实际宽度大于指定宽度则实际宽度不受影响；
> (如果操作超出类型范围，会发生out of range错误提示；每一个整数类型都有可选的UNSIGNED,指示无符号;此外，整数还有AUTO_INCREMENT属性，每一个表只能指定一列，且为NOT NULL,并定义为PRIMARY　KEY或UNIQUE键)
> (定点数DECIMAL(M,D)在内部以字符串形式存储，表示一共M位数字，D位位于小数点后，更加精确，精度超出时进行四舍五入［出现截断warning，传统SQLMode出现错误］，默认精度M=10,标度D=0)，注意：在浮点数后加(M,D)为非标准用法；
＞　（BIT(M)类型存储多位二进制，M=1~64,默认为１，直接SELECT不会显示，需要使用bin()或hex()函数进行读取，插入时首先转换为二进制，位数超出出现错误；
### 日期时间类型:
* DATE：4字节，范围(年－月－日)：1000-01-01 ~ 9999-12-31
* DATETIME:８字节，范围(年－月－日　时：分：秒):1000-01-01 00:00:00 ~ 9999-12-31 23:59:59
* TIMESTAMP：４字节，范围:19700101080001 ~ 2038年某个时刻 (第一个默认值为当前系统时间（只有一个，第二个默认值为０），适合需要经常插入或更新为当前系统时间的场景，默认返回DATETIME格式， 如需返回数值应添加＂＋０＂)
* TIME：３字节，范围：-838:59:59 ~ 838: 59:59
* YEAR：１字节，范围:分为两位格式(70~99)和默认的四位格式（1901 ~ 2155）
> （默认SQLMode下，超出范围，错误提示，并以对应格式零值存储；TIMESTAMP与时区相关，插入时先转换为本地时区存放，取出时转换为本地时区后显示）
> 日期格式：任意标点符分割模式，（[YY]YY.[M]M.[D]D [H]H.[M]M.[S]S）；无间隔字符串:[YY]YYMMDDHHMMSS;数值格式：[YY]YYMMDDHHMMSS/[YY]YYMMDD;NOW(),CURRENT_DATE返回当前时间；

### 字符串类型
* CHAR(M):M=0~255指定字节数，长度固定为M字节,插入时不保留尾部空格;
* VARCHAR(M):M=0~65535指定，可边长字符串，占用０～M+1个字节（５.0.3以前版本M最大为255）,插入时保留尾部空格;
* BINARY:二进制字符串，与CHAR类似，插入时最后填充零字节0x00以达到字段定义长度;
* VARBINARY:二进制字符串，与VARCHAR类似，插入时最后填充零字节0x00以达到字段定义长度;
* ENUM:枚举类型，取值范围创建时通过枚举方式指定，最多允许65535个成员，０-255个成员需要一个字节　255-65535个成员需要２个字节，插入时忽略大小写，统一为大写，对于插入不在枚举范围的值不会报错，而是插入定义时的第一个值
* SET类型:与EMUM类似，包含０~64个成员，１－８个成员占１个字节，９－１６个成员占２个字节，以此逐渐增加字节数，取值可以选择其中多个元素组合，超出允许值范围不可以插入，插入时会进行去重；
* 此外字符串类型还有：TINYBLOB,BLOB,MEDIUMBLOB,LONGBLOB,TINYTEXT,TEXT,MEDIUMTEXT,LONGTEXT


## MYSQL中的运算符:
* 算术运算符:加（+）,减(-)乘(*)除(/,DIV)取余(%,MOD);(注：在除法和取余运算中，当除数为０时，除数非法，均返回NULL)
* 比较运算符：等于(=),不等于(<>,!=),NULL安全的等于(<=>)，小于(<),小于或等于(<=)，大于(>)，大于或等于(>=),存在于指定范围(BETWEEN)，存在于指定集合(IN)，为NULL(IS NULL),不为NULL*(IS NOT NULL),通配符匹配(LIKE),正则表达式匹配(REGEXP,RLIKE)
> (执行SELECT语句时，可以对操作数进行比较，结果为真返回１，结果为假返回０，结果不确定返回NULL；可以用于比较数字，字符串或表达式；其中数字以浮点数进行比较，字符串不区分大小写)
* 逻辑运算符(又称为布尔运算符)：与(AND,&&),或(OR,||),非(NOT,!),异或(XOR)
* 位运算符:按位与(&),按位或(|),按位取反(~),按位异或(^),位右移(>>),位左移(<<)
> (将给定操作数转化位二进制，将二进制的每一位进行逻辑运算，然后将得到的二进制结果转化为十进制即为位运算的结果)
> (多种运算符混合参与运算，按照指定优先级确定运算顺序)

## MYSQL中常用函数:
### 字符串函数
* CONCAT(s1,s2,...sn):字符串连接(任何字符串与NULL连接，结果为NULL)
* INSERT(str,x,y,instr):将字符串str的第x位开始y个字符长度的子串替换为instr
* UPPER(str):所有字符变为大写
* LOWER(str):所有字符变为小写
* LEFT(str,x):返回字符串的最左面的x个字符；
* RIGHT(str,x):返回字符串的最右面的x个字符
* LPAD(str,n,pad):用字符串pad对字符串str左面进行填充直到长度为n个字符
* RPAD(str,n,pad):用字符串pad对字符串str右面进行填充直到长度为n个字符
* LTRIM(str):去掉左侧的空格
* RTRIM(str):去掉行尾的空格
* REPEAT(str,x)：返回str重复x次的结果
* REPLACE(str,a,b):用字符串b替换字符串中所有出现的a
* STRCMP(s1,s2):比较字符串s1和s2(比较ASCII码值大小)
* TRIM(str):去掉字符串行尾和行头的空格
* SUBSTRING(str,x,y):返回字符串从x位置起y个字符长度的子字符串；


### 数值函数：
* ABS(x):返回绝对值
* CEIL(x):返回大于x的最小整数值
* FLOOR(x):返回小于x的最大整数值
* MOD(x,y):返回x/ｙ的模
* RAND():返回0-1内的随机值
* ROUND(x,y):返回x的四舍五入有y位小数的值
* TRUNCATE(x,y):返回数字x被截断为y位小数的结果

### 日期和时间函数
* CURDATE()：返回当前日期
* CURTIME():返回当前时间
* NOW():放回当前日期和时间
* UNIX_TIMESTAMP(date):返回日期date的UNIX时间戳
* FROM_UNIXTIME:返回UNIX时间戳的日期值
* WEEK(date);返回日期date为一年中的第几周
* YEAR(date):返回date的年份
* HOUR(time):返回time的小时值
* MINUTE(time):返回time的分钟值
* MONTHNAME(date):返回date的月份名
* DATE_FORMAT(date,fmt):返回字符串fmt格式的date值
* DATE_ADD(date,INTERVAL expr type):返回一个日期或时间值加上一个时间间隔的时间值(INTERVAL间隔关键字，expr间隔表达指示后面间隔的含义)
* DATEDIFF(expr,expr2):返回起始时间expr和结束时间expr2之间的天数
#### Mysql日期和时间格式说明符：
* %S和％s:两位数字形式的秒
* %i:两位数字形式的分钟
* %H:两位数字形式的小时，24小时
* %h或%I:两位数字形式的小时,12小时
* %k:数字形式的小时,24小时
* %l:数字形式的小时,12小时
* %T:24小时的时间形式(hh:mm:ss)
* %r:12小时的时间形式（hh:mm:ss AM/PM）
* %p:AM或PM
* %W:一周中每一天的名称
* %a:一周中每一天的名称的缩写
* %d:两位数字形式表示月中天数(01,02....31)
* %e:数字形式表示月中的天数（１,2....31）
* %D：英文后缀表示月中天数（１st,2nd,3rd.....）
* %w：以数字形式表示周中的天数(0=Sundaty...)
* %j:以三位数字表示年中的天数
* %U:周(0,1,52)，其中Sunday为周中的第一天
* %u:周(0,1,52),其中Monday为周中的第一天
* %M:月名
* %b:缩写的月名
* %m:两位数字表示的月份
* %c：数字表示的月份
* %Y:四位数字表示的年份
* %y:两位数字表示的年份
* ％％：表示字符值‘％’
### 流程函数
* IF(value,t f):如果value值为真返回t，否则为假返回f
* IFNULL(value1,value2):如果value1的值不为NUL,返回value1，否则value1为NULL返回value2
* CASE WHEN [value1] THEN[result1]...ELSE[default] END:如果value1的值为真返回result1，否则返回default
* CASE [expr]WHEN[value1]THEN[result1]...ELSE[dafault]END:如果expr的值等于value1则返回result1，否则返回default
### 其他常用函数
* DATABASE():返回当前登录的数据库名
* VERSION():返回当前数据库版本
* USER()：返回当前登录的用户名
* INET_ATON(IP):返回IP地址的数字表示
* INET_NTOA(num):返回数字代表的IP地址
* PASSWORD(str):返回字符串str的加密版本
* MD5():返回字符串str的MD5值

