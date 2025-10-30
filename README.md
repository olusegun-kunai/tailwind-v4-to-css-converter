# Tailwind to CSS Module Converter

A TypeScript-based tool that converts Tailwind CSS classes to CSS Modules, with semantic class naming and component-based organization.

## 🚀 Features Implemented

- **JSX Parser**: Successfully parses React components and extracts Tailwind classes
- **UnoCSS Integration**: Uses UnoCSS for accurate Tailwind class processing
- **CSS Module Generation**: Converts Tailwind classes to CSS Modules with semantic naming
- **Diff Visualization**: Shows changes between original and converted code in both console and HTML formats
- **Component Analysis**: Provides detailed analysis of component structure and class usage
- **Type Safety**: Full TypeScript support with proper type definitions

## 📦 Project Structure

```
src/
├── lib/
│   ├── jsxParser.ts    # JSX parsing and analysis
│   ├── unoGenerator.ts # UnoCSS integration
│   └── converter.ts    # Main conversion logic
├── utils/
│   └── cli.ts         # CLI utilities
└── types/
    └── index.ts       # TypeScript type definitions
```

## 🛠️ CLI Usage

```bash
# Convert a single file (outputs updated component + CSS module in the output dir)
npm run convert -- -i path/to/component.tsx -o path/to/output [-v] [-d] [--html-diff]

# Example
npm run convert -- -i ./qwik-convertor-test/src/routes/tailwind/index.tsx -o ./test-output/qwik -v -d
```

## 🔗 Qwik Demo (Tailwind → Vanilla)

This repo includes a Qwik app to demonstrate converting Tailwind classes to vanilla CSS.

### Quickstart

```bash
# 1) Install repo deps
npm install

# 2) Generate CSS from the Tailwind route and update the vanilla route
npm run update:vanilla

# 3) Install and start the Qwik demo app
npm --prefix ./qwik-convertor-test install
npm --prefix ./qwik-convertor-test run dev
```

Open `http://localhost:5173/tailwind` and `http://localhost:5173/vanilla` and ensure they look identical.

### Notes
- The converter uses UnoCSS internally; your target project can keep using Tailwind.
- If you prefer a one-liner: `npm run demo:all` will update CSS and start the dev server.

## ✅ What to expect
- Converter generates a `.module.css` file and updates the component to use it.
- The included script updates the demo’s `vanilla.css` to match the Tailwind route.
- The converter focuses on static Tailwind classes.

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

MIT 