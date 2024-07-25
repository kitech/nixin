package nixcmd

import (
	"log"
	"testing"
)

func TestNCB1(t *testing.T) {
	nix := NewNix("")
	nix.Offline().Json()
	out := nix.Line()
	log.Println(out, nix.Lineslc())
}
