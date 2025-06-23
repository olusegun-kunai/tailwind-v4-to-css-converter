import fs from 'fs/promises';
import path from 'path';
import { ConversionResult, JSXNode, GeneratedCSS } from '../types/index.js';
import { UnoGenerator } from './unoGenerator.js';
import { JSXParser } from './jsxParser.js';
import { DiffGenerator, ChangeReport } from './diffGenerator.js';

export class QwikTailwindConverter {
  private unoGenerator: UnoGenerator;
  private jsxParser: JSXParser;
  private diffGenerator: DiffGenerator;

  constructor() {
    this.unoGenerator = new UnoGenerator();
    this.jsxParser = new JSXParser();
    this.diffGenerator = new DiffGenerator();
  }

  /**
   * Convert a Qwik component file to CSS modules
   * @param inputPath - Path to the input .tsx file
   * @param outputDir - Directory to write output files
   * @param generateDiff - Whether to generate diff report (default: false)
   * @returns Conversion result with file paths and content
   */
  async convertFile(inputPath: string, outputDir: string, generateDiff: boolean = false): Promise<ConversionResult> {
    try {
      // Read input file
      const jsxContent = await fs.readFile(inputPath, 'utf-8');
      
      // Parse JSX to extract nodes with classes
      console.log('📄 Parsing JSX nodes...');
      const jsxNodes = this.jsxParser.parseMultiLineClasses(jsxContent);
      
      if (jsxNodes.length === 0) {
        throw new Error('No JSX nodes with class attributes found');
      }

      console.log(`🔍 Found ${jsxNodes.length} nodes with classes`);
      
      // Generate CSS for each node
      console.log('🎨 Generating CSS with UnoCSS...');
      const generatedCSS = await this.generateCSSForNodes(jsxNodes);
      
      // Create CSS modules content
      const cssModulesContent = this.createCSSModules(generatedCSS);
      
      // Update JSX to use CSS modules
      const updatedJSX = this.updateJSXWithModules(jsxContent, jsxNodes, inputPath);
      
      // Ensure output directory exists
      await fs.mkdir(outputDir, { recursive: true });
      
      // Generate output file names
      const baseName = path.basename(inputPath, '.tsx');
      const cssModulesPath = path.join(outputDir, `${baseName}.module.css`);
      const componentPath = path.join(outputDir, `${baseName}.tsx`);
      
      // Write output files
      await fs.writeFile(cssModulesPath, cssModulesContent);
      await fs.writeFile(componentPath, updatedJSX);
      
      // Generate diff report if requested
      let changeReport: ChangeReport | undefined;
      if (generateDiff) {
        changeReport = this.diffGenerator.generateChangeReport(
          jsxContent,
          updatedJSX,
          cssModulesContent,
          jsxNodes,
          inputPath
        );
      }
      
      return {
        cssModulesPath,
        componentPath,
        cssContent: cssModulesContent,
        updatedComponent: updatedJSX,
        changeReport
      };
      
    } catch (error) {
      throw new Error(`Conversion failed: ${error}`);
    }
  }

  /**
   * Generate CSS for all JSX nodes using UnoCSS
   * @param nodes - Array of JSX nodes
   * @returns Array of generated CSS rules
   */
  private async generateCSSForNodes(nodes: JSXNode[]): Promise<GeneratedCSS[]> {
    const generatedCSS: GeneratedCSS[] = [];
    
    for (const node of nodes) {
      console.log(`  🔄 Processing ${node.tagName} (${node.semanticName})`);
      
      try {
        const result = await this.unoGenerator.generateCSS(node.className);
        
        if (result.css) {
          // Parse the CSS to separate base rules and modifiers
          const parsed = this.unoGenerator.parseModifiers(result.css);
          
          const cssRule: GeneratedCSS = {
            selector: node.semanticName,
            rules: this.extractCSSRules(parsed.base),
            modifiers: parsed.modifiers.map(mod => ({
              pseudo: mod.pseudo,
              rules: mod.rules
            }))
          };
          
          generatedCSS.push(cssRule);
        }
        
      } catch (error) {
        console.warn(`⚠️  Failed to generate CSS for ${node.tagName}: ${error}`);
      }
    }
    
    return generatedCSS;
  }

  /**
   * Extract CSS rules from generated CSS string
   * @param css - Raw CSS string
   * @returns Clean CSS rules
   */
  private extractCSSRules(css: string): string {
    // Split CSS by layers and only get the default layer (actual utility classes)
    const defaultLayerMatch = css.match(/\/\* layer: default \*\/\s*([\s\S]*?)(?=\/\*|$)/);
    
    if (defaultLayerMatch) {
      const defaultLayerCSS = defaultLayerMatch[1].trim();
      
      // Extract just the CSS properties from each rule
      const rules: string[] = [];
      const ruleMatches = defaultLayerCSS.match(/\.[^{]+\{([^}]+)\}/g);
      
      if (ruleMatches) {
        for (const rule of ruleMatches) {
          const propertyMatch = rule.match(/\{([^}]+)\}/);
          if (propertyMatch) {
            rules.push(propertyMatch[1].trim());
          }
        }
      }
      
      return rules.join('; ');
    }
    
    return '';
  }

  /**
   * Create CSS modules content from generated CSS
   * @param generatedCSS - Array of generated CSS rules
   * @returns CSS modules content as string
   */
  private createCSSModules(generatedCSS: GeneratedCSS[]): string {
    let cssContent = '/* Generated CSS Modules from Tailwind classes */\n\n';
    
    for (const css of generatedCSS) {
      // Base styles
      if (css.rules.trim()) {
        cssContent += `.${css.selector} {\n`;
        // Clean up the CSS rules formatting
        const cleanRules = css.rules
          .split(';')
          .map(rule => rule.trim())
          .filter(rule => rule)
          .join(';\n  ');
        cssContent += `  ${cleanRules};\n`;
        cssContent += '}\n\n';
      }
      
      // Modifier styles (focus, hover, etc.)
      if (css.modifiers && css.modifiers.length > 0) {
        for (const modifier of css.modifiers) {
          cssContent += `.${css.selector}:${modifier.pseudo} {\n`;
          const cleanModRules = modifier.rules
            .split(';')
            .map(rule => rule.trim())
            .filter(rule => rule)
            .join(';\n  ');
          cssContent += `  ${cleanModRules};\n`;
          cssContent += '}\n\n';
        }
      }
    }
    
    return cssContent;
  }

  /**
   * Update JSX content to use CSS modules
   * @param jsxContent - Original JSX content
   * @param nodes - Array of JSX nodes
   * @param inputPath - Original input file path for naming
   * @returns Updated JSX with CSS module imports
   */
  private updateJSXWithModules(jsxContent: string, nodes: JSXNode[], inputPath: string): string {
    let updatedContent = jsxContent;
    
    // Add CSS modules import at the top  
    const baseName = inputPath.split('/').pop()?.replace('.tsx', '') || 'component';
    const importStatement = `import styles from './${baseName}.module.css';\n`;
    
    // Find the first import line and add our import after it
    const importRegex = /^import.*?from.*?;$/gm;
    const imports = updatedContent.match(importRegex);
    
    if (imports) {
      const lastImport = imports[imports.length - 1];
      const lastImportIndex = updatedContent.indexOf(lastImport) + lastImport.length;
      updatedContent = 
        updatedContent.slice(0, lastImportIndex) + 
        '\n' + importStatement + 
        updatedContent.slice(lastImportIndex);
    } else {
      // If no imports found, add at the beginning
      updatedContent = importStatement + '\n' + updatedContent;
    }
    
    // Replace class attributes with CSS module references
    for (const node of nodes) {
      // Create regex to match the specific class attribute for this node
      const escapedTagName = node.tagName.replace(/\./g, '\\.');
      const classRegex = new RegExp(
        `(<${escapedTagName}[^>]*)((?:class|className)=")([^"']+)(")`,
        'g'
      );
      
      updatedContent = updatedContent.replace(classRegex, `$1className={styles.${node.semanticName}}`);
    }
    
    return updatedContent;
  }
} 