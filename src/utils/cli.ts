import { CliArgs } from '../types/index.js';

export function parseArgs(): CliArgs {
  const args = process.argv.slice(2);
  
  const result: CliArgs = {
    input: '',
    output: '',
    verbose: false,
    diff: false,
    htmlDiff: false
  };

  for (let i = 0; i < args.length; i++) {
    const arg = args[i];
    
    switch (arg) {
      case '-i':
      case '--input':
        result.input = args[++i];
        break;
      case '-o':
      case '--output':
        result.output = args[++i];
        break;
      case '-v':
      case '--verbose':
        result.verbose = true;
        break;
      case '-d':
      case '--diff':
        result.diff = true;
        break;
      case '--html-diff':
        result.htmlDiff = true;
        result.diff = true; // HTML diff implies diff
        break;
      case '-h':
      case '--help':
        printHelp();
        process.exit(0);
        break;
      default:
        if (arg.startsWith('-')) {
          console.error(`Unknown option: ${arg}`);
          printHelp();
          process.exit(1);
        }
    }
  }

  if (!result.input || !result.output) {
    console.error('❌ Missing required arguments');
    printHelp();
    process.exit(1);
  }

  return result;
}

function printHelp() {
  console.log(`
🔄 Qwik Tailwind to CSS Modules Converter

Usage:
  tsx src/index.ts -i <input-file> -o <output-dir> [options]

Options:
  -i, --input     Input Qwik component file (.tsx)
  -o, --output    Output directory for generated files
  -v, --verbose   Enable verbose logging
  -d, --diff      Show conversion changes in console
  --html-diff     Generate HTML diff report
  -h, --help      Show this help message

Examples:
  tsx src/index.ts -i ./examples/checkbox.tsx -o ./output
  tsx src/index.ts -i ./components/button.tsx -o ./dist -v
  tsx src/index.ts -i ./components/button.tsx -o ./dist --diff
  tsx src/index.ts -i ./components/button.tsx -o ./dist --html-diff
`);
} 