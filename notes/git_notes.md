# Command
## Object exploring
```
# Show the content of a blob, tree, or commit object
git cat-file -p <object-hash>

# Show the type of an object
git cat-file -t <object-hash>

# Show the size of an object
git cat-file -s <object-hash>
```

Create custom object
```
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:04:34] echo "Hello, Git internals" | git hash-object -w --stdin

268556b6603a1a3ffe8db19a5f47975c8c8e1de0
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:15:27]

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:15:34] ls .git/objects/26
8556b6603a1a3ffe8db19a5f47975c8c8e1de0 ffeab31c732eaa025bb4895d092f38c6857bde
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:15:41] git cat-file -p 268556b6603a1a3ffe8db19a5f47975c8c8e1de0
Hello, Git internals
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:15:47] git cat-file -t 268556b6603a1a3ffe8db19a5f47975c8c8e1de0
blob
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:15:50]

```


# Concepts


## Blobs and Trees
- Blob Object : Represents the content of a file (just the data, not the filename or permissions)
- Blobs are immutable - any change to a file creates a new blob object

### See blob of ReadMe.md file
```

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:27:44] git ls-files --stage README.md
100644 f1a3d37520534ec3d3c32e7438f704d1ddaedf4c 0	README.md

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:34:40] git hash-object ReadMe.md
f1a3d37520534ec3d3c32e7438f704d1ddaedf4c
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:34:49]

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:27:52]

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:27:17] git cat-file -p f1a3d37520534ec3d3c32e7438f704d1ddaedf4c | wc -l
     245
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:27:23] git cat-file -t f1a3d37520534ec3d3c32e7438f704d1ddaedf4c
blob
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:27:41] git cat-file -s f1a3d37520534ec3d3c32e7438f704d1ddaedf4c
6589

```

### Tree object for directories
```
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:31:59] git ls-tree HEAD | grep src
040000 tree 82803084085ca790428ee491e88d10b79f36854c	src

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:32:24] git rev-parse HEAD:src/
82803084085ca790428ee491e88d10b79f36854c
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:33:11]

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:32:04] git cat-file -t 82803084085ca790428ee491e88d10b79f36854c
tree
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:32:16] git cat-file -s 82803084085ca790428ee491e88d10b79f36854c
171
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:32:20] git cat-file -p 82803084085ca790428ee491e88d10b79f36854c
040000 tree 22af126647a7bd51867fa930dfec97df48973670	cmd
040000 tree c4467a44710ccfc162fc4ee9bb81d3a0fbb3dcf4	config
040000 tree 7faa78752b00e09cc3a0450cd332cc3ac2ec76e1	internal
040000 tree ada7ffaf73afd10984b6dce8281e1a05df12c535	pkg
040000 tree ba8ae03f81b25b4da4eece118f5d069ecfe3cc8f	superfile_config
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:32:24]
```

### Root level tree sha 

```
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:51:28] git rev-parse HEAD^{tree}
ca378cf20ab0ac7e6da99966a4b9e04fdcfc9885
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:51:38] git cat-file -p HEAD | grep tree
tree ca378cf20ab0ac7e6da99966a4b9e04fdcfc9885
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:51:57]
```

Root level tree for a commit
```
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:52:02] git rev-parse HEAD
488cf8da8643619eadb2bbd9e71aa9c1d3af1db7
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:53:40] git rev-parse HEAD~1
e654b25b8296765fbcd8b2f7837e9bfdfb62fbdc
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:53:43] git cat-file -t 488cf8da8643619eadb2bbd9e71aa9c1d3af1db7
commit
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:53:49] git cat-file -t e654b25b8296765fbcd8b2f7837e9bfdfb62fbdc
commit
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:53:53] git cat-file -p 488cf8da8643619eadb2bbd9e71aa9c1d3af1db7 | grep tree
tree ca378cf20ab0ac7e6da99966a4b9e04fdcfc9885
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:54:18] git rev-parse 488cf8da8643619eadb2bbd9e71aa9c1d3af1db7^{tree}
ca378cf20ab0ac7e6da99966a4b9e04fdcfc9885
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:54:31] git rev-parse e654b25b8296765fbcd8b2f7837e9bfdfb62fbdc^{tree}
7018cccb91bda7cf040b5ba380035e71367b06ef
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [9:54:39]
```

## Commit object
Represents a specific version of your project
Contains:
- A pointer to the top-level tree (representing project root)
- Pointers to parent commit(s)
Author and committer information
Commit message
Timestamp

## Storage

- Loose objects .git/objects/<first-2-chars-of-hash>/<remaining-38-chars>
- Packed objects in .git/objects/pack
- References are pointers to commits stored in simple text files:
  - Branches: Stored in .git/refs/heads/
  - Each branch file contains the SHA-1 of the latest commit on that branch
  - Tags: Stored in .git/refs/tags/
  - Remote branches: Stored in .git/refs/remotes/
  - HEAD: Special reference (usually in .git/HEAD) pointing to the current branch

## Git's three areas
Git manages your files in three main areas:

### Working Tree (Working Directory)

The actual files on your filesystem that you can see, edit, and interact with
When you run ls in your terminal, you're seeing the working tree
Files here may be tracked or untracked by Git

### Staging Area (Index)

`.git/index` , which actually contains information about all tracked files in your current branch's HEAD commit. This means it holds:
A list of all tracked files
Their blob hashes (SHA-1)
Permissions and other metadata
Statistics like timestamps and file sizes

- Acts as a middle ground between working tree and repository
- The `git ls-files --stage` command shows the contents of the staging area (index)

### Repository (.git directory)

Where Git permanently stores your project history as objects
Contains all commits, trees, blobs and other Git objects
The .git directory is what makes a directory a Git repository

# Exploration

## Create a sample object and add it to index, then commit it and update the reference

```bash
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:04:34] echo "Hello, Git internals" | git hash-object -w --stdin

268556b6603a1a3ffe8db19a5f47975c8c8e1de0
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev)

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:15:41] git cat-file -p 268556b6603a1a3ffe8db19a5f47975c8c8e1de0
Hello, Git internals
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:15:47] git cat-file -t 268556b6603a1a3ffe8db19a5f47975c8c8e1de0
blob
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:15:50]

# Add to index
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) [10:16:55] git update-index --add --cacheinfo 100644 268556b6603a1a3ffe8db19a5f47975c8c8e1de0 filename.txt
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:17:41]

# Verification
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:18:50] git ls-files --stage | grep filename.txt
100644 268556b6603a1a3ffe8db19a5f47975c8c8e1de0 0	filename.txt
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:18:54] git status
On branch lazysegtree_dev
Your branch is up to date with 'origin/lazysegtree_dev'.

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	new file:   filename.txt

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    filename.txt

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:18:56] ls
asset           cd_on_quit      flake.nix       gomod2nix.toml  main.go         src             website
bin             CONTRIBUTING.md go.mod          hi              README.md       testsuite
build.sh        flake.lock      go.sum          LICENSE         release         vhs
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:19:00]

# Write a tree object from the current index

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:19:00] git write-tree
0ad1fdf62b25cea0075eab821602a016f5691b47
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:19:26] ls
asset           cd_on_quit      flake.nix       gomod2nix.toml  main.go         src             website
bin             CONTRIBUTING.md go.mod          hi              README.md       testsuite
build.sh        flake.lock      go.sum          LICENSE         release         vhs
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:19:30] git status
On branch lazysegtree_dev
Your branch is up to date with 'origin/lazysegtree_dev'.

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
	new file:   filename.txt

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    filename.txt

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:19:53] git cat-file -t 0ad1fdf62b25cea0075eab821602a016f5691b47
tree
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:20:05] git cat-file -p 0ad1fdf62b25cea0075eab821602a016f5691b47 | grep filename
100644 blob 268556b6603a1a3ffe8db19a5f47975c8c8e1de0	filename.txt
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:20:15]


# Update ref to that commit, 

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:20:15] # Create a commit object from a tree
echo "Initial commit" | git commit-tree 0ad1fdf62b25cea0075eab821602a016f5691b47
0f4865248fbd5be60c3145e9f10c6fa03fb1d6fd
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:20:46] git symbolic-ref HEAD
refs/heads/lazysegtree_dev
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:21:20] git update-ref refs/heads/lazysegtree_dev 0f4865248fbd5be60c3145e9f10c6fa03fb1d6fd
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:21:36] git status
On branch lazysegtree_dev
Your branch and 'origin/lazysegtree_dev' have diverged,
and have 1 and 1766 different commits each, respectively.
  (use "git pull" to merge the remote branch into yours)

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    filename.txt

no changes added to commit (use "git add" and/or "git commit -a")
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:21:46] ls
asset           cd_on_quit      flake.nix       gomod2nix.toml  main.go         src             website
bin             CONTRIBUTING.md go.mod          hi              README.md       testsuite
build.sh        flake.lock      go.sum          LICENSE         release         vhs
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:21:51]


# The above  commit has no parent, so it will be bad. Create a new commit, with last
# lazysegtree_dev commit as parent commit

➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:23:39] cat .git/refs/remotes/origin/lazysegtree_dev
488cf8da8643619eadb2bbd9e71aa9c1d3af1db7
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:23:42] echo "File commit" | git commit-tree 0ad1fdf62b25cea0075eab821602a016f5691b47 -p 488cf8da8643619eadb2bbd9e71aa9c1d3af1db7
e94cb28f5d0d88f7ac7d76e1f7922ef5eaf2130e
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:24:10] git update-ref refs/heads/lazysegtree_dev e94cb28f5d0d88f7ac7d76e1f7922ef5eaf2130e
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:24:19] git status
On branch lazysegtree_dev
Your branch is ahead of 'origin/lazysegtree_dev' by 1 commit.
  (use "git push" to publish your local commits)

Changes not staged for commit:
  (use "git add/rm <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
	deleted:    filename.txt

no changes added to commit (use "git add" and/or "git commit -a")
➜  ~/Workspace/kuknitin/superfile git:(lazysegtree_dev) ✗ [10:24:22]

```