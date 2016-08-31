#!/bin/bash

# Fetch Go 1.7
sudo curl -s https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz >> /home/vagrant/go.tar.gz

# Unpack the archive
(cd /home/vagrant && sudo tar -xvf go.tar.gz)

# Create a symlink so 'go' works from the command line
sudo ln -s /home/vagrant/go/bin/go /usr/bin/go

# Create $GOPATH
touch /home/vagrant/.bashrc
echo "export GOPATH=/home/vagrant/gotools" >> /home/vagrant/.bashrc
echo "export GOROOT=/home/vagrant/go" >> /home/vagrant/.bashrc
source /home/vagrant/.bashrc
