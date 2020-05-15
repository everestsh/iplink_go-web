```
{
    "url": "OpenAirInterface安装教程",
    "time": "2018/01/20 19:00",
    "tag": "SRD, LTE,OAI"
}
```

# OpenAirInterface安装教程

安装时可同时参考官网文档、本文档、OAI安裝與協定觀察、OpenAirInterface安裝流程與設定。因为本人搭建时发现每个文档或多或少都有点细节问题，或者就是自己手上的设备跟文档中的不太一致。
老版本安装可参考：Tutorial for Open Air Interface

====以下内容99%转自OAI 完全搭建手册，也有很多细节问题，请大家结合上述自行斟酌====

开源软件OAI搭建终于成功，主要参照官网操作手册，并不断解决出现的各种问题，现将安装配置过程和遇到的主要问题整理在此文档，以供参考学习。

## 0x01 代码解析
OAI开源代码分为两个部分：eNB模块openairinterface5g、核心网模块openair-cn，老版本代码的核心网部分包含EPC和HSS两个模块，新版本核心网包含MME、HSS、SPGW三个模块。推荐安装master branch，因为官网确保master是最稳定的版本，有其他开发兴趣的可以check develop branch。

## 0x02 系统环境

我们的核心网安装在ThinkStation 64位主机上，系统Ubuntu 14.04，内核版本4.7；
eNB安装使用64位华硕主板，硬件要求corei5,i7，主频3.0GHz以上，USB3.0（射频前端要求），系统Ubuntu14.04.3 LTS，内核版本3.19。建议使用华硕主板，因为相关CPU frequency的配置较为容易。

###### 官网链接：https://gitlab.eurecom.fr/oai/openairinterface5g/wikis/OpenAirSystemRequirements

## 0x03 安装步骤

## 0x31 安装内核

#### eNB：

###### 1) 安装Ubuntu 14.04.3系统

```
sudo apt-get update
sudo apt-get -y install vim
```

###### 2) 在BIOS里关掉所有功耗管理性能(sleep states, in particular C-states，CPU frequencyscaling )：

在`/etc/default/grub`中添加`intel_pstate=disable`，也可以添加`processor.max_cstate=1intel_idle.max_cstate=0 idle=poll（optional）`
执行
```
sudo update-grub
```

在`/etc/modprobe.d/blacklist.conf`最末添加`blacklistintel_powerclamp`，如果文件不存在，新建一个。
最后确认在BIOS 关闭hyperthreading,CPU frequency control, C-States, P-States 和其他功耗管理。

###### 安装 cpufrequtils:
```
sudo apt-get -y install cpufrequtils
```
编辑以下文件(如不存在新建一个):
```
sudo vim /etc/default/cpufrequtils
```
添加下面这行到文件中：
```
GOVERNOR=”performance”
```
保存退出

###### 关闭ondemand daemon，防止重启后被改写.
```
sudo update-rc.d ondemand disable
```

###### 3) 安装3.19lowlatency内核：

```
sudo apt-get -y install linux-image-3.19.0-61-lowlatency linux-headers-3.19.0-61-lowlatency
```

###### 4) Reboot

#### Corenetwork：

###### 1) 系统为Ubuntu14.04.5 LTS

###### 2) 安装4.7.x版本内核，可选择官网pre-compile版本：

```
sudo apt-get -y install subversion git
git clone https://gitlab.eurecom.fr/oai/linux-4.7.x.git
cd linux-4.7.x
sudo dpkg -i linux-headers-4.7.7-oaiepc_4.7.7-oaiepc-10.00.Custom_amd64.deb linux-image-4.7.7-oaiepc_4.7.7-oaiepc-10.00.Custom_amd64.deb
```

###### 3) Reboot，从4.7.7-oaiepc内核启动

##### 官网链接：https://gitlab.eurecom.fr/oai/openairinterface5g/wikis/OpenAirKernelMainSetup

###### 问题：
之前未重新安装4.7内核时，运行spgw会报错error inloading gtp kernel module；后从源码安装带有gtp module的4.7.x内核，登录不了图形界面，原因是xorg server不能成功运行，且spgw运行时报错相同。登录界面问题通过crtl+alt+F1进入命令行界面，安装gnome服务解决。Gtp module加载失败的问题未解决，怀疑系统和内核未很好的安装所致，故重新安装了oai的pre-compile的4.7.x内核，gtp问题解决。
我自己装的时候，系统启动时没有BootLoader信息，所以没办法修改grub引导内核的顺序，遂直接把其他内核删了，只留下了4.7.7-oaiepc内核。

## 0x32 下载代码

###### 1) 安装Git：

```
sudo apt-get update
sudo apt-get -y install subversion git
```

###### 2) 添加许可证：
```
sudo su
echo -n | openssl s_client -showcerts -connect gitlab.eurecom.fr:443 2>/dev/null | sed -ne '/-BEGIN CERTIFICATE-/,/-END CERTIFICATE-/p' >> /etc/ssl/certs/ca-certificates.crt
```
###### 3) 下载源码：


#### eNB：
```
git clone https://gitlab.eurecom.fr/oai/openairinterface5g.git
```

#### corenetwork：

```
git clone https://gitlab.eurecom.fr/oai/openair-cn.git
```

#### 官网链接：https://gitlab.eurecom.fr/oai/openairinterface5g/wikis/GetSources

## 0x33 安装mysql-server

理论上讲，按照0x34步骤执行./build_hss –i 时会自动安装MySQL-server和phpmyadmin数据库管理软件，但是在实际安装时并没有自动弹出安装mysql-server，导致phpmyadmin也不能成功安装，故需要在此之前在核心网机器上单独安装mysql-server。

#### Core network：
```
sudo apt-get install mysql-server
```
安装过程中会弹出如下对话框，设置mysql的root密码。本机设置为Linux。

如果安装过程出现任何问题，无法成功配置root密码，则还可在命令行进行配置：
###### 1) 打开my.cnf文件添加语句skip-grant-tables

```
cd /etc/mysql
vim my.cnf
service mysqld restart
```

###### 2) 用空密码登录，切换至mysql库

```
mysql –u root –p
mysql> use mysql
```

###### 3) 重设密码

```
mysql> update user set password=PASSWORD(“new_pass”) where user=’root’;
```

###### 4) 回到vim /etc/mysql/my.cnf，把刚才加入的那一行“skip-grant-tables”注释或删除掉

###### 5) 重新启动使用新密码登录
```
sudo service mysql restart
mysql –u root -p
```

## 0x34 Autobuild安装缺少的包

#### eNB：

```
cd openairinterface5g #YOUR_openairinterface5g_DIRECTORY
source oaienv #Very important. It sets the correct environment variables.
cd cmake_targets
./build_oai -I --install-optional-packages #need to run only once
./build_oai –h #show help
```

#### Core network:

```
cd openair-cn #YOUR_openair-cn_DIRECTORY
cd SCRIPTS
./build_mme -i #Need to run only once to install missing packages
./build_hss -i #Need to run only once to install missing packages
./build_spgw -i #Need to run only once
```

在核心网安装此过程中会弹出对话框进行phpmyadmin的配置：

安装phpmyadmin数据库，选择ok
此处输入mysql的root密码linux：
输入phpmyadmin的密码，可随意设置。
Web server选择为apache：
##### 问题：

###### 1. 安装过程中报错error2002：can’t connect to local mysql server through socket。
检查后发现mysql-server并没有自动安装。为了避免后续导入数据库出错，重新安装了系统和内核，在执行此步骤前先手动安装mysql-server，问题解决。
###### 2. 在浏览器输入127.0.0.1/phpmyadmin,出现not found界面。
需要在Apache配置中添加phpmyadmin
打开apache.conf文件：

```
vim /etc/apache2/apache2.conf
```

添加下面语句：

```
include /etc/phpmyadmin/apache.conf
```

重新启动apache服务：

```
/etc/init.d/apache2 restart
```

##### 官网链接：https://gitlab.eurecom.fr/oai/openairinterface5g/wikis/AutoBuild

## 0x35 修改配置文件

#### eNB：

###### 1) 修改eNB配置文件，位置：`~/openairinterface5g/targets/PROJECTS/GENERIC-LTE-EPC/CONF/enb.band7.tm1.usrpb210.conf`

```
tracking_area_code  =  "1";
mobile_country_code =  "466";
mobile_network_code =  "68";

////////// MME parameters:

    mme_ip_address      = ( { ipv4       = "192.168.31.134";
                              ipv6       = "192:168:30::17";
                              active     = "yes";
                              preference = "ipv4";
                            }
                          );

    NETWORK_INTERFACES :
    {
        ENB_INTERFACE_NAME_FOR_S1_MME         = "eth0";
        ENB_IPV4_ADDRESS_FOR_S1_MME           = "192.168.31.210/24";
        ENB_INTERFACE_NAME_FOR_S1U            = "eth0";
        ENB_IPV4_ADDRESS_FOR_S1U              = "192.168.31.210/24";
        ENB_PORT_FOR_S1U                      = 2152; # Spec 2152
    };
```

此文件需要更改两处，第一是与sim卡信息对应的mmc、mnc值，第二是eNB本机的网卡与ip地址与核心网MME的ip地址。我们所使用的sim卡对应的mmc为466，mnc为68。eNB端网卡为eth0，ip address：192.168.31.210；MME的ip address为192.168.31.134.

###### 2) 修改openairinterface5g/openair3/UTILS/mcc-mnc-itu.c文件，添加{466.68}.

#### Corenetwork：

###### 1) 配置FQDN
域名需要为MME正确配置，我们可以通过hosts查看：

```
$ cat /etc/hostname
openair-cn
```

可按照如下输出修改etc/hosts文件：

```
$ cat /etc/hosts
127.0.0.1 localhost
127.0.1.1 openair-cn.openair4G.eur openair-cn
127.0.1.1 hss.openair4G.eur hss
```

修改host后要重启机器。

在以下信息完全符合时方可继续！

```
$ hostname
Openair-cn

$ hostname -f
Openair-cn.openair4G.eur
```

###### 2) 修改core network配置文件
将EPC的配置文件复制到/usr/local/etc/oai目录中：

```
sudo mkdir -p /usr/local/etc/oai/freeDiameter
sudo cp ~/openair-cn/ETC/mme.conf /usr/local/etc/oai
sudo cp ~/openair-cn/ETC/hss.conf /usr/local/etc/oai
sudo cp ~/openair-cn/ETC/spgw.conf /usr/local/etc/oai
sudo cp ~/openair-cn/ETC/acl.conf /usr/local/etc/oai/freeDiameter
sudo cp ~/openair-cn/ETC/mme_fd.conf /usr/local/etc/oai/freeDiameter
sudo cp ~/openair-cn/ETC/hss_fd.conf /usr/local/etc/oai/freeDiameter
```

此后只修改此目录下的配置文件。

###### 3) hss.conf

```
HSS :
{ ## MySQL mandatory options
MYSQL_server = "127.0.0.1";
MYSQL_user = "root";
MYSQL_pass = "linux"; #此处为配置mysql是设置的root密码
MYSQL_db = "oai_db";

## HSS options
OPERATOR_key = "11111111111111111111111111111111"; # OP key for oai_db.sql
RANDOM = "true";
## Freediameter options
FD_conf = "/usr/local/etc/oai/freeDiameter/hss_fd.conf";
};
```

此文件中注意配置好mysql的登录名和密码，OPERATOR_key是sim卡中确定的。本卡中为32位的1.

###### 4) hss_fd.conf

```
Identity = "hss.openair4G.eur";
Realm = "openair4G.eur";
```

###### 5) mme_fd.conf

```
Identity = "openair-cn.openair4G.eur";
Realm = "openair4G.eur";
ConnectPeer= "hss.openair4G.eur" { ConnectTo = "127.0.0.1"; No_SCTP ; No_IPv6; Prefer_TCP; No_TLS; port = 3868;  realm = "openair4G.eur";};
```

###### 6) mme.conf

```
REALM = "openair4G.eur";

S6A :
    {
        S6A_CONF                   = "/usr/local/etc/oai/freeDiameter/mme_fd.conf"; # YOUR MME freeDiameter config file path
        HSS_HOSTNAME               = "hss";                                         # THE HSS HOSTNAME
    };

GUMMEI_LIST = (
        {MCC="466" ; MNC="68"; MME_GID="4" ; MME_CODE="1"; }                   # YOUR GUMMEI CONFIG HERE
     );

TAI_LIST = (
{MCC="466" ; MNC="68";  TAC = "1"; }                              # YOUR PLMN CONFIG HERE
);

NETWORK_INTERFACES :
    {
        # MME binded interface for S1-C or S1-MME  communication (S1AP), can be ethernet interface, virtual ethernet interface, we don't advise wireless interfaces
        MME_INTERFACE_NAME_FOR_S1_MME         = "eth0";                        # YOUR NETWORK CONFIG HERE
        MME_IPV4_ADDRESS_FOR_S1_MME           = "192.168.31.134/24";            # YOUR NETWORK CONFIG HERE

        # MME binded interface for S11 communication (GTPV2-C)
        MME_INTERFACE_NAME_FOR_S11_MME        = "lo";                          # YOUR NETWORK CONFIG HERE
        MME_IPV4_ADDRESS_FOR_S11_MME          = "127.0.11.1/8";                # YOUR NETWORK CONFIG HERE
        MME_PORT_FOR_S11_MME                  = 2123;                          # YOUR NETWORK CONFIG HERE
    };

S-GW :
{
    # S-GW binded interface for S11 communication (GTPV2-C), if none selected the ITTI message interface is used
    SGW_IPV4_ADDRESS_FOR_S11                = "127.0.11.2/8";            # YOUR NETWORK CONFIG HERE
};
```

此文件中要注意修改MCC，MNC以及MME的网卡及ip地址。

###### 7) spgw.conf

```
S-GW :
{
   NETWORK_INTERFACES :
    {
        # S-GW binded interface for S11 communication (GTPV2-C), if none selected the ITTI message interface is used
        SGW_INTERFACE_NAME_FOR_S11              = "lo";                         # STRING, interface name, YOUR NETWORK CONFIG HERE
        SGW_IPV4_ADDRESS_FOR_S11                = "127.0.11.2/8";               # STRING, CIDR, YOUR NETWORK CONFIG HERE
        # S-GW binded interface for S1-U communication (GTPV1-U) can be ethernet interface, virtual ethernet interface, we don't advise wireless interfaces

        SGW_INTERFACE_NAME_FOR_S1U_S12_S4_UP    = "eth0";                       # STRING, interface name, YOUR NETWORK CONFIG HERE, USE "lo" if S-GW run on eNB host
        SGW_IPV4_ADDRESS_FOR_S1U_S12_S4_UP      = "192.168.31.134/24";           # STRING, CIDR, YOUR NETWORK CONFIG HERE
        SGW_IPV4_PORT_FOR_S1U_S12_S4_UP         = 2152;                         # INTEGER, port number, PREFER NOT CHANGE UNLESS YOU KNOW WHAT YOU ARE DOING

        # S-GW binded interface for S5 or S8 communication, not implemented, so leave it to none
        SGW_INTERFACE_NAME_FOR_S5_S8_UP         = "none";                       # STRING, interface name, DO NOT CHANGE (NOT IMPLEMENTED YET)
        SGW_IPV4_ADDRESS_FOR_S5_S8_UP           = "0.0.0.0/24";                 # STRING, CIDR, DO NOT CHANGE (NOT IMPLEMENTED YET)
};

P-GW =
{
NETWORK_INTERFACES :
    {        # P-GW binded interface for S5 or S8 communication, not implemented, so leave it to none
        PGW_INTERFACE_NAME_FOR_S5_S8          = "none";                         # STRING, interface name, DO NOT CHANGE (NOT IMPLEMENTED YET)
        # P-GW binded interface for SGI (egress/ingress internet traffic)
        PGW_INTERFACE_NAME_FOR_SGI            = "eth0";                         # STRING, YOUR NETWORK CONFIG HERE
        PGW_MASQUERADE_SGI                    = "yes";                           # STRING, {"yes", "no"}. YOUR NETWORK CONFIG HERE, will do NAT for you if you put "yes".
        UE_TCP_MSS_CLAMPING                   = "no";                           # STRING, {"yes", "no"}.
    };
```

此中要注意修改网卡和ip地址。

###### 8) 修改openair-cn/SRC/UTILS/mcc-mnc-itu.c文件，添加{466.68}.

官网链接：https://gitlab.eurecom.fr/oai/openairinterface5g/wikis/HowToConnectCOTSUEwithOAIeNBNew

## 0x36 Compile

#### eNB：

将usrpb210插上USB3.0口进行编译

```
cd ~/openairinterface
source oaienv
cd cmake_targets
./build_oai -I --eNB -x --install-system-files -w USRP --install-optional-packages #for USRP
```

-I: installs required packages.

–eNB: installs eNB, i.e., lte-softmodem.

-x: adds a software oscilloscope feature to the producedbinaries.

–install-system-files: installs OAI required files inLinux system.

-w: adds the hardware support, which is USRP in our case.

–install-optional-packages: Installs optional packages

#### Corenetwork：

###### 1) 安装许可证

```
cd ~/openair-cn/SCRIPTS
./check_hss_s6a_certificate /usr/local/etc/oai/freeDiameter/ hss.openair4G.eur
./check_mme_s6a_certificate /usr/local/etc/oai/freeDiameter/ openair-cn.openair4G.eur
```

###### 2) 编译hss，安装数据库

```
cd ~/openair-cn
cd SCRIPTS
./build_hss -c
./run_hss -i ~/openair-cn/SRC/OAI_HSS/db/oai_db.sql #Run only once to install database
```

此处oai会通过oai_db.sql脚本安装oai_db数据库，登录数据库可查看相关table。如果未成功安装，可后续通过此脚本进行手动添加数据库。

###### 3) 编译mme

```
cd ~/openair-cn/SCRIPTS
./build_mme -c
```

###### 4) 编译spgw

```
cd ~/openair-cn
cd SCRIPTS
./build_spgw -c
```

## 0x37 修改数据库

核心网本机浏览器输入127.0.0.1/phpmyadmin登录phpmyadmin界面对oai_db数据库进行修改，以匹配sim卡信息与MME信息。

###### 1) Table users

选中左侧数据库oai_db，打开表users。点击上方Structure标签，将key，rand，OPc三个值的类型从varbinary改为binary。再点击上方SQL标签，在打开的空白框中输入

```
INSERT INTO users (`imsi`, `msisdn`, `imei`, `imei_sv`, `ms_ps_status`, `rau_tau_timer`, `ue_ambr_ul`, `ue_ambr_dl`, `access_restriction`, `mme_cap`, `mmeidentity_idmmeidentity`, `key`, `RFSP-Index`, `urrp_mme`, `sqn`, `rand`, `OPc`) VALUES ('466680000000001',  '33611123456', '35609204079299', NULL, 'NOT_PURGED', '120', '40000000', '100000000', '47', '0000000000', '1', 0x8BAF473F2F8FD09487CCCBD7097C6862, '1', '0', '', 0x00000000000000000000000000000000, '11111111111111111111111111111111');

```
执行后成功在表中添加了此user。其中imsi为sim卡对应信息（466680000000001）；key值为sim中确定信息（8BAF473F2F8FD09487CCCBD7097C6862）；mmeidentity设为1，要与mmeidentity中信息对应；sqn值可暂时不设，rand值随机设为32位的hex数；OPc值在鉴权中至关重要，它通过数据库中的key值与hss配置文件中设定的OP值计算得出，必须要与sim卡中此值相符，鉴权才可通过。故在配置数据库时一定要确定自己sim卡中的key，OP值，并在数据库和hss.conf文件中正确填写。

##### 问题：
在修改数据库时我们遇到了一些问题，主要是所填信息与sim卡信息不符，填写方式不对等原因造成的。
###### 1.在phpmyadmin数据库中直接点击上方insert进行添加，报错key，rand和OPc三个值数据被截断；
###### 2.直接修改现有的数据，报错数据被截断；
###### 3.在mysql命令行中直接执行添加命令，有warning，原因还是数据被截断；
###### 4.修改key，rand和OPc的数据长度，从varbinary16改为varbinary32，添加数据不会被截断，但是运行hss时数据会自动变化，导致鉴权依旧失败。

最终通过学习查找数据类型的不同，将数据改为binary，并且添加时写成0x的16进制数，才添加成功，同时这也是我们调试系统跑通前解决的最后一个问题。

###### 2) Table mmeidentity
将表格中idmmeidentity=1的那行mmehost改为`openair-cn.openair4G.eur`, mmerealm改为`openair4G.eur`, 与mme配置文件中相符。

###### 3) Table pdn
新添加一行或修改一行，添加数据`users_imsi=466680000000001`.

###### 4) Table pgw
将表中id=3（对应table pdn中的pgw_id值）的那行ipv4改为mme本机地址`192.168.31.134`，ipv6写0.

## 0x38 运行

在所有修改都完成后，可重复编译并运行。在所有模块中一定要先启动hss。

#### Corenetwork：

```
cd ~/openair-cn/SCRIPTS
./build_hss -c
./run_hss #Run for all subsequent runs
cd ~/openair-cn/SCRIPTS
./build_mme -c
./run_mme
cd ~/openair-cn/SCRIPTS
./build_spgw -c
./run_spgw
```

#### eNB:

```
cd ~/openairinterface5g/cmake_targets/lte_build_oai/build
sudo -E ./lte-softmodem -O ~openairinterface5g/targets/PROJECTS/GENERIC-LTE-EPC/CONF/enb.band7.tm1.usrpb210.conf -d
sudo -E ./lte-softmodem -h #(to see help options)
```

各模块运行成功后将sim卡放入HUAWEI Dongle中，断开笔记本电脑的网络连接，将dongle插入，笔记本电脑可正常上网。

当前网络性能不稳定，打开多个网页后会断开连接，还需后续调试。
