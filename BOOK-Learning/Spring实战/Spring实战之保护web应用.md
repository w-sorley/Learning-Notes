---
title: 保护Web应用
date: 2017-06-05
tags: [Spring MVC,framework,]
categories: FrameWork
---

# 保护Web应用

## Spring Security简介
* Spring Security是为基于Spring的应用程序提供声明式安全保护的安全性框架，充分利用了依赖注入(dependency injection,DI)和面向切面的技术，提供了完整的安全性解决方案,它能够在Web请求级别和方1法调用级别处理身份认证和授权。
* Spring Security引入了全新的、与安全性相关的XML命名空间,连同注解和一些合理的默认设置简化了安全性的配置，Spring Security 3.0 融入了 SpEL ,这进一步简化了安全性的配置
* 它使用Servlet规范中的Filter保护Web请求并限制URL级别的访问，使用Spring AOP保护方法调用（借助于对象代理和使用通知）能够确保只有具备适当权限的用户才能访问安全保护的方法

### 理解 Spring Security 的模块
* Spring　Security 3.2分为11个模块：
<br/>　－ACL：支持通过访问控制列表(access control list,ACL)为域对象提供安全性
<br/>　－切面(Aspects)：一个很小的模块,当使用Spring Security注解时,会使用基于AspectJ的切面,而不是使用标准的Spring AOP
<br/>　－CAS客户端(CAS Client)：提供与Jasig的中心认证服务(Central Authentication Service , CAS)进行集成的功能
<br/>　－配置(Configuration)：包含通过XML和Java配置Spring Security的功能支持
<br/>　－核心(Core)：提供Spring Security基本库
<br/>　－加密(Cryptography)：提供了加密和密码编码的功能
<br/>　－LDAP：支持基于LDAP进行认证
<br/>　－OpenID：支持使用OpenID进行集中式认证Remoting提供了对Spring Remoting的支持
<br/>　－标签库(Tag Library)：Spring Security的JSP标签库
<br/>　－Web：提供了Spring Security基于Filter的Web安全性支持
> (应用程序的类路径下至少要包含 Core 和 Configuration 这两个模块)

### 过滤Web请求
* DelegatingFilterProxy 是一个特殊的 Servlet Filter,将工作委托给一个作为<bean>注册在Spring应用的上下文中javax.servlet.Filter实现类,会拦截发往应用中的请求,并将请求委托给ID为 springSecurityFilterChain的bean
* 可以在web.xml中进行配置
```
<filter>
<filter-name>springSecurityFilterChain</filter-name>
<filter-class>org.springframework.security.web.context.AbstractSecurityWebApplicationInitializer</filter-class>
</filter>
```
* 也可以借助WebApplicationInitializer以Java的方式来配置
```
import org.springframework.security.web.context.AbstractSecurityWebApplicationInitializer
public class SecurityWebInitlializer extends AbstractSecurityWebApplicationInitializer{}
//AbstractSecurityWebApplicationInitializer 实现了 WebApplication-Initializer ,因此 Spring 会发现它,并用它在 Web容器中注册DelegatingFilterProxy
```
* springSecurityFilterChain是另一个特殊的Filter,也被称为FilterChainProxy.可以链接任意一个或多个其他的Filter，从而Spring Security依赖一系列Servlet Filter来提供不同的安全特性

### 编写简单的安全性配置
```
import org.springframework.conetext.annotation.Configuration;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
@Configuration
@EnableWebSecurity  //启用Web安全功能,若使用Spring MVC应使用@EnableWebMvcSecurity替代它
public class SecurityConfig extends WebSecurityConfigurerAdapter{}  //Spring Security必须配置在一个实现了WebSecurityConfigurer的bean中
```
> 注：在Spring应用上下文中,任何实现了WebSecurityConfigurer的bean都可以用来配置Spring Security,但是最为简单的方式是扩展WebSecurityConfigurerAdapter类
<br/>此外@EnableWebMvcSecurity注解还配置了一个Spring MVC参数解析解析器(argument resolver),因此处理器方法就能够通过带有@AuthenticationPrincipal注解的参数获得认证用户的principal (或username)。同时还配置了一个bean,在使用Spring表单绑定标签库来定义表单时,这个bean会自动添加一个隐藏的跨站请求伪造(cross-site request forgery,CSRF)token输入域。
* 可以通过重载 WebSecurityConfigurerAdapter 的三个 configure() 方法来配置 Web 安全性，使用传递进来的参数设置行为:
<br/>　－configure(WebSecurity)：通过重载,配置Spring Security的Filter链
<br/>　－configure(HttpSecurity)：通过重载,配置如何通过拦截器保护请求
<br/>　－configure(AuthenticationManagerBuilder)：通过重载,配置user-detail服务


* Spring Security能够基于各种数据存储来认证用户。它内置了多种常见的用户存储场景,如内存、关系型数据库以及LDAP，此外也可以编写并插入自定义的用户存储实现

## 用户查询信息服务选择
### 如配置Spring Security使用内存用户存储:
```
import org.springframework.conetext.annotation.Configuration;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.web.configuration.EnableWebSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;
@Configuration
@EnableWebSecurity  //启用Web安全功能,若使用Spring MVC应使用@EnableWebMvcSecurity替代它
public class SecurityConfig extends WebSecurityConfigurerAdapter{  //Spring Security必须配置在一个实现了WebSecurityConfigurer的bean中
    @Override
    protected void configure(AuthenticationManagerBuilder auth) throws Exception{
        auth.
            inMemoryAuthentication()    //启用内存用户存储
            .withUser("user").password("password").roles("USER").and()   //启用内存用户存储添加新用户以及为给定用户授予角色权限
            .withUser("admin").password("password").roles("USER","ADMIN");　　//and() 方法能够将多个用户的配置连接
            //注：roles()方法是authorities()方法的简写形式,所给定的值都会添加一个“ROLE_”前缀,并将其作为权限授予给用户
    }
}
```
* 配置用户详细信息的方法包括:
<br/>　－and():用来连接配置
<br/>　－accountExpired(boolean):定义账号是否已经过期
<br/>　－accountLocked(boolean):定义账号是否已经锁定
<br/>　－authorities(GrantedAuthority...):授予某个用户一项或多项权限
<br/>　－authorities(List<? extends GrantedAuthority>):授予某个用户一项或多项权限
<br/>　－authorities(String...):授予某个用户一项或多项权限
<br/>　－credentialsExpired(boolean):定义凭证是否已经过期
<br/>　－disabled(boolean):定义账号是否已被禁用
<br/>　－password(String):定义用户的密码
<br/>　－roles(String...):授予某个用户一项或多项角色


### 基于数据库表进行认证:
* 使用jdbcAuthentication()方法配置Spring Security使用以JDBC为支撑的用户存储:
```
@Autowired
DataSource datasource;   //必须配置一个datasource以访问数据库，这里通过自动自动注入方法得到

@Override
protected void configure(AuthenticationManagerBuilder auth) throws Exception{
    auth
       .jdbcAuthentication()
          .dataSource(dataSource);
}
```
* 若定义的数据库模式与默认配置不匹配，需要配置自己的查询：
```
@Override
protected void configure(AuthenticationManagerBuilder auth) throws Exception {
    auth
        .jdbcAuthentication()
        .datesource(datasource)
        .usersByUsernameQuery("SELECT_By_USERNAME_SQL")   //配置用户认证信息SQL查询
        .authoritiesByUsernameQuery("SELECT_By_USERNAME_SQL");　　　//配置用户权限信息SQL查询
        .passwordEncoder(new StandardPasswordEncoder(""))　　//借助该方法指定一个密码转码器(encoder),从而使用转码后的密码
}
```
* passwordEncoder()方法可以接受Spring Security中PasswordEncoder接口的任意实现。Spring Security的加密模块包括了三个这样的实现: BCryptPasswordEncoder、NoOpPasswordEncoder和StandardPasswordEncoder,此外还可以提供自定义的实现
```
//PasswordEncode接口定义
public interface PasswordEncode{
    String encode(CharSequence rawPassword);
    boolean match(CharSquence rawPassword,CharSequence encodePassword);
}
```

### 基于LDAP进行认证
* 可以使用ldapAuthentication()方法(在功能上类似于jdbcAuthentication()),让Spring Security使用基于LDAP的认证:
```
@Override
protected void configure(AuthenticationManagerBuilder auth) throws Exception{
    auth
        .ldapAuthentication()
        .userSearchBase("ou")  //指定查询基础,在名为people的组织单元下搜索,默认搜索会在LDAP层级结构的根开始
        .userSearchFilter("(uid={0})")
        .groupSearchBase("ou=groups")  //在名为 groups 的组织单元下搜索
        .groupSearchFilter("member={0}") //为基础LDAP查询提供过滤条件
        .passwordCompare();  //设置通过密码比对进行认证
        .passwordEncoder(new Md5PasswordEncoder())  //指定密码转码器即加密策略
        .passwordAttribute("password") //声明密码属性的名称
        .contextSource().url("ldap://remote_url:port/dc=,dc="); //配置远程ldap服务器地址，默认监听本机的33389端口
        .contextSource().root()  //或指定嵌入式服务器的根前缀，当LDAP服务器启动时,会尝试在类路径下寻找LDIF文件来加载数据
                  //LDIF(LDAP Data Interchange Format,LDAP数据交换格式)是以文本文件展现LDAP数据的标准方式,每条记录可以有一行或多行,每项包含一个名值对,记录之间通过空行进行分割
                       .ldif("classpath:users.ldif")   //明确指定加载的LDIF文件。默认从整个根路径下搜索
}
```
> 注：基于LDAP进行认证的默认策略是进行绑定操作,直接通过 LDAP 服务器认证用户。另一种可选的方式是进行比对操作，涉及将输入的密码发送到LDAP目录上,并要求服务器将这个密码和用户的密码进行比对，比对在LDAP服务器内完成的,实际的密码能保持私密。默认情况下,在登录表单中提供的密码将会与用户的LDAP条目中的userPassword属性进行比对，如果密码被保存在不同的属性中,可以通
过 passwordAttribute()方法来声明密码属性的名称；

### 配置自定义的用户服务
如果现有的用户存储不能满足应用需求,需要创建并配置自定义的用户详细信息服务，需要提供一个自定义的UserDetailsService 接口实现，实现其实现loadUserByUsername()方法
```
//UserDetailsService 接口定义
public interface UserDetailsService{
    UserDetails loadUserByUsername(String username) throws UsernameNotFoundException;　　//返回代表给定用户的UserDetails对象
}
```
## 拦截请求
* 有时对于不同的请求可能有不同的认证和权限需求，对每个请求进行细粒度安全性控制的关键在于重载configure(HttpSecurity)方法：
```
//如对不同URL路径有选择的应用安全性
＠Override
protected void configure(HttpSecurity http) throws Exception{
    http
        .authorizeRequests()　　//配置HTTP安全性
        .antMatchers("/URL").authenticated()  //指定路径请求认证,路径支持Ant风格的通配符，支持一次指定多个
        .antMatchers(HttpMethod.POST/GET,"/URL").authenticated()  //指定路径和请求方法认证，此外而regexMatchers()方法则能够接受正则表达式来定义请求路径
        .anyRequest().permitAll();　　//说明其他所有的请求不需要认证和任何的权限
```
* authenticated() 要求在执行该请求时,必须已经登录了应用，否则Spring Security的Filter将会捕获该请求,并将用户重定向到应用的登录页面。同时,permitAll()方法允许请求没有任何的安全限制。此外，还有其他的一些方法能够用来定义该如何保护请求：
<br/> -access(String):如果给定的SpEL表达式计算结果为true,就允许访问
<br/> -anonymous():允许匿名用户访问
<br/> -authenticated():允许认证过的用户访问
<br/> -denyAll():无条件拒绝所有访问
<br/> -fullyAuthenticated():如果用户是完整认证的话(不是通过Remember-me功能认证的),就允许访问
<br/> -hasAnyAuthority(String...):如果用户具备给定权限中的某一个的话,就允许访问
<br/> -hasAnyRole(String...):如果用户具备给定角色中的某一个的话,就允许访问
<br/> -hasAuthority(String):如果用户具备给定权限的话,就允许访问
<br/> -hasIpAddress(String):如果请求来自给定IP地址的话,就允许访问
<br/> -hasRole(String):如果用户具备给定角色的话,就允许访问
<br/> -not():对其他访问方法的结果求反
<br/> -permitAll():无条件允许访问
<br/> -rememberMe():如果用户是通过Remember-me功能认证的,就允许访问
* 此外可以要求用户不仅需要认证,还要具备 ROLE_SPITTER 权限
```
＠Override
protected void configure(HttpSecurity http) throws Exception{
    http
        .authorizeRequests()
        .antMatchers("/URL").hasAuthority("ROLE_USER"); //此外还可以使用hasRole()方法,会自动使用“ROLE_”前缀
}
```

注：将任意数量的antMatchers()、regexMatchers()和anyRequest()连接起来,可以满足Web应用安全规则的需要，这些规则会按照给定的顺序发挥作用，因此在配置时一般将最为具体的请求路径放在前面,而最不具体的路径(如anyRequest())放在最后面

### 使用Spring表达式(Spring Expression Language,SpEL)进行安全保护
* 借助access()方法,可以将SpEL作为声明访问限制的一种方式,不再局限于基于用户的权限进行访问限制,如：.antMatchers("/url").access("SpEL"),此外Spring Security 支持其他SpEL表达式：
<br/>　－authentication：用户的认证对象
<br/>　－denyAll：结果始终为false
<br/>　－hasAnyRole(list of roles)：如果用户被授予了列表中任意的指定角色,结果为true
<br/>　－hasRole(role)：如果用户被授予了指定的角色,结果为true
<br/>　－hasIpAddress(IP Address)：如果请求来自指定IP的话,结果为true
<br/>　－isAnonymous()：如果当前用户为匿名用户,结果为true
<br/>　－isAuthenticated()：如果当前用户进行了认证的话,结果为true
<br/>　－isFullyAuthenticated()：如果当前用户进行了完整认证的话(不是通过Remember-me 功能进行的认证),结果为true
<br/>　－isRememberMe()：如果当前用户是通过Remember-me自动认证的,结果为true
<br/>　－permitAll：结果始终为 true
<br/>　－principal：用户的 principal 对象

### 强制通道的安全性
借助HttpSecurity对象的requiresChannel()能够为各种URL模式声明所要求的通道：
```
@Override
protected void configure(HttpSecurity http) throws Exception{
    http
        .authorizeRequests()
        .......
        .and()
        .requiresChannel()
        .antMatchers("/URL").requiresSecure(); //指定的路径需要https,可以使用requiresInsecure()指定路径始终通过HTTP传送
}
```

### 防止跨站请求伪造
* 如果一个站点欺骗用户提交请求到其他服务器的话,就会发生CSRF(cross-site request forgery,跨站请求伪造)攻击,
* 从Spring Security 3.2开始,默认会启用CSRF防护,通过一个同步 token 的方式,拦截状态变化的请求(例如,非GET、HEAD、OPTIONS和TRACE的请求)并检查CSRF token
* 如果请求中不包含 CSRF token 的话,或者token不能与服务器端的token相匹配,请求将会失败,并抛出CsrfException异常。
> (意味着表单必须在一个“_csrf”域中提交token,且这个token必须要与服务器端计算并存储的 token 一致)

Spring Security简化了将token放到请求的属性中的操作,若使用Thymeleaf作为页面模板，只需<form>标签的action属性添加了Thymeleaf命名空间前缀，就会自动生成一个“_csrf”隐藏域,如果使用 Spring的表单绑定标签, <sf:form>标签会自动为我们添加隐藏的CSRF token标签,此外可以在配置中通过调用csrf().disable()禁用Spring Security的CSRF防护功能
```
<!--Thymeleaf请求的属性中添加token  -->
<form method="POST" thymeleaf:action="">
<!-- JSP请求的属性中添加token -->
<input type="hiden" name="${_csrf.variable}" value="${_csrf.token}" />
//禁用 Spring Security 的 CSRF 防护功能
@Override
protected void configure(HttpSecurity http) throws Exception{
    http
        ....
        .csrf()
        .disable();
}
```
## 认证用户/退出
* 利用formLogin()方法启用了基本的登录页功能,可以添加自定义的登录页
* HTTPBasic认证( HTTP Basic Authentication )会直接通过HTTP请求本身,对要访问应用程序的用户进行认证，在HttpSecurity对象上调用httpBasic()启用HTTP Basic认证,还可以通过调用realmName()方法指定域
* 退出功能是通过Servlet容器中的Filter实现的(默认情况下),这个Filter会拦截针对 “/logout” 的请求,可以在configure()中配置退出后的重定向页面；
* 可以通过调用logoutUrl()方法重写默认的 LogoutFilter 拦截路径
```
@Override
protected void configure(HttpSecurity http) throws Exception{
    http
        .formLogin()　　//启用了基本的登录页功能
         .loginPage("/login_url_path")  //添加自定义的登录页
         .and()
         .logout()    //设置退出后的重定向页面
         .logoutSuccessUrl("/redirect_url")  
         .logoutUrl("/filter_url")  //重写默认的LogoutFilter拦截路径
         .and()
         . httpBasic()
         .realname("")
         .....
         .remeberMe()    //启用Rememberme功能,只需在HttpSecurity对象上调用rememberMe()即可  
           .tokenValiditySeconds(time) //配置有效时间，默认通过在cookie中存储一个token完成配置
           .key(keyname)  //设置私钥的名
}
```

## 保护视图
### 使用 Spring Security 的 JSP 标签库: 
* 命名空间声明：<%@ taglib uri="http://www.springframework.org/security/tags" prefix="security" %>
<br/>　－<security:accesscontrollist>：如果用户通过访问控制列表授予了指定的权限,那么渲染该标签体中的内容
<br/>　－<security:authentication>：渲染当前用户认证对象的详细信息
<br/>　－<security:authorize>：如果用户被授予了特定的权限或者SpEL表达式的计算结果为true,那么渲染该标签体中的内容

* 访问认证信息的细节:<security:authentication property="object.property" var="variable_name" />  并将其赋值给var指定的变量，默认是定义在页面作用域，以通过scope属性设置
* property 用来标示用户认证对象的一个属性,取决于用户认证的方式,此外可依赖通用认证属性,如：
<br/>　－authorities：一组用于表示用户所授予权限的 GrantedAuthority 对象
<br/>　－Credentials：用于核实用户的凭证(通常,这会是用户的密码)
<br/>　－details：认证的附加信息( IP 地址、证件序列号、会话 ID 等)
<br/>　－principal：用户的基本信息对象
* <security:authorize>JSP 标签能够根据用户被授予的权限有条件地渲染页面的部分内容：<security:authorize access="SpEL">....</security:authoriz>SpEL 表达式的值将确定标签主体内的内容是否渲染,此外url属性不用明确声明安全性限制，对一个给定的URL模式会间接引用其安全性约束


### 使用Thymeleaf的Spring Security方言


* Thymeleaf的安全方言提供了条件化渲染和显示认证细节的能力:
<br/>　－sec:authentication:渲染认证对象的属性。类似于Spring Security的<sec:authentication/>JSP标签
<br/>　－sec:authorize:基于表达式的计算结果,条件性的渲染内容。类似于Spring Security的<sec:authorize/>JSP标签
<br/>　－sec:authorize-acl:基于表达式的计算结果,条件性的渲染内容。类似于Spring Security的<sec:accesscontrollist/>JSP标签
<br/>　－sec:authorize-expr:sec:authorize属性的别名
<br/>　－sec:authorize-url:基于给定URL路径相关的安全规则,条件性的渲染内容。类似于Spring Security的<sec:authorize/>JSP标签使用url属性时的场景
> （使用安全方言，需确保Thymeleaf Extras Spring Security位于应用的类路径下，还需要在配置中使用SpringTemplateEngine来注册SpringSecurity Dialect）
```
@Bean
pubic SpringTemplateEngine templateEngine(TemplateResolver templateResolver){
    SpringTemplateEngine templateEngine=new SpringTemplateEngine();
    templateEngine.setTemplateResolver(templateResolver);
    templateEngine.addDialect(new SpringSecurityDialect()); //注册安全方言
    return templateEngine;
}
//在thymeleaf模板中声明安全命名空间：
<html 
....
xmlns:sec="http://www.thymeleaf.org/thymeleaf-extras-springsecurity3">

```
* sec:authorize属性会接受一个 SpEL 表达式。如果表达式的计算结果为 true ,那么元素的主体内容就会渲染:
```
<div sec:authorize="SpEL">
...
</div>
```
