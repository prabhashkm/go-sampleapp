# Install Golang

```bash
brew install go
```

# Configure Golang

First create the Golang directory structure.

```bash
mkdir -p ~/go/{bin,src,pkg}
```

Add GOPATH environment variable.

```bash
cat << EOF >>~/.bash_profile
export GOPATH=$HOME/go
export PATH=\$PATH:\${GOPATH}/bin
EOF
```

# Install gb

install gb to manage dependencies http://getgb.io

```bash
source ~/.bash_profile
go get github.com/constabulary/gb/...
```

# Add gofmt to Vim

It's recommended to run gofmt as a post job when saving files, here is the required syntax for Vim. This make's
sure your code follows the standardized formating.

```
cat << EOF >>~/.vimrc
au BufWritePost *.go !gofmt -w %
EOF
```

# Create go-sampleapp

```bash
mkdir -p ~/code/go-sampleapp
cd ~/code/go-sampleapp
git init
```

# Create Makefile

Create Makefile.

```
cat << EOF | unexpand >Makefile
all: build

clean:
        rm -rf pkg bin

test: clean
        gb test

build: test
        gb build all

update:
        gb vendor update --all
EOF
```

# Add go-sampleapp

Copy the sample application from this repository.

```bash
git archive --remote=https://github.com/mickep76/go-sampleapp.git master src | tar xvf -
```

# Vendor third-party packages

```bash
gb vendor fetch github.com/Sirupsen/logrus
gb vendor fetch github.com/jessevdk/go-flags
gb vendor fetch gopkg.in/yaml.v2
gb vendor fetch github.com/BurntSushi/toml
```

# Build application

```bash
make
```

# Test application

```
git archive --remote=https://github.com/mickep76/go-sampleapp.git master examples | tar xvf -
bin/go-sampleapp -h
bin/go-sampleapp -v -f examples/sample.yaml
bin/go-sampleapp -v -f examples/sample.json
bin/go-sampleapp -v -f examples/sample.toml
```

# Activate git hooks

First install golint.

```bash
go get -u github.com/golang/lint/golint
```

Now download the githooks and activate them.

```bash
git archive --remote=https://github.com/mickep76/go-sampleapp.git master .githooks | tar xvf -
.githooks/activate
```

If you want to skip the hooks just use "--no-verify".
