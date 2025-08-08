# GitHub MCP 服务器

GitHub MCP 服务器将 AI 工具直接连接到 GitHub 平台。这使得 AI
代理、助手和聊天机器人能够读取仓库和代码文件、管理问题和
PR、分析代码并自动化工作流程。所有操作都通过自然语言交互完成。

### 使用场景

- **仓库管理**：浏览和查询代码、搜索文件、分析提交记录，并了解您有权访问的任何仓库的项目结构。
- **问题与 PR 自动化**：创建、更新和管理问题与拉取请求。让 AI
  帮助分类错误、审查代码更改和维护项目面板。
- **CI/CD 与工作流程智能**：监控 GitHub Actions
  工作流程运行、分析构建失败、管理发布，并获取开发流水线的洞察。
- **代码分析**：检查安全发现、审查 Dependabot
  警报、理解代码模式，并获取代码库的全面洞察。
- **团队协作**：访问讨论、管理通知、分析团队活动，并为您的团队简化流程。

为希望将 AI 工具连接到 GitHub
上下文和功能的开发者构建，从简单的自然语言查询到复杂的多步骤代理工作流程。

---

## 远程 GitHub MCP 服务器

[![在 VS Code 中安装](https://img.shields.io/badge/VS_Code-安装服务器-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&config=%7B%22type%22%3A%20%22http%22%2C%22url%22%3A%20%22https%3A%2F%2Fapi.githubcopilot.com%2Fmcp%2F%22%7D)
[![在 VS Code Insiders 中安装](https://img.shields.io/badge/VS_Code_Insiders-安装服务器-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&config=%7B%22type%22%3A%20%22http%22%2C%22url%22%3A%20%22https%3A%2F%2Fapi.githubcopilot.com%2Fmcp%2F%22%7D&quality=insiders)

远程 GitHub MCP 服务器由 GitHub 托管，提供最简单的入门方法。如果您的 MCP
主机不支持远程 MCP
服务器，不用担心！您可以使用[本地版本的 GitHub MCP 服务器](https://gitee.com/masx200/github-mcp-server?tab=readme-ov-file#local-github-mcp-server)。

### 前提条件

1. 支持远程服务器的兼容 MCP 主机（VS Code 1.101+、Claude
   Desktop、Cursor、Windsurf 等）
2. 任何适用的[策略已启用](https://gitee.com/masx200/github-mcp-server/blob/main/docs/policies-and-governance.md)

### 在 VS Code 中安装

快速安装，请使用上方的一键安装按钮之一。完成该流程后，切换代理模式（位于 Copilot
Chat
文本输入框旁），服务器将启动。确保您使用[VS Code 1.101](https://code.visualstudio.com/updates/v1_101)或更高版本以获得远程
MCP 和 OAuth 支持。

或者，要手动配置 VS Code，请从以下示例中选择适当的 JSON
块并添加到您的主机配置中：

<table>
<tr><th>使用 OAuth</th><th>使用 GitHub PAT</th></tr>
<tr><th align=left colspan=2>VS Code（1.101 或更高版本）</th></tr>
<tr valign=top>
<td>

```json
{
  "servers": {
    "github": {
      "type": "http",
      "url": "https://api.githubcopilot.com/mcp/"
    }
  }
}
```

</td>
<td>

```json
{
  "servers": {
    "github": {
      "type": "http",
      "url": "https://api.githubcopilot.com/mcp/",
      "headers": {
        "Authorization": "Bearer ${input:github_mcp_pat}"
      }
    }
  },
  "inputs": [
    {
      "type": "promptString",
      "id": "github_mcp_pat",
      "description": "GitHub 个人访问令牌",
      "password": true
    }
  ]
}
```

</td>
</tr>
</table>

### 在其他 MCP 主机中安装

- **[其他 IDE 中的 GitHub Copilot](/docs/installation-guides/install-other-copilot-ides.md)** -
  JetBrains、Visual Studio、Eclipse 和 Xcode 的 GitHub Copilot 安装
- **[Claude 应用程序](/docs/installation-guides/install-claude.md)** - Claude
  Web、Claude Desktop 和 Claude Code CLI 的安装指南
- **[Cursor](/docs/installation-guides/install-cursor.md)** - Cursor IDE
  的安装指南
- **[Windsurf](/docs/installation-guides/install-windsurf.md)** - Windsurf IDE
  的安装指南

> **注意：** 每个 MCP 主机应用程序需要配置 GitHub App 或 OAuth App 以支持通过
> OAuth 的远程访问。任何支持远程 MCP 服务器的主机应用程序都应该支持使用 PAT
> 认证的远程 GitHub
> 服务器。配置细节和支持级别因主机而异。请务必参考主机应用程序的文档获取更多信息。

> ⚠️ **公开预览状态：** **远程** GitHub MCP
> 服务器目前处于公开预览阶段。在预览期间，访问可能根据认证类型和界面进行限制：
>
> - OAuth：在 GA 之前受 GitHub Copilot 编辑器预览策略约束
> - PAT：通过您组织的 PAT 策略控制
> - MCP 服务器在 Copilot 策略中：启用/禁用 VS Code 中的所有 MCP 服务器访问，其他
>   Copilot 编辑器将在未来几个月迁移到此策略

### 配置

有关如何向远程 GitHub MCP
服务器传递额外配置设置的信息，请参阅[远程服务器文档](/docs/remote-server.md)。

---

## 本地 GitHub MCP 服务器

[![在 VS Code 中使用 Docker 安装](https://img.shields.io/badge/VS_Code-安装服务器-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&inputs=%5B%7B%22id%22%3A%22github_token%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22GitHub%20个人访问令牌%22%2C%22password%22%3Atrue%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-e%22%2C%22GITHUB_PERSONAL_ACCESS_TOKEN%22%2C%22ghcr.io%2Fgithub%2Fgithub-mcp-server%22%5D%2C%22env%22%3A%7B%22GITHUB_PERSONAL_ACCESS_TOKEN%22%3A%22%24%7Binput%3Agithub_token%7D%22%7D%7D)
[![在 VS Code Insiders 中使用 Docker 安装](https://img.shields.io/badge/VS_Code_Insiders-安装服务器-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&inputs=%5B%7B%22id%22%3A%22github_token%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22GitHub%20个人访问令牌%22%2C%22password%22%3Atrue%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-e%22%2C%22GITHUB_PERSONAL_ACCESS_TOKEN%22%2C%22ghcr.io%2Fgithub%2Fgithub-mcp-server%22%5D%2C%22env%22%3A%7B%22GITHUB_PERSONAL_ACCESS_TOKEN%22%3A%22%24%7Binput%3Agithub_token%7D%22%7D%7D&quality=insiders)

### 前提条件

1. 要在容器中运行服务器，您需要安装 [Docker](https://www.docker.com/)。
2. 安装 Docker 后，您还需要确保 Docker
   正在运行。镜像是公开的；如果在拉取时遇到错误，您可能有过期的令牌，需要
   `docker logout ghcr.io`。
3. 最后，您需要[创建 GitHub 个人访问令牌](https://github.com/settings/personal-access-tokens/new)。MCP
   服务器可以使用许多 GitHub API，因此启用您愿意授予 AI
   工具的权限（要了解更多关于访问令牌的信息，请查看[文档](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)）。

<details><summary><b>安全处理 PAT</b></summary>

### 环境变量（推荐）

为了保持您的 GitHub PAT 安全并在不同的 MCP 主机中重复使用：

1. **将您的 PAT 存储在环境变量中**

   ```bash
   export GITHUB_PAT=your_token_here
   ```

   或创建 `.env` 文件：

   ```env
   GITHUB_PAT=your_token_here
   ```

2. **保护您的 `.env` 文件**

   ```bash
   # 添加到 .gitignore 以防止意外提交
   echo ".env" >> .gitignore
   ```

3. **在配置中引用令牌**

   ```bash
   # CLI 使用
   claude mcp update github -e GITHUB_PERSONAL_ACCESS_TOKEN=$GITHUB_PAT

   # 在配置文件中（在支持的地方）
   "env": {
     "GITHUB_PERSONAL_ACCESS_TOKEN": "$GITHUB_PAT"
   }
   ```

> **注意**：环境变量支持因主机应用和 IDE 而异。某些应用程序（如
> Windsurf）需要在配置文件中硬编码令牌。

### 令牌安全最佳实践

- **最小权限**：仅授予必要的权限
  - `repo` - 仓库操作
  - `read:packages` - Docker 镜像访问
- **分离令牌**：为不同项目/环境使用不同的 PAT
- **定期轮换**：定期更新令牌
- **永不提交**：将令牌排除在版本控制之外
- **文件权限**：限制对包含令牌的配置文件的访问
  ```bash
  chmod 600 ~/.your-app/config.json
  ```

</details>

## 安装

### 在 GitHub Copilot on VS Code 中安装

快速安装，请使用上方的一键安装按钮之一。完成该流程后，切换代理模式（位于 Copilot
Chat 文本输入框旁），服务器将启动。

更多关于在 VS Code 中使用 MCP
服务器工具的信息，请参阅[代理模式文档](https://code.visualstudio.com/docs/copilot/chat/mcp-servers)。

在其他 IDE（JetBrains、Visual Studio、Eclipse 等）的 GitHub Copilot 中安装

将以下 JSON 块添加到您 IDE 的 MCP 设置中。

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "github_token",
        "description": "GitHub 个人访问令牌",
        "password": true
      }
    ],
    "servers": {
      "github": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "-e",
          "GITHUB_PERSONAL_ACCESS_TOKEN",
          "ghcr.io/github/github-mcp-server"
        ],
        "env": {
          "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}"
        }
      }
    }
  }
}
```

可选地，您可以将类似的示例（即不包含 mcp 键）添加到工作区中的 `.vscode/mcp.json`
文件中。这将允许您与其他接受相同格式的主机应用程序共享配置。

<details>
<summary><b>不包含 MCP 键的示例 JSON 块</b></summary>
<br>

```json
{
  "inputs": [
    {
      "type": "promptString",
      "id": "github_token",
      "description": "GitHub 个人访问令牌",
      "password": true
    }
  ],
  "servers": {
    "github": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e",
        "GITHUB_PERSONAL_ACCESS_TOKEN",
        "ghcr.io/github/github-mcp-server"
      ],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}"
      }
    }
  }
}
```

</details>

### 在其他 MCP 主机中安装

对于其他 MCP 主机应用程序，请参阅我们的安装指南：

- **[其他 IDE 中的 GitHub Copilot](/docs/installation-guides/install-other-copilot-ides.md)** -
  JetBrains、Visual Studio、Eclipse 和 Xcode 的 GitHub Copilot 安装
- **[Claude Code 和 Claude Desktop](docs/installation-guides/install-claude.md)** -
  Claude Code 和 Claude Desktop 的安装指南
- **[Cursor](/docs/installation-guides/install-cursor.md)** - Cursor IDE
  的安装指南
- **[Windsurf](/docs/installation-guides/install-windsurf.md)** - Windsurf IDE
  的安装指南

有关所有安装选项的完整概述，请参阅我们的**[安装指南索引](docs/installation-guides/installation-guides.md)**。

> **注意：** 任何支持本地 MCP 服务器的主机应用程序都应该能够访问本地 GitHub MCP
> 服务器。但是，具体的配置过程、语法和集成的稳定性将因主机应用程序而异。虽然许多可能遵循与上述示例类似的格式，但这不能保证。请参考您的主机应用程序的文档以获取正确的
> MCP 配置语法和设置过程。

### 从源代码构建

如果您没有 Docker，您可以使用 `go build` 在 `cmd/github-mcp-server`
目录中构建二进制文件，并使用 `github-mcp-server stdio` 命令，并将
`GITHUB_PERSONAL_ACCESS_TOKEN`
环境变量设置为您的令牌。要指定构建的输出位置，请使用 `-o`
标志。您应该配置您的服务器以使用构建的可执行文件作为其 `command`。例如：

```JSON
{
  "mcp": {
    "servers": {
      "github": {
        "command": "/path/to/github-mcp-server",
        "args": ["stdio"],
        "env": {
          "GITHUB_PERSONAL_ACCESS_TOKEN": "<YOUR_TOKEN>"
        }
      }
    }
  }
}
```

## 使用 streamable-http 模式启动服务器

您还可以使用 streamable-http 模式运行服务器：

```bash
go run -v ./main.go http --address ":38888"
```

或者

```bash
go build -v ./main.go

./main.exe http --address ":38888"
```

这将在端口 38888 上使用 streamable-http 传输启动服务器。

### 日志格式控制

服务器支持通过 `--pretty` 命令行参数控制日志输出格式：

- `--pretty=false`：输出压缩的 JSON 格式日志（单行，无缩进）
- `--pretty=true`：输出美化的 JSON 格式日志（多行，带缩进和换行）

示例：

```bash
# 使用压缩 JSON 日志格式
go run -v ./main.go http --address ":38888" --pretty=false

# 使用美化 JSON 日志格式（默认）
go run -v ./main.go http --address ":38888" --pretty=true
```

## 工具配置

GitHub MCP 服务器支持通过 `--toolsets` 标志启用或禁用特定功能组。这允许您控制 AI
工具可用的 GitHub API 功能。仅启用您需要的工具集可以帮助 LLM
进行工具选择并减少上下文大小。

_工具集不仅限于工具。相关的 MCP 资源和提示也包含在适用的地方。_

### 可用工具集

以下工具集可用（默认全部开启）：

<!-- START AUTOMATED TOOLSETS -->

| 工具集              | 描述                                                             |
| ------------------- | ---------------------------------------------------------------- |
| `context`           | **强烈推荐**：提供关于当前用户和您正在操作的 GitHub 上下文的工具 |
| `actions`           | GitHub Actions 工作流程和 CI/CD 操作                             |
| `code_security`     | 代码安全相关工具，如 GitHub 代码扫描                             |
| `dependabot`        | Dependabot 工具                                                  |
| `discussions`       | GitHub 讨论相关工具                                              |
| `experiments`       | 尚未被视为稳定的实验性功能                                       |
| `gists`             | GitHub Gist 相关工具                                             |
| `issues`            | GitHub 问题相关工具                                              |
| `notifications`     | GitHub 通知相关工具                                              |
| `orgs`              | GitHub 组织相关工具                                              |
| `pull_requests`     | GitHub 拉取请求相关工具                                          |
| `repos`             | GitHub 仓库相关工具                                              |
| `secret_protection` | 密钥保护相关工具，如 GitHub 密钥扫描                             |
| `users`             | GitHub 用户相关工具                                              |

<!-- END AUTOMATED TOOLSETS -->

## 工具

<!-- START AUTOMATED TOOLS -->
<details>

<summary>操作</summary>

- **cancel_workflow_run** - 取消工作流程运行

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **delete_workflow_run_logs** - 删除工作流程日志

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **download_workflow_run_artifact** - 下载工作流程构件

  - `artifact_id`: 构件的唯一标识符（数字，必需）
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）

- **get_job_logs** - 获取作业日志

  - `failed_only`: 当为 true 时，获取 run_id
    中所有失败作业的日志（布尔值，可选）
  - `job_id`: 工作流程作业的唯一标识符（单个作业日志必需）（数字，可选）
  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `return_content`: 返回实际日志内容而不是 URL（布尔值，可选）
  - `run_id`: 工作流程运行 ID（使用 failed_only 时必需）（数字，可选）
  - `tail_lines`: 从日志末尾返回的行数（数字，可选）

- **get_workflow_run** - 获取工作流程运行

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **get_workflow_run_logs** - 获取工作流程运行日志

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **get_workflow_run_usage** - 获取工作流程使用情况

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **list_workflow_jobs** - 列出工作流程作业

  - `filter`: 根据 completed_at 时间戳过滤作业（字符串，可选）
  - `owner`: 仓库所有者（字符串，必需）
  - `page`: 分页页码（最小 1）（数字，可选）
  - `perPage`: 分页每页结果数（最小 1，最大 100）（数字，可选）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **list_workflow_run_artifacts** - 列出工作流程构件

  - `owner`: 仓库所有者（字符串，必需）
  - `page`: 分页页码（最小 1）（数字，可选）
  - `perPage`: 分页每页结果数（最小 1，最大 100）（数字，可选）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **list_workflow_runs** - 列出工作流程运行

  - `actor`:
    返回某人的工作流程运行。使用创建工作流程运行的用户的登录名。（字符串，可选）
  - `branch`: 返回与分支关联的工作流程运行。使用分支名称。（字符串，可选）
  - `event`: 返回特定事件类型的工作流程运行（字符串，可选）
  - `owner`: 仓库所有者（字符串，必需）
  - `page`: 分页页码（最小 1）（数字，可选）
  - `perPage`: 分页每页结果数（最小 1，最大 100）（数字，可选）
  - `repo`: 仓库名称（字符串，必需）
  - `status`: 返回具有检查运行状态的工作流程运行（字符串，可选）
  - `workflow_id`: 工作流程 ID 或工作流程文件名（字符串，必需）

- **list_workflows** - 列出工作流程

  - `owner`: 仓库所有者（字符串，必需）
  - `page`: 分页页码（最小 1）（数字，可选）
  - `perPage`: 分页每页结果数（最小 1，最大 100）（数字，可选）
  - `repo`: 仓库名称（字符串，必需）

- **rerun_failed_jobs** - 重新运行失败的作业

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **rerun_workflow_run** - 重新运行工作流程

  - `owner`: 仓库所有者（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `run_id`: 工作流程运行的唯一标识符（数字，必需）

- **run_workflow** - 运行工作流程
  - `inputs`: 工作流程接受的输入（对象，可选）
  - `owner`: 仓库所有者（字符串，必需）
  - `ref`: 工作流程的 git 引用。引用可以是分支或标签名称。（字符串，必需）
  - `repo`: 仓库名称（字符串，必需）
  - `workflow_id`: 工作流程 ID 或工作流程文件名（字符串，必需）

</details>
# GitHub MCP Server

The GitHub MCP Server connects AI tools directly to GitHub's platform. This
gives AI agents, assistants, and chatbots the ability to read repositories and
code files, manage issues and PRs, analyze code, and automate workflows. All
through natural language interactions.

### Use Cases

- Repository Management: Browse and query code, search files, analyze commits,
  and understand project structure across any repository you have access to.
- Issue & PR Automation: Create, update, and manage issues and pull requests.
  Let AI help triage bugs, review code changes, and maintain project boards.
- CI/CD & Workflow Intelligence: Monitor GitHub Actions workflow runs, analyze
  build failures, manage releases, and get insights into your development
  pipeline.
- Code Analysis: Examine security findings, review Dependabot alerts, understand
  code patterns, and get comprehensive insights into your codebase.
- Team Collaboration: Access discussions, manage notifications, analyze team
  activity, and streamline processes for your team.

Built for developers who want to connect their AI tools to GitHub context and
capabilities, from simple natural language queries to complex multi-step agent
workflows.

---

## Remote GitHub MCP Server

[![Install in VS Code](https://img.shields.io/badge/VS_Code-Install_Server-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&config=%7B%22type%22%3A%20%22http%22%2C%22url%22%3A%20%22https%3A%2F%2Fapi.githubcopilot.com%2Fmcp%2F%22%7D)
[![Install in VS Code Insiders](https://img.shields.io/badge/VS_Code_Insiders-Install_Server-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&config=%7B%22type%22%3A%20%22http%22%2C%22url%22%3A%20%22https%3A%2F%2Fapi.githubcopilot.com%2Fmcp%2F%22%7D&quality=insiders)

The remote GitHub MCP Server is hosted by GitHub and provides the easiest method
for getting up and running. If your MCP host does not support remote MCP
servers, don't worry! You can use the
[local version of the GitHub MCP Server](https://gitee.com/masx200/github-mcp-server?tab=readme-ov-file#local-github-mcp-server)
instead.

### Prerequisites

1. A compatible MCP host with remote server support (VS Code 1.101+, Claude
   Desktop, Cursor, Windsurf, etc.)
2. Any applicable
   [policies enabled](https://gitee.com/masx200/github-mcp-server/blob/main/docs/policies-and-governance.md)

### Install in VS Code

For quick installation, use one of the one-click install buttons above. Once you
complete that flow, toggle Agent mode (located by the Copilot Chat text input)
and the server will start. Make sure you're using
[VS Code 1.101](https://code.visualstudio.com/updates/v1_101) or
[later](https://code.visualstudio.com/updates) for remote MCP and OAuth support.

Alternatively, to manually configure VS Code, choose the appropriate JSON block
from the examples below and add it to your host configuration:

<table>
<tr><th>Using OAuth</th><th>Using a GitHub PAT</th></tr>
<tr><th align=left colspan=2>VS Code (version 1.101 or greater)</th></tr>
<tr valign=top>
<td>

```json
{
  "servers": {
    "github": {
      "type": "http",
      "url": "https://api.githubcopilot.com/mcp/"
    }
  }
}
```

</td>
<td>

```json
{
  "servers": {
    "github": {
      "type": "http",
      "url": "https://api.githubcopilot.com/mcp/",
      "headers": {
        "Authorization": "Bearer ${input:github_mcp_pat}"
      }
    }
  },
  "inputs": [
    {
      "type": "promptString",
      "id": "github_mcp_pat",
      "description": "GitHub Personal Access Token",
      "password": true
    }
  ]
}
```

</td>
</tr>
</table>

### Install in other MCP hosts

- **[GitHub Copilot in other IDEs](/docs/installation-guides/install-other-copilot-ides.md)** -
  Installation for JetBrains, Visual Studio, Eclipse, and Xcode with GitHub
  Copilot
- **[Claude Applications](/docs/installation-guides/install-claude.md)** -
  Installation guide for Claude Web, Claude Desktop and Claude Code CLI
- **[Cursor](/docs/installation-guides/install-cursor.md)** - Installation guide
  for Cursor IDE
- **[Windsurf](/docs/installation-guides/install-windsurf.md)** - Installation
  guide for Windsurf IDE

> **Note:** Each MCP host application needs to configure a GitHub App or OAuth
> App to support remote access via OAuth. Any host application that supports
> remote MCP servers should support the remote GitHub server with PAT
> authentication. Configuration details and support levels vary by host. Make
> sure to refer to the host application's documentation for more info.

> ⚠️ **Public Preview Status:** The **remote** GitHub MCP Server is currently in
> Public Preview. During preview, access may be gated depending on
> authentication type and surface:
>
> - OAuth: Subject to GitHub Copilot Editor Preview Policy until GA
> - PAT: Controlled via your organization's PAT policies
> - MCP Servers in Copilot policy: Enables/disables access to all MCP servers in
>   VS Code, with other Copilot editors migrating to this policy in the coming
>   months.

### Configuration

See [Remote Server Documentation](/docs/remote-server.md) on how to pass
additional configuration settings to the remote GitHub MCP Server.

---

## Local GitHub MCP Server

[![Install with Docker in VS Code](https://img.shields.io/badge/VS_Code-Install_Server-0098FF?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&inputs=%5B%7B%22id%22%3A%22github_token%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22GitHub%20Personal%20Access%20Token%22%2C%22password%22%3Atrue%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-e%22%2C%22GITHUB_PERSONAL_ACCESS_TOKEN%22%2C%22ghcr.io%2Fgithub%2Fgithub-mcp-server%22%5D%2C%22env%22%3A%7B%22GITHUB_PERSONAL_ACCESS_TOKEN%22%3A%22%24%7Binput%3Agithub_token%7D%22%7D%7D)
[![Install with Docker in VS Code Insiders](https://img.shields.io/badge/VS_Code_Insiders-Install_Server-24bfa5?style=flat-square&logo=visualstudiocode&logoColor=white)](https://insiders.vscode.dev/redirect/mcp/install?name=github&inputs=%5B%7B%22id%22%3A%22github_token%22%2C%22type%22%3A%22promptString%22%2C%22description%22%3A%22GitHub%20Personal%20Access%20Token%22%2C%22password%22%3Atrue%7D%5D&config=%7B%22command%22%3A%22docker%22%2C%22args%22%3A%5B%22run%22%2C%22-i%22%2C%22--rm%22%2C%22-e%22%2C%22GITHUB_PERSONAL_ACCESS_TOKEN%22%2C%22ghcr.io%2Fgithub%2Fgithub-mcp-server%22%5D%2C%22env%22%3A%7B%22GITHUB_PERSONAL_ACCESS_TOKEN%22%3A%22%24%7Binput%3Agithub_token%7D%22%7D%7D&quality=insiders)

### Prerequisites

1. To run the server in a container, you will need to have
   [Docker](https://www.docker.com/) installed.
2. Once Docker is installed, you will also need to ensure Docker is running. The
   image is public; if you get errors on pull, you may have an expired token and
   need to `docker logout ghcr.io`.
3. Lastly you will need to
   [Create a GitHub Personal Access Token](https://github.com/settings/personal-access-tokens/new).
   The MCP server can use many of the GitHub APIs, so enable the permissions
   that you feel comfortable granting your AI tools (to learn more about access
   tokens, please check out the
   [documentation](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens)).

<details><summary><b>Handling PATs Securely</b></summary>

### Environment Variables (Recommended)

To keep your GitHub PAT secure and reusable across different MCP hosts:

1. **Store your PAT in environment variables**

   ```bash
   export GITHUB_PAT=your_token_here
   ```

   Or create a `.env` file:

   ```env
   GITHUB_PAT=your_token_here
   ```

2. **Protect your `.env` file**

   ```bash
   # Add to .gitignore to prevent accidental commits
   echo ".env" >> .gitignore
   ```

3. **Reference the token in configurations**

   ```bash
   # CLI usage
   claude mcp update github -e GITHUB_PERSONAL_ACCESS_TOKEN=$GITHUB_PAT

   # In config files (where supported)
   "env": {
     "GITHUB_PERSONAL_ACCESS_TOKEN": "$GITHUB_PAT"
   }
   ```

> **Note**: Environment variable support varies by host app and IDE. Some
> applications (like Windsurf) require hardcoded tokens in config files.

### Token Security Best Practices

- **Minimum scopes**: Only grant necessary permissions
  - `repo` - Repository operations
  - `read:packages` - Docker image access
- **Separate tokens**: Use different PATs for different projects/environments
- **Regular rotation**: Update tokens periodically
- **Never commit**: Keep tokens out of version control
- **File permissions**: Restrict access to config files containing tokens
  ```bash
  chmod 600 ~/.your-app/config.json
  ```

</details>

## Installation

### Install in GitHub Copilot on VS Code

For quick installation, use one of the one-click install buttons above. Once you
complete that flow, toggle Agent mode (located by the Copilot Chat text input)
and the server will start.

More about using MCP server tools in VS Code's
[agent mode documentation](https://code.visualstudio.com/docs/copilot/chat/mcp-servers).

Install in GitHub Copilot on other IDEs (JetBrains, Visual Studio, Eclipse,
etc.)

Add the following JSON block to your IDE's MCP settings.

```json
{
  "mcp": {
    "inputs": [
      {
        "type": "promptString",
        "id": "github_token",
        "description": "GitHub Personal Access Token",
        "password": true
      }
    ],
    "servers": {
      "github": {
        "command": "docker",
        "args": [
          "run",
          "-i",
          "--rm",
          "-e",
          "GITHUB_PERSONAL_ACCESS_TOKEN",
          "ghcr.io/github/github-mcp-server"
        ],
        "env": {
          "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}"
        }
      }
    }
  }
}
```

Optionally, you can add a similar example (i.e. without the mcp key) to a file
called `.vscode/mcp.json` in your workspace. This will allow you to share the
configuration with other host applications that accept the same format.

<details>
<summary><b>Example JSON block without the MCP key included</b></summary>
<br>

```json
{
  "inputs": [
    {
      "type": "promptString",
      "id": "github_token",
      "description": "GitHub Personal Access Token",
      "password": true
    }
  ],
  "servers": {
    "github": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e",
        "GITHUB_PERSONAL_ACCESS_TOKEN",
        "ghcr.io/github/github-mcp-server"
      ],
      "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}"
      }
    }
  }
}
```

</details>

### Install in Other MCP Hosts

For other MCP host applications, please refer to our installation guides:

- **[GitHub Copilot in other IDEs](/docs/installation-guides/install-other-copilot-ides.md)** -
  Installation for JetBrains, Visual Studio, Eclipse, and Xcode with GitHub
  Copilot
- **[Claude Code & Claude Desktop](docs/installation-guides/install-claude.md)** -
  Installation guide for Claude Code and Claude Desktop
- **[Cursor](docs/installation-guides/install-cursor.md)** - Installation guide
  for Cursor IDE
- **[Windsurf](docs/installation-guides/install-windsurf.md)** - Installation
  guide for Windsurf IDE

For a complete overview of all installation options, see our
**[Installation Guides Index](docs/installation-guides/installation-guides.md)**.

> **Note:** Any host application that supports local MCP servers should be able
> to access the local GitHub MCP server. However, the specific configuration
> process, syntax and stability of the integration will vary by host
> application. While many may follow a similar format to the examples above,
> this is not guaranteed. Please refer to your host application's documentation
> for the correct MCP configuration syntax and setup process.

### Build from source

If you don't have Docker, you can use `go build` to build the binary in the
`cmd/github-mcp-server` directory, and use the `github-mcp-server stdio` command
with the `GITHUB_PERSONAL_ACCESS_TOKEN` environment variable set to your token.
To specify the output location of the build, use the `-o` flag. You should
configure your server to use the built executable as its `command`. For example:

```JSON
{
  "mcp": {
    "servers": {
      "github": {
        "command": "/path/to/github-mcp-server",
        "args": ["stdio"],
        "env": {
          "GITHUB_PERSONAL_ACCESS_TOKEN": "<YOUR_TOKEN>"
        }
      }
    }
  }
}
```

## start server using streamable-http mode

You can also run the server using streamable-http mode:

```bash
go run -v ./main.go http --address ":38888"
```

This will start the server on port 38888 using streamable-http transport.

## Tool Configuration

The GitHub MCP Server supports enabling or disabling specific groups of
functionalities via the `--toolsets` flag. This allows you to control which
GitHub API capabilities are available to your AI tools. Enabling only the
toolsets that you need can help the LLM with tool choice and reduce the context
size.

_Toolsets are not limited to Tools. Relevant MCP Resources and Prompts are also
included where applicable._

### Available Toolsets

The following sets of tools are available (all are on by default):

<!-- START AUTOMATED TOOLSETS -->

| Toolset             | Description                                                                                                         |
| ------------------- | ------------------------------------------------------------------------------------------------------------------- |
| `context`           | **Strongly recommended**: Tools that provide context about the current user and GitHub context you are operating in |
| `actions`           | GitHub Actions workflows and CI/CD operations                                                                       |
| `code_security`     | Code security related tools, such as GitHub Code Scanning                                                           |
| `dependabot`        | Dependabot tools                                                                                                    |
| `discussions`       | GitHub Discussions related tools                                                                                    |
| `experiments`       | Experimental features that are not considered stable yet                                                            |
| `gists`             | GitHub Gist related tools                                                                                           |
| `issues`            | GitHub Issues related tools                                                                                         |
| `notifications`     | GitHub Notifications related tools                                                                                  |
| `orgs`              | GitHub Organization related tools                                                                                   |
| `pull_requests`     | GitHub Pull Request related tools                                                                                   |
| `repos`             | GitHub Repository related tools                                                                                     |
| `secret_protection` | Secret protection related tools, such as GitHub Secret Scanning                                                     |
| `users`             | GitHub User related tools                                                                                           |

<!-- END AUTOMATED TOOLSETS -->

## Tools

<!-- START AUTOMATED TOOLS -->
<details>

<summary>Actions</summary>

- **cancel_workflow_run** - Cancel workflow run

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **delete_workflow_run_logs** - Delete workflow logs

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **download_workflow_run_artifact** - Download workflow artifact

  - `artifact_id`: The unique identifier of the artifact (number, required)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)

- **get_job_logs** - Get job logs

  - `failed_only`: When true, gets logs for all failed jobs in run_id (boolean,
    optional)
  - `job_id`: The unique identifier of the workflow job (required for single job
    logs) (number, optional)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `return_content`: Returns actual log content instead of URLs (boolean,
    optional)
  - `run_id`: Workflow run ID (required when using failed_only) (number,
    optional)
  - `tail_lines`: Number of lines to return from the end of the log (number,
    optional)

- **get_workflow_run** - Get workflow run

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **get_workflow_run_logs** - Get workflow run logs

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **get_workflow_run_usage** - Get workflow usage

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **list_workflow_jobs** - List workflow jobs

  - `filter`: Filters jobs by their completed_at timestamp (string, optional)
  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **list_workflow_run_artifacts** - List workflow artifacts

  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **list_workflow_runs** - List workflow runs

  - `actor`: Returns someone's workflow runs. Use the login for the user who
    created the workflow run. (string, optional)
  - `branch`: Returns workflow runs associated with a branch. Use the name of
    the branch. (string, optional)
  - `event`: Returns workflow runs for a specific event type (string, optional)
  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)
  - `status`: Returns workflow runs with the check run status (string, optional)
  - `workflow_id`: The workflow ID or workflow file name (string, required)

- **list_workflows** - List workflows

  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)

- **rerun_failed_jobs** - Rerun failed jobs

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **rerun_workflow_run** - Rerun workflow run

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `run_id`: The unique identifier of the workflow run (number, required)

- **run_workflow** - Run workflow
  - `inputs`: Inputs the workflow accepts (object, optional)
  - `owner`: Repository owner (string, required)
  - `ref`: The git reference for the workflow. The reference can be a branch or
    tag name. (string, required)
  - `repo`: Repository name (string, required)
  - `workflow_id`: The workflow ID (numeric) or workflow file name (e.g.,
    main.yml, ci.yaml) (string, required)

</details>

<details>

<summary>Code Security</summary>

- **get_code_scanning_alert** - Get code scanning alert

  - `alertNumber`: The number of the alert. (number, required)
  - `owner`: The owner of the repository. (string, required)
  - `repo`: The name of the repository. (string, required)

- **list_code_scanning_alerts** - List code scanning alerts
  - `owner`: The owner of the repository. (string, required)
  - `ref`: The Git reference for the results you want to list. (string,
    optional)
  - `repo`: The name of the repository. (string, required)
  - `severity`: Filter code scanning alerts by severity (string, optional)
  - `state`: Filter code scanning alerts by state. Defaults to open (string,
    optional)
  - `tool_name`: The name of the tool used for code scanning. (string, optional)

</details>

<details>

<summary>Context</summary>

- **get_me** - Get my user profile
  - No parameters required

</details>

<details>

<summary>Dependabot</summary>

- **get_dependabot_alert** - Get dependabot alert

  - `alertNumber`: The number of the alert. (number, required)
  - `owner`: The owner of the repository. (string, required)
  - `repo`: The name of the repository. (string, required)

- **list_dependabot_alerts** - List dependabot alerts
  - `owner`: The owner of the repository. (string, required)
  - `repo`: The name of the repository. (string, required)
  - `severity`: Filter dependabot alerts by severity (string, optional)
  - `state`: Filter dependabot alerts by state. Defaults to open (string,
    optional)

</details>

<details>

<summary>Discussions</summary>

- **get_discussion** - Get discussion

  - `discussionNumber`: Discussion Number (number, required)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)

- **get_discussion_comments** - Get discussion comments

  - `after`: Cursor for pagination. Use the endCursor from the previous page's
    PageInfo for GraphQL APIs. (string, optional)
  - `discussionNumber`: Discussion Number (number, required)
  - `owner`: Repository owner (string, required)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)

- **list_discussion_categories** - List discussion categories

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)

- **list_discussions** - List discussions
  - `after`: Cursor for pagination. Use the endCursor from the previous page's
    PageInfo for GraphQL APIs. (string, optional)
  - `category`: Optional filter by discussion category ID. If provided, only
    discussions with this category are listed. (string, optional)
  - `direction`: Order direction. (string, optional)
  - `orderBy`: Order discussions by field. If provided, the 'direction' also
    needs to be provided. (string, optional)
  - `owner`: Repository owner (string, required)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name. If not provided, discussions will be queried at the
    organisation level. (string, optional)

</details>

<details>

<summary>Gists</summary>

- **create_gist** - Create Gist

  - `content`: Content for simple single-file gist creation (string, required)
  - `description`: Description of the gist (string, optional)
  - `filename`: Filename for simple single-file gist creation (string, required)
  - `public`: Whether the gist is public (boolean, optional)

- **list_gists** - List Gists

  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `since`: Only gists updated after this time (ISO 8601 timestamp) (string,
    optional)
  - `username`: GitHub username (omit for authenticated user's gists) (string,
    optional)

- **update_gist** - Update Gist
  - `content`: Content for the file (string, required)
  - `description`: Updated description of the gist (string, optional)
  - `filename`: Filename to update or create (string, required)
  - `gist_id`: ID of the gist to update (string, required)

</details>

<details>

<summary>Issues</summary>

- **add_issue_comment** - Add comment to issue

  - `body`: Comment content (string, required)
  - `issue_number`: Issue number to comment on (number, required)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)

- **add_sub_issue** - Add sub-issue

  - `issue_number`: The number of the parent issue (number, required)
  - `owner`: Repository owner (string, required)
  - `replace_parent`: When true, replaces the sub-issue's current parent issue
    (boolean, optional)
  - `repo`: Repository name (string, required)
  - `sub_issue_id`: The ID of the sub-issue to add. ID is not the same as issue
    number (number, required)

- **assign_copilot_to_issue** - Assign Copilot to issue

  - `issueNumber`: Issue number (number, required)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)

- **create_issue** - Open new issue

  - `assignees`: Usernames to assign to this issue (string[], optional)
  - `body`: Issue body content (string, optional)
  - `labels`: Labels to apply to this issue (string[], optional)
  - `milestone`: Milestone number (number, optional)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `title`: Issue title (string, required)

- **get_issue** - Get issue details

  - `issue_number`: The number of the issue (number, required)
  - `owner`: The owner of the repository (string, required)
  - `repo`: The name of the repository (string, required)

- **get_issue_comments** - Get issue comments

  - `issue_number`: Issue number (number, required)
  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)

- **list_issues** - List issues

  - `direction`: Sort direction (string, optional)
  - `labels`: Filter by labels (string[], optional)
  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)
  - `since`: Filter by date (ISO 8601 timestamp) (string, optional)
  - `sort`: Sort order (string, optional)
  - `state`: Filter by state (string, optional)

- **list_sub_issues** - List sub-issues

  - `issue_number`: Issue number (number, required)
  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (default: 1) (number, optional)
  - `per_page`: Number of results per page (max 100, default: 30) (number,
    optional)
  - `repo`: Repository name (string, required)

- **remove_sub_issue** - Remove sub-issue

  - `issue_number`: The number of the parent issue (number, required)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `sub_issue_id`: The ID of the sub-issue to remove. ID is not the same as
    issue number (number, required)

- **reprioritize_sub_issue** - Reprioritize sub-issue

  - `after_id`: The ID of the sub-issue to be prioritized after (either after_id
    OR before_id should be specified) (number, optional)
  - `before_id`: The ID of the sub-issue to be prioritized before (either
    after_id OR before_id should be specified) (number, optional)
  - `issue_number`: The number of the parent issue (number, required)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `sub_issue_id`: The ID of the sub-issue to reprioritize. ID is not the same
    as issue number (number, required)

- **search_issues** - Search issues

  - `order`: Sort order (string, optional)
  - `owner`: Optional repository owner. If provided with repo, only
    notifications for this repository are listed. (string, optional)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `query`: Search query using GitHub issues search syntax (string, required)
  - `repo`: Optional repository name. If provided with owner, only notifications
    for this repository are listed. (string, optional)
  - `sort`: Sort field by number of matches of categories, defaults to best
    match (string, optional)

- **update_issue** - Edit issue
  - `assignees`: New assignees (string[], optional)
  - `body`: New description (string, optional)
  - `issue_number`: Issue number to update (number, required)
  - `labels`: New labels (string[], optional)
  - `milestone`: New milestone number (number, optional)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `state`: New state (string, optional)
  - `title`: New title (string, optional)

</details>

<details>

<summary>Notifications</summary>

- **dismiss_notification** - Dismiss notification

  - `state`: The new state of the notification (read/done) (string, optional)
  - `threadID`: The ID of the notification thread (string, required)

- **get_notification_details** - Get notification details

  - `notificationID`: The ID of the notification (string, required)

- **list_notifications** - List notifications

  - `before`: Only show notifications updated before the given time (ISO 8601
    format) (string, optional)
  - `filter`: Filter notifications to, use default unless specified. Read
    notifications are ones that have already been acknowledged by the user.
    Participating notifications are those that the user is directly involved in,
    such as issues or pull requests they have commented on or created. (string,
    optional)
  - `owner`: Optional repository owner. If provided with repo, only
    notifications for this repository are listed. (string, optional)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Optional repository name. If provided with owner, only notifications
    for this repository are listed. (string, optional)
  - `since`: Only show notifications updated after the given time (ISO 8601
    format) (string, optional)

- **manage_notification_subscription** - Manage notification subscription

  - `action`: Action to perform: ignore, watch, or delete the notification
    subscription. (string, required)
  - `notificationID`: The ID of the notification thread. (string, required)

- **manage_repository_notification_subscription** - Manage repository
  notification subscription

  - `action`: Action to perform: ignore, watch, or delete the repository
    notification subscription. (string, required)
  - `owner`: The account owner of the repository. (string, required)
  - `repo`: The name of the repository. (string, required)

- **mark_all_notifications_read** - Mark all notifications as read
  - `lastReadAt`: Describes the last point that notifications were checked
    (optional). Default: Now (string, optional)
  - `owner`: Optional repository owner. If provided with repo, only
    notifications for this repository are marked as read. (string, optional)
  - `repo`: Optional repository name. If provided with owner, only notifications
    for this repository are marked as read. (string, optional)

</details>

<details>

<summary>Organizations</summary>

- **search_orgs** - Search organizations
  - `order`: Sort order (string, optional)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `query`: Organization search query. Examples: 'microsoft',
    'location:california', 'created:>=2025-01-01'. Search is automatically
    scoped to type:org. (string, required)
  - `sort`: Sort field by category (string, optional)

</details>

<details>

<summary>Pull Requests</summary>

- **add_comment_to_pending_review** - Add review comment to the requester's
  latest pending pull request review

  - `body`: The text of the review comment (string, required)
  - `line`: The line of the blob in the pull request diff that the comment
    applies to. For multi-line comments, the last line of the range (number,
    optional)
  - `owner`: Repository owner (string, required)
  - `path`: The relative path to the file that necessitates a comment (string,
    required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)
  - `side`: The side of the diff to comment on. LEFT indicates the previous
    state, RIGHT indicates the new state (string, optional)
  - `startLine`: For multi-line comments, the first line of the range that the
    comment applies to (number, optional)
  - `startSide`: For multi-line comments, the starting side of the diff that the
    comment applies to. LEFT indicates the previous state, RIGHT indicates the
    new state (string, optional)
  - `subjectType`: The level at which the comment is targeted (string, required)

- **create_and_submit_pull_request_review** - Create and submit a pull request
  review without comments

  - `body`: Review comment text (string, required)
  - `commitID`: SHA of commit to review (string, optional)
  - `event`: Review action to perform (string, required)
  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **create_pending_pull_request_review** - Create pending pull request review

  - `commitID`: SHA of commit to review (string, optional)
  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **create_pull_request** - Open new pull request

  - `base`: Branch to merge into (string, required)
  - `body`: PR description (string, optional)
  - `draft`: Create as draft PR (boolean, optional)
  - `head`: Branch containing changes (string, required)
  - `maintainer_can_modify`: Allow maintainer edits (boolean, optional)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `title`: PR title (string, required)

- **delete_pending_pull_request_review** - Delete the requester's latest pending
  pull request review

  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **get_pull_request** - Get pull request details

  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **get_pull_request_comments** - Get pull request comments

  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **get_pull_request_diff** - Get pull request diff

  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **get_pull_request_files** - Get pull request files

  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **get_pull_request_reviews** - Get pull request reviews

  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **get_pull_request_status** - Get pull request status checks

  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **list_pull_requests** - List pull requests

  - `base`: Filter by base branch (string, optional)
  - `direction`: Sort direction (string, optional)
  - `head`: Filter by head user/org and branch (string, optional)
  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)
  - `sort`: Sort by (string, optional)
  - `state`: Filter by state (string, optional)

- **merge_pull_request** - Merge pull request

  - `commit_message`: Extra detail for merge commit (string, optional)
  - `commit_title`: Title for merge commit (string, optional)
  - `merge_method`: Merge method (string, optional)
  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **request_copilot_review** - Request Copilot review

  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **search_pull_requests** - Search pull requests

  - `order`: Sort order (string, optional)
  - `owner`: Optional repository owner. If provided with repo, only
    notifications for this repository are listed. (string, optional)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `query`: Search query using GitHub pull request search syntax (string,
    required)
  - `repo`: Optional repository name. If provided with owner, only notifications
    for this repository are listed. (string, optional)
  - `sort`: Sort field by number of matches of categories, defaults to best
    match (string, optional)

- **submit_pending_pull_request_review** - Submit the requester's latest pending
  pull request review

  - `body`: The text of the review comment (string, optional)
  - `event`: The event to perform (string, required)
  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

- **update_pull_request** - Edit pull request

  - `base`: New base branch name (string, optional)
  - `body`: New description (string, optional)
  - `draft`: Mark pull request as draft (true) or ready for review (false)
    (boolean, optional)
  - `maintainer_can_modify`: Allow maintainer edits (boolean, optional)
  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number to update (number, required)
  - `repo`: Repository name (string, required)
  - `reviewers`: GitHub usernames to request reviews from (string[], optional)
  - `state`: New state (string, optional)
  - `title`: New title (string, optional)

- **update_pull_request_branch** - Update pull request branch
  - `expectedHeadSha`: The expected SHA of the pull request's HEAD ref (string,
    optional)
  - `owner`: Repository owner (string, required)
  - `pullNumber`: Pull request number (number, required)
  - `repo`: Repository name (string, required)

</details>

<details>

<summary>Repositories</summary>

- **create_branch** - Create branch

  - `branch`: Name for new branch (string, required)
  - `from_branch`: Source branch (defaults to repo default) (string, optional)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)

- **create_or_update_file** - Create or update file

  - `branch`: Branch to create/update the file in (string, required)
  - `content`: Content of the file (string, required)
  - `message`: Commit message (string, required)
  - `owner`: Repository owner (username or organization) (string, required)
  - `path`: Path where to create/update the file (string, required)
  - `repo`: Repository name (string, required)
  - `sha`: Required if updating an existing file. The blob SHA of the file being
    replaced. (string, optional)

- **create_repository** - Create repository

  - `autoInit`: Initialize with README (boolean, optional)
  - `description`: Repository description (string, optional)
  - `name`: Repository name (string, required)
  - `private`: Whether repo should be private (boolean, optional)

- **delete_file** - Delete file

  - `branch`: Branch to delete the file from (string, required)
  - `message`: Commit message (string, required)
  - `owner`: Repository owner (username or organization) (string, required)
  - `path`: Path to the file to delete (string, required)
  - `repo`: Repository name (string, required)

- **fork_repository** - Fork repository

  - `organization`: Organization to fork to (string, optional)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)

- **get_commit** - Get commit details

  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)
  - `sha`: Commit SHA, branch name, or tag name (string, required)

- **get_file_contents** - Get file or directory contents

  - `owner`: Repository owner (username or organization) (string, required)
  - `path`: Path to file/directory (directories must end with a slash '/')
    (string, optional)
  - `ref`: Accepts optional git refs such as `refs/tags/{tag}`,
    `refs/heads/{branch}` or `refs/pull/{pr_number}/head` (string, optional)
  - `repo`: Repository name (string, required)
  - `sha`: Accepts optional commit SHA. If specified, it will be used instead of
    ref (string, optional)

- **get_tag** - Get tag details

  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)
  - `tag`: Tag name (string, required)

- **list_branches** - List branches

  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)

- **list_commits** - List commits

  - `author`: Author username or email address to filter commits by (string,
    optional)
  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)
  - `sha`: Commit SHA, branch or tag name to list commits of. If not provided,
    uses the default branch of the repository. If a commit SHA is provided, will
    list commits up to that SHA. (string, optional)

- **list_tags** - List tags

  - `owner`: Repository owner (string, required)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `repo`: Repository name (string, required)

- **push_files** - Push files to repository

  - `branch`: Branch to push to (string, required)
  - `files`: Array of file objects to push, each object with path (string) and
    content (string) (object[], required)
  - `message`: Commit message (string, required)
  - `owner`: Repository owner (string, required)
  - `repo`: Repository name (string, required)

- **search_code** - Search code

  - `order`: Sort order for results (string, optional)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `query`: Search query using GitHub's powerful code search syntax. Examples:
    'content:Skill language:Java org:github', 'NOT is:archived language:Python
    OR language:go', 'repo:github/github-mcp-server'. Supports exact matching,
    language filters, path filters, and more. (string, required)
  - `sort`: Sort field ('indexed' only) (string, optional)

- **search_repositories** - Search repositories
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `query`: Repository search query. Examples: 'machine learning in:name
    stars:>1000 language:python', 'topic:react', 'user:facebook'. Supports
    advanced search syntax for precise filtering. (string, required)

</details>

<details>

<summary>Secret Protection</summary>

- **get_secret_scanning_alert** - Get secret scanning alert

  - `alertNumber`: The number of the alert. (number, required)
  - `owner`: The owner of the repository. (string, required)
  - `repo`: The name of the repository. (string, required)

- **list_secret_scanning_alerts** - List secret scanning alerts
  - `owner`: The owner of the repository. (string, required)
  - `repo`: The name of the repository. (string, required)
  - `resolution`: Filter by resolution (string, optional)
  - `secret_type`: A comma-separated list of secret types to return. All default
    secret patterns are returned. To return generic patterns, pass the token
    name(s) in the parameter. (string, optional)
  - `state`: Filter by state (string, optional)

</details>

<details>

<summary>Users</summary>

- **search_users** - Search users
  - `order`: Sort order (string, optional)
  - `page`: Page number for pagination (min 1) (number, optional)
  - `perPage`: Results per page for pagination (min 1, max 100) (number,
    optional)
  - `query`: User search query. Examples: 'john smith', 'location:seattle',
    'followers:>100'. Search is automatically scoped to type:user. (string,
    required)
  - `sort`: Sort users by number of followers or repositories, or when the
    person joined GitHub. (string, optional)

</details>
<!-- END AUTOMATED TOOLS -->

### Additional Tools in Remote Github MCP Server

<details>

<summary>Copilot coding agent</summary>

- **create_pull_request_with_copilot** - Perform task with GitHub Copilot coding
  agent
  - `owner`: Repository owner. You can guess the owner, but confirm it with the
    user before proceeding. (string, required)
  - `repo`: Repository name. You can guess the repository name, but confirm it
    with the user before proceeding. (string, required)
  - `problem_statement`: Detailed description of the task to be performed (e.g.,
    'Implement a feature that does X', 'Fix bug Y', etc.) (string, required)
  - `title`: Title for the pull request that will be created (string, required)
  - `base_ref`: Git reference (e.g., branch) that the agent will start its work
    from. If not specified, defaults to the repository's default branch (string,
    optional)

</details>

#### Specifying Toolsets

To specify toolsets you want available to the LLM, you can pass an allow-list in
two ways:

1. **Using Command Line Argument**:

   ```bash
   github-mcp-server --toolsets repos,issues,pull_requests,actions,code_security
   ```

2. **Using Environment Variable**:
   ```bash
   GITHUB_TOOLSETS="repos,issues,pull_requests,actions,code_security" ./github-mcp-server
   ```

The environment variable `GITHUB_TOOLSETS` takes precedence over the command
line argument if both are provided.

### Using Toolsets With Docker

When using Docker, you can pass the toolsets as environment variables:

```bash
docker run -i --rm \
  -e GITHUB_PERSONAL_ACCESS_TOKEN=<your-token> \
  -e GITHUB_TOOLSETS="repos,issues,pull_requests,actions,code_security,experiments" \
  ghcr.io/github/github-mcp-server
```

### The "all" Toolset

The special toolset `all` can be provided to enable all available toolsets
regardless of any other configuration:

```bash
./github-mcp-server --toolsets all
```

Or using the environment variable:

```bash
GITHUB_TOOLSETS="all" ./github-mcp-server
```

## Dynamic Tool Discovery

**Note**: This feature is currently in beta and may not be available in all
environments. Please test it out and let us know if you encounter any issues.

Instead of starting with all tools enabled, you can turn on dynamic toolset
discovery. Dynamic toolsets allow the MCP host to list and enable toolsets in
response to a user prompt. This should help to avoid situations where the model
gets confused by the sheer number of tools available.

### Using Dynamic Tool Discovery

When using the binary, you can pass the `--dynamic-toolsets` flag.

```bash
./github-mcp-server --dynamic-toolsets
```

When using Docker, you can pass the toolsets as environment variables:

```bash
docker run -i --rm \
  -e GITHUB_PERSONAL_ACCESS_TOKEN=<your-token> \
  -e GITHUB_DYNAMIC_TOOLSETS=1 \
  ghcr.io/github/github-mcp-server
```

## Read-Only Mode

To run the server in read-only mode, you can use the `--read-only` flag. This
will only offer read-only tools, preventing any modifications to repositories,
issues, pull requests, etc.

```bash
./github-mcp-server --read-only
```

When using Docker, you can pass the read-only mode as an environment variable:

```bash
docker run -i --rm \
  -e GITHUB_PERSONAL_ACCESS_TOKEN=<your-token> \
  -e GITHUB_READ_ONLY=1 \
  ghcr.io/github/github-mcp-server
```

## GitHub Enterprise Server and Enterprise Cloud with data residency (ghe.com)

The flag `--gh-host` and the environment variable `GITHUB_HOST` can be used to
set the hostname for GitHub Enterprise Server or GitHub Enterprise Cloud with
data residency.

- For GitHub Enterprise Server, prefix the hostname with the `https://` URI
  scheme, as it otherwise defaults to `http://`, which GitHub Enterprise Server
  does not support.
- For GitHub Enterprise Cloud with data residency, use
  `https://YOURSUBDOMAIN.ghe.com` as the hostname.

```json
"github": {
    "command": "docker",
    "args": [
    "run",
    "-i",
    "--rm",
    "-e",
    "GITHUB_PERSONAL_ACCESS_TOKEN",
    "-e",
    "GITHUB_HOST",
    "ghcr.io/github/github-mcp-server"
    ],
    "env": {
        "GITHUB_PERSONAL_ACCESS_TOKEN": "${input:github_token}",
        "GITHUB_HOST": "https://<your GHES or ghe.com domain name>"
    }
}
```

## i18n / Overriding Descriptions

The descriptions of the tools can be overridden by creating a
`github-mcp-server-config.json` file in the same directory as the binary.

The file should contain a JSON object with the tool names as keys and the new
descriptions as values. For example:

```json
{
  "TOOL_ADD_ISSUE_COMMENT_DESCRIPTION": "an alternative description",
  "TOOL_CREATE_BRANCH_DESCRIPTION": "Create a new branch in a GitHub repository"
}
```

You can create an export of the current translations by running the binary with
the `--export-translations` flag.

This flag will preserve any translations/overrides you have made, while adding
any new translations that have been added to the binary since the last time you
exported.

```sh
./github-mcp-server --export-translations
cat github-mcp-server-config.json
```

You can also use ENV vars to override the descriptions. The environment variable
names are the same as the keys in the JSON file, prefixed with `GITHUB_MCP_` and
all uppercase.

For example, to override the `TOOL_ADD_ISSUE_COMMENT_DESCRIPTION` tool, you can
set the following environment variable:

```sh
export GITHUB_MCP_TOOL_ADD_ISSUE_COMMENT_DESCRIPTION="an alternative description"
```

## Library Usage

The exported Go API of this module should currently be considered unstable, and
subject to breaking changes. In the future, we may offer stability; please file
an issue if there is a use case where this would be valuable.

## License

This project is licensed under the terms of the MIT open source license. Please
refer to [MIT](./LICENSE) for the full terms.
