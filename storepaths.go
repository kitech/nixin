package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

type Storepaths struct {
	*PageBase
}

func NewStorepaths(parent ...qtwidgets.QWidgetITF) *Storepaths {
	me := &Storepaths{}
	me.PageBase = NewPageBase(0, parent...)

	// me.setupmore()
	time.AfterFunc(gopp.DurandMs(100, 200), me.setupmore)
	return me
}

func (me *Storepaths) setupmore() {
	// nix path-info /nix/store/0253sjdzfspshhhn6zicnvnk9l0hz92x-libffi-3.4.6-dev/ --json --offline
	ets, err := os.ReadDir("/nix/store/")
	gopp.ErrPrint(err)
	var ets2 []fs.DirEntry
	filecnt := 0
	dircnt := 0

	var pkgnames []string
	gopp.Mapdo(ets, func(idx int, etx any) {
		et := etx.(fs.DirEntry)
		if strings.HasSuffix(et.Name(), ".lock") {
			return
		}
		if et.Name() == ".links" {
			return
		}
		if !et.IsDir() {
			// return
		}
		ets2 = append(ets2, et)
		filecnt += gopp.IfElse2(et.IsDir(), 0, 1)
		dircnt += gopp.IfElse2(et.IsDir(), 1, 0)

		log.Println(idx, et.Name())

	})
	line := fmt.Sprintf("Path Info: dircnt: %d filecnt: %d, pkgcnt: %d", dircnt, filecnt, len(pkgnames))
	qtrt.RunonUithread(func() { me.ccte.Append(line) })

	// 太慢了，不能一次读取
	for i, et := range ets2 {
		break
		name := et.Name()
		cmdline := fmt.Sprintf("nix path-info /nix/store/%s --json --offline", name)
		log.Println(cmdline)
		lines, err := gopp.RunCmd(".", strings.Split(cmdline, " ")...)
		gopp.ErrPrint(err)
		_ = lines
		_ = i

		// text := fmt.Sprintf("%d : %s", i, gopp.FirstofGv(lines))
		// me.ccte.Append(text)
		// qtrt.RunonUithread(func() { me.ccte.Append(text) })
	}

	// todo nix profile <list|history>
}
