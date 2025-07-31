# Generate feature request

## Goal

To guide an AI assistant in creating a feature request markdown file.

## Input

The user will provide the feature request contents or file name.

<feature>
$ARGUMENTS
</feature>

## Output

- **Format:** Markdown (`.md`)
- **Location:** `PRPs/INITIAL.md` in the projects root directory
- **Filename:** `INITIAL.md`

## Process

1. Truncate or create the output file
1. Read the input and come up with a feature name. Use less than 5 words.
1. Extract the feature into the following format:

```md
# <FEATURE_NAME>

## FEATURE

[Describe what you want to build - be specific about functionality and requirements]

## EXAMPLES

[List any example files in the examples/ folder and explain how they should be used]

## DOCUMENTATION

[Include links to relevant documentation, APIs, or MCP server resources]

## OTHER CONSIDERATIONS

[Mention any gotchas, specific requirements, or things AI assistants commonly miss]
```
