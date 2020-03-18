<!--
 * @Author: your name
 * @Date: 2019-12-08 21:26:46
 * @LastEditTime: 2019-12-08 21:28:01
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /leetcode/home/w_sorley/Documents/Learning-Notes/temp.md
 -->
### 1.tmux


### 二、使用 nvm 安装node
* 安装nvm
    - wget -qO- https://raw.githubusercontent.com/creationix/nvm/v0.33.6/install.sh | bash
* nvm 的使用,常用的nvm 指令有这几个
- nvm ls： 列出本地已经安装的node版本
- nvm ls-remote ： 列出所有的node版本
- nvm install --lts : 安装lts版本
- nvm install <version> ： 安装指定版本
- nvm use <version> ： 使用指定版本
 
* go get 使用ssh协议
 - git config --global url.git@github.com:.insteadOf https://github.com/
# 解决git下载慢错误
- git config --global http.lowSpeedLimit 0
- git config --global http.lowSpeedTime 999999
