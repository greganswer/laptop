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
├── frontend
│   └── CLAUDE.md
├── PRPs
│   └── archive
│       └── MVP
│           └── requirements.md
├── CLAUDE.md
├── Profile.dev
└── start.sh
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

Ensure `start.sh` includes the following:

```sh
lsof -ti:3001 -ti:3000 | xargs kill -9 && foreman start -f Procfile.dev
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
