sudo chown -hR $USER:$USER $HOME/.Xilinx/
source  /opt/Xilinx/Vivado/2015.4/settings64.sh
source /opt/Xilinx/SDK/2015.4/settings64.sh
cd ~/fpga/usrp3/top/x300/
source setupenv.sh
viv_jtag_program /usr/local/share/uhd/images/usrp_x310_fpga_HG.bit
