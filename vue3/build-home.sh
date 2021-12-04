echo 同步远程仓库
git pull
echo 编译主页
npm run build_home
echo 删除源文件
rm -r ../js ../css ../img ../fonts
echo 拷贝主页
cp -rf ../home/. ../
echo 删除编译文件
rm -r ../home
echo 完成