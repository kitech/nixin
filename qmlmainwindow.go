package main

import (
	"log"

	"github.com/qtui/qtqml"
)

var qmlape *qtqml.QQmlApplicationEngine

func qmlmain() {
	qmlape = qtqml.NewQQmlApplicationEngine(nil)
	qmlape.Load("qmlmain.qml")
	// todo load done callback

	robjx := qmlape.RootObject()
	romo := robjx.MetaObject()
	log.Println(robjx)
	log.Println(robjx.Dbgstr())
	qvar := robjx.Property("objectName")
	log.Println(qvar)
	log.Println(qvar.ToString())

	itm := qtqml.NewQQuickItem()
	itm.Dtor()

	qvar = robjx.Property("contentItem")
	log.Println(qvar)
	log.Println(qvar.ToLongLong())
	log.Println(romo.ClassName())
	supmo := romo.SuperClass()
	log.Println(supmo.ClassName())

	robj := qtqml.QQuickApplicationWindowFromptr(robjx.GetCthis())
	log.Println(robj.ContentItem())
	itm = robj.ContentItem()

	ta := qtqml.NewQQuickTextArea(itm)
	log.Println(ta)
	ta.SetProperty("visible", true)
	ta.SetWidth(200)
	ta.SetHeight(200)
	ta.SetProperty("color", "white")
	ta.SetProperty("text", "whitealiefieaf瀑 fewer iaefjoeff粗盐")
}
