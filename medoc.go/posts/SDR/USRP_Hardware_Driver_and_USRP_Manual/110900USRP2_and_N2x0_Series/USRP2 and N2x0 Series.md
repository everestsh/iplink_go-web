```
{
    "url": "USRP_X3x0_Series",
    "time": "2018/01/06 19:00",
    "tag": "SRD, UHD,USRP Hardware Driver and USRP Manual "
}
```


# USRP2 and N2x0 Series

## Comparative features list

* Hardware Capabilities:
 * 1 transceiver card slot
 * External PPS reference input
 * External 10 MHz reference input
 * MIMO cable shared reference
 * Fixed 100 MHz clock rate
 * Internal GPSDO option (N2x0 only)
* FPGA Capabilities:
 * 2 RX DDC chains in FPGA
 * 1 TX DUC chain in FPGA
 * Timed commands in FPGA (N2x0 only)
 * Timed sampling in FPGA
 * 16-bit and 8-bit sample modes (sc8 and sc16)
   * Up to 25 MHz of RF BW with 16-bit samples
   * Up to 50 MHz of RF BW with 8-bit samples

## Load the Images onto the SD card (USRP2 only)
**Warning!** Use usrp2_card_burner with caution. If you specify the wrong device node, you could overwrite your hard drive. Make sure that --dev= specifies the SD card.

**Warning!** It is possible to use 3rd party SD cards with the USRP2. However, certain types of SD cards will not interface with the CPLD:

* Cards can be SDHC, which is not a supported interface.
* Cards can have unexpected timing characteristics.
For these reasons, we recommend that you use the SD card that was supplied with the USRP2.

Use the card burner tool (UNIX)

```
sudo <install-path>/lib/uhd/utils/usrp2_card_burner_gui.py
-- OR --
cd <install-path>/lib/uhd/utils
sudo ./usrp2_card_burner.py --dev=/dev/sd<XXX> --fpga=<path_to_fpga_image>
sudo ./usrp2_card_burner.py --dev=/dev/sd<XXX> --fw=<path_to_firmware_image>
```
Use the --list option to get a list of possible raw devices. The list result will filter out disk partitions and devices too large to be the sd card. The list option has been implemented on Linux, Mac OS X, and Windows.

### Use the card burner tool (Windows)
```
<path_to_python.exe> <install-path>/lib/uhd/utils/usrp2_card_burner_gui.py
```
## Load the Images onto the On-board Flash (USRP-N Series only)

The USRP-N Series can be reprogrammed over the network to update or change the firmware and FPGA images. When updating images, always burn both the FPGA and firmware images before power cycling. This ensures that when the device reboots, it has a compatible set of images to boot into.

### Use the image loader

Use default images:
```
uhd_image_loader --args="type=usrp2,addr=<IP address>"
```
Use custom-built images:
```
uhd_image_loader --args="type=usrp2,addr=<IP address>" --fw-path="<firmware path>" --fpga-path="<FPGA path>"
```
If you immediately want to apply this image, add the reset argument:
```
uhd_image_loader --args="type=usrp2,addr=<IP address>,reset"
```
The USRP will drop off the network for a brief moment and reboot.

**Note:** Different hardware revisions require different FPGA images. Determine the revision number from the sticker on the rear of the chassis. Use this number to select the correct FPGA image for your device.

### Device recovery and bricking

Its possible to put the device into an unusable state by loading bad images. Fortunately, the USRP-N Series can be booted into a safe (read-only) image. Once booted into the safe image, the user can once again load images onto the device.

The safe-mode button is a pushbutton switch (S2) located inside the enclosure. To boot into the safe image, hold-down the safe-mode button while power-cycling the device. Continue to hold-down the button until the front-panel LEDs blink and remain solid.

When in safe-mode, the USRP-N device will always have the IP address `192.168.10.2`.

For more information on using external tools to unbrick your device when even this fails, see [Unbricking an N-Series Device]().

## Setup Networking

The USRP2 only supports Gigabit Ethernet and will not work with a 10/100 Mbps interface. However, a 10/100 Mbps interface can be connected indirectly to a USRP2 through a Gigabit Ethernet switch.

### Setup the host interface

The USRP2 communicates at the IP/UDP layer over the gigabit Ethernet. The default IP address of the USRP2 is `192.168.10.2`. You will need to configure the host's Ethernet interface with a static IP address to enable communication. An address of `192.168.10.1` and a subnet mask of `255.255.255.0` is recommended.

On a Linux system, you can set a static IP address very easily by using the 'ifconfig' command:
```
sudo ifconfig <interface> 192.168.10.1
```
Note that interface is usually something like `eth0`. You can discover the names of the network interfaces in your computer by running `ifconfig` without any parameters:
```
ifconfig -a
```
**Note:** When using UHD software, if an IP address for the USRP2 is not specified, the software will use UDP broadcast packets to locate the USRP2. On some systems, the firewall will block UDP broadcast packets. It is recommended that you change or disable your firewall settings.

### Multiple devices per host

For maximum throughput, one Ethernet interface per USRP2 is recommended, although multiple devices may be connected via a Gigabit Ethernet switch. In any case, each Ethernet interface should have its own subnet, and the corresponding USRP2 device should be assigned an address in that subnet. Example:

* Configuration for USRP2 device 0:
 * Ethernet interface IPv4 address: `192.168.10.1`
 * Ethernet interface subnet mask: `255.255.255.0`
 * USRP2 device IPv4 address: `192.168.10.2`
* Configuration for USRP2 device 1:
 * Ethernet interface subnet mask: `255.255.255.0`
 * USRP2 device IPv4 address: `192.168.20.2`

### Change the USRP2's IP address

You may need to change the USRP2's IP address for several reasons:

* to satisfy your particular network configuration
* to use multiple USRP2s on the same host computer
* to set a known IP address into USRP2 (in case you forgot)

**Method 1**

To change the USRP2's IP address, you must know the current address of the USRP2, and the network must be setup properly as described above. Run the following commands: :
```
cd <install-path>/lib/uhd/utils
./usrp_burn_mb_eeprom --args=<optional device args> --values="ip-addr=192.168.10.3"
```

**Method 2 (Linux Only)**

This method assumes that you do not know the IP address of your USRP2. It uses raw Ethernet packets to bypass the IP/UDP layer to communicate with the USRP2. Run the following commands:
```
cd <install-path>/lib/uhd/utils
sudo ./usrp2_recovery.py --ifc=eth0 --new-ip=192.168.10.3
```

## Communication Problems

When setting up a development machine for the first time, you may have various difficulties communicating with the USRP device. The following tips are designed to help narrow down and diagnose the problem.

### RuntimeError: no control response

This is a common error that occurs when you have set the subnet of your network interface to a different subnet than the network interface of the USRP device. For example, if your network interface is set to `192.168.20.1`, and the USRP device is `192.168.10.2` (note the difference in the third numbers of the IP addresses), you will likely see a 'no control response' error message.

Fixing this is simple - just set the your host PC's IP address to the same subnet as that of your USRP device. Instructions for setting your IP address are in the previous section of this documentation.

### Firewall issues

When the IP address is not specified, the device discovery broadcasts UDP packets from each Ethernet interface. Many firewalls will block the replies to these broadcast packets. If disabling your system's firewall or specifying the IP address yields a discovered device, then your firewall may be blocking replies to UDP broadcast packets. If this is the case, we recommend that you disable the firewall or create a rule to allow all incoming packets with UDP source port `49152`.

### Ping the device

The USRP device will reply to ICMP echo requests. A successful ping response means that the device has booted properly and that it is using the expected IP address.
```
ping 192.168.10.2
```

### Monitor the serial output

Read the serial port to get debug verbose output from the embedded microcontroller. The microcontroller prints useful information about IP addresses, MAC addresses, control packets, fast-path settings, and bootloading. Use a standard USB to 3.3v-level serial converter at 230400 baud. Connect `GND` to the converter ground, and connect `TXD` to the converter receive. The `RXD` pin can be left unconnected as this is only a one-way communication.

* **USRP2:** Serial port located on the rear edge
* **N210:** Serial port located on the left side

### Monitor the host network traffic

Use Wireshark to monitor packets sent to and received from the device.

## Addressing the Device

### Single device configuration

In a single-device configuration, the USRP device must have a unique IPv4 address on the host computer. The USRP can be identified through its IPv4 address, resolvable hostname, or by other means. See the application notes on Device Identification. Please note that this addressing scheme should also be used with the **multi_usrp** interface.

Example device address string representation for a USRP2 with IPv4 address `192.168.10.2`:
```
addr=192.168.10.2
```

### Multiple device configuration

In a multi-device configuration, each USRP device must have a unique IPv4 address on the host computer. The device address parameter keys must be suffixed with the device index. Each parameter key should be of the format <key><index>. Use this addressing scheme with the uhd::usrp::multi_usrp interface.

* The order in which devices are indexed corresponds to the indexing of the transmit and receive channels.
* The key indexing provides the same granularity of device identification as in the single device case.
Example device address string representation for 2 USRP2s with IPv4 addresses `192.168.10.2` and `192.168.20.2`:
```
addr0=192.168.10.2, addr1=192.168.20.2
```

## Using the MIMO Cable

The MIMO cable allows two USRP devices to share reference clocks, time synchronization, and the Ethernet interface. One of the devices will sync its clock and time references to the MIMO cable. This device will be referred to as the slave, and the other device, the master.

* The slave device acquires the clock and time references from the master device.
* The master and slave may be used individually or in a multi-device configuration.
* External clocking is optional and should only be supplied to the master device.

### Shared Ethernet mode

In shared Ethernet mode, only one device in the configuration can be attached to the Ethernet.

* Clock reference, time reference, and data are communicated over the MIMO cable.
* Master and slave must have different IPv4 addresses in the same subnet.

### Dual Ethernet mode

In dual Ethernet mode, both devices in the configuration must be attached to the Ethernet.

Only clock reference and time reference are communicated over the MIMO cable.
The master and slave must have different IPv4 addresses in different subnets.

### Configuring the slave

In order for the slave to synchronize to the master over MIMO cable, the following clock configuration must be set on the slave device: :
```
usrp->set_time_source("mimo", slave_index);
usrp->set_clock_source("mimo", slave_index);
```

## Alternative stream destination

It is possible to program the USRP device to send RX packets to an alternative IP/UDP destination.

### Set the subnet and gateway

To use an alternative streaming destination, the device needs to be able to determine if the destination address is within its subnet, and ARP appropriately. Therefore, the user should ensure that subnet and gateway addresses have been programmed into the device's EEPROM.

Run the following commands:
```
cd <install-path>/lib/uhd/utils
./usrp_burn_mb_eeprom --args=<optional device args> --values="subnet=255.255.255.0, gateway=192.168.10.2"
```

### Create a receive streamer

Set the stream args "addr" and "port" values to the alternative destination. Packets will be sent to this destination when the user issues a stream command.
```
//create a receive streamer, host type does not matter
uhd::stream_args_t stream_args("fc32");
//resolvable address and port for a remote udp socket
stream_args.args["addr"] = "192.168.10.42";
stream_args.args["port"] = "12345";
//create the streamer
uhd::rx_streamer::sptr rx_stream = usrp->get_rx_stream(stream_args);
//issue stream command
uhd::stream_cmd_t stream_cmd(uhd::stream_cmd_t::STREAM_MODE_NUM_SAMPS_AND_DONE);
stream_cmd.num_samps = total_num_samps;
stream_cmd.stream_now = true;
usrp->issue_stream_cmd(stream_cmd);
```
**Note:** Calling recv() on this streamer object should yield a timeout.

## Hardware Setup Notes

### Front panel LEDs

The LEDs on the front panel can be useful in debugging hardware and software issues. The LEDs reveal the following about the state of the device:

* **LED A:** transmitting
* **LED B:** MIMO cable link
* **LED C:** receiving
* **LED D:** firmware loaded
* **LED E:** reference lock
* **LED F:** CPLD loaded

### Ref Clock - 10 MHz

Using an external 10 MHz reference clock, a square wave will offer the best phase noise performance, but a sinusoid is acceptable. The reference clock requires the following power level:

* **USRP2** 5 to 15 dBm
* **N2XX** 0 to 15 dBm

### PPS - Pulse Per Second

Using a PPS signal for timestamp synchronization requires a square wave signal with the following amplitude:

* **USRP2** 5Vpp
* **N2XX** 3.3 to 5Vpp

Test the PPS input with the following app:

* <args> are device address arguments (optional if only one USRP device is on your machine)
```
cd <install-path>/lib/uhd/examples ./test_pps_input ???args=<args>
```

### Internal GPSDO

Please see [Internal GPSDO (USRP-N2x0/E1X0 Models)]() for information on configuring and using the internal GPSDO.

## Miscellaneous

### Available Sensors

The following sensors are available for the USRP2/N-Series motherboards; they can be queried through the API.

* **mimo_locked** - clock reference locked over the MIMO cable
* **ref_locked** - clock reference locked (internal/external)
* other sensors are added when the GPSDO is enabled

### Multiple RX channels

There are two complete DDC chains in the FPGA. In the single channel case, only one chain is ever used. To receive from both channels, the user must set the **RX** subdevice specification. This hardware has only one daughterboard slot, which has been aptly named slot **A**.

In the following example, a TVRX2 is installed. Channel 0 is sourced from subdevice RX1, and channel 1 is sourced from subdevice RX2 (RX1 and RX2 are the antenna ports on the TVRX2 daughterboard):
```
usrp->set_rx_subdev_spec("A:RX1 A:RX2");
```

### Unbricking an N-Series Device

You'll need:

* JTAG programmer: please connect it to the JTAG connector on the motherboard as shown in the attachment
* Xilinx 'iMPACT': launch and cancel the new project wizards. You should be left with the screen which shows a single FPGA chip in the main document (auto-detected by the programmer).

![N2xx-JTAG](imgs/N2xx-JTAG.jpg)

<center>N2x0 JTAG Connection<center>


Download the latest FPGA images, e.g. using `uhd_images_downloader`.

There is a sub-directory in the archive below the firmware/images called 'bit'. Use Impact to load `usrp_n210_r4_fpga.bit` via the programmer (the filename may be different depending on your device type and revision).

The USRP should now be able to communicate on the network (you'll see some LEDs light up and network link be established). The next step is to flash the device and program the serial number. Both these steps can be done with UHD (the JTAG step is complete).

To be sure, run `uhd_find_devices` and it should appear in the list - remember this IP address for the image loader utility (should be 192.168.10.2 - make sure your network settings enable to you communicate with that subnet!).

The first step is to flash the unit's safe-mode image, and then do a normal flash - both with the UHD Image Loader utility.

Make sure you have UHD installed, and the images from before, and follow the instructions in [Load the Images onto the SD card (USRP2 only)](). You can combine the `--fw-path` and `--fpga-path` arguments into the single invocation of the image loader.

You will probably use ``usrp_n210_fw.bin`` for the firmware and ``usrp_n210_r4_fpga.bin`` for the FPGA image parameters (use the full/relative file path if your current directory is not that of the images).
```
uhd_image_loader --args="type=usrp2,addr=192.168.10.2,overwrite-safe" --fw-path=usrp_n210_fw.bin --fpga-path=usrp_n210_r4_fpga.bin
```
Use the overwrite-safe option the first time, and then repeat without it for the second time. Don't forget to power-cycle the device after it has been flashed.

You can change the normal IP address by following the instructions in [Change the USRP2's IP address]().

If you run `uhd_usrp_probe`, you can see the EEPROM keys at the top. Example:
```
Mboard: N210r4
 hardware: 2577
 mac-addr: a0:36:fa:25:34:a7
 ip-addr: 192.168.10.4
 subnet: 255.255.255.255
 gateway: 255.255.255.255
 gpsdo: none
 serial: EAR14U7UP
```

If you need to change any of there, you should then be able to run:
```
usrp_burn_mb_eeprom --key=<key> --val=<val>
```

to set the `mac-addr`, `serial` and `Mboard`.
