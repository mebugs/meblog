#!/usr/bin/env bash
export ENV=prod
export NODE=APP01
app="smart"
logsFolder="../logs"

echo "Killing $app process..."
ps -ef | grep -i $app |grep -v "grep"
killall $app

echo "Starting $app process ..."
mkdir -p $logsFolder

nohup ./$app >> $logsFolder/$app.log 2>&1 &
ps -ef | grep -i $app |grep -v "grep"

echo "Sleep 5 seconds..."
sleep 5
tail -n 30 $logsFolder/$app.log
