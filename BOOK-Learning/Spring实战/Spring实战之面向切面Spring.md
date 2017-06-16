---
title: Spring实战之面向切面Spring
date: 2017-05-31
tags: [JAVA,framewoke,AOP]
cetegories: FrameWork
---

# Spring实战之面向切面Spring
<br/>
<br/>

## 面向切面编程：
* 在软件开发中,散布于应用中多处的功能被称为横切关注点(cross-cutting concern)通常,横切关注点从概念上是与应用的业务逻辑相分离(但是往往会直接嵌入到应用的业务逻辑之中)。以横切关注点与业务逻辑相分离为目标的编程模式称为面向切面编程(AOP)
* DI有助于应用对象之间的解耦;而AOP 可以实现横切关注点与它们所影响的对象之间的解耦,横切关注点可以被描述为影响应用多处的功能，切面能帮助我们模块化横切关注点,适用场景包括日志，声明式事务、安全和缓存等。

* 对于通用功能的重用,最常见的面向对象技术是继承(inheritance)或委托(delegation)，但在整个应用中使用相同的基类,继承往往会导致一个脆弱的对象体系;而使用委托可能需要对委托对象进行复杂的调用。
* 此外利用切面技术可以一个地方定义通用功能,通过声明的方式定义这个功能要以何种方式在何处应用(无需修改受影响的类)。从而横切关注点可以被模块化为被称为切面(aspect)特殊的类；
> 优点：每个关注点都集中于一个地方;服务模块更简洁;

## AOP术语
* 描述切面的常用术语有通知(advice)、切点(pointcut)和连接点(join point):
> （在一个或多个连接点上,可以把切面的功能(通知)织入到程序的执行过程中）

### 通知(Advice)
* 在AOP术语中,切面必须要完成的工作被称为通知，定义了切面是什么以及何时使用;
* Spring切面可以应用5种类型的通知：
<br/>　－前置通知(Before):在目标方法被调用之前调用通知功能;
<br/>　－后置通知(After):在目标方法完成之后调用通知,此时不会关心方法的输出是什么;
<br/>　－返回通知(After-returning):在目标方法成功执行之后调用通知;
<br/>　－异常通知(After-throwing):在目标方法抛出异常后调用通知;
<br/>　－环绕通知(Around):通知包裹了被通知的方法,在被通知的方法调用之前和调用之后执行自定义的行为。

### 连接点(Join point）
* 应用可能有数以千计的时机应用通知，这些时机被称为连接点.
* 是在应用执行过程中能够插入切面的一个点．这个点可以是调用方法时、抛出异常时、甚至修改一个字段时。切面代码可以利用这些点插入到应用的正常流程之中,并添加新的行为。

### 切点(Poincut)
* 切点有助于缩小切面所通知的连接点的范围(一个切面并不需要通知应用的所有连接点)，切点的定义会匹配通知所要织入的一个或多个连接点.
> (通常使用明确的类和方法名称,或是利用正则表达式定义所匹配的类和方法名称来指定这些切点)

### 切面(Aspect)
* 切面是通知和切点的结合。通知和切点共同定义了切面的全部内容-它是什么,在何时和何处完成其功能。

### 引入(Introduction)
* 通过引入向现有的类添加新方法或属性
### 织入(Weaving)
* 织入是把切面应用到目标对象并创建新的代理对象的过程.切面在指定的连接点被织入到目标对象中。
* 在目标对象的生命周期里有多个点可以进行织入:
<br/>　　－编译期:需要特殊的编译器,AspectJ的织入编译器就是以这种方式织入切面的。
<br/>　　－类加载期:在目标类加载到JVM时被织入,需要特殊的类加载器(ClassLoader),可以在目标类被引入应用之前增强该目标类的字节码。 AspectJ 5的加载时织入(load-time weaving, LTW)支持以这种方式织入切面。
<br/>　　－运行期:在应用运行的某个时刻被织入。一般情况下在织入切面时, AOP容器会为目标对象动态地创建一个代理对象。 SpringAOP就是以这种方式织入切面。

## Spring对AOP的支持
* Spring提供了4种类型的AOP支持:
<br>　－基于代理的经典Spring AOP;
<br>　－纯POJO切面;
<br>　－@AspectJ注解驱动的切面;
<br>　－注入式AspectJ切面(适用于Spring各版本)
> （前三种都是Spring AOP实现的变体, Spring AOP构建在动态代理基础之上,因此,Spring对AOP的支持局限于方法拦截）
<br/>可以在XML配置中,借助Spring的aop命名空间,可以声明式地将纯POJO转换为切面．
<br/>同时Spring借鉴了AspectJ的切面,提供注解驱动的AOP
<br/>（本质上,它依然是Spring基于代理的AOP ,但是编程模型几乎与编写成熟的AspectJ注解切面完全一致，避免了使用xml）
* 如果AOP需求超过了简单的方法调用(如构造器或属性拦截),那么需要考虑使用注入式AspectJ切面将值注入到 AspectJ驱动的切面中

### Spring AOP框架
* Spring通知由Java编写,定义通知所应用的切点通常会使用注解或在Spring配置文件里采用XML来编写
* Spring在运行时通知对象,通过在代理类中包裹切面,在运行期把切面织入到 Spring管理的bean中,从而代理类封装了目标类,并拦截被通知方法的调用,再把调用转发给真正的目标bean 。
* 当代理拦截到方法调用时,在调用目标bean方法之前,会执行切面逻辑。同时直到应用需要被代理的bean时, Spring才创建代理对象.
* 如果使用的是ApplicationContext,在ApplicationContext从BeanFactory中加载所有bean时, Spring才会创建被代理的对象（因为Spring运行时才创建代理对象,所以不需要特殊的编译器来织入Spring AOP的切面）。
> 注：Spring基于动态代理,所以Spring只支持方法连接点,而AspectJ 和JBoss,除了方法切点,它们还提供了字段和构造器接入点，从而可以创建细粒度的通知（例如拦截对象字段的修改），可以在bean创建时应用通知。

#### 通过切点来选择连接点
* 在Spring AOP中,使用AspectJ的切点表达式语言来定义切点，Spring仅支持AspectJ切点指示器(pointcut designator)的一个子集(使用AspectJ其他指示器时,将会抛出 IllegalArgument-Exception 异常)：
<br/>　－arg():限制连接点匹配参数为指定类型的执行方法
<br/>　－@args():限制连接点匹配参数由指定注解标注的执行方法
<br/>　－execution()：用于匹配是连接点的执行方法
<br/>　－this()：限制连接点匹配 AOP 代理的 bean 引用为指定类型的类
<br/>　－target：限制连接点匹配目标对象为指定类型的类
<br/>　－@target()：限制连接点匹配指定的执行对象，这些对象对应的类要具有指定类型的注解
<br/>　－within():限制连接点匹配指定的类型
<br/>　－@within():限制连接点匹配指定注解所标注的类型（当使用Spring AOP时，方法定义在由指定的注解所标注的类里）
<br/>　－@annotation:限定匹配带有指定注解的连接点
> (此外Spring还引入了一个新的 bean() 指示器，允许在切点表达式中使用bean的ID来标识bean，使用bean ID或bean名称作为参数来限制切点只匹配特定的bean使用注解创建切面)


### 编写切点
* 形式：pointcut_designator(return_type package.Class.method(param..))&&...... :'*'表示
> 注：可以使用逻辑操作符与( and，&&),或（or，||）,非(not,!)连接多个指示器（在Spring的XML配置中，“&”有特殊含义，描述切点时,使用and来代替 “&&” 。）

* 使用@AspectJ注解标注POJO为切面
* 使用AspectJ注解声明通知方法:
<br/>　－@After：通知方法会在目标方法返回或抛出异常后调用
<br/>　－@AfterReturning：通知方法会在目标方法返回后调用
<br/>　－@AfterThrowing通知方法会在目标方法抛出异常后调用
<br/>　－@Around通知方法会将目标方法封装起来
<br/>　－@Before通知方法会在目标方法调用之前执行
<br/>（给定了一个切点表达式作为它的值，方法上添加@Pointcut注解能够在一个@AspectJ切面内定义可重用的切点）
> 注：最后需要在配置类的类级别上通过使用EnableAspectJ-AutoProxy注解（或在xml中使用Spring aop命名空间中的 \<aop:aspectj-autoproxy\> 元素）启用自动代理功能，以上切面配置才会生效；
<br/>另外Spring的AspectJ自动代理仅仅使用@AspectJ作为创建切面的指导，会为使用@Aspect注解的bean创建一个代理,这个代理会围绕着所有该切面的切点所匹配的bean，因此切面依然是基于代理的

#### 创建环绕通知
* 环绕通知是最为强大的通知类型，能够让编写的逻辑将被通知的目标方法完全包装起来，就像在一个通知方法中同时编写前置通知和后置通知
> (通知方法接受ProceedingJoinPoint 作为参数，在通知中通过它来调用被通知的方法，如果不调proceed()方法的话,那么通知实际上会阻塞对被通知方法的调用处理)
* 通知中的参数args(..) 限定符表明传递给通知的方法的参数也会传递到通知中去,切点定义中的参数与切点方法中的参数名称是一样的从而完成了从命名切点到通知方法的参数转移

### 通过注解引入新功能
* 一些编程语言（如Ruby和Groovy）有开放类的理念，可以不用直接修改对象或类的定义就能够为对象或类增加新的方法．java本身并不支持，但可以借助AOP实现类似的功能
> (当引入接口的方法被调用时,代理会把此调用委托给实现了新接口的某个其他对象，从而一个 bean 的实现被拆分到了多个类中。)

* 创建一个切面，通过@DeclareParents注解,将需要添加的接口引入到目标bean中.@DeclareParents 注解由三部分组成
<br/>　－value属性:指定了哪种类型的 bean 要引入该接口。
<br/>　－defaultImpl属性:指定了为引入功能提供实现的类。
<br/>　－@DeclareParents注解:标注的静态属性指明了要引入了接口。

* 在Spring应用中将定义的切面接口声明为一个bean

* Spring的自动代理机制将会获取到它的声明,当Spring发现一个bean使用了@Aspect注解时, Spring就会创建一个代理,然后将调用委托给被代理的bean或被引入的实现(这取决于调用的方法属于被代理的 bean还是属于被引入的接口)。

### 在XML中声明切面
* 在Spring的aop命名空间中,提供了多个元素用来在 XML 中声明切面：
<br/>　　－\<aop:advisor\>: 定义AOP通知器
<br/>　　－\<aop:after\>:定义AOP后置通知(不管被通知的方法是否执行成功)
<br/>　　－\<aop:after-returning\>:定义AOP返回通知
<br/>　　－\<aop:after-throwing\>:定义AOP异常通知
<br/>　　－\<aop:around\>:定义AOP环绕通知
<br/>　　－\<aop:aspect\>:定义一个切面
<br/>　　－\<aop:aspectj-autoproxy\>:启用@AspectJ注解驱动的切面
<br/>　　－\<aop:before\>:定义一个AOP前置通知
<br/>　　－\<aop:config\>:顶层的AOP配置元素。大多数的\<aop:*\>元素必须包含在\<aop:config\>元素内
<br/>　　－\<aop:declare-parents\>:以透明的方式为被通知的对象引入额外的接口
<br/>　　－\<aop:pointcut\>:定义一个切点
* 举例：
```
<aop:config>　　　　　　　　　　　　　　//大多数的 AOP 配置元素必须在 <aop:config> 元素的上下文内使用
<aop:aspect ref="bean_id">       //明了一个切面,ref 元素所引用的 bean 提供了在切面中通知所调用的方法
<aop:pointcut                   //定义切点，可一处定义多处引用，此处作用范围为一个本切面
    id=".."
    expression="" />

<aop:before               //前置通知，后置通知after-returning,异常通知after-throwing，环绕通知around类似
    pointcut=".."
    method=".."/>
<aop:aspect>
<aop:delare-parents                 //声明了此切面所通知的 bean 要在它的对象层次结构中拥有新的父类型
    type-matching="package.class+"
    implement-interface="package.interface"
    dafault-impl="package.interface_impl_class">   //直接标识委托或者使用delegate-ref＝＂bean_id＂属性标识符
</aop:aspect>
</aop:config>
```
## 注入AspectJ切面
* 在使用AspectJ切面时，如果在执行通知时,切面依赖于一个或多个类,可以借助Spring的依赖注入把bean装配进AspectJ切面中
> 所有的 AspectJ 切面都提供了一个静态的aspectOf() 方法,该方法返回切面的一个单例,可以使用 factory-method 来调用asepctOf()方法获得切面的实例，然后像 <bean> 元素规定的那样在该对象上执行依赖注入