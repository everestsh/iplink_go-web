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
#TEST:./DeFile_Python.py argv[1] ...

import os
import sys

def main(argv):
    for arg in sys.argv[1:]:
        os.remove(arg)
if __name__ == '__main__':
    main(sys.argv[1:])
