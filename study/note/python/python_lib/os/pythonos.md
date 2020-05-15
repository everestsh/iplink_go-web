# python基础之模块之os模块

os模块

os模块的作用：

　　os，语义为操作系统，所以肯定就是操作系统相关的功能了，可以处理文件和目录这些我们日常手动需要做的操作，就比如说：显示当前目录下所有文件/删除某个文件/获取文件大小……

　　另外，os模块不受平台限制，也就是说：当我们要在linux中显示当前命令时就要用到pwd命令，而Windows中cmd命令行下就要用到这个，额...我擦，我还真不知道，（甭管怎么着，肯定不是pwd），这时候我们使用python中os模块的os.path.abspath(name)功能，甭管是linux或者Windows都可以获取当前的绝对路径。

os模块的常用功能：

1  os.name      #显示当前使用的平台

```python
>>> os.name
'nt'                  #这表示Windows
>>> os.name
'posix'               #这表示Linux
```

2 os.getcwd()      #显示当前python脚本工作路径

```python
>>> os.getcwd()
'C:\\Users\\Capital-D\\PycharmProjects\\untitled'    #使用pycharm

>>> os.getcwd()
'/root'               #Linux平台在/root目录直接使用python3命令
```

3  os.listdir('dirname')        #返回指定目录下的所有文件和目录名

```python
#相对于os.getcwd路径下的文件
>>> os.listdir('./')
['.idea', 'test']

>>> os.listdir('./')
['.bash_logout', 'Python-3.4.4', '.mysql_history', '.tcshrc', 'Python-3.4.4.tar.xz', '.bash_profile', '.lesshst', 'install.log.syslog', '.cshrc', '04.sql', 'anaconda-ks.cfg', 'test', '.viminfo', 'phpMyAdmin-4.4.15-all-languages.tar.bz2', '1test', '.bashrc', 'binlog.sql', 'back.sql', 'install.log', 'binlog4.sql', '.bash_history', 'backup.sql', 'text.py', '.rnd', 'test1']
```
4  os.remove('filename')       #删除一个文件

```
[root@slyoyo ~]# touch hahaha
[root@slyoyo ~]# ls
04.sql           back.sql     binlog.sql   install.log.syslog                       Python-3.4.4.tar.xz  text.py
1test            backup.sql   hahaha       phpMyAdmin-4.4.15-all-languages.tar.bz2  test
anaconda-ks.cfg  binlog4.sql  install.log  Python-3.4.4                             test1
#hahaha（粉色字体）存在
[root@slyoyo ~]# python3
Python 3.4.4 (default, Apr  5 2016, 04:23:19)
[GCC 4.4.7 20120313 (Red Hat 4.4.7-4)] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> import os
>>> os.remove('hahaha')
>>> exit()
[root@slyoyo ~]# ls
04.sql  anaconda-ks.cfg  backup.sql   binlog.sql   install.log.syslog                       Python-3.4.4         test   text.py
1test   back.sql         binlog4.sql  install.log  phpMyAdmin-4.4.15-all-languages.tar.bz2  Python-3.4.4.tar.xz  test1
#hahaha已被删
```

5  os.makedirs('dirname/dirname')     #可生成多层递规目录

```
[root@slyoyo ~]# ls
04.sql  anaconda-ks.cfg  backup.sql   binlog.sql   install.log.syslog                       Python-3.4.4         test   text.py
1test   back.sql         binlog4.sql  install.log  phpMyAdmin-4.4.15-all-languages.tar.bz2  Python-3.4.4.tar.xz  test1
[root@slyoyo ~]# python3
Python 3.4.4 (default, Apr  5 2016, 04:23:19)
[GCC 4.4.7 20120313 (Red Hat 4.4.7-4)] on linux
Type "help", "copyright", "credits" or "license" for more information.
>>> import os
>>> os.makedirs('hahaha/linghuchong')
>>> exit()
[root@slyoyo ~]# ls
04.sql           back.sql     binlog.sql   install.log.syslog                       Python-3.4.4.tar.xz  text.py
1test            backup.sql   hahaha       phpMyAdmin-4.4.15-all-languages.tar.bz2  test
anaconda-ks.cfg  binlog4.sql  install.log  Python-3.4.4                             test1
[root@slyoyo ~]# ls hahaha/
linghuchong
[root@slyoyo ~]# ls hahaha/linghuchong/
[root@slyoyo ~]#
```

6  os.rmdir('dirname')     #删除单级目录

```

```








FAQ:
1

报错：
```
>>> import os

>>> os.listdir()
Traceback (most recent call last):
File "<stdin>", line 1, in <module>
TypeError: listdir() takes exactly 1 argument (0 given)
```

原因：

listdir()函数需要一个入参，但是只给了0个入参。

解决方案：

加一个入参

```
>>> os.listdir('/home/autotest')
['hello.py', 'email126pro']
```
