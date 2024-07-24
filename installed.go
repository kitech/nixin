package main

import (
	"fmt"
	"time"

	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

type Installed struct {
	*PageBase
}

func NewInstalled(parent ...qtwidgets.QWidgetITF) *Installed {
	me := &Installed{}
	me.PageBase = NewPageBase(parent...)

	// me.setupmore()
	time.AfterFunc(gopp.DurandMs(100, 200), qtrt.RunonUithreadfn(me.setupmore))
	return me
}

func (me *Installed) setupmore() {

	lines, err := gopp.RunCmd(".", "nix-env", "-q")
	gopp.ErrPrint(err)

	me.ccte.Append("Installed nixpkgs: " + gopp.ToStr(len(lines)))

	for i, line := range lines {
		text := fmt.Sprintf("%d : %s", i, line)
		me.ccte.Append(text)
	}

	// todo nix profile <list|history>
}
