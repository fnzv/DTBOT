package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/telegram-bot-api.v4"
)

var (
	telegram_token string
	chatid         int64
)

type Config struct {
	TelegramToken string
	ChatID        int64
}

// Reads info from config file
func ReadConfig() Config {
	var configfile = "/etc/dtbot/dtbot.conf"
	_, err := os.Stat(configfile)
	if err != nil {
		log.Fatal("Config file is missing: ", configfile)
	}

	var config Config
	if _, err := toml.DecodeFile(configfile, &config); err != nil {
		log.Fatal(err)
	}
	return config
}

func exec_shell(command string) string {
	out, err := exec.Command("/bin/bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

var command string

func main() {
	f, err := os.OpenFile("/var/log/dtbot.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)

	var config = ReadConfig()
	telegram_token = config.TelegramToken
	chatid = config.ChatID
	schatid := strconv.FormatInt(chatid, 10)
	command := ""
	bot, err := tgbotapi.NewBotAPI(telegram_token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Chat.ID == chatid {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "..")


			if strings.Contains(update.Message.Text, "/info") {
					words := strings.Fields(update.Message.Text)
					log.Println("Command: /info")
					log.Println(len(words))
					if len(words) == 2 {
					total_nodes := words[1]
					command = `source /etc/dtbot/os_creds &&  ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook /etc/dtbot/playbooks/info.yaml --extra-vars "telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + ` total_nodes=`+total_nodes+`"`
					log.Println("MESSAGGIO : " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Starting gathering info on nodes ..")
					bot.Send(msg)
					exec_shell(command)
					}  else {

                                        log.Println("Bot missing parameter")
                                }
			}

			if strings.Contains(update.Message.Text, "/load_custom") {
					words := strings.Fields(update.Message.Text)
					log.Println("Command: /load_custom")
					log.Println(len(words))
					if len(words) == 3 {
					url := words[1]
					total_nodes := words[2]
					command = `source /etc/dtbot/os_creds &&  ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook /etc/dtbot/playbooks/ddosb.yaml --extra-vars "telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + ` url=`+url+ ` total_nodes=`+total_nodes+`"`
					log.Println("MESSAGGIO : " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Starting stresser nodes and magic ..")
					bot.Send(msg)
					exec_shell(command)
					} else {

                                        log.Println("Bot missing parameter")
                                }
			}

			if strings.Contains(update.Message.Text, "/loadj") {
					words := strings.Fields(update.Message.Text)
					log.Println("Command: /loadj")
					log.Println(len(words))
					if len(words) == 3 {
					url := words[1]
					total_nodes := words[2]
					command = `source /etc/dtbot/os_creds &&  ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook /etc/dtbot/playbooks/ddosj.yaml --extra-vars "telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + ` url=`+url+ ` total_nodes=`+total_nodes+`"`
					log.Println("MESSAGGIO : " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Starting stresser nodes and magic ..")
					bot.Send(msg)
					exec_shell(command)
							}  else {

                                        log.Println("Bot missing parameter")
                                }
			}


			if strings.Contains(update.Message.Text, "/load_custom_aws ") {
					words := strings.Fields(update.Message.Text)
					log.Println("Command: /load_custom_aws")
					if len(words) == 2 {
					url := words[1]
					command = `source /etc/dtbot/aws_creds &&  ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i /etc/dtbot/ec2.py /etc/dtbot/playbooks/ddosb-aws.yaml --extra-vars "telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + ` url=`+url+ `"`
					log.Println("MESSAGGIO : " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Starting stresser nodes and magic ..")
					bot.Send(msg)
					exec_shell(command)

					} else {

                                        log.Println("Bot missing parameter")
                                }
			}
			if strings.Contains(update.Message.Text, "/loadj_aws ") {
					words := strings.Fields(update.Message.Text)
					log.Println("Command: /loadj_aws")
					if len(words) == 2 {
					url := words[1]
					command = `source /etc/dtbot/aws_creds &&  ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i /etc/dtbot/ec2.py /etc/dtbot/playbooks/ddosj-aws.yaml --extra-vars "telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + ` url=`+url+ `"`
					log.Println("MESSAGGIO : " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Starting stresser nodes and magic ..")
					bot.Send(msg)
					exec_shell(command)

					}  else {

                                        log.Println("Bot missing parameter")
                                }
			}
			if strings.Contains(update.Message.Text, "/load_aws") {

				words := strings.Fields(update.Message.Text)
				if len(words) == 4 {

					domain := words[1]
					concurrency := words[2]
					duration := words[3]
					command = `source /etc/dtbot/aws_creds &&  ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -i /etc/dtbot/ec2.py /etc/dtbot/playbooks/ddos-aws.yaml --extra-vars "concurrency=` + concurrency + ` telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + ` url=` + domain + ` duration=` + duration + `"`
					log.Println("MESSAGGIO : " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Starting stresser nodes and magic ..")
					bot.Send(msg)
					exec_shell(command)

				} else {

					log.Println("Bot missing parameter")
				}
			}

			if strings.Contains(update.Message.Text, "/load ") {

				words := strings.Fields(update.Message.Text)
				if len(words) == 5 {

					domain := words[1]
					concurrency := words[2]
					duration := words[3]
					total_nodes := words[4]
					command = `source /etc/dtbot/os_creds &&  ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook /etc/dtbot/playbooks/ddos.yaml --extra-vars "total_nodes=` + total_nodes + ` concurrency=` + concurrency + ` telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + ` url=` + domain + ` duration=` + duration + `"`
					log.Println("MESSAGGIO : " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Starting stresser nodes and magic ..")
					bot.Send(msg)
					exec_shell(command)

				} else {

					log.Println("Bot missing parameter")
				}

			}

			if strings.Contains(update.Message.Text, "/help") {
				log.Println("Command: /help")
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "/help - shows this output \nExamples: \n /load <URL> <Num Clients> <Seconds> <Num Instances used> \n /create <Num Instances> \n /destroy <Num Instances> \n /stop <Num Instances> ... The Num instances is a sequential number, if i deploy 10 Instances with /create 10 the next time i re-launch the command will not deploy another 10 instances but will only check that instances from 0-10 are present... The same for /stop /load and /destroy\n/create_aws <Num Instances>.. every time the command is launched new instances will be deployed \n/destory_aws (shuts off all VMS Accessible via the ssh key provided so create a separate key.. Stresser nodes will be deleted on shutoff) \n/load_aws <URL> <Num Clients> <Seconds>.. \n/loadj_aws - Downloads custom jmx and execute it on aws .. \n/loadj <URL> <Total nodes> - Downloads custom jmx and execute it on Openstack. .. \n/load_custom <URL> <Total nodes> .. start loadstressers and executes custom bash script provided on Openstack creds .. \n/load_custom_aws <URL> .. start loadstressers and executes custom bash script provided on AWS nodes...")
				bot.Send(msg)
			}

			if strings.Contains(update.Message.Text, "/stop_aws") {
				command = `source /etc/dtbot/aws_creds && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook  -i /etc/dtbot/ec2.py /etc/dtbot/playbooks/aws-stop.yaml --extra-vars="telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + `"&`
				log.Println("Command: " + command)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Force stopping stresser nodes and magic..")
				bot.Send(msg)
				exec_shell(command)

			}

			if strings.Contains(update.Message.Text, "/stop ") {
				text := strings.Fields(update.Message.Text)
				if len(text) == 2 {
					total_nodes := text[1]
					command = `source /etc/dtbot/os_creds && ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook  /etc/dtbot/playbooks/stop.yaml --extra-vars="total_nodes=` + total_nodes + ` telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + `"`
					log.Println("Command: " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Force stopping stresser nodes and magic..")
					bot.Send(msg)
					exec_shell(command)
				} else {

					log.Println("Bot missing parameter")
				}

			}
			if strings.Contains(update.Message.Text, "/create ") {

				text := strings.Fields(update.Message.Text)
				if len(text) == 2 {
					total_nodes := text[1]
					command = `source /etc/dtbot/os_creds && ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -vv /etc/dtbot/playbooks/create-infra.yaml --extra-vars="total_nodes=` + total_nodes + ` telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + `"`
					log.Println("Command: " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Deploying "+total_nodes+" stresser nodes .. Wait a few minutes before starting an attack")
					bot.Send(msg)
					exec_shell(command)
				} else {

					log.Println("Bot missing parameter")
				}
			}

			if strings.Contains(update.Message.Text, "/create_aws ") {

				text := strings.Fields(update.Message.Text)
				if len(text) == 2 {
					total_nodes := text[1]
					command = `source /etc/dtbot/aws_creds && ansible_python_interpreter=/usr/bin/python3 ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -vv /etc/dtbot/playbooks/aws-create-infra.yaml --extra-vars="total_nodes=` + total_nodes + ` telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + `"`
					log.Println("Command: " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Deploying "+total_nodes+" AWS stresser nodes .. Wait a few minutes before starting an attack")
					bot.Send(msg)
					exec_shell(command)
				} else {

					log.Println("Bot missing parameter")
				}
			}
			if strings.Contains(update.Message.Text, "/destroy_aws") {
				command = `source /etc/dtbot/aws_creds && ANSIBLE_HOST_KEY_CHECKING=False ansible_python_interpreter=/usr/bin/python3 ansible-playbook -i /etc/dtbot/ec2.py -vv /etc/dtbot/playbooks/aws-destroy-infra.yaml --extra-vars="telegramtoken=bot` + telegram_token + ` telegramchatid=` + schatid + `"`
				log.Println("Command: " + command)
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Deleting All AWS stresser nodes ...")
				bot.Send(msg)
				exec_shell(command)
			}
			if strings.Contains(update.Message.Text, "/destroy ") {
				text := strings.Fields(update.Message.Text)
				if len(text) == 2 {
					total_nodes := text[1]
					command = `source /etc/dtbot/os_creds && ANSIBLE_HOST_KEY_CHECKING=False ansible_python_interpreter=/usr/bin/python3 ansible-playbook -vv /etc/dtbot/playbooks/destroy-infra.yaml --extra-vars="total_nodes=` + total_nodes + `"`
					log.Println("Command: " + command)
					msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Deleting stresser nodes ...")
					bot.Send(msg)
					exec_shell(command)
				} else {

					log.Println("Bot missing parameter")
				}
			}
		} // close chatid if
	}
}
