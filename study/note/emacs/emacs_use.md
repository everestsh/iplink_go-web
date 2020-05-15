# emacs use

**一些默认快捷键**

1>

一些默认快捷键

C-代表按住Ctrl键

M-代表按住Alt(mac是command)键，或先按ESC键再按x键，ESC和x不是同时按下
```
C-x C-c        退出emacs
```

2>

```
最先要记住的
M-x <cmd>      输入指令执行，在输入时用Tab可以自动补全或列出项目
C-g            取消当前操作指令
C-h k <key>    查看当前按键绑定的指令及介绍
```



emacs默认是没有打开语法高亮显示功能的，可通过以下命令开启emacs的语法高亮显示，启动emacs后按下M-x global-font-lock-mode,语法高亮显示就打开了，或者在你的.emacs里面加一句
(global-font-lock-mode t)
M-x：M代表Meta键盘，可用Alt键代替或者先按ESC键再按x键，ESC和x不是同时按下


