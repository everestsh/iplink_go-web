```
{
    "url": "usrp n2xx ADS-B",
    "time": "2018/01/03 19:00",
    "tag": "SRD, LDE,环境部署"
}
```

# usrp设备跟踪飞机轨迹

**update FW and FPGA bin file**

```
cd /microembedded/burn/
chmod 777 ./usrp_n210_image_burn.sh
./usrp_n210_image_burn.sh
```

**run ADS-B**

```
google-earth
cd /microembedded/examples/mode-s/gr-air-modes-master/apps/
./modes_rx -K out.kml
```
