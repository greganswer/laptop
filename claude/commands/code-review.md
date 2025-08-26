# Code Review Command

Analyze codebase and provide prioritized, actionable feedback.

## Review Focus

**Scan for:**

- Security vulnerabilities
- Performance issues
- Code quality problems
- Architecture concerns
- Testing gaps
- Bug risks

**Priority levels:**

- Must: Security, breaking bugs, major performance
- Should: Code quality, architecture issues, minor bugs, style, missing tests
- Could: Documentation, optimizations

## Task Management

1. Read existing `TASKS.md`
1. Add findings to appropriate priority sections
1. Use the same format as existing tasks
1. If no format exits, use format: `**[TYPE]** Description in file:line`
1. IMPORTANT: Don't duplicate existing tasks
1. Mark verified completions as `[x]`

## Process

1. Scan entire codebase
2. Read current `TASKS.md`
3. Categorize findings by priority
4. Update `TASKS.md` with actionable tasks
5. Provide review summary

## Guidelines

- Be specific: Include file paths and line numbers
- Explain impact and suggested solutions
- Focus on security, performance, maintainability
- Apply language/framework best practices
- Consider effort-to-benefit ratio
