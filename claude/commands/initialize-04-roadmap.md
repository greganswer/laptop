# Generate files and folders from an MVP

## Goal

To guide an AI assistant in creating a set of files and folders that will be used for code development.

## Input

The user will provide the MVP contents or file name.

<mvp>
$ARGUMENTS
</mvp>

## Output

- **Format:** Markdown (`.md`)
- **Location:** `planning/[ordered-id]-[feature-name]/requirements.md` in the projects root directory. (e.g., `planning/001-user-authentication/requirements.md`)
- **Filename:** `requirements.md`
- Keep the folder name less than 5 words long

## Process

1. Verify that there are one or more user stories. If not, abort and notify the user.
1. Determine the logical order in which the stories should be completed.
    - Consider things like dependencies, complexity, and user flow.
    - Prefer to develop the database schema, then the API, then the frontend.
1. Create folders as described in the "Output" section.
    - Each folder is numbered sequentially based on the order the user stories should be completed
    - If the user stories are not numbered, assign them a number based on the order they should be completed
    - If the location already contains any folders with a number, increment the highest number by 1 and create new folders sequentially
1. Use subagents to add the relevant content to the `requirements.md` file in each folder, including:
    - A brief overview of the feature
    - The user stories in Gherkin format with Acceptance Criteria beneath each story
    - Any relevant information from the `<mvp>` section that is needed to complete the feature

## Final instructions

1. Do NOT write any new information. Only use the information provided in the MVP.
