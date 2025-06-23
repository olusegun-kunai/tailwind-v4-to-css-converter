#!/usr/bin/env node

import { QwikTailwindConverter } from './lib/converter.js';
import { parseArgs } from './utils/cli.js';
import fs from 'fs/promises';

async function main() {
  try {
    const args = parseArgs();
    
    const converter = new QwikTailwindConverter();
    
    console.log('🚀 Starting Qwik Tailwind to CSS Modules conversion...\n');
    
    const result = await converter.convertFile(args.input, args.output, args.diff);
    
    console.log('✅ Conversion completed successfully!');
    console.log(`📁 CSS Modules: ${result.cssModulesPath}`);
    console.log(`📄 Updated Component: ${result.componentPath}`);
    
    // Show diff if requested
    if (args.diff && result.changeReport) {
      const { DiffGenerator } = await import('./lib/diffGenerator.js');
      const diffGenerator = new DiffGenerator();
      
      // Console diff
      console.log(diffGenerator.generateConsoleDiff(result.changeReport));
      
      // HTML diff if requested  
      if (args.htmlDiff) {
        const htmlDiff = diffGenerator.generateHTMLDiff(result.changeReport);
        const htmlPath = result.componentPath.replace('.tsx', '.diff.html');
        await fs.writeFile(htmlPath, htmlDiff);
        console.log(`🌐 HTML Diff Report: ${htmlPath}`);
      }
    }
    
  } catch (error) {
    console.error('❌ Conversion failed:', error);
    process.exit(1);
  }
}

// Run if this file is executed directly
if (import.meta.url === `file://${process.argv[1]}`) {
  main();
} 