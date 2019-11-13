#!/bin/bash
# -------------------------------------------------------------------------------
# Filename:    run.sh
# Revision:    1.0
# Date:        2018-09-06
# Author:      ch-yk
# Email:       commonheart.yk@gmail.com
# Website:     www.commonheart-yk.com
# Description: 管理 yk_cgi 的启动和终止
# Notes:       This plugin uses the "./run.sh -h", "./run.sh start", "./run.sh stop", "./run.sh status" command
# -------------------------------------------------------------------------------
# Copyright:   2018 (c) ch-yk
# License:     Apache 2.0
# -------------------------------------------------------------------------------
#Version 1.0
## 1. 提供多个 option 选项
## 2. -h, start, status, stop, restart


SERVER="yk_cgi"
BASE_DIR=$PWD
INTERVAL=2

# 命令行参数，需要手动指定(shell文件里写死)，这里可以指定 init.conf 文件位置
ARGS=""

function start()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER already running"
		exit 1
	fi

	nohup $BASE_DIR/$SERVER $ARGS &>/dev/null &

	echo "starting..." &&  sleep $INTERVAL

	# check status
	if [ "`pgrep $SERVER -u $UID`" == "" ];then
		echo "$SERVER start failed"
		exit 1
	fi
}

function status()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo $SERVER is running
	else
		echo $SERVER is not running
	fi
}

function stop()
{
	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		kill -9 `pgrep $SERVER -u $UID`
	fi

	echo "stoping..." &&  sleep $INTERVAL

	if [ "`pgrep $SERVER -u $UID`" != "" ];then
		echo "$SERVER stop failed"
		exit 1
	fi
}

case "$1" in
	'start')
	start
	;;
	'stop')
	stop
	;;
	'status')
	status
	;;
	'restart')
	stop && start
	;;
	*)
	echo "usage: $0 {start|stop|restart|status}"
	exit 1
	;;
esac