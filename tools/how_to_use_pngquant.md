# 使用pngquant命令近乎无损压缩PNG图片大小减少70%左右


## 安装
```
brew install libpng

git clone git://github.com/pornel/pngquant.git
cd pngquant
```
③ make 回车后就会看到一大串命令，然后看到生成pngquant执行文件以及其他相关文件，即安装成功。
```
wangnaiyideMacBook-Pro:pngquant One$ make
此处省略若干。。。
ar rv libimagequant.a pam.o mediancut.o blur.o mempool.o viter.o nearest.o libimagequant.o
ar: creating archive libimagequant.a
a - pam.o
a - mediancut.o
a - blur.o
a - mempool.o
a - viter.o
a - nearest.o
a - libimagequant.o
gcc pngquant.o rwpng.o rwpng_cocoa.o lib/libimagequant.a -O3 -fno-math-errno -funroll-loops -fomit-frame-pointer -Wall -std=c99 -I. -DNDEBUG -DUSE_SSE=1 -msse -Wno-unknown-pragmas -mmacosx-version-min=10.6 -DUSE_COCOA=1 -I/usr/local/Cellar/libpng/1.6.19/include/libpng16 -I/usr/include -mmacosx-version-min=10.6 -framework Cocoa -L/usr/local/Cellar/libpng/1.6.19/lib -lpng16 -L/usr/lib -lz -lm -o pngquant
```
## 使用示例
```
 ls *.png | xargs -L1 -t /Users/One/pngquant/pngquant --ext .png --force 256 --speed 1 --quality=50-60
 ```
 
## 详细了解见官网：https://pngquant.org/
