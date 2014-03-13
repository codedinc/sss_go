# SSS ported to Go

* To get started on this project, first do steps 1 and 2 below if you don't have go or golex installed, then start on the usage

## Usage

    $ make
    $ ./sss samples/plain.css

##Step 1 - How to Install Go

###Install gvm

```
bash < <(curl -s https://raw.github.com/moovweb/gvm/master/binscripts/gvm-installer)
```

###Figure out which version of go you want to install

```
gvm listall
```

###Install it

E.g., if you want to install version 1.2.1 of go:

```
gvm install go1.2.1
```

##Step 2 - Install Golex once you installed go via gvm

```
go get github.com/cznic/golex
```
