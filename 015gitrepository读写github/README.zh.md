https://blog.csdn.net/m0_51573433/article/details/126449687
# Go 原生的 git 实现库：go-git



1. go-git 介绍
2. go-git 使用
2.1）go-git 的安装：
2.2）拉取仓库：
2.3）获取 last commit hash：
2.4）对 commit 历史信息进行遍历
1. go-git 介绍
一个用 Go 语言编写的 git 实现库：它的官方仓库地址：go-git

为什么我们需要它？
举个例子：

如果我们需要获取 git log 的信息，需要通过 Go 调用 cmd 命令来获取：获取 git 的 last commit hash 的话需要耗费的时间大概在 50ms 左右（不同运行环境可能不一样）

它的慢原因是什么呢 ？
Go 调用 cmd 命令本身存在着性能原因
其次，如果想要获取其它仓库的 git log 信息，还存在着跨文件调用的消耗

所以这里 go-git 的作用就体现出来了！由于它是 go 编写的 git 实现库，所以对 cmd 命令调用的消耗就可以避免了。测试之后，使用时间在 500μs 左右，巨大的提升！

2. go-git 使用
2.1）go-git 的安装：
go get -u github.com/go-git/go-git/v5
1
2.2）拉取仓库：
通过 github 地址
r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
    URL: "仓库地址",
})
1
2
3
通过本地仓库
r, err := git.PlainOpen("本地仓库路径")
1
2.3）获取 last commit hash：
rHead, _ := r.Head()
// 将获得的 
rHeadStr := rHead.String()
rHeadIdx := strings.Index(rHeadStr, " ")
lastCommitHash = rHeadStr[:rHeadIdx]
1
2
3
4
5
r.Head() 获取的是一个结构体存有的数据，数据格式可以阅读源码获得：

// Reference is a representation of git reference
type Reference struct {
	t      ReferenceType
	n      ReferenceName
	h      Hash
	target ReferenceName
}
// Hash SHA1 hashed content
type Hash [20]byte
func (h Hash) String() string {
	return hex.EncodeToString(h[:])
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
所以可以使用 String() 将获得的信息转为字符串，结尾有空格需要去掉。

2.4）对 commit 历史信息进行遍历
cIter, _ := r.Log(&git.LogOptions{From: ref.Hash()})
err = cIter.ForEach(func(c *object.Commit) error {
	return nil
})
1
2
3
4
LogOptions 的填写也可以通过阅读源码获取：

type LogOptions struct {
	// When the From option is set the log will only contain commits
	// reachable from it. If this option is not set, HEAD will be used as
	// the default From.
	From plumbing.Hash
	// The default traversal algorithm is Depth-first search
	// set Order=LogOrderCommitterTime for ordering by committer time (more compatible with `git log`)
	// set Order=LogOrderBSF for Breadth-first search
	Order LogOrder
	// Show only those commits in which the specified file was inserted/updated.
	// It is equivalent to running `git log -- <file-name>`.
	// this field is kept for compatility, it can be replaced with PathFilter
	FileName *string
	// Filter commits based on the path of files that are updated
	// takes file path as argument and should return true if the file is desired
	// It can be used to implement `git log -- <path>`
	// either <path> is a file path, or directory path, or a regexp of file/directory path
	PathFilter func(string) bool
	// Pretend as if all the refs in refs/, along with HEAD, are listed on the command line as <commit>.
	// It is equivalent to running `git log --all`.
	// If set on true, the From option will be ignored.
	All bool
	// Show commits more recent than a specific date.
	// It is equivalent to running `git log --since <date>` or `git log --after <date>`.
	Since *time.Time
	// Show commits older than a specific date.
	// It is equivalent to running `git log --until <date>` or `git log --before <date>`.
	Until *time.Time
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
18
19
20
21
22
23
24
25
26
27
28
29
可以通过：Order 指定遍历的方式，Since 和 Until 指定遍历时间范围。

cIter 是一个 object.Commiter 结构体，拥有以下方法，可以根据需要进行使用。

// CommitIter is a generic closable interface for iterating over commits.
type CommitIter interface {
	Next() (*Commit, error)
	ForEach(func(*Commit) error) error
	Close()
}
1
2
3
4
5
6
怎么根据只有一个 Hash 值获取 commit 呢 ？
thisHash := plumbing.NewHash(string("Hash 值"))
thisCommit, _ := object.GetCommit(r.Storer, thisHash)
1
2
根据 commit 结构体又可以获得别的信息：

type Commit struct {
	// Hash of the commit object.
	Hash plumbing.Hash
	// Author is the original author of the commit.
	Author Signature
	// Committer is the one performing the commit, might be different from
	// Author.
	Committer Signature
	// PGPSignature is the PGP signature of the commit.
	PGPSignature string
	// Message is the commit message, contains arbitrary text.
	Message string
	// TreeHash is the hash of the root tree of the commit.
	TreeHash plumbing.Hash
	// ParentHashes are the hashes of the parent commits of the commit.
	ParentHashes []plumbing.Hash
	s storer.EncodedObjectStorer
}
type Signature struct {
	// Name represents a person name. It is an arbitrary string.
	Name string
	// Email is an email, but it cannot be assumed to be well-formed.
	Email string
	// When is the timestamp of the signature.
	When time.Time
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
18
19
20
21
22
23
24
25
26
例如：获取 commit 的时间

thisCommit.Committer.When.String()
1

go-git 的出现对 Go 语言实现 git 命令产生了挺大的优化，更多使用方法可以阅读官方仓库进行学习：go-git
————————————————
版权声明：本文为CSDN博主「CSPsy」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/m0_51573433/article/details/126449687