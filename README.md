# Qwik Tailwind to CSS Modules Converter (TypeScript Implementation)

A TypeScript-based CLI tool that converts Qwik components from Tailwind classes to CSS modules using UnoCSS for reliable CSS generation.

> **Note**: This is the TypeScript implementation branch. For the original Go implementation, see the `main` branch.

## 🚀 Features

- **Node-by-Node Processing**: Parses each JSX element individually with semantic naming
- **Semantic CSS Class Names**: Component parts get meaningful names (`.trigger`, `.indicator`) while HTML elements get sequential names (`.node0`, `.node1`)
- **UnoCSS Integration**: Uses UnoCSS generator for reliable Tailwind-to-CSS conversion
- **CSS Modules Output**: Generates properly scoped CSS modules with clean formatting
- **Modifier Support**: Handles pseudo-selectors like `:hover`, `:focus`, `:disabled` 
- **Component Update**: Automatically updates JSX to use CSS module imports

## 🛠️ Installation & Setup

```bash
npm install
```

## 📖 Usage

### Command Line Interface

```bash
# Basic usage
npm run dev -- -i ./examples/checkbox.tsx -o ./output

# With verbose logging
npm run dev -- -i ./examples/checkbox.tsx -o ./output -v

# Using the demo command
npm run demo
```

### Available Scripts

```bash
npm run dev          # Run the converter CLI
npm run build        # Compile TypeScript to dist/
npm run demo         # Demo conversion with checkbox example
npm run test:all     # Run all tests
npm run test:basic   # Test UnoCSS generator only
npm run test:jsx     # Test JSX parser only  
npm run test:converter # Test full conversion pipeline
npm run clean        # Clean output directories
```

## 📁 Input/Output Example

### Input Component (`checkbox.tsx`)
```tsx
import { component$ } from "@builder.io/qwik";
import { Checkbox } from "@kunai-consulting/qwik";

export default component$(() => {
  return (
    <Checkbox.Root>
      <div class="flex items-center gap-2">
        <Checkbox.Trigger class="size-[25px] rounded-lg relative bg-gray-500">
          <Checkbox.Indicator class="data-[checked]:flex justify-center items-center">
            <LuCheck />
          </Checkbox.Indicator>
        </Checkbox.Trigger>
        <Checkbox.Label class="text-sm">
          This is a trusted device
        </Checkbox.Label>
      </div>
    </Checkbox.Root>
  );
});
```

### Generated CSS Modules (`checkbox.module.css`)
```css
/* Generated CSS Modules from Tailwind classes */

.node0 {
  display:flex;
  align-items:center;
  gap:0.5rem;
}

.trigger {
  position:relative;
  width:25px;
  height:25px;
  border-radius:0.5rem;
  --un-bg-opacity:1;
  background-color:rgb(107 114 128 / var(--un-bg-opacity));
}

.indicator {
  position:absolute;
  inset:0;
  display:flex;
  align-items:center;
  justify-content:center;
}

.label {
  font-size:0.875rem;
  line-height:1.25rem;
}
```

### Updated Component (`checkbox.tsx`)
```tsx
import { component$ } from "@builder.io/qwik";
import { Checkbox } from "@kunai-consulting/qwik";
import styles from './checkbox.module.css';

export default component$(() => {
  return (
    <Checkbox.Root>
      <div class={styles.node0}>
        <Checkbox.Trigger class={styles.trigger}>
          <Checkbox.Indicator class={styles.indicator}>
            <LuCheck />
          </Checkbox.Indicator>
        </Checkbox.Trigger>
        <Checkbox.Label class={styles.label}>
          This is a trusted device
        </Checkbox.Label>
      </div>
    </Checkbox.Root>
  );
});
```

## 🏗️ Architecture

### Core Components

- **`src/index.ts`**: CLI entry point with argument parsing
- **`src/lib/converter.ts`**: Main orchestrator class
- **`src/lib/unoGenerator.ts`**: UnoCSS wrapper for CSS generation
- **`src/lib/jsxParser.ts`**: JSX parsing and node extraction
- **`src/utils/cli.ts`**: Command-line argument handling
- **`src/types/index.ts`**: TypeScript type definitions

### Processing Pipeline

1. **JSX Parsing**: Extract JSX nodes with class attributes
2. **Semantic Naming**: Assign meaningful class names based on component structure
3. **CSS Generation**: Use UnoCSS to convert Tailwind classes to CSS
4. **CSS Modules Creation**: Format generated CSS as scoped modules
5. **Component Update**: Replace class attributes with CSS module references

## 🔧 Technical Details

### Semantic Naming Logic
- **Components**: `<Checkbox.Trigger>` → `.trigger`
- **HTML Elements**: `<div>`, `<span>` → `.node0`, `.node1`, etc.

### UnoCSS Integration
- Uses `@unocss/core` with `@unocss/preset-wind` (Tailwind preset)
- Reliable CSS generation via programmatic API
- Handles complex Tailwind features like arbitrary values and modifiers

### CSS Processing
- Extracts only the "default" layer (utility classes)
- Removes UnoCSS preflights and variable declarations
- Clean formatting with proper indentation

## 🧪 Testing

The project includes comprehensive tests:

- **Basic Tests**: UnoCSS generator functionality
- **JSX Parser Tests**: Node extraction and class parsing
- **Converter Tests**: End-to-end conversion pipeline

## 📦 Dependencies

- **Core**: `@unocss/core`, `@unocss/preset-wind`
- **Development**: `tsx`, `typescript`, `@types/node`
- **Runtime**: Node.js 18+ with ES modules support

## 🎯 Project Status

✅ **Completed Features**:
- Full conversion pipeline
- Semantic naming system
- UnoCSS integration
- CSS modules generation
- CLI interface
- Comprehensive testing

## 🌿 Branch Information

- **Current Branch**: `typescript-converter` - Clean TypeScript implementation
- **Main Branch**: Contains the original Go-based converter implementation
- **Migration**: This TypeScript implementation replaces the previous @apply-based approach with a more reliable UnoCSS-based solution

This converter successfully transitions from the previous @apply-based approach to a more reliable UnoCSS-based solution, providing clean CSS modules output with semantic class names. 