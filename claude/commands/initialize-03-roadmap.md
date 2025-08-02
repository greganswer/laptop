---
description: Break down an MVP document into organized PRP files for sequential feature development
---

# Generate PRP files from an MVP

## Goal

To guide an AI assistant in creating a set of files and folders that will be used for code development.

## Input

The user will provide the MVP contents or file name.

<mvp>
$ARGUMENTS
</mvp>

## Output

- **Format:** Markdown (`.md`)
- **Location:** `PRPs` in the root directory
- **Filename:** `[ordered-id]-feature-[feature-name].md`. (e.g., `001-feature-user-authentication.md`)
- Keep the feature name less than 5 words long

## Process

1. Verify that there are one or more user stories. If not, abort and notify the user.
1. Determine the logical order in which the stories should be completed.
    - Consider things like dependencies, complexity, and user flow.
    - Prefer to develop one small fullstack feature at a time (including minimal database schema, backend API, and frontend UI).
1. Create files as described in the "Output" section.
    - Each file is numbered sequentially based on the order the user stories should be completed
    - If the user stories are not numbered, assign them a number based on the order they should be completed
    - If the location already contains any files with a number, increment the highest number by 1 and create new files sequentially
1. Use subagents to add the relevant content to the correct files, including:
    - A brief overview of the feature
    - The user stories in Gherkin format with Acceptance Criteria beneath each story
    - Any relevant information from the `<mvp>` section that is needed to complete the feature

## Final instructions

1. Do NOT write any new information. Only use the information provided in the MVP.
