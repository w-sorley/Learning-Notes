---
title: Spring实战之JDBC操作数据库
date: 2017-06-30
tags: [JDBC，DataBase,Spring]
categories: Framework
---

# Spring实战之通过Spring和JDBC操作数据库

Spring自带一组数据访问框架,集成了多种数据访问技术，在通过JDBC或Hibernate等对象关系映射(object-relational mapping,ORM )框架实现数据持久化时，Spring能够消除持久化代码中单调枯燥的数据访问逻辑。
Spring 的数据访问
将底层的数据访问功能集成放在数据访问对象DAO(又称Repository)中，对外提供数据访问接口。从而实现服务与数据访问实现的解绑定，将数据访问的持久化技术隐藏在接口后，易于测试，便于以后的修改，更加灵活性

为了将数据访问层与应用程序的其他部分隔离开来,Spring采用的方式之一就是提供统一的异常体系
Spring 的数据访问异常体系
编写过JDBC代码时一般需要强制捕获SQLException，可能导致抛出 SQLException 的常见问题包括:
应用程序无法连接数据库;应用程序无法连接数据库;
要执行的查询存在语法错误;
查询中所使用的表和 / 或列不存在;
试图插入或更新的数据违反了数据库约束。
SQLException 被视为处理数据访问所有问题的通用异常，所有的数据访问问题都会抛出 SQLException，一些持久化框架提供了相对丰富的异常体系，分别对应于特定的数据访问问题。这样就可以针对想
处理的异常编写 catch 代码块。Hibernate 的异常是其本身所特有的，无法实现将特定的持久化机制独立于数据访问层。因此需要的数据访问异常要具有描述性而且又与特定的持久化框架无关

Spring 所提供的平台无关的持久化异常
Spring 为读取和写入数据库的几乎所有错误都提供了异常，如：
JDBC 的异常
BatchUpdateException
DataTruncation
SQLException
SQLWarning
Spring 的数据访问异常
BadSqlGrammarException
CannotAcquireLockException
CannotSerializeTransactionException
CannotGetJdbcConnectionException
CleanupFailureDataAccessException
ConcurrencyFailureException
DataAccessException
DataAccessResourceFailureException
DataIntegrityViolationException
DataRetrievalFailureException
DataSourceLookupApiUsageException
DeadlockLoserDataAccessException
DuplicateKeyException
EmptyResultDataAccessException
IncorrectResultSizeDataAccessException
IncorrectUpdateSemanticsDataAccessException
InvalidDataAccessApiUsageException
InvalidDataAccessResourceUsageException
InvalidResultSetAccessException
JdbcUpdateAffectedIncorrectNumberOfRowsException
LbRetrievalFailureException

JDBC 的异常
BatchUpdateException
DataTruncation
SQLException
SQLWarning
Spring 的数据访问异常
NonTransientDataAccessResourceException
OptimisticLockingFailureException
PermissionDeniedDataAccessException
PessimisticLockingFailureException
QueryTimeoutException
RecoverableDataAccessException
SQLWarningException
SqlXmlFeatureNotImplementedException
TransientDataAccessException
TransientDataAccessResourceException
TypeMismatchDataAccessException
UncategorizedDataAccessException
UncategorizedSQLException

这些异常都继承自 DataAccessException它是一个非检查型异常，即没有必要捕获 Spring 所抛出的数据访问异常
为了利用 Spring 的数据访问异常,我们必须使用 Spring 所支持的数据访问模板
数据访问模板化

使用了模板方法模式，模板方法将过程中与特定实现相关的部分委托给接口,而这个接口的不同实现定
义了过程中的具体行为，Spring 将数据访问过程中固定的和可变的部分明确划分为两个不同的类:模板( template )和回调( callback )模板管理过程中固定的部分(事务控制、管理资源以及处理异常准备资源，开始事务，提高回滚事务，关闭资源处理错误等),而回调处理自定义的数据访问代码(语句、绑定参数以及整理结果集事务执行，返回数据)

针对不同的持久化平台, Spring 提供了多个可选的模板
jca.cci.core.CciTemplate JCA CCI 连接
jdbc.core.JdbcTemplate JDBC 连接
jdbc.core.namedparam.NamedParameterJdbcTemplate 支持命名参数的 JDBC 连接
jdbc.core.simple.SimpleJdbcTemplate 通过 Java 5 简化后的 JDBC 连接( Spring 3.1 中已经废弃)
orm.hibernate3.HibernateTemplate Hibernate 3.x 以上的 Session
orm.ibatis.SqlMapClientTemplate iBATIS SqlMap 客户端
orm.jdo.JdoTemplate Java 数据对象( Java Data Object )实现
orm.jpa.JpaTemplate Java 持久化 API 的实体管理器


配置数据源
Spring 所支持的大多数持久化功能都依赖于数据源。因此,在声明模板和 Repository 之前,我们需要在 Spring 中配置一个数据源用来连接数据库，Spring 提供了在 Spring 上下文中配置数据源 bean 的多种方式：
通过 JDBC 驱动程序定义的数据源;
通过 JNDI 查找的数据源;
连接池的数据源

使用 JNDI 数据源

好处在于数据源完全可以在应用程序之外进行管理，此外在应用服务器中管理的数据源通常以池的方式组织,从而具备更好的性能,并且还支持系统管理员对其进行热切换。
<jee:jndi-lookup> 元素可以用于检索 JNDI 中的任何对象(包括数据源)并将其作为 Spring 的 bean
```
<jee:jndi-lookup id="" jndi-name="指定 JNDI 中资源的名称" resource-ref="Java 应用设置为 true给定的 jndi-name 将会自动添加 “java:comp/env/” 前缀" />
```
或借助 JndiObjectFactoryBean使用 Java 配置
```
@Bean
public JndiObjectFactoryBean dataSource(){
    JndiObjectFactoryBean jofb=new JndiObjectFactoryBean();
    jofb.setJndiName("..");
    jofb.setResourceRef(true/false);
    jofb.setProxyInterface(javax.sql.DataSource.class);
    return jofb;
}

```

使用数据源连接池

Spring 并没有提供数据源连接池实现,但是有多项可用的方案(在一定程度上与 Spring 自带的DriverManagerDataSource或SingleConnectionDataSource很类似),如：
[Apache Commons DBCP](http://jakarta.apache.org/commons/dbcp) ;
[c3p0](http://sourceforge.net/projects/c3p0/) ;
[BoneCP](http://jolbox.com/)
如配置DBCP BasicDataSource的方式:
```
<bean id="datasource" class="org.apache.commons.dbcp.BasicDataSource"
p:DriverClassName="org.h2.Driver"
p:url="jdbc:h2:tcp://localhost/../database"
p:username=".."
p:password=".."
p:initialSize=".."
p:maxActive=".."  />
```
或使用java配置：
```
@Bean
public BasicDataSource datesource(){
    BasicDataSource ds=new BasicDataSource();
    ds.setDriverClassName("..");
    ds.setUrl("..");
    ds.setUserName("..");
    ds.setPassword("..");
    ds.InitialSize(..);
    ds.setMaxActive(..);
    return ds;
}
```
BasicDataSource 的池配置属性
initialSize 池启动时创建的连接数量
maxActive 同一时间可从池中分配的最多连接数。如果设置为 0 ,表示无限制
maxIdle 池里不会被释放的最多空闲连接数。如果设置为 0 ,表示无限制
maxOpenPreparedStatements 在同一时间能够从语句池中分配的预处理语句( prepared statement )的最大数量。如果设置为 0 ,表示无限制
maxWait 在抛出异常之前,池等待连接回收的最大时间(当没有可用连接时)。如果设置为 -1 ,表示无限等待
minEvictableIdleTimeMillis 连接在池中保持空闲而不被回收的最大时间
minIdle 在不创建新连接的情况下,池中保持空闲的最小连接数
poolPreparedStatements 是否对预处理语句( prepared statement )进行池管理(布尔值)


基于 JDBC 驱动的数据源
Spring 提供了三个JDBC驱动定义数据源(均位于org.springframework.jdbc.datasource包中,配置与 DBCPBasicDataSource 的配置类似):
DriverManagerDataSource :在每个连接请求时都会返回一个新建的连接。与 DBCP 的 BasicDataSource 不同,由 DriverManagerDataSource 提供的连接并没有进行池化管理;
SimpleDriverDataSource :与 DriverManagerDataSource 的工作方式类似,但是它直接使用 JDBC 驱动,来解决在特定环境下的类加载问题,这样的环境包括 OSGi 容器;
SingleConnectionDataSource :在每个连接请求时都会返回同一个的连接。尽管 SingleConnectionDataSource 不是严格意义上的连接池数据源,但是你可以将其视为只有一个连接的池。



使用嵌入式的数据源

嵌入式数据库作为应用的一部分运行,而不是应用连接的独立数据库服务器,适用于对于开发和测试场景(每次重启应用或运行测试的时候,都能够重新填充测试数据)
如使用jdbc命名空间来配置嵌入式的H2数据库：
```
<jdbc:embedded-database id="datasource" type="H2">  //要确保 H2 位于应用的类路径下,此外也可为DERBY使用嵌入式的 Apache Derby 数据库
<jdbc:scripts location="/path/*.sql"/>  //可以配置多个<jdbc:script>元素来搭建数据库
...
</jdbc:embedded-database>

```
或基于Java配置嵌入式数据库，使用 EmbeddedDatabaseBuilder 来构建 DataSource
```
@Bean
public DataSource datasource(){
    return new EmbeddedDatabaseBuilder()
                                    .setType(EmbeddedDtabaseType.H2)
                                    addScripts(classpath:file.sql)
                                    ....
                                    .build();
}

```

使用 profile 选择数据源

将每个数据源配置在不同的 profile 中,使用与不同阶段不同场景，从而能够在运行时选择数据源：
```

```





