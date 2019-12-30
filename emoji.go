package main

import (
	"bytes"
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

	var b bytes.Buffer

	translator := emojispeller.NewTranslator(&b)
	translator.Write([]byte(msg))

	if translator.Didsomething {
		Privmsg(tgt, fmt.Sprintf("[DEmojified] %s", b.String()))
	}

	return nil
}
