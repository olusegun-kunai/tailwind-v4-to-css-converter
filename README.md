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

## 🛠️ Usage

```bash
# Convert a single file
npm run convert -- path/to/component.tsx

# Convert a directory
npm run convert -- path/to/components/directory
```

## 🎯 Current Status

### Completed
- ✅ Basic JSX parsing and class extraction
- ✅ UnoCSS integration for Tailwind processing
- ✅ CSS Module generation with semantic naming
- ✅ Diff visualization (console and HTML)
- ✅ Component analysis and statistics
- ✅ TypeScript type definitions
- ✅ Basic CLI interface

### In Progress
- 🔄 Enhanced class name generation
- 🔄 Improved component analysis
- 🔄 Better error handling and reporting

### To Do
- [ ] Support for nested components
- [ ] Handling of dynamic classes
- [ ] Support for Tailwind plugins
- [ ] Performance optimizations
- [ ] More comprehensive testing
- [ ] Documentation improvements
- [ ] Example components and use cases
- [ ] Integration with build tools

## 🧪 Testing

```bash
# Run all tests
npm test

# Run specific test file
npm test -- tests/jsx-parser.test.ts
```

## 📝 Notes

- The converter currently focuses on static Tailwind classes
- Dynamic classes (using template literals or variables) are not yet supported
- Nested component analysis is in progress
- Performance optimizations are planned for larger codebases

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

MIT 