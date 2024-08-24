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
	"time"

	spjson "github.com/bitly/go-simplejson"
	"github.com/kitech/gopp"
	_ "github.com/spf13/cobra"

	_ "embed"
)

const nixstdir = "/nix/store"
const nixusrenv = "/nix/store/ia81gjfsjrq2rzf52j4klcw7vxqbdvh3-env-manifest.nix"
const websourl = "https://search.nixos.org/backend/latest-42-nixos-24.05/_search"

var nixpkgs_cachedir = os.Getenv("HOME") + "/nixos/nixpkgs-6132b0f6e344ce2fe34fc051b72fb46e34f668e0"

//go:embed websereq.json
var websereqtmpl string

func init() {
	// nixpkgs_cachedir = fmt.Sprintf(nixpkgs_cachedir, os.Getenv("HOME"))
	log.SetFlags(log.Flags() ^ log.Ldate ^ log.Ltime)
}

func main() {
	flag.Parse()
	log.Println("args:", flag.Args())
	btime := time.Now()
	defer func() { log.Println("Used:", time.Since(btime)) }()

	cmd := flag.Arg(0)
	switch cmd {
	case "sow", "sew": // search web
		// var data = gopp.MapSA{"from": 0, "size": 50}
		// gopp.NewHttpClient().Post(websourl).BodyJson()

		words := flag.Arg(1)
		reqdata := websereqtmpl
		reqdata = strings.ReplaceAll(reqdata, "\"size\": 50,", "\"size\": 20,")
		reqdata = strings.ReplaceAll(reqdata, "\"aerc\"", fmt.Sprintf("\"%s\"", words))
		reqdata = strings.ReplaceAll(reqdata, "\"multi_match_aerc\"", fmt.Sprintf("\"multi_match_%s\"", words))
		reqdata = strings.ReplaceAll(reqdata, "\"*Aerc*\"", fmt.Sprintf("\"*%s*\"", gopp.Title(words)))
		reqdata = strings.ReplaceAll(reqdata, "\"*aerc*\"", fmt.Sprintf("\"*%s*\"", (words)))

		_, err := spjson.NewJson([]byte(reqdata))
		gopp.ErrPrint(err, reqdata)
		// log.Println(reqdata)

		htcli := gopp.NewHttpClient().HeaderKV("referer", "https://search.nixos.org/packages").Post(websourl).BodyRaw([]byte(reqdata))
		htcli.HeaderKV("Authorization", "Basic YVdWU0FMWHBadjpYOGdQSG56TDUyd0ZFZWt1eHNmUTljU2g=")
		htcli.HeaderKV("content-type", gopp.HttpCTJson)
		resp, err := htcli.Do()
		gopp.ErrPrint(err, resp == nil)
		if err != nil {
			break
		}
		defer resp.Body.Close()

		bcc, err := io.ReadAll(resp.Body)
		gopp.ErrPrint(err)
		if resp.StatusCode >= 400 {
			log.Println(resp.Status, websourl, string(bcc))
		}

		jso, err := spjson.NewJson(bcc)
		gopp.ErrPrint(err, jso == nil)
		// log.Println(jso)
		hitsx := jso.GetPath("hits", "hits")
		for i := 0; i < len(hitsx.MustArray()); i++ {
			ox := hitsx.GetIndex(i).Get("_source")
			// log.Println(i, ox)
			// log.Println(i, ox.Get("package_attr_name"), ox.Get("package_pversion"), ox.Get("package_description"), ox.Get("package_homepage"))
			// 没有更新时间还感觉缺少点啥
			pkgname := ox.Get("package_attr_name").MustString()
			pkgver := ox.Get("package_pversion").MustString()
			pkgdesc := ox.Get("package_description").MustString()
			pkgurl := ox.Get("package_homepage").Interface()

			log.Println(i, pkgname, pkgver, pkgdesc)
			// log.Println(i, pkgdesc)
			log.Println(i, strings.Repeat(" ", 4), pkgurl)
			log.Println()
		}

	case "soc", "sec": // search nixpkgs full cache
		// search works.nix in metadb ~/nixos/nixpkgs-*/pkgs
		// little slow, about 5s

		keyword := flag.Arg(1)
		cachedir := nixpkgs_cachedir + "/pkgs"

		var secnter = 0
		var dotnixs []string
		filepath.WalkDir(cachedir, func(path string, d fs.DirEntry, err error) error {
			secnter++
			fmt.Printf("%4d: %v%v\r", secnter, d.Name(), strings.Repeat(" ", 26))
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
		log.Println("rc", len(dotnixs), secnter)

		vec := gopp.Mapdo(dotnixs, func(idx int, vx any) []any {
			v := vx.(string)
			bname := filepath.Base(v)
			uhome, _ := os.UserHomeDir()
			v2 := gopp.IfElseStr(strings.HasPrefix(v, uhome), "~"+v[len(uhome):], v)
			fmt.Println("")
			log.Println(idx, bname, v2)

			dftnix := v[:len(v)-len(bname)] + "default.nix"
			dftnix = strings.Replace(dftnix, uhome, "~", 1)
			log.Println(gopp.FileExist(dftnix), dftnix)

			return nil
		})
		log.Println(gopp.Lenof(vec))

	case "so": // search /nix/store dir, already installed
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
			log.Println(v.Type(), nixstdir+"/"+v.Name())
			return nil
		})

	case "envshow":
		files, err := filepath.Glob("/nix/store/*-env-manifest.nix")
		gopp.ErrPrint(err, files == nil)
		gopp.Mapdo(files, func(idx int, vx any) []any {
			val := vx.(string)
			log.Println("Reading...", idx, val)
			scc := gopp.ReadFileMust(val)
			jso, err := spjson.NewJson([]byte(scc)) // todo, its not json indeed
			gopp.ErrPrint(err, scc)

			gopp.Mapdo(jso.MustArray(), func(vx any) []any {
				log.Println(vx)
				return nil
			})

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

	default:
		log.Println("so, soc, sow, dlar, envshow")
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
