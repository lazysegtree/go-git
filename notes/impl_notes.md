- Storer is for storing and managing .git
- Remote object needs storer to update refs.
- Shared clone is cloning from a repo alread on your disk, while saving space to not copy all objects
- transport.Endpoint is a struct that encapsulates all the connection details for a Git repository

## Logs
### Simple repo clone 
```
➜  ~/Workspace/kuknitin/go-git/_examples/lst_repo_info git:(v6-transport-lazysegtree) ✗ [8:30:26] go run main.go https://www.github.com/lazysegtree/temp-share
[LST] cleanup true, cleanupParent false, err <nil>
[LST] path : /tmp/lst_test/list_commits, branch :
[LST] dot : /tmp/lst_test/list_commits/.git, statErr : stat /tmp/lst_test/list_commits/.git: no such file or directory
[LST] InitWithOptions(), branch : refs/heads/master
[LST] New repository object r, ref is nil
[LST] Added new remote with name : origin, config : &{origin [https://www.github.com/lazysegtree/temp-share] false false [] [+refs/heads/*:refs/remotes/origin/*] <nil>}
[LST] Fetching from remote URL: https://www.github.com/lazysegtree/temp-share, RefSpecs: [+refs/heads/*:refs/remotes/origin/*]
[LST] Creating new client with protocol: https
[LST] Created new client with endpoint: https://www.github.com/lazysegtree/temp-share, supported protocols: [0 1]
[LST] Handshaking with HTTP session, req : &{GET https://www.github.com/lazysegtree/temp-share/info/refs?service=git-upload-pack HTTP/1.1 1 1 map[Accept:[application/x-git-upload-pack-result] Content-Type:[application/x-git-upload-pack-request] Host:[www.github.com] User-Agent:[go-git/6.x]] <nil> <nil> 0 [] false www.github.com map[] map[] <nil> map[]   <nil> <nil> <nil>  {{}} <nil> [] map[]}
[LST] Handshaking with HTTP session, status code : 200
[LST] before modifyRedirect ep : https://www.github.com/lazysegtree/temp-share
[LST] after modifyRedirect ep : https://github.com/lazysegtree/temp-share.git
[LST] before decoding, ar : &{<nil>  map[] map[] []}, s.IsSmart() : true
[LST] read line "# service=git-upload-pack\n"
[LST] reply : {git-upload-pack}
[LST] s.version : 0
[LST] state function: github.com/go-git/go-git/v6/plumbing/protocol/packp.decodeFirstHash
[LST] decodeFirstHash, p.line : 23f6c99cb342bcc39febf74a24e91c83385e8d5c HEADmulti_ack thin-pack side-band side-band-64k ofs-delta shallow deepen-since deepen-not deepen-relative no-progress include-tag multi_ack_detailed allow-tip-sha1-in-want allow-reachable-sha1-in-want no-done symref=HEAD:refs/heads/main filter object-format=sha1 agent=git/github-413959fa1280
[LST] state function: github.com/go-git/go-git/v6/plumbing/protocol/packp.decodeFirstRef
[LST] decodeFirstRef, l.line :  HEADmulti_ack thin-pack side-band side-band-64k ofs-delta shallow deepen-since deepen-not deepen-relative no-progress include-tag multi_ack_detailed allow-tip-sha1-in-want allow-reachable-sha1-in-want no-done symref=HEAD:refs/heads/main filter object-format=sha1 agent=git/github-413959fa1280
[LST] state function: github.com/go-git/go-git/v6/plumbing/protocol/packp.decodeCaps
[LST] decodeCaps, p.line : multi_ack thin-pack side-band side-band-64k ofs-delta shallow deepen-since deepen-not deepen-relative no-progress include-tag multi_ack_detailed allow-tip-sha1-in-want allow-reachable-sha1-in-want no-done symref=HEAD:refs/heads/main filter object-format=sha1 agent=git/github-413959fa1280
[LST] state function: github.com/go-git/go-git/v6/plumbing/protocol/packp.decodeOtherRefs
[LST] decodeOtherRefs, p.line : 23f6c99cb342bcc39febf74a24e91c83385e8d5c refs/heads/main
[LST] state function: github.com/go-git/go-git/v6/plumbing/protocol/packp.decodeOtherRefs
[LST] decodeOtherRefs, p.line :
[LST] after decoding, ar : &{23f6c99cb342bcc39febf74a24e91c83385e8d5c multi_ack thin-pack side-band side-band-64k ofs-delta shallow deepen-since deepen-not deepen-relative no-progress include-tag multi_ack_detailed allow-tip-sha1-in-want allow-reachable-sha1-in-want no-done symref=HEAD:refs/heads/main filter object-format=sha1 agent=git/github-413959fa1280 map[refs/heads/main:23f6c99cb342bcc39febf74a24e91c83385e8d5c] map[] []}
[LST] Handshake with HTTP session, server version: 0
[LST] Getting remote refs from HTTP session, s.refs : &{23f6c99cb342bcc39febf74a24e91c83385e8d5c multi_ack thin-pack side-band side-band-64k ofs-delta shallow deepen-since deepen-not deepen-relative no-progress include-tag multi_ack_detailed allow-tip-sha1-in-want allow-reachable-sha1-in-want no-done symref=HEAD:refs/heads/main filter object-format=sha1 agent=git/github-413959fa1280 map[refs/heads/main:23f6c99cb342bcc39febf74a24e91c83385e8d5c] map[] []}
[LST] Fetched ref from remote(name : HEAD, hash : 0000000000000000000000000000000000000000, target : refs/heads/main)
[LST] Fetched ref from remote(name : refs/heads/main, hash : 23f6c99cb342bcc39febf74a24e91c83385e8d5c, target : )
[LST] Fetching with HTTP session, req : &{<nil> [23f6c99cb342bcc39febf74a24e91c83385e8d5c] [] 0  false}
[LST] performance: 1.071910750 s: git command: git clone https://www.github.com/lazysegtree/temp-share
remotes count : 1
➜  ~/Workspace/kuknitin/go-git/_examples/lst_repo_info git:(v6-transport-lazysegtree) ✗ [8:34:09]
```


### Clone a repo with many branches
See [./superfile_clone.md]