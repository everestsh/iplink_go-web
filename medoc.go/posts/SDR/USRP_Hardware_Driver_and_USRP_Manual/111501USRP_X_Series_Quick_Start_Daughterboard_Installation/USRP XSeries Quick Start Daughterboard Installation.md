```
{
    "url": "USRP_X_Series_Quick_Start_Daughterboard_Installation",
    "time": "2018/01/07 19:00",
    "tag": "SRD, UHD,USRP Hardware Driver and USRP Manual "
}
```

# USRP X 系列 快速 入门 (子板安装)

## Abstract

This application note is a detailed step-by-step guide to install a daughterboard into the USRP X300/X310.

## Overview

This Quick Start is meant to show you how to put together your new X300/310. We will start at the point when you have yet to unpack the boxes and go all the way to being able to ping the device, performing a quick software probe to verify hardware components and finally running a simple FFT demo. This Quick Start does not cover the installation of software on the host computer. If you have not installed UHD/Gnuradio on your system, please reference the Building and Installing the USRP Open-Source Toolchain (UHD and GNU Radio) on Linux, OS X and Windows Application Notes. You may also use the Live SDR Environment to perform the verification steps for your USRP. Detailed information on the Live SDR Environment is available at the Live SDR Environment Getting Started Guides page.

## Tools Required
* Philips Screwdriver
* 5/16” wrench

## Pre-installed Software
* UHD Latest
* GNU Radio

## Box Contents

#### USRP Box
* 1 x USRP X300/X310
* 1 x SFP Adapter for 1 GigE
* 1 x Power Supply and US Cord
* 1 x USB 2.0 JTAG Debug Cable
* 1 x Gigabit Ethernet Cable
* 4 x SMA-Bulkhead Cables
* 16 x Daughterboard Screws
* Daughterboard Boxes
* 2 x SBX Daughterboards
* Antenna Boxes
* One or more Antennas
![Xseries quickstart figure 1](../../../static/uploads/imgs/Xseries_quickstart_figure_1.png)
![Xseries quickstart figure 2](../../../static/uploads/imgs/Xseries_quickstart_figure_2.png)

## Proper Care and Handling

All Ettus Research products are individually tested before shipment. The USRP™ is guaranteed to be functional at the time it is received by the customer. Improper use or handling of the USRP™ can easily cause the device to become non-functional. Ettus Research recommends you perform the installation with no power to the USRP and using ESD equipment. Listed below are some examples of actions which can prevent damage to the unit:

* Never allow metal objects to touch the circuit board while powered.
* Always properly terminate the transmit port with an antenna or 50Ω load.
* Always handle the board with proper anti-static methods.
* Never allow the board to directly or indirectly come into contact with any voltage spikes.
* Never allow any water, or condensing moisture, to come into contact with the boards.
* Always use caution with FPGA, firmware, or software modifications.

**Caution:** Never apply more than -15 dBm of power into any RF input.
**Caution:** Always use at least 30dB attenuation if operating in loopback configuration

## Installation Process
1. Unscrew the 2 screws on the top of the USRP and remove cover. (Lift up about 15 degrees and wiggle back as there is a flange on the front part of the cover)
![Xseries quickstart figure 3](../../../static/uploads/imgs/Xseries_quickstart_figure_3.png)
![Xseries quickstart figure 4](../../../static/uploads/imgs/Xseries_quickstart_figure_4.png)
![Xseries quickstart figure 5](../../../static/uploads/imgs/Xseries_quickstart_figure_5.png)
2. Line up the 8 screw holes on the Daughterboard with the USRP Motherboard standoffs (they only go one way).
![Xseries quickstart figure 6](../../../static/uploads/imgs/Xseries_quickstart_figure_6.png)
![Xseries quickstart figure 7](../../../static/uploads/imgs/Xseries_quickstart_figure_7.png)
3. After you have aligned the Daughterboard correctly you can press the Daughterboard on to the connectors below them (you will feel them snap into place).
![Xseries quickstart figure 8](../../../static/uploads/imgs/Xseries_quickstart_figure_8.png)
![Xseries quickstart figure 9](../../../static/uploads/imgs/Xseries_quickstart_figure_9.png)
4. Put 8 of the screws provided in the daughterboard.
![Xseries quickstart figure 10](../../../static/uploads/imgs/Xseries_quickstart_figure_10.png)
5. Repeat steps 2- 4 for the second Daughterboard.
6. It is recommended to connect the bulkhead cables one at a time to avoid confusion. The Daughterboards and front panel of the X300/310 are clearly labeled as to which cable goes where.
![Xseries quickstart figure 11](../../../static/uploads/imgs/Xseries_quickstart_figure_11.png)
![Xseries quickstart figure 12](../../../static/uploads/imgs/Xseries_quickstart_figure_12.png)
![Xseries quickstart figure 13](../../../static/uploads/imgs/Xseries_quickstart_figure_13.png)
![Xseries quickstart figure 14](../../../static/uploads/imgs/Xseries_quickstart_figure_14.png)
![Xseries quickstart figure 15](../../../static/uploads/imgs/Xseries_quickstart_figure_15.png)
![Xseries quickstart figure 16](../../../static/uploads/imgs/Xseries_quickstart_figure_16.png)
7. Repeat step 6 for the other bulkhead cables.
![Xseries quickstart figure 17](../../../static/uploads/imgs/Xseries_quickstart_figure_17.png)
![Xseries quickstart figure 18](../../../static/uploads/imgs/Xseries_quickstart_figure_18.png)
8. Install USRP cover with screws.
![Xseries quickstart figure 19](../../../static/uploads/imgs/Xseries_quickstart_figure_19.png)
9. Connect the SFP 1 GigE adapter into USRP SFP port 0.
![Xseries quickstart figure 20](../../../static/uploads/imgs/Xseries_quickstart_figure_20.png)
![Xseries quickstart figure 21](../../../static/uploads/imgs/Xseries_quickstart_figure_21.png)
10. Connect the Gigabit Ethernet cable and power cord provided.
![Xseries quickstart figure 22](../../../static/uploads/imgs/Xseries_quickstart_figure_22.png)
11. Attach any Antennas you may have purchased.
![Xseries quickstart figure 23](../../../static/uploads/imgs/Xseries_quickstart_figure_23.png)
12. On the computer(host) you plan to use to connect to the USRP set the Ethernet adapter to have an IP address of 192.168.10.1 with a subnet mask of 255.255.255.0. Connect the other end of the Gigabit Ethernet cable to your computer.
![Xseries quickstart figure 24](../../../static/uploads/imgs/Xseries_quickstart_figure_24.png)
13. Power on the USRP (button on the front right of the USRP)
14. Ping the device from host computer:
```
   $ ping 192.168.10.2
```
![Xseries quickstart figure 25](../../../static/uploads/imgs/Xseries_quickstart_figure_25.png)
15. Assuming you have properly installed the UHD driver you can now run this command in a terminal/command window:
```
   $ uhd_usrp_probe
```
This will tell you about the hardware inside of your USRP. The output will look like the following:

```
   $ uhd_usrp_probe
   linux; GNU C++ version 4.8.4; Boost_105400; UHD_003.010.git-202-g9e0861e1

   -- X300 initialization sequence...
   -- Determining maximum frame size... 1472 bytes.
   -- Setup basic communication...
   -- Loading values from EEPROM...
   -- Setup RF frontend clocking...
   -- Radio 1x clock:200
   -- Detecting internal GPSDO.... No GPSDO found
   -- Initialize Radio0 control...
   -- Performing register loopback test... pass
   -- Initialize Radio1 control...
   -- Performing register loopback test... pass
     _____________________________________________________
    /
   |       Device: X-Series Device
   |     _____________________________________________________
   |    /
   |   |       Mboard: X300
   |   |   revision: 7
   |   |   revision_compat: 7
   |   |   product: 30518
   |   |   mac-addr0: ff:ff:ff:ff:ff:ff
   |   |   mac-addr1: ff:ff:ff:ff:ff:ff
   |   |   gateway: 255.255.255.255
   |   |   ip-addr0: 255.255.255.255
   |   |   subnet0: 255.255.255.255
   |   |   ip-addr1: 255.255.255.255
   |   |   subnet1: 255.255.255.255
   |   |   ip-addr2: 255.255.255.255
   |   |   subnet2: 255.255.255.255
   |   |   ip-addr3: 255.255.255.255
   |   |   subnet3: 255.255.255.255
   |   |   serial: FFFFFFF
   |   |   FW Version: 4.0
   |   |   FPGA Version: 20.0
   |   |   
   |   |   Time sources: internal, external, gpsdo
   |   |   Clock sources: internal, external, gpsdo
   |   |   Sensors: ref_locked
   |   |     _____________________________________________________
   |   |    /
   |   |   |       RX DSP: 0
   |   |   |   Freq range: -100.000 to 100.000 MHz
   |   |     _____________________________________________________
   |   |    /
   |   |   |       RX DSP: 1
   |   |   |   Freq range: -100.000 to 100.000 MHz
   |   |     _____________________________________________________
   |   |    /
   |   |   |       RX Dboard: A
   |   |   |   ID: SBX (0x0054)
   |   |   |   Serial: FFFFFF
   |   |   |     _____________________________________________________
   |   |   |    /
   |   |   |   |       RX Frontend: 0
   |   |   |   |   Name: SBXv3 RX
   |   |   |   |   Antennas: TX/RX, RX2, CAL
   |   |   |   |   Sensors: lo_locked
   |   |   |   |   Freq range: 400.000 to 4400.000 MHz
   |   |   |   |   Gain range PGA0: 0.0 to 31.5 step 0.5 dB
   |   |   |   |   Bandwidth range: 40000000.0 to 40000000.0 step 0.0 Hz
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
   |   |   |   ID: SBX (0x0054)
   |   |   |   Serial: FFFFFF
   |   |   |     _____________________________________________________
   |   |   |    /
   |   |   |   |       RX Frontend: 0
   |   |   |   |   Name: SBXv3 RX
   |   |   |   |   Antennas: TX/RX, RX2, CAL
   |   |   |   |   Sensors: lo_locked
   |   |   |   |   Freq range: 400.000 to 4400.000 MHz
   |   |   |   |   Gain range PGA0: 0.0 to 31.5 step 0.5 dB
   |   |   |   |   Bandwidth range: 40000000.0 to 40000000.0 step 0.0 Hz
   |   |   |   |   Connection Type: IQ
   |   |   |   |   Uses LO offset: No
   |   |   |     _____________________________________________________
   |   |   |    /
   |   |   |   |       RX Codec: B
   |   |   |   |   Name: ads62p48
   |   |   |   |   Gain range digital: 0.0 to 6.0 step 0.5 dB
   |   |     _____________________________________________________
   |   |    /
   |   |   |       TX DSP: 0
   |   |   |   Freq range: -100.000 to 100.000 MHz
   |   |     _____________________________________________________
   |   |    /
   |   |   |       TX DSP: 1
   |   |   |   Freq range: -100.000 to 100.000 MHz
   |   |     _____________________________________________________
   |   |    /
   |   |   |       TX Dboard: A
   |   |   |   ID: SBX (0x0055)
   |   |   |   Serial: FFFFFF
   |   |   |     _____________________________________________________
   |   |   |    /
   |   |   |   |       TX Frontend: 0
   |   |   |   |   Name: SBXv3 TX
   |   |   |   |   Antennas: TX/RX, CAL
   |   |   |   |   Sensors: lo_locked
   |   |   |   |   Freq range: 400.000 to 4400.000 MHz
   |   |   |   |   Gain range PGA0: 0.0 to 31.5 step 0.5 dB
   |   |   |   |   Bandwidth range: 40000000.0 to 40000000.0 step 0.0 Hz
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
   |   |   |   ID: SBX (0x0055)
   |   |   |   Serial: FFFFFF
   |   |   |     _____________________________________________________
   |   |   |    /
   |   |   |   |       TX Frontend: 0
   |   |   |   |   Name: SBXv3 TX
   |   |   |   |   Antennas: TX/RX, CAL
   |   |   |   |   Sensors: lo_locked
   |   |   |   |   Freq range: 400.000 to 4400.000 MHz
   |   |   |   |   Gain range PGA0: 0.0 to 31.5 step 0.5 dB
   |   |   |   |   Bandwidth range: 40000000.0 to 40000000.0 step 0.0 Hz
   |   |   |   |   Connection Type: QI
   |   |   |   |   Uses LO offset: No
   |   |   |     _____________________________________________________
   |   |   |    /
   |   |   |   |       TX Codec: B
   |   |   |   |   Name: ad9146
   |   |   |   |   Gain Elements: None
```

## UHD FFT
Try the UHD_FFT demo that comes with GNU Radio

1. Connect one antenna to RX2 on the left Daughterboard (Daughterboard A)
![Xseries quickstart figure 26](../../../static/uploads/imgs/Xseries_quickstart_figure_26.png)
2. From the terminal/command window:
```
   $ uhd_fft --ant RX2
```
![Xseries quickstart figure 27](../../../static/uploads/imgs/Xseries_quickstart_figure_27.png)
![Xseries quickstart figure 28](../../../static/uploads/imgs/Xseries_quickstart_figure_28.png)


Success
Congratulations! You have successfully setup and verified your new USRP X300/X310. A more detailed verification guide is at the Verifying the Operation of the USRP Using UHD and GNU Radio application note. For additional step-by-step guides to using your USRP X300/X310, see the Application Notes section of the [Ettus Research Knowledge Base](https://kb.ettus.com/Application_Notes)

[link](https://kb.ettus.com/USRP_X_Series_Quick_Start_(Daughterboard_Installation))

## links
   * [目录](<../100000TableOfContents/Table Of Contents.md>)
   * 上一章: [目录](<../100000TableOfContents/Table Of Contents.md>)
   * 下一节: [X系列硬件设备恢复](<111502X300&X310_Device_Recovery.md>)
