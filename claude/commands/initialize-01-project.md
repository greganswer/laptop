---
description: Initialize a full-stack project with Next.js frontend, Rails backend, and Playwright e2e tests
---

# Project Initialization Command

Use sub agents where applicable.

Initialize the project with the following structure:

```plain
.
├── backend
│   ├── spec
│   │   └── CLAUDE.md
│   └── CLAUDE.md
├── e2e
│   └── CLAUDE.md
├── frontend
│   └── CLAUDE.md
├── PRPs
│   └── archive
│       └── MVP
│           └── requirements.md
├── CHANGELOG.md
├── CLAUDE.md
├── Procfile.dev
├── TASKS.md
└── package.json
```

- Initialize a Next.js 15+ project in `frontend/` with TypeScript and shadcn components
- Add a Next.js `frontend/.gitignore` file

- Initialize a Ruby on Rails 8+ API only project in `backend/` with PostgreSQL, Redis, RSpec, and Sidekiq
- Add a Ruby on Rails `backend/.gitignore` file
- Remove the `backend/test` folder

- Initialize a Playwright project in `e2e/`
- Add a Playwright `e2e/.gitignore` file
- Ensure Playwight is configured to take screenshots on failure

## File contents

Ensure `TASKS.md` includes the following:

```md
# Tasks

## 🔴 Critical Priority

## 🟠 High Priority

## 🟡 Medium Priority

## 🟢 Low Priority
```

Ensure `Procfile.dev` includes the following:

```yaml
backend: cd backend && RAILS_LOG_LEVEL=info bundle exec rails s -p 3000
frontend: cd frontend && PORT=3001 npm run dev
```

Ensure `backend/Gemfile` includes the following (merge any existing groups):

```ruby
gem "active_model_serializers"
gem "dotenv-rails"
gem "kaminari" # pagination
gem "lograge"
gem "pundit", "~> 2.3"
gem "rack-cors" 
gem "redis"

group :development, :test do
  gem "factory_bot_rails"
  gem "faker"
  gem "rspec-rails"
  gem "rubocop-rspec", require: false
  gem "rubocop-rspec_rails", require: false
end

group :test do
  gem "pundit-matchers"
  gem "rspec-sidekiq"
  gem "shoulda-matchers"
end
```

Ensure `backend/.rubocop.yml` includes the following:

```yaml
# https://github.com/rubocop/ruby-style-guide

# Omakase Ruby styling for Rails
inherit_gem: { rubocop-rails-omakase: rubocop.yml }

require:
  - rubocop-rspec
  - rubocop-rspec_rails

ExtraSpacing:
  Enabled: true

# Layout
# https://docs.rubocop.org/rubocop/cops_layout.html
Layout/IndentationConsistency:
  Enabled: true
  EnforcedStyle: normal
Layout/IndentationWidth:
  Enabled: true
Layout/EmptyLines:
  Enabled: true
Layout/MultilineHashBraceLayout:
  Enabled: true
Layout/MultilineHashKeyLineBreaks:
  Enabled: true
Layout/ArgumentAlignment:
  Enabled: true
Layout/EmptyLinesAroundAccessModifier:
  Enabled: true
  EnforcedStyle: around
Layout/EmptyLineBetweenDefs:
  Enabled: true
Layout/EmptyLinesAroundAttributeAccessor:
  Enabled: true
Layout/AccessModifierIndentation:
  Enabled: true
Layout/SpaceAroundOperators:
  Enabled: true

# Style
Style/BlockComments:
  Enabled: false

# RSpec
RSpec/EmptyLineAfterHook:
  Enabled: true
  AllowConsecutiveOneLiners: false
RSpec/EmptyLineAfterSubject:
  Enabled: true
RSpec/EmptyLineAfterFinalLet:
  Enabled: true
```

Ensure `frontend/CLAUDE.md` includes the following:

```markdown
# Frontend Claude rules

- Always use shadcn components
- Use strict TypeScript with explicit return types
- Define interfaces for all data structures
- Use meaningful component and variable names
- Prefer async/await over promise chains
- Use TSDoc for component documentation
```

Ensure `backend/CLAUDE.md` includes the following:

```markdown
# Backend Claude Rules

## 1. Core Principles
- RESTful endpoints (e.g., `GET /users?search=jo`).
- Pundit guards every action.
- Skinny controllers; fat models and service objects.

## 2. Code Organization
- Queries live in models: `User.search(params)`.
- Business logic lives in `app/commands`.
- Shared behavior → concerns.
- Long-running work → Sidekiq jobs.

## 3. Style & Conventions
- Idiomatic Ruby 3.x; follow Rails guides.
- `snake_case` files & methods, `CamelCase` classes.
- No duplicated code; prefer modules/iterations.

## 4. Error Handling
- Raise only for exceptional cases.
- Model validations + helpful messages.
- Log errors with context.

## 5. Performance
- Add indexes.
- Eager-load to kill N+1.
- Use `includes`, `joins`, `select` wisely.

## 6. Testing
- RSpec (FactoryBot, Shoulda, Pundit, Sidekiq).
- TDD/BDD preferred.

## 7. Security
- Devise (Token/JWT) + Pundit.
- Strong params.
- Guard against XSS, CSRF, SQLi.

## 8. Stack Snapshot
- Rails 8.0.1 + PostgreSQL + Redis.
- JSON: ActiveModel::Serializer, Kaminari pagination, Rack-CORS.
- Background: Sidekiq (+ Cron).
- Payments: Pay + Stripe.
- Auditing: Paranoia, PaperTrail.
- Dev/Test: RSpec, Faker, RuboCop, Brakeman, Lograge.

## 9. Configuration
- `.env` via Dotenv.
- Always read environment variables with `ENV.fetch` and raise an error if not set

## 10. Controller Size Limits & Extraction Rules

### Size Limits

- Controllers MUST NOT exceed 100 lines
- Actions MUST NOT exceed 20 lines
- Private methods MUST NOT exceed 15 lines

### Mandatory Extractions

When controllers exceed limits, extract in this order:

1. **Service Objects** (`app/services/`) - Complex business logic
2. **Form Objects** (`app/forms/`) - Parameter validation & processing
3. **Concerns** (`app/controllers/concerns/`) - Shared controller behavior
4. **Value Objects** (`app/values/`) - Data structures & formatting

### Controller Responsibility

Controllers should ONLY:

- Route requests to appropriate services
- Handle HTTP responses/redirects
- Manage session state
- Set cookies/headers
```

Ensure `backend/spec/rails_helper.rb` includes the following:

```ruby
  # Call `create` instead of `FactoryBot.create` in spec files
  config.include FactoryBot::Syntax::Methods
  
  # Shoulda matchers. https://matchers.shoulda.io/docs/v6.4.0
  Shoulda::Matchers.configure do |config|
   config.integrate do |with|
     with.test_framework :rspec
     with.library :rails
   end
 end
```

Ensure `backend/spec/CLAUDE.md` includes the following:

```markdown
# Backend Tests

## RSpec Best Practices (betterspecs.org)

### Testing Structure and Organization

1. **Describe your methods** - Keep your description clear and concise
2. **Use context blocks** - Group related tests together with context
3. **Use shared examples** - DRY your test suite up with shared examples
4. **Test valid, edge and invalid cases** - Cover all possible inputs
5. **Use factories, not fixtures** - Prefer Factory Bot for test data
6. **Use Subject for the Object Under Test** - Explicitly define the subject when helpful

### Modern RSpec Syntax

- **Use expect syntax** - Always use `expect().to` over `should`
- **Configure RSpec** - Only accept the new syntax to avoid mixing syntaxes
- **Use readable matchers** - Choose descriptive matchers like `be_empty` over `eq([])`
- **Shoulda matchers** - Prefer Shoulda matchers when available

### Data Management in Tests

- **Use let statements** - Prefer `let` over instance variables in `before` blocks
- **Avoid before(:all)** - Prevents data leakage between tests
- **Use let! for side effects** - When you need eager evaluation
- **Keep setup minimal** - Only create what's needed for the test

### Test Organization

- **One assertion per test** - Each test should verify one behavior
- **Use descriptive test names** - Names should explain what is being tested
- **Group related tests** - Use `describe` and `context` blocks effectively
- **Test behavior, not implementation** - Focus on what the code does, not how

### Model Testing

- **Test validations** - Verify all model validations work correctly
- **Test associations** - Ensure relationships are properly defined
- **Test scopes** - Verify named scopes return expected results
- **Test callbacks** - Ensure callbacks execute as expected

### Controller Testing

- **Test requests** - Prefer request specs over controller specs
- **Test responses** - Verify HTTP status codes and response formats
- **Test redirects** - Ensure proper redirection behavior
- **Test authorization** - Verify access control works correctly
- **Avoid over-testing** - Focus on controller-specific logic

### Integration Testing

- **Test user workflows** - Verify complete user journeys
- **Test edge cases** - Verify error handling and edge conditions
- **Keep tests independent** - Each test should be able to run in isolation

### Performance and Maintenance

- **Keep tests fast** - Use tools like `database_cleaner` efficiently
- **Use shared contexts** - Reduce duplication with shared setup
- **Mock external services** - Use stubs and mocks for external dependencies
- **Regular refactoring** - Keep test code clean and maintainable

### Tools and Configuration

- **Use fuubar** - Better test output formatting
- **Configure database_cleaner** - Proper test database cleanup
- **Use shoulda-matchers** - Simplify common Rails testing patterns
- **Set up CI/CD** - Automate test running and reporting
```

Ensure `e2e/CLAUDE.md` contains the following:

```md
# E2E Testing Guidelines

## Project Context

This is a TypeScript project using Playwright for end-to-end testing with real API integration.

## Test Strategy

- Focus on critical user journeys (authentication, core workflows)
- Validate complete user flows from start to finish using real backend APIs
- Test both happy paths and error scenarios
- Use real API calls for authentic integration testing
- Implement proper test data setup and cleanup

## Best Practices

- **Real APIs**: Use actual backend APIs instead of mocks for true integration testing
- **Selectors**: Use `data-testid` attributes or semantic selectors
- **Structure**: Group related tests in `test.describe` blocks
- **Setup/Cleanup**: Use `test.beforeEach`/`test.afterEach` for test data management
- **Data Isolation**: Create unique test data with timestamps to avoid conflicts
- **Cleanup**: Always clean up created test data to prevent database pollution
- **Naming**: Write descriptive test names that explain the behavior
- **Focus**: Keep test files to 3-5 focused tests maximum

## Test Data Management

- **Setup**: Create test data via API calls in `beforeEach` or individual tests
- **Unique Data**: Use timestamps or UUIDs to ensure unique test data
- **Cleanup**: Delete test data in `afterEach` hooks
- **Authentication**: Set up test users and valid JWT tokens for API calls

