import { UnoGeneratorResult, ThemeConfig } from '../types/index.js';
export declare class UnoGenerator {
    private generator;
    private customTheme;
    constructor(theme?: ThemeConfig);
    private initGenerator;
    /**
     * Generate CSS from Tailwind classes
     * @param classes - Space-separated Tailwind classes
     * @returns Generated CSS string and matched classes
     */
    generateCSS(classes: string): Promise<UnoGeneratorResult>;
    /**
     * Extract individual CSS rules from generated CSS
     * @param css - Generated CSS string
     * @returns Array of CSS rules
     */
    extractRules(css: string): string[];
    /**
     * Parse CSS to extract base rules and modifiers
     * @param css - Generated CSS string
     * @returns Object with base rules and modifiers
     */
    parseModifiers(css: string): {
        base: string;
        modifiers: Array<{
            pseudo: string;
            rules: string;
        }>;
    };
}
