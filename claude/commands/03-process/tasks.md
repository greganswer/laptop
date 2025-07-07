
# Task List Management

Guidelines for managing task lists in markdown files to track progress on completing a PRD

## Task Implementation

- **One sub-task at a time:** Do **NOT** start the next sub‑task until you complete the previous
- **Completion protocol:**  
  1. When you finish a **sub‑task**, immediately mark it as completed by changing `[ ]` to `[x]`.  
  2. If **all** subtasks underneath a parent task are now `[x]`, also mark the **parent task** as completed.  
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
8. Follow Test Driven Development (TDD) principles to ensure code quality and maintainability.

## Test Driven Development (TDD)

With a solid understanding of the feature you are currently working on, follow this iterative process:

1. Red - Write a failing test that enforces new desired behavior. Run it, it should fail. Do not modify non-test code in the Red phase!
2. Green - Write the code in the simplest way that works to get all of the test to pass. Do not modify tests in the Green phase!
test & commit to git (only when all tests are passing) but don't push
3. Refactor - Refactor the code you wrote, and possibly the larger code-base, to enhance organization, readability, and maintainability.
This is an opportunity to improve the quality of the code. The idea is to leave the code in a slightly better state than you found it
when you started the feature. Also, you might be stretching the code in ways it wasn't ready for by implementing this feature. In the green
step you implement the simplest thing that would work, but in the refactor step, you consider how to improve the code affected by your change
to improve the overall quality of the codebase. Follow Martin Fowler's guidance on this.
a minor refactor can be committed in a single commit
major refactorings should be committed in stages
test & commit to git (only when all tests are passing)

Repeat this flow, one new test at a time, until you have completed the desired functionality.

