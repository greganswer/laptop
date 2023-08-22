# Workflow

Tired of forgetting to transition Jira tickets? Want to have consistent Merge Request info in GitLab? Look no further!

Workflow is a set of Bash scripts that helps automate these admin tasks so you can focus on coding.

## Installation

Download one or more of the scripts you want to use into a folder of your choice. You can symlink them into your binary folder with the following:

```bash
"$PWD/start" /usr/local/bin/gstart
"$PWD/review" /usr/local/bin/greview
"$PWD/draft" /usr/local/bin/gdraft
```

I prefixed them with `g` so that they don't conflict with any other binaries.

## Usage/Examples

My workflow looks like the following:

```bash
gstart c2-123
gdraft
greview
```

## FAQ

### Why not write this in language X?

In the past, I have written this in Ruby, Python, Go, and Typescript. They all had too many dependencies to install. Too many files to read. I just want something I can load on my machine and start using without too many version compatibility issues.

Bash, though hard to read, keeps it pretty simple. You can use the AI tool-of-the-day to help you read/write Bash.

### Can you add feature X?

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

With the MIT license, you're also free to modify your local scripts to your heart's desire!

## License

[MIT](https://choosealicense.com/licenses/mit/)

