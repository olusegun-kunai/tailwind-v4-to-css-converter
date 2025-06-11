import { createGenerator } from '@unocss/core';
import presetWind from '@unocss/preset-wind';
import { UnoGeneratorResult } from '../types/index.js';

export class UnoGenerator {
  private generator: any;

  constructor() {
    // UnoCSS createGenerator returns a Promise
    this.generator = null;
    this.initGenerator();
  }

  private async initGenerator() {
    this.generator = await createGenerator({
      presets: [presetWind()], // Tailwind 4 preset
      theme: {
        // Add any custom theme configuration here
      }
    });
  }

  /**
   * Generate CSS from Tailwind classes
   * @param classes - Space-separated Tailwind classes
   * @returns Generated CSS string and matched classes
   */
  async generateCSS(classes: string): Promise<UnoGeneratorResult> {
    try {
      // Ensure generator is initialized
      if (!this.generator) {
        await this.initGenerator();
      }

      // UnoCSS generates CSS from a set of tokens
      const tokens = new Set(classes.split(/\s+/).filter(c => c.trim()));
      const result = await this.generator.generate(tokens);
      
      return {
        css: result.css, // This is a getter that returns the CSS string
        matched: result.matched
      };
    } catch (error) {
      console.error('Error generating CSS with UnoCSS:', error);
      throw new Error(`UnoCSS generation failed: ${error}`);
    }
  }

  /**
   * Extract individual CSS rules from generated CSS
   * @param css - Generated CSS string
   * @returns Array of CSS rules
   */
  extractRules(css: string): string[] {
    // Split CSS into individual rules
    const rules = css.split('}').filter(rule => rule.trim());
    return rules.map(rule => rule.trim() + '}');
  }

  /**
   * Parse CSS to extract base rules and modifiers
   * @param css - Generated CSS string
   * @returns Object with base rules and modifiers
   */
  parseModifiers(css: string): { base: string; modifiers: Array<{ pseudo: string; rules: string }> } {
    const base: string[] = [];
    const modifiers: Array<{ pseudo: string; rules: string }> = [];

    const rules = this.extractRules(css);
    
    for (const rule of rules) {
      if (rule.includes(':')) {
        // Check if it's a pseudo-class rule
        const match = rule.match(/\.([\w-]+):([\w-]+)\s*\{([^}]+)\}/);
        if (match) {
          const [, , pseudo, cssRules] = match;
          modifiers.push({
            pseudo: pseudo,
            rules: cssRules.trim()
          });
        } else {
          base.push(rule);
        }
      } else {
        base.push(rule);
      }
    }

    return {
      base: base.join('\n'),
      modifiers
    };
  }
} 