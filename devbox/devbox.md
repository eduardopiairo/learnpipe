# Devbox

[Devbox](https://www.jetify.com/devbox) is a command-line tool that lets you easily create isolated shells for development.

Devbox requires the [Nix Package Manager](https://nixos.org/download/). If Nix is not detected when running a command, Devbox will install it in multi-user mode for macOS.


## 0. Install Devbox
```
curl -fsSL https://get.jetify.com/devbox | bash
```


## 1. Initilize DevBox
```
devbox init
``` 
This creates a `devbox.json` file in the current directory. You should commit it to source control.

## 2. Add a package to your project
```
devbox add <package>@<version>
```

## 3. Launch your environment
```
devbox shell
```

## 4. To exit the Devbox shell and return to your regular shell:
```
exit
```

To share your project and shell, make sure to check in your `devbox.json` and `devbox.lock` file into source control. These files will ensure that developers get the same packages and environment when they run your project.

