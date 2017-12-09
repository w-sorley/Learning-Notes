---
title: Spring实战之构建SpringWeb应用
date: 2017-06-01
tags: [Web,Spring MVC,Spring]
categories: FrameWork
---

# Spring实战之构建SpringWeb应用
<br/>
<br/>
## 概述
Spring MVC是基于模型-视图-控制器(Model-View-Controller, MVC)模式实现,利用它从而能够构建构建灵活和松耦合的Web应用程序；
spring对请求的处理过程大概是：Spring将请求在调度Servlet、处理器映射(handler mapping)、控制器以及视图解析器(view resolver)之间移动；请求从客户端发起,经过Spring MVC中的组件,最终再返回到客户端的

## 请求处理

* 当用户在Web浏览器中点击链接或提交表单的时候,请求开始工作（会将信息从一个地方带到另一个地方）
* 首先经过Spring的DispatcherServlet前端控制器Servlet(front controller,一个单实例的Servlet)，DispatcherServlet会查询一个或多个处理器映射(handler mapping,处理器映射会根据请求所携带的 URL 信息来进行决策) 来确定将请求发送给哪个Spring MVC控制器(controller);
* 请求到了控制器,会卸下其负载(用户提交的信息)并耐心等待控制器处理．控制器在完成逻辑处理后,通常会产生一些需要返回给用户并在浏览器上显示的信息，这些信息被称为模型(model)．信息需要发送给一个格式化视图(view),通常会是JSP
* 最后控制器将模型数据打包,并且标示出用于渲染输出的视图名,然后将请求连同模型和视图名发送回DispatcherServlet
* DispatcherServlet将会使用视图解析器(view resolver) 来将逻辑视图名匹配为一个特定的视图实现(可能是也可能不是JSP)
* 最后是视图的实现(可能是JSP) ,在这里它交付模型数据。视图将使用模型数据渲染输出,这个输出会通过响应对象传递给客户端

## 搭建Spring MVC
### 配置DispatcherServlet
使用Java将DispatcherServlet配置在Servlet容器中(而不使用web.xml文件)
> 注：扩展AbstractAnnotation-ConfigDispatcherServletInitializer的任意类都会自动地配置Dispatcher-Servlet和Spring应用上下文;Spring的应用上下文会位于应用程序的Servlet上下文之中
#### AbstractAnnotationConfigDispatcherServletInitializer剖析
* 在Servlet 3.0环境中,容器会在类路径中查找实现javax.servlet.ServletContainerInitializer接口的类,如果能发现的话,就会用它来配置Servlet容器
* 而Spring中名为SpringServletContainerInitializer恰好提供了这个接口的实现,,这个类反过来又会查找实现WebApplicationInitializer的类并将配置的任务交给它们来完成
* 而AbstractAnnotationConfigDispatcherServletInitializer类恰好为Spring 3.2引入的一个WebApplicationInitializer类的基础实现
* 从而当自定义的配置类扩展了AbstractAnnotationConfigDispatcherServletInitializer(同时也就实现了WebApplicationInitializer),当部署到Servlet 3.0容器中的时候,容器会自动发现它,并用它来配置Servlet上下文。
```
/*DispatcherServlet前端控制器实现，配置servlet应用上下文，Spring应用上下文*/
import org.springframework.web.servlet.support.AbstractAnnotationConfigDispatcherServletInitializer

public class WebAppInitializer_name 
        extends AbstractAnnotationConfigDispatcherServletInitializer{
    @Override
    protected String[] getServletMapping(){
        return new String[] { "/" }      //将DispatcherServlet映射到"/"，标识默认servlet，会处理所有对应用的请求
    }
    @Override
    protected Class<?> getRootConfigClasses(){       //返回的带有@Configuration注解的类将会用来配置ContextLoaderListener创建的应用上下文中的bean。
        return new CLass<?> { RootConfigClassName.class; };
    }
    @Override
    protected Class<?> getServletConfigClasses(){　　　　　　　//返回的带有@Configuration注解的类，用来定义DispatcherServlet创建的spring应用上下文中的bean。
        return new Class<?> { ServletConfigClassName.class; };
    }
}

```
#### 应用上下文创建
* AbstractAnnotationConfigDispatcherServletInitializer会同时创建DispatcherServlet和Servlet监听器（ContextLoaderListener）。 
* 当DispatcherServlet启动时,它会创建Spring应用上下文,并加载配置文件或配置类（由getServletConfigClasses()实现方法指定）中所声明的bean，主要加载包含Web组件的bean,如控制器、视图解析器以及处理器映射等
* 而Servlet应用上下文由ContextLoaderListener创建，并加载配置文件或配置类（由getRootConfigClasses()实现方法指定）中所声明的bean，主要加载驱动应用后端的中间层和数据层组件等
> 注：通过AbstractAnnotationConfigDispatcherServletInitializer来配置DispatcherServlet,只能部署到支持Servlet 3.0的服务器中才能正常工作(如Tomcat 7或更高版本）,是传统web.xml方式的替代方案，同时在配置中亦可以时包含 web.xml和AbstractAnnotationConfigDispatcher-ServletInitializer

#### 启用Spring MVC
* 若Spring是使用XML进行配置的,你可以使用<mvc:annotation-driven>启用注解驱动的Spring MVC。
* 也可基于Java进行配置:利用一个带有@EnableWebMvc注解的配置类(最简单的Spring MVC配置)启用Spring MVC


```
/*Spring MVC 配置,创建Spring应用上下文，加载相关的bean*/
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.ViewResolver;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;
import org.springframework.web.servlet.view.InternalResourceViewResolver;

@Configuration
@EnableWebMvc     //启用Spring MVC
@ComponentScan("package_name")      //启用组件扫描，将会扫描package_name指定的包来查找组件
public class WebConfig extends WebMvcConfigurerAdapter {
	@Bean
	public ViewResolver viewResolver() {　　　　//配置视图解析器
		InternalResourceViewResolver resolver = new InternalResourceViewResolver();
		resolver.setPrefix("/WEB-INF/views");　　　//设置视图查找前缀
		resolver.setSuffix(".jsp");　　　　　　　　//设置视图查找前缀
		resolver.setExposeContextBeansAsAttributes(true);　　//是否应用上下文中所有的Spring bean都可以作为请求参数（？？）
		return resolver;
	}
    ＠Override                                            //配置静态资源的处理
    public void configureDefaultServletHandling(DefaultServletHandlerConfigurer config )
    {　　　　　　　　　　　　　　　　　　　//要求 DispatcherServlet 将对静态资源的请求转发到 Servlet 容器中默认的 Servlet 上    
        config.enable();            //而不是使用 DispatcherServlet 本身来处理此类请求
    }
}
```
> 注意：
<br/>－若使用最简单的Spring MVC配置, DispatcherServlet会映射为应用的默认Servlet,所以它会处理所有的请求,包括对静态资源的请求,如图片和样式表(在大多数情况下,这可能并不是想要的效果)。
<br/>－若没有配置视图解析器，Spring默认会使用BeanNameView-Resolver,这个视图解析器会查找ID与视图名称匹配且实现View接口的bean,以这样的方式来解析视图。
<br/>－若没有启用组件扫描，Spring只能找到显式声明在配置类中的控制器。

* Servlet相关配置：
```
/*创建Servlet上下文，加载相关的bean*/
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.ComponentScan.Filter;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.FilterType;
import org.springframework.web.servlet.config.annotation.EnableWebMvc;

@Configuration   //标注配置类
@ComponentScan(basePackages = { "spitter" }, excludeFilters = {     //指定要扫描的包加载其中的bean,(??)
		@Filter(type = FilterType.ANNOTATION, value = EnableWebMvc.class) })         //
public class RootConfig {
}
```



#### 创建控制器：
* @Controller注解：标注控制器,因为此注解基于@Component注解，是一个构造型( stereotype )的注解，会使其标注的类成为组件扫描时的候选bean，因此不需要在配置类中显式声明任何的控制器．
* @RequestMapping注解：其value属性指定了这个方法所要处理的请求路径, method属性指定所处理的HTTP方法；
* 视图解析器会根据控制器的返回结果生成视图名称，查找JSP文件，将逻辑名称解析为实际的视图（可能会根据视图解析器的配置在视图名称上加一个特定的前缀和后缀）

```
import static org.springframework.web.bind.annotation.RequestMethod.GET;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
@Controller    //标注控制器
@RequestMapping(value = "/")　　　//指定类级别的所要处理的请求路径，会应用到控制器的所有处理器方法上，value属性能够接受一个 String 类型的数组映射到多个路径
public class ControllerName {
	@RequestMapping(method = GET)　　　　//HTTP方法依然映射在方法级别上,会对类级别上的 @RequestMapping 的声明进行补充
	public return_type controllerMethodName() {
		return "ViewName";    //视图解析器会根据控制器的返回结果和自身配置生成视图逻辑名称
	}
}
```
### 传递模型数据到视图中
* 首先,需要定义一个数据访问的接口,并在稍后实现它
```
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;

@Controller
@RequestMapping("/url_path")
public class ControllerName {
	private interface interfaceName;  //私有数据访问接口属性
	@Autowired
	public SpittleController(DataClass data_interface) {　　　//构造注入，将
		this.interfaceName = data_interface;
	}

	@RequestMapping(method =HEEP_METHOD)
	public String spittles(Model model) {　　　//模型为key-value对的集合,会传递给视图,从而可以将数据渲染到客户端
		model.addAttribute(interfaceName.accessDateMethod());  //通过数据访问接口获得数据，传入模型中
		return "viewName";
	}
}
```
### Spring MVC测试：
* Spring包含可以针对控制器执行HTTP请求的机制的mock Spring MVC，因此可以不启动Web服务器和Web浏览器，对控制器进行测试：
```
/*利用Mock Spring MVC进行测试*/
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.view;
import static org.springframework.test.web.servlet.setup.MockMvcBuilders.standaloneSetup;
import org.springframework.test.web.servlet.MockMvc;
import org.junit.Test;
public class ControllerTestName {
	@Test
	public void testMethodName() throws Exception {
		TestClass TestClassName = new TestClass();  //创建测试对象实例
		MockMvc mockMvc = standaloneSetup(TestClassName).build();　　　//根据测试对象实例，构建MockMvc实例
		mockMvc.perform(get("/")).andExpect(view().name("expectedViewName"));  //使用MockMvc实例执行特定的http请求，并检验是否得到期望的视图逻辑名称
	}
}
```
```
/*对于传递模型数据到视图的控制器的测试*/
import static org.hamcrest.CoreMatchers.hasItems;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.when;
import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.model;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.view;
import static org.springframework.test.web.servlet.setup.MockMvcBuilders.standaloneSetup;
import org.springframework.test.web.servlet.MockMvc;
import org.springframework.web.servlet.view.InternalResourceView;
import java.util.ArrayList;
import java.util.Date;
import java.util.List;
import org.junit.Test;

public class ControllerTestName {
	@Test
	public void testMethodName() throws Exception {		
		SpittleRepository mockName = mock(interfaceName.class);  //创建访问接口的mock实现
		when(mockName.accessDateMethodName(..)).thenReturn(expectedData);   //打桩：配置接口方法的输入与输出

		ControllerClass controllerName = new ControllerClass(mockName);
		MockMvc mockMvc = standaloneSetup(controllerName)
				.setSingleView(new InternalResourceView("viewName")).build();　　　／／指定mock框架就不用解析控制器中的视图名（？？）

		mockMvc.perform(get("/url_path")).andExpect(view().name("viewName"))　　　//执行指定HTTP请求，断言控制器返回指定视图的名称
				.andExpect(model().attributeExists("param_name"))     //断言模型中包含指定名称属性
				.andExpect(model().attribute("param_name", hasItems(return_data));　　//断言预期的数据内容
	}
}
```

### 接受请求的输入
* Spring MVC允许以多种方式将客户端中的数据传送到控制器的处理器方法中,包括:
　－查询参数(Query Parameter)
　－表单参数(Form Parameter)
　－路径变量(Path Variable)
#### 查询参数：
* 处理查询参数:@RequestParam注解标注在控制器请求处理方法的传入参数声明前：value属性指定参数名称；defaultValue属性表示请求中参数不存在时以此设定参数默认值（给定的是String类型，会根据参数类型进行类型转换）
```
	@RequestMapping(method = GET/POST)
	public return_type processMethod(@RequestParam(value = "..", defaultValue = "..") type param, ..)
```
#### 路径参数：
* 通过路径参数接受输入:Spring MVC允许在@RequestMapping路径中添加占位符，占位符的名称要用大括号( “{” 和 “}” )括起来，路径中的其他部分要与所处理的请求完全匹配,但是占位符部分可以是任意的值;
* 控制器方法参数上添加了@PathVariable("param") 注解,表明在请求路径中,不管占位符部分的值是什么都会传递到处理器方法的param参数中（若方法的参数名与占位符的名称相同，可以去掉 @PathVariable中的value属性，自动匹配）
```
	@RequestMapping(value = "/{param_name}", method = RequestMethod.GET/POST)
	public return_type processMethod(@PathVariable("param_name") type param,...)
```
> （路径变量适合传递请求中少量的数据。若需要传递很多的数据(也许是表单提交的数据),可以使用表单参数等其他方法）
#### 处理表单
* 控制器对于POST请求的处理方法通常接受一个对象作为参数属性，将会使用请求中同名的参数进行填充（在方法中可以对数据对象进行简单的校验处理，可以调用数据访问接口的相关方法将数据对象保存）
```
public return_type processMethod(ClassName classInstance) 
```
> 注：当InternalResourceViewResolver看到视图格式中的“redirect:” 前缀时,会将其解析为重定向的规则，当它发现视图格式中以“forward:”作为前缀时,请求将会前往(forward)指定的URL路径,而不再是重定向。

### 校验表单
* 从Spring 3.0开始,在 Spring MVC中提供了对Java校验API的支持，只需保证在类路径下包含这个Java API的实现即可（如Hibernate Validator）。
* Java校验API定义了多个注解,这些注解可以放到数据对象类定义的属性上,从而限制这些属性的值(所有的注解都位于javax.validation.constraints包中)
　－@AssertFalse 所注解的元素必须是 Boolean 类型,并且值为false
　－@AssertTrue 所注解的元素必须是 Boolean 类型,并且值为true
　－@DecimalMax 所注解的元素必须是数字,并且它的值要小于或等于给定的BigDecimalString值
　－@DecimalMin 所注解的元素必须是数字,并且它的值要大于或等于给定的BigDecimalString值
　－@Digits 所注解的元素必须是数字,并且它的值必须有指定的位数
　－@Future 所注解的元素的值必须是一个将来的日期
　－@Max 所注解的元素必须是数字,并且它的值要小于或等于给定的值
　－@Min 所注解的元素必须是数字,并且它的值要大于或等于给定的值
　－@NotNull 所注解元素的值必须不能为null
　－@Null 所注解元素的值必须为null
　－@Past 所注解的元素的值必须是一个已过去的日期
　－@Pattern 所注解的元素的值必须匹配给定的正则表达式
　－@Size 所注解的元素的值必须是String、集合或数组,并且它的长度要符合给定的范围
> （此外Java 校验 API 的实现可能还会提供额外的校验注解，也可以定义自己的限制条件）
* 控制器的方法参数添加了@Valid注解:告知Spring需要确保这个对象满足校验限制，如果有校验出现错误的话,那么这些错误可以通过Errors对象（作为处理方法的参数，要紧跟在带有@Valid注解的参数后面）进行访问;
```
public class ClassName {
@NotNull(message = "...") 
@Size(message = "...",min = M, max = N)
private parameter1
.......
}
public return_type processMethod(@Valid Spittler spittler, Errors errors) { 
		if (errors.hasErrors()) { 
			......
		}
}
```

