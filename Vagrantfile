
VAGRANTFILE_API_VERSION = "2"

ARM_COUNT=2
FINGER_COUNT=1

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|

  config.vm.box = "precise64"
  config.vm.box_url = "hashicorp/precise64"

  config.vm.define 'head' do |machine|
    machine.vm.hostname = 'head'
    machine.vm.network "private_network", ip: "192.168.88.100"
    machine.vm.provision "ansible" do |ansible|
      ansible.playbook = "provisioning/playbook.yml"
      ansible.extra_vars = {
        consul_master: 'true',
        consul_bind_addr: "192.168.88.100",
        consul_client_addr: "192.168.88.100"
       }
    end
  end

  (1..ARM_COUNT).each do |arm_num|
    config.vm.define "arm-#{arm_num}" do |machine|
      machine.vm.hostname = "arm-#{arm_num}"
      machine.vm.network "private_network", ip: "192.168.88.15#{arm_num}"
      machine.vm.provision "ansible" do |ansible|
        ansible.playbook = "provisioning/playbook.yml"
        ansible.extra_vars = {
          consul_bind_addr: "192.168.88.15#{arm_num}",
          consul_client_addr: "192.168.88.15#{arm_num}"
         }
      end
    end
  end

  (1..FINGER_COUNT).each do |finger_num|
    config.vm.define "finger-#{finger_num}" do |machine|
      machine.vm.hostname = "finger-#{finger_num}"
      machine.vm.network "private_network", ip: "192.168.88.20#{finger_num}"
      machine.vm.provision "ansible" do |ansible|
        ansible.playbook = "provisioning/playbook.yml"
        ansible.extra_vars = {
          consul_bind_addr: "192.168.88.20#{finger_num}",
          consul_client_addr: "192.168.88.20#{finger_num}"
         }
      end
    end
  end

end
