-- README.md --
# newtoniansheep

A command-line tool for managing image links in markdown files.

## Install

On macOS/Linux:
```bash
brew install gkwa/homebrew-tools/newtoniansheep
```

On Windows:
```powershell
TBD
```

## Expected File Format

The input file should be a markdown file containing image links in the following format:

```markdown
![](http://image.hosting.com/myrecipe1/1.jpg)

![](http://image.hosting.com/myrecipe1/2.jpg)

![](http://image.hosting.com/myrecipe1/3.jpg)
```

Each image link should be on its own line. The tool processes these types of image links.

## Cheatsheet

Deduplicate image links:
```bash
newtoniansheep deduplicate '/path/to/your/file.md'
```

Randomize image links:
```bash
newtoniansheep randomize '/path/to/your/file.md'
```

Get version information:
```bash
newtoniansheep version
```

For more detailed information on each command:
```bash
newtoniansheep help
newtoniansheep help deduplicate
newtoniansheep help randomize
```

## Example Output

After running a command, you'll see output similar to this:

```
/Users/username/Documents/file.md is 1.7 MB with 26,163 lines and 13,078 links and 2 duplicates removed
```

This shows the file path, size, number of lines, number of links, and how many duplicates were removed (for the deduplicate command).