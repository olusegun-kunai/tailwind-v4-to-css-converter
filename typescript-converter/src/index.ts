#!/usr/bin/env node

import { QwikTailwindConverter } from './lib/converter.js';
import { parseArgs } from './utils/cli.js';

async function main() {
  try {
    const args = parseArgs();
    
    const converter = new QwikTailwindConverter();
    
    console.log('🚀 Starting Qwik Tailwind to CSS Modules conversion...\n');
    
    const result = await converter.convertFile(args.input, args.output);
    
    console.log('✅ Conversion completed successfully!');
    console.log(`📁 CSS Modules: ${result.cssModulesPath}`);
    console.log(`📄 Updated Component: ${result.componentPath}`);
    
  } catch (error) {
    console.error('❌ Conversion failed:', error);
    process.exit(1);
  }
}

// Run if this file is executed directly
if (import.meta.url === `file://${process.argv[1]}`) {
  main();
} 