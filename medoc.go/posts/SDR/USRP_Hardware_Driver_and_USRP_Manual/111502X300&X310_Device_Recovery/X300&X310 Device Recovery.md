```
{
    "url": "X300&X310_Device_Recovery",
    "time": "2018/01/08 19:00",
    "tag": "SRD, UHD,USRP Hardware Driver and USRP Manual "
}
```


# X系列硬件设备恢复

## 摘要
本文介绍了通过JTAG恢复USRP X300 / X310的详细信息.

## 概述
本文介绍了通过JTAG接口刷写FPGA映像文件来恢复USRP X300 / X310的过程.

**Note:** Linux only.

## 手册
参考文档，用户X300/X310手册，链接为.

http://files.ettus.com/manual/page_usrp_x3x0.html

## 所需工具
具有USB2 / 3和1GbE或10GbE接口的计算机
Ubuntu 14.04或16.x安装
UHD安装
USB2电缆
SFP+/RJ45适配器
以太网电缆

## 预备知识
本文假设您在一个Ubuntu 14.x或16.x Linux安装。您还应该安装UHD.如果您没有安装UHD，请参考[在Linux上构建和安装USRP开源工具链（UHD和GNU Radio）](000000000Building_and_Installing_the_USRP_Open-Source_Toolchain_UHD_and_GNU_Radio_on_Linux).

在进行UHD安装之前，您将需要下载匹配的FPGA映像。如果您没有下载FPGA映像，请运行命令:
```
sudo uhd_images_downloader
```

验证您是否已经下载FPGA映像:
```
ls -alh /usr/local/share/uhd/images/usrp_x3*
```
![X3x0 recovery 1](../../../static/uploads/imgs/X3x0_recovery_1.png)

## 安装 Xilinx Vivado  2015.4
你需要下载并安装Xilinx Vivado 2015.4为了通过JTAG线缆刷写的USRP X300 / x310。Xilinx Vivado 2015.4版可以在以下链接下载:

[Xilinx Vivado 2015.4](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X3XX_FPGA/xilinx2015.4.tar.bz2)

下一步, 解压缩文件到指定位置:
```
sudo tar xvf xilinx2015.4.tar.bz2  -C /opt/
```
下一步, 进入新到目录并运行 xsetup 进行安装 (必须以 sudo 的权限进行安装):
```
cd /opt/xilinx2015.4/Downloads/2015.4/
sudo xsetup
```
![X3x0 recovery 2](../../../static/uploads/imgs/X3x0_recovery_2.png)
(上图略微在目录位置有点不正确)

如下将运行xilinx2015.4安装
![X3x0 recovery 3](../../../static/uploads/imgs/X3x0_recovery_3.png)
（上图略微在目录位置有点不正确）

你会被提示有更新的版本，忽略这个弹出窗口，点击继续安装Xilinx Vivado 2015.4。如果有连不上internet，点击忽略。
![X3x0 recovery 4](../../../static/uploads/imgs/X3x0_recovery_4.png)

然后将提示您接受各种许可协议，单击``Next``.
![X3x0 recovery 5](../../../static/uploads/imgs/X3x0_recovery_5.png)

接着选择安装类型如图所示，单击``Next``.
![X3x0 recovery 6](../../../static/uploads/imgs/X3x0_recovery_6.png)

再选择如图所示，单击``Next``.
![X3x0 recovery 7](../../../static/uploads/imgs/X3x0_recovery_7.png)

再下来选择如图所示，单击``Next``.
![X3x0 recovery 8](../../../static/uploads/imgs/X3x0_recovery_8.png)

 最后，开始安装，Click ``Install``.
![X3x0 recovery 9](../../../static/uploads/imgs/X3x0_recovery_9.png)

安装过程只需要几分钟。
![X3x0 recovery 10](../../../static/uploads/imgs/X3x0_recovery_10.png)

然后会提示安装成功。单击OK，安装程序将关闭。

![X3x0 recovery 11](../../../static/uploads/imgs/X3x0_recovery_11.png)

## 安装 Digilent Cable 驱动
为了使用内置的USRP X300/x310前面板的JTAG接口，您需要安装 Digilent Cable 驱动。它包括Xilinx Vivado包装。

进入该文件夹
``/opt/Xilinx/Vivado_Lab/2015.4/data/xicom/cable_drivers/lin64/install_script/install_drivers``,并运行程序脚本.
```
cd /opt/Xilinx/Vivado_Lab/2015.4/data/xicom/cable_drivers/lin64/install_script/install_drivers
sudo ./install_digilent.sh
```

![X3x0 recovery 12](../../../static/uploads/imgs/X3x0_recovery_12.png)

## 配置网络接口

您将需要将要连接到USRP X300/X310的以太网接口设置为静态IP地址192.168.10.1，并设置MTU为1500.
![X3x0 recovery 13](../../../static/uploads/imgs/X3x0_recovery_13.png)
![X3x0 recovery 14](../../../static/uploads/imgs/X3x0_recovery_14.png)

## 连接X300/X310

使用USB2电缆通过前面板上的JTAG端口连接USRP X300/X310。您还可以将SFP + / RJ45适配器连接到端口0，并通过以太网连接计算机。

启动USRP X300/X310.

## Starting Xilinx Vivado  Edition
进入home目录:
```
cd ~/
source  /opt/Xilinx/Vivado/2015.4/settings64.sh
source /opt/Xilinx/SDK/2015.4/settings64.sh
```
下一步, 运行 Xilinx Vivado
```
/opt/Xilinx/Vivado/2015.4/bin/vivado
```
![X3x0 recovery 15](../../../static/uploads/imgs/X3x0_recovery_15_e.png)
![X3x0 recovery 16](../../../static/uploads/imgs/X3x0_recovery_16_e.png)

打开``Open Hardware Manager``
![X3x0 recovery 17](../../../static/uploads/imgs/X3x0_recovery_17_e.png)
![X3x0 recovery 18](../../../static/uploads/imgs/X3x0_recovery_18_e.png)

下一步,  ``Hardware Manager``的菜单中选择``Tools``->``Auto Connect``.
![X3x0 recovery 19](../../../static/uploads/imgs/X3x0_recovery_19_e.png)
FPGA的细节会出现在``Hardware Manager``左侧的窗口
![X3x0 recovery 20](../../../static/uploads/imgs/X3x0_recovery_20_e.png)
右键点击FPGA列表,并且选择``Program Device``.
![X3x0 recovery 21](../../../static/uploads/imgs/X3x0_recovery_21_e.png)

这将弹出一个新窗口。单击文件选择按钮并导航到UHD FPGA映像的位置，并为您的设备选择正确的FPGA映像。(``/usr/local/share/uhd/images``)

**注意:** 选择使用以 ``.bit`` 扩展名的与(``_x300``或``_x310``)相匹配的FPGA映像.建议选择``_HG``FPGA映像，将将端口0初始化为1 GbE，将端口1初始化为10 GbE。使用双10 GbE操作的高级用户可以选择``_XG``图像，但是您需要调整双10GbE配置（如：IP地址192.168.40.1，MTU设置等）。

![X3x0 recovery 22](../../../static/uploads/imgs/X3x0_recovery_22_e.png)
![X3x0 recovery 23](../../../static/uploads/imgs/X3x0_recovery_23_e.png)
![X3x0 recovery 24](../../../static/uploads/imgs/X3x0_recovery_24_e.png)
![X3x0 recovery 25](../../../static/uploads/imgs/X3x0_recovery_25_e.png)
下一步, 点击 ``Program``.
![X3x0 recovery 26](../../../static/uploads/imgs/X3x0_recovery_26_e.png)

随着FPGA被编程，弹出一个进度条。
![X3x0 recovery 27](../../../static/uploads/imgs/X3x0_recovery_27_e.png)

编程完成后，关闭Vivado。
![X3x0 recovery 28](../../../static/uploads/imgs/X3x0_recovery_28_e.png)


---
以上为用图形界面下载bit文件。
以下是用命令行来下载FPGA的bit文件：
下载脚本[035_burn_jtag_x310.sh](https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X310%E5%88%B7%E5%86%99%E8%84%9A%E6%9C%AC/install%26burn_x31x_uhd_gnuradio/035_burn_jtag_x310.sh)
```
./035_burn_jtag_x310.sh
```
 脚本文件内容是：
 ```
sudo chown -hR $USER:$USER $HOME/.Xilinx/
source  /opt/Xilinx/Vivado/2015.4/settings64.sh
source /opt/Xilinx/SDK/2015.4/settings64.sh
cd ~/fpga/usrp3/top/x300/
source setupenv.sh
viv_jtag_program /usr/local/share/uhd/images/usrp_x310_fpga_HG.bit
 ```
![X3x0 recovery 28_1](../../../static/uploads/imgs/X3x0_recovery_28_1.png)

返回终端，尝试ping USRP X300/X310。

```
ping 192.168.10.2
```
![X3x0 recovery 29](../../../static/uploads/imgs/X3x0_recovery_29.png)

停止ping用``CTRL-C``.

此时，如果您能ping通USRP X300/X310，请尝试运行UHD实用程序``uhd_usrp_probe``
```
uhd_usrp_probe
```
``uhd_usrp_probe``的屏幕显示内容:

```
user@host:~$ uhd_usrp_probe
linux; GNU C++ version 5.4.0 20160609; Boost_105800; UHD_003.010.001.HEAD-0-gc705922a

-- X300 initialization sequence...
-- Determining maximum frame size... 1472 bytes.
-- Setup basic communication...
-- Loading values from EEPROM...
-- Setup RF frontend clocking...
-- Radio 1x clock:200
-- [DMA FIFO] Running BIST for FIFO 0... pass (Throughput: 1304.3MB/s)
-- [DMA FIFO] Running BIST for FIFO 1... pass (Throughput: 1300.5MB/s)
-- [RFNoC Radio] Performing register loopback test... pass
-- [RFNoC Radio] Performing register loopback test... pass
-- [RFNoC Radio] Performing register loopback test... pass
-- [RFNoC Radio] Performing register loopback test... pass
-- Performing timer loopback test... pass
-- Performing timer loopback test... pass
  _____________________________________________________
 /
|       Device: X-Series Device
|     _____________________________________________________
|    /
|   |       Mboard: X310
|   |   revision: 8
|   |   revision_compat: 7
|   |   product: 30818
|   |   mac-addr0: 00:00:00:00:00:00
|   |   mac-addr1: 00:00:00:00:00:00
|   |   gateway: 192.168.10.1
|   |   ip-addr0: 192.168.10.2
|   |   subnet0: 255.255.255.0
|   |   ip-addr1: 192.168.20.2
|   |   subnet1: 255.255.255.0
|   |   ip-addr2: 192.168.30.2
|   |   subnet2: 255.255.255.0
|   |   ip-addr3: 192.168.40.2
|   |   subnet3: 255.255.255.0
|   |   serial: xxxxxxxx
|   |   FW Version: 5.1
|   |   FPGA Version: 33.0
|   |   RFNoC capable: Yes
|   |   
|   |   Time sources:  internal, external, gpsdo
|   |   Clock sources: internal, external, gpsdo
|   |   Sensors: ref_locked
|   |     _____________________________________________________
|   |    /
|   |   |       RX Dboard: A
|   |   |   ID: UBX-160 v1 (0x007a)
|   |   |   Serial: xxxxxxxx
|   |   |     _____________________________________________________
|   |   |    /
|   |   |   |       RX Frontend: 0
|   |   |   |   Name: UBX RX
|   |   |   |   Antennas: TX/RX, RX2, CAL
|   |   |   |   Sensors: lo_locked
|   |   |   |   Freq range: 10.000 to 6000.000 MHz
|   |   |   |   Gain range PGA0: 0.0 to 31.5 step 0.5 dB
|   |   |   |   Bandwidth range: 160000000.0 to 160000000.0 step 0.0 Hz
|   |   |   |   Connection Type: IQ
|   |   |   |   Uses LO offset: No
|   |   |     _____________________________________________________
|   |   |    /
|   |   |   |       RX Codec: A
|   |   |   |   Name: ads62p48
|   |   |   |   Gain range digital: 0.0 to 6.0 step 0.5 dB
|   |     _____________________________________________________
|   |    /
|   |   |       RX Dboard: B
|   |   |     _____________________________________________________
|   |   |    /
|   |   |   |       RX Frontend: 0
|   |   |   |   Name: Unknown (0xffff) - 0
|   |   |   |   Antennas:
|   |   |   |   Sensors:
|   |   |   |   Freq range: 0.000 to 0.000 MHz
|   |   |   |   Gain Elements: None
|   |   |   |   Bandwidth range: 0.0 to 0.0 step 0.0 Hz
|   |   |   |   Connection Type: IQ
|   |   |   |   Uses LO offset: No
|   |   |     _____________________________________________________
|   |   |    /
|   |   |   |       RX Codec: B
|   |   |   |   Name: ads62p48
|   |   |   |   Gain range digital: 0.0 to 6.0 step 0.5 dB
|   |     _____________________________________________________
|   |    /
|   |   |       TX Dboard: A
|   |   |   ID: UBX-160 v1 (0x0079)
|   |   |   Serial: xxxxxxxx
|   |   |     _____________________________________________________
|   |   |    /
|   |   |   |       TX Frontend: 0
|   |   |   |   Name: UBX TX
|   |   |   |   Antennas: TX/RX, CAL
|   |   |   |   Sensors: lo_locked
|   |   |   |   Freq range: 10.000 to 6000.000 MHz
|   |   |   |   Gain range PGA0: 0.0 to 31.5 step 0.5 dB
|   |   |   |   Bandwidth range: 160000000.0 to 160000000.0 step 0.0 Hz
|   |   |   |   Connection Type: QI
|   |   |   |   Uses LO offset: No
|   |   |     _____________________________________________________
|   |   |    /
|   |   |   |       TX Codec: A
|   |   |   |   Name: ad9146
|   |   |   |   Gain Elements: None
|   |     _____________________________________________________
|   |    /
|   |   |       TX Dboard: B
|   |   |     _____________________________________________________
|   |   |    /
|   |   |   |       TX Frontend: 0
|   |   |   |   Name: Unknown (0xffff) - 0
|   |   |   |   Antennas:
|   |   |   |   Sensors:
|   |   |   |   Freq range: 0.000 to 0.000 MHz
|   |   |   |   Gain Elements: None
|   |   |   |   Bandwidth range: 0.0 to 0.0 step 0.0 Hz
|   |   |   |   Connection Type: IQ
|   |   |   |   Uses LO offset: No
|   |   |     _____________________________________________________
|   |   |    /
|   |   |   |       TX Codec: B
|   |   |   |   Name: ad9146
|   |   |   |   Gain Elements: None
|   |     _____________________________________________________
|   |    /
|   |   |       RFNoC blocks on this device:
|   |   |   
|   |   |   * DmaFIFO_0
|   |   |   * Radio_0
|   |   |   * Radio_1
|   |   |   * DDC_0
|   |   |   * DDC_1
|   |   |   * DUC_0
|   |   |   * DUC_1
```

如果运行``uhd_usrp_probe``成功，请继续使用UHD实用程序刷写FPGA映像 ``uhd_image_loader``.

**注意:** 通过JTAG刷写FPGA映像不会将FPGA映像写入EEPROM，因此必须运行``uhd_image_loader``将FPGA映像写入内部EEPROM。
```
uhd_image_loader --args "type=x300,addr=192.168.10.2,type=HG"
```
![X3x0 recovery 30](../../../static/uploads/imgs/X3x0_recovery_30.png)

当``uhd_image_loader``完成刷写的过程时，建议您重新启动USRP X300/X310。

关闭USRP X300/X310电源，卸下JTAG USB电缆，然后打开USRP X300/X310电源。

USRP X300/X310现已恢复。你应该可以``ping``，正常运行``uhd_usrp_probe``和任何其他的UHD实用程序/应用程序.




## links
   * [目录](<../100000TableOfContents/Table Of Contents.md>)
   * 上一章: [USRP X 系列 快速 入门 (子板安装)](<111501USRP_X_Series_Quick_Start_Daughterboard_Installation.md>)
   * 下一节: [](<>)
