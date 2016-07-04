require 'yaml'

MACHINES_FILE='group_vars/all'

all_config = YAML::load_file(MACHINES_FILE)
machines = all_config['machines']
template = all_config['vagrant_template']

Vagrant.configure('2') do |config|
  machines.each do |box|
    config.vm.define box['role'] do |machine|
      machine.vm.hostname = "#{box['role']}.vagrant"
      machine.vm.network :private_network, :ip => box['ip']
      machine.vm.box = template
    end
  end

  config.vm.provision "ansible" do |ansible|
    # ansible.verbose = "vvv"
    ansible.playbook = "site.yaml"
    ansible.groups = {
      "web" => ["web"],
      "app" => ["app[1:2]"],
    }
    ansible.extra_vars = {
      run_tests: true,
    }
  end
end
