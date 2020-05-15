#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Created on May 31 2017
#Filename:
#File tag:
#Function:
#Autor:sooof
#Email:sooof@me.com
#Date:
#Version:0.0003
#TEST:./CreatFile_Python.py argv[1] ...


import sys

def main(argv):
    for arg in sys.argv[1:]:
        f=open(arg,'w')
        f.close()
if __name__ == '__main__':
    main(sys.argv[1:])
