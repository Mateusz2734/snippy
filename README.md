# snippy

## Introduction
`snippy` is a CLI tool designed to help developers manage and utilize code snippets efficiently. With `snippy`, you can add, retrieve, edit, and organize your snippets, making code reuse a breeze. It is thoroughly tested (around 80% coverage) and easy to use.

## Available Commands
- `snippy init` - Initialize local snippy source
- `snippy add` - Add a new snippet (get content from stdin, file, clipboard or flag)
- `snippy get` - Get a snippet
- `snippy list` - List all snippets
- `snippy search` - Search for a snippet (by name or extension)
- `snippy delete` - Delete a snippet
- `snippy edit` - Edit a snippet (change extension or content)
- `snippy backup` - Group of commands for managing backups
    - `snippy backup create` - Create a backup file
    - `snippy backup restore` - Restore snippets from a backup file
- `snippy favorite` - Group of commands for managing favorites
    - `snippy favorite add` - Add a snippet to favorites
    - `snippy favorite delete` - Remove a snippet from favorites
    - `snippy favorite list` - List all favorite snippets

## Technologies
1. **Languages:**
   - Go
2. **Libraries:**
   - github.com/stretchr/testify
   - github.com/urfave/cli/v2

## Future Improvements And Features
- Colorful output
- Tags for snippets

## License
`snippy` is released under the MIT License. See the [LICENSE.md](LICENSE.md) file for more details.