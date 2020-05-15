```
{
    "url": "usrp U&O mean",
    "time": "2018/01/03 19:00",
    "tag": "SRD, LDE,环境部署"
}
```

# USRP使用过程中出现“U”、“O”的意义

## USRP简单介绍

USRP（Universal Software RadioPeripheral）通用软件无线电外设，是构成SDR（Software Defined Radio）系统的主要部件（关于SDR系统的原理介绍，可以参考【1】），其中通用的无线电外设主要有HackRF，bladeRF和USRP三种（关于三种外设的区别，可以参考【2】）。本文中以下所说的USRP，指的是Ettus公司的产品USRP，到目前为止，其已经有众多系列，例如B系列、E系列、N系列、X系列等，基本上字母越靠后，价格越高，目前在实验室主要使用的是USRP B系列的产品B200/B210，其采用USB3.0与GPP（General Purpose Processor，通用处理器）连接，并直接使用USB供电。

最近一直在跑OAI的代码来进行 LTE系统的仿真工作，其所支持的无线电外设就包括了USRP，在运行LTE系统的 eNB 的代码的时候，会打印“U”、“L”等字母，刚开始一直以为是程序中的printf输出，而且本身打印输出的也并不多，所以没有太在意，但近期在仿真运行LTE系统的 1.4MHz带宽的OAI eNB 的时候，出现了大量的“U”，而且UE侧也无法接入，所以才注意到其问题（OAI 目前1.4MHz无法跑通，比较稳定的是带宽5M以及10M），在这里整理一番，与大家共享。mar
![USRP U&O](../../../static/uploads/imgs/USRP_U_O.png)


## Figure 1 OAI eNB输出打印“U”

## USRP使用中出现“U”、“O”的意义

在Linux系统下，USRP作为硬件设备，是由UHD来进行驱动的，所以在使用之前，需要先安装UHD驱动（有时间写一下UHD的安装教程吧！）。USRP在使用过程中，会由于与PC机之间的一些“不协调”而输出打印一些字符，而明白这些字符的意义，有助于解决USRP与PC之间的连接问题。如下是经常会输出打印的字符的意义：

###### 1）’a’：audio，表示声卡；

###### 2）’u’：USRP；

###### 3）’O’：overrun，PC not keeping up with received data from USRP of audio card，表示“超速、溢出”，也就是说 PC无法同步地去接收USRP上的数据;

###### 4）’U’：underrun，PC not providing data quickly，与’O’相反，表示PC无法快速的提供数据；

###### 5）’L’：latency，a bunch of late packet / late transmit packet，表示PC与USRP之间的信号传输时延较大。

当然，也可能会出现 ‘aU’ 、‘aO’、’uU’、’uO’等组合，理解的话也就是以上意义的组合吧。

关于出现这些字符的解决方法，因为以上问题都是PC与USRP之间信号的速率不匹配的问题，所以提高PC的性能是一种解决的方法，另外一个的话，就是修改、优化程序，减少程序的复杂度，具体可以参考3) 4)中所提到的思路。
