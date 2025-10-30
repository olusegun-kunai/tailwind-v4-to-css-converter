## Support Tailwind v4 Theme Configuration (CSS) — 5-Day Work Plan

### Objective
Enhance the converter CLI so generated vanilla CSS respects Tailwind v4 theme configuration defined in CSS files (e.g., `theme.css`). The CLI should accept a theme path, parse it, and apply the tokens during conversion. Fallback to default theme when no config is provided.

### Acceptance Criteria
- CLI accepts `--theme ./path/to/theme.css` (absolute/relative).
- Converter parses the Tailwind v4 CSS theme file and applies values (colors, spacing, fonts, etc.).
- Generated vanilla CSS reflects custom theme values.
- Default behavior unchanged when `--theme` is omitted.
- Clear error handling for invalid/missing theme files.
- Documentation updated with theme usage examples.

### References
- UnoCSS wind v4 tests: [preset-wind4.test.ts](https://github.com/unocss/unocss/blob/main/test/preset-wind4.test.ts)
- Tailwind v4 theme concepts: [Tailwind theme docs](https://tailwindcss.com/docs/theme)

---

### Day 1 — Research, Design, and CLI Spec
- Review Tailwind v4 theme-in-CSS conventions (e.g., `@theme` blocks, CSS custom properties, layering).
- Investigate parsing approach for CSS theme files (PostCSS vs css-tree). Choose PostCSS for robust parsing.
- Map CSS theme tokens to UnoCSS `theme` structure expected by `createGenerator({ presets: [presetWind()], theme })`:
  - Colors: map CSS variables (e.g., `--color-red-500`) and `@theme` tokens into `theme.colors.red.500`.
  - Spacing: map to `theme.spacing`.
  - Font sizes/weights: map to `theme.fontSize` / `theme.fontWeight`.
- Define CLI flag: `--theme <path>` (optional). Resolution rules: resolve from CWD; support `.css` only for this ticket.
- Deliverable: design notes, CLI arg updates scaffolded (no runtime behavior yet).

### Day 2 — Theme CSS Loader and Normalizer
- Implement `src/utils/themeCssLoader.ts`:
  - Read CSS file, parse with PostCSS.
  - Extract `@theme` blocks and relevant custom properties.
  - Normalize to an UnoCSS-compatible `theme` object (deep object with colors, spacing, fontSize, fontWeight).
  - Validate schema, produce friendly errors with file/line context.
- Unit tests for loader: valid file, missing file, malformed CSS, minimal/no-op theme.
- Deliverable: loader returns a sanitized `theme` object or throws clear errors.

### Day 3 — Integration with UnoCSS and CLI Wiring
- Update `UnoGenerator` to accept an optional `theme` in constructor and pass into `createGenerator({ presets: [presetWind()], theme })`.
- Update `src/index.ts` to parse `--theme`, call loader, and pass theme to `QwikTailwindConverter` -> `UnoGenerator`.
- Add verbose logs summarizing applied overrides when `-v` is set.
- E2E smoke test: override `colors.red.500` and `fontSize.2xl` in a sample `theme.css`; confirm generated CSS uses overrides.
- Deliverable: end-to-end conversion respects theme overrides.

### Day 4 — Examples, Docs, and Edge Cases
- Add `examples/themes/theme.css` demonstrating overrides (colors, spacing, font size).
- Update `README.md`, `DOCUMENTATION.md`, and `EXAMPLES.md` with `--theme` usage and example commands:
  - `npm run convert -- -i ./path/comp.tsx -o ./out --theme ./examples/themes/theme.css`
- Validate with Qwik demo:
  - Run converter with `--theme`, `npm run update:vanilla` (optionally extended to accept `--theme`), compare `/tailwind` vs `/vanilla` — visuals should reflect custom theme.
- Edge cases: ensure deep merge (theme overrides only replace specified keys), warn on unknown keys, handle multiple `@theme` blocks.
- Deliverable: documented examples and stable behavior across common cases.

### Day 5 — Testing, Polish, and Release Prep
- Expand tests:
  - Nested color maps; spacing scale overrides; missing keys fallback.
  - Invalid CSS produces actionable errors.
- DX polish: improve error messages; add `--html-diff` validation with themed outputs.
- Final doc polish and changelog entry; bump minor version.
- Deliverable: ready-to-release feature with tests and docs.

---

### Validation Plan
- Unit: `themeCssLoader` parsing and normalization.
- Integration: `UnoGenerator` receives and uses injected theme.
- E2E: Convert known inputs and assert CSS contains themed values.
- Visual: Qwik demo `/tailwind` vs `/vanilla` matches and reflects theme overrides.

### Risks & Mitigations
- Tailwind v4 theme-in-CSS nuances → Start with well-scoped tokens (colors, spacing, fontSize); add more iteratively.
- Parser complexity → Use PostCSS; add robust tests and fixtures.
- Preset-wind expectations → Pin compatible versions and cross-check against wind v4 tests.


