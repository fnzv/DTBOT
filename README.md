

## DTbot
A loadstressing bot managed via Telegram supporting Openstack & AWS
![](img/dtbot-logo.png?raw=true)

-----------------

## Requirements

- Ubuntu 16.04
- Ansible (current version: ansible 2.5.3), Golang (if you want to compile it yourself)


### Quickstart

Run the bash script (install.sh) to install all the required dependencies.

```bash install_script.sh```

<br>
This install script will copy the current compiled binary from /bin into your current system and configure systemd.
Installed packages:<br>
- python-pip (modules: ansible,shade,boto)<br>

To start the bot process run:
```service dtbot start```
<br>

If you encounter any issue check the log file: /var/log/dtbot.log



### Usage

All the available commands can be viewed by sending the message /help to the Bot:
```
/help - shows this output
Examples:
/load <URL> <Num Clients> <Seconds> <Num Instances used>
/create <Num Instances>
/destroy <Num Instances>
/stop <Num Instances> ... The Num instances is a sequential number, if i deploy 10 Instances with /create 10 the next time i re-launch the command will not deploy another 10 instances but will only check that instances from 0-10 are present... The same for /stop /load and /destroy
/create_aws <Num Instances>.. every time the command is launched new instances will be deployed
/destory_aws (shuts off all VMS Accessible via the ssh key provided so create a separate key.. Stresser nodes will be deleted on shutoff)
/load_aws <URL> <Num Clients> <Seconds>..
/loadj_aws - Downloads custom jmx and execute it on aws ..
/loadj <URL> <Total nodes> - Downloads custom jmx and execute it on Openstack. ..
/load_custom <URL> <Total nodes> .. start loadstressers and executes custom bash script provided on Openstack creds ..
/load_custom_aws <URL> .. start loadstressers and executes custom bash script provided on AWS nodes...
```


## Work Flow
This is the basic concept explained with a diagram:
![](img/dtbot-diagram.png?raw=true)

User --> Telegram bot --> DTbot (Golang) --> Ansible

## Screenshots
![](img/dtbot-tg.png?raw=true)


Tested on Ubuntu 16.04 and Ansible 2.5.3

## Contributors

Feel free to open issues or send me an email


## License

Code distributed under GPLv3 licence.
