package main

import (
	"log"

	"github.com/kitech/gopp"
	"github.com/qtui/qtcore"
	"github.com/qtui/qtwidgets"
)

type PageBase struct {
	*qtwidgets.QWidget

	/////
	ccte    *qtwidgets.QTextEdit
	cctw    *qtwidgets.QTableWidget
	mode    int
	lbtitle *qtwidgets.QLabel
	opbtn1  *qtwidgets.QPushButton
	opbtn2  *qtwidgets.QPushButton
	opbtn3  *qtwidgets.QPushButton
}

func NewPageBase(mode int, parent ...qtwidgets.QWidgetITF) *PageBase {
	me := &PageBase{}
	me.mode = mode
	me.QWidget = qtwidgets.NewQWidget(gopp.FirstofGv(parent))

	me.setupui()
	// me.setupmore()
	// time.AfterFunc(gopp.DurandMs(100, 200), qtrt.RunonUithreadfn(me.setupmore))
	return me
}

func (me *PageBase) setupui() {

	hbox := qtwidgets.NewQHBoxLayout(nil)
	btn := qtwidgets.NewQPushButtonz0(nil)
	me.opbtn1 = btn
	hbox.AddWidget(btn)
	btn = qtwidgets.NewQPushButtonz0(nil)
	me.opbtn2 = btn
	hbox.AddWidget(btn)
	btn = qtwidgets.NewQPushButtonz0(nil)
	me.opbtn3 = btn
	btn.SetText("refresh")
	// qtrt.Connect(btn, "clicked(bool)", func(b bool) {
	// 	log.Println("refresh what???")
	// })
	hbox.AddWidget(btn)

	vbox := qtwidgets.NewQVBoxLayout(nil)
	vbox.AddLayout(hbox)

	lb := qtwidgets.NewQLabel("hehhee", nil)
	vbox.AddWidget(lb)
	me.lbtitle = lb
	lb.SetTextInteractionFlags(qtcore.Qt__TextSelectableByKeyboard | qtcore.Qt__TextSelectableByMouse)

	if me.mode == 1 {
		tw := qtwidgets.NewQTableWidget()
		vbox.AddWidget(tw)
		me.cctw = tw
	} else {
		te := qtwidgets.NewQTextEdit()
		te.SetReadOnly(true)
		te.SetText("*Profile History\n")
		vbox.AddWidget(te)
		me.ccte = te
	}

	spc := qtwidgets.NewQSpacerItem(40, 20, 0, qtwidgets.QSizePolicy__Minimum)
	vbox.AddItem(spc)

	log.Println(vbox, hbox)
	me.SetLayout(vbox)
}
func (me *PageBase) settitle(title string) {
	me.lbtitle.SetText(title)
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
