VM_MEMORY = 1024
VM_CPUS = 1

Vagrant.configure("2") do |config|
  config.vm.box = "generic/centos8"
  config.ssh.insert_key = false

  # database
  config.vm.define "database" do |db|
    db.vm.hostname = "database"
    db.vm.network "private_network", ip: "192.168.56.101"
    db.vm.provider "virtualbox" do |vb|
      vb.memory = VM_MEMORY
      vb.cpus = VM_CPUS
    end
  end

  # backend-1
  config.vm.define "app_instance_1" do |app|
    app.vm.hostname = "app-instance-1"
    app.vm.network "private_network", ip: "192.168.56.102"
    app.vm.provider "virtualbox" do |vb|
      vb.memory = VM_MEMORY
      vb.cpus = VM_CPUS
    end
  end

  # backend-2
  config.vm.define "app_instance_2" do |app|
    app.vm.hostname = "app-instance-2"
    app.vm.network "private_network", ip: "192.168.56.103"
    app.vm.provider "virtualbox" do |vb|
      vb.memory = VM_MEMORY
      vb.cpus = VM_CPUS
    end
  end

  # load-balancer
  config.vm.define "load_balancer" do |web|
    web.vm.hostname = "load-balancer"
    web.vm.network "private_network", ip: "192.168.56.104"
    web.vm.provider "virtualbox" do |vb|
      vb.memory = VM_MEMORY
      vb.cpus = VM_CPUS
    end
  end
end