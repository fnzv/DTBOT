


mkdir /etc/dtbot/

rsync -avz playbooks/ /etc/dtbot/playbooks/
rsync -avz conf/ /etc/dtbot/
cp bin/dtbot_ctrl /usr/bin/

mkdir /etc/dtbot/key

echo "Installing Ansible and requirements.."
apt-get install -y python-pip
pip install ansible
pip install shade
pip install boto

echo "Generating SSH-key for dtbot.."
ssh-keygen -t rsa -N "" -f /etc/dtbot/key/key

echo "Configuring dtbot service.."
cp setup/dtbot.service /etc/systemd/system/dtbot.service
systemctl daemon-reload

echo "After you changed the conf under /etc/dtbot/ you can start dtbot with 'service dtbot start'"
echo "Log file is /var/log/dtbot.log"
