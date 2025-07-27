Please review this background worker for:

- Ensure Idempotentcy (Example: duplicates, cancellations, etc.)
- Handling retries exhausted (logs, data cleanup, etc.)
- Execution time
  - Under 10 seconds for workers that make API calls
  - Under 1 second for all other workers
