## Conversion Examples and Recipes

This guide shows concrete conversion runs and what to expect afterwards.

What happens after a conversion:
- Writes an updated component that imports a generated CSS module.
- Writes a `.module.css` file with the equivalent CSS rules.
- You can either keep CSS Modules, or copy the CSS properties into your own stylesheet.

General command:
```bash
npm run convert -- -i <path/to/component.tsx> -o <path/to/output> [-v] [-d] [--html-diff]
```

---

### 1) Minimal Qwik route (static classes)
- Input: `qwik-convertor-test/src/routes/tailwind/index.tsx`
- Classes: `bg-red-500 text-blue-300 text-2xl font-bold`
- Run:
```bash
npm run convert -- -i ./qwik-convertor-test/src/routes/tailwind/index.tsx -o ./test-output/qwik -v -d
```
- Output:
  - `test-output/qwik/index.module.css`
  - `test-output/qwik/index.tsx`

Tip: use `npm run update:vanilla` to copy the generated properties into `qwik-convertor-test/src/routes/vanilla/vanilla.css`.

---

### 2) React Button (simple static classes)
- Input: `examples/button-simple.tsx`
- Run:
```bash
npm run convert -- -i ./examples/button-simple.tsx -o ./output/button-simple -v -d
```
- Output:
  - `output/button-simple/button-simple.module.css`
  - `output/button-simple/button-simple.tsx`

What to check: hover/focus rules and base styles mapped into CSS.

---

### 3) React Button with variants (composed classes)
- Input: `examples/button.tsx`
- Note: The converter focuses on static class strings. Template composition is flattened at runtime. For demonstration, run conversion to inspect the generated module and updated component:
```bash
npm run convert -- -i ./examples/button.tsx -o ./output/button -v -d
```

What to check: CSS module includes the union of utilities used by the rendered variant.

---

### 4) Qwik Checkbox (nested components, modifiers)
- Input: `examples/checkbox.tsx`
- Run:
```bash
npm run convert -- -i ./examples/checkbox.tsx -o ./output/checkbox -v -d
```
- Output:
  - `output/checkbox/checkbox.module.css`
  - `output/checkbox/checkbox.tsx`
  - Add `--html-diff` to also create an HTML report

What to check: semantic selectors for parts (e.g., `trigger`, `indicator`), and generated rules for `focus-visible`, `data-[checked]`, etc.

---

### 5) HTML diff report
Append `--html-diff`:
```bash
npm run convert -- -i ./examples/checkbox.tsx -o ./output/checkbox -v -d --html-diff
```
This will place an HTML report next to the updated component file.

---

### Validation tips
- Use the Qwik demo: open `/tailwind` and `/vanilla` and confirm visuals match.
- Verify hover/focus states.
- If you see `--un-*-opacity` variables, you may replace them with concrete rgb/hex if desired; visuals should match either way.


