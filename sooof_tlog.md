# sooof日志


##2017年05月07日 星期日

```
django-admin startproject sblog  建立sblog项目
cd sblog
python manage.py startapp blog 建立blog子模块
python manage.py runserver 运行
python manage.py migrate 同步数据库
```

##2017年05月10日 星期三

```
git add --all
git commit -m "beeblog00.0403"
git push origin master
```
##2017年05月21日 星期六 US

**twilio**

1.install Flask

```
sudo pip install Flask
sudo pip install twilio
sudo pip install virtualenv

```
2.

```
cd Workspace/twilio/
$ virtualenv --no-site-packages .
	New python executable in /Users/raylea/Workspace/twilio/1/	bin/python
	Installing setuptools, pip, wheel...
	done.
$source bin/activate
$ vim requirements.txt
	Flask>=0.12
	twilio~=6.0.0
$bin/pip  install -r requirements.txt

```

3.

run.py

```
from flask import Flask
app = Flask(__name__)

@app.route("/")
def hello():
    return "Hello World!"

if __name__ == "__main__":
    app.run(debug=True)
```



**Getting Started on Heroku with Python**

```
sudo pip install virtualenv

```


```
pip install Flask
sudo pip install twilio

```
4.running

```
$python run.py
* Running on http://127.0.0.1:5000/
```


**2017年05月23日 星期一 US**



```
(twilio) raylea:twilio raylea$ cd heroku/
(twilio) raylea:heroku raylea$ heroku login
Enter your Heroku credentials:
Email: sooof@me.com
Password: *********
Logged in as sooof@me.com


git clone https://github.com/heroku/python-getting-started.git
cd python-getting-started/

(twilio) raylea:python-getting-started raylea$ heroku create
Creating app... done, ⬢ glacial-ridge-63594
https://glacial-ridge-63594.herokuapp.com/ | https://git.heroku.com/glacial-ridge-63594.git

(twilio) raylea:python-getting-started raylea$ git push heroku master
Counting objects: 287, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (130/130), done.
Writing objects: 100% (287/287), 44.49 KiB | 0 bytes/s, done.
Total 287 (delta 138), reused 287 (delta 138)
remote: Compressing source files... done.
remote: Building source:
remote:
remote: -----> Python app detected
remote: -----> Installing python-3.6.0
remote: -----> Installing pip
remote: -----> Installing requirements with pip
remote:        Collecting dj-database-url==0.4.1 (from -r /tmp/build_4988ef384c8db88090d81c1c2e0d18e6/requirements.txt (line 1))
remote:          Downloading dj-database-url-0.4.1.tar.gz
remote:        Collecting Django==1.9.7 (from -r /tmp/build_4988ef384c8db88090d81c1c2e0d18e6/requirements.txt (line 2))
remote:          Downloading Django-1.9.7-py2.py3-none-any.whl (6.6MB)
remote:        Collecting gunicorn==19.6.0 (from -r /tmp/build_4988ef384c8db88090d81c1c2e0d18e6/requirements.txt (line 3))
remote:          Downloading gunicorn-19.6.0-py2.py3-none-any.whl (114kB)
remote:        Collecting psycopg2==2.6.2 (from -r /tmp/build_4988ef384c8db88090d81c1c2e0d18e6/requirements.txt (line 4))
remote:          Downloading psycopg2-2.6.2.tar.gz (376kB)
remote:        Collecting whitenoise==2.0.6 (from -r /tmp/build_4988ef384c8db88090d81c1c2e0d18e6/requirements.txt (line 5))
remote:          Downloading whitenoise-2.0.6-py2.py3-none-any.whl
remote:        Installing collected packages: dj-database-url, Django, gunicorn, psycopg2, whitenoise
remote:          Running setup.py install for dj-database-url: started
remote:            Running setup.py install for dj-database-url: finished with status 'done'
remote:          Running setup.py install for psycopg2: started
remote:            Running setup.py install for psycopg2: finished with status 'done'
remote:        Successfully installed Django-1.9.7 dj-database-url-0.4.1 gunicorn-19.6.0 psycopg2-2.6.2 whitenoise-2.0.6
remote:
remote: -----> $ python manage.py collectstatic --noinput
remote:        58 static files copied to '/tmp/build_4988ef384c8db88090d81c1c2e0d18e6/gettingstarted/staticfiles', 58 post-processed.
remote:
remote: -----> Discovering process types
remote:        Procfile declares types -> web
remote:
remote: -----> Compressing...
remote:        Done: 62.7M
remote: -----> Launching...
remote:        Released v4
remote:        https://glacial-ridge-63594.herokuapp.com/ deployed to Heroku
remote:
remote: Verifying deploy... done.
To https://git.heroku.com/glacial-ridge-63594.git
 * [new branch]      master -> master

(twilio) raylea:python-getting-started raylea$ heroku ps:scale web=1
Scaling dynos... done, now running web at 1:Free

heroku ps:scale web=1
heroku open

heroku logs --tail

heroku ps:scale web=0
heroku ps:scale web=1


virtualenv venv
source venv/bin/activate

faq1:

pip install -r requirements.txt



sudo  pip install dj-database-url

faq2:
sudo pip install whitenoise

python manage.py collectstatic

heroku local web -f Procfile.windows

faq3:  
echo gunicorn==0.16.1 > requirements.txt
venv/bin/pip install -r requirements.txt
source venv/bin/activate

(venv) raylea:python-getting-started raylea$ heroku local web -f Procfile
[OKAY] Loaded ENV .env File as KEY=VALUE Format
03:21:16 web.1   |  2017-05-23 03:21:16 [35267] [INFO] Starting gunicorn 0.16.1
03:21:16 web.1   |  2017-05-23 03:21:16 [35267] [INFO] Listening at: http://0.0.0.0:5000 (35267)
03:21:16 web.1   |  2017-05-23 03:21:16 [35267] [INFO] Using worker: sync
03:21:16 web.1   |  2017-05-23 03:21:16 [35268] [INFO] Booting worker with pid: 35268




Modify hello/views.py so that it imports the requests module at the start:
import requests
Now modify the index method to make use of the module. Try replacing the current index method with the following code:
def index(request):
    r = requests.get('http://httpbin.org/status/418')
    print(r.text)
    return HttpResponse('<pre>' + r.text + '</pre>')
Now test locally:
$ pip install -r requirements.txt
$ heroku local






http://localhost:5000/



git add .
git commit -m "Demo"
git push heroku master
heroku open





```


FAQ1:

```
sudo mkdir -p /etc/paths.d &&
echo /Applications/Postgres.app/Contents/Versions/latest/bin | sudo tee /etc/paths.d/postgresapp
```


faq2:

```
raylea:python-getting-started raylea$ venv/bin/pip install -r requirements.txt
Collecting dj-database-url==0.4.1 (from -r requirements.txt (line 1))
  Using cached dj-database-url-0.4.1.tar.gz
Collecting Django==1.9.7 (from -r requirements.txt (line 2))
  Using cached Django-1.9.7-py2.py3-none-any.whl
Collecting gunicorn==19.6.0 (from -r requirements.txt (line 3))
  Using cached gunicorn-19.6.0-py2.py3-none-any.whl
Collecting psycopg2==2.6.2 (from -r requirements.txt (line 4))
  Using cached psycopg2-2.6.2.tar.gz
Collecting whitenoise==2.0.6 (from -r requirements.txt (line 5))
  Downloading whitenoise-2.0.6-py2.py3-none-any.whl
Building wheels for collected packages: dj-database-url, psycopg2
  Running setup.py bdist_wheel for dj-database-url ... done
  Stored in directory: /Users/raylea/Library/Caches/pip/wheels/a6/ee/0b/fa5aa1269e9e877fc925294ecd7752e9265f42ee18d38c37dd
  Running setup.py bdist_wheel for psycopg2 ... done
  Stored in directory: /Users/raylea/Library/Caches/pip/wheels/49/47/2a/5c3f874990ce267228c2dfe7a0589f3b0651aa590e329ad382
Successfully built dj-database-url psycopg2
Installing collected packages: dj-database-url, Django, gunicorn, psycopg2, whitenoise
Successfully installed Django-1.9.7 dj-database-url-0.4.1 gunicorn-19.6.0 psycopg2-2.6.2 whitenoise-2.0.6
```

FAQ3:

```
raylea:python-getting-started raylea$ heroku local web -f Procfile
[OKAY] Loaded ENV .env File as KEY=VALUE Format
03:20:36 web.1   |  /bin/sh: gunicorn: command not found
[DONE] Killing all processes with signal  null
03:20:36 web.1   Exited with exit code 127
```
**2017年05月31日 星期二 US**

**python**

```
sudo apt-get install python3
sudo update-alternatives --install /usr/bin/python python /usr/bin/python2.7
sudo update-alternatives --install /usr/bin/python python /usr/bin/python3.4 1


ubuntu:
#update-alternatives --display python ##查看ubuntu安装了那些版本的python
#python --version ##查看python的当前版本
sudo update-alternatives --config python #配置python的当前版本
```

**find的使用:**

```
find ../ -name .DS_Store -exec rm {} \;
find ~/ -name '.DS_Store' -delete 删除当前目录的.DS_store
find . -type f -exec stat --format '%Y :%y %n' "{}" \; | grep -v cache | sort -nr | cut -d: -f2- | head
find . -type f -name "*.c" | xargs stat --format "%y %n" |sort -rn |head

搜索隐藏文件，并删除
find ./ -name "\.*"
find ./ -name "\.*" -exec rm {} \;
```
**搞定讨厌的.DS_store**

删除当前目录的.DS_store

```
find . -name '*.DS_Store' -type f -delete
```

删除所有目录的.DS_store

```
sudo find / -name “.DS_Store” -depth -exec rm {} \;
```

禁止.DS_store生成：

```
defaults write com.apple.desktopservices DSDontWriteNetworkStores -bool TRUE
```

恢复.DS_store生成：

```
defaults delete com.apple.desktopservices DSDontWriteNetworkStores
```

**2017年07月14日 星期五 CN**

**2017年07月15日 星期六 CN**
FAQ:
me@me-1T61p:~$ cat Desktop/1
Error mounting /dev/sdb1 at /media/me/G: Command-line `mount -t "exfat" -o "uhelper=udisks2,nodev,nosuid,uid=1000,gid=1000,iocharset=utf8,namecase=0,errors=remount-ro,umask=0077" "/dev/sdb1" "/media/me/G"' exited with non-zero exit status 32: mount: unknown filesystem type 'exfat'
me@me-1T61p:~$ sudo apt-get install exfat-fuse exfat-utils
[sudo] password for me:


**2017年12月19日 星期二 CN**
wget ftp://ftp.gnutls.org/gcrypt/gnutls/v3.1/gnutls-3.1.23.tar.xz
unxz gnutls-3.1.23.tar.xz && tar -xvf gnutls-3.1.23.tar
cd gnutls-3.1.23/
./configure
make
sudo make install
sudo ln -s /usr/local/lib/libgnutls.so.28 /usr/lib/libgnutls.so.28
gnutls-cli -v
