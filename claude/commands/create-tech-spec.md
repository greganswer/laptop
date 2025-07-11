# Generate a Technical Specification Document

## Goal

To guide an AI assistant in creating a detailed Technical Specification Document in Markdown format, based on a brief user description. The tech spec should focus on implementation details, be clear and actionable, and suitable for a junior developer to understand how to build the feature.

## Process

1. **Receive Initial Description:** The user provides a brief description of what they want to implement (e.g., "Add Cypress E2E test framework")
2. **Ask Clarifying Questions:** Before writing the tech spec, the AI *must* ask clarifying questions to understand the technical requirements and constraints. The goal is to understand the "how" of implementation, architectural decisions, and integration patterns.
3. **Generate Tech Spec:** Based on the initial description and the user's answers to the clarifying questions, generate a technical specification using the structure outlined below.
4. **Save Tech Spec:** Save the generated document as `tech-spec.md` in the current directory or as specified by the user.

## Clarifying Questions (Examples)

The AI should adapt its questions based on the feature description, but here are key areas to explore:

* **Current State:** "What is the current architecture/tech stack of the system?" or "Are there existing patterns or frameworks we should follow?"
* **Integration Points:** "What existing systems/modules will this need to integrate with?"
* **Data Flow:** "How will data flow through this feature? What are the inputs and outputs?"
* **Technology Choices:** "Are there specific technologies/libraries you prefer or want to avoid?"
* **Performance Requirements:** "What are the expected performance characteristics? (e.g., response time, throughput, concurrent users)"
* **Security Requirements:** "What security considerations should we address? (e.g., authentication, authorization, data encryption)"
* **Scalability:** "What are the expected scale requirements? Will this need to handle growth?"
* **Testing Approach:** "What types of tests should we include? (unit, integration, E2E)"
* **Deployment:** "How will this be deployed? Are there specific deployment constraints?"
* **Monitoring:** "What metrics should we track? How will we monitor this in production?"

## Tech Spec Structure

The generated tech spec should include the following sections:

1. **Overview:** Brief description of what's being implemented and why (1-2 paragraphs)
2. **Architecture Design:** High-level architecture decisions and system design
   * Component diagram or description (use mermaid.js)
   * Data flow diagram or description (use mermaid.js)
   * Key architectural patterns used
3. **Database Schema:** Any new tables, columns, or schema changes (include SQL or schema definitions)
4. **API Specifications:**
   * Endpoint definitions (REST/GraphQL/gRPC)
   * Request/Response formats
   * Error handling
5. **Integration Patterns:** How this integrates with existing systems
   * External services
   * Internal modules
   * Event/messaging patterns
6. **Security Considerations:**
   * Authentication/Authorization approach
   * Data validation
   * Security best practices
7. **Performance Requirements:**
   * Expected load
   * Response time targets
   * Caching strategy
8. **Testing Strategy:**
   * Unit test approach
   * Integration test approach
   * E2E test scenarios
9. **Deployment & Infrastructure:**
   * Deployment process
   * Environment requirements
   * Configuration management
10. **Monitoring & Observability:**
    * Key metrics to track
    * Logging approach
    * Alert conditions
11. **Dependencies & Risks:**
    * External dependencies
    * Technical risks and mitigation strategies
12. **Future Considerations:** What might need to change as the system grows?

Also add a table of contents at the top of the document.

## Code Examples

Include code examples only when necessary to clarify complex concepts or demonstrate critical integration patterns. Code should be:

* Minimal and focused
* Well-commented
* Language-agnostic pseudocode unless a specific language is required

## Target Audience

Assume the primary reader is a **junior developer**. Therefore:

* Use clear, simple language
* Avoid unnecessary jargon
* Explain technical decisions and trade-offs
* Provide enough detail for implementation without being overwhelming

## Output

* **Format:** Markdown (`.md`)
* **Location:** Current directory or as specified by the user
* **Filename:** `[feature-name]-tech-spec.md` exampe: `cypress-e2e-test-framework-tech-spec.md`

## Final Instructions

1. Do NOT start implementing the technical specification
2. Make sure to ask the user clarifying questions based on their initial description
3. Take the user's answers and create a comprehensive technical specification
4. Focus on "how" to build it, not "what" to build (that's the PRD's job)
5. Keep language clear and accessible for junior developers
6. Ensure all mandatory sections are included

