<!--
 * @Description: Android Learning Notes
 * @Author: wangshouli
 * @Date: 2019-10-14 23:36:11
 * @LastEditTime: 2019-10-26 22:45:47
 * @LastEditors: Please set LastEditors
 -->
# Android学习笔记:

## 架构
* Linux内核
* 程序库
    - 安卓程序库: 专门为Android开发的基于Java的程序库，如:
	    - ndroid.app - 提供应用程序模型的访问，是所有 Android 应用程序的基石。
	    - android.content - 方便应用程序之间，应用程序组件之间的内容访问，发布，消息传递。
	    - android.database - 用于访问内容提供者发布的数据，包含 SQLite 数据库管理类。
	    - android.opengl - OpenGL ES 3D 图片渲染 API 的 Java 接口。
	    - android.os - 提供应用程序访问标注操作系统服务的能力，包括消息，系统服务和进程间通信。
	    - android.text - 在设备显示上渲染和操作文本。
	    - android.view - 应用程序用户界面的基础构建块。
	    - android.widget - 丰富的预置用户界面组件集合，包括按钮，标签，列表，布局管理，单选按钮等。
	    - android.webkit - 一系列类的集合，允许为应用程序提供内建的 Web 浏览能力。
    - 安卓运行时: 提供名为 Dalvik 虚拟机的关键组件
* 应用框架
    - 以Java类的形式为应用程序提供许多高级的服务，如：
        - 活动管理者 - 控制应用程序生命周期和活动栈的所有方面。
        - 内容提供者 - 允许应用程序之间发布和分享数据。
        - 资源管理器 - 提供对非代码嵌入资源的访问，如字符串，颜色设置和用户界面布局。
        - 通知管理器 - 允许应用程序显示对话框或者通知给用户。
        - 视图系统 - 一个可扩展的视图集合，用于创建应用程序用户界面。
* 应用程序
	- 基本构建块Android应用程序组件：由应用清单文件松耦合的组织,如:
	   - Activities 	描述UI，并且处理用户与机器屏幕的交互。
       - Services 	处理与应用程序关联的后台操作。
       - Broadcast Receivers 	处理Android操作系统和应用程序之间的通信。
       - Content Providers 	处理数据和数据库管理方面的问题。
## 配置文件
* MainActivity.java主程序入口
	- setContentView(R.layout.xxx)引用自res/layout目录下的xxx.xml文件
* manifest.xml文件中声明所有的组件,包含以下标签:
	- 声明应用：
	```
	<application andnroid:icon=“属性指出位于res/drawable-hdpi下面的应用程序图标">
       <activity android:name=".MainActivity">
	      <intent-filter>
             <action android:name="android.intent.action.MAIN" />
             <category android:name="android.intent.category.LAUNCHER"/>
          </intent-filter>
		</activity>
	</application>
	```
* @string指的是res/value文件夹下的strings.xml
* activity_main.xml是一个在res/layout目录下的layout文件。当应用程序构建它的界面时被引用。你将非常频繁的修改这个文件来改变应用程序的布局

## 资源(Resources)访问
* anim/ 	定义动画属性的XML文件。它们被保存在res/anim/文件夹下，通过R.anim类访问
* color/ 	定义颜色状态列表的XML文件。它们被保存在res/color/文件夹下，通过R.color类访问
* drawable/ 	图片文件，如.png,.jpg,.gif或者XML文件，被编译为位图、状态列表、形状、动画图片。它们被保存在res/drawable/文件夹下，通过R.drawable类访问
* layout/ 	定义用户界面布局的XML文件。它们被保存在res/layout/文件夹下，通过R.layout类访问
* menu/ 	定义应用程序菜单的XML文件，如选项菜单，上下文菜单，子菜单等。它们被保存在res/menu/文件夹下，通过R.menu类访问
* raw/ 	任意的文件以它们的原始形式保存。需要根据名为R.raw.filename的资源ID，通过调用Resource.openRawResource()来打开raw文件
* values/ 	包含简单值(如字符串，整数，颜色等)的XML文件。这里有一些文件夹下的资源命名规范。
	- arrays.xml代表数组资源，通过R.array类访问；
	- integers.xml代表整数资源，通过R.integer类访问；              
	- bools.xml代表布尔值资源，通过R.bool类访问；
	- colors.xml代表颜色资源，通过R.color类访问；
	- dimens.xml代表维度值，通过R.dimen类访问；
	- strings.xml代表字符串资源，- 通过R.string类访问；
	- styles.xml代表样式资源，通过R.style类访问
* xml/ 	可以通过调用Resources.getXML()来在运行时读取任意的XML文件。可以在这里保存运行时使用的各种配置文件
* 替代资源：运行时，Android 检测当前设备配置，并为应用程序加载合适的资源，在res/ 下创建一个新的目录，以 resourcename_configqualifier 的方式命名
* 使用R类(编译时生成)，通过子类+资源名或者直接使用资源 ID 来访问资源，可以通过代码还是通过 XML 文件

## 活动(Activity)

# 第一行代码-学习笔记
* build.gradle配置编译打包信息
	- jcenter依赖包括第三方安卓开源库
	- dependencies闭包，依赖类型包括:本地依赖、库依赖、远程依赖(如android-support支持库)
	- android:defaultConfig配置使用的SDK版本信息
	- android:buildType配置编译打包类型，包括relesae和debug
* logt + TAB 生成TAG常量
* logcat添加过滤器，调试更方便



## 活动Activity
* 包含用户界面，主要用于和用户交互
* 安卓强调逻辑和视图分离
* 新建活动:
     - 新建对应类，重写onCreate
     - 创建并使用关联的布局文件
     - 在Mainfest中注册
* 在活动中使用Toast:
	- Toast类主要用于弹出提醒信息(自动消失)
	- Toast.makeText(context,string, showtime)创建Toast对象，show方法弹出显示
* 在活动中使用Menu:
page:48


## 布局
* layout布局文件
* 元素类型:<Button><TextView>
* match_parent:和父元素一样
* wrap_content:只要能包含子元素即可
* @id/XXXXX引用对应资源，@+id/xxxx定义一个新的ID

## 其他
* Avtivity.findViewById(R.id.XXX)可以根据xml定义生成view组件
* View可设置各种事件监听器(Listener)
* Activity本身即为Context
