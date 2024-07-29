package main

import (
	"fmt"
	"log"

	"github.com/qtui/qtqml"
	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

var qmlape *qtqml.QQmlApplicationEngine

func qmlmain() {
	// 可能是要尽快初始化QQmlApplicationEngine，在QApplication之后，
	qapp := qtwidgets.QApp()
	qmlape = qtqml.NewQQmlApplicationEngine(qapp)
	qtrt.Connect(qmlape, "objectCreated", func(objx voidptr, urlx voidptr) {
		log.Println(".qml loaded", objx, urlx)
		// obj := qtcore.QObjectFromptr(objx)
		// url := qtcore.QUrlFromptr(urlx)
		// log.Println(".qml loaded", obj.Dbgstr(), url)

		qtrt.RunonUithread(qmluiloaded)
		// qmluiloaded()
	})
	log.Println(qmlape.Dbgstr())
	qmlape.Load("qmlmain.qml")
	// todo load done callback
	// qmlape.Dtor()
	// gopp.PauseAk()
}
func qmluiloaded() {
	robjx := qmlape.RootObject()
	romo := robjx.MetaObject()
	log.Println(robjx)
	log.Println(robjx.Dbgstr())
	qvar := robjx.Property("objectName")
	log.Println(qvar)
	log.Println(qvar.ToString())

	{
		itm := qtqml.NewQQuickItem()
		itm.Dtor()
	}

	qvar = robjx.Property("contentItem")
	log.Println(qvar)
	log.Println(qvar.ToLongLong())
	log.Println(romo.ClassName())
	supmo := romo.SuperClass()
	log.Println(supmo.ClassName())

	robj := qtqml.QQuickApplicationWindowFromptr(robjx.GetCthis())
	log.Println(robj.ContentItem())

	myw := qtqml.NewQQuickApplicationWindow()
	log.Println(myw.Dbgstr())

	// 效果不一样，全都堆在一起了
	obj := robj.FindChild("vertlo")
	log.Println(obj.Dbgstr())
	lbp := qtqml.QQuickItemFromptr(obj.GetCthis())

	log.Println("=======")
	for i := 0; i < 3; i++ {
		lb := qtqml.NewQQuickLabel(lbp)
		// lb.sette
		lb.SetText(fmt.Sprintf("hehe%d", i))
		lb.SetProperty("color", "red")

		// ta := qtqml.NewQQuickTextArea(vlo)
		// // gopp.PauseAk()
		// log.Println(ta)
		// ta.SetProperty("visible", true)
		// ta.SetWidth(100)
		// ta.SetHeight(100)
		// ta.SetProperty("color", "white")
		// ta.SetProperty("text", "whitealiefieaf瀑 fewer iaefjoeff粗盐")
	}
	log.Println(lbp.ChildItems().Size())
	// gopp.PauseAk()
}
