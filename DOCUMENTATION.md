## Tailwind → CSS Modules Converter: Project Guide

### 1) What this is
- **Purpose**: Convert components that use Tailwind classes into components that use plain CSS Modules, keeping the same visuals and structure.
- **Who it’s for**: Teams that want to ship plain CSS (or CSS Modules) while still authoring with Tailwind utilities during development.
- **Key value**: No Tailwind or UnoCSS runtime is required by the target project; the tool depends on UnoCSS internally to generate equivalent CSS.

### 2) Problem statement
- **Pain**: Many codebases want the speed of Tailwind utilities, but prefer compiled, framework-agnostic CSS at runtime (for portability, performance, or policy).
- **Gaps**:
  - Manual migration from Tailwind classes to CSS is tedious and error-prone.
  - Ensuring pixel-perfect parity after migration is hard.
  - Projects not using UnoCSS/Tailwind at runtime still want to author with utilities.

### 3) High-level solution
- Parse JSX to find elements/components with `class`/`className` attributes.
- Use UnoCSS (with Tailwind v4 preset) to generate the exact CSS properties for those classes.
- Emit a CSS Module with semantic selectors and update the component to reference those selectors instead of Tailwind classes.
- Provide a CLI to run conversion and an optional diff view for transparency.

### 4) Architecture and flow
- **JSX Parser (`src/lib/jsxParser.ts`)**
  - Detects JSX tags with `class`/`className` and extracts the class list.
  - Generates a semantic selector (e.g., `node0`, or for components like `Checkbox.Indicator` → `indicator`).

- **UnoCSS Generator (`src/lib/unoGenerator.ts`)**
  - Uses `@unocss/core` + `@unocss/preset-wind` to translate Tailwind classes into CSS.
  - Splits base and modifier rules (e.g., `:hover`, `:focus`).

- **Converter (`src/lib/converter.ts`)**
  - Orchestrates parsing, CSS generation, CSS Module creation, and component rewriting.
  - Writes `Component.module.css` and an updated `Component.tsx` that uses `styles.<selector>`.

- **Diff (`src/lib/diffGenerator.ts`)**
  - Optional console and HTML reports showing what changed (imports added, class replacements, counts).

### 5) CLI usage

```bash
# Convert a single file (outputs updated component + CSS module in the output dir)
npm run convert -- -i path/to/component.tsx -o path/to/output [-v] [-d] [--html-diff]

# Example
npm run convert -- -i ./qwik-convertor-test/src/routes/tailwind/index.tsx -o ./test-output/qwik -v -d
```

- **Flags**
  - `-i, --input`: Input component file (.tsx)
  - `-o, --output`: Output directory for the updated component and `.module.css`
  - `-v, --verbose`: Extra logs
  - `-d, --diff`: Show console diff summary
  - `--html-diff`: Also emit an HTML diff report next to the output component

### 6) Demo app (Qwik) and validation workflow
- Location: `qwik-convertor-test/`
- Routes:
  - `/tailwind` uses Tailwind classes directly.
  - `/vanilla` imports `vanilla.css` and should render the same visuals.

#### One-command update for the demo
```bash
npm run update:vanilla
```
This will:
- Convert `src/routes/tailwind/index.tsx` into a CSS Module in `test-output/qwik/`.
- Extract the generated CSS properties and write them into `src/routes/vanilla/vanilla.css` under `.vanilla-container`.

#### Run the demo server
```bash
npm --prefix ./qwik-convertor-test install
npm --prefix ./qwik-convertor-test run dev
```
Open:
- `http://localhost:5173/tailwind` (or the port shown)
- `http://localhost:5173/vanilla`

They should look identical. If not, see “Troubleshooting” below.

### 7) What “parity” means (validation checklist)
- Background color matches
- Text color matches
- Font size and line-height match
- Font weight matches
- No spacing/layout shifts or specificity issues
- Responsive behavior remains consistent (for classes that imply it)

### 8) Example (conceptual)
- Input JSX snippet:
  ```tsx
  <div class="bg-red-500 text-blue-300 text-2xl font-bold">Tailwind</div>
  ```
- Generated CSS Module (simplified):
  ```css
  .node0 {
    background-color: rgb(239 68 68);
    color: rgb(147 197 253);
    font-size: 1.5rem;
    line-height: 2rem;
    font-weight: 700;
  }
  ```
- Updated JSX:
  ```tsx
  <div className={styles.node0}>Tailwind</div>
  ```

### 9) Install and prerequisites
- Node.js 18+ recommended
- Repo installs `@unocss/core` and `@unocss/preset-wind` but the target project does not need UnoCSS or Tailwind at runtime

### 10) Limitations and assumptions
- Focuses on static class strings (dynamic template strings/variables are not yet supported).
- CSS specificity should match Tailwind defaults for common cases; verify with the parity checklist.
- Generated classnames are semantic but simple; some projects may want to adjust naming rules.

### 11) Troubleshooting
- If `/vanilla` and `/tailwind` differ:
  - Ensure `vanilla.css` contains exactly the generated properties.
  - If you see `--un-*-opacity` variables, you can keep them or replace with concrete rgb/hex; both should render equivalently.
  - Check global styles (`src/global.css`) for unexpected overrides.
  - Re-run the update: `npm run update:vanilla` and refresh.

### 12) FAQ
- **Do I need UnoCSS in my project?** No. The tool uses UnoCSS internally for conversion only.
- **Can I convert entire folders?** The CLI currently targets files; you can script batch conversions.
- **Does it support plugins/responsive/variants?** Common modifiers (e.g., hover/focus) are parsed. Plugin coverage varies; verify via demo.

### 13) Files to know about
- `src/index.ts`: CLI entry
- `src/lib/*`: Parser, generator, converter, and diff logic
- `scripts/update-vanilla-css.ts`: Helper to update the demo’s `vanilla.css`
- `qwik-convertor-test/*`: Demo app to visually validate parity


