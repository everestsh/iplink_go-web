#!/usr/bin/env python
# -*- coding: utf-8 -*-
# Created on May 31 2017
# 
import os
import requests

print(os.getcwd())

r = requests.get("http://www.baidu.com")

print(r.url) 
