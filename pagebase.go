package main

import (
	"log"

	"github.com/kitech/gopp"
	"github.com/qtui/qtwidgets"
)

type PageBase struct {
	*qtwidgets.QWidget

	/////
	ccte *qtwidgets.QTextEdit
}

func NewPageBase(parent ...qtwidgets.QWidgetITF) *PageBase {
	me := &PageBase{}
	me.QWidget = qtwidgets.NewQWidget(gopp.FirstofGv(parent))

	me.setupui()
	// me.setupmore()
	// time.AfterFunc(gopp.DurandMs(100, 200), qtrt.RunonUithreadfn(me.setupmore))
	return me
}

func (me *PageBase) setupui() {

	hbox := qtwidgets.NewQHBoxLayout(nil)
	btn := qtwidgets.NewQPushButtonz0(nil)
	hbox.AddWidget(btn)
	btn = qtwidgets.NewQPushButtonz0(nil)
	hbox.AddWidget(btn)
	btn = qtwidgets.NewQPushButtonz0(nil)
	hbox.AddWidget(btn)

	vbox := qtwidgets.NewQVBoxLayout(nil)
	vbox.AddLayout(hbox)

	te := qtwidgets.NewQTextEdit()
	te.SetReadOnly(true)
	te.SetText("*Profile History\n")
	vbox.AddWidget(te)
	me.ccte = te

	spc := qtwidgets.NewQSpacerItem(40, 20, 0, qtwidgets.QSizePolicy__Minimum)
	vbox.AddItem(spc)

	log.Println(vbox, hbox)
	me.SetLayout(vbox)
}

// func (me *PageBase) setupmore() {

// 	lines, err := gopp.RunCmd(".", "nix", "profile", "history")
// 	gopp.ErrPrint(err)

// 	me.ccte.Append("Profile History: " + gopp.ToStr(len(lines)))

// 	for i, line := range lines {
// 		text := fmt.Sprintf("%d : %s", i, line)
// 		me.ccte.Append(text)
// 	}

// 	// todo nix profile <list|history>
// }
