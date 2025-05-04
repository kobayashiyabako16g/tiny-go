# Tiny Go Project

## Prerequisites

- [Go](https://go.dev/dl/)
- [Neovim](https://neovim.io/)
- A Neovim plugin manager (e.g., [lazy.nvim](https://github.com/folke/lazy.nvim), [packer.nvim](https://github.com/wbthomason/packer.nvim))

## 1. Install Go

Download and install Go from the official website:

👉 https://go.dev/dl/

Make sure the `go` command is available in your terminal:

```bash
go version
```

## 2. Install gopls (Go Language Server)

Install `gopls` with:

```bash
go install golang.org/x/tools/gopls@latest
```

Make sure the Go bin directory is in your `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## 3. Install Neovim

Follow the instructions on the official Neovim site:

👉 https://neovim.io/

## 4. Configure Neovim for Go

Install [nvim-lspconfig](https://github.com/neovim/nvim-lspconfig) and set up `gopls`. Example:

```lua
require('lspconfig').gopls.setup{}
```

You may also want to install:

- `nvim-treesitter/nvim-treesitter` (syntax highlighting)
- `hrsh7th/nvim-cmp` (completion)
- `neovim/nvim-lspconfig` (LSP support)

## 5. Verify

Open a `.go` file in Neovim and check that autocompletion, go-to-definition, etc., are working. Use `:LspInfo` to verify `gopls` is attached.

---

## Troubleshooting

- Check `gopls version` to verify installation.
- Use `:LspInfo` in Neovim to confirm LSP status.
- Ensure `$GOPATH/bin` or `~/go/bin` is in your `PATH`.

## Resources

- [Go Official Site](https://go.dev/)
- [gopls Documentation](https://github.com/golang/tools/tree/master/gopls)
- [nvim-lspconfig](https://github.com/neovim/nvim-lspconfig)

