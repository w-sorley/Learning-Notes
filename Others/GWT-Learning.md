---
title: GWT学习记录
date: 2017-07-25 11:05 
---

# GWT学习记录
## 
* 创建工程：gwt_path/webAppCreator \-out \<project_name\> \[\-junit \<junit_jar_path\>\] \<base_package_name\>
* 以开发模式运行:1.进入project root path 2.执行：ant devmode(需要安装ant构建工具)

## GWT　ｍodule xml file:project_root/base_package/module_name.gwt.xml
```
<module rename-to="module_name">  /／模块命名
  <inherits name="com.google.gwt.user.User" />  //模块继承，可默认集成自gwt提供基础模块，同时也可以继承其他自定义模块(也可以在html页面指定）
  <inherits name='com.google.gwt.user.theme.standard.Standard'/>  //可以在此指定要应用从css样式，此处指定自带的standard.css，在编译时会自动在应用中绑定指定的静
  　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　　//资源态,称为自动资源包含Automatic Resource Inclusion，易于移植／共享
  <entry-point class="..." />  //指定入口类。如果一个module没有entry-point只能被其他模块继承。此外可能有多个entry-point按次序执行。
</module>  
```


* 可以使用html页面的div标签向页面中动态嵌入页面内容： RootPanel.get("div.id").add(...); //将RootPanel定义的内容绑定到id指定的页面div中
* 推荐通过css文件指定页面的样式

## GWT UI
* 面板Panel,可通过add()方法放置widgets
    * RootPanel,可放置其他Panel，可与html页面中的指定div绑定
    * VerticalPanel/HorizontalPanel
FlexTable
TextBox
Button
Label

## GWT页面样式
* 通常利用css指定页面样式，可以在模块文件(.gwt.xml)中指定，也可以在html页面中引入；
* 通常包括基础主题样式和自定义样式，也可以覆盖原有的自带基础样式，自定义自己的个性化页面样式，也可以继承已有的样式，在此基础上添加/修改
* 每个通过GWT自动生成的widgets对应html元素，都会有一个与之关联的样式名(class　属性)，可以通过其在css文件中自定义该元素的样式(Buttton 会生成\<button class="gwt-Button"\>)
* 可以通过widget的addStyleName()方法，自定义生成元素的class属性，用于在css文件中自定义样式
* 添加的class标签与原来的标签组合作为新的class标签:gwt-Button　gwt-Button-remove此时元素class含有两个值(Button.addStyleDependentName("remove");）
> 此外可以通过java直接设置对应的html标签属性，如:FlexTable.setCellPadding(6);

## 事件管理Manage Events
* GWT基于事件，可以响应用户通过鼠标，键盘或应用接口触发的事件
* 事件监听：
```
Widget.addEvent(new EventHandler{
    @Override 
    public void onEvent(Event event){
        //TODO 处理事件,可调用其他函数处理
    }
})
```
*　提供定时事件执行：com.google.gwt.user.client.Timer;
```
 Timer newTimer = new Timer() {
        @Override
        public void run() {
          //TODO 要定时器触发执行的任务，可指定执行其他函数
        }
      };
newTimer.scheduleRepeating()指定事件间隔执行
```
> 通过String.matches("regular_expression")  //通过正则表达式判断字符创是否匹配指定模式，若匹配返回true

>com.google.gwt.i18n.client.NumberFormat／DateTimeFormat：提高格式化字符
