<h1 align="center" style="border-bottom: none">
    <b>
        <a href="https://docker.nsddd.top">github-label-syncer</a><br>
    </b>
</h1>
<h3 align="center" style="border-bottom: none">
      ⭐️  sync labels between repos and org.  ⭐️ <br>
<h3>


<p align=center>
<a href="https://goreportcard.com/report/github.com/kubecub/github-label-syncer"><img src="https://goreportcard.com/badge/github.com/kubecub/github-label-syncer" alt="A+"></a>
<a href="https://github.com/kubecub/github-label-syncer/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A%22good+first+issue%22"><img src="https://img.shields.io/github/issues/kubecub/github-label-syncer/good%20first%20issue?logo=%22github%22" alt="good first"></a>
<a href="https://github.com/kubecub/github-label-syncer"><img src="https://img.shields.io/github/stars/kubecub/github-label-syncer.svg?style=flat&logo=github&colorB=deeppink&label=stars"></a>
<a href="https://join.slack.com/t/kubecub/shared_invite/zt-1se0k2bae-lkYzz0_T~BYh3rjkvlcUqQ"><img src="https://img.shields.io/badge/Slack-100%2B-blueviolet?logo=slack&amp;logoColor=white"></a>
<a href="https://github.com/kubecub/github-label-syncer/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-green"></a>
<a href="https://golang.org/"><img src="https://img.shields.io/badge/Language-Go-blue.svg"></a>
</p>

</p>

<p align="center">
    <a href="./README.md"><b>English</b></a> •
    <a href="./README_zh-CN.md"><b>中文</b></a>
</p>

</p>

----

## 🧩 Awesome features

This document discusses the development of a Go language actions that synchronizes labels on GitHub repositories. This actions is not limited to synchronizing labels on your own repository, but can also retrieve all of the API objects of a target repository, which can then be saved to a YAML file or synchronized to your own GitHub project.

Labels are an important way of organizing and categorizing issues and pull requests on GitHub. They allow you to quickly and easily identify the status, priority, and type of each issue or pull request.

To retrieve the API objects of a GitHub project's labels, you can use the following URLs:

- To retrieve the first page of objects: `https://api.github.com/repos/<owner>/<repo>/labels?page=1&sort=name-asc`
- To retrieve all objects: `https://api.github.com/repos/<owner>/<repo>/labels?per_page=100&page=1&sort=name-asc`

Simply replace `<owner>` and `<repo>` with the owner and repository name, respectively. These URLs can be used to retrieve all of the labels associated with a project, which can then be synchronized or saved as needed.

github-label-syncer is a tool to sync labels between repos and org. The main features are:

- [x]  Feel free to pick a remote repository and fetch all its labels to a local `yaml`, `table` AND `json`
- [x]  Feel free to pick a remote repository and sync all of its labels to your desired project repository
- [x]  Sync labels across repositories according to configuration 
- [x]  Create, update and delete labels 
- [x]  Dry run to preview which labels would be created, updated or deleted 
- [x]  Handle label collisions by renaming or skipping existing labels 
- [x]  Support labeling pull requests and issues


**Labels denger:**
The labels are designed semantic and standard. We provide labels for priority, type, status etc. You can refer to the [label doc](https://github.com/kubecub/github-label-syncer/labels) to get more info.

**github actions


## 🛫 Quick start

> **Note**: You can get started quickly with github-label-syncer.


#### Work with Makefile

```bash
❯ make help    # show help
❯ make build   # build binary exporter and syncer
❯ echo 'export PATH=$PATH:/path/to/_output/platforms/linux/amd64/' | tee -a ~/.zshrc;source ~/.zshrc
```


#### Use Github-Label-Syncer

You can set your own `GITHUB_TOKEN` via env or via the `export GITHUB_TOKEN` environment variable, or use one of the default tokens we provide, which is `TOEKN` for our free automated [🤖 robot](https://github.com/kubbot)

```bash
❯ ./_output/bin/platforms/linux/amd64/exporter --help
usage: exporter [<flags>] <owner> <repo>


Flags:
      --[no-]help    Show context-sensitive help (also try --help-long and --help-man).
  -y, --[no-]yaml    Use the YAML format.
  -j, --[no-]json    Use the JSON format.
  -t, --[no-]table   Use the table format.
  -x, --[no-]xml     Use the XML format.
      --[no-]toml    Use the TOML format.
      --[no-]ini     Use the INI format.
      --[no-]csv     Use the CSV format.
  -f, --file=FILE    Export labels to file.
      --token=TOKEN  GitHub token. ($GITHUB_TOKEN)

Args:
  <owner>  Owner of the repository.
  <repo>   Repository whose wanted labels.

❯ export GITHUB_TOKEN="ghp_****************************************0z"
❯ exporter kubecub log --json | jq
[
  {
    "name": "major version",
    "description": "Automatically create a new major version tag after PR is merged",
    "color": "1E8DE7"
  },
  {
    "name": "merge when passing",
    "description": "Merge the PR automatically once all status checks have passed",
    "color": "FF851B"
  },
  {
    "name": "minor version",
    "description": "Automatically create a new minor version tag after PR is merged",
    "color": "6EBAF7"
  },
  {
    "name": "patch version",
    "description": "Automatically create a new patch version tag after PR is merged",
    "color": "99cef9"
  }
]

❯ ./_output/bin/platforms/linux/amd64/exporter kubernetes kubernetes --yaml -f ./_output/tmp/kubernetes_labels.yaml
❯ tree _output/
_output/
├── bin
│   └── platforms
│       └── linux
│           └── amd64
│               ├── exporter
│               └── syncer
├── tmp
│   └── kubernetes_labels.yaml
└── tools
    ├── addlicense
    └── go-gitlint

❯ ./syncer --help
```


#### Actions Auto Sync labels

For a more convenient way, we can use actions to sync labels directly, which is faster and easier

An example workflow is here:

```yaml
name: Sync labels
on:
  push:
    branches:
      - master
    paths:
      - path/to/manifest/labels.yml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Github lables pull and synchronize
        uses: kubecub/github-label-syncer@v2.0.0
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          manifest: path/to/manifest/labels.yml
```


You can sync labels to multiple repositories at the same time, which is more convenient, but you'll need to provide a token.

For example, the following yaml file takes a `local.github/sync_labeler.yml` file and syncs it to `kubecub/github-label-syncer`, `kubecub/log`, etc. `BOT_GITHUB_TOKEN` is used as the token


```yaml
name: Kubecub Sync labels

on:
  push:
    branches:
      - main
    # paths:
    #   - .github/sync_labels.yml

jobs:
    build:
      runs-on: ubuntu-latest
      steps:
        - name: Checkout
          uses: actions/checkout@1.0.0
        - name: Github lables pull and synchronize
          uses: kubecub/github-label-syncer@v2.0.0
          with:
            manifest: .github/sync_labeler.yml
            token: ${{ secrets.BOT_GITHUB_TOKEN }}
            repository: |
              kubecub/github-label-syncer
              kubecub/log
          env:
            GITHUB_TOKEN: ${{ secrets.BOT_GITHUB_TOKEN }}
```

> **Note**:
> The new labels and colors will replace the old ones

You can add jobs.<job_id>.steps.with.prune: false in order to preserver all existing labels which is not mentioned in manifest, in this case when a label will be renamed old label will be not deleted.


#### Work with Docker

```bash
❯ make deploy
```


## 🕋 architecture diagram
```mermaid
graph LR
A[Config file] -->B(github-label-syncer)
B --> C{Sync labels}
C -->|Yes| D[Create/Update labels] 
C -->|No | E[Delete labels]
D --> F[Dry run]
F --> |Yes| G[Preview]
F --> |No| H[Apply changes]
E --> I[Dry run]
I --> |Yes| J[Preview]
I --> |No| K[Apply changes]
```

**MVC Architecture Design:**
```mermaid
flowchart TB

A[Config file]
A --> B[github-label-syncer]
B --> C{Sync labels}
C -->|Yes|D[<font color='green'>Create/Update </font> <br>labels]
C -->|No|E[<font color='red'>Delete </font><br> labels]  
D -.-> F[<font color='blue'>Dry run</font><br>Preview]
F --> G[Apply <br>changes]
E -.-> H[<font color='blue'>Dry run</font><br>Preview]
H --> I[Apply<br> changes]
B --> J[Label <br>PRs & <br>Issues]
```

## 🗓️ community meeting

We welcome everyone to join us and contribute to github-label-syncer, whether you are new to open source or professional. We are committed to promoting an open source culture, so we offer community members neighborhood prizes and reward money in recognition of their contributions. We believe that by working together, we can build a strong community and make valuable open source tools and resources available to more people. So if you are interested in github-label-syncer, please join our community and start contributing your ideas and skills!

We take notes of each [biweekly meeting](https://github.com/kubecub/github-label-syncer/issues/2) in [GitHub discussions](https://github.com/kubecub/github-label-syncer/discussions/categories/meeting), and our minutes are written in [Google Docs](https://docs.google.com/document/d/1nx8MDpuG74NASx081JcCpxPgDITNTpIIos0DS6Vr9GU/edit?usp=sharing).

github-label-syncer maintains a [public roadmap](https://github.com/kubecub/community/tree/main/roadmaps). It gives a a high-level view of the main priorities for the project, the maturity of different features and projects, and how to influence the project direction.

## 🤼‍ Contributing & Development

kubecub Our goal is to build a top-level open source community. We have a set of standards, in the [Community repository](https://github.com/kubecub/community).

If you'd like to contribute to this github-label-syncer repository, please read our [contributor documentation](https://github.com/kubecub/github-label-syncer/blob/main/CONTRIBUTING.md).

Before you start, please make sure your changes are in demand. The best for that is to create a [new discussion](https://github.com/kubecub/github-label-syncer/discussions/new/choose) OR [Slack Communication](https://join.slack.com/t/kubecub/shared_invite/zt-1se0k2bae-lkYzz0_T~BYh3rjkvlcUqQ), or if you find an issue, [report it](https://github.com/kubecub/github-label-syncer/issues/new/choose) first.


## 🚨 License

Kubecub is licensed under the MIT License, Version 2.0. See [LICENSE](https://github.com/kubecub/github-label-syncer/tree/main/LICENSE) for the full license text.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fkubecub%2Fkubecub.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fkubecub%2Fgithub-label-syncer?ref=badge_large)


## 🔮 Thanks to our contributors!

<a href="https://github.com/kubecub/github-label-syncer/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=kubecub/github-label-syncer" />
</a>
