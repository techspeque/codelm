# Use a stable Ubuntu LTS base image
FROM ubuntu:24.04

# Avoid interactive prompts during build
ENV DEBIAN_FRONTEND=noninteractive

# Install essential packages, including Zsh, Python, Node.js, and npm
RUN apt-get update && apt-get install -y \
    zsh \
    git \
    curl \
    wget \
    python3 \
    python3-pip \
    nodejs \
    npm \
    && rm -rf /var/lib/apt/lists/*

# Install the correct AI CLIs using npm
RUN npm install -g @google/gemini-cli
RUN npm install -g @anthropic-ai/claude-code
RUN npm install -g @openai/codex

# Set up a non-root user 'devuser' for better security
RUN useradd --create-home --shell /bin/zsh devuser
USER devuser
WORKDIR /home/devuser

# Install Oh My Zsh and key plugins
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended
RUN git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
RUN git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

# Configure .zshrc with plugins and the safe import logic
RUN echo 'export ZSH="$HOME/.oh-my-zsh"' > /home/devuser/.zshrc && \
    echo 'ZSH_THEME="robbyrussell"' >> /home/devuser/.zshrc && \
    echo 'plugins=(git zsh-autosuggestions zsh-syntax-highlighting)' >> /home/devuser/.zshrc && \
    echo 'source $ZSH/oh-my-zsh.sh' >> /home/devuser/.zshrc && \
    echo '\n# --- Safe import for user settings ---' >> /home/devuser/.zshrc && \
    echo 'if [ -f /home/devuser/.zshrc_user ]; then' >> /home/devuser/.zshrc && \
    echo '    source /home/devuser/.zshrc_user' >> /home/devuser/.zshrc && \
    echo 'fi' >> /home/devuser/.zshrc

# Create the workspace directory and set it as the default
WORKDIR /workspace

# Set the default command to Zsh
CMD [ "/bin/zsh" ]