###
 # @Description: build the new golang project  
 # @aUTHOR:wang.shouli 
 # @Date: 2019-06-15 15:23:02
 # @LastEditTime: 2019-06-15 15:56:00
 # @LastEditors: Please set LastEditors
###

# set -e

# 获取工程名
if [ $# -lt 1 ] || [ "$1" = "" ]; then
    echo -n "please enter the new project name:"
    read name
else
    name="$1"
fi

if [ -n  $name ];then
    name="new_project"
fi

# 创建最上层工程目录，失败则退出
mkdir "$name" || { echo "build failed!!"; exit 1; }

# 创建子目录
cd "$name"
mkdir "src"
mkdir "bin"
mkdir "pkg"

echo "create successful!!";
exit 0;
