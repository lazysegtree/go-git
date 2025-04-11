# Core todos
- [x] Use the go-git module in your own example code.
- [x] Local build - Use go-git with your custom changes
- [x] Understand how logging works.

# Work-stream 
- [ ] Finish Learning Part 1 and Part 2, and Questions Part 1
- [ ] Transport package unit tests
- [ ] Fix PlainInitWithOptions()
- [ ] 

# Learning Part 1
- [ ] Full walkthrough of Clone functionality with logs, etc. 
  - Init -> Create Remote -> Create Client -> Create session 

# Learning Part 2

- [x] Custom logs at some places
- [ ] Perform some plumbing operations using git
- [ ] Understand the various directories and content of .git directory

# Questions part 1
- [ ] What is the concept of work tree ? What is the difference between working tree and what is actually visible to me when I do ls ?
- [ ] "git add command adds file contents to the index" What is index here ?
- [ ] Understand refs, objects, packs, index, worktree, packfile pack protocol, sideband packets ?
- [ ] What is bare clone ?


# Questions part 2
- [ ] Understand this ```
➜  ~/Workspace/kuknitin/go-git git:(v6-transport-lazysegtree) [7:55:43] gitpu
Enumerating objects: 47, done.
Counting objects: 100% (47/47), done.
Delta compression using up to 8 threads
Compressing objects: 100% (31/31), done.
Writing objects: 100% (37/37), 4.83 KiB | 4.83 MiB/s, done.
Total 37 (delta 19), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (19/19), completed with 9 local objects.
remote:
remote: Create a pull request for 'v6-transport-lazysegtree' on GitHub by visiting:
remote:      https://github.com/lazysegtree/go-git/pull/new/v6-transport-lazysegtree
remote:
To https://github.com/lazysegtree/go-git
 * [new branch]        v6-transport-lazysegtree -> v6-transport-lazysegtree
branch 'v6-transport-lazysegtree' set up to track 'origin/v6-transport-lazysegtree'.
➜  ~/Workspace/kuknitin/go-git git:(v6-transport-lazysegtree) [7:55:47]
```

