```
{
    "url": "IUbuntu1604_Install_Pybombs",
    "time": "2018/01/21 19:00",
    "tag": "SRD, UHD,gnuradio"
}
```

# Ubuntu1604 Install Pybombs

ubuntu-16.04.3-desktop-amd64.iso


```
sudo  apt-get update
sudo apt-get install git -y
sudo apt-get install python-pip  (只能安装出v8版本，必须更新，否则以后会出很多很奇怪的问题，)
sudo -H pip install --upgrade pip
sudo pip -V
sudo -H pip install PyBOMBS
sudo pip install git+https://github.com/gnuradio/pybombs.git

pybombs auto-config
pybombs recipes add-defaults
pybombs prefix init ~/prefix -a myprefix -R gnuradio-default

```
