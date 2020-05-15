#FAQ: python django

**FAQ:**

不能创建python manage.py makemigrations数据库表文件？

1、在setting.py文件中

```
INSTALLED_APPS = (
    'django.contrib.admin',
    'django.contrib.auth',
    'django.contrib.contenttypes',
    'django.contrib.sessions',
    'django.contrib.messages',
    'django.contrib.staticfiles',
    'blog',//添加app名
)
```
和：

```
pip install pillow
```

2、

```
$>python manage.py makemigrations 建立版本
manage.py@sblog > makemigrations 
bash -cl "/System/Library/Frameworks/Python.framework/Versions/2.7/bin/python2.7 /Applications/PyCharm.app/Contents/helpers/pycharm/django_manage.py makemigrations /Users/raylea/Workspace/python/PycharmProjects/sblog"
Migrations for 'blog':
  0001_initial.py:
    - Create model Article
    - Create model Author
    - Add field author to article

$>python manage.py migrate 同步数据库
manage.py@sblog > migrate
bash -cl "/System/Library/Frameworks/Python.framework/Versions/2.7/bin/python2.7 /Applications/PyCharm.app/Contents/helpers/pycharm/django_manage.py migrate /Users/raylea/Workspace/python/PycharmProjects/sblog"
Operations to perform:
  Apply all migrations: admin, blog, contenttypes, auth, sessions
Running migrations:
  Applying blog.0001_initial... OK

test:
sqlite3 db.sqlite3 
sqlite> .tables
auth_group                  blog_article              
auth_group_permissions      blog_author               
auth_permission             django_admin_log          
auth_user                   django_content_type       
auth_user_groups            django_migrations         
auth_user_user_permissions  django_session 


manage.py@sblog > syncdb
bash -cl "/System/Library/Frameworks/Python.framework/Versions/2.7/bin/python2.7 /Applications/PyCharm.app/Contents/helpers/pycharm/django_manage.py syncdb /Users/raylea/Workspace/python/PycharmProjects/sblog"
Operations to perform:
  Apply all migrations: admin, blog, contenttypes, auth, sessions
Running migrations:
  No migrations to apply.

You have installed Django's auth system, and don't have any superusers defined.
Would you like to create one now? (yes/no):  yes
Username (leave blank to use 'raylea'):  sooof
Email address:  sooof@me.com
Warning: Password input may be echoed.
Password:  QAZwsx123
Warning: Password input may be echoed.
Password (again):  QAZwsx123
Superuser created successfully.
===python manage.py createsuperuser
```

