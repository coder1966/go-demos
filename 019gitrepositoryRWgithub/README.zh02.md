

https://blog.csdn.net/u014686399/article/details/108334869
# golang使用go-git上传代码到gitlab远端仓库


直接上代码了
//git clone
var gitRep *git.Repository
// gitUser: git的用户名  password: git的用户密码
var gitAuth = &http.BasicAuth{Username:gitUser, Password: gitPassword}
//gitTmpDir 是指clone到本地后，本地的目录
gitRep, endError = git.PlainClone(gitTmpDir, false,  &git.CloneOptions{
    // gitCloneUrl是你的项目在gitlab上的地址
	URL: gitCloneUrl,
	RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	Auth: gitAuth,
	// refs/heads是必须有的， %s 可以是你的分支的名字
	ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", nameInfoObj.Version)),
})
1
2
3
4
5
6
7
8
9
10
11
12
13
//git add
//获取工程树
var tree *git.Worktree
tree, endError = gitRep.Worktree()
if endError != nil{
	return HandleException(endError, "get local work tree")
}
//拷贝文件过来
....此处省略无数行代码
//ansiblePath 是你的目录的名字，注意这里是指在clone后的工程目录里相对路径,比如 如下文clone后的目录是redis， 然后我拷贝了ansible文件夹到了这个目录，当需要add ansible到暂存区，那么ansiblePath就是ansible：
//    --redis
 //     ---- .git
 //     ---- ansible
_, endError = tree.Add(ansiblePath)
if endError!=nil{
	return HandleException(endError, "add ansible to gitlab")
}
1
2
3
4
5
6
7
8
9
10
11
12
13
14
15
16
17
//git commit
var treeCommit plumbing.Hash
treeCommit, endError = tree.Commit("The first commit by chengying", &git.CommitOptions{
	All: true,
	Author: &object.Signature{
		Name:  "John Doe",
		//Email: "john@doe.org",
		When:  time.Now(),
	},
})
_, endError = gitRep.CommitObject(treeCommit)
if endError!=nil{
	return HandleException(endError, "commit Rep to gitlab")
}
1
2
3
4
5
6
7
8
9
10
11
12
13
14
//git push
//refs/heads 是必须要有的
// +的含义我还没有整明白，就先不写了，如果觉得有问题可以去掉
endError = gitRep.Push(&git.PushOptions{RemoteName: "origin",
	RefSpecs: []config.RefSpec{config.RefSpec("+refs/heads/"+nameInfoObj.Version+":refs/heads/"+nameInfoObj.Version)},
	Auth: gitAuth,
})
1
2
3
4
5
6
7
然后检测一下远端仓库的代码，看看有没有推送过去
————————————————
版权声明：本文为CSDN博主「runing_an_min」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/u014686399/article/details/108334869