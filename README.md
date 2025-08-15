# codelm

*A CLI for isolated, containerized environments for AI-assisted coding.*

`codelm` is a command-line utility that manages Docker containers with pre-installed AI tools. It provides a consistent, secure, and isolated development workspace, ensuring your projects and tools are neatly separated from your host machine.

-----

## ✨ Features

* **Isolated Workspace**: Leverages Docker to create a self-contained Ubuntu-based environment, preventing conflicts with your local setup.
* **AI Tools Included**: The environment comes pre-packaged with command-line tools for popular AI models, including Google Gemini, Anthropic Claude, and OpenAI Codex.
* **Simple Configuration**: A quick `configure` command interactively sets up your API keys and preferences in a clean YAML file.
* **Project Persistence**: Mounts your local project directory directly into the container and remembers your last-used project for quick starts.
* **Seamless Shell Integration**: Attach to the container's pre-configured Zsh shell, complete with Oh My Zsh, auto-suggestions, and syntax highlighting. You can even import your host's `.zshrc` and `.gitconfig` for a familiar feel.
* **Correct File Permissions**: The container runs using your local user and group ID, so any files you create inside the container have the correct ownership on your host machine.
* **Environment Health Checks**: Includes a `codelm doctor` command to diagnose common issues with your Docker installation and configuration.

-----

## 🚀 Installation (macOS)

### Prerequisites

* **Docker**: You must have Docker Desktop for Mac installed and the **Docker daemon running**.
* **macOS**: This application is currently distributed for macOS.

### Instructions

1. **Download the latest release**: Go to the [**Releases page**](https://github.com/techspeque/codelm/releases). At the moment we only support Apple Silicon (M1/2/3/4).

2. **Make the binary executable**: Open your Terminal, navigate to your `Downloads` folder, and run the `chmod` command.

    ```sh
    # If you downloaded the Apple Silicon version (arm64)
    chmod +x codelm
    ```

3. **Move the binary to your PATH**: To run the `codelm` command from anywhere, move the executable to `/usr/local/bin` and rename it.

    ```sh
    # Make sure to use the filename you downloaded
    sudo mv codelm /usr/local/bin/codelm
    ```

    *You'll be prompted for your password as this command requires administrator privileges.*

4. **Verify the installation**: Close and reopen your Terminal, then run the doctor command:

    ```sh
    codelm doctor
    ```

    If it runs successfully, you're all set\!

-----

## 🛠️ Getting Started

Follow these steps to get your first `codelm` session running.

### 1\. Configure API Keys

Set up your API keys and preferences with the interactive configuration command.

```sh
codelm configure
```

You will be prompted to enter your API keys and choose whether to import your personal shell settings.

### 2\. Start a Session

Start a new session by pointing `codelm` to your project directory.

```sh
codelm start /path/to/your/project
```

This command builds the Docker image if needed, starts the container, and mounts your project directory to `/workspace`. If you've run `start` before, you can omit the path to resume the last session.

### 3\. Enter the Shell

Attach a Zsh shell to the running container.

```sh
codelm shell
```

You are now inside the container at the `/workspace` directory.

### 4\. Stop the Session

When you are finished, stop and remove the container.

```sh
codelm stop
```

This command cleans up the container resources but keeps your project files and the Docker image intact.

-----

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

## Building From source

### Requirements

* **Go**: Version 1.24.4 or newer.
* **Docker**: The Docker daemon must be running. `codelm` will not work without it.
* **Git**: For cloning the repository.

### Steps

1. **Clone the repository:**

    ```sh
    git clone https://github.com/techspeque/codelm.git
    cd codelm
    ```

2. **Build the binary:**

    ```sh
    go build -o codelm .
    ```

3. **Move the binary to your PATH:**

    ```sh
    sudo mv codelm /usr/local/bin/
    ```

-----

## 📜 License

This project is licensed under the Apache License, Version 2.0.

-----
