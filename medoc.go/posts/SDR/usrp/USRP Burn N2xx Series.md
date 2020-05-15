```
{
    "url": "usrp n2xx 的烧写",
    "time": "2018/01/02 19:00",
    "tag": "SRD, LDE,环境部署"
}
```

#usrp n2xx 的烧写
1>
**update FW and FPGA bin file**

```
cd /microembedded/burn/
chmod 777 ./usrp_n210_image_burn.sh
./usrp_n210_image_burn.sh
```
2>
###### USRP FAQ:

```
root@ubuntu:/home/ubuntu# sudo /usr/local/lib/uhd/utils/usrp_n2xx_simple_net_burner --addr="192.168.10.2"
linux; GNU C++ version 4.6.3; Boost_104800; UHD_003.007.000-0-unknown

Searching for USRP N2XX with IP address 192.168.10.2.
Found n200_r4.

Searching for specified images.

Will burn the following images:
 * Firmware: /usr/local/share/uhd/images/usrp_n200_fw.bin
 * FPGA:     /usr/local/share/uhd/images/usrp_n200_r4_fpga.bin

Querying n200_r4 for flash information.
 * Flash size:  4194304
 * Sector size: 65536

Erasing FPGA image.
 * Successfully erased 1572864 bytes at 1572864.
Writing FPGA image (100%).
 * Successfully wrote 890592 bytes.
Verifying FPGA image (100%).
 * Successful.

Erasing firmware image.
 * Successfully erased 31744 bytes at 3145728.
Writing firmware image (100%).
 * Successfully wrote 16383 bytes.
Verifying firmware image (100%).
 * Successful.

Image burning successful. Reset USRP (Y/n)? y

Resetting USRP.
```
3>
