#!/bin/bash

echo "cp python File"
cp -a python/* ~/Workspace/bin
echo "export PATH=$PATH:~/Workspace/bin"
echo "export PATH=$PATH:~/Workspace/bin">> ~/.bash_profile
source  ~/.bash_profile
exit 0;
