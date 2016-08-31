Vagrant.configure(2) do |config|

  config.vm.box = "phusion/ubuntu-14.04-amd64"
  config.vm.network "private_network", ip: "192.168.33.10"
  config.vm.synced_folder ".", "/home/vagrant/docktor"

  # Install Git 2.x
  config.vm.provision "shell", inline: "apt-get install git -y"

  # Install Docker v1.x
  config.vm.provision "shell", path: "docker.sh"

  # Install Go 1.7
  config.vm.provision "shell", path: "go.sh" 

end