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

## 贡献

欢迎对 SubsGenie 进行贡献！无论是功能请求、错误报告还是代码贡献，请随时联系我们或提交拉取请求。

## 许可证

SubsGenie 采用 MIT 许可证开源。有关更多详情，请查看 LICENSE 文件。

---

SubsGenie：自动化管理你的原神虚空终端订阅，让你在提瓦特大陆的冒险之旅更加无缝顺畅。
