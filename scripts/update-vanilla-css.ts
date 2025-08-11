#!/usr/bin/env tsx

import fs from 'fs/promises';
import path from 'path';
import { fileURLToPath } from 'url';
import { QwikTailwindConverter } from '../src/lib/converter.js';

async function main() {
  const thisFilePath = fileURLToPath(import.meta.url);
  const thisDir = path.dirname(thisFilePath);
  const repoRoot = path.resolve(thisDir, '..');
  const tailwindRoute = path.join(repoRoot, 'qwik-convertor-test/src/routes/tailwind/index.tsx');
  const vanillaCssPath = path.join(repoRoot, 'qwik-convertor-test/src/routes/vanilla/vanilla.css');
  const tempOutDir = path.join(repoRoot, 'test-output/qwik');

  const converter = new QwikTailwindConverter();

  console.log('🔄 Converting Tailwind route to CSS...');
  const result = await converter.convertFile(tailwindRoute, tempOutDir, false);

  const css = await fs.readFile(result.cssModulesPath, 'utf-8');
  const match = css.match(/\.([^{\s]+)\s*\{([\s\S]*?)\}\s*$/m);

  if (!match) {
    throw new Error('Failed to extract CSS block from generated module.');
  }

  const propertiesBlock = match[2]
    .split('\n')
    .map((l) => l.trim())
    .filter(Boolean)
    .join('\n  ');

  const vanillaTemplate = `/* Generated CSS from Tailwind classes */\n/* Converted using Tailwind to CSS Modules Converter */\n\n.vanilla-container {\n  ${propertiesBlock}\n}`;

  await fs.writeFile(vanillaCssPath, vanillaTemplate, 'utf-8');

  console.log(`✅ Updated vanilla CSS: ${path.relative(repoRoot, vanillaCssPath)}`);
}

main().catch((err) => {
  console.error('❌ Failed to update vanilla CSS:', err);
  process.exit(1);
});


