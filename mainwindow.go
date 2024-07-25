package main

import (
	"log"

	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

type MainWindow struct {
	*qtwidgets.QMainWindow

	welcm *Welcome

	split1 *qtwidgets.QSplitter
	split2 *qtwidgets.QSplitter
	split3 *qtwidgets.QSplitter

	lftw1      *qtwidgets.QWidget
	rgtopstkw1 *qtwidgets.QStackedWidget
	stkcuritm  string

	lovbox1 *qtwidgets.QVBoxLayout // left nav
	lohbox1 *qtwidgets.QVBoxLayout // right top nav

	stkw1 *qtwidgets.QStackedWidget

	tabcw1 *qtwidgets.QTabWidget
	tabcw2 *qtwidgets.QTabWidget
	tab1   *qtwidgets.QWidget
	tab2   *qtwidgets.QWidget
	tab3   *qtwidgets.QWidget

	lftbtnmenus []*qtwidgets.QPushButton

	ccpages map[string]qtwidgets.QWidgetITF
}

func NewMainWindow() *MainWindow {
	me := &MainWindow{}
	me.ccpages = map[string]qtwidgets.QWidgetITF{}
	me.setupui()
	return me
}
func (me *MainWindow) setupui() {
	me.stkcuritm = "welcome"

	me.QMainWindow = qtwidgets.NewQMainWindow(nil, 0)
	me.QMainWindow.Resize(800, 600)
	me.split1 = qtwidgets.NewQSplitter(me.QMainWindow)
	me.SetCentralWidget(me.split1)

	me.lovbox1 = qtwidgets.NewQVBoxLayout(me.split1)

	btnmenudats := []string{
		"welcome", "installed", "history",
		"profilelist", "storepaths", "hehehhe666",
		"hehehhe777", "hehehhe888", "hehehhe999",
		"hehehheaaa", "hehehhebbb", "hehehheccc",
		"hehehheddd",
	}

	for i := 0; i < len(btnmenudats); i += 1 {
		name := btnmenudats[i]

		btn := qtwidgets.NewQPushButtonz0()
		btn.SetObjectName(name)
		btn.SetText(name)

		fn := func(b bool) { me.switchcontentpage(b, btn) }
		qtrt.Connect(btn, "clicked(bool)", fn)

		me.lftbtnmenus = append(me.lftbtnmenus, btn)
		me.lovbox1.AddWidget(btn)
	}

	me.lftw1 = qtwidgets.NewQWidget(nil)
	me.lftw1.SetLayout(me.lovbox1)
	me.split1.AddWidget(me.lftw1)

	/// right
	me.split2 = qtwidgets.NewQSplitter(me)
	me.split2.SetOrientation(3)
	me.split1.AddWidget(me.split2)
	me.split1.SetStretchFactor(1, 99)

	me.rgtopstkw1 = qtwidgets.NewQStackedWidget()
	me.tabcw1 = qtwidgets.NewQTabWidget()

	me.split2.AddWidget(me.rgtopstkw1)
	me.split2.AddWidget(me.tabcw1)
	me.split2.SetStretchFactor(0, 99)

	// top
	me.welcm = NewWelcome()
	me.rgtopstkw1.AddWidget(me.welcm)
	me.ccpages["welcome"] = me.welcm
	// me.rgtopstkw1.SetCurrentIndex(0)

	// bottom
	{
		btn := qtwidgets.NewQPushButtonz0(nil)
		me.tabcw1.AddTab(btn, "tab1")
		btn = qtwidgets.NewQPushButtonz0(nil)
		me.tabcw1.AddTab(btn, "tab2")
		btn = qtwidgets.NewQPushButtonz0(nil)
		me.tabcw1.AddTab(btn, "tab3")
	}
}

func (me *MainWindow) switchcontentpage(b bool, btn *qtwidgets.QPushButton) {
	objname := btn.ObjectName()
	log.Println(objname, b)

	pagex, ok := me.ccpages[objname]
	if ok {
		idx := me.rgtopstkw1.IndexOf(pagex)
		curidx := me.rgtopstkw1.CurrentIndex()
		if idx != curidx {
			me.rgtopstkw1.SetCurrentWidget(pagex)
		}
		return
	}
	switch objname {
	case "welcome":
	case "installed":
		page := NewInstalled()
		idx := me.rgtopstkw1.AddWidget(page)
		me.rgtopstkw1.SetCurrentIndex(idx)
		me.ccpages[objname] = page
	case "history":
		page := NewHistory()
		idx := me.rgtopstkw1.AddWidget(page)
		me.rgtopstkw1.SetCurrentIndex(idx)
		me.ccpages[objname] = page
	case "profilelist":
		page := NewProfilelist()
		idx := me.rgtopstkw1.AddWidget(page)
		me.rgtopstkw1.SetCurrentIndex(idx)
		me.ccpages[objname] = page
	case "storepaths":
		page := NewStorepaths()
		idx := me.rgtopstkw1.AddWidget(page)
		me.rgtopstkw1.SetCurrentIndex(idx)
		me.ccpages[objname] = page

	}
}
