//   Account Generator Bot
//   Copyright (C) 2021 AnonyIndian (@xnony)

//   This program is distributed in the hope that it will be useful,
//   but WITHOUT ANY WARRANTY; without even the implied warranty of
//   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//   GNU Affero General Public License for more details.


package main

import (
	"fmt"
	"math/rand"
	"strings"
	"strconv"
	"time"
	"os"
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

var MSG string

func main() {
	// Put Your Bot Token via ENV Vars
	b, err := gotgbot.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	// Create updater and dispatcher.
	updater := ext.NewUpdater(b, nil)
	dispatcher := updater.Dispatcher

	// Handlers for runnning commands.
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("gen", gen))

	err = updater.StartPolling(b, &ext.PollingOpts{Clean: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\nMade with ❤️ by @FlayeRNetworK.\n", b.User.Username)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}


func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		// Put your accounts here
		// "email:pass",
		// Following are some demo accounts
		"jack.verma159@gmail.com:Bringo159",
"sandeeppahariya@gmail.com:mantra123",
"sujeetmandal@gmail.com:123456",
"ridhimaplah@gmail.com:bollywood",
"manav_motwani@ymail.com:manavrox",
"torbpatil@yahoo.com:r1a2j3a4",
"badgujar.deepak@gmail.com:deepak@123",
"tapashalder21@gmail.com:tapas@1970",
"bibhu.padhi1@gmail.com:Epuninky@123",
"nitishraj27@gmail.com:sagar123",
"ashwinijagtap11295@gmail.com:9673999133",
"panditavibha@gmail.com:panditavibha",
"vinodbhadala@gmail.com:4656again",
"ankitamona12@gmail.com:saroj1435",
"hsinha343@gmail.com:Inspire@110",
"bajaj.manav@gmail.com:abeer2029",
"neerajjamwal9727@gmail.com:Swift@9727",
"pattabhi.ram007@gmail.com:Kathyayani$1",
"mirk2005@gmail.com:Qatar@123",
	}

	return formats[rand.Intn(len(formats))]
}

func start(ctx *ext.Context) error {
	// To ensure bot does not reply outside of private chats
	if ctx.EffectiveChat.Type != "private" {
		return nil
	}
	// Following string is replied to cmd user on /start 
	MSG = "*Hi %v,\n" +
		"I am an Account Generator Bot\n" +
		"-------------------------------------------------\n" +
		"I can provide premium accounts of different services\n" +
		"--------------------------------------------------\n" +
		"Do /gen to generate an account\n" +
		"--------------------------------------------------\n" +
		"❤️Brought to You By @FlayeRNetworK❤️\n*"

	user := ctx.EffectiveUser
	channel_id, cerror := strconv.Atoi(os.Getenv("CHANNEL_ID"))
	if cerror != nil {fmt.Println("Please Provide me a valid Channel ID")}
	member, eror := ctx.Bot.GetChatMember(int64(channel_id), user.Id)
	if eror != nil {
		ctx.Bot.SendMessage(ctx.EffectiveChat.Id, "*Bot not admin in JoinCheck Channel.*", nil)
		return nil
	}
	// For Checking either user joined channel or not
	if member.Status == "member" || member.Status == "administrator" || member.Status == "creator" {
		_, err := ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf(MSG, user.FirstName), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})
		if err != nil {
			fmt.Println("failed to send: " + err.Error())
		}
	} else {
		// An Error message replied to command user if he's not member of the JoinCheck Channel
		ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf("*You must join %v To use me.*", os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{ParseMode: "Markdown"})
	}
	if strings.ToLower(os.Getenv("LOGS")) != "false"{
		logs, log_err := strconv.Atoi(os.Getenv("LOGS"))
		if log_err != nil{fmt.Println(log_err.Error())}
		// Following message is sent in Logs Group (if set)
		ctx.Bot.SendMessage(int64(logs), fmt.Sprintf("#Start\n\nBot Started By %v(%v)", user.FirstName, user.Id), nil)
	}
	return nil
}

func gen(ctx *ext.Context) error {
	// To ensure bot does not reply outside of private chats
	if ctx.EffectiveChat.Type != "private" {
		return nil
	}
	Combo := strings.Split(randomFormat(), ":")
	// Following string is replied to cmd user on /gen 
	MSG = "𝙃𝙚𝙧𝙚 𝙄𝙨 𝙔𝙤𝙪𝙧 %v 𝘼𝙘𝙘𝙤𝙪𝙣𝙩" +
		"\n\n𝙀𝙢𝙖𝙞𝙡: `%v`" +
		"\n𝙋𝙖𝙨𝙨: `%v`" +
		"\n𝙂𝙚𝙣𝙚𝙧𝙖𝙩𝙚𝙙 𝘽𝙮: *%v*" +
		"\n\n𝙏𝙝𝙖𝙣𝙠 𝙮𝙤𝙪 𝙛𝙤𝙧 𝙪𝙨𝙞𝙣𝙜 𝙢𝙚!\n❤️𝙎𝙝𝙖𝙧𝙚 & 𝙎𝙪𝙥𝙥𝙤𝙧𝙩 *%v*❤️"

	user := ctx.EffectiveUser
	channel_id, cerror := strconv.Atoi(os.Getenv("CHANNEL_ID"))
	if cerror != nil {fmt.Println("Please Provide me a valid Channel ID")}
	member, eror := ctx.Bot.GetChatMember(int64(channel_id), user.Id)
	if eror != nil {
		ctx.Bot.SendMessage(ctx.EffectiveChat.Id, "Bot not admin in JoinCheck Channel", nil)
		return nil
	}
	// For Checking either user joined channel or not
	if member.Status == "member" || member.Status == "administrator" || member.Status == "creator" {
		_, err := ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf(MSG, os.Getenv("ACC_NAME"), Combo[0], Combo[1], user.FirstName, os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{
			ParseMode: "Markdown",
		})
		if err != nil {
			fmt.Println("failed to send: " + err.Error())
		}
	} else {
		// An Error message replied to command user if he's not member of the JoinCheck Channel
		ctx.EffectiveMessage.Reply(ctx.Bot, fmt.Sprintf("*You must join %v to use me.*", os.Getenv("CHANNEL_USERNAME")), &gotgbot.SendMessageOpts{ParseMode: "Markdown"})
	}
	if strings.ToLower(os.Getenv("LOGS")) != "false"{
		logs, log_err := strconv.Atoi(os.Getenv("LOGS"))
		if log_err != nil{fmt.Println(log_err.Error())}
		// Following message is sent in Logs Group (if set)
		ctx.Bot.SendMessage(int64(logs), fmt.Sprintf("#AccClaimed\n\n%v generated a new account.", user.FirstName), nil)
	}
	return nil
}
