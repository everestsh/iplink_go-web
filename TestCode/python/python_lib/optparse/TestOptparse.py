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
#TEST:./

import sys
from optparse import OptionParser
def main():
    usage = "usage: %prog [options] arg"
    parser = OptionParser(usage)
    parser.add_option("-f", "--file", dest="filename",
                      help="read data from FILENAME")
    parser.add_option("-v", "--verbose",
                      action="store_true", dest="verbose")
    parser.add_option("-q", "--quiet",
                      action="store_false", dest="verbose")

    (options, args) = parser.parse_args()
    if len(args) != 1:
        sys.stderr.write("You must specify -? !!!\n")
        #print "reading %s..." % options.filename
        parser.print_help(sys.stderr)
        sys.exit(1)
#        parser.error("incorrect number of arguments")
#    if options
    if options.verbose:
        print "reading %s..." % options.filename

if __name__ == "__main__":
    main()
