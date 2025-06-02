# Tailwind v4+ to CSS Converter

A Go-based tool that converts Tailwind CSS v4+ classes in HTML/JSX/TSX files to semantic CSS modules with improved maintainability and performance.

## Features

- ✅ Parse HTML, JSX, TSX, and Vue component files
- ✅ Extract Tailwind classes and convert to vanilla CSS
- ✅ Generate semantic CSS modules (.module.css files)
- ✅ Update component files with CSS module imports
- ✅ Deduplicate CSS properties automatically
- ✅ Support for responsive classes and pseudo-states
- ✅ Comprehensive Tailwind class mapping
- ✅ CLI interface for batch processing

## Installation

```bash
git clone https://github.com/your-org/tailwind-v4-to-css-converter.git
cd tailwind-v4-to-css-converter
go mod tidy
go build -o tailwind-converter
```

## Usage

### Basic Usage

```bash
./tailwind-converter --input ./src/components --output ./dist
```

### With Verbose Output

```bash
./tailwind-converter --input ./src/components --output ./dist --verbose
```

### Convert Single Directory

```bash
./tailwind-converter --input ./src --output ./dist
```

## Example Conversion

**Input (otp-input.tsx):**
```jsx
export default component$(() => {
  return (
    <div class="flex flex-col items-center gap-4">
      <div class="flex flex-col items-center text-center">
        <h3 class="text-sm font-semibold">Two-step verification</h3>
      </div>
    </div>
  );
});
```

**Output (otp-input.tsx):**
```jsx
import styles from './otp-input.module.css';

export default component$(() => {
  return (
    <div class={styles.div_layout_1}>
      <div class={styles.div_container_2}>
        <h3 class={styles.h3_text_3}>Two-step verification</h3>
      </div>
    </div>
  );
});
```

**Generated CSS (otp-input.module.css):**
```css
/* Generated CSS Module */
/* Converted from Tailwind CSS classes */

.div_layout_1 {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 1rem;
}

.div_container_2 {
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.h3_text_3 {
  font-size: 0.875rem;
  font-weight: 600;
}
```

## Testing

To verify the solution works:

1. **Build the tool:**
   ```bash
   go mod tidy
   go build -o tailwind-converter
   ```

2. **Create test files:**
   ```bash
   mkdir -p test/src test/output
   # Add your HTML/JSX/TSX files to test/src/
   ```

3. **Run conversion:**
   ```bash
   ./tailwind-converter --input ./test/src --output ./test/output --verbose
   ```

4. **Verify output:**
   - Check that `.module.css` files are generated
   - Verify component files have CSS module imports
   - Ensure classes are properly converted and deduplicated

## Supported Tailwind Classes

### Core Utilities
- **Display**: `flex`, `grid`, `block`, `inline`, `hidden`
- **Flexbox**: `flex-col`, `items-center`, `justify-between`
- **Spacing**: `p-4`, `m-2`, `gap-4`, `px-6`, `py-2`
- **Sizing**: `w-full`, `h-64`, `w-4`, `h-4`
- **Typography**: `text-sm`, `font-bold`, `font-semibold`
- **Colors**: `bg-blue-500`, `text-red-600`, `border-green-200`
- **Layout**: `container`, `mx-auto`
- **Visual**: `bg-white`, `border`, `rounded-md`, `shadow-lg`

### Advanced Features
- **Responsive**: `md:grid-cols-2`, `lg:grid-cols-3`
- **Pseudo-states**: `hover:bg-blue-700`, `focus:outline-none`
- **Grid**: `grid-cols-1`, `grid-cols-4`
- **Border**: `border-b`, `border-blue-200`, `rounded-lg`
- **Shadow**: `shadow`, `shadow-md`, `shadow-lg`

## Architecture

```
internal/
├── parser/
│   ├── html.go          # Parse HTML/JSX/TSX files
│   └── classes.go       # Extract Tailwind classes
├── generator/
│   ├── css.go           # Generate CSS modules
│   └── html.go          # Update HTML with semantic classes
converter/
├── converter.go         # Main conversion logic with deduplication
├── mappings.go          # Comprehensive Tailwind to CSS mappings
└── modern.go            # Modern CSS features support
```

## Improvements Made

- **Fixed duplicate class references** in generated components
- **Enhanced Tailwind mappings** to support 200+ classes
- **Improved semantic naming** based on element type and class categories
- **Better CSS property deduplication** to prevent conflicts
- **Comprehensive color support** for blue, red, green, gray, purple palettes
- **Responsive and pseudo-state handling** for modern CSS features

## Current Limitations

- Dynamic class names not supported (e.g., template literals with variables)
- Some complex Tailwind plugins may need manual conversion
- Single file processing requires directory input
- AI integration for unknown classes not yet implemented

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Submit a pull request

## License

MIT License - see LICENSE file for details.