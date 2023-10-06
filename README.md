# Jira-Sync

Jira-Sync is a lightweight command-line application that allows you to efficiently add spent time to your Jira tasks without the usual Jira interface freezes. It fetches your tasks where time is not tracked and provides a seamless way to update them.

## Features

- Retrieve and list tasks with untracked time.
- Add spent time to tasks without Jira interface slowdowns.
- Cross-platform support: Available for various OS/Architectures.
- Easy setup using your Jira domain, project name, email, and generated Jira token.

## Installation

1. Download the Jira-Sync application for your OS/Architecture from the [Releases](https://github.com/denysovkos/jira-sync/releases) page.

2. Extract the downloaded archive to a location of your choice.

## Setup

Before using Jira-Sync, you'll need to set it up with your Jira account:

1. Open a terminal window and navigate to the directory where you extracted the Jira-Sync application.

2. Run the application:

```bash
./jira-sync
```
You will be prompted to enter the following information:

* Jira Domain: The domain of your Jira instance (e.g., https://yourcompany.atlassian.net).
* Project Name: The name of the Jira project where you want to add spent time.
* Email: Your Jira account email address.
* Jira Token: To generate a Jira token, follow the instructions here.
  
After entering the required information, the application will save your configuration.

Close the application using CTRL-C and run it again to start using Jira-Sync.

## Usage
Once you have set up Jira-Sync, you can use it to manage your Jira tasks:

1. Run the application:
```bash
./jira-sync
```
2. The application will fetch tasks with untracked time.

3. Select a task from the table and enter the spent time for that task.

4. Confirm the update (click Submit), and the application will add the spent time to the selected task.

5. Repeat the process for other tasks as needed.


## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments
Thanks to the Go community for providing a powerful development environment.
Contributing
Contributions are welcome! Feel free to open an issue or submit a pull request.

## Support
If you encounter any issues or have questions, please open an issue on the GitHub repository.
