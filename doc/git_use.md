$ ssh-keygen -t rsa -C "1476048558@qq.com"

$ ssh -T git@github.com

$ git config --global user.email "you@example.com"
$ git config --global user.name "Your Name"


$ git remote add origin git@github.com:1476048558@qq.com/BWF

$ git add *
$ git commit -m "info"
$ git push origin master

合并分支：
创建一个叫做“feature_x”的分支，并切换过去：
$ git checkout -b feature_x
切换回主分支：
$ git checkout master
再把新建的分支删掉：
$ git branch -d feature_x
除非你将分支推送到远端仓库，不然该分支就是 不为他人所见的：
$ git push origin <branch>

将master合并到当前分支main 
$ git merge master

fatal: 拒绝合并无关的历史:
git merge master --allow-unrelated-histories


fork:
git remote -v 

git remote add upstream https://github.com/XXX/XXX.git

git fetch upstream 

git merge upstream/master