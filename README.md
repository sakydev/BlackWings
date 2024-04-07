# BlackWings
![Go](https://img.shields.io/badge/Go-00ADD8?logo=go&logoColor=white) ![Chi](https://img.shields.io/badge/Chi-02A9E0?logo=go&logoColor=white) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-336791?logo=postgresql&logoColor=white) ![Cobra](https://img.shields.io/badge/Cobra-02A9E0?logo=go&logoColor=white) ![Docker](https://img.shields.io/badge/Docker-2496ED?logo=docker&logoColor=white)

A command line tool to search across your entire digital footprint. This will include essential apps like Gmail, Drive, Dropbox, Jira, Slack, Confluence and more.

### Why
I often encounter situations where I remember something but can't remember where it exists. This can include documentation, technical discussions, and more. I felt the need for a uniform tool that has wings into my essential apps and can retrieve what I'm looking for. 

BlackWings is a Go-powered CLI built with Cobra and integrates with popular APIs. It authenticates once and uses tokens for further refreshses. You can add an unlimited number of accounts

### Installation
This will install Goose for database migrations and all other packages required and setup database.
```
git clone git@github.com:sakydev/black-wings.git
cd black-wings && make install
```

### Commands
#### 1. Search
Search command allows you to find items matching criteria

```bash
# list options
blackwings search -h

Usage:
  blackwings search [flags]

Flags:
      --after string         Results after date
  -a, --apps strings         Apps to search
      --before string        Results before date
  -e, --exclude string       Results must exclude
  -t, --file-types strings   File types to search
  -h, --help                 help for search
  -i, --include string       Results must include
  -l, --limit int            Results limit (default 20)
      --offset int           Results offset
  -o, --order string         Order results by (default "desc")
  -q, --query string         Search query
  -s, --sort string          Sort results by (default "relevance")
```

Examples:
```bash

# search across all connected apps
blackwings search --query "search query"

# limit search to specific apps
 blackwings search --query "hello" --apps "app1,app2,app3"

# include or exclude items that contain certain word
blackwings search --query "hello" --include "include1,include2" --exclude "excludeA,excludeB"

# search limit and pagination
blackwings search --query "search query" --limit 10 --offset 10
```

#### 2. Accounts
Accounts commands allow adding and removing accounts
```bash
# list options
blackwings accounts -h

# connect a new account in interactive mode
blackwings accounts connect

# list all accounts
blackwings accounts list

# list accounts for a specific provider only e.g gmail
blackwings accounts list --provider "gmail"

# remove an account in interactive mode
blackwings accounts disconnect 
```
