---
description: Execute a Project Request Plan (PRP) file to implement a feature with validation
---

# Execute PRP

Execute a Project Request Plan (PRP) file to implement features systematically with validation.

## PRP File: $ARGUMENTS

## Execution Process

### 1. Pre-flight Checks

- Check git status for uncommitted changes - abort if dirty working directory
- Verify PRP file exists and is readable
- Create feature branch from `main` branch

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
- Re-run validation until all checks pass
- Verify all PRP requirements implemented

### 6. Completion

- Review PRP file to confirm all items addressed
- Run final validation suite
- Create pull request with:
  - Descriptive title referencing PRP
  - Summary of changes made
  - Link to PRP file
  - Any additional context needed

### Error Handling

- If validation fails, analyze error patterns from PRP
- Fix issues and retry validation
- If blocked, reference PRP for guidance or escalate

### Notes

- PRP file can be referenced at any time during execution
- Maintain clean commit history with conventional commits
- Abort execution if critical errors encountered
