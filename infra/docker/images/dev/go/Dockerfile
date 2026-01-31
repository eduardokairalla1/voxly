# Base Go image
FROM golang:1.25-alpine

# Install system dependencies and build tools
RUN apk update && apk add --no-cache \
    gcc \
    g++ \
    git \
    curl \
    vim \
    openssh \
    bash \
    make \
    musl-dev \
    && rm -rf /var/cache/apk/*

# Configure SSH for passwordless root access (dev only!)
RUN ssh-keygen -A && \
    echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config && \
    echo 'PermitEmptyPasswords yes' >> /etc/ssh/sshd_config && \
    echo 'StrictModes no' >> /etc/ssh/sshd_config && \
    passwd -d root

# Install Go development tools
RUN go install github.com/daixiang0/gci@v0.13.7 && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 && \
    go install github.com/air-verse/air@v1.64.4

# Add Go binaries to PATH
ENV PATH="/usr/local/go/bin:/go/bin:${PATH}"

# Configure PATH for SSH sessions
RUN echo 'export PATH="/usr/local/go/bin:/go/bin:$PATH"' >> /root/.bashrc && \
    echo 'export PATH="/usr/local/go/bin:/go/bin:$PATH"' >> /root/.profile

# Set working directory to root home
WORKDIR /root

# Expose SSH port
EXPOSE 22

# Start SSH daemon in foreground
CMD ["/usr/sbin/sshd", "-D"]
