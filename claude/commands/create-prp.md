# Create PRP

## Feature file

> [!IMPORTANT]
> If `PRPs/INITIAL.md` file does not exist, abort and notify the user

Using `PRPs/INITIAL.md`, generate a complete PRP for general feature implementation with thorough research. Ensure context is passed to the AI agent to enable self-validation and iterative refinement. Read the feature file first to understand what needs to be created, how the examples provided help, and any other considerations.

The AI agent only gets the context you are appending to the PRP and training data. Assume the AI agent has access to the codebase and the same knowledge cutoff as you, so its important that your research findings are included or referenced in the PRP. The Agent has Websearch capabilities, so pass urls to documentation and examples.

## Research Process

1. **Codebase Analysis**
   - Search for similar features/patterns in the codebase
   - Identify files to reference in PRP
   - Note existing conventions to follow
   - Check test patterns for validation approach

2. **External Research**
   - Search for similar features/patterns online
   - Library documentation (include specific URLs)
   - Implementation examples (GitHub/StackOverflow/blogs)
   - Best practices and common pitfalls

3. **User Clarification** (if needed)
   - Specific patterns to mirror and where to find them?
   - Integration requirements and where to find them?

## PRP Generation

> [!IMPORTANT]
> If the `PRPs/templates/prp_base.md` template file does not exist, abort and notify the user

Using the template file:

### Critical Context to Include and pass to the AI agent as part of the PRP

- **Documentation**: URLs with specific sections
- **Code Examples**: Real snippets from codebase
- **Gotchas**: Library quirks, version issues
- **Patterns**: Existing approaches to follow

### Implementation Blueprint

- Start with pseudocode showing approach
- Reference real files for patterns
- Include error handling strategy
- list tasks to be completed to fulfill the PRP in the order they should be completed

### Validation Gates (Must be Executable) eg for python

```bash
# Syntax/Style
ruff check --fix && mypy .

# Unit Tests
uv run pytest tests/ -v

```

***CRITICAL AFTER YOU ARE DONE RESEARCHING AND EXPLORING THE CODEBASE BEFORE YOU START WRITING THE PRP***

***ULTRATHINK ABOUT THE PRP AND PLAN YOUR APPROACH THEN START WRITING THE PRP***

## Output

- Each file is numbered sequentially based on the order the feature should be completed
- If the files are not numbered, assign them a number based on the order they should be completed
- If the location already contains any files with a number, increment the highest number by 1 and create new files sequentially

- **Format:** Markdown (`.md`)
- **Location:** `PRPs` in the root directory
- **Filename:** `[ordered-id]-[feature-name].md`. (e.g., `001-user-authentication.md`)
- Keep the feature name less than 5 words long

## Quality Checklist

- [ ] All necessary context included
- [ ] Validation gates are executable by AI
- [ ] References existing patterns
- [ ] Clear implementation path
- [ ] Error handling documented
- [ ] For implementation, focus on KISS, YAGNI, and DRY principles.

Score the PRP on a scale of 1-10 (confidence level to succeed in one-pass implementation using claude codes)

Remember: The goal is one-pass implementation success through comprehensive context.
