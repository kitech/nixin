package main

import (
	"log"

	"github.com/kitech/gopp"
	"github.com/qtui/qtwidgets"
)

type Welcome struct {
	*qtwidgets.QWidget

	/////
}

func NewWelcome(parent ...qtwidgets.QWidgetITF) *Welcome {
	me := &Welcome{}
	me.QWidget = qtwidgets.NewQWidget(gopp.FirstofGv(parent))

	me.setupui()
	return me
}

func (me *Welcome) setupui() {

	hbox := qtwidgets.NewQHBoxLayout(nil)
	btn := qtwidgets.NewQPushButtonz0(nil)
	hbox.AddWidget(btn)
	btn = qtwidgets.NewQPushButtonz0(nil)
	hbox.AddWidget(btn)
	btn = qtwidgets.NewQPushButtonz0(nil)
	hbox.AddWidget(btn)

	vbox := qtwidgets.NewQVBoxLayout(nil)
	vbox.AddLayout(hbox)

	spc := qtwidgets.NewQSpacerItem(40, 20, 0, 1|2|4)
	vbox.AddItem(spc)

	log.Println(vbox, hbox)
	me.SetLayout(vbox)
}
