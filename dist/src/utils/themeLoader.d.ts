import { ThemeLoadResult } from '../types/index.js';
/**
 * Load and validate theme configuration from various file formats
 */
export declare class ThemeLoader {
    /**
     * Load theme configuration from file
     * @param themePath - Path to theme configuration file (.json, .js, .ts)
     * @param basePath - Base path to resolve relative paths from (default: process.cwd())
     * @returns Promise<ThemeLoadResult> - Loaded and validated theme configuration
     */
    loadTheme(themePath: string, basePath?: string): Promise<ThemeLoadResult>;
    /**
     * Validate that the theme file exists
     * @param filePath - Absolute path to theme file
     */
    private validateFileExists;
    /**
     * Load theme from JSON file
     * @param filePath - Absolute path to JSON file
     * @returns Parsed JSON object
     */
    private loadJsonTheme;
    /**
     * Load theme from ESM module (.js or .ts)
     * @param filePath - Absolute path to module file
     * @returns Exported theme object
     */
    private loadEsmTheme;
    /**
     * Normalize and validate theme structure for UnoCSS compatibility
     * @param themeData - Raw theme data
     * @returns Normalized ThemeConfig
     */
    private normalizeTheme;
    /**
     * Validate a specific theme section
     * @param sectionName - Name of the theme section
     * @param section - Theme section data
     */
    private validateThemeSection;
    /**
     * Validate colors section
     * @param colors - Colors configuration
     */
    private validateColors;
    /**
     * Validate fontSize section
     * @param fontSize - Font size configuration
     */
    private validateFontSize;
    /**
     * Validate spacing section
     * @param spacing - Spacing configuration
     */
    private validateSpacing;
    /**
     * Validate borderRadius section
     * @param borderRadius - Border radius configuration
     */
    private validateBorderRadius;
    /**
     * Validate fontWeight section
     * @param fontWeight - Font weight configuration
     */
    private validateFontWeight;
    /**
     * Print theme summary for verbose mode
     * @param themeResult - Loaded theme result
     */
    printThemeSummary(themeResult: ThemeLoadResult): void;
}
export declare const themeLoader: ThemeLoader;
