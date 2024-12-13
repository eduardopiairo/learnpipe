# Zsh 

## Plugins

### Syntax Highlight

[zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting)

Install command:
```
brew install zsh-syntax-highlighting
```

Enable the plugin by running:
```
echo "source $(brew --prefix)/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh" >> ${ZDOTDIR:-$HOME}/.zshrc
```


### Auto Suggestions

[zsh-autosuggestions](https://github.com/zsh-users/zsh-autosuggestions4)

Install command:
```
brew install zsh-autosuggestions
```

Activate autosuggestions:
```
echo "source $(brew --prefix)/share/zsh-autosuggestions/zsh-autosuggestions.zsh" >> ${ZDOTDIR:-$HOME}/.zshrc
```


### History Commands

[zsh-history-substring-search](https://github.com/zsh-users/zsh-history-substring-search)

Install command:
```
brew install zsh-history-substring-search
```

```
echo "source $(brew --prefix)/share/zsh-history-substring-search/zsh-history-substring-search.zsh" >> ~/.zshrc
```