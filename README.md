# SubsGenie README.md

Welcome to SubsGenie, your automated companion for fetching and updating Genshin Impact Void Terminal (Clash M) subscription files with ease. SubsGenie is designed to minimize the waiting time for updating your subscriptions by automatically requesting the latest files from subconverter and uploading them to a specified gist.

## Features

- **Automatic Updates:** SubsGenie regularly checks for and retrieves the latest subscription files, ensuring your Void Terminal is always up to date.
- **Customizable Subscriptions:** Support for a primary subscription URL along with additional ones, including the option for community-provided URLs.
- **Gist Integration:** Seamlessly upload and manage your subscription files on GitHub Gist, keeping your configurations private and accessible.
- **Flexible Configuration:** Tailor the service to your needs with various environment variables, allowing for detailed control over the request process.
- **Logging:** Track the operation of SubsGenie with configurable log levels, ensuring transparency and ease of troubleshooting.

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

## Contributing

Contributions to SubsGenie are welcome! Whether it's feature requests, bug reports, or code contributions, please feel free to reach out or submit a pull request.

## License

SubsGenie is open-sourced under the MIT license. See the LICENSE file for more details.

---

SubsGenie: Automating your Genshin Impact Void Terminal subscription management, making your adventures in Teyvat as seamless as possible.
