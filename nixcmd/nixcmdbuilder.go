package nixcmd

import (
	"strings"
)

type Nix struct {
	Cmd     string
	Subcmds []string
	Options []string
}

const (
	CmdNix = "nix"
	CmdEnv = "nix-env"
	CmdSh  = "nix-shell"
	CmdCh  = "nix-channel"
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
	case "":
		fallthrough
	default:
		me.Cmd = CmdNix
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

func (me *Nix) Verbose() *Nix {
	me.Putarg("--verbose")
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
