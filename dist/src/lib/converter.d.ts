import { ConversionResult, ThemeConfig } from '../types/index.js';
export declare class QwikTailwindConverter {
    private unoGenerator;
    private jsxParser;
    private diffGenerator;
    constructor(theme?: ThemeConfig);
    /**
     * Convert a Qwik component file to CSS modules
     * @param inputPath - Path to the input .tsx file
     * @param outputDir - Directory to write output files
     * @param generateDiff - Whether to generate diff report (default: false)
     * @returns Conversion result with file paths and content
     */
    convertFile(inputPath: string, outputDir: string, generateDiff?: boolean): Promise<ConversionResult>;
    /**
     * Generate CSS for all JSX nodes using UnoCSS
     * @param nodes - Array of JSX nodes
     * @returns Array of generated CSS rules
     */
    private generateCSSForNodes;
    /**
     * Extract CSS rules from generated CSS string
     * @param css - Raw CSS string
     * @returns Clean CSS rules
     */
    private extractCSSRules;
    /**
     * Create CSS modules content from generated CSS
     * @param generatedCSS - Array of generated CSS rules
     * @returns CSS modules content as string
     */
    private createCSSModules;
    /**
     * Update JSX content to use CSS modules
     * @param jsxContent - Original JSX content
     * @param nodes - Array of JSX nodes
     * @param inputPath - Original input file path for naming
     * @returns Updated JSX with CSS module imports
     */
    private updateJSXWithModules;
}
