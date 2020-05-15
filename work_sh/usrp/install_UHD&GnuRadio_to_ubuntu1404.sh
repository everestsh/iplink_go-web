#!/bin/bash
sudo tar xvf archives.tar -C /
sudo apt-get update
sudo apt-get install -y git vim openssh-server openssh-client
git clone https://github.com/sooof/iplinkme_go-web.git
#cp iplinkme_go-web/work_sh/usrp/install\&burn_x31x_uhd_gnuradio/*.sh .
cp iplinkme_go-web/work_sh/usrp/install\&burn_x31x_uhd_gnuradio/020_install_uhd.sh .
cp iplinkme_go-web/work_sh/usrp/install\&burn_x31x_uhd_gnuradio/025_install_gnuradio.sh .
#wget https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/X3XX_FPGA/xilinx2015.4.tar.bz2
wget https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/x3x0_src/gnuradio.tar.bz2
wget https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/x3x0_src/uhd.tar.bz2
#wget https://s3.cn-north-1.amazonaws.com.cn/microembedded/USRP%E4%BA%A7%E5%93%81%E6%8A%80%E6%9C%AF%E8%B5%84%E6%96%99/USRP+X310/X310%E8%B5%84%E6%96%99/x3x0_src/fpga.tar
sudo ./020_install_uhd.sh
sudo ./025_install_gnuradio.sh
sudo /usr/local/lib/uhd/utils/uhd_images_downloader.py
