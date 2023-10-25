package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/go-git/go-git/v5"
	. "github.com/go-git/go-git/v5/_examples"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/cache"
	commitgraph_fmt "github.com/go-git/go-git/v5/plumbing/format/commitgraph"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/object/commitgraph"
	"github.com/go-git/go-git/v5/storage/filesystem"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

// Example how to resolve a revision into its commit counterpart
// 如何将修订解析为其对应的提交的示例
func main() {
	CheckArgs("<path>", "<revision>", "<tree path>")

	path := os.Args[1]
	revision := os.Args[2]
	treePath := os.Args[3]

	// We instantiate a new repository targeting the given path (the .git folder)
	// 我们实例化一个针对给定路径（.git 文件夹）的新存储库
	fs := osfs.New(path)
	if _, err := fs.Stat(git.GitDirName); err == nil {
		fs, err = fs.Chroot(git.GitDirName)
		CheckIfError(err)
	}

	s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})
	r, err := git.Open(s, fs)
	CheckIfError(err)
	defer s.Close()

	// Resolve revision into a sha1 commit, only some revisions are resolved look at the doc to get more details
	// 将修订解决为 sha1 提交，仅解决了部分修订，请查看文档以获取更多详细信息
	Info("git rev-parse %s", revision)

	h, err := r.ResolveRevision(plumbing.Revision(revision))
	CheckIfError(err)

	commit, err := r.CommitObject(*h)
	CheckIfError(err)

	tree, err := commit.Tree()
	CheckIfError(err)
	if treePath != "" {
		tree, err = tree.Tree(treePath)
		CheckIfError(err)
	}

	var paths []string
	for _, entry := range tree.Entries {
		paths = append(paths, entry.Name)
	}

	commitNodeIndex, file := getCommitNodeIndex(r, fs)
	if file != nil {
		defer file.Close()
	}

	commitNode, err := commitNodeIndex.Get(*h)
	CheckIfError(err)

	revs, err := getLastCommitForPaths(commitNode, treePath, paths)
	CheckIfError(err)
	for path, rev := range revs {
		// Print one line per file (name hash message)
		// 每个文件打印一行（名称哈希消息）
		hash := rev.Hash.String()
		line := strings.Split(rev.Message, "\n")
		fmt.Println(path, hash[:7], line[0])
	}
}

func getCommitNodeIndex(r *git.Repository, fs billy.Filesystem) (commitgraph.CommitNodeIndex, io.ReadCloser) {
	file, err := fs.Open(path.Join("objects", "info", "commit-graph"))
	if err == nil {
		index, err := commitgraph_fmt.OpenFileIndex(file)
		if err == nil {
			return commitgraph.NewGraphCommitNodeIndex(index, r.Storer), file
		}
		file.Close()
	}

	return commitgraph.NewObjectCommitNodeIndex(r.Storer), nil
}

type commitAndPaths struct {
	commit commitgraph.CommitNode
	// Paths that are still on the branch represented by commit
	// 仍在提交代表的分支上的路径
	paths []string
	// Set of hashes for the paths
	// 路径的哈希集
	hashes map[string]plumbing.Hash
}

func getCommitTree(c commitgraph.CommitNode, treePath string) (*object.Tree, error) {
	tree, err := c.Tree()
	if err != nil {
		return nil, err
	}

	// Optimize deep traversals by focusing only on the specific tree
	// 通过仅关注特定树来优化深度遍历
	if treePath != "" {
		tree, err = tree.Tree(treePath)
		if err != nil {
			return nil, err
		}
	}

	return tree, nil
}

func getFullPath(treePath, path string) string {
	if treePath != "" {
		if path != "" {
			return treePath + "/" + path
		}
		return treePath
	}
	return path
}

func getFileHashes(c commitgraph.CommitNode, treePath string, paths []string) (map[string]plumbing.Hash, error) {
	tree, err := getCommitTree(c, treePath)
	if err == object.ErrDirectoryNotFound {
		// The whole tree didn't exist, so return empty map
		// 整棵树不存在，因此返回空地图
		return make(map[string]plumbing.Hash), nil
	}
	if err != nil {
		return nil, err
	}

	hashes := make(map[string]plumbing.Hash)
	for _, path := range paths {
		if path != "" {
			entry, err := tree.FindEntry(path)
			if err == nil {
				hashes[path] = entry.Hash
			}
		} else {
			hashes[path] = tree.Hash
		}
	}

	return hashes, nil
}

func getLastCommitForPaths(c commitgraph.CommitNode, treePath string, paths []string) (map[string]*object.Commit, error) {
	// We do a tree traversal with nodes sorted by commit time
	// 我们对按提交时间排序的节点进行树遍历
	heap := binaryheap.NewWith(func(a, b interface{}) int {
		if a.(*commitAndPaths).commit.CommitTime().Before(b.(*commitAndPaths).commit.CommitTime()) {
			return 1
		}
		return -1
	})

	resultNodes := make(map[string]commitgraph.CommitNode)
	initialHashes, err := getFileHashes(c, treePath, paths)
	if err != nil {
		return nil, err
	}

	// Start search from the root commit and with full set of paths
	// 从根提交开始搜索并使用完整的路径集
	heap.Push(&commitAndPaths{c, paths, initialHashes})

	for {
		cIn, ok := heap.Pop()
		if !ok {
			break
		}
		current := cIn.(*commitAndPaths)

		// Load the parent commits for the one we are currently examining
		// 加载我们当前正在检查的父提交
		numParents := current.commit.NumParents()
		var parents []commitgraph.CommitNode
		for i := 0; i < numParents; i++ {
			parent, err := current.commit.ParentNode(i)
			if err != nil {
				break
			}
			parents = append(parents, parent)
		}

		// Examine the current commit and set of interesting paths
		// 检查当前提交和一组有趣的路径
		pathUnchanged := make([]bool, len(current.paths))
		parentHashes := make([]map[string]plumbing.Hash, len(parents))
		for j, parent := range parents {
			parentHashes[j], err = getFileHashes(parent, treePath, current.paths)
			if err != nil {
				break
			}

			for i, path := range current.paths {
				if parentHashes[j][path] == current.hashes[path] {
					pathUnchanged[i] = true
				}
			}
		}

		var remainingPaths []string
		for i, path := range current.paths {
			// The results could already contain some newer change for the same path, so don't override that and bail out on the file early.
			// 结果可能已经包含同一路径的一些更新的更改，因此不要覆盖它并尽早放弃文件。
			if resultNodes[path] == nil {
				if pathUnchanged[i] {
					// The path existed with the same hash in at least one parent so it could not have been changed in this commit directly.
					// 该路径至少在一个父级中存在相同的哈希值，因此无法在此提交中直接更改它。
					remainingPaths = append(remainingPaths, path)
				} else {
					// There are few possible cases how can we get here:
					// 有几种可能的情况我们如何到达这里：
					// - The path didn't exist in any parent, so it must have been created by this commit.
					// - 该路径在任何父级中都不存在，因此它必须是由该提交创建的。
					// - The path did exist in the parent commit, but the hash of the file has changed.
					// - 路径确实存在于父提交中，但文件的哈希值已更改。
					// - We are looking at a merge commit and the hash of the file doesn't match any of the hashes being merged. This is more common for directories, but it can also happen if a file is changed through conflict resolution.
					// - 我们正在查看合并提交，文件的哈希值与任何正在合并的哈希值都不匹配。这对于目录来说更常见，但如果通过冲突解决更改文件，也可能会发生这种情况。
					resultNodes[path] = current.commit
				}
			}
		}

		if len(remainingPaths) > 0 {
			// Add the parent nodes along with remaining paths to the heap for further processing.
			// 将父节点以及剩余路径添加到堆中以进行进一步处理。
			for j, parent := range parents {
				// Combine remainingPath with paths available on the parent branch and make union of them
				// 将剩余路径与父分支上可用的路径组合起来并进行合并
				remainingPathsForParent := make([]string, 0, len(remainingPaths))
				newRemainingPaths := make([]string, 0, len(remainingPaths))
				for _, path := range remainingPaths {
					if parentHashes[j][path] == current.hashes[path] {
						remainingPathsForParent = append(remainingPathsForParent, path)
					} else {
						newRemainingPaths = append(newRemainingPaths, path)
					}
				}

				if remainingPathsForParent != nil {
					heap.Push(&commitAndPaths{parent, remainingPathsForParent, parentHashes[j]})
				}

				if len(newRemainingPaths) == 0 {
					break
				} else {
					remainingPaths = newRemainingPaths
				}
			}
		}
	}

	// Post-processing
	result := make(map[string]*object.Commit)
	for path, commitNode := range resultNodes {
		var err error
		result[path], err = commitNode.Commit()
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
