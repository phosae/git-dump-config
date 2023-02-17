package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

func main() {
	// github public repo example
	dumpToConfig("https://github.com/kubernetes-sigs/kind", "hack/build")
	// gitlab auth with token example
	token := "<your-access-token-hear>" // scopes: api, read_repository(not read_registry)
	cdeploys := fmt.Sprintf("https://oauth2:%s@gitlab.qiniu.io/qbox/c-deploy.git", token)
	dumpToConfig(cdeploys, "callisto/templates/dora/base/service-gate")
}

func dumpToConfig(repoURL, path string) {
	fmt.Println("start clone github repo:", repoURL)
	repo, err := clone(repoURL)
	if err != nil {
		panic(err)
	}
	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	fmt.Println("git head ref:", ref)
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		panic(err)
	}

	tree, err := commit.Tree()
	if err != nil {
		panic(err)
	}

	var configs = make(map[string]string, 32)

	err = dump(path, tree.Files(), configs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("dump git files(in '%s') success:\n", path)
	for name := range configs {
		fmt.Printf("\t%s\n", name)
	}
}

func dump(dir string, files *object.FileIter, to map[string]string) error {
	var addFile = func(f *object.File) error {
		content, err := f.Contents()
		if err != nil {
			return err
		}
		to[path.Base(f.Name)] = content
		return nil
	}
	return files.ForEach(func(f *object.File) (err error) {
		switch dir {
		case ".":
			fallthrough
		case "":
			err = addFile(f)
		default:
			if !strings.HasPrefix(f.Name, dir) {
				return nil
			}
			err = addFile(f)
		}
		return
	})
}

// clone a git Repository in memory
// branch selection is not supported now, can be added similar to https://github.com/src-d/go-git/issues/777
func clone(url string) (*git.Repository, error) {
	// clone to dir
	// repo, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return r, err
}
