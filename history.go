package main

import (
	"fmt"
	"time"

	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

type History struct {
	*PageBase
}

func NewHistory(parent ...qtwidgets.QWidgetITF) *History {
	me := &History{}
	me.PageBase = NewPageBase(0, parent...)

	// me.setupmore()
	time.AfterFunc(gopp.DurandMs(100, 200), qtrt.RunonUithreadfn(me.setupmore))
	return me
}

func (me *History) setupmore() {

	lines, err := gopp.RunCmd(".", "nix", "profile", "history", "--offline")
	gopp.ErrPrint(err)

	me.ccte.Append("Profile History: " + gopp.ToStr(len(lines)))

	for i, line := range lines {
		text := fmt.Sprintf("%d : %s", i, line)
		me.ccte.Append(text)
	}

	// todo nix profile <list|history>
}
