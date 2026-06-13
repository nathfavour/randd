# Agent Skills
Custom skills and tools for AI agents, following the **Agent Skills Open Standard** and the **`.agents/skills/`** filesystem convention.

## Structure
Each skill directory follows this standard layout:
- `SKILL.md`: Metadata (YAML frontmatter) and core instructions.
- `scripts/`: Executable code (Python, Bash, Node, etc.) used by the skill.
- `references/`: Supporting documentation or API specifications.
- `assets/`: Templates, schemas, or static assets.
- `tests/`: Test cases and validation logic.

## Conventions
- **Discovery:** Modern agentic CLIs automatically discover skills placed in `.agents/skills/`.
- **Progressive Disclosure:** Keep `SKILL.md` concise. Offload bulky data to `references/`.
- **Atomic Activations:** Use trigger-optimized descriptions in the YAML frontmatter.
- **Deterministic Logic:** Prefer scripts for complex parsing or data manipulation.
