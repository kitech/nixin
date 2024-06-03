package main

import (
	"flag"
	"io"
	"io/fs"
	"log"
	"os"

	spjson "github.com/bitly/go-simplejson"
	"github.com/kitech/gopp"
	_ "github.com/spf13/cobra"
)

const nixstdir = "/nix/store"
const nixusrenv = "/nix/store/ia81gjfsjrq2rzf52j4klcw7vxqbdvh3-env-manifest.nix"

func main() {
	flag.Parse()
	cmd := flag.Arg(0)
	log.Println(flag.Args())
	switch cmd {
	case "se":
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
