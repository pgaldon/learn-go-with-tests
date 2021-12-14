FROM alpine:latest

# Versions
ARG TERRAFORM_VERSION=0.14.2
ARG TFLINT_VERSION=0.17.0

# Configure apt and install packages
# Verify git, process tools, lsb-release (common in install instructions for CLIs), unzip, curl, wget installed
RUN apk add bash git unzip curl file wget vim make go \
    # Install Terraform
    && mkdir -p /tmp/docker-downloads \
    && curl -sSL -o /tmp/docker-downloads/terraform.zip https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip \
    && ls -la /tmp/docker-downloads/ \
    && unzip /tmp/docker-downloads/terraform.zip \
    && mv terraform /usr/local/bin \
    # Install tflint
    && curl -sSL -o /tmp/docker-downloads/tflint.zip https://github.com/terraform-linters/tflint/releases/download/v${TFLINT_VERSION}/tflint_linux_amd64.zip \
    && unzip /tmp/docker-downloads/tflint.zip \
    && mv tflint /usr/local/bin \
    # Clean up downloaded files
    && cd ~ \ 
    && rm -rf /tmp/docker-downloads 

# Bash customization
RUN \
    ## Git command prompt
    git clone https://github.com/magicmonty/bash-git-prompt.git ~/.bash-git-prompt --depth=1 \
    && echo "if [ -f \"$HOME/.bash-git-prompt/gitprompt.sh\" ]; then GIT_PROMPT_ONLY_IN_REPO=1 && source $HOME/.bash-git-prompt/gitprompt.sh; fi" >> "/root/.bashrc" \
    ## Add aliasreeee
    && echo "alias tf=terraform" >> "/root/.bashrc"

COPY .terraformrc /root/.terraformrc

EXPOSE 27228    