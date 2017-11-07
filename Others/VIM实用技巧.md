---
title: VIM实用技巧
tags: [编辑器, VIM, 编程技巧, 提高效率]
---

# VIM 实用技巧

### 常用命令
:h vimtutor ，打开vim向导
:h key-notation ,自定义按键映射项
:h feature-list, 查看功能列表
> 备注：vim的功能集分为small,normal,big,huge,在编译期可选，如，参数--with-feature=tiny只会编译最基本的功能
:h gui, 查看vim GUI版相关信息 

vim -u NONE -N :不加载vimrc(-u指定加载文件路径，默认会进入vi兼容模式，导致很多功能被禁用，-u指定加载文件路径),同时使能“nocompatible”选项，防止进入vim兼容模式
激活vim内置插件最小配置：
```
set nocompatible
filetype plugin on
```

“/”， 用命令行模式执行正向查找；
“?”, 用命令行模式执行反向查找
"="， 用命令行模式，对一个vim脚本表达式进行求值