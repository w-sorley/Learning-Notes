---
title: JavaScripts学习记录
date: 2017-06-01
tags: [语法,基础，JS,Javascripts]
categories: JavaScripts
---


# JavaScripts学习记录

## 概述：
* JavaScript代码可以直接嵌在网页的任何地方(通常放到<head>中),由<script>...</script>包含(或放到一个单独的.js文件，通过<script src="..."></script>引入)
> （可以在一个页面中引入多个JS文件，浏览器按照顺序依次执行）

## 数据类型
* Number数值类型:JavaScript不区分整数和浮点数，统一用Numbe表示:
<br/>　　－MeN：表示科学计数法
<br/>　　－NaN：表示无法计算结果
<br/>　　－Infinity：表示无限大

### 字符串:
* 用''或""括起来;反引号 ` ... ` 表示原格式字符串
* 模板字符串：可以用+号连接；｀string1${变量名}string2｀在字符串中嵌入变量，自动替换字符串中的变量值
* 操作字符串:
<br/>　　－string.length:获得字符串长度
<br/>　　－string[index]:用索引获得包含的字符
<br/>　　－indexOf():返回指定字符串第一个字符的索引
<br/>　　－substring():返回指定索引区间的子串
> (字符串是不可变的，对字符串的某个索引赋值不会有任何错误也没有任何效果)

### 数组：
* 一个数组中可以包括任意数据类型，通过＇[]＇或Array()函数创建，通过索引访问
* 常用操作:
<br/>　　－indexOf():返回指定元素的索引
<br/>　　－slice()：截取Array的部分元素，然后返回一个新的Array
<br/>　　－push()：向Array的末尾添加若干元素；
<br/>　　－pop()则把Array的最后一个元素删除掉
<br/>　　－unshift()：方法往Array的头部添加若干元素
<br/>　　－shift():把Array的第一个元素删掉
<br/>　　－sort()：对当前Array进行排序
<br/>　　－reverse()：把整个Array的元素反转
<br/>　　－splice()：从指定的索引开始删除若干元素，然后再从该位置添加若干元素
<br/>　　－concat():把当前的Array和另一个Array连接起来，并返回一个新的Array
<br/>　　－join():把当前Array的每个元素都用指定的字符串连接起来，然后返回连接后的字符串
> 注：直接给Array的length赋一个新的值会导致Array大小的变化，如果通过索引赋值时，索引超过了范围，同样会引起Array大小的变化    

### 对象:
* 一组由字符串类型的键和任意数据类型的值组成的无序集合，通过'对象变量.属性名',获取一个对象的属性(也可以用Object['property']来访问)
> (如果属性名包含特殊字符，就必须用''括起来)
* 可以自由地给一个对象添加或删除属性:Object['new_property']=new_value/delete Object['property']
* in操作符检测对象是否拥有某一属性：'property' in Object
> (此属性可能是xiaoming继承得到，hasOwnProperty()，判断一个属性是否是自身拥有）

### 其他
* Map:一组键值对的结构（键值可以是非字符串类型），具有极快的查找速度
* Set：一组不能重复的key
* null表示一个“空”的值，undefined表示值未定义
* 变量：用＇var　变量名＇声明一个变量，只能申明一次（同其他动态语言一样，变量本身类型不固定）
> (JS本身并不强制要求用var申明变量，且没有通过var申明就被使用的变量自动被申明为全局变量，作用在在同一个页面的不同的JavaScript文件．使用var申明的变量，作用范围被限制在该变量被申明的函数体内，为解决全局变量带来的混乱，在strict模式下('use strict')，强制通过var申明变量)


### 常用操作：
* 遍历:可以用循环for(in)/可以用iterable:for( of )(同时iterable类型内置forEach方法)

* 比较运算符:JavaScript允许对任意数据类型做比较:
<br/>　－'=='会自动转换数据类型再比较
<br/>　－'==='不会自动转换数据类型，如果数据类型不一致，返回false，如果一致，再比较
<br/>　－(NaN与所有其他值都不相等(包括自身)，可通过isNaN()函数判断)


