
## Claude Code hooks

In the `claude/settings.json` file, please create the following hooks:

- When Claude Code is awaiting my input, then say "Greg, I need your input."
- When Claude Code finishes a task, then say "Task Complete."
- When Claude Code finishes writing to a ruby file, then run Rubocop on that file.
- When Claude Code finishes writing to a TypeScript file, then run prettier on that file.
