# CodeLM

*A CLI for isolated, containerized environments for AI-assisted coding.*

`codelm` is a command-line utility that manages Docker containers with pre-installed AI tools. It provides a consistent, secure, and isolated development workspace, ensuring your projects and tools are neatly separated from your host machine.

## ✨ Features

* **Isolated Workspace**: Leverages Docker to create a self-contained Ubuntu-based environment, preventing conflicts with your local setup.
* **AI Tools Included**: The environment comes pre-packaged with command-line tools for popular AI models, including Google Gemini, Anthropic Claude, and OpenAI Codex.
* **Simple Configuration**: A quick `configure` command interactively sets up your API keys and preferences in a clean YAML file.
* **Project Persistence**: Mounts your local project directory directly into the container and remembers your last-used project for quick starts.
* **Seamless Shell Integration**: Attach to the container's pre-configured Zsh shell, complete with Oh My Zsh, auto-suggestions, and syntax highlighting. You can even import your host's `.zshrc` and `.gitconfig` for a familiar feel.
* **Correct File Permissions**: The container runs using your local user and group ID, so any files you create inside the container have the correct ownership on your host machine.
* **Environment Health Checks**: Includes a `codelm doctor` command to diagnose common issues with your Docker installation and configuration.

## Prerequisites

Before you begin, ensure you have the following installed:

* **Go**: Version 1.24.4 or newer.
* **Docker**: The Docker daemon must be running. `codelm` will not work without it.
* **Git**: For cloning the repository.

## Installation (from Source)

To build `codelm` from the source code, follow these steps:

1. **Clone the repository:**

    ```sh
    git clone [https://github.com/your-username/codelm.git](https://github.com/techspeque/codelm.git)
    cd codelm
    ```

2. **Build the binary:**

    ```sh
    go build -o codelm .
    ```

3. **Move the binary to your PATH:**
    To make the command accessible from anywhere, move it to a directory in your system's PATH.

    ```sh
    # For macOS and Linux
    sudo mv codelm /usr/local/bin/
    ```

Now you can run the `codelm` command from any terminal.

## 🚀 Getting Started

Follow these steps to get your first `codelm` session running.

### 1\. Check Your Environment

Run the doctor command to ensure Docker is running and your system is ready.

```sh
codelm doctor
```

This will check for the Docker installation and daemon status.

### 2\. Configure API Keys

Set up your API keys and preferences with the interactive configuration command.

```sh
codelm configure
```

You will be prompted to enter your OpenAI and Anthropic API keys and choose whether to import your personal shell settings.

### 3\. Start a Session

Start a new session by pointing `codelm` to your project directory.

```sh
codelm start /path/to/your/project
```

This command will:

* Build the `codelm:latest` Docker image if it doesn't already exist.
* Start a container named `codelm_session`.
* Mount your project directory to `/workspace` inside the container.
* Save the project path for future use.

If you've already run `start` once, you can omit the path to resume the last session:

```sh
codelm start
```

### 4\. Enter the Shell

Attach a Zsh shell to the running container.

```sh
codelm shell
```

You are now inside the container at the `/workspace` directory, with all AI tools available in your PATH.

### 5\. Stop the Session

When you are finished, stop and remove the container.

```sh
codelm stop
```

This command cleans up the container resources but keeps your project files and the Docker image intact.

## 📖 Command Reference

| Command                  | Description                                                              |
| ------------------------ | ------------------------------------------------------------------------ |
| `codelm configure`       | Interactively create the `codelm` configuration file.      |
| `codelm configure show`  | Show the current configuration with redacted API keys.        |
| `codelm start [path]`    | Start a new session, using the last project if no path is given. |
| `codelm shell`           | Attach a shell to the running `codelm` container.           |
| `codelm stop`            | Stop and remove the `codelm` container session.              |
| `codelm status`          | Show the status and details of the container session.      |
| `codelm rebuild`         | Forcefully remove and rebuild the `codelm` Docker image.         |
| `codelm doctor`          | Check your local environment for potential problems.       |

## 📜 License

This project is licensed under the Apache License, Version 2.0.
