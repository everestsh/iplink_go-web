#!/bin/bash

echo -e "installing wget...\n"
#brew install wget
echo -e "wget pysim_cmd...\n"
mkdir pysim
cd pysim
git init
git remote add -f origin https://github.com/sooof/iplinkme_go-web.git
git config core.sparsecheckout true
echo "work_sh"      >> .git/info/sparse-checkout
git pull origin master
exit 0;
