# Devcontainer base with Go support
FROM mcr.microsoft.com/devcontainers/go:1.21

# Install AWS CLI and Pulumi CLI
USER root

RUN rm -f /etc/apt/sources.list.d/yarn.list /etc/apt/sources.list.d/yarn.list.save \
    && apt-get update \
    && apt-get install -y --no-install-recommends curl unzip python3 awscli \
    && curl -fsSL https://get.pulumi.com | sh \
    && ln -s /root/.pulumi/bin/pulumi /usr/local/bin/pulumi \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

ENV PATH="/root/.pulumi/bin:${PATH}"

USER vscode
