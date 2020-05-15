# install and test django

**install django:**

Here is the code:

```
wget https://bootstrap.pypa.io/get-pip.py

chmod a+x get-pip.py
./get-pip.py
pip install Django 
or   sudo pip install Django==1.7.7
```
**test diango：**

```
raylea:Workspace raylea$ python
Python 2.7.10 (default, Feb  6 2017, 23:53:20) 
[GCC 4.2.1 Compatible Apple LLVM 8.0.0 (clang-800.0.34)] on darwin
Type "help", "copyright", "credits" or "license" for more information.
>>> import django
>>> print(django.get_version())
1.7.7
>>> django.VERSION
(1, 7, 7, 'final', 0)
>>> 
```