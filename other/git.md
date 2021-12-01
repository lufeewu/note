# 简介
git 是目前最流行的代码托管工具。通过 git 的代码管理，可以进行许多不同类型的流程的开发，如主干开发、分支开发以及 fork 工作流。

## fork 工作流
fork 工作流使用, 这种工作流不是使用单个服务端仓库作为中央代码基线，而让各位开发者都有一个服务端仓库。一个本地私有、一个服务端公开的。开发者 push 到自己的服务端，项目维护者才能 push 到正式仓库。这样项目维护者可以接受任何开发者的提交而不需要给正式代码库的写权限。

### fork 工作方式
开发者 fork 了代码仓库后，通过如下流程进行开发:
1. git clone 代码仓库(fork 仓库)到本地私有开发环境
2. 提交本地修改后，push 到自己的仓库中
3. 给正式仓库提交一个 pull request
4. 正式仓库维护者同意后，将 pull 贡献者的变更到私有代码仓库中，合并变更到本地的 master 分支
5. 正式仓库维护者提交 push master 分支到服务器的正式仓库中

git remote add upstream && git fetch upstream


## 参考
1. [Forking工作流](https://github.com/oldratlee/translations/blob/master/git-workflows-and-tutorials/workflow-forking.md)