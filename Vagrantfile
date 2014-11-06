
VAGRANTFILE_API_VERSION = "2"

ARM_COUNT = ENV["ARM_COUNT"] ? ENV["ARM_COUNT"].to_i : 2
FINGER_COUNT = ENV["FINGER_COUNT"] ? ENV["FINGER_COUNT"].to_i : 2
RAM = ENV["RAM"] ? ENV["RAM"].to_i : 256
CPU = ENV["CPU"] ? ENV["CPU"].to_i : 1

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box = "ubuntu/trusty64"
  config.vm.box_url = "ubuntu/trusty64"

  config.vm.define 'head' do |machine|
    machine.vm.hostname = 'head'
    machine.vm.network "private_network", ip: "192.168.93.100"
    machine.vm.provision "ansible" do |ansible|
      ansible.playbook = "provisioning/playbook.yml"
      ansible.extra_vars = {
        consul_node_name: "head",
        consul_ui_require_auth: true,
        consul_is_ui: true,
        consul_is_server: true,
        consul_bootstrap_expect: 3,
        consul_bind_addr: "192.168.93.100",
        consul_client_addr: "192.168.93.100"
       }
    end
  end

  (1..ARM_COUNT).each do |arm_num|
    config.vm.define "arm-#{arm_num}" do |machine|
      machine.vm.hostname = "arm-#{arm_num}"
      machine.vm.network "private_network", ip: "192.168.93.15#{arm_num}"
      machine.vm.provision "ansible" do |ansible|
        ansible.playbook = "provisioning/playbook.yml"
        ansible.extra_vars = {
          consul_node_name: "arm-#{arm_num}",
          consul_is_server: true,
          consul_bootstrap_expect: 3,
          consul_bind_addr: "192.168.93.15#{arm_num}",
          consul_client_addr: "192.168.93.15#{arm_num}"
         }
      end
    end
  end

  (1..FINGER_COUNT).each do |finger_num|
    config.vm.define "finger-#{finger_num}" do |machine|
      machine.vm.hostname = "finger-#{finger_num}"
      machine.vm.network "private_network", ip: "192.168.93.20#{finger_num}"
      machine.vm.provision "ansible" do |ansible|
        ansible.playbook = "provisioning/playbook.yml"
        ansible.extra_vars = {
          consul_node_name: "finger-#{finger_num}",
          consul_bind_addr: "192.168.93.20#{finger_num}",
          consul_client_addr: "192.168.93.20#{finger_num}"
         }
      end
    end
  end

  # Provider specifics
  config.vm.provider "virtualbox" do |v|
    v.customize ["modifyvm", :id, "--memory", RAM]
    v.customize ["modifyvm", :id, "--cpus", CPU]
  end

end
