package bootstrap

import (
	"code.rongyi.io/LER0ever/EverVim/utils"
	"github.com/dixonwille/wmenu"
	log "github.com/sirupsen/logrus"
)

type AutoCompletionPlugin int

const (
	_ AutoCompletionPlugin = iota - 1
	YouCompleteMe
	Deoplete
	NeoComplete
	LSP
)

func askForAutoCompletion() int {
	completer := utils.MultipleChoice(
		"Which AutoCompletion Plugin do you prefer",
		[]string{
			"YouCompleteMe",
			"Deoplete",
			"NeoComplete",
			"Language Server Protocol",
		},
	)
	log.Debugf("User choosed %d", completer)
	return completer
}

func askForAutoCompletionOld() int {
	choice := -1
	menu := wmenu.NewMenu("Which AutoCompletion Plugin do you prefer?")
	menu.Action(func(opts []wmenu.Opt) error {
		log.Infof("AutoComplete: %s\n", opts[0].Text)
		choice = int(opts[0].Value.(AutoCompletionPlugin))
		return nil
	})
	menu.Option("YouCompleteMe", YouCompleteMe, true, nil)
	menu.Option("Deoplete", Deoplete, false, nil)
	menu.Option("NeoComplete", NeoComplete, false, nil)
	menu.Option("Language Server Protocol", LSP, false, nil)
	err := menu.Run()
	if err != nil {
		log.Fatalf("Error running menu: %s\n", err.Error())
		return -1
	}
	log.Debugf("Choice is %d\n", choice)
	return choice
}
