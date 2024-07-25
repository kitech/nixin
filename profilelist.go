package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	spjson "github.com/bitly/go-simplejson"
	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

type Profilelist struct {
	*PageBase
}

func NewProfilelist(parent ...qtwidgets.QWidgetITF) *Profilelist {
	me := &Profilelist{}
	me.PageBase = NewPageBase(1, parent...)

	me.setupui2()
	// me.setupmore()
	time.AfterFunc(gopp.DurandMs(100, 200), me.setupmore)
	return me
}

func (me *Profilelist) setupui2() {
	me.cctw.SetColumnCount(5)
	// me.cctw.SetRowCount(3)

	qtrt.Connect(me.opbtn3, "clicked(bool)", func(b bool) {
		log.Println("refresh what???")
	})
}

func (me *Profilelist) setupmore() {

	// todo "--json"
	cmdslc := gopp.Sliceof("nix", "profile", "list", "--json", "--offline")
	lines, err := gopp.RunCmd(".", cmdslc...)
	gopp.ErrPrint(err)

	title := fmt.Sprintf("%v\nProfile History: %v", strings.Join(cmdslc, " "), len(lines))
	qtrt.RunonUithreadx(me.settitle, title)
	// me.ccte.Append("Profile History: " + gopp.ToStr(len(lines)))

	jsonline := gopp.FirstofGv(lines)
	jso, err := spjson.NewJson([]byte(jsonline))
	gopp.ErrPrint(err, gopp.SubStr(jsonline, 32))
	elems := jso.Get("elements")
	var cnter = 0
	var narsztot int64 = 0
	for k, _ := range elems.MustMap() {
		vx := elems.Get(k)
		log.Println(k, vx)

		pathsx := vx.Get("storePaths").Interface()
		path0x := gopp.Firstof(pathsx)
		stpath0 := path0x.Str()

		tbitm0 := qtwidgets.NewQTableWidgetItem(k)
		tbitm1 := qtwidgets.NewQTableWidgetItem(gopp.ToStr(vx.Get("active")))
		tbitm2 := qtwidgets.NewQTableWidgetItem(gopp.ToStr(vx.Get("priority")))
		tbitm3 := qtwidgets.NewQTableWidgetItem(gopp.ToStr(vx.Get("storePaths")))
		cmdline := fmt.Sprintf("nix path-info %v --json --offline", stpath0)
		lines, err := gopp.RunCmd(".", strings.Split(cmdline, " ")...)
		gopp.ErrPrint(err, cmdline)
		jo, err := spjson.NewJson([]byte(gopp.FirstofGv(lines)))
		gopp.ErrPrint(err, lines)
		narsz := jo.Get(stpath0).Get("narSize").MustInt64()
		narsztot += narsz
		tbitm4 := qtwidgets.NewQTableWidgetItem(gopp.Bytes2Humz(narsz))

		row := cnter
		qtrt.RunonUithread(func() {
			me.cctw.InsertRow(row)
			me.cctw.SetItem(row, 0, tbitm0)
			me.cctw.SetItem(row, 1, tbitm1)
			me.cctw.SetItem(row, 2, tbitm2)
			me.cctw.SetItem(row, 4, tbitm3)
			me.cctw.SetItem(row, 3, tbitm4)
		})
		cnter++
	}

	title = fmt.Sprintf("%v\nProfile History: %v, size: %v", strings.Join(cmdslc, " "), cnter, gopp.Bytes2Humz(narsztot))
	qtrt.RunonUithreadx(me.settitle, title)

	if true {
		return
	}

	for i, line := range lines {
		text := fmt.Sprintf("%d : %s", i, line)
		me.ccte.Append(text)
	}

	// todo nix profile <list|history>
}
