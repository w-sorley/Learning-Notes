---
title: Spring实战之高级装配
date:　2017-05-30
tags: [Web,Spring,FrameWork]
categories: FrameWork
---
<br/>
<br/>
＃ Spring实战之高级装配

## 环境与profile
* 在软件开发过程中,会涉及将应用程序从一个环境迁移到另外一个环境的问题,为保证在不同环境下程序都能够正确运行,可能会需要针对环境进行不同的配置．
* 其中一种方式就是在单独的配置类(或XML文件)中配置每个bean,然后在构建阶段(可能会使用Maven的profiles),针对不同的环境确定要将哪一个配置编译到可部署的应用中.
> （同时为每种环境重新构建应用的方式会增加工作量，不够灵活)

* 另外可以利用Spring的特性:运行时根据环境(激活配置)决定该创建哪个bean,从而同一个部署单元(可能会是WAR文件)能够适用于所有的环境,没有必要进行重新构建(由激活配置针对不同环境，决定创建不同的bean)
### bean-profile：
* 在3.1版本中,Spring引入了bean　profile功能:首先将所有不同的bean定义整理到一个或多个profile之中,在将应用部署到每个环境时,要确保对应的profile处于激活(active)的状态。
* 使用@Profile注解指定某个bean属于哪一个profile,告诉Spring这个配置类中的bean只有在对应的profile激活时才会创建(同时没有指定 profile 的 bean 始终都会被创建） 
(注：在Spring 3.1中,只能在类级别上使用@Profile注解,但从Spring 3.2开始,也可以在方法级别上使用@Profile注解)

### 在XML中配置 profile
在XML中可以通过<beans>元素的 profile属性配置profile bean：<beans profile="...">
同时支持在根<beans>元素中嵌套定义<beans>元素,从而一个profile XML文件中可定义多个环境的配置

### 激活profile
* Spring在确定哪个profile处于激活状态时,需要依赖两个独立的属性: spring.profiles.active 和 spring.profiles.default
<br/>　　－1.通过设置spring.profiles.active属性,确定哪个profile激活；
<br/>　　－2.如果没有设置 spring.profiles.active 属性的话,Spring将会查找spring.profiles.default的值,确定哪个profile激活
<br/>　　－3.如果spring.profiles.active和spring.profiles.default均没有设置的话,则没有激活的profile,因此只会创建那些没有定义在profile中的bean 。
* 同时有多种方式来设置spring.profiles.active 和 spring.profiles.default这两个属性:
<br/>　　－1.作为DispatcherServlet的初始化参数;
<br/>　　－2.作为Web应用的上下文参数;
<br/>　　－3.作为JNDI条目;
<br/>　　－4.作为环境变量;
<br/>　　－5.作为JVM的系统属性;
<br/>　　－6.在集成测试类上,使用@ActiveProfiles注解设置
* 如在Web应用中,通过web.xml设置 spring.profiles.default作为Web应用的上下文参数（可以同时激活多个 profile，以逗号分隔）
```
<context-param>
  <param-name>spring.profiles.default</param-name>
  <param-value>...</param-value>
</context-param>
<servlet>
  <init-param>
    <param-name>spring.profiles.default<param-name>
    <param-value>...</param-value>
  </init-param>
</servlet>

```
* 如使用profile进行测试：Spring 提供 @ActiveProfiles 注解，可用来指定运行测试时要激活哪个 profile

### 条件化的bean
* Spring4引入了新的@Conditional注解,可以用到带有@Bean注解的方法上：如果给定的条件计算结果为true,就会创建这个bean,否则的话,这个bean会被忽略；
> @Conditional将会通过Condition接口进行条件对比，设置给@Conditional的类可以是任意实现了Condition接口的类型，只需提供matches()，方法返回true,那么就会创建带有 @Conditional 注解的bean<br/>
> 同时从Spring4开始, @Profile 注解也进行了重构,使其基于 @Conditional 和 Condition 实现    

## 处理自动装配的歧义性
* 当Spring试图自动装配依赖参数时，若没有唯一、无歧义的可选值,出现自动装配的歧义性，Spring会抛出NoUniqueBeanDefinitionException异常
* 当确实发生歧义性的时候,Spring提供了多种可选方案来解决这样的问题。
<br/>　　－1.可以将可选bean中的某一个设为首选(primary)的bean;
<br/>　　－2.使用限定符(qualifier),令Spring将可选的bean的范围缩小到只有一个bean

### 标示首选的bean：
* 1.通过@Primary与@Component组合,用在组件扫描的bean上/也可以与@Bean组合用在Java配置的bean声明中;
* 2.通过xml配置文件中<bean primary="true"> 元素primary 属性用来指定首选的 bean;
### 限定自动装配的bean
* Spring的限定符能够在所有可选的bean上进行缩小范围的操作,最终能够达到只有一个 bean满足所规定的限制条件
* @Qualifier注解是使用限定符的主要方式。可以与@Autowired和@Inject协同使用,在注入的时候指定想要注入进去的是哪个bean
* 用法:@Qualifier("bead_id") 所引用的bean要具有String类型限定符(所有的bean都会给定一个默认的限定符,与bean的ID相同)
> 注意:同时注入方法上所指定的限定符与要注入的bean的名称是紧耦合的。对类名称的任意改动都会导致限定符失效

#### 创建自定义的限定符
* bean声明上添加@Qualifier注解可为bean设置自己的限定符,而不是依赖于将bean ID作为限定符
* 由于Java不允许在同一个条目上重复出现相同类型的多个注解，但可以创建自定义的限定符注解(本身要使用 @Qualifier 注解来标注)
> (注意:Java 8 允许出现重复的注解,只要这个注解本身在定义的时候带有 @Repeatable 注解就可以。不过, Spring 的 @Qualifier 注解并没有在定义时添加 @Repeatable 注解)
## bean的作用域
* 在默认情况下, Spring应用上下文中所有bean都是作为以单例(singleton)的形式创建的
* 但有时所使用的类可能是易变的(mutable),它们会保持一些状态,因此重用是不安全的,对象会被污染,出现意想不到的问题,因此需要指定bean的作用域；
* Spring定义了多种作用域,可以基于这些作用域创建bean,包括：
<br/>　　－1.单例(Singleton):在整个应用中,只创建 bean 的一个实例。
<br/>　　－2.原型(Prototype):每次注入或者通过Spring应用上下文获取的时候,都会创建一个新的bean实例:@Scope("prototype")或传入ConfigurableBeanFactory．SCOPE_PROTOTYPE常量； 
<br/>　　－3.会话(Session):在Web应用中,为每个会话创建一个bean实例:利用WebApplicationContext．SCOPE_SESSION常量指定；
<br/>　　－4.请求(Rquest):在Web应用中,为每个请求创建一个bean实例
> 注：可以利用@Scope注解,可与@Component或@Bean一起使用来选择其他的作用域(默认是单例)，也可以使用<bean>元素的 scope 属性来设置作用域

* 如使用会话和请求作用域，@Scope的proxyMode属性设置成ScopedProxyMode.INTERFACES ，这个属性解决了将会话或请求作用域的 bean 注入到单例 bean 中所遇到的问题：Spring并不会将实际的bean注入,会注入一个到bean 的代理．这个代理会暴露与Bean相同的方法,当调用代理的方法时,代理会对其进行懒解析并将调用委托给会话作用域内真正的bean；
> 注：(如果注入的对象是接口而不是类,这是可以的(也是最为理想的代理模式)，但如果注入的对象是一个具体的类,Spring没有办法创建基于接口的代理了，此时,必须使用CGLib来生成基于类的代理，所以,如果bean类型是具体类的话,我们必须要将proxyMode属性设置为 ScopedProxyMode.TARGET_CLASS,以此来表明要以生成目标类扩展的方式创建代理（请求作用域的 bean 应该也以作用域代理的方式进行注入）

### 在 XML 中声明作用域代理
* <bean> 元素的 scope 属性能够设置 bean 的作用域，如设置回话作用域:
```
<bean id=".." class=".." scope="..">
<aop:scoped-proxy>
</bean>
```
## 运行时值注入
* 利用构造器注入或xml配置文件时，在实现的时候是将值硬编码在配置类中，有时,可能会希望避免硬编码值,而是想让这些值在运行时再确定,Spring提供了两种在运行时求值的方式:
<br/>　　－1.属性占位符(Property placeholder)。
<br/>　　－2.Spring表达式语言(SpEL)。

### 注入外部的值
* 处理外部值的最简单方式就是声明属性源并通过Spring的Environment来检索属性
```
@PropertySource("properties_file_path")  //声明属性源
public class ClassName{a
    @Autowired
    Environment env;
    @Bean
    public BeanClass basic(){
        return new BeanClass(env.getProperty(".."));      //检索属性值
    }
}
//属性文件会加载到 Spring 的 Environment 中,稍后可以从中检索属性
```
#### 深入学习Spring的Environment
* getProperty() 方法有四个重载的变种形式:
<br/>　　－1.String getProperty(String key)
<br/>　　－2.String getProperty(String key, String defaultValue)
<br/>　　－3.T getProperty(String key, Class<T> type)
<br/>　　－4.T getProperty(String key, Class<T> type, T defaultvalue)

* Environment.containsProperty() 方法,检查一下某个属性是否存在
* getPropertyAsClass() 方法将属性解析为类
* String[] getActiveProfiles() :返回激活 profile 名称的数组;
* String[] getDefaultProfiles() :返回默认 profile 名称的数组;
* boolean acceptsProfiles(String... profiles) :如果 environment 支持给定 profile 的话,就返回 true 。

### 占位符装配属性
* Spring也提供了通过占位符装配属性的方法,这些占位符的值会来源于一个属性源。
* 解析属性占位符:Spring支持将属性定义到外部的属性的文件中,并使用占位符值将其插入到Spring bean 中。在Spring 装配中,占位符的形式为使用 “${
.包装的属性名称.. }” 
* 如果我们依赖于组件扫描和自动装配来创建和初始化应用组件,需要使用@Value注解解析占位符,它的使用方式与@Autowired注解非常相似;
> 为了使用占位符,我们必须要配置PropertyPlaceholderConfigurer bean或PropertySourcesPlaceholderConfigurer　bean
<br/>(从Spring 3.1开始，PropertySourcesPlaceholderConfigurer能够基于 Spring Environment 及其属性源来解析占位符)
<br/>若使用XML配置, Spring context命名空间中的<context:propertyplaceholder>元素将会生成PropertySourcesPlaceholderConfigurer bean
<br/>(解析外部属性能够将值的处理推迟到运行时,但是它的关注点在于根据名称解析来自于 Spring Environment 和属性源的属性)

### 使用Spring表达式语言进行装配
* Spring3引入了Spring表达式语言(Spring Expression Language,SpEL),能够以一种强大和简洁的方式将值装配到 bean 属性和构造器参数中,同时在这个过程中所使用的表达式会在运行时计算得到值
* SpEL拥有很多特性,包括:
* 　－使用 bean 的 ID 来引用 bean ;
* 　－调用方法和访问对象的属性;
* 　－对值进行算术、关系和逻辑运算;
* 　－正则表达式匹配;
* 　－集合操作
* SpEL使用：
> SpEL表达式要放到 “#{ ... }” 之中
```
public construtor_name(              //在构造注入中使用
  @Value("#{class['property']}") Class param, ...)
｛　this.param=param;　｝
//在xml配置中使用
<bean id=".." class=".." c:_param="#{class['property']}" />
```
#### SpEL基础表达式
* 表示字面值:＃{int}/#{double}/#{string}/#{boolean}
* 通过 ID引用 bean 、属性和方法:#{bean_id}/#{bean_id.property_name}/#{bean_id.method_name}
* 在表达式中使用类型:T() 运算符如：T(java.lang.Math)返回一个Class对象，真正价值在于它能够访问目标类型的静态方法和常量，如T(java.lang.Math).PI/T(java.lang.Math).random()
* SpEL运算符:支持常见的算术(支持幂运算^)/逻辑(and,or,not,|)/比较(还支持lt,gt,eq,le,ge)/条件(?:)运算符,支持正则表达(matches 运算符对String 类型的文本(作为左边参数)应用正则匹配)
* 计算集合：支持集合和数组有关操作，提供了查询运算符( .?[] ),它会用来对集合进行过滤,得到集合的一个子集（“.^[]” 和 “.$[]” ,它们分别用来在集合中查询第一个匹配项和最后一个匹配项）


