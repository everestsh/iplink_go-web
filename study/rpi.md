#RPI 学习

**安装**

```
vi install.sh
git clone git://git.drogon.net/wiringPi
cd wiringPi/
git pull origin
./build
gpio -v
gpio readall

cd ..
git clone https://github.com/sunfounder/Sunfounder_SuperKit_C_code_for_RaspberryPi.git
git clone https://github.com/sunfounder/Sunfounder_SuperKit_Python_code_for_RaspberryPi.git
git clone https://github.com/sunfounder/Sunfounder_SuperKit_Python_code_for_RaspberryPi
#git clone https://github.com/sunfounder/SunFounder_SensorKit_for_RPi2.git

sudo apt-get install vim -y
```

**LED**

```
cd /home/pi/Sunfounder_SuperKit_C_code_for_RaspberryPi/01_LED/

```

```
gcc led.c -o led -lwiringPi

~/SunFounder_SensorKit_for_RPi2/C/04_relay_module $ gcc relay.c  -lwiringPi -lpthread
gcc dule_color_led.c   -lwiringPi -lpthread
```

