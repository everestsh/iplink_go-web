#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Created on May 31 2017
#Filename:
#File tag:
#Function:
#Autor:sooof
#Email:sooof@me.com
#Date:
#Version:0.0001
#TEST:
#20:58:42_@raylea:getopt$ ./TestGetopt.py -ip=192.168.0.1 --port=80 123 a
#ip is: p=192.168.0.1
#port is: 80
#error args: ['123', 'a']



import getopt
import sys

def main(argv):
    try:
            options, args = getopt.getopt(argv, "hp:i:", ["help", "ip=", "port="])
    except getopt.GetoptError:
            sys.exit()

    for option, value in options:
        if option in ("-h", "--help"):
            print("help")
        if option in ("-i", "--ip"):
            print("ip is: {0}".format(value))
        if option in ("-p", "--port"):
            print("port is: {0}".format(value))

    print("error args: {0}".format(args))

if __name__ == '__main__':
    main(sys.argv[1:])
