package main

import (
	"bufio"
	"flag"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	var (
		dryrun bool
		remote bool
		ignore string
	)

	flag.BoolVar(&dryrun, "dryrun", false, "run as dryrun")
	flag.BoolVar(&remote, "remote", false, "delete branches on remote, too")
	flag.StringVar(&ignore, "ignore", "", "regexp pattern to ignore")
	flag.Parse()

	branches, err := listMergedBranches()
	if err != nil {
		fmt.Println("failed to list merged branches:", err)
		return
	}

	var ignorePattern *regexp.Regexp
	if ignore != "" {
		ignorePattern, err = regexp.Compile(ignore)
		if err != nil {
			fmt.Printf("invalid regexp=%s: %s\n", ignore, err)
			return
		}
	}

	for _, branch := range branches {
		if ignorePattern != nil && ignorePattern.MatchString(branch) {
			return
		}

		fmt.Println(branch)

		if dryrun {
			continue
		}

		if err := deleteBranch(branch); err != nil {
			fmt.Printf("failed to delete branch=%s: %s\n", branch, err)
		}

		if !remote {
			continue
		}

		if deleteRemoteBranch(branch); err != nil {
			fmt.Printf("failed to delete remote branch=%s: %s\n", branch, err)
		}
	}
}

func listMergedBranches() (branches []string, err error) {
	cmd := exec.Command("git", "branch", "--merged")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}

	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		branch := scanner.Text()

		if strings.HasPrefix(branch, "*") {
			continue
		}

		branches = append(branches, strings.TrimLeft(branch, " "))
	}

	cmd.Wait()

	return
}

func deleteBranch(branch string) error {
	cmd := exec.Command("git", "branch", "--delete", branch)
	return cmd.Run()
}

func deleteRemoteBranch(branch string) error {
	cmd := exec.Command("git", "push", "origin", ":"+branch)
	return cmd.Run()
}
