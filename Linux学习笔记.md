# Linux 后端开发常用命令笔记

| 分类               | 命令    | 作用说明                                                                 | 常用示例                                                                 |
|--------------------|---------|--------------------------------------------------------------------------|--------------------------------------------------------------------------|
| **目录与文件操作** | `pwd`   | 查看当前所在路径                                                         | `pwd`                                                                    |
|                    | `ls`    | 列出目录内容                                                             | `ls` / `ls -l`（详情） / `ls -a`（含隐藏文件）                           |
|                    | `cd`    | 切换目录                                                                 | `cd /home` / `cd ..`（上一级） / `cd ~`（回到家目录）                   |
|                    | `mkdir` | 创建目录                                                                 | `mkdir mydir` / `mkdir -p a/b/c`（递归创建多级目录）                     |
|                    | `rm`    | 删除文件/目录（-rf 慎用，强制递归删除）                                  | `rm file.txt` / `rm -rf dir/`                                            |
|                    | `touch` | 创建空文件                                                               | `touch main.go`                                                          |
|                    | `cat`   | 查看文件全部内容                                                         | `cat main.go`                                                            |
| **权限与执行**     | `chmod` | 修改文件/目录权限（r=4, w=2, x=1）                                       | `chmod 755 app`（所有者rwx，其他rx）                                     |
|                    | `sudo`  | 以管理员权限执行命令                                                     | `sudo apt install golang-go`                                             |
| **网络与端口**     | `ping`  | 测试网络连通性                                                           | `ping baidu.com`                                                         |
|                    | `netstat` | 查看网络连接与端口占用                                                 | `netstat -tulnp`（查看监听端口）                                         |
| **进程与服务**     | `ps`    | 查看进程状态                                                             | `ps aux`（列出所有进程） / `ps aux \| grep go`（过滤Go进程）             |
|                    | `systemctl` | 管理系统服务（启动/停止/查看状态）                                   | `systemctl start nginx` / `systemctl status nginx`                       |
| **调试与日志查看** | `curl`  | 发送HTTP请求/测试接口                                                     | `curl http://localhost:8080/health`                                      |
|                    | `tail`  | 查看文件尾部内容，`-f` 可实时滚动查看日志                                | `tail -f app.log`（Ctrl+C 退出）                                         |