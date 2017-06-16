---
title: Spring实战之装配Bean
date: 2017-05-28
tags: [JAVA,IoC,DI,Spring,Web,FrameWork]
categories: FrameWork
---
# Spring实战之装配Bean
<br/>
<br/>
## 概述
* 任何一个成功的应用都是由多个为了实现某一个业务目标而相互协作的组件构成的。这些组件必须彼此了解,并且相互协作来完成工作；
* 创建应用对象之间关联关系的传统方法(通过构造器或者查找)通常会导致结构复杂的代码,对象彼此之间高度耦合,从而难以复用和测试；
* 而在Spring中,由容器负责把需要相互协作的对象引用赋予各个对象,创建应用对象之间协作关系，这种行为通常称为装配(wiring),也是依赖注入( DI )的本质.
### Spring三种主要的装配机制:
* 在XML中进行显式配置。
* 在Java中进行显式配置。
* 隐式的bean发现机制和自动装配
> 注:不同配置风格是可以互相搭配，但应该尽可能地使用自动配置的机制，尽量使用JavaConfig(类型安全,且比XML更加强大)

## 注解自动化装配bean：
* Spring从两个角度来实现自动化装配:
<br/>　　　－组件扫描(component scanning): Spring会自动发现应用上下文中所创建的bean。
<br/>　　　－自动装配(autowiring):Spring自动满足bean之间的依赖
> (如CD播放器依赖于CD才能完成它的使命，可以定义CD的一个接口定义了CD播放器对一盘CD所能进行的操作，同时不同种类的CD有CD接口的多个实现，则CD和CD播放机均为组件Bean,同时CD播放机的各种行为又依赖于不同种类的CD)
#### 创建可发现的Bean：
* @Component注解：表明该类为组件类,并告知Spring要为这个类创建对应的bean（一般标识接口的实现类）
* @ComponentScan注解:标识在配置类上，表明在Spring中启用组件扫描,@ComponentScan默认会扫描与配置类相同的包(@Configuration注解标识Spring配置类)
> 注：或通过xml文件中的<context:component-scan base-package="package_name">，启动组件扫描，并指示扫描的包名；
<br/>（组件扫描默认是不启用的,需要显式配置Spring,从而命令它去寻找带有 @Component 注解的类,并为其创建 bean）

> 注：在JUnit测试中＠Runwith(SpringJUnit4ClassRunner.class)配置测试执行器以便在测试开始的时候自动创建Spring的应用上下文；注解@ContextConfiguration(clsses=ConfigurationClass.class)标识配置类，指明从哪里加载配置

#### 为组件扫描的bean命名
* Spring应用上下文中所有的bean都会给定一个ID(默认会将类名的第一个字母变为小写作为id),可以通过@Component（"bean_id"）为组件指定id,或使用Java依赖注入规范(Java Dependency Injection)中所提供的 @Named（"bean_id"） 注解指定；

#### 设置组件扫描的基础包
* @ComponentScan注解默认会以配置类所在的包作为基础包扫描组件,同时可以通过ComponentScan("base_package_name")(一个)或ComponentScan(basePackages=｛"base_package_name"，．．．｝)（多个）配置扫描基础包
* 除了将包设置为简单的 String 类型(类型不安全(not type-safe))之外，还可以将其指定为包中所包含的类或接口ComponentScan(basePackageClasses=｛"ClassName.class"，．．．｝)
> (可以在包中创建一个用来进行扫描的空标记接口(marker interface)。通过标记接口的方式,保持对重构友好的接口引用,避免引用任何实际的应用程序代码)
#### 注解实现自动装配
* 自动装配：将组件扫描得到的bean和它们的依赖装配在一起，让Spring自动满足bean依赖;
* @Autowired注解：构造器上添加了@Autowired注解,这表明当Spring创建此Bean的时候,会通过这个构造器来进行实例化并且会传入一个可设置给传入参数类型的bean（@Autowired 注解不仅能够用在构造器上,还能用在属性的 Setter 方法上，此外@Autowired 注解可以用在类的任何方法上，根据方法的参数注入相应的Bean,尝试满足方法参数上所声明的依赖）
```
@Component    //标识组件bean
public ClassName implements InterfaceName{
    private Dependance_Bean dependance_bean;   //依赖属性
    @Autowired　　　　　　　　　　　　　　//自动装配
    public ClassName(Dependance_Bean db){
        this.dependance_bean=db;
    }
    public void action(){
        dependance_bean.action();
    }
}
```
>　注意:
<br/>　　　－当有且只有一个bean匹配依赖需求,这个bean将会被装配进来；
<br/>　　　－当没有匹配的bean,在应用上下文创建的时候,Spring会抛出一个异常(或＠Autowired(required=false)Spring会尝试执行自动装配,但是如果没有匹配的bean的话,Spring将会让这个bean处于未装配的状态,避免异常的出现,但是如果在你的代码中没有进行null检查,这个处于未装配状态的属性可能会出现NullPointerException异常)；
<br/>　　　－当有多个bean能满足依赖关系, Spring将会抛出一个异常,表明没有明确指定要选择哪个bean进行自动装配；
> 此外还可以使用@Inject注解实现自动装配（@Inject 注解同@Named 注解一样来源于 Java 依赖注入规范）

## JavaConfig显式配置方案:
* 有时需要明确配置Spring，如需要将第三方库中的组件装配到你的应用中，则没有办法在它的类上添加@Component和@Autowired注解，因此就不能使用自动化装配的方案
* 有两种可选方案: Java 和 XML（显式配置时JavaConfig是更好的方案,因为它更为强大、类型安全并且对重构友好)
* @Configuration注解表明：这个类是一个配置类;
简单bean声明：在JavaConfig中声明bean,需要编写一个添加@Bean注解的方法创建返回所需类型的实例(默认情况下,bean的ID与带有@Bean注解的方法名是一样的,可以通过 name 属性指定一个不同的名字＠Bean(name=""))
* 借助JavaConfig实现注入:最简单方式就是引用创建bean的方法，调用需要传入对象的构造器来创建返回实例;
```
import org.spring.framework.context.annotation.Configuration;
@Configuration   //标识配置类
public class ConfigurateClassName{
    @Bean(name="..")    //声明简单Bean     
    public BeanClassName1 bean_name1(){
        ．．．．．．．
        return new BeanClassName1();
    }
    ＠Bean(name="..")     //将bean1装配到bean2实现注入方法1
    public BeanClassName2 bean_name2(){
        ...........
        return new BeanClassName2(bean_name1());
    }
    或
    ＠Bean(name="..")     //将bean1装配到bean2实现注入方法2(推荐，不依赖@Bean方法，可将bean的声明放在多个配置类中)
    public BeanClassName2 bean_name2(BeanClassName bcn){
        ...........
        return new BeanClassName2(bcn);
    }
}
```
> 注：标有@Bean注解的方法,Spring将会拦截所有对它的调用,并确保直接返回该方法所创建的bean,而不是每次都对其进行实际的调用，默认情况下,Spring中的bean都是单例的

## 通过XML装配 bean
* 声明bean：<bean id=".." class="...">class的指定应使用权限定的类名
> (bean 的类型以字符串的形式设置,不能从编译期的类型检查中发现错误，可借助一些工具检查xml文件的合法性，如Spring Tool Suite)
### 依赖注入:
#### 构造器注入:
* <constructor-arg ref="构造器传入依赖参数bean_id">（对象注入）／ <constructor-arg value="构造器传入依赖参数"> (字面量注入)
* 使用 Spring 3.0 所引入的 c- 命名空间：<bean...c:paramater_name(构造器参数名)-ref="构造器传入依赖参数bean_id" >或c:paramater_index(构造器参数索引)-ref(对象注入),c:_paramater=".."/或c:_index＝".."字面量注入
* 其中集合装配注入，只能由构造器注入实现，C-命名空间无法实现（<list>对象Java中List集合,<set>对应Set集合）:
```
／<constructor-arg>
<list>
..
<value>".."</value>(或<ref bean=" "/>)
..
</list>(或<set>)
</constructor-arg>
```
设置属性
通常对强依赖使用构造器注入,而对可选性的依赖使用属性注入，与构造器注入类似属性注入也有两种<property> 元素为属性的 Setter 方法注入参数和p-命名空间,使用与构造注入类似；

## 导入和混合配置
#### 在JavaConfig中引用其他配置:
* 使用 @Import(ConfigurationClass.class...) 注解导入其他JavaConfig类配置
* 使用@ImportResource("xml_file_path") 注解引用xml文件配置
#### 在XML配置中引用其他配置:
* 通过<bean class="CongurationClassName"> 元素,导入JavaConfig类的配置
* 通过<import resource="xml_file_path">引用其他xml配置











