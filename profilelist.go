package main

import (
	"fmt"
	"time"

	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

type Profilelist struct {
	*PageBase
}

func NewProfilelist(parent ...qtwidgets.QWidgetITF) *Profilelist {
	me := &Profilelist{}
	me.PageBase = NewPageBase(parent...)

	// me.setupmore()
	time.AfterFunc(gopp.DurandMs(100, 200), qtrt.RunonUithreadfn(me.setupmore))
	return me
}

func (me *Profilelist) setupmore() {

	// todo "--json"
	lines, err := gopp.RunCmd(".", "nix", "profile", "list", "--offline")
	gopp.ErrPrint(err)

	me.ccte.Append("Profile History: " + gopp.ToStr(len(lines)))

	for i, line := range lines {
		text := fmt.Sprintf("%d : %s", i, line)
		me.ccte.Append(text)
	}

	// todo nix profile <list|history>
}
