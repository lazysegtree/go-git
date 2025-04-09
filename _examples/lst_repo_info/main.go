package main

import (
	"os"
)
import "github.com/go-git/go-git/v5"
import . "github.com/go-git/go-git/v5/_examples"

func main() {
	CheckArgs("<url>")
	url := os.Args[1]
	dir := "/tmp/lst_test/list_commits"
	os.RemoveAll(dir)
	os.MkdirAll(dir, 0755)
	//fs := memfs.New()
	//storer := memory.NewStorage()
	//Info("cloning")
	//_, err := git.Clone(storer, fs, &git.CloneOptions{URL: url})
	//
	//CheckIfError(err)

	repo, err := git.PlainClone(dir, false, &git.CloneOptions{URL: url})
	CheckIfError(err)

	remotes, err := repo.Remotes()
	CheckIfError(err)
	Info("remotes count : %d", len(remotes))

	/*citer, err := repo.CommitObjects()
	CheckIfError(err)

	citer.ForEach(func(c *object.Commit) error {
		Info("Commit %s, authored by %s, committed by %s on %v", c.Hash.String(),
			c.Author.Name,
			c.Committer.Name, c.Committer.When)
		return nil
	})

	biter, err := repo.BlobObjects()
	CheckIfError(err)
	biter.ForEach(func(b *object.Blob) error {
		Info("Blob of type %v, size : %d", b.Type(), b.Size)
		return nil
	})*/
}
