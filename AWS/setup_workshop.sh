#!/bin/bash

echo "Setting up"

sudo apt-get update
sudo apt-get install -y mysql-server mysql-client curl docker.io

wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz
sudo ln -s /usr/local/go/bin/go /bin/go

curl --proto '=https' --tlsv1.2 -sSf https://tiup-mirrors.pingcap.com/install.sh | sh
. .bashrc
tiup install playground pd tidb tikv tiflash grafana prometheus ctl

if [ ! -e dm-master ]; then
  wget https://download.pingcap.org/tidb-community-toolkit-v7.2.0-linux-amd64.tar.gz
  tar zxf tidb-community-toolkit-v7.2.0-linux-amd64.tar.gz \
  	tidb-community-toolkit-v7.2.0-linux-amd64/dmctl-v7.2.0-linux-amd64.tar.gz \
  	tidb-community-toolkit-v7.2.0-linux-amd64/dm-master-v7.2.0-linux-amd64.tar.gz \
  	tidb-community-toolkit-v7.2.0-linux-amd64/dm-worker-v7.2.0-linux-amd64.tar.gz
  
  tar zxf tidb-community-toolkit-v7.2.0-linux-amd64/dm-master-v7.2.0-linux-amd64.tar.gz dm-master/dm-master -C . --strip-components=1
  tar zxf tidb-community-toolkit-v7.2.0-linux-amd64/dm-worker-v7.2.0-linux-amd64.tar.gz dm-worker/dm-worker -C . --strip-components=1
  tar zxf tidb-community-toolkit-v7.2.0-linux-amd64/dmctl-v7.2.0-linux-amd64.tar.gz dmctl/dmctl -C . --strip-components=1
fi

if [ ! -e demoblog ]; then
  git clone https://github.com/dveeden/demoblog.git
fi

sudo docker pull docker.redpanda.com/redpandadata/redpanda:latest

sudo mysql -e "CREATE USER 'blog'@'%' IDENTIFIED BY 'blog'"
sudo mysql -e 'CREATE SCHEMA blog COLLATE utf8mb4_general_ci'
sudo mysql -e "GRANT ALL ON *.* TO 'blog'@'%'"
sudo mysql < demoblog/sql/0001_schema.sql
sudo mysql < demoblog/sql/0002_data.sql 
sudo mysql < demoblog/sql/0003_index.sql

pushd demoblog
go build
pushd loadGen
go build
popd
pushd ticketStat
go build
popd
popd

# Cleanup
rm -rf go1.20.6.linux-amd64.tar.gz tidb-community-toolkit-v7.2.0-linux-amd64 tidb-community-toolkit-v7.2.0-linux-amd64.tar.gz


