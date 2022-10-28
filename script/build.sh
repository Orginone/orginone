#!/bin/bash

echo "停止并移除之前的进程"
devpath=$orginone/release
if [ -d $devpath ]; then
    cd $devpath/development
    bash stop.sh
    cd ../../
fi
rm -rf $devpath/development/*

echo "编译新版本"
go build -o $devpath/development/system_api.run $orginone/app/system/cmd/api/system.go
go build -o $devpath/development/system_rpc.run $orginone/app/system/cmd/rpc/system.go
go build -o $devpath/development/user_api.run $orginone/app/user/cmd/api/user.go
go build -o $devpath/development/user_rpc.run $orginone/app/user/cmd/rpc/user.go
go build -o $devpath/development/market_api.run $orginone/app/market/cmd/api/market.go
go build -o $devpath/development/market_rpc.run $orginone/app/market/cmd/rpc/market.go
cd $devpath/development
cp -r ../scripts/* .
chmod +x *
mkdir logs

echo "启动etcd"
etcdNum=`ps -ef|grep etcd|wc -l`
if [ $etcdNum -lt 2 ];then
    if [ ! -d $devpath/etcd ]; then
        mkdir $devpath/etcd
    fi
    nohup /usr/bin/etcd --advertise-client-urls 'http://localhost:2380' --listen-client-urls 'http://localhost:1003' --data-dir $devpath/etcd >> $devpath/etcd/etcd.log  2>&1 &
fi

echo "启动服务"
bash start.sh

echo "启动nginx"
nginxNum=`ps -ef|grep nginx|wc -l`
if [ $nginxNum -lt 2 ];then
    nginx
fi
echo "查看进程"
pids=`cat $orginone/release/development/pid.log |xargs | tr ' ' ','`
top -p $pids -d 0.5