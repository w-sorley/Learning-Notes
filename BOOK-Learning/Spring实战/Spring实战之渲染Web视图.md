---
title: Spring实战之渲染Web视图.
date: 2017-06-02
tags: [Spring, MVC, Spring MVC,Web,framework]
categories: FrameWork
---

# Spring实战之渲染Web视图
## 理解视图解析
* 将控制器中请求处理的逻辑和视图中的渲染实现解耦是Spring MVC的一个重要特性:
<br/>　－Spring MVC定义了一个名为ViewResolver的接口传入一个视图名和Locale对象时,它会返回一个View实例
<br/>　－view为一个接口接受模型以及Servlet的request和response对象,并将输出结果渲染到response中，展现到客户端
<br/>　－一般只需编写ViewResolver和View的实现,将要渲染的内容放到response中,进而展现到用户的浏览器中
* Spring提供了多个内置ViewResolver的实现：
<br/>　－BeanNameViewResolver：将视图解析为Spring应用上下文中的bean,其中bean的ID与视图的名字相同
<br/>　－ContentNegotiatingViewResolver：通过考虑客户端需要的内容类型来解析视图,委托给另外一个能够产生对应内容类型的视图解析器
<br/>　－FreeMarkerViewResolver：将视图解析为FreeMarker模板
<br/>　－InternalResourceViewResolver：将视图解析为Web应用的内部资源(一般为JSP)
<br/>　－JasperReportsViewResolver：将视图解析为JasperReports定义
<br/>　－ResourceBundleViewResolver：将视图解析为资源bundle(一般为属性文件)
<br/>　－TilesViewResolver:将视图解析为Apache Tile定义,其中tile ID与视图名称相同(注:有两个不同的TilesViewResolver实现(Tiles 2.0 和Tiles 3.0))
<br/>　－UrlBasedViewResolver:直接根据视图的名称解析视图,视图的名称会匹配一个物理视图的定义
<br/>　－VelocityLayoutViewResolver 将视图解析为Velocity布局,从不同的Velocity模板中组合页面
<br/>　－VelocityViewResolver:将视图解析为Velocity模板
<br/>　－XmlViewResolver:将视图解析为特定XML文件中的bean定义。类似于BeanNameViewResolver
<br/>　－XsltViewResolver 将视图解析为XSLT转换后的结果
（上述13种视图解析器几乎每一项都对应Java Web应用中特定的某种视图技术,其中Spring 3.1不支持TilesViewResolver视图解析器）

## 创建JSP视图
* Spring提供了两种支持JSP视图的方式:
* InternalResourceViewResolver:会将视图名解析为JSP文件。
 > （如果JSP页面中使用了JSP标准标签库(JavaServer Pages Standard Tag Library,JSTL),InternalResourceViewResolver能够将视图名解析为JstlView形式的JSP文件,从而将JSTL本地化和资源bundle变量暴露给JSTL的格式化(formatting)和信息(message)标签。)
* Spring 提供了两个 JSP 标签库,一个用于表单到模型的绑定,另一个提供了通用的工具类特性

### 配置适用于JSP的视图解析器
* InternalResourceViewResolver采取的方式遵循一种约定,在视图名上添加前缀和后缀,进而确定一个 Web应用中视图资源的物理路径
> (一些视图解析器如ResourceBundleViewResolver会直接将逻辑视图名映射为特定的View接口实现)
* 此外还可以基于XML的Spring配置的方式配置InternalResourceViewResolver
```
<bean id="viewResolver" class="org.springframework.web.servlet.view.InternalResourceViewResolver" p:prefix="前缀"　p:suffix="后缀"　／>
```
### 解析JSTL视图
* JSTL的格式化标签需要一个Locale对象,以便于恰当地格式化地域相关的值,如日期和货币。
* 信息标签可以借助Spring的信息资源和Locale,从而选择适当的信息渲染到HTML之中。
* 通过解析JstlView ,JSTL能够获得Locale对象以及Spring中配置的信息资源。
* 如果想让InternalResourceViewResolver将视图解析为JstlView ,而不是InternalResourceView的话,只需设置它的viewClass属性即可：InternalResourceViewResolver.setViewClass(org.springframework.web.servlet.view.JstlView.class),从而能确保 JSTL 的格式化和信息标签能够获得 Locale 对象以及 Spring 中配置的信息资源
* 同样也可以使用XML配置：
```
<bean id="viewResolver" class="org.springframework.web.servlet.view.InternalResourceViewResolver"
 p:prefix="前缀"　p:suffix="后缀"　p:viewClass="org.springframework.web.servlet.view.JstlView.class"　　／>
```
### 使用Spring的JSP库
* 标签库能够避免在脚本块中直接编写Java代码,实现为JSP添加功能.其中Spring提供了两个JSP标签库:
* 一个可以绑定模型model中的某个属性,用来渲染HTML表单,从而表单就可以预先填充属性值,并且在表单提交失败后,能够展现校验错误（一个标签库包含了一些工具类标签）
#### 将表单绑定到模型上:
* 首先在JSP页面声明表单绑定库：
```
<%@ taglib uri="http://www.springframework.org/tags/form" prefix="sf" %>
```
* Spring的表单绑定JSP标签库包含了14个标签:
<br/>　－\<sf:checkbox\>：渲染成一个HTML\<input\>标签,其中type属性设置为checkbox
<br/>　－\<sf:checkboxes\>：渲染成多个HTML\<input\>标签,其中type属性设置为checkbox
<br/>　－\<sf:hidden\>：渲染成一个HTML\<input\>标签,其中type属性设置为hidden
<br/>　－\<sf:input\>：渲染成一个HTML\<input\>标签,其中type属性设置为text
<br/>　－\<sf:password\>：渲染成一个HTML\<input\>标签,其中type属性设置为password
<br/>　－\<sf:radiobutton\>：渲染成一个HTML\<input\>标签,其中type属性设置为radio
<br/>　－\<sf:radiobuttons\>：渲染成多个HTML\<input\>标签,其中type属性设置为radio
<br/>　－\<sf:errors\>：在一个HTML\<span\>中渲染输入域的错误
<br/>　－\<sf:form\>：渲染成一个HTML\<form\>标签,并为其内部标签暴露绑定路径,用于数据绑定
<br/>　－\<sf:label\>：渲染成一个HTML\<label\>标签
<br/>　－\<sf:option\>：渲染成一个HTML\<option\>标签,其selected属性根据所绑定的值进行设置
<br/>　－\<sf:options\>：按照绑定的集合、数组或Map,渲染成一个HTML\<option\>标签的列表
<br/>　－\<sf:select\>：渲染为一个HTML\<select\>标签
<br/>　－\<sf:textarea\>：渲染为一个HTML<\textarea\>标签
* 应用举例:
```
<sf:form method="POST/GET" commandName="binded_class_instance_name">   //commandName属性构建针对某个模型对象的上下文信息,确定需要绑定的模型对象（指定为模型中对象的key值，可
　                                                                     //能需要在相应的控制器中传入模型参数，并设置对象）
<sf:label path="bind_attribute" cssErrorClass="error">outputWord1:</sf:label><sf:input path="bind_attribute" /> <br/>　　　　　　　　　　
 //<sf:label>标签，用path来指定绑定模型对象中的属性，没有校验错误渲染为普通<label>元素
 //cssErrorClass属性，指定所绑定的属性有错误,渲染得到的<label>class 属性将会被设置为error,同样适用于<sf:input>标签
<sf:errors path="bind_attribute">  //如果有校验错误会在一个HTML<span>标签中显示错误信息（校验注解message属性指定,可使用大括号指定为属性文件中的某一个属性）,path="*"可显示所有绑定属性的输入错误信息,此外添加//element="div"可渲染为<div>标签，属性文件ValidationMessages.properties放在根类路径之下，可使用{min}｛max}引用@Size注解中设置的最值
outputWord2:<sf:password path="bind_attribute" /> <br/>   //input标签的value属性值将会设置为模型对象中 path 属性所对应的值，单独显示
.....
</sf:form>
//注：从Spring3.1开始,<sf:input>标签能够允许我们指定type属性,除了其他可选的类型外,还能指定 HTML5特定类型的文本域,如date、range和email
```
#### Spring通用的标签库：
* Spring 通用的标签库声明:
```
<%@ taglib uri="http://www.springframework.org/tags" prefix="s" %>
```
* Spring的JSP标签库中提供了多个便利的标签,包括一些遗留的数据绑定标签:
<br/>　－\<s:bind\>:将绑定属性的状态导出到一个名为status的页面作用域属性中,与\<s:path\>组合使用获取绑定属性的值
<br/>　－\<s:escapeBody\>:将标签体中的内容进行HTML和/或JavaScript转义
<br/>　－\<s:hasBindErrors\>:根据指定模型对象(在请求属性中)是否有绑定错误,有条件地渲染内容
<br/>　－\<s:htmlEscape\>:为当前页面设置默认的HTML转义值
<br/>　－\<s:message\>:根据给定的编码获取信息,然后要么进行渲染(默认行为),要么将其设置为页面作用域、请求作用域、会话作用域或应用作用域的变量(通过使用var和scope属性实现)
<br/>　－\<s:nestedPath\>:设置嵌入式的path,用于\<s:bind\>之中
<br/>　－\<s:theme\>:根据给定的编码获取主题信息,然后要么进行渲染(默认行为),要么将其设置为页面作用域、请求作用域、会话作用域或应用作用域的变量(通过使用var和scope属性实现)
<br/>　－\<s:transform\>:使用命令对象的属性编辑器转换命令对象中不包含的属性
<br/>　－\<s:url\>:创建相对于上下文的URL,支持URI模板变量以及HTML/XML/JavaScript转义。可以渲染URL(默认行为),也可以将其设置为页面作用域、请求作用域、会话作用域或应用作用域的变量(通过使用var和scope属性实现)
<br/>　－\<s:eval>:计算符合Spring表达式语言(Spring Expression Language,SpEL)语法的某个表达式的值,然后要么进行渲染(默认行为),要么将其设置为页作用域、请求作用域、会话作用域或应用作用域的变量(通过使用var和scope属性实现)
##### 展现国际化信息
* 文本能够位于一个或多个属性文件中，从而借助<s:message>,我们可以将硬编码的一些信息替换为<s:message code="Class.attribute"/>从而实现根据key为Class.attribute的信息源渲染文本(需要配置相应的信息源)
* Spring 有多个信息源的类，它们都实现了 MessageSource 接口:
* ResourceBundleMessageSource类会从一个属性文件（属性文件的名称是根据基础名称(base name)衍生而来）中加载信息，配置如下：
```
@Bean
public MessageSource messageSource(){
    ResourceBoundleMessageSource messageSource=new ResourceBoundleMessageSource;
    messageSource.setBasename("base_name")  //设置基础名称base name,会根据其得到属性文件的名称
}
```
* 此外ReloadableResourceBundleMessageSource类工作方式与ResourceBundleMessageSource类似,但是它能够重新加载信息属性,而不必重新编译或重启应用；
> 注:basename属性可以设置为在类路径下(以“classpath:”作为前缀)、文件系统中(以“file:”作为前缀)或Web应用的根路径下(没有前缀)查找属性
##### 创建URL
* \<s:url\>标签主要用来创建 URL,然后将其赋值给一个变量或者渲染到响应中,是JSTL中\<c:url\>标签的替代,如：
```
//接受一个相对于Servlet上下文的URL,并在渲染的时候,预先添加上Servlet上下文路径
<a href="<s:url href="/source_path"/>">..</a>    //若当前servlet上下文为/basepath则会渲染为<a href="/basepath/source">
//可以使用<s:url> 创建 URL ,并将其赋值给一个变量供模板在稍后使用
<s:url href="/moudle_path" var="moudle_URL" />  //默认作用在页面作用域内，可通过scope属性作用在应用作用域内、会话作用域内或请求作用域内
<a href="${moudle_URL}"></a>
//在URL上添加参数
<s:url href="/module_path/{param_name}" var="moudle_URL"htmlEscape＝“y=true”  //htmlEscape属性为true指示URL转义将渲染得到的URL内容展现在Web页面上
javaScriptEscape＝”true“　> //javaScriptEscape属性为true,指示可以在JavaScript代码中使用URL
<s:param name="param_name" value=".." />  //当href属性中的占位符匹配<s:param>中所指定的参数时,参数会插入到占位符的位置中,否则参数作为查询参数
</s:url>
```
##### 转义内容
<s:escapeBody>标签为通用的转义标签，会渲染标签体中内嵌的内容,并且在必要的时候进行转义，如：
```
／／将HTML转义后原样输出到浏览器
<s:escapeBody javaScriptEscape＝”true“>  //javaScriptEscape属性为true,指示支持javaScript转义
<h1>........<h1>  //会转义为“&lt;h1&gt;.......&lt;/h1&gt;
<s:escapeBody>
```
## 使用Apache Tiles视图定义布局:
Apache Tiles为一种布局引擎，可以定义适用于所有页面的通用页面布局，同时Spring MVC以视图解析器的形式为Apache Tiles提供了支持,能够将逻辑视图名解析为Tile定义；
### 配置 Tiles 视图解析器:
* TilesConfigurer bean,负责定位和加载Tile定义并协调生成Tiles,并与Apache Tiles协作
```
@Bean
TilesConfigurer tilesConfigurer(){
    TilesConfigurer tiles=new TilesConfigurer();
    tiles.setDefinitions(new String[]{"/WEB-INF/layout/titles.xml"});  //指定title定义的位置，可指定多个，可使用通配符
    titles.setCheckRefresh(true);       //启用刷新功能
    return tiles;
}

```

针对Apache Tiles 2的TilesConfigurer/TilesViewResolver位于org.springframework.web.servlet.view.tiles2 包中,而针对 Tiles 3 的组件位于 org.springframework.web.servlet.view.tiles3 包中
* TilesViewResolver bean将逻辑视图名称解析为Tile定义的视图（通过查找与逻辑视图名称相匹配的Tile定义）:
```
//spring3
@Bean
public ViewResolver viewResolver(){
    return new ViewResolver();
}
//spring4
@Override
public void configureViewResolvers(ViewResolverRegistry registry) {
    TilesViewResolver viewResolver = new TilesViewResolver();
    registry.viewResolver(viewResolver);
}

```
* 此外还可以在XML文件中配置TilesConfigurer和TilesViewResolver：
```
<bean id="titlesConfigurer" class="org.springframework.web.servlet.view.title3.TitlesConfigurer">
    <property name="definitions">
        <list>
            <value>/WEB-INF/layout/titles.xml</value>
            .....
        </list>
    </property>
</bean>

<bean id="viewResolver" class="org.springframework.web.servlet.view.tiles3.TilesViewResolver" />
```
#### 定义 Tiles
* Apache Tiles提供了一个文档类型定义(document type definition , DTD),用来在XML文件中指定Tile的定义,每个定义中需要包含一个\<definition\>元素,这个元素会有一个或多个\<put-attribute\>元素:
```
<tiles-definitions>
    <definition name=“tiles_name” extends="other_defined_tiles" template="/WEB-INF/layout/moudleName.jsp">
        <put-attribute name="moduleName" value="/path/fileName.jsp">    //在JSP文件中使用Tile标签库中的 <t:insert Attribute name="moduleName">标签来插入其他的模板
        .....                     //extend扩展Tiles,会继承其它Tiles的attribute定义，也可以覆盖
    </definition>
    .....
</tiles-definitions>
```
## Thymeleaf
* JSP缺点：JSP模板都是采用HTML的形式,同时又掺杂上各种JSP标签库的标签,使其变得很混乱缺乏良好格式。同时JSP规范与Servlet规范紧密耦合，意味着它只能用在基于Servlet的 Web应用之中，不能作为通用的模板，也不能用于非Servlet的Web应用。
* Thymeleaf模板：原生的,不依赖于标签库，能在接受原始HTML的地方进行编辑和渲染。没有与Servlet规范耦合，通用性更强.
### 配置 Thymeleaf 视图解析器
* 在Spring中使用Thymeleaf ,需配置三个启用Thymeleaf与Spring集成的bean:
```
//注入applicationContext
@Override
public void setApplicationContext(ApplicationContext applicationContext) {
    this.applicationContext = applicationContext;
}
```
* ThymeleafViewResolver :将逻辑视图名称解析为 Thymeleaf 模板视图
```
@Bean
public ViewResolver viewResolver(SpirngTemplateEngine templateEngine)P{
    ThymeleafViewResolver viewResolver=new ThymeleafViewResolver();
    viewResolver.setTemplateEngine(templateEngine); //注入了一个对SpringTemplateEngine bean的引用,会在Spring中启用Thymeleaf引擎,来解析模板，并基于这些模板渲染结果
    return viewResolver;
}
//sping4 thymeleaf 3
@Bean
public ViewResolver htmlViewResolver() {
    ThymeleafViewResolver resolver = new ThymeleafViewResolver();
    resolver.setTemplateEngine(templateEngine(htmlTemplateResolver()));
    resolver.setContentType("text/html");
    resolver.setCharacterEncoding(UTF8);
    resolver.setViewNames(array("*.html"));
    return resolver;
}
```
> 注:ThymeleafViewResolver是Spring MVC中ViewResolver的一个实现类,像其他的视图解析器一样,它会接受一个逻辑视图名称,并将
其解析为视图，在此,视图是一个Thymeleaf 模板

* SpringTemplateEngine :处理模板并渲染结果;
```

@Bean
public TemplateEngine templateEngine(TemplateResolver templateResolver){
    SpringTemplateEngine templateEngine=new SpringTemplateEngine();
    templateEngine.setTemplateResolver(templateResolver);　//注入了一个TemplateResolver bean的引用,TemplateResolver会最终定位和查找模板
    return templateEngine;
}
```
* TemplateResolver :加载 Thymeleaf 模板
```
@Bean
public TemplateResolver templateResolver(){
    TemplateResolver templateResolver=new TemplateResolver();
    templateResolver.setPrefix("WEB-INF/template/");　／／前缀后缀与逻辑视图名组合，用来查找定位实际模板视图
    templateResolver.setSuffix(".html");
    templateResolver.setTemplateMode("HTML5");  //templateMode 属性设置为HTML 5,表明预期要解析的模板会渲染成HTML5输出
    return templateResolver;
}
//spring 4 thymeleaf 3
private ITemplateResolver htmlTemplateResolver() {
    SpringResourceTemplateResolver resolver = new SpringResourceTemplateResolver();
    resolver.setApplicationContext(applicationContext);
    resolver.setPrefix("/WEB-INF/templates/");
    resolver.setTemplateMode(TemplateMode.HTML);
    return resolver;
}

```
* 或使用xml文件配置相关的bean:
```
<!--- 配置模板解析器 ---->
<bean id="templateResolver" class="org.thymeleaf.templateresolver.ServletContextTemplateResolver"
        p:prefix="/WEB_INF/template/"
        p:suffix=".html"
        p:templateMode=“HTML5” />

<!--- 配置引擎 --->
<bean id="templateEngine" class="org.thymeleaf.spring3.SpringgTemplateEngine"
        p:templateResolver-ref="templateResolver" />
<!--- 配置Thymeleaf视图解析器 --->
<bean id="viewResolver" class="org.thymeleaf.spring3.view.ThymeleafViewResolver"
        p:templateEngine-ref="templateEngine" />
```




### 定义Thymeleaf模板
* Thymeleaf在很大程度上就是HTML文件,与JSP不同,没有什么特殊的标签或标签库。通过自定义的命名空间,为标准的HTML标签集合添加Thymeleaf属性,如:
```
<html xmlns="http://www.w3.org/1999/xhtml"
      xmlns="http://www.thymeleaf.org">　　　<!--- 声明Thymeleaf命名空间  --->
<head>
<title>....</title>
<link rel="..." type="..."
        th:href="@{/resource/style.css}"></link>      <!--- 到样式表的th:href链接  --->
</head>
<body>
........
<a th:href="@{/path}">..</a>   <!--- 到页面的th:href链接，@{}表达式,用来计算相对于URL的路径  --->
........
</body>
</html>
```
> 注：th:href属性与对应的原生HTML属性href类似,并且可以按照相同的方式来使用.不同之处在于th:href的值中可以包含Thymeleaf表达式,用来计算动态的值。最终它会渲染成一个标准的href属性,其中包
含渲染时动态创建得到的值：这是 Thymeleaf 命名空间中很多属性的运行方式（对应标准的HTML属性,并且具有相同的名称,但是会渲染一些计算后得到的值），
* 借助Thymeleaf实现表单绑定
表单绑定:能够将表单提交的数据填充到命令对象中,并将其传递给控制器,而在展现表单的时候,表单中也会填充命令对象中的值,是Spring MVC的一项重要特性
```                   <!--- 检查fieldName域有没有校验错误,如果有则class属性在渲染时的值为error,没有则不会渲染class属性 --->
<lael th:class="${#fields.hasErrors('fieldName')}? 'error'">...</label>  <!--- th:class属性会渲染为一个class属性,它的值是根据给定的表达式计算得到的 --->
<input type="text" th:field="*{fieldName}" th:class="${#fields.hasErrors('fieldName')}? 'error'"/>　　<!-- th:field属性引用后端对象的fieldName域 -->
```
th:if 属性来检查是否有校验错误
<li> 标签上的 th:each 属性将会通知 Thymeleaf 为每项错误都渲染一个 <li> ,在每次迭代中会将当前错误设置到一个名为 err 的变量中
th:text 属性。这个命令会通知 Thymeleaf 计算某一个表达式(在本例中,也就是 err 变量)并将它的值渲染为 <li> 标签的内容体。实际上的效果就是每项错误对应一个 <li> 元素,并展现错误的文本。
“${}” 表达式(如 ${spitter} )是变量表达式( variable expression )。一般来讲,它们会是对象图导航语言( Object-Graph Navigation Language , OGNL )表达式但在使用 Spring 的时候,它们是 SpEL 表达式
“*{}” 表达式,它们是选择表达式( selection expression )。变量表达式是基于整个 SpEL 上下文计算的,而选择表达式是基于某一个选中对象计算的









