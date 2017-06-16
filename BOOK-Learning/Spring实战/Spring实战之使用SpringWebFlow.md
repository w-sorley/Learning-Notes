---
title: 使用Spring Web Flow
date: 2017-06-05
tags: [Spring MVC,framework,]
categories: FrameWork
---

# 使用Spring Web Flow
* Spring Web Flow是Spring MVC的扩展,它支持开发基于流程的应用程序。它将流程的定义与实现流程行为的类和视图分离.
## 在Spring中配置 Web Flow
* Spring Web Flow是构建于Spring MVC基础之上的,这意味着所有的流程请求都需要首先经过Spring MVC的DispatcherServlet,需要在Spring应用上下文中配置一些bean来处理流程请求并执行流程,这些只能在XML中对其进行配置：
### 1.声明命名空间
* 在上下文定义XML文件中添加Spring Web Flow命名空间声明
```
<beans xmlns:flow="http://www.springframework.org/schema/webflow-config"
xsi:schemaLocation="http:www.springframework.org/schema/webflow-config
                    http:www.springframework.org/schema/webflow-config/[CA]
                    spring-webflow-config-2.3.xsd...."
```
### 2.装配流程执行器
* 流程执行器(flow executor)驱动流程的执行:当用户进入一个流程时,流程执行器会为用户创建并启动一个流程执行实例;当流程暂停的时候(如为用户展示视图时),流程执行器会在用户执行操作后恢复流程
```
<flow:flow-executor id="flowExecutor" />
```
### 3.配置流程注册表
* 尽管流程执行器负责创建和执行流程,但它并不负责加载流程定义.流程注册表 ( flow registry )的工作是加载流程定义并让流程执行器能够使用它们
```
<flow:flow-registry id="flowRegistry" base-path="/flow_file_path">　　//指定流程注册表目录
    <flow：flow-location-pattern value="*-flow.xml">   //指定任何文件名以 “-flow.xml” 结尾的 XML 文件都将视为流程定义
</flow:flow-registry>
//所有的流程都是通过其 ID 来进行引用的：流程id=流程注册表基本路径/流程id/流程定义
//也可以去除base-path属性,而显式声明流程定义文件的位置
<flow:flow-registry id="flowRegistry">
    <flow:flow-location [id=".."] pahth="/path/file.xml">
<flow:flow-registry>
```
### 4.处理流程请求
* 对于流程而言,需要一个FlowHandlerMapping来帮助DispatcherServlet将流程请求发送给Spring Web Flow,在Spring应用上下文中,FlowHandlerMapping的配置如下:
```
<bean class="org.springframework.webflow.mvc.servlet.FlowHandlerMapping">
    <property name="flowRegistry" ref="flowRegistry"> //装配了流程注册表的引用,从而知道如何将请求的 URL 匹配到流程上
</bean>
```
FlowHandlerMapping仅仅是将流程请求定向到Spring Web Flow，FlowHandlerAdapter会响应发送的流程请求并对其进行处理(等同于Spring MVC的控制器),会处理流程请求并管理基于这些请求的流程
```
<bean class="org.springframework.webflow.mvc.servlet.FlowHandlerAdapter">
    <property name="flowExcutor" ref="flowExcutor" />
</bean>
```
## 流程的组件
* 在Spring Web Flow中,流程是由三个主要元素定义的:状态、转移和流程数据
### 状态
* 状态(State)是流程中事件发生的地点,流程中的状态是业务逻辑执行、做出决策或将页面展现给用户的地方.Spring Web Flow定义了五种不同类型的状态：
<br/>(通过选择 Spring Web Flow 的状态几乎可以把任意的安排功能构造成会话式的Web 应用,不同类型的状态组合起来形成一个完整的流程)
<br/>　－行为(Action)：行为状态是流程逻辑发生的地方
<br/>　－决策(Decision)：决策状态将流程分成两个方向,它会基于流程数据的评估结果确定流程方向
<br/>　－结束(End)：结束状态是流程的最后一站。一旦进入End状态,流程就会终止
<br/>　－子流程(Subflow)：子流程状态会在当前正在运行的流程上下文中启动一个新的流程
<br/>　－视图(View)：视图状态会暂停流程并邀请用户参与流程
#### 视图状态
* 视图状态用于为用户展现信息并使用户在流程中发挥作用,实际的视图实现可以是 Spring 支持的任意视图类型,但通常是用 JSP 来实现.在流程定义的 XML 文件中,<view-state>用于定义视图状态:
```
<view-state id="view_id" /> //id 属性在流程内标示这个状态,也指定了流程到达这个状态时要展现的逻辑视图名(也可以用view属性指定，也可以；；利用model属性为表单绑定对象)
```
### 行为状态
行为状态一般会触发 Spring 所管理 bean 的一些方法并根据方法调用的执行结果转移到另一个状态,是应用程序自身在执行任务,使用 <action-state> 元素来声明:
```
<action-state id="action_id">
    <evaluate expresssion="action_method_name" />  //给出了行为状态要做的事情,expression属性指定了进入这个状态时要评估的表达式
    <transition to="dest" />
</action-state>
```

> 注：Spring Web Flow与表达式语言:在1.0 版本的时候,Spring Web Flow使用的是对象图导航语言(Object-Graph Navigation Language,OGNL),2.0版本换成统一表达式语言(Unified Expression Language,Unified EL)。2.1版本中,Spring Web Flow使用SpEL.
决策状态
流程在某一个点根据流程的当前情况进入不同的分支(也有可能流程会完全按照线性执行,从一个状态进入另一个状态,没有其他的替代路线)。决策状态能够在流程执行时产生两个分支，将评估一个 Boolean 类型的表达式,然后在两个状态转移中选择一个，通过 <decision-state> 元素进行定义：
```
<decision-state id="decision_id">
<if test="boolean_decision_menthod()" then="one_branch" else=”other_branch“>
</decision-state>
```
### 子流程状态
同将应用程序的所有逻辑分散到多个类,方法以及其他结构中相似，也可以将流程分成独立的部分，<subflow-state> 允许在一个正在执行的流程中调用另一个流程：
```
<subflow-state id="sub_id" subflow="">
    <input name="input_name" value="input_Object"/>  //用于传递子流程的输入
    <transition on="now_static" to="dest_static" />
</subflow-state>
```
结束状态
当流程转移到结束状态时所处的行为，<end-state>元素指定了流程的结束：
```
<end-state id="">
```
* 当到达<end-state>状态,流程会结束。接下来会发生什么取决于几个因素:
<br/>   －如果结束的流程是一个子流程,那调用它的流程将会从<subflow-state>处继续执行<end-state>的ID将会用作事件触发从<subflow-state>开始的转移。
<br/>   －如果<end-state>设置了view属性,指定的视图将会被渲染。视图可以是相对于流程路径的视图模板,如果添加“externalRedirect:”前缀的话,将会重定向到流程外部的页面,如果添加“flowRedirect:”将重定向到另一个流程中。
<br/>   －如果结束的流程不是子流程,也没有指定 view 属性,那这个流程只是会结束而已。浏览器最后将会加载流程的基本 URL 地址,当前已没有活动的流程,所以会开始一个新的流程实例
> 注:流程可能会有不止一个结束状态。子流程的结束状态ID确定了激活的事件,通过多种结束状态来结束子流程,从而能够在调用流程中触发不同的事件(即使不是在子流程中,也有可能在结束流程后,根据流程的执行情况有多个显示页面供选择)





## 转移
* 转移连接了流程中的状态。流程中除结束状态之外的每个状态,至少都需要一个或多个转移,确定一旦这个状态完成时流程要去向哪里，转移使用 <transition> 元素来进行定义，作为各种状态元素的子元素：
```
<transition to="dest_static" />  //to属性指定当前状态的默认转移选项

//基于事件的触发来进行转移定义(视图状态,事件通常是用户采取的动作;行为状态,事件是评估表达式得到的结果;子流程状态,事件取决于子流程结束状态的ID),用on属性来指定触发转移的事件:
<transition on="triger_event" to="dest_static">  //事件triger_event发生时进入dest_static状态
<transition on-exception="package.Exception" to="dest_static">  //当指定异常发生时，进入dest_static状态
```
#### 全局转移
* 避免在多个状态中都重复通用的转移，可以将<transition>元素作为<global-transitions>的子元素,从而把它们定义为全局转移：
```
<global-transition>
<transition on="universal_event" to="universal_static"> //流程中的所有状态都会默认拥有这个转移
</global-transition>
```

### 流程数据
* 当流程从一个状态进行到另一个状态时,它会带走一些数据。有时候,这些数据只需要很短的时间(可能只要展现页面给用户)。有时候,这些数据会在整个流程中传递并在流程结束的时候使用。
#### 声明变量
* 流程数据保存在变量中,而变量可以在流程的各个地方进行引用,可以以多种方式创建：
```
//如使用<var>元素创建流程数据,作用在定义变量的流程内
<var name="variable_name" "stored_Class_instance">
//使用<evaluate>元素来创建变量,作用域通过result属性的前缀指定
<evaluate result="variable_name" expresssion=“T(SpEL_expression)”>  //计算了一个表达式(SpEL表达式)并将结果放到了名为variable_name的变量
//可以利用<set>元素设置变量的值，作用域通过name属性的前缀指定
<set name="variable_name" value="setted_value_expression" />    //将变量设置为表达式计算的结果
```
#### 定义流程数据的作用域
* 流程中携带的数据会拥有不同的生命作用域和可见性,这取决于保存数据的变量本身的作用域.Spring Web Flow定义了五种不同作用域:
<br/>　－Conversation：最高层级的流程开始时创建,在最高层级的流程结束时销毁。被最高层级的流程和其所有的子流程所共享
<br/>　－Flow：当流程开始时创建,在流程结束时销毁。只有在创建它的流程中是可见的
<br/>　－Request：当一个请求进入流程时创建,在流程返回时销毁
<br/>　－Flash：当流程开始时创建,在流程结束时销毁。在视图状态渲染后,它也会被清除
<br/>　－View：当进入视图状态时创建,当这个状态退出时销毁。只在视图状态内是可见的

## 保护Web流程
* Spring Web Flow 中的状态、转移甚至整个流程都可以借助 <secured> 元素实现安全性,该元素会作为这些元素的子元素，如：
```
//保护对一个视图状态的访问
<view-state id="view_id">
    <secured attributes="PERMISSION" match="any/all">  //attributes 属性授予访问权限(使用逗号分隔的权限列表)
</view-state>
```