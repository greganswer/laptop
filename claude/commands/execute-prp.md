---
description: Execute a Project Request Plan (PRP) file to implement a feature with validation
---

# Execute BASE PRP

Implement a feature using using the PRP file.

## PRP File: $ARGUMENTS

## Execution Process

1. **Create a branch**
    - Create the branch off the `main` branch
    - If the user has unsaved changes, abort and notify the user
1. **Load PRP**
   - Read the specified PRP file
   - Understand all context and requirements
   - Follow all instructions in the PRP and extend the research if needed
   - Ensure you have all needed context to implement the PRP fully
   - Do more web searches and codebase exploration as needed
1. **ULTRATHINK**
   - Think hard before you execute the plan. Create a comprehensive plan addressing all requirements.
   - Break down complex tasks into smaller, manageable steps using your todos tools.
   - Use the TodoWrite tool to create and track your implementation plan.
   - Identify implementation patterns from existing code to follow.
1. **Execute the plan**
   - Execute the PRP
   - Implement all the code
1. **Validate**
   - Run each validation command
   - Fix any failures
   - Re-run until all pass
1. **Complete**
   - Ensure all checklist items done
   - Run final validation suite
   - Report completion status
   - Read the PRP again to ensure you have implemented everything
1. **Reference the PRP**
   - You can always reference the PRP again if needed
1. **Create a pull request**
   - Create a pull request with the branch you created
   - Ensure the PR description references the PRP file and includes any relevant information
   - Include a link to the PR in the PRP file for future reference

Note: If validation fails, use error patterns in PRP to fix and retry.
