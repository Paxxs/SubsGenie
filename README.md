# SubsGenie

SubsGenie is an automation tool for updating and merging Clash (Mihomo) proxy subscriptions using GitHub Actions. It works by pre-processing and combining multiple subscription sources, then saving the results to GitHub Gist, significantly reducing the waiting time when clients (like Shadowrocket, Clash, etc.) fetch subscriptions.

[中文文档](README_cn.md)

## Features

- **Automated Merging and Updating**: Runs on a schedule via GitHub Actions to automatically retrieve and merge multiple subscription sources without manual intervention.
- **Reduced Client Waiting Time**: Pre-processes subscription content and stores it in Gist, making it almost instantaneous for clients to fetch, preventing timeouts in apps like Shadowrocket on iOS which fail if waiting time is too long.
- **Multiple Subscription Support**: Set up primary subscription URLs and additional subscription URLs to easily combine nodes from multiple sources.
- **Gist Storage and Distribution**: Stores processed subscription files on GitHub Gist, ensuring both privacy and reliable access.
- **Reverse Proxy Support**: Provides Cloudflare Workers reverse proxy functionality to solve access issues in regions where GitHub Gist is restricted.
- **Flexible Configuration**: Control the merging and processing through environment variables, adjusting as needed.

## Configuration

To get started with SubsGenie, you need to configure the following environment variables:

- **CORE_SUBSCRIPTION_URL**: Your primary subscription URL.
- **EXTRA_PARAMS**: Set to `"&expand=false&classic=true&new_name=true&udp=true&append_type=true"` to customize the request parameters.
- **GIST_ID**: Create a private gist and use its ID here.
- **LOG_LEVEL**: Set to `INFO` for standard logging (adjust as needed for more or less verbosity).
- **GITHUB_TOKEN**: Generate a token with gist permissions for accessing GitHub Gist.

### Optional Environment Variables:

- **SUBCONVERT_SERVICE_URL**: The backend service URL, e.g., `https://api.dler.io/sub`.
- **CONFIG_URL**: URL for a configuration file. Leave blank to use the provided default.
- **CF_SUBSCRIPTION_URL**: A benevolent provider's subscription URL.
- **OTHER_SUBSCRIPTION_URLS**: Additional subscription URLs, separated by commas.

## Getting Started

1. **Environment Setup**: Ensure all required environment variables are set according to your preferences and needs.
2. **Gist Preparation**: Create a private gist on GitHub and note its ID for the `GIST_ID` variable.
3. **Token Generation**: Generate a GitHub token with gist permissions and set it as your `GITHUB_TOKEN`.
4. **Configuration**: Optionally, customize your setup with additional environment variables for more tailored functionality.

## Usage

Once configured, SubsGenie will automatically manage the process of fetching the latest subscription files and uploading them to your specified gist. Ensure your system has access to the internet and the necessary permissions to execute scripts and access GitHub Gist.

## GitHub Proxy Reverse Proxy Feature

To assist users in regions where direct access to GitHub Gist is restricted, SubsGenie includes a GitHub Proxy reverse proxy feature. By deploying a simple proxy service on Cloudflare Workers, users can seamlessly access the Gist content managed by SubsGenie without any barriers.

### Configuration Steps

1. **Navigate to the Proxy Service Directory**: Locate and enter the folder for the proxy service within the root directory of the project.

2. **Install Dependencies**: `pnpm install`

3. **Login to Your Cloudflare Account**: `pnpm wrangler login`

Follow the prompts to log in to your Cloudflare account.

4. **Create a Configuration File**: In the folder, create a `wrangler.toml` configuration file with the following content:

```toml
name = "subs-genie-worker"
main = "src/index.ts"
compatibility_date = "2024-03-04" # Use your deployment date

[env.production]
vars = { ENVIRONMENT = "production" }
routes = [{ pattern = "your domain", custom_domain = true }, "your domain/*"]
```

Replace `your domain` with your actual domain information.

5. **Deploy the Proxy Service**: `pnpm run deploy`

This command deploys your proxy service to Cloudflare Workers.

### How to Use

After a successful deployment, access your Gist content via the following URL format:

```
https://your-website/api/v1/github/gist_id?name=coreAndCF.txt&user=yourGitHubUsername
```

Replace `your-website`, `gist_id`, and `yourGitHubUsername` with your actual information.

With this proxy service, you can smoothly retrieve the content uploaded by SubsGenie on a scheduled basis, regardless of whether your region supports direct access to GitHub Gist or not.

---

We hope this feature helps all users better utilize SubsGenie for their network subscription management needs.

## Contributing

Contributions to SubsGenie are welcome! Whether it's feature requests, bug reports, or code contributions, please feel free to reach out or submit a pull request.

## License

SubsGenie is open-sourced under the MIT license. See the LICENSE file for more details.

---

SubsGenie: Automating your Clash (Mihomo) proxy subscription management for a smoother, more reliable network experience.
