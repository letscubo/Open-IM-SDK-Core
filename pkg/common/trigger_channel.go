package common

import (
	"errors"
	"open_im_sdk/pkg/constant"
	"open_im_sdk/pkg/log"
	"open_im_sdk/pkg/utils"
	"time"
)

func TriggerCmdNewMsgCome(msg utils.ArrMsg, conversationCh chan Cmd2Value) error {
	c2v := Cmd2Value{Cmd: constant.CmdNewMsgCome, Value: msg}
	return sendCmd(conversationCh, c2v, 1)
}
func TriggerCmdDeleteConversationAndMessage(sourceID, conversationID string, sessionType int, conversationCh chan Cmd2Value) error {
	c2v := Cmd2Value{
		Cmd:   constant.CmdDeleteConversation,
		Value: deleteConNode{SourceID: sourceID, ConversationID: conversationID, SessionType: sessionType},
	}

	return sendCmd(conversationCh, c2v, 1)
}
func TriggerCmdUpdateConversation(node updateConNode, conversationCh chan Cmd2Value) error {
	c2v := Cmd2Value{
		Cmd:   constant.CmdUpdateConversation,
		Value: node,
	}

	return sendCmd(conversationCh, c2v, 1)
}

type deleteConNode struct {
	SourceID       string
	ConversationID string
	SessionType    int
}
type updateConNode struct {
	ConId  string
	Action int //1 Delete the conversation; 2 Update the latest news in the conversation or add a conversation; 3 Put a conversation on the top;
	// 4 Cancel a conversation on the top, 5 Messages are not read and set to 0, 6 New conversations
	Args interface{}
}

type Cmd2Value struct {
	Cmd   string
	Value interface{}
}

func unInitAll(conversationCh chan Cmd2Value) {
	c2v := Cmd2Value{Cmd: constant.CmdUnInit}
	_ = sendCmd(conversationCh, c2v, 1)
}

type goroutine interface {
	work(cmd Cmd2Value)
	getCh() chan Cmd2Value
}

func doListener(Li goroutine) {
	log.Info("internal", "doListener start.", Li.getCh())
	for {
		log.Info("doListener for.")
		select {
		case cmd := <-Li.getCh():
			if cmd.Cmd == constant.CmdUnInit {
				log.Info("doListener goroutine.")
				return
			}
			log.Info("doListener work.")
			Li.work(cmd)
		}
	}
}

func sendCmd(ch chan Cmd2Value, value Cmd2Value, timeout int64) error {
	var flag = 0
	select {
	case ch <- value:
		flag = 1
	case <-time.After(time.Second * time.Duration(timeout)):
		flag = 2
	}
	if flag == 1 {
		return nil
	} else {
		return errors.New("send cmd timeout")
	}
}
