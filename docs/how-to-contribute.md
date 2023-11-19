## 5. How to contribute

### Fork

从项目主页(https://github.com/wengsy150943/OpenSODAExcitingT2)点击`Fork`，将项目fork到自己的代码仓库。

### Clone

将上一步中fork的代码仓库clone到本地

```
git clone https://github.com/USERNAME/OpenSODAExcitingT2
cd OpenSODAExcitingT2
```

### Create local branch

创建自己的本地branch，在新的branch上进行开发

```
git checkout -b new_branch
```

### Start development

使用`git status`查看当前仓库分支状态

```bash
 git status
On branch test
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)
    modified:   README.md
Untracked files:
  (use "git add <file>..." to include in what will be committed)
    test
no changes added to commit (use "git add" and/or "git commit -a")
```

在service文件夹新建你需要实现功能的文件，并新建该功能的测试文件。

### Build and test

遵循README中项目构建步骤

```bash
go build
go install
```

自行编写测试代码用例，使用Go Test测试。

### Commit and push

使用下面的命令完成提交。

```bash
git commit -m "your commit info"
```

保持本地仓库最新

需要获取仓库最新代码并更新当前分支。

push到远程仓库

```bash
git push origin new_branch
```

### Pull request

点击 new pull request，选择本地分支和目标分支。在 PR 的描述说明中，填写该 PR 所完成的功能。接下来等待 review，如果有需要修改的地方，参照上述步骤更新 origin 中的对应分支即可。

#####　删除分支

在 PR 被 merge 进主仓库后，我们可以在 PR 的页面删除远程仓库的分支。

也可以使用 `git push origin :分支名` 删除远程分支，如：

```
git push origin :new_branch
```

##### 删除本地分支

```
# 切换到 main 分支，否则无法删除当前分支
git checkout main

# 删除 new_branch 分支
git branch -D new_branch
```

