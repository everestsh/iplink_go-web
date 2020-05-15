# install&burn_x31x_uhd_gnuradio

1. 安装14.04
2. 更新
```
sudo apt-get update
```
3. 下载[xilinx2015.4.tar.bz2](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X3XX_FPGA/xilinx2015.4.tar.bz2)、 [gnuradio.tar.bz2](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/x3x0_src/gnuradio.tar.bz2)、 [uhd.tar.bz2](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/x3x0_src/uhd.tar.bz2)、 [fpga.tar](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/x3x0_src/fpga.tar)
```
 cp xilinx2015.4.tar.bz2  gnuradio.tar.bz2 uhd.tar.bz2 fpga.tar $HOME
```
4. 如果你是用微嵌的X310 live DVD, 需要运行第4、5步。
下载脚本[010_uninstall_gnuradio.sh](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X310%E5%88%B7%E5%86%99%E8%84%9A%E6%9C%AC/install%26burn_x31x_uhd_gnuradio/010_uninstall_gnuradio.sh)
```
./010_uninstall_gnuradio.sh
```
5. 如果你是用微嵌的X310 live DVD, 需要运行第4、5步。
下载脚本[015_uninstall_uhd.sh](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X310%E5%88%B7%E5%86%99%E8%84%9A%E6%9C%AC/install%26burn_x31x_uhd_gnuradio/015_uninstall_uhd.sh)
```
./015_uninstall_uhd.sh
```
6. 下载脚本[020_install_uhd.sh](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X310%E5%88%B7%E5%86%99%E8%84%9A%E6%9C%AC/install%26burn_x31x_uhd_gnuradio/020_install_uhd.sh)
```
./020_install_uhd.sh
```
8. 下载脚本[025_install_gnuradio.sh](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X310%E5%88%B7%E5%86%99%E8%84%9A%E6%9C%AC/install%26burn_x31x_uhd_gnuradio/025_install_gnuradio.sh)
```
./025_install_gnuradio.sh
```
8. 安装xilinx2015.4,
```
sudo uhd_images_downloader
sudo tar xvf xilinx2015.4.tar.bz2  -C /opt/
```
9. 运行以下命令安装：
```
cd /opt/xilinx2015.4/Downloads/2015.4/
sudo xsetup
```
[详细过程](https://github.com/sooof/me_doc/blob/master/SDR/USRP_Hardware_Driver_and_USRP_Manual/111502X300%26X310_Device_Recovery.md)。

10. 插好USB JTDAG线，检测工具是否安装。
下载脚本[030_test_usb_jtag.sh](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X310%E5%88%B7%E5%86%99%E8%84%9A%E6%9C%AC/install%26burn_x31x_uhd_gnuradio/030_test_usb_jtag.sh)
```
./030_test_usb_jtag.sh
```
11. 下载FPGA文件到主板（此过程下载的文件断电会丢失。）
下载脚本[035_burn_jtag_x310.sh](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X310%E5%88%B7%E5%86%99%E8%84%9A%E6%9C%AC/install%26burn_x31x_uhd_gnuradio/035_burn_jtag_x310.sh)
```
./035_burn_jtag_x310.sh
```
12. 设置固定ip,
13. 刷新映像（此过程下载的文件断电不会丢失。）
下载脚本[035_burn_jtag_x310.sh](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X310%E5%88%B7%E5%86%99%E8%84%9A%E6%9C%AC/install%26burn_x31x_uhd_gnuradio/040_burn_net_x310.sh)
```
./040_burn_net_x310.sh
```
