package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/kitech/gopp"
	"github.com/qtui/qtrt"
	"github.com/qtui/qtwidgets"
)

type Welcome struct {
	*PageBase
}

func NewWelcome(parent ...qtwidgets.QWidgetITF) *Welcome {
	me := &Welcome{}
	me.PageBase = NewPageBase(0, parent...)
	// me.setupmore()
	time.AfterFunc(gopp.DurandMs(100, 200), qtrt.RunonUithreadfn(me.setupmore))
	return me
}

func (me *Welcome) setupmore() {
	me.ccte.Append("Install Mode: nixpkgs")

	lines, err := gopp.RunCmd(".", "nix-env", "--version")
	gopp.ErrPrint(err)
	log.Println(lines)
	me.ccte.Append("Version: " + strings.Join(lines, "\n"))

	lines, err = gopp.RunCmd(".", "nix", "store", "info", "--offline")
	gopp.ErrPrint(err)
	me.ccte.Append("Version: " + strings.Join(lines, "\n"))

	{
		lines, err := gopp.RunCmd(".", "which", "nix-env")
		gopp.ErrPrint(err)
		cmdfile := gopp.FirstofGv(lines)
		pdir := filepath.Dir(filepath.Dir(cmdfile))
		instdir := "/" + gopp.FirstofGv(strings.Split(pdir, "/")[1:])

		me.ccte.Append("Install dir: " + instdir)
		me.ccte.Append("Profile dir: " + pdir)

		// manifile := pdir + "/manifest.nix"

	}

	{
		lines, err := gopp.RunCmd(".", "nix-env", "-q")
		gopp.ErrPrint(err)

		me.ccte.Append("Installed nixpkgs: " + gopp.ToStr(len(lines)))
	}

	{
		lines, err := gopp.RunCmd(".", "nix-env", "--list-generations")
		gopp.ErrPrint(err)

		me.ccte.Append("Total generations: " + gopp.ToStr(len(lines)))
	}

	{
		me.ccte.Append("")
		envnames := gopp.Sliceof("NIX_STATE_DIR", "IN_NIX_SHELL", "NIX_PATH", "NIX_IGNORE_SYMLINK_STORE", "NIX_STORE_DIR", "NIX_DATA_DIR", "NIX_LOG_DIR", "NIX_CONF_DIR", "NIX_CONFIG", "NIX_USER_CONF_FILES", "TMPDIR", "NIX_REMOTE", "NIX_SHOW_STATS", "NIX_COUNT_CALLS", "GC_INITIAL_HEAP_SIZE")
		for _, envname := range envnames {
			envval := os.Getenv(envname)
			me.ccte.Append(envname + ": " + envval)
		}
	}

	// todo nix profile <list|history>
}
