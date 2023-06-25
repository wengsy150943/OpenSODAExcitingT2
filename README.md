

![Untitled](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/Untitled.png)

## 1.Introduction

### 1.1 项目背景

本项目的目标是开发一个命令行工具，为用户提供一种交互式的方式，用于查询和获取 OpenDigger 平台上的数据。OpenDigger 是一个提供各种统计型和网络型指标的平台，用户可以通过修改 HTTPS URL 来获取特定仓库或开发者在各项指标上的数据结果。为了更方便地查询和浏览这些数据，我们希望开发一种新型的指标结果查询方式，即通过命令行实现可交互的指标结果查询。

本命令行工具的核心功能包括查询特定仓库在特定指标上的数据、查询特定仓库在特定自然月上在特定指标上的数据，以及查询特定仓库在特定自然月份上的整体报告。同时，该工具还支持将查询结果导出到本地文件。

通过该工具，用户可以方便地在终端中进行各项指标的查询，探索仓库或开发者的数据情况。用户可以根据自己的需求进行定制化的查询，例如只查询特定指标或只查询特定类型的数据。

该命令行工具的设计旨在提供简单易用、高效便捷的查询功能，帮助用户更好地了解和分析 GitHub 上的开源项目和开发者的数据表现。通过这种交互式的查询方式，用户可以快速获取所需的数据结果，并进行进一步的分析和研究。

### 1.2 命令行框架 Cobra

本次项目基于Cobra框架进行开发，Cobra 是一个 Go 语言开发的命令行框架，它提供了**简洁、灵活且强大**的方式来创建命令行程序。它包含一个用于创建命令行程序的库，以及一个用于快速生成基于 Cobra 库的命令行程序工具。Cobra具有以下特性和优势：

- 简单易用：Cobra提供了直观的API和清晰的设计，使得构建命令行应用程序变得简单而直观。开发者可以快速定义命令、子命令和相关的标志和参数。
- 子命令支持：Cobra支持创建具有层次结构的命令行应用程序，即主命令下可以有多个子命令。这种结构使得应用程序的命令组织更清晰、更易扩展。
- 命令行补全：Cobra支持命令行补全功能，用户可以通过按Tab键自动补全命令、子命令、标志和参数，提高了交互性和用户体验

## 2. Usage

Exciting-Opendigger提供了多样的查询功能。为了更好地满足精细化的查询需求，我们给出了一系列的命令选项，其整体语法如下：

![USAGESTMT.svg](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/USAGESTMT.svg)

通过这一系列命令，Exciting-Opendigger满足了开源爱好者在查询贡献指标上的不同需求。如下表所示：

| 模块 | 命令 | 结果 | 需求 |
| --- | --- | --- | --- |
| 单点查询 | SHOW | 单个仓库或开发者在特定指标、特定月份的数据或整体报告 | 对开源指标的简单浏览 |
| 比较查询 | COMPARE | 对两个仓库、开发者、月份的比较结果 | 对开源指标的差异分析 |
| 文件下载 | DOWNLOAD | 查询结果的文件和可视化分析 | 对开源指标的详细分析和可视化 |
| 批量分析 | BATCH | 批量对象的数据 | 对Github社区的宏观分析 |
| 日志查询 | LOG | 历史查询日志的查看和回放 | 连续使用本工具的便捷性 |
| 帮助菜单 | HELP | 使用说明 | 辅助工具的便捷使用 |

### 2.1 **基础查询**

**2.1.1 单点查询**

其中，`SHOW`命令提供了基础的查询功能，即对特定对象的指标查询。此处我们提供一系列的查询参数如下：

![SearchOpt.svg](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/SearchOpt.svg)

在基础查询中，我们支持对仓库和开发者两个维度的指标查询，并且可以使用两个参数，`month`和`metric`，约束查询对应的月份和指标。用户需要至少提供一个参数以约束查询结果。选择不同参数约束的结果如下：

| month | metric | 结果 |
| --- | --- | --- |
| 填写 | 填写 | 特定指标在特定月份上的结果 |
| 填写 | 忽略 | 特定月份上的整体报告 |
| 忽略 | 填写 | 特定指标的长期趋势 |

> 开发日志：该模块需要与网络交互，访问OpenDigger提供的网络API接口。这部分已经基本完成。
> 

### 2.2 **指标比较**

除了单个仓库的查询，Exciting-Opendigger也提供多个对象间的比较。当启用`COMPARE`子命令时，开源爱好者可以通过追加新的对象或月份的方式比较两个指标数据的差异。

> 开发日志：该模块可以复用单点查询的接口完成，正在开发中。
> 

### 2.3 文件**下载**

Exciting-Opendigger支持使用`DOWNLOAD`命令下载基础查询的结果。下载需要提供文件的输出位置。同时，考虑到数据下载往往用于详细分析，我们为下载提供`DRAW`选项，用于**追加可视化的数据分析，**可以选择不同的画图方法**。**文件下载支持的格式有html、PDF以及csv，其中PDF文件由html转化。

![DownloadClause.svg](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/DownloadClause.svg)

> 开发日志：该模块基于html模板生成报告，并基于`signintech/gopdf`转化为PDF文件。这部分已经基本完成。
> 

### 2.4 批量分析

Exciting-Opendigger为开源爱好者进一步提供了批量分析的功能，从而鸟瞰Github社区的整体情况。

![BatchClause.svg](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/BatchClause.svg)

当使用`BATCH`命令进行批量分析时，用户可以请求两种不同的数据来源。首先，可以使用`TOP`参数请求查询当前Github最活跃仓库和开发者；其次，也可以采用提供文件的方式，重点关注特定的仓库或开发者。由于批量分析的数据量往往较多，直接输出到屏幕不利于分析数据。我们同样提供了文件下载的方式提供查询的结果。除了简单的批量查询，Exciting-Opendigger也支持对数据进行初步地过滤和排序处理，助力后续的专业数据分析。

![BatchOpt.svg](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/BatchOpt.svg)

Exciting-Opendigger支持对特定指标进行过滤，以及依据指标或月份进行排序。

> 开发日志：该模块可以复用单点查询的接口完成，为了提高查询效率，可以利用`goroutine`进行并发查询，这部分准备开始开发。
> 

### 2.5 日志查询

为了提供更为连续的操作体验，Exciting-Opendigger支持使用`CACHE`参数导出命令的历史日志到文件。随后，用户可以使用`LOG`命令查询之前的指标查询历史和结果。首先，用户可以使用`DISPLAY`参数展示被导出的一系列历史日志；然后，用户可以用日志号查询之前的查询命令，并回放对应的结果。

![LogClause.svg](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/LogClause.svg)

> 开发日志：Exciting-Opendigger作为非侵入用户系统的工具，计算通过SQLIite3缓存查询日志，避免依赖于操作系统的缓存设计。因此用户可以自行指定存储文件的位置，拥有对使用数据的完全管理权限，这部分已完成前期方案调研，准备开始开发。
> 

### 2.6 帮助菜单

由于Exciting-Opendigger提供了多样化的查询服务，其查询命令和参数较为详细。为便于用户使用这一工具，我们提供了全面的帮助菜单以展示可选操作和参数，解除了用户的记忆压力，降低使用门槛。

> 开发日志：该部分基于`Cobra`架构生成，已完成。
> 

## 3. Demo

命令行接口展示，此处展示了工具入口的帮助菜单，显示了Exciting-Opendigger提供的命令类型和用法说明。

![Untitled](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/Untitled%201.png)

html版本报告demo：

输出的报告与用户需求有关，将查询的数据进行格式整理以html格式输出，查询的指标和参数可以灵活设置，并且支持排序和筛选及其他自定义输出。

下图的demo展示了X-lab2017/open-digger项目的一个分析报告，这里以openrank指标为例，输出了X-lab2017/open-digger项目近五个月的openrank值，并以柱状图的形式输出了X-lab2017/open-digger项目立项以来openrank的月份变化趋势。

![report_html.jpg](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/report_html.jpg)

PDF版本报告demo：

即上方html版本报告的PDF版本

![report_pdf.jpg](https://github.com/wengsy150943/OpenSODAExcitingT2/blob/main/fig/report_pdf.jpg)

## 4.项目成员和分工

- 翁思扬   交互模块
- 钱堃       输出模块
- 梁辉       查询模块
- 孙印政   查询模块

## 5.项目依赖
### 5.1 wkhtmltox
项目中用到了go-wkhtmltopdf这个库，这个库依赖于wkhtmltox，Linux以及windows的安装方法见如下链接
链接：https://github.com/adrg/go-wkhtmltopdf