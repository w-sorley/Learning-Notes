---
title: Spring实战学习之Spring核心
date: 2017-05-26
tags: [Spring，web,FrameWork,IoC,DI]
categories:　FrameWork
---

# Spring实战学习之Spring核心

## Spring概述
* Spring是为了解决企业级应用开发的复杂性而创建的，相对于EJB来说, Spring 提供了更加轻量级和简单的编程模型。它增强了简单老式Java对象(Plain Old Java object ,POJO)的功能,使其具备了之前只有EJB和其他企业级Java规范才具有的功能．核心思想是采用了依赖注入( Dependency Injection , DI )和面向切面编程( Aspect-Oriented Programming , AOP )的理念．
* 同时spring不仅仅局限于服务器端开发,任何 Java 应用都能在简单性、可测试性和松耦合等方面从Spring中获益。 Spring也正在涉足其他创新的领域，如移动开发、社交 API 集成、 NoSQL 数据库、云计算以及大数据．
* 为了降低 Java 开发的复杂性, Spring 采取了以下 4 种关键策略:
<br/>　　　－1.基于 POJO 的轻量级和最小侵入性编程;
<br/>　　　－2.通过依赖注入和面向接口实现松耦合;
<br/>　　　－3.基于切面和惯例进行声明式编程;
<br/>　　　－4.通过切面和模板减少样板式代码。
## Spring关键策略
### 最小侵入性编程,激发 POJO 的潜能：
* Spring 不会强迫你实现 Spring 规范的接口或继承 Spring 规范的类,意味着这个类在Spring应用和非 Spring 应用中都可以发挥同样的作用,通过 DI 装配,对象彼此之间保持松散耦合.




### 依赖注入
* 多个类相互之间进行协作来完成特定的业务逻辑。按照传统的做法,每个对象负责管理与自己相互协作的对象(即它所依赖的对象)的引用,这将会导致高度耦合和难以测试的代码．
* 耦合具有两面性( two-headed beast )：
<br/>　　　－一方面,紧密耦合的代码难以测试、难以复用、难以理解,并且典型地表现出 “ 打地鼠 ” 式的 bug 特性(修复一个 bug ,将会出现一个或者更多新的 bug )。
<br/>　　　－另一方面,一定程度的耦合又是必须的，不同的类必须以适当的方式进行交互；
* 通过 DI ,对象的依赖关系将由系统中负责协调各对象的第三方组件在创建对象的时候进行设定。对象无需自行创建或管理它们的依赖关系，依赖关系将被自动注入到需要它们的对象当中去,实现对象间的松耦合．<br/>(对依赖进行替换的一个最常用方法就是在测试的时候使用 mock 实现可以使用 mock 框架 Mockito 去创建一个接口的 mock 实现（））
* 依赖注入的方式之一,即构造器注入( constructor injection )，传入所有探险任务都必须实现的一个接口，对象能够响应任意的接口实现
* 创建应用组件之间协作的行为通常称为装配( wiring )，采用 XML 是很常见的一种装配方式：
```
/*将id２(实现具体接口)对象注入到id1的依赖中*/
＜bean id="bean_id1" class="class_path"＞    //被注入对象bean
＜constructor-arg ref="Dependence_Injection_id"/＞  //需要的注入bean
</bean>
＜bean id="bean_id2" class="class_path"＞    //注入对象bean
＜constructor-arg value="constructor_args"/＞ //构造器传入参数
</bean>
```
* 此外Spring 还支持使用 Java 来描述配置:
```
/*dependenced_bean对象依赖interface_bean接口*/
@Configuration
public class ConfigClassName{
    ＠Bean
    public dependenced_bean dependenced(){
        return new dependenced_bean( interface_bean());  //被注入bean
    }
    @Bean
    public interface_bean interface_bean(){
        return new achieved_interface_bean(constructor args);  //注入的bean
    }
}
```
* Spring 通过应用上下文( Application Context )装载 bean 的定义并把它们组装起来(自带了多种应用上下文实现,主要的区别仅仅在于如何加载配置)
如ClassPathXmlApplicationContext该类加载位于应用程序类路径下的一个或多个 XML 配置文件:
```
ClassPathXmlApplicationContext context=new ClassPathXmlApplicationContext（"xml_file_path"）  //加载Spring上下文
POJO pojo_name=context.getBean(POJO.Class)　　／／获取对象实例
context.close()
```
### 应用切面
* 面向切面编程( aspect-oriented programming , AOP )允许你把遍布应用各处的功能分离出来形成可重用的组件.
* 系统由许多不同的组件组成,每一个组件各负责一块特定功能。除了实现自身核心的功能之外,这些组件还经常承担着额外的职责。诸如日志、事务管理和安全这样的系统服务经常融入到自身具有核心业务逻辑的组件中去,这些系统服务通常被称为横切关注点,因为它们会跨越系统的多个组件,这样将会带来双重的复杂性．
* 而AOP 能够使这些服务模块化,并以声明的方式将它们应用到它们需要影响的组件中去。所造成的结果就是这些组件会具有更高的内聚性并且会更加关注自身的业务,完全不需要了解涉及系统服务所带来复杂性，确保 POJO 的简单性.
> （可以把切面想象为覆盖在很多组件之上的一个外壳，借助 AOP ,可以使用各种功能层去包裹核心业务层）
* 切面声明：
```
＜bean id="aspect_bean_id" class="class_path"＞   //声明切面作为普通的bean
<constructor-arg value＝"constructor args">
</bean>
<aop:config>
<aop:aspect ref="aspect_bean_id">                 
<aop:pointcut id="pointcut_id" expression="execution(*,*,method_name)＂ />   //定义切点
<aop:before pintcut-ref="pointcut_id" method="aspect_bean_method_name">　　　　//声明前置通知
<aop:after pintcut-ref="pointcut_id" method="aspect_bean_method_name">　　　　//声明后置通知
</aop:aspect>
</aop:config>
```
### 使用模板消除样板式代码
* 通常为了实现通用的和简单的任务,不得不编写一些相似的样板式的代码( boilerplate code )，Spring 旨在通过模板封装来消除样板式代码，使程序仅关注于核心逻辑,
* 如数据库操作模板JdbcTemplate:
```
jdbcTemplate.queryForObject("sql_string",new RowMapper<Object>(){...},sql_paramater)   //查询数据库，并转化为对象实例
```

## Spring 容器( container )
* 在基于Spring的应用中,应用对象生存于Spring容器( container )中,由Spring 容器负责创建对象,装配它们,配置它们并管理它们的整个生命周期．
* 容器是 Spring 框架的核心，使用 DI 管理构成应用的组件，会创建相互协作的组件之间的关联
* Spring 自带了多个容器实现，可以归为两种不同的类型。
<br/>　　　－bean 工厂(由 org.springframework. beans.factory.eanFactory 接口定义)是最简单的容器,提供基本的 DI 支持。　
<br/>　　　－应用上下文(由 org.springframework.context.ApplicationContext 接口定义)基于 BeanFactory 构建,并提供应用框架级别的服务,例如从属性文件解析文本信息以及发布应用事件给感兴趣的事件监听者。(使用更广泛)

> Spring 自带了多种类型的应用上下文，如:
* AnnotationConfigApplicationContext :从一个或多个基于 Java 的配置类中加载 Spring 应用上下文。
* AnnotationConfigWebApplicationContext :从一个或多个基于 Java 的配置类中加载 Spring Web 应用上下文。
* ClassPathXmlApplicationContext :从类路径下的一个或多个 XML 配置文件中加载上下文定义,把应用上下文的定义文件作为类资源。
* FileSystemXmlapplicationcontext :从文件系统下的一个或多个 XML 配置文件中加载上下文定义。
* XmlWebApplicationContext :从 Web 应用下的一个或多个 XML 配置文件中加载上下文定义。

### bean的生命周期:
* 1 . Spring 对 bean 进行实例化;
* 2 . Spring 将值和 bean 的引用注入到 bean 对应的属性中;
* 3 .如果 bean 实现了 BeanNameAware 接口, Spring 将 bean 的 ID 传递给 setBean-Name() 方法;
* 4 .如果 bean 实现了 BeanFactoryAware 接口, Spring 将调用 setBeanFactory() 方法,将 BeanFactory 容器实例传入;
* 5 .如果 bean 实现了 ApplicationContextAware 接口, Spring 将调用 setApplicationContext() 方法,将 bean 所在的应用上下文的引用传入进来;
* 6 .如果 bean 实现了 BeanPostProcessor 接口, Spring 将调用它们的 post-ProcessBeforeInitialization() 方法;
* 7 .如果 bean 实现了 InitializingBean 接口, Spring 将调用它们的 after-PropertiesSet() 方法。类似地,如果 bean 使用 init-method 声明了初始化方法,该方法也会被调用;
* 8 .如果 bean 实现了 BeanPostProcessor 接口, Spring 将调用它们的 post-ProcessAfterInitialization() 方法;
* 9 .此时, bean 已经准备就绪,可以被应用程序使用了,它们将一直驻留在应用上下文中,直到该应用上下文被销毁;
* 10 .如果 bean 实现了 DisposableBean 接口, Spring 将调用它的 destroy() 接口方法。同样,如果 bean 使用 destroy-method 声明了销毁方法,该方法也会被调用

## Spring框架:
* Spring框架不仅关注于通过DI 、 AOP和消除样板式代码来简化企业级Java开发，在Spring框架之外还存在一个构建在核心框架之上的庞大生态圈,它将Spring扩展到不同的领域,例如Web服务、REST移动开发以及 NoSQL :
### Spring主要包括6个主要模块分类:
* 数据访问与集成(JDBC,Transaction,ORM,OXM,Messaging,JMS):JDBC 和DAO ( Data Access Object )模块抽象了数据库访问样板式代码，同时在多种数据库服务的错误信息之上构建了一个语义丰富的异常层；ORM 模块对许多流行的ORM 框架进行了集成；同时包含了在JMS(Java Message Service )之上构建的 Spring 抽象层,它会使用消息以异步的方式与其他应用集成；除此之外,本模块会使Spring AOP 模块为 Spring 应用中的对象提供事务管理服务
* Web与远程调用(Web,Web servlet,Web portlet,WebSocket)：自带MVC框架,有助于在 Web 层提升应用的松耦合水平．除了面向用户的Web应用,该模块还提供了多种构建与其他应用交互的远程调用方案
* 面向切面编程(AOP,Aspects):对面向切面编程提供了丰富的支持,是 Spring 应用系统中开发切面的基础
* Instrumentation(Instrument,Instrument Tomcat):提供了为 JVM 添加代理( agent )的功能
* Spring核心容器(Beans,Core,Context,Expression,Context support):除了bean工厂和应用上下文,该模块也提供了许多企业服务,例如E-mail、JNDI访问,EJB 集成和调度.所有的Spring模块都构建于核心容器之上
* 测试（text）:为使用 JNDI 、 Servlet 和 Portlet 编写单元测试提供了一系列的 mock 对象实现,为加载 Spring应用上下文中的 bean 集合以及与 Spring 上下文中的 bean 进行交互提供了支持

### Spring Portfolio
> 整个 Spring Portfolio 包括多个构建于核心 Spring 框架之上的框架和类库,几乎为每一个领域的 Java 开发都提供了 Spring 编程模型
* Spring Web Flow:建立于 Spring MVC 框架之上,它为基于流程的会话式 Web 应用提供了支持,[参考](http://projects.spring.io/spring-webflow/)   
* Spring Web Service:提供了契约优先的 Web Service 模型,服务的实现都是为了满足服务的契约而编写的.[参考](http://docs.spring.io/spring- ws/site/) 
* Spring Security:为 Spring 应用提供了声明式的安全机制.[参考](http://projects.spring.io/spring-security/)
* Spring Integration：提供了多种通用应用集成模式的 Spring 声明式风格实现．[参考](http://projects.spring.io/spring-integration/)
* Spring Batch:为批处理应用开发提供面向 POJO 的编程模型．[参考](http://projects.spring.io/ spring-batch/)
* Spring Data：使得在 Spring 中使用任何数据库都变得非常容易，为多种数据库类型提供了一种自动化的 Repository 机制,负责为你创建 Repository 的实现
* Spring Socia:一个社交网络扩展模块,更多的是关注连接( connect )[参考-facebok](https://spring.io/guides/gs/accessing- facebook/)/[参考－twitter](https://spring.io/guides/gs/accessing-twitter/)
* Spring Mobile：是 Spring MVC 新的扩展模块,用于支持移动 Web 应用开发．
* Spring for Android:旨在通过 Spring 框架为开发基于 Android 设备的本地应用提供某些简单的支持.[参考](http://projects.spring.io /spring-android/)
* Spring Boot：以 Spring 的视角,致力于简化 Spring 本身，大量依赖于自动配置技术,它能够消除大部分(在很多场景中,甚至是全部) Spring 配置．
