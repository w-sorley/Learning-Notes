---
title: Spring MVC的高级技术
date: 2017-06-05
tags: [Spring MVC,framework,]
categories: FrameWork
---

# Spring MVC 的高级技术

## Spring MVC配置的替代方案  
在很多Spring 应用中除了DispatcherServlet以外,可能还需要额外的Servlet和Filter;可能还需要对DispatcherServlet本身做一些额外的配置，从而需要将DispatcherServlet配置到传统的 web.xml中

## 自定义DispatcherServlet配置
* 除了三个必须要重载的abstract方法，AbstractAnnotationConfigDispatcherServletInitializer还有更多的方法可以进行重载,从而实现额外的配置;其中customizeRegistration()方法将DispatcherServlet注册到Servlet容器中之后调用，并将Servlet注册后得到的Registration.Dynamic传递进来。通过重载该方法,可以对 DispatcherServlet进行额外的配置。如：
```
@Override
public void customizeRegistration(Dynamic registration){
    registration.setMultiparConfig(new MultipartConfigElement("/path"))   //启用multipart请求,设置上传文件的临时存储目录
}
```
* 借助ServletRegistration.Dynamic够完成多项任务，包括：
<br/>　－通过调用setLoadOnStartup()设置load-on-startup优先级
<br/>　－通过setInitParameter()设置初始化参数
## 添加其他的Servlet和Filter：
* 基于Java的初始化器(initializer)可以定义任意数量的初始化器类，因此如果想要添加其他的Servlet和Filter只需创建一个新的初始化器
```
//通过实现WebApplicationInitializer来注册Servle(最简单)
import javax.servlet.ServletContext;
import javax.servlet.ServletException
import javax.servlet.ServletRegistration.Dynamic
import org.springframework.web.webApplicationInitializer;
public MyServletInitializer implements WebApplicationInitializer{
    @Override
    pulic void onStarup(ServletContext servletContext) throws ServletException{
        Dynamic myServlet=servletContext.addServlet("myServlet" myServlet.class)   //注册servlet
        myServlet.addMapping("/custom/***")            //映射servlet
    }
}
```
* 类似地可以创建新的WebApplicationInitializer实现来注册Listener和Filter:
```
import javax.servlet.FilterRegistration.Dynamic;
public void onStartup(ServletContext servletContext) throws ServletException{
    Dynamic filter=servletContext.addFilter("myFlilter",myFilter,class);　　　　／／注册filter
    filter,addMappingForUrlPatterns(null,false,"/custom/**");    //添加filter映射路径
}
```
* 若只是注册Filter,并且该Filter只会映射到DispatcherServlet上，可以仅重载AbstractAnnotationConfigDispatcherServletInitializer的getServletFilters()方法来实现：
```
@Override
public Filter[] getServletFilter(){
    return new Filter[]{ new myFilter() }; //可以返回任意数量的Filter,无需声明映射路径，返回的所有 Filter 都会映射到 DispatcherServlet 上
}
```
## 在 web.xml中声明DispatcherServlet
```
<cotext-param>
    <param-name>contextConfigurationLocation</parm-name>
    <param-value>/WEB-INF/spring/root-context.xml</param-value>   <!-- 指定上下文配置文件 -->
</cotext-param>
<!-- 注册ContextLoaderListener -->
<listener>
<listener-class>org.springframework.web.context.ContextLoaderListener</listener-class>
</listener>
<!--  注册DispatcherServlet -->
<servlet>
    <servlet-name>dispatcherServlet</servlet-name>
    <servlet-class>org.springframework.web.servlet.DispatcherServlet</servlet-class>
    <!-- 指定DisparcherSrvlet上下文配置文件，可省略默认根据context-param属性 -->
    <init-param>
        <param-name>contextName</param-name>
        <param-value>/path</param-value>
    </init-param>　　　
    <load-on-startup>1</load-on-startup>
<servlet>

<servlet-mapping>
    <servlet-name>dispatcherServlet</servlet-name>   //设置DispatcherServlet映射路径为／
    <url-pattern>/</url-pattern>
</servlet-mapping>
//DispatcherServlet 会根据 Servlet 的名字找到一个文件,并基于该文件从上下文配置文件加载应用上下文
```
## 设置web.xml使用基于Java的配置
* 要在Spring MVC中使用基于Java的配置,需要告诉DispatcherServlet和ContextLoaderListener使用AnnotationConfigWebApplicationContext(是一个WebApplicationContext 的实现类,会加载Java配置类,而不是使用XML),通过设置contextClass上下文参数以及DispatcherServlet的初始化参数实现：
```
<context-param>
<param-name>contextClass<param-name>
<param-value>org.springframework.web.context.support.AnnotationConfigWebApplicationContext</param-value>  <!-- 指定使用Java配置 -->
</context-param>

<context-param>
<param-name>contextConfigLocation</param-name>  <!-- 指定根配置类 -->
<param-value>package.RootConfigClass</param-value>
</context-param>

<listener>
<listener-class>org.springframework.web.context.ContextLoaderListener</listener-class>
</listener>
<servlet>
    <servlet-name>appServletName<servlet-name>
    <servlet-class>org.springframework.web.servlet.DispatcherServlet<servlet-class>
    <init-param>
        <param-name>ContextClass</param-name>                <!-- 使用java配置 -->
        <param-value>org.springframework.web.context.support.AnnotationConfigWebApplicationContext</param-value>
    </init-param>
    <init-param>
    <param-name>contextConfigLocation</param-name>　　　<!-- 指定DispatcherServlet配置类 -->
    <param-value>package.WebConfig<param-value>
    </init-param>
    <load-on-startup>1<load-on-startup>
<servlet>
<servlet-mapping>...</servlet-mapping>  <!-- servlet映射路径 -->
```


## 处理multipart形式的数据
* 一般表单提交所形成的请求结果是以“&”符分割的多个name-value对,不适合对于传送二进制数据,而multipart格式的数据会将一个表单拆分为多个部分(part),每个部分对应一个输入域。在一般的表单输入域中,它所对应的部分中会放置文本型数据,但是如果上传文件的话,它所对应的部分可以是二进制。其中：ContentType:表示数据类型

### 配置multipart解析器

* DispatcherServlet将multipart的处理任务委托给了Spring中MultipartResolver策略接口的实现，通过这个实现类来解析multipart请求中的内容，Spring 3.1内置了两个 MultipartResolver 的实现：
<br/>　－CommonsMultipartResolver:　使用 Jakarta Commons FileUpload 解析 multipart 请求;
<br/>　－StandardServletMultipartResolver:　依赖于Servlet 3.0 对multipart请求的支持(始于 Spring 3.1)（优选使用Servlet 3.0 解析 multipart请求）
```
@Bean
public MultipartResolver throws IOException{
    return new StandardServletMultipartResolver();
}
```
* 必须要在web.xml或Servlet初始化类中,将multipart的具体细节作为DispatcherServlet配置的一部分,如:
* 采用Servlet初始化类的方式来配置DispatcherServlet,可以在Servlet registration上调用setMultipartConfig()方法,传入一个MultipartConfig-Element实例,如：
```
DispatecherServlet ds=new DispatcherServlet();
Dynamic registration=context.addServlet("appServlet",ds);
registration.addMapping("/");
registration.setMultipartConfig(new MultipartConfigElement("/path"))  //设置临时缓存路径path
```
* 如果配置DispatcherServlet的Servlet初始化类继承了AbstractAnnotationConfigDispatcherServletInitializer或AbstractDispatcher-ServletInitializer 的话可以通过重载 customizeRegistration() 方法(它会得到一个 Dynamic 作为参数)来配置 multipart 的具体细节：
```
@Override
protected void customizeRegistration(Dynamic registration){
    registration.setMultipartCongig(new MultipartConfigElement("/path"))      //设置临时缓存路径path
}
```
* 除了临时路径的位置,还有其他的构造器参数,如(path,file_size,request_size,buffued_size):
<br/>　－file_size:上传文件的最大容量(以字节为单位)。默认是没有限制的。
<br/>　－request_size:整个multipart请求的最大容量(以字节为单位),不会关心有多少个part以及每个part的大小。默认是没有限制的。
<br/>　－buffued_size:在上传的过程中,如果文件大小达到了一个指定最大容量(以字节为单位),将会写入到临时文件路径中。默认值为0,也就是所有上传的文件都会写入到磁盘上。
* 此外还可以使用web.xml文件来配置MultipartConfigElement:
```
<servlet>
    <servlet-name>appServlet</servlet-name>
    <servlet-class>org.springframework.web.servlet.Dispatecher<servlet-calss>
    <load-on-startup>1</load-on-startup>
    <multipart-config>
        <location>/path</location>      <!-- 临时缓存路径 -->
        <max-file-size>NUM</max-file-size>　　　<!-- 上传文件最大容量 -->
        <max-request-size>NUM</max-request-size>　　　<!-- 请求的最大容量 -->
    </multipart-config>
</servlet>
```
### 配置 Jakarta Commons FileUpload multipart 解析器
* Spring内置了CommonsMultipartResolver,可以作为 StandardServletMultipartResolver的替代方案(也可以自定义multipartresolver):
* 将CommonsMultipartResolver声明为Spring bean:
```
@Bean
public MultipartResolver multipartResolver() throws IOException {
    CommoonsMultipartResolver multipartResolver = new CommonsMultipartResolver();
    multipartResolver.setUploadTempDir(new FileSystemResource("/path"))  //可选，设定临时缓存路径
    multipartResolver.setMaxUploadSize(NUM);  //可选，设置上传文件最大容量
　　 multipartResolver.setMaxInMemorySize(NUM)  //可选，设置最大内存大小
    return multipartResolver;
}
```
* CommonsMultipart-Resolver不会强制要求设置临时文件路径。默认为Servlet容器的临时目录，可以通过设置uploadTempDir属性,将其指定为一个不同的位置

处理 multipart 请求
最常见的方式就是在某个控制器方法参数上添加@RequestPart注解来处理接收上传的文件，同时<form>标签的enctype属性设置为multipart/form-data，告诉浏览器以multipart数据的形式提交表单
```

<form method="POST" enctype="multipart/form-data">
....
<input type="file" name="file_name" accept="accept_file_type" />
....
</form>
@RequestMapping(value="/url_path",method=POST)    //byte数组中包含了请求中对应part的数据
public String ControllerMethodName(@RequestPart("file_name") byte[] fileName) {

}
```
接受 MultipartFile
Spring 还提供了 MultipartFile 接口,它为处理 multipart 数据提供了内容更为丰富的对象(使用上传文件的原始 byte 比较简单但是功能有限)
```
public interface MultiparfFile{
    String getName();
    String getOriginalFileName();
    String getContentType();
    boolean isEmpty();
    long getSize();
    byte[] getBytes() throws IOException;
    void transferTo(File dest) throws IOException;
}
```
将文件保存到云端，如将文件保存到Amazon S3：
```
private savaMultipartImage(MultipartFile filename) throws ImageUploadException{
    try{
        AWSCredentials  awsCredentials=new AWSCredentials(s3Accesskey,s2SecretKey); //创建AWS凭证
        S3Service s3=new RestS3Service(awscredentials);
        S3Buket bucket=s3.getBuket("filename");
        S3Object fileObject=new S3OBject(fimename.getOriginalFileName());
        fileObject.setDataInptstream(filename.getInputstream())    //设置文件数据
        fileObject.setContentLength(filename.getSize());
        fileObject.setContentType(filename.getContentType());

        AccessControlList acl=new AccessControlList();       //设置权限
        acl.setOwner(bucket.getOwner());
        acl.grantPermission(GroupGrantee.ALL_USERS,Permission.PERMISSION_READ);
        fileObject.setAcl(acl);
        s3.putObject(bucket,fileObject) //保存图片
    }catch(Exception e){
        thrown  new Image_UoloadException("information");
    }

}
```


### 以Part的形式接受上传的文件
* Spring MVC 也能接受javax.servlet.http.Part作为控制器方法的参数,
```
public interface Part{
    public InputStream getInputStream() throws IOException;
    public String getContentType();
    public String getName();
    public String　SubmittedFileName();  //获得原始文件名
    public long getSize();
    public void write(String filename) throws Exception;　　//写入文件系统
    public void delete() throw IOException;
    public String   getHeader(String name);
    public Collection<String> getHeader(String name);
    public Collection<String> getHeaderNames();
}
```

* 如果使用Part来替换MultipartFile,那么processRegistration()的方法签名将会变成如下的形式（此时不必配置MultipartResolver）:
```
RequestMapping(value="/URL_path",method=POST)
public String ControllerMethod(@RequestPart("filename") Part filename){   
}
```

## 处理异常
* Servlet请求的输出为一个Servlet响应。如果在请求处理的时候,出现了异常,常必须要以某种方式转换为响应，其中Spring提供了多种方式将异常转换为响应：
<br/>　－特定的Spring异常将会自动映射为指定的HTTP状态码;
<br/>　－异常上可以添加@ResponseStatus注解,从而将其映射为某一个HTTP状态码;
<br/>　－在方法上可以添加@ExceptionHandler注解,使其用来处理异常
#### 1.将异常映射为HTTP状态码
* 在默认情况下, Spring会将自身的一些异常自动转换为合适的状态码,如：
<br/>　－BindException:400-Bad Request
<br/>　－ConversionNotSupportedException:500-Internal Server Error
<br/>　－HttpMediaTypeNotAcceptableException:406-Not Acceptable
<br/>　－HttpMediaTypeNotSupportedException:415-Unsupported Media Type
<br/>　－HttpMessageNotReadableException:400-Bad Request
<br/>　－HttpMessageNotWritableException:500-Internal Server Error
<br/>　－HttpRequestMethodNotSupportedException:405-Method Not Allowed
<br/>　－MethodArgumentNotValidException:400-BadRequest
<br/>　－MissingServletRequestParameterException:400 - Bad Request
<br/>　－MissingServletRequestPartException:400 - Bad Request
<br/>　－NoSuchRequestHandlingMethodException:404 - Not Found
<br/>　－TypeMismatchException:400 - Bad Request
* 2.此外，Spring提供了一种机制,能够通过@ResponseStatus注解将异常映射为HTTP状态码(默认响应一般都会带有500状态码)，如：
```
import org.springframework.http.HttpStatus;
import org.spirngframework.web.bin.annotation.ResponseStatus;
@ResponseStatus(value=HttpStatus.NOT_FOUND,reason="output_information")
public class ExpectionName extends RuntimeException{        //将ExceptionName异常映射为404NOT_FOUND  
}
```
* 3.编写异常处理的方法
* 异常处理方法上添加了@ExceptionHandler注解,当抛出指定的异常的时,将会委托该方法来处理：
```
@ExceptionHandler(ExceptionName.class)
public String ExceptionHandlerMethod(){
    return ".."
}
```
@ExceptionHandler 注解标注的方法,能处理同一个控制器中所有处理器方法所抛出的异常,此外将其定义到控制器通知类中能够处理所有控制器中处理器方法所抛出的异常
## 为控制器添加通知
* 如果控制器类的特定切面能够运用到整个应用程序的所有控制器中，如希望一个异常处理方法能够处理所有控制器的特定异常，为了避免重复,我们会创建一个基础的控制器类,所有控制器类要扩展这个类,从而继承通用的@ExceptionHandler方法，此外Spring 3.2为这类问题引入了一个新的解决方案:控制器通知。
* 控制器通知(controller advice)是任意带有@ControllerAdvice注解的类,这个类会包含一个或多个如下类型的方法:
<br/>　－@ExceptionHandler 注解标注的方法;
<br/>　－@InitBinder 注解标注的方法;
<br/>　－@ModelAttribute 注解标注的方法。
* 上所述的这些方法会运用到整个应用程序所有控制器中带有@RequestMapping注解的方法上,此外ControllerAdvice注解本身已经使用了@Component,因此@ControllerAdvice注解所标注的类将会自动被组件扫描获取到。

## 跨重定向请求传递数据
* “redirect:”前缀能够用来指导浏览器进行重定向,此外Spring为重定向功能还提供了一些其他的辅助功能：
* 一般来讲,当一个处理器方法完成之后,该方法所指定的模型数据将会复制到请求中,并作为请求中的属性,请求会转发( forward )到视图上进行渲染。因为控制器方法和视图所处理的是同一个请求,所以在转发的过程中,请求属性能够得以保存。
* 当控制器的结果是重定向的话,原始的请求就结束了,并且会发起一个新的 GET 请求。原始请求中所带有的模型数据也就随着请求一起消亡了。对于重定向来说,模型并不能用来传递数据，需要利用其他方法能够从发起重定向的方法传递数据给处理重定向方法中：
<br/>　－使用URL模板以路径变量和查询参数的形式传递数据;
<br/>　－通过flash属性发送数据。

### 通过URL模板进行重定向
* 通过路径变量和查询参数传递数据:除了连接String的方式来构建重定向URL,Spring还提供了使用模板的方式来定义重定向URL：return “／.../｛variable｝”

### 使用flash属性
* 会话能够长期存在,并且能够跨多个请求,因此将 Spitter 放到会话并在重定向后,从会话中将其取出，在重定向后在会话中将其清理掉，可以实现一些大的的数据（如整个对象）在重定向时的传递.
* Spring提供了将数据发送为flash属性(flash attribute)的功能,flash 属性会一直携带这些数据直到下一次请求,然后才会消失。
* Spring 提供了通过 RedirectAttributes 设置 flash 属性的方法:
```
@RequestMapping(value="/URL_path",method=POST/GET)
public String controllerMethod(..Model model){
    model.addFlashAttribute("ObjectName",Onject); //添加传递数据对象
    if(!model.containsAttribute("ObjectName"))　　//查看是否存在制定对象
}
`
```