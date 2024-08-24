package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	spjson "github.com/bitly/go-simplejson"
	"github.com/kitech/gopp"
	_ "github.com/spf13/cobra"

	_ "embed"
)

const nixstdir = "/nix/store"
const nixusrenv = "/nix/store/ia81gjfsjrq2rzf52j4klcw7vxqbdvh3-env-manifest.nix"
const websourl = "https://search.nixos.org/backend/latest-42-nixos-24.05/_search"

var nixpkgs_cachedir = "%s/nixos/nixpkgs-6132b0f6e344ce2fe34fc051b72fb46e34f668e0"

func init() {
	nixpkgs_cachedir = fmt.Sprintf(nixpkgs_cachedir, os.Getenv("HOME"))
	log.SetFlags(log.Flags() ^ log.Ldate ^ log.Ltime)
}

func main() {
	flag.Parse()
	log.Println("args:", flag.Args())

	cmd := flag.Arg(0)
	switch cmd {
	case "sow", "sew": //web
		// var data = gopp.MapSA{"from": 0, "size": 50}
		// gopp.NewHttpClient().Post(websourl).BodyJson()
	case "soc", "sec":
		keyword := flag.Arg(1)
		cachedir := nixpkgs_cachedir + "/pkgs"

		var dotnixs []string
		filepath.WalkDir(cachedir, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				return nil
			}
			if !strings.HasSuffix(d.Name(), ".nix") {
				return nil
			}
			if gopp.StrHaveNocase(d.Name(), keyword) {
				dotnixs = append(dotnixs, path)
			}
			return nil
		})
		// log.Println(len(dotnixs), dotnixs, len(dotnixs))
		log.Println("rc", len(dotnixs))

		vec := gopp.Mapdo(dotnixs, func(idx int, vx any) []any {
			v := vx.(string)
			bname := filepath.Base(v)
			uhome, _ := os.UserHomeDir()
			v2 := gopp.IfElseStr(strings.HasPrefix(v, uhome), "~"+v[len(uhome):], v)
			fmt.Println("")
			log.Println(idx, bname, v2)

			dftnix := v[:len(v)-len(bname)] + "default.nix"
			log.Println(gopp.FileExist(dftnix), dftnix)

			return nil
		})
		log.Println(gopp.Lenof(vec))

	case "so":
		items, err := os.ReadDir(nixstdir)
		gopp.ErrPrint(err)
		// log.Println(items)

		vec := gopp.Mapdo(items, func(vx any) []any {
			v := vx.(fs.DirEntry)
			// log.Println(v)

			match := true
			sekeys := flag.Args()[1:]
			for _, sekey := range sekeys {
				if !gopp.StrHaveNocase(v.Name(), sekey) {
					match = false
					break
				}
			}
			if match {
				return []any{vx}
			}

			return nil
		})
		// log.Println(gopp.Lenof(vec), vec)

		gopp.Mapdo(vec, func(vx any) []any {
			v := vx.(fs.DirEntry)
			log.Println(v)
			return nil
		})

	case "envshow":
		scc := gopp.ReadFileMust(nixusrenv)
		jso, err := spjson.NewJson([]byte(scc))
		gopp.ErrPrint(err)
		gopp.Mapdo(jso.MustArray(), func(vx any) []any {
			log.Println(vx)
			return nil
		})

	case "dlar":
		var hturl = "https://channels.nixos.org/flake-registry.json"
		// https://github.com/NixOS/nixpkgs/archive/6132b0f6e344ce2fe34fc051b72fb46e34f668e0.tar.gz // 40M
		log.Println("DL...", hturl)
		hc := gopp.NewHttpClient()
		resp, err := hc.Get(hturl).Do()
		gopp.ErrPrint(err, hturl)
		bcc, err := io.ReadAll(resp.Body)
		gopp.ErrPrint(err, hturl)
		log.Println(string(bcc))
	}

}

// ////
// 没法手工解析。。。
// 没办法找到包名，版本号。。。
type nixfile struct {
	fpath string
}

func nixfilenew(fpath string) *nixfile {
	me := &nixfile{}
	me.fpath = fpath
	return me
}

func (me *nixfile) parse() *nixfile {
	bcc, err := os.ReadFile(me.fpath)
	gopp.ErrPrint(err, me.fpath)

	lines := strings.Split(string(bcc), "\n")
	log.Println(lines)

	return me
}
