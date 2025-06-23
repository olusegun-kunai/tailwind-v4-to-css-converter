import { QwikTailwindConverter } from '../src/lib/converter.js';
import { DiffGenerator } from '../src/lib/diffGenerator.js';
import fs from 'fs/promises';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

describe('Diff Generation', () => {
  let converter: QwikTailwindConverter;
  let diffGenerator: DiffGenerator;
  
  beforeEach(() => {
    converter = new QwikTailwindConverter();
    diffGenerator = new DiffGenerator();
  });

  test('should generate change report', async () => {
    const inputPath = path.join(__dirname, '../examples/checkbox.tsx');
    const outputDir = path.join(__dirname, '../temp-diff-test');
    
    // Clean up output directory
    try {
      await fs.rm(outputDir, { recursive: true });
    } catch (error) {
      // Directory might not exist
    }
    
    // Run conversion with diff generation
    const result = await converter.convertFile(inputPath, outputDir, true);
    
    // Verify change report was generated
    expect(result.changeReport).toBeDefined();
    expect(result.changeReport!.changes).toHaveLength(3); // import + 2 class replacements
    expect(result.changeReport!.summary.totalNodes).toBe(2);
    expect(result.changeReport!.summary.classesConverted).toBe(2);
    expect(result.changeReport!.summary.importsAdded).toBe(1);
    
    // Verify changes include import and class replacements
    const changes = result.changeReport!.changes;
    const importChange = changes.find(c => c.type === 'import-added');
    const classChanges = changes.filter(c => c.type === 'class-replacement');
    
    expect(importChange).toBeDefined();
    expect(importChange!.modified).toContain('import styles from');
    expect(classChanges).toHaveLength(2);
    
    // Clean up
    await fs.rm(outputDir, { recursive: true });
  });

  test('should generate console diff output', async () => {
    const inputPath = path.join(__dirname, '../examples/checkbox.tsx');
    const outputDir = path.join(__dirname, '../temp-diff-test');
    
    try {
      await fs.rm(outputDir, { recursive: true });
    } catch (error) {
      // Directory might not exist
    }
    
    const result = await converter.convertFile(inputPath, outputDir, true);
    
    const consoleDiff = diffGenerator.generateConsoleDiff(result.changeReport!);
    
    expect(consoleDiff).toContain('🔄 Conversion Summary:');
    expect(consoleDiff).toContain('Components: 2');
    expect(consoleDiff).toContain('Classes converted: 2');
    expect(consoleDiff).toContain('📝 Changes made:');
    expect(consoleDiff).toContain('import styles from');
    expect(consoleDiff).toContain('styles.trigger');
    expect(consoleDiff).toContain('styles.indicator');
    
    // Clean up
    await fs.rm(outputDir, { recursive: true });
  });

  test('should generate HTML diff', async () => {
    const inputPath = path.join(__dirname, '../examples/checkbox.tsx');
    const outputDir = path.join(__dirname, '../temp-diff-test');
    
    try {
      await fs.rm(outputDir, { recursive: true });
    } catch (error) {
      // Directory might not exist
    }
    
    const result = await converter.convertFile(inputPath, outputDir, true);
    
    const htmlDiff = diffGenerator.generateHTMLDiff(result.changeReport!);
    
    expect(htmlDiff).toContain('<!DOCTYPE html>');
    expect(htmlDiff).toContain('Tailwind → CSS Modules Conversion');
    expect(htmlDiff).toContain('Conversion Summary');
    expect(htmlDiff).toContain('Components Processed');
    expect(htmlDiff).toContain('Classes Converted');
    expect(htmlDiff).toContain('CSS Rules Generated');
    expect(htmlDiff).toContain('import styles from');
    expect(htmlDiff).toContain('class-replacement');
    expect(htmlDiff).toContain('import-added');
    
    // Clean up
    await fs.rm(outputDir, { recursive: true });
  });

  test('should write HTML diff file', async () => {
    const inputPath = path.join(__dirname, '../examples/checkbox.tsx');
    const outputDir = path.join(__dirname, '../temp-diff-test');
    
    try {
      await fs.rm(outputDir, { recursive: true });
    } catch (error) {
      // Directory might not exist
    }
    
    const result = await converter.convertFile(inputPath, outputDir, true);
    const htmlDiff = diffGenerator.generateHTMLDiff(result.changeReport!);
    const htmlPath = path.join(outputDir, 'checkbox.diff.html');
    
    await fs.writeFile(htmlPath, htmlDiff);
    
    // Verify file was written
    const fileExists = await fs.access(htmlPath).then(() => true).catch(() => false);
    expect(fileExists).toBe(true);
    
    // Verify content
    const content = await fs.readFile(htmlPath, 'utf-8');
    expect(content).toContain('<!DOCTYPE html>');
    expect(content).toContain('checkbox.tsx');
    
    // Clean up
    await fs.rm(outputDir, { recursive: true });
  });
}); 