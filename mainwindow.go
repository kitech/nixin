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

	lovbox1 *qtwidgets.QVBoxLayout // left nav
	lohbox1 *qtwidgets.QVBoxLayout // right top nav

	stkw1 *qtwidgets.QStackedWidget

	tabcw1 *qtwidgets.QTabWidget
	tabcw2 *qtwidgets.QTabWidget
	tab1   *qtwidgets.QWidget
	tab2   *qtwidgets.QWidget
	tab3   *qtwidgets.QWidget

	lftbtnmenus []*qtwidgets.QPushButton
}

func NewMainWindow() *MainWindow {
	me := &MainWindow{}
	me.setupui()
	return me
}
func (me *MainWindow) setupui() {
	me.QMainWindow = qtwidgets.NewQMainWindow(nil, 0)
	me.split1 = qtwidgets.NewQSplitter(me.QMainWindow)
	me.SetCentralWidget(me.split1)

	me.lovbox1 = qtwidgets.NewQVBoxLayout(me.split1)

	{
		btn := qtwidgets.NewQPushButton("hehee111")
		qtrt.Connect(btn, "clicked(bool)", func(b bool) {
			log.Println("works???", b)
		})
		// btn.SetFlat(true)
		// btn.SetText("hehhe111")
		me.lovbox1.AddWidget(btn)
	}
	{
		btn := qtwidgets.NewQPushButtonz0()
		qtrt.Connect(btn, "clicked(bool)", func(b bool) {
			log.Println("works???222", b)
		})
		// btn.SetFlat(true)
		btn.SetText("hehhe222")
		me.lovbox1.AddWidget(btn)
	}

	btnmenudats := []any{
		"hehehhe333", func(b bool) { log.Println("hehehhe333") },
		"hehehhe444", func(b bool) { log.Println("hehehhe444") },
		"hehehhe555", func(b bool) { log.Println("hehehhe555") },
		"hehehhe666", func(b bool) { log.Println("hehehhe666") },
		"hehehhe777", func(b bool) { log.Println("hehehhe777") },
		"hehehhe888", func(b bool) { log.Println("hehehhe888") },
		"hehehhe999", func(b bool) { log.Println("hehehhe999") },
		"hehehheaaa", func(b bool) { log.Println("hehehheaaa") },
		"hehehhebbb", func(b bool) { log.Println("hehehhebbb") },
		"hehehheccc", func(b bool) { log.Println("hehehheccc") },
		"hehehheddd", func(b bool) { log.Println("hehehheddd") },
	}

	for i := 0; i < len(btnmenudats); i += 2 {
		name := btnmenudats[i].(string)
		fn := btnmenudats[i+1].(func(bool))

		btn := qtwidgets.NewQPushButtonz0()
		btn.SetText(name)
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
