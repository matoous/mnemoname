# Mnemoname

[![Build Status](https://github.com/matoous/mnemoname/workflows/Tests/badge.svg)](https://github.com/matoous/mnemoname/actions) 
[![Build Status](https://github.com/matoous/mnemoname/workflows/Lint/badge.svg)](https://github.com/matoous/mnemoname/actions) 
[![GoDoc](https://godoc.org/github.com/matoous/mnemoname?status.svg)](https://godoc.org/github.com/matoous/mnemoname)
[![GitHub issues](https://img.shields.io/github/issues/matoous/mnemoname.svg)](https://github.com/matoous/mnemoname/issues)
[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/matoous/mnemoname/LICENSE)


**Mnemoname** is tiny golang utility for obtaining mnemonic names. It is based on following
comment from HN by `curtis3389`: https://news.ycombinator.com/item?id=26054951 on a discussion
of [RFC 1178: Choosing a name for your computer (1990)](https://news.ycombinator.com/item?id=26054014).
The list of the mnemonic names itself is from [following article](https://web.archive.org/web/20090918202746/http://tothink.com/mnemonic/wordlist.html)
by _Oren Tirosh_.

> _mnemonic_ - a system such as a pattern of letters, ideas, or associations which assists in remembering something.

---

DISCLAIMER: The generated names are not cryptographically secure and shouldn't be used as such.

NOTE: [A little copying is better than a little dependency](https://go-proverbs.github.io/) so you
might be better of with just copying these few files into your project 🤷‍.

## Install

Via go get tool

``` bash
> go get github.com/matoous/mnemoname
```

## Usage

Generate random mnemonic name

``` go
name := mnemoname.New()
```

Get multiple mnemonic names at ones:

``` go
names, err := mnemoname.NewN(13)
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
