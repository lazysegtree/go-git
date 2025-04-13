# To-dos
- [ ] Understand transport package
- [ ] Write come example program using transport package
- [ ] Add logs in transport code and the see the logs in your example program
- [ ] Go through the transport code and try to understand it
- [ ] Find out current code coverage
- [ ] Understand some of the existing tests
- [ ] Write one small unit test
- [ ] Get reviews ?


# Notes
## Transport package 

## Coverage

```bash
(cd plumbing/transport; go test -covermode=count -coverprofile=cover.out  ./...)

# HTML report
(cd plumbing/transport; go test -covermode=count -coverprofile=cover.out  ./...; go tool cover -html=cover.out -o cover.html; open cover.html)
```

```
➜  ~/Workspace/kuknitin/go-git/plumbing/transport git:(v6-transport-lazysegtree) ✗ [2:18:16] go test -covermode=count -coverprofile=cover.out  ./...
	github.com/go-git/go-git/v6/plumbing/transport/ssh/sshagent		coverage: 0.0% of statements
ok  	github.com/go-git/go-git/v6/plumbing/transport	1.444s	coverage: 24.1% of statements
ok  	github.com/go-git/go-git/v6/plumbing/transport/file	2.620s	coverage: 16.3% of statements
ok  	github.com/go-git/go-git/v6/plumbing/transport/git	4.510s	coverage: 3.2% of statements
ok  	github.com/go-git/go-git/v6/plumbing/transport/http	13.650s	coverage: 35.4% of statements
ok  	github.com/go-git/go-git/v6/plumbing/transport/ssh	5.466s	coverage: 71.6% of statements
ok  	github.com/go-git/go-git/v6/plumbing/transport/ssh/knownhosts	5.715s	coverage: 94.2% of statements
➜  ~/Workspace/kuknitin/go-git/plumbing/transport git:(v6-transport-lazysegtree) ✗ [2:19:33]
```