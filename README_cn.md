# SubsGenie

SubsGenie 是一个为原神的虚空终端（Clash M）订阅文件自动更新而设计的自动化工具。SubsGenie 旨在通过自动请求 subconverter 获取最新的订阅文件并上传至指定的 GitHub Gist 中，从而减少更新原神虚空终端订阅时的等待时间。

## 特性

- **自动更新**：SubsGenie 定期检查并获取最新的订阅文件，确保你的虚空终端始终保持最新状态。
- **自定义订阅**：支持设置主要订阅 URL 和额外订阅 URL，包括大善人提供的超大杯节点。
- **Gist 集成**：无缝地在 GitHub Gist 上上传和管理你的订阅文件，保持配置的私密性和可访问性。
- **灵活配置**：通过各种环境变量进行详细控制，根据需要调整请求过程。
- **日志记录**：通过可配置的日志级别跟踪 SubsGenie 的操作，确保透明性和易于故障排除。

## 配置

开始使用 SubsGenie 前，你需要配置以下环境变量：

- **CORE_SUBSCRIPTION_URL**：你的主要订阅地址。
- **EXTRA_PARAMS**：设置为 `"&expand=false&classic=true&new_name=true&udp=true&append_type=true"`，以自定义请求参数。
- **GIST_ID**：在 GitHub 上创建一个私有 Gist，并使用其 ID。
- **LOG_LEVEL**：设置为 `INFO` 以进行标准日志记录（根据需要调整，以获取更多或更少的详细信息，**千万别在 Action 中使用 DEBUG 级别**）。
- **GITHUB_TOKEN**：生成具有 Gist 权限的令牌，用于访问 GitHub Gist。

### 可选环境变量：

- **SUBCONVERT_SERVICE_URL**：后端服务地址，例如 `https://api.dler.io/sub`。
- **CONFIG_URL**：配置文件地址。留空则使用默认配置。
- **CF_SUBSCRIPTION_URL**：大善人的订阅地址。
- **OTHER_SUBSCRIPTION_URLS**：其他订阅地址，用英文逗号分隔。

## 开始使用

1. **环境设置**：确保根据你的偏好和需求设置所有必需的环境变量。
2. **Gist 准备**：在 GitHub 上创建一个私有 Gist，并记下其 ID 用于 `GIST_ID` 变量。
3. **令牌生成**：生成一个具有 Gist 权限的 GitHub 令牌，并将其设置为你的 `GITHUB_TOKEN`。
4. **配置**：可选地，使用额外的环境变量自定义设置，以获得更加定制化的功能。

## 使用方式

一旦完成配置，SubsGenie 将自动管理获取最新订阅文件并上传到你指定的 Gist 的过程。确保你的系统能够访问互联网并且具有执行脚本和访问 GitHub Gist 的必要权限。

## GitHub Proxy 反向代理功能

为了帮助一些地区的用户解决无法直接访问 GitHub Gist 的问题，SubsGenie 现新增了 GitHub Proxy 反向代理功能。通过在 Cloudflare Workers 上部署一个简单的代理服务，用户可以无障碍地访问由 SubsGenie 管理的 Gist 内容。

### 配置步骤

1. **进入代理服务目录**：在项目根目录中找到并进入代理服务所在的文件夹。

2. **安装依赖**：`pnpm install`

3. **登录 Cloudflare 账号**：`pnpm wrangler login`

按照提示登录你的 Cloudflare 账号。

4. **创建配置文件**：在文件夹内创建名为 `wrangler.toml` 的配置文件，并填写以下内容：

> 如果没有，就删掉`routes` 这句话，试试用大善人送的域名。

```
name = "subs-genie-worker"
main = "src/index.ts"
compatibility_date = "2024-03-04" # 使用你的部署日期

[env.production]
vars = { ENVIRONMENT = "production" }
routes = [{ pattern = "填写你的域名", custom_domain = true }, "你的域名/*"]
```

将 `填写你的域名` 替换为你实际的域名信息。

5. **部署代理服务**：`pnpm run deploy`

执行此命令将代理服务部署到 Cloudflare Workers。

### 使用方法

部署成功后，通过以下格式的网址访问你的 Gist 内容：

```
https://你的网址/api/v1/github/gist_id?name=coreAndCF.txt&user=你GitHub用户名
```

将 `你的网址`、`gist_id`、`你GitHub用户名` 替换为实际的信息。

通过这个代理服务，无论你所在的地区是否直接支持访问 GitHub Gist，都能顺畅地获取 SubsGenie 定时上传的内容。

---

我们希望这个新功能能够帮助所有用户更好地使用 SubsGenie，享受无缝的原神虚空终端订阅管理体验。


## 贡献

欢迎对 SubsGenie 进行贡献！无论是功能请求、错误报告还是代码贡献，请随时联系我们或提交拉取请求。

## 许可证

SubsGenie 采用 MIT 许可证开源。有关更多详情，请查看 LICENSE 文件。

---

SubsGenie：自动化管理你的原神虚空终端订阅，让你在提瓦特大陆的冒险之旅更加无缝顺畅。
