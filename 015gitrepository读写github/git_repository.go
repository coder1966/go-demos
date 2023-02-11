package main

// import (
// 	"fmt"

// 	"github.com/go-git/go-git"
// 	"github.com/go-git/go-git/storage/memory"
// )

// func main() {

// 	gitCloneMEM()

// 	// gitClone()

// 	// gitAdd()

// 	// gitCommit()

// 	// gitPush()

// }

// func gitCloneMEM() {
// 	// Clones the given repository in memory, creating the remote, the local
// 	// branches and fetching the objects, exactly as:
// 	// Info("git clone https://github.com/go-git/go-billy")

// 	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
// 		URL: "https://github.com/go-git/go-billy",
// 	})

// 	// CheckIfError(err)
// 	if err != nil {
// 		fmt.Println(" error: ", err)
// 	}

// 	// Gets the HEAD history from HEAD, just like this command:
// 	// Info("git log")

// 	// ... retrieves the branch pointed by HEAD
// 	ref, err := r.Head()
// 	// CheckIfError(err)
// 	if err != nil {
// 		fmt.Println(" error: ", err)
// 	}

// 	// ... retrieves the commit history
// 	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
// 	// CheckIfError(err)
// 	if err != nil {
// 		fmt.Println(" error: ", err)
// 	}

// 	_ = cIter

// 	// ... just iterates over the commits, printing it
// 	// err = cIter.ForEach(func(c *object.Commit) error {
// 	// 	fmt.Println(c)
// 	// 	return nil
// 	// })
// 	// // CheckIfError(err)
// 	// if err != nil {
// 	// 	fmt.Println(" error: ", err)
// 	// }
// }

// /*
// func gitClone() {

// 	Info("git clone https://github.com/go-git/go-billy")
// 	//git clone
// 	// var gitRep *git.Repository
// 	// gitUser: git的用户名  password: git的用户密码
// 	var gitAuth = &http.BasicAuth{Username: gitUser, Password: gitPassword}
// 	//gitTmpDir 是指clone到本地后，本地的目录
// 	gitRep, endError := git.PlainClone(gitTmpDir, false, &git.CloneOptions{
// 		// gitCloneUrl是你的项目在gitlab上的地址
// 		URL:               gitCloneUrl,
// 		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
// 		Auth:              gitAuth,
// 		// refs/heads是必须有的， %s 可以是你的分支的名字
// 		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", nameInfoObj.Version)),
// 	})
// }

// func gitAdd() {
// 	//git add
// 	//获取工程树
// 	var tree *git.Worktree
// 	tree, endError = gitRep.Worktree()
// 	if endError != nil {
// 		return HandleException(endError, "get local work tree")
// 	}
// 	//拷贝文件过来
// 	// ....此处省略无数行代码
// 	//ansiblePath 是你的目录的名字，注意这里是指在clone后的工程目录里相对路径,比如 如下文clone后的目录是redis， 然后我拷贝了ansible文件夹到了这个目录，当需要add ansible到暂存区，那么ansiblePath就是ansible：
// 	//    --redis
// 	//     ---- .git
// 	//     ---- ansible
// 	_, endError = tree.Add(ansiblePath)
// 	if endError != nil {
// 		return HandleException(endError, "add ansible to gitlab")
// 	}

// }

// func gitCommit() {
// 	//git commit
// 	var treeCommit plumbing.Hash
// 	treeCommit, endError = tree.Commit("The first commit by chengying", &git.CommitOptions{
// 		All: true,
// 		Author: &object.Signature{
// 			Name: "John Doe",
// 			//Email: "john@doe.org",
// 			When: time.Now(),
// 		},
// 	})
// 	_, endError = gitRep.CommitObject(treeCommit)
// 	if endError != nil {
// 		return HandleException(endError, "commit Rep to gitlab")
// 	}

// }

// func gitPush() {
// 	//git push
// 	//refs/heads 是必须要有的
// 	// +的含义我还没有整明白，就先不写了，如果觉得有问题可以去掉
// 	endError = gitRep.Push(&git.PushOptions{RemoteName: "origin",
// 		RefSpecs: []config.RefSpec{config.RefSpec("+refs/heads/" + nameInfoObj.Version + ":refs/heads/" + nameInfoObj.Version)},
// 		Auth:     gitAuth,
// 	})

// }

// // Branches returns all the References that are Branches.
// func Branches(r *git.Repository) (storer.ReferenceIter, error) {
// 	_ = r
// 	return nil, nil
// }
// */
