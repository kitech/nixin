package nixcmd

import (
	"strings"

	"github.com/kitech/gopp"
)

type Nix struct {
	Cmd     string
	Subcmds []string
	Options []string
	Args    []string

	short bool
}

const (
	CmdNix = "nix"
	CmdEnv = "nix-env"
	CmdSh  = "nix-shell"
	CmdCh  = "nix-channel"
	CmdSt  = "nix-store"
	CmdGc  = "nix-collect-garbage"
)

func NewNix(cmd string) *Nix {
	me := &Nix{}

	switch cmd {
	case "env":
		me.Cmd = CmdEnv
	case "sh":
		me.Cmd = CmdSh
	case "ch":
		me.Cmd = CmdCh
	case "st":
	case "":
		me.Cmd = CmdNix
	default:
		me.Cmd = cmd
	}

	return me
}

func (me *Nix) Line() string {
	line := strings.Join(me.Lineslc(), " ")
	return line
}
func (me *Nix) Lineslc() (rets []string) {
	rets = append(rets, me.Cmd)
	rets = append(rets, me.Subcmds...)
	rets = append(rets, me.Options...)
	rets = append(rets, me.Args...)
	return
}

func (me *Nix) Putarg(arg string) *Nix {
	me.Options = append(me.Options, arg)
	return me
}
func (me *Nix) Putsubcmd(arg string) *Nix {
	me.Subcmds = append(me.Subcmds, arg)
	return me
}
func (me *Nix) Arg(arg string) *Nix {
	me.Args = append(me.Args, arg)
	return me
}

func (me *Nix) Short() *Nix {
	me.short = true
	return me
}

func (me *Nix) Verbose() *Nix {
	me.Putarg("--verbose")
	me.Putarg("-v")
	return me
}
func (me *Nix) Version() *Nix {
	me.Putarg("--version")
	return me
}
func (me *Nix) Json() *Nix {
	me.Putarg("--json")
	return me
}
func (me *Nix) Offline() *Nix {
	me.Putarg("--offline")
	return me
}
func (me *Nix) Gc() *Nix {
	me.Putarg("--gc")
	return me
}
func (me *Nix) MaxJob(j int) *Nix {
	me.Putarg("--max-jobs")
	me.Putarg(gopp.ToStr(j))
	me.Putarg("-j")
	me.Putarg(gopp.ToStr(j))
	return me
}
func (me *Nix) Update() *Nix {
	me.Putarg("--update")
	return me
}
func (me *Nix) Install() *Nix {
	me.Putarg("--install")
	me.Putarg("-i")
	return me
}
func (me *Nix) Attr() *Nix {
	me.Putarg("--attr")
	me.Putarg("-A")
	return me
}

func (me *Nix) Delold() *Nix {
	me.Putarg("--delete-old")
	me.Putarg("-d")
	return me
}
func (me *Nix) Dryrun() *Nix {
	me.Putarg("--dry-run")
	return me
}

// ///
func (me *Nix) Store() *Nix {
	me.Putsubcmd("store")
	return me
}
func (me *Nix) Profile() *Nix {
	me.Putsubcmd("profile")
	return me
}
func (me *Nix) Pathinfo() *Nix {
	me.Putsubcmd("path-info")
	return me
}
func (me *Nix) List() *Nix {
	me.Putsubcmd("list")
	return me
}
func (me *Nix) Info() *Nix {
	me.Putsubcmd("info")
	return me
}
