## Theme Configuration Support — 5-Day Implementation Plan

### Objective
Enable the converter CLI to respect a user-provided theme configuration (Tailwind-like), so generated CSS reflects custom tokens (colors, spacing, font sizes, etc.). The tool remains generic and runnable in any repo.

### References
- UnoCSS wind v4 tests: [preset-wind4.test.ts](https://github.com/unocss/unocss/blob/main/test/preset-wind4.test.ts)
- Tailwind CSS theme docs: [Tailwind theme](https://tailwindcss.com/docs/theme)

---

### Deliverables
- CLI option `--theme <path>` that accepts JSON/JS/TS exporting a theme object
- Loader + validation for the theme object
- UnoCSS generator updated to merge the provided theme
- Examples demonstrating overrides (e.g., `colors.red.500`, `fontSize.2xl`)
- Tests (unit + e2e) validating generated CSS reflects overrides
- Docs: README/DOCUMENTATION updates and an example theme config

### Acceptance Criteria
- Passing `--theme` modifies generated CSS according to overrides
- Works with absolute and relative paths; clear errors for invalid configs
- Backwards compatible when `--theme` is omitted
- Demo parity: overrides visible in `/tailwind` vs `/vanilla` after update script

---

### Day 1 — Design & Scaffolding
- Research UnoCSS theme merging behavior with `preset-wind`
- Decide supported formats: `.json`, `.js`, `.ts` (ESM)
- Define minimal schema (subset aligned with Tailwind):
  - `theme.colors`, `theme.fontSize`, `theme.spacing` (stretch: `borderRadius`, `fontWeight`)
- Design CLI flag: `--theme <path>`, documented behavior and resolution rules (resolve from CWD)
- Scaffolding:
  - Create `src/utils/themeLoader.ts` with interface + stubs
  - Add example theme files under `examples/themes/`
- Outcome: Architectural notes + stubbed loader and CLI arg parsing (no runtime changes yet)

### Day 2 — Loader & Validation
- Implement `themeLoader.ts`:
  - Resolve path, load JSON/ESM module (use `import()`), support `.json`, `.js`, `.ts`
  - Normalize structure to UnoCSS-compatible `theme` object
  - Validate allowed keys; informative errors with exact path to bad fields
- Add unit tests for loader (valid/invalid/empty/missing file)
- Outcome: Reliable theme loader returning a sanitized `theme` object

### Day 3 — Generator Integration & CLI Plumbing
- Update `UnoGenerator`:
  - Accept optional `theme` in constructor
  - Pass `theme` into `createGenerator({ presets: [presetWind()], theme })`
- Update CLI entry (`src/index.ts`):
  - Parse `--theme`, call loader, instantiate `QwikTailwindConverter` with theme
- E2E test using a simple theme override (e.g., `colors.red.500` → distinct color)
- Outcome: End-to-end conversion reflects theme overrides in generated CSS

### Day 4 — Examples, Docs, and Edge Cases
- Add example theme files:
  - `examples/themes/custom.json` (colors + fontSize)
  - `examples/themes/custom.ts` (function/JS object export)
- Update `EXAMPLES.md` and `DOCUMENTATION.md` with `--theme` usage
- Demo validation:
  - Run converter against Qwik `/tailwind` with `--theme`, update `/vanilla` CSS, verify visual changes match theme
- Edge cases:
  - Deep merging vs replace semantics: implement deep merge with sensible defaults
  - Clear warnings for unknown keys; no crashes
- Outcome: Usable docs and examples with stable behavior

### Day 5 — Polish, Tests, and Release Prep
- Add additional tests:
  - Multiple keys overridden, nested color maps
  - Missing keys fall back to defaults
- Improve DX:
  - Helpful error messages and `--verbose` logging path + keys applied
  - Print a short summary of active overrides when `-v`
- Finalize docs (README quickstart + theme section)
- Cut a minor version bump and changelog entry

---

### Validation Plan
- Unit: theme loader schema and path resolution
- Integration: UnoGenerator uses the injected theme
- E2E: Convert known inputs and assert CSS contains override values
- Visual: Qwik demo `/tailwind` vs `/vanilla` with custom theme shows expected colors/sizes

### Risks & Mitigations
- Incorrect theme merge semantics → Add tests mirroring Tailwind expectations for covered keys
- ESM import issues for `.ts` configs → Use `tsx` in dev context; document JSON/JS as primary
- Unexpected preset-wind changes → Pin compatible versions; add lockfile and tests

### Post-Plan Stretch (Optional)
- Support Tailwind-like `extend` semantics
- Config discovery (look for `theme.*` in project root if `--theme` omitted)
- Watch mode for theme changes


