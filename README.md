# devops-demo

This is a sample code for a technical interview.

The requirements consisted of deploying two app nodes and one reverse proxy service with configuration management tool of choice and to demonstrate that round-robin works across the app nodes.

**Requirements**

- Vagrant 1.8+
- Virtualbox 5.0+
- Ansible 2.1+

**Usage**

To start:

```
> vagrant up --provider virtualbox
```

To cleanup:

```
> vagrant destroy -f
```

**Configuration**

IP addresses for the VMs are configured in ```group_vars/all``` and can be changed in case when the ```192.168.33.0/24``` network is used locally.

Sample app code is present in ```roles/app/files/app.go```. It can be changed to demonstrate that the code can be redeployed to achieve continous improvement using ```vagrant provision```.

More verbose output from Ansible can be configured in ```Vagrantfile``` by adjusting ```ansible.verbose``` parameter to expected value.

**Demo**

```
> vagrant up --provider virtualbox
Bringing machine 'app1' up with 'virtualbox' provider...
Bringing machine 'app2' up with 'virtualbox' provider...
Bringing machine 'web' up with 'virtualbox' provider...
==> app1: Importing base box 'centos/7'...
[...]
TASK [web : Run tests] *********************************************************
[...]

TASK [web : Fetching data from web] ********************************************
ok: [web]

TASK [web : Expecting app1 in return] ******************************************
ok: [web]

TASK [web : Fetching data from web] ********************************************
ok: [web]

TASK [web : Expecting app2 in return] ******************************************
ok: [web]

TASK [web : Fetching data from web] ********************************************
ok: [web]

TASK [web : Expecting app1 in return] ******************************************
ok: [web]

TASK [web : Fetching data from web] ********************************************
ok: [web]

TASK [web : Expecting app2 in return] ******************************************
ok: [web]
[...]
```

You can also check the output by hand:
```
> curl 192.168.33.100
Hi there, I'm served from app1.vagrant!
> curl 192.168.33.100
Hi there, I'm served from app2.vagrant!
```
