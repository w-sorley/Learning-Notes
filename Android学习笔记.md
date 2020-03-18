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
