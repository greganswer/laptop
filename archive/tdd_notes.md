
## 10. Test Driven Development (TDD)

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
