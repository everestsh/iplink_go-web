```
{
    "url": "Install_Xilinx_Vivado_SDK_Lin_2015",
    "time": "2018/01/01 19:00",
    "tag": "SRD, LDE,环境部署"
}
```

# Install Xilinx Vivado 2015.4 on Ubuntu 14.04

The user guide for Xilinx Vivado 2015.4 installation is [UG973](http://www.xilinx.com/support/documentation/sw_manuals/xilinx2015_4/ug973-vivado-release-notes-install-license.pdf). More info about Vivado Design Suite can be found on [Xilinx Products](http://www.xilinx.com/products/design-tools/vivado.html) page.

Firstly, go to the [Xilinx Downloads](http://www.xilinx.com/support/download/index.html/content/xilinx/en/downloadNav/vivado-design-tools.html) page to obtain the installer. Select version 2015.4 on the left sidebar. I use the “Single File Download” option and choose “Vivado HLx 2015.4: Full Installer For Linux Single File Download Image Including SDK”. It is a tarball that is 8.68 GB large. Note: you have to be a registered user to download it.

Now, untar the tarball, cd into the extracted directory, and execute the GUI installer.

```
tar -zxvf Xilinx_Vivado_SDK_Lin_2015.4_1118_2.tar.gz
cd Xilinx_Vivado_SDK_Lin_2015.4_1118_2
sudo ./xsetup
```

Agree to the terms and choose Vivado HL WebPACK Edition. Next, tick also Xilinx Software Development Kit (SDK) on the next page, since it’s free and very useful.

At the end of the installation, the license manager will ask for a license. The “Obtain a license” button in the license manager didn’t work for me, so I just go to Xilinx Licensing site directly and get a WebPACK license and install it.

Done that, install the JTAG cable drivers that are needed for many purposes e.g. programming the hardware. Note: by default, Vivado is installed into the /opt/Xilinx/Vivado/2015.4 directory.

```
cd /opt/Xilinx/Vivado/2015.4/data/xicom/cable_drivers/lin64/install_script/install_drivers
sudo ./install_drivers
```

Now, change the ownership of the .Xilinx directory so that you may use Vivado without superuser privilege:

```
sudo chown -hR $USER:$USER $HOME/.Xilinx/
```

Every time you want to fire up Vivado, remember to source the “settings” scripts to have the right environment variables:

```
source /opt/Xilinx/Vivado/2015.4/settings64.sh
source /opt/Xilinx/SDK/2015.4/settings64.sh
```

Finally, fire up Vivado:

```
vivado
```

or SDK:

```
xsdk
```

(alternatively, you can call SDK from within Vivado.
