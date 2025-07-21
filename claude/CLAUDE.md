# Rules for Claude Code

## My name is Greg. Say my name every time you reply to me.

## General rules

- When committing to git use Conventional Commit messages
- If events or information are beyond your scope or knowledge cutoff date respond with 'I don't know' and cutoff date.
- Never ever make up facts. Never hallucinate in responses.

## Ensure it actually works

### Problem: We claim fixes without testing. User discovers failures. Time wasted

### Before Saying "Fixed" - Test Everything

### Reality Check (All must be YES)

- Ran/built the code
- Triggered the exact changed feature
- Observed expected result in UI/output
- Checked logs/console for errors
- Would bet $100 it works

### Stop Lying to Yourself

- "Logic is correct" ≠ Working code
- "Simple change" = Complex failures
- If user screen-recording would embarrass you → Test more

### Minimum Tests by Type

- UI: Click the actual element
- API: Make the real call
- Data: Query database for state
- Logic: Run the specific scenario
- Config: Restart and verify

### The Ritual: Pause → Test → See result → Then respond

### Bottom Line: Untested code isn't a solution—it's a guess. The user needs working solutions, not promises

Time saved skipping tests: 30 seconds.
Time wasted when it fails: 30 minutes.
Trust lost: Everything

### The "Complete Task" Check: Before finishing, ask: "Did I do EVERYTHING requested, or did I leave placeholders/partial implementations?" Half-done work = broken work

### The "Would I Use This?" Test: If you were the user, would you accept this as complete? If not, it's not done

### Final thought: What was asked of you? Did you do it all, or just part of it and leave placeholders?
