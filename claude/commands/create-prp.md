---
description: Generate a Project Request Plan (PRP) from initial requirements
---

# Create PRP

Generate a comprehensive PRP from `PRPs/INITIAL.md` for systematic feature implementation.

## Quick Reference

`Read Initial → Research → Plan → Generate PRP`

## Process

### 1. Pre-flight Checks

- Verify `PRPs/INITIAL.md` exists (abort if missing)
- Check for `PRPs/templates/prp_base.md` template (abort if missing)
- Read initial requirements completely

### 2. Research Phase

**Codebase Discovery:**

- Find similar patterns/features
- Identify conventions and test patterns
- Note EXACT file paths and line numbers to reference

**External Research (if needed):**

- Documentation URLs (be specific)
- Implementation examples
- Known issues and solutions

### 3. PRP Content Requirements

**Context Section:**

- Relevant documentation links
- Code snippets from codebase
- Library versions and quirks
- Existing patterns to follow

**Implementation Plan:**

- Clear, numbered task list (1, 2, 3...)
- Concrete code examples (not pseudocode)
- Specific error cases to handle
- Exact file:line references for patterns

**Validation Commands:**

- Must be executable
- Include all test/lint commands
- Example format varies by language

### 4. Output Specification

**Location:** `PRPs/` directory
**Format:** `###-feature-name.md`

- Sequential numbering (001, 002, etc.)
- Feature name under 5 words
- Check existing PRPs for next number

## Critical Steps

1. **ULTRATHINK** before writing PRP
2. Research thoroughly - be SPECIFIC with references
3. Include ALL context for one-pass success
4. Make validation gates executable
5. Follow KISS, YAGNI, DRY principles

## Quality Gates

- [ ] Initial requirements understood
- [ ] Research findings documented
- [ ] Context enables self-validation
- [ ] Clear implementation path
- [ ] Executable validation commands

The PRP must be self-contained - the executing agent only has access to what you provide.
