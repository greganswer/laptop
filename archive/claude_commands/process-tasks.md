# Task List Management

Guidelines for managing task lists in markdown files to track progress on completing a PRD

## Task Implementation

- **Create and switch to the branch:** if it does not already exist, create a new branch for the feature. The branch name should be the same as the feature folder name in `/planning`.
- **One sub-task at a time:** Do **NOT** start the next sub‑task until you complete the previous
- **Completion protocol:**  
  1. When you finish a **sub‑task**, immediately mark it as completed by changing `[ ]` to `[x]`.  
  2. If **all** subtasks underneath a parent task are now `[x]`, also mark the **parent task** as completed. If we're not on the main branch, commit and push the changes.
  3. If **all** parent tasks are now `[x]`, move the feature folder from `/planning` to `/archive` folder in the root directory. DO NOT move any other files or folders.
  4. Start working on the next sub-task or task in the file. DO NOT wait for user input within the same task

## Task List Maintenance

1. **Update the task list as you work:**
   - Mark tasks and subtasks as completed (`[x]`) per the protocol above.
   - Add new tasks as they emerge.

2. **Maintain the “Relevant Files” section:**
   - List every file created or modified.
   - Give each file a one‑line description of its purpose.

## AI Instructions

When working with task lists, the AI must:

1. Clear the current context before starting a new task.
2. Regularly update the task list file after finishing any significant work.
3. Follow the completion protocol:
   - Mark each finished **sub‑task** `[x]`.
   - Mark the **parent task** `[x]` once **all** its subtasks are `[x]`.
4. Add newly discovered tasks.
5. Keep “Relevant Files” accurate and up to date.
6. Before starting work, check which sub‑task is next.
7. After implementing a sub‑task, update the file
