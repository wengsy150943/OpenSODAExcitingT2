## 1. How to install

本项目采用golang开发，需要运行环境具有golang 1.18 及以上的版本 / Docker部署。

1. 编译源代码获得可执行文件。

   ```bash
   go build
   ```

   可以进一步将可执行文件安装进系统中。
   ```bash
   go install
   ```
2. Docker部署
   拉取docker image。
   ```bash
   docker pull yinzhengsun/exciting-opendigger:v1.0
   ```

   进入容器
   ```bash
   docker run -it yinzhengsun/exciting-opendigger:v1.0 /bin/bash
   ```

3. 执行可执行文件

   ```bash
   ./exciting-opendigger commands
   ```
   在将文件安装进系统后，可以在任意位置直接运行该文件
   ```bash
   exciting-opendigger commands
   ```
   
   具体的执行命令见下，其中均假定未将文件安装入系统中：

## 2. Usage

具体usage可运行./exciting-opendigger -h进行查看

| 模块 | 命令 | 结果 | 需求 |
| --- | --- | --- | --- |
| 单点查询 | SHOW | 单个仓库或开发者在特定指标、特定月份的数据或整体报告 | 对开源指标的简单浏览 |
| 比较查询 | COMPARE | 对两个仓库、开发者、月份的比较结果 | 对开源指标的差异分析 |
| 文件下载 | DOWNLOAD | 查询结果的文件和可视化分析 | 对开源指标的详细分析和可视化 |
| 批量分析 | BATCH | 批量对象的数据 | 对Github社区的宏观分析 |
| 日志查询 | LOG | 历史查询日志的查看和回放 | 连续使用本工具的便捷性 |
| 帮助菜单 | HELP | 使用说明 | 辅助工具的便捷使用 |
