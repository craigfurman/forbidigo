package main

import (
	"flag"
	"go/ast"
	"log"
	"os"

	"github.com/ashanbrown/disallow/disallow"
	"golang.org/x/tools/go/packages"
)

func main() {
	log.SetFlags(0) // remove log timestamp

	setExitStatus := flag.Bool("set_exit_status", false, "Set exit status to 1 if any issues are found")
	flag.Parse()

	patterns := disallow.DefaultPatterns()

	firstPkg := 0
	for n, arg := range flag.Args() {
		if arg == "--" {
			firstPkg = n + 1
			break
		}
		patterns = append(patterns, arg)
	}

	cfg := packages.Config{
		Mode: packages.NeedSyntax | packages.NeedName | packages.NeedFiles | packages.NeedTypes,
	}
	pkgs, err := packages.Load(&cfg, flag.Args()[firstPkg:]...)
	if err != nil {
		log.Fatalf("Could not load packages: %s", err)
	}
	linter, err := disallow.NewLinter(patterns...)
	if err != nil {
		log.Fatalf("Could not create linter: %s", err)
	}

	var issues []disallow.Issue // nolint:prealloc // we don't know how many there will be
	for _, p := range pkgs {
		nodes := make([]ast.Node, 0, len(p.Syntax))
		for _, n := range p.Syntax {
			nodes = append(nodes, n)
		}
		newIssues, err := linter.Run(p.Fset, nodes...)
		if err != nil {
			log.Fatalf("failed: %s", err)
		}
		issues = append(issues, newIssues...)
	}

	for _, issue := range issues {
		log.Println(issue)
	}

	if *setExitStatus && len(issues) > 0 {
		os.Exit(1)
	}
}