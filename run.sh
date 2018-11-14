#! /bin/sh
#
# run.sh
# Copyright (C) 2018 lijiaocn <lijiaocn@foxmail.com>
#
# Distributed under terms of the GPL license.
#

nohup gitbook serve --port 4001 2>&1 >/tmp/book-go.log &
