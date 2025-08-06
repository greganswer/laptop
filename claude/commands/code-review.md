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

- 🔴 Critical: Security, breaking bugs, major performance
- 🟠 High: Code quality, architecture issues
- 🟡 Medium: Minor bugs, style, missing tests
- 🟢 Low: Documentation, optimizations

## Task Management

1. Read existing `TASKS.md`
2. Add findings to appropriate priority sections
3. Use format: `**[TYPE]** Description in file:line`
4. IMPORTANT: Don't duplicate existing tasks
5. Mark verified completions as `[x]`

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
