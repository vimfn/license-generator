<h1 align="center">License Generator</h1> 

![iamagif](https://user-images.githubusercontent.com/102473837/210200111-54029bc3-faf6-4123-8164-ffe1afc03e01.gif)

### Overview

This is a blazing fast ⚡, command line license generator for your open source projects written in Go.

I know that GitHub has a great GUI to add licenses to projects but I always found myself doing too much work. First you have to go to GitHub, create a file, type 'LICENSE', pick a license, push it and then pull it locally. With this you can just generate the license locally and push it to GitHub.

This is my first Go program, so I made this simple project (still very beginner).

### Installation

#### From source

```bash
git clone https://github.com/its-ag/license-generator
cd license-generator
go build gen-license.go
```

If you don't have go installed, you can download the executable from the [releases](hhttps://github.com/its-ag/license-generator/releases/tag/Latest) section.

### Usage

```bash
gen-license
```

### Contributing

- Fork the repository
- Create a branch
  ```bash
  git checkout -b fix/amazingFix
  ```
- Commit your changes and push to your branch

  ```bash
  git commit -m "made an amazingFix"
  git push origin fix/amazingFix
  ```
- Open a pull request

[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fits-ag%2Flicense-generator&count_bg=%2379C83D&title_bg=%23555555&icon=visualstudiocode.svg&icon_color=%23E7E7E7&title=&edge_flat=false)](https://hits.seeyoufarm.com)
