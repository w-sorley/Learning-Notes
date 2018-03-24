---
title: JAX-RS学习记录
date: 2017-11-13
tags: [JAX-RS, Restful]
---
* REST:Representational State Transfer 的缩写,一种基于 HTTP，URI，以及 XML 这些现有的广泛流行的协议和标准的开发 Web 应用的架构风格，基于 REST 的 Web 服务遵循一些基本的设计原则:
    * 每一个对象(资源)都可以通过一个唯一的URI来进行寻址
    * 以遵循 RFC-2616 所定义的协议的方式显式地使用 HTTP 方法CRUD分别对应GET,PUT,POST,DElETE方法
    * 资源以使用不同的形式加以表示（如XML，JSON,具体的表现形式取决于访问资源的客户端),具体的表现形式取决于访问资源的客户端
# JAX-RS：Java API for RESTful Web Services
* JavaEE6引入了对JSR-311的支持,旨在定义一个统一的规范，使得Java程序员可以使用一套固定的接口来开发REST应用.JAX-RS 使用POJO编程模型和基于标注的配置，并集成了JAXB，从而可以有效缩短REST应用的开发周期。
* Resource 类和 Resource 方法
    * Web资源作为一个Resource类来实现(其中对资源的请求由Resource方法来处理),Resource类或Resource方法被打上了Path标注(指示一个用于资源定位的相对的URI路径，可被子类继承)
    * Resource类分为根Resource类和子Resource类,区别在于子Resource类没有打在类上的Path标注。
    * Resource类的实例方法打上了Path标注，则为Resource方法或子Resource定位器，区别在于子Resource定位器上没有标注HTTP请求类型
    *  Resource 方法参数的标注包括：@PathParam(用于将@Path中的模板变量映射到方法参数，模板变量支持使用正则表达式)、@MatrixParam、@QueryParam、@FormParam、@HeaderParam、@CookieParam、@DefaultValue 和 @Encoded(JAX-RS规定Resource方法中只允许有一个参数没有打上任何的参数标注，该参数称为实体参数，用于映射请求体)
    * Resource方法参数与返回值类型:
        * 方法参数类型:
            * 原生类型;
            * 构造函数接收单个字符串参数或者包含接收单个字符串参数的静态方法 valueOf 的任意类型;
            * List<T>，Set<T>，SortedSet<T>（T 为以上的 2 种类型）;
            * 用于映射请求体的实体参数
        *返回值类型:
            * void(状态码 204 和空响应体);
            * Response(status属性指定了状态码，entity属性映射为响应体);
            * GenericEntity(entity属性映射为响应体，entity属性为空则状态码为204，非空则状态码为200）;
            * 其它类型(返回的对象实例映射为响应体，实例为空则状态码为204，非空则状态码为200)
    * 对于错误处理，Resource方法可以抛出非受控异常WebApplicationException或者返回包含了适当的错误码集合的Response对象
    * Context标注:通过Context标注，根Resource类的实例字段可以被注入如下类型的上下文资源:Request、UriInfo、HttpHeaders、Providers、SecurityContext,HttpServletRequest、HttpServletResponse、ServletContext、ServletConfig
    * 内容协商(Content Negotiation)与数据绑定：
         * @Produces("application/json")：用于指定响应体的数据格式（MIME 类型）
         * @Consumes("application/json"):用于指定请求体的数据格式
         * (JAX-RS依赖于MessageBodyReader和MessageBodyWriter的实现来自动完成返回值到响应体的序列化以及请求体到实体参数的反序列化工作，可以使用 Provider标注来注册使用自定义的MessageBodyProvider,其中，XML格式的请求／响应数据与Java对象的自动绑定依赖于JAXB的实现)
* JAX-RS与JPA的结合:都使用了基于POJO和标注的编程模型,同一个POJO类上既有JAXB的标注，也有JPA的标注(或者还有Gson的标注),可以在JAX-RS与JPA之间得到复用
* 参考:
    * [JSR 311的规范文档](http://jcp.org/en/jsr/detail?id=311)
    * [JAX-RS的API文档](https://jsr311.dev.java.net/nonav/releases/1.0/index.html)
    * [JAX-RS 的参考实现Jersey](https://jersey.dev.java.net/)
    * Apache JAX-WS、JAX-RS 等的开源实现[Apache CXF](http://cxf.apache.org/)
    * 用于将 Java 对象与 JSON 格式的数据进行双向转换的开源类库[Google Gson](http://code.google.com/p/google-gson/)
    * Apache下提供了 JPA 规范的开源实现[Apache OpenJPA](http://openjpa.apache.org/)
    * [Java 技术专区](http://www.ibm.com/developerworks/cn/java/)




