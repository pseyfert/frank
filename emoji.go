package main

import (
	"fmt"
	emojispeller "github.com/pseyfert/go-terminal-speller/pkg"
	"gopkg.in/sorcix/irc.v2"
	"strings"
)

func emojiSpellout(parsed *irc.Message) error {
	tgt := Target(parsed)
	msg := parsed.Trailing()
	if !strings.HasPrefix(tgt, "#") {
		// only answer to this in channels
		return nil
	}

	translation, err := emojispeller.StringForceTranslate(msg)
	if err == nil {
		Privmsg(tgt, fmt.Sprintf("[DEmojified] %s", translation))
	}

	return nil
}
