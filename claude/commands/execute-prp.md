---
description: Execute a Project Request Plan (PRP) file to implement a feature with validation (auto-selects next PRP if no file specified)
---

# Execute PRP

Execute a Project Request Plan (PRP) file to implement features systematically with validation.

## Quick Reference

`Read PRP → Plan → Implement → Validate → Create PR`

## PRP File Selection

- **If no argument provided:** Automatically selects the next PRP in sequential order from `PRPs/` directory
- **If argument provided:** Uses the specified PRP file: $ARGUMENTS

## Execution Process

### 1. Pre-flight Checks

- Check git status for uncommitted changes - abort if dirty working directory
- **PRP Selection:**
  - If no argument provided:
    - List all PRP files in `PRPs/` directory (format: `###-*.md`)
    - Exclude files in `PRPs/archive/` subdirectory
    - Sort numerically and select the lowest numbered PRP
    - Display: "Auto-selected PRP: ###-feature-name.md"
    - Abort if no PRPs found with message: "No PRP files found in PRPs/ directory"
  - If argument provided:
    - Use the specified PRP file path
    - Verify the file exists and is readable
- Create branch from `main` branch (format: `prp-###-feature-name`, e.g., `prp-003-user-auth`)

### 2. Load & Analyze PRP

- Read the PRP file completely
- Understand all requirements, context, and acceptance criteria
- Perform additional research if gaps identified:
  - Web searches for external dependencies/patterns
  - Codebase exploration for existing patterns
  - Documentation review

### 3. Plan Implementation (ULTRATHINK)

- Create comprehensive implementation plan using TodoWrite tool
- Break complex tasks into manageable steps
- Identify existing code patterns to follow
- Map out dependencies and execution order
- Consider edge cases and error scenarios

### 4. Execute Implementation

- Work through todo list systematically
- Implement all required code changes
- Follow existing code conventions and patterns
- Update todo status as work progresses

### 5. Validation & Testing

- Run all validation commands specified in PRP
- Fix any test failures or linting issues
  - If tests fail, use @claude/agents/debug-specialist to diagnose and fix issues
- Re-run validation until all checks pass
- Verify all PRP requirements implemented

### 6. Completion

- Review PRP file to confirm all items addressed
- Run final validation suite
- Move PRP file to archive folder:
  - Create `PRPs/archive` folder if it doesn't exist
  - Move the executed PRP file to `PRPs/archive/`
- Update CHANGELOG.md if present:
  - Check if CHANGELOG.md exists in project root
  - If exists: Add entry under "[Unreleased]" section
  - Format: `- [PRP-###] Description` under appropriate category (Added/Changed/Fixed/etc.)
  - Commit with message: `docs: update CHANGELOG for PRP-###`
  - If not present: Skip (no error)
- Create pull request with:
  - **PR Title Format**: `<type>: [PRP-###] <description>`
    - Types: `feat`, `fix`, `docs`, `style`, `refactor`, `test`, `chore`
    - Example: `feat: [PRP-003] Add user authentication system`
  - Summary of changes made
  - Any additional context needed

### Error Handling

- Validation fails → Fix and retry
- Critical blocker → Check PRP guidance or abort

### Notes

- PRP file can be referenced at any time during execution
- Maintain clean commit history with conventional commits
- Abort execution if critical errors encountered
