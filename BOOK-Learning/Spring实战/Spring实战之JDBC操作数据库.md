---
title: Spring实战之JDBC操作数据库
date: 2017-06-30
tags: [JDBC，DataBase,Spring]
categories: Framework
---

# Spring实战之通过Spring和JDBC操作数据库

Spring自带一组数据访问框架,集成了多种数据访问技术，在通过JDBC或Hibernate等对象关系映射(object-relational mapping,ORM )框架实现数据持久化时，Spring能够消除持久化代码中单调枯燥的数据访问逻辑。
Spring 的数据访问
将底层的数据访问功能集成放在数据访问对象DAO(又称Repository)中，对外提供数据访问接口。从而实现服务与数据访问实现的解绑定，将数据访问的持久化技术隐藏在接口后，易于测试，便于以后的修改，更加灵活性

为了将数据访问层与应用程序的其他部分隔离开来,Spring采用的方式之一就是提供统一的异常体系
Spring 的数据访问异常体系



















