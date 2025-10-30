import fs from 'fs/promises';
import path from 'path';
/**
 * Load and validate theme configuration from various file formats
 */
export class ThemeLoader {
    /**
     * Load theme configuration from file
     * @param themePath - Path to theme configuration file (.json, .js, .ts)
     * @param basePath - Base path to resolve relative paths from (default: process.cwd())
     * @returns Promise<ThemeLoadResult> - Loaded and validated theme configuration
     */
    async loadTheme(themePath, basePath = process.cwd()) {
        try {
            // Resolve the absolute path
            const resolvedPath = path.isAbsolute(themePath)
                ? themePath
                : path.resolve(basePath, themePath);
            // Check if file exists
            await this.validateFileExists(resolvedPath);
            // Determine file type and load accordingly
            const ext = path.extname(resolvedPath).toLowerCase();
            let themeData;
            switch (ext) {
                case '.json':
                    themeData = await this.loadJsonTheme(resolvedPath);
                    break;
                case '.js':
                case '.ts':
                    themeData = await this.loadEsmTheme(resolvedPath);
                    break;
                default:
                    throw new Error(`Unsupported theme file format: ${ext}. Supported formats: .json, .js, .ts`);
            }
            // Validate and normalize theme structure
            const normalizedTheme = this.normalizeTheme(themeData);
            return {
                theme: normalizedTheme,
                path: resolvedPath
            };
        }
        catch (error) {
            throw new Error(`Failed to load theme from "${themePath}": ${error}`);
        }
    }
    /**
     * Validate that the theme file exists
     * @param filePath - Absolute path to theme file
     */
    async validateFileExists(filePath) {
        try {
            await fs.access(filePath);
        }
        catch {
            throw new Error(`Theme file not found: ${filePath}`);
        }
    }
    /**
     * Load theme from JSON file
     * @param filePath - Absolute path to JSON file
     * @returns Parsed JSON object
     */
    async loadJsonTheme(filePath) {
        try {
            const content = await fs.readFile(filePath, 'utf-8');
            return JSON.parse(content);
        }
        catch (error) {
            throw new Error(`Invalid JSON in theme file: ${error}`);
        }
    }
    /**
     * Load theme from ESM module (.js or .ts)
     * @param filePath - Absolute path to module file
     * @returns Exported theme object
     */
    async loadEsmTheme(filePath) {
        try {
            // Convert to file:// URL for import()
            const fileUrl = new URL('file://' + filePath).href;
            const module = await import(fileUrl);
            // Handle both default and named exports
            return module.default || module.theme || module;
        }
        catch (error) {
            throw new Error(`Failed to import theme module: ${error}`);
        }
    }
    /**
     * Normalize and validate theme structure for UnoCSS compatibility
     * @param themeData - Raw theme data
     * @returns Normalized ThemeConfig
     */
    normalizeTheme(themeData) {
        if (!themeData || typeof themeData !== 'object') {
            throw new Error('Theme must be an object');
        }
        const validKeys = ['colors', 'fontSize', 'spacing', 'borderRadius', 'fontWeight'];
        const normalized = {};
        // Validate and copy known theme keys
        for (const [key, value] of Object.entries(themeData)) {
            if (validKeys.includes(key)) {
                if (typeof value === 'object' && value !== null && !Array.isArray(value)) {
                    normalized[key] = value;
                }
                else {
                    console.warn(`⚠️  Theme key "${key}" should be an object, skipping`);
                }
            }
            else {
                console.warn(`⚠️  Unknown theme key "${key}", skipping (supported: ${validKeys.join(', ')})`);
            }
        }
        // Validate specific theme sections
        this.validateThemeSection('colors', normalized.colors);
        this.validateThemeSection('fontSize', normalized.fontSize);
        this.validateThemeSection('spacing', normalized.spacing);
        this.validateThemeSection('borderRadius', normalized.borderRadius);
        this.validateThemeSection('fontWeight', normalized.fontWeight);
        return normalized;
    }
    /**
     * Validate a specific theme section
     * @param sectionName - Name of the theme section
     * @param section - Theme section data
     */
    validateThemeSection(sectionName, section) {
        if (!section)
            return;
        // Basic validation for common theme sections
        switch (sectionName) {
            case 'colors':
                this.validateColors(section);
                break;
            case 'fontSize':
                this.validateFontSize(section);
                break;
            case 'spacing':
                this.validateSpacing(section);
                break;
            case 'borderRadius':
                this.validateBorderRadius(section);
                break;
            case 'fontWeight':
                this.validateFontWeight(section);
                break;
        }
    }
    /**
     * Validate colors section
     * @param colors - Colors configuration
     */
    validateColors(colors) {
        for (const [colorName, colorValue] of Object.entries(colors)) {
            if (typeof colorValue === 'string') {
                // Simple color value like "red": "#ff0000"
                continue;
            }
            else if (typeof colorValue === 'object' && colorValue !== null) {
                // Color scale like "red": { "500": "#ff0000", "600": "#dc2626" }
                for (const [shade, value] of Object.entries(colorValue)) {
                    if (typeof value !== 'string') {
                        console.warn(`⚠️  Invalid color value at colors.${colorName}.${shade}: expected string, got ${typeof value}`);
                    }
                }
            }
            else {
                console.warn(`⚠️  Invalid color value at colors.${colorName}: expected string or object, got ${typeof colorValue}`);
            }
        }
    }
    /**
     * Validate fontSize section
     * @param fontSize - Font size configuration
     */
    validateFontSize(fontSize) {
        for (const [sizeName, sizeValue] of Object.entries(fontSize)) {
            if (typeof sizeValue === 'string') {
                // Simple size like "xl": "1.25rem"
                continue;
            }
            else if (Array.isArray(sizeValue) && sizeValue.length >= 1) {
                // Array format like "xl": ["1.25rem", "1.75rem"] or ["1.25rem", { lineHeight: "1.75rem" }]
                continue;
            }
            else {
                console.warn(`⚠️  Invalid fontSize value at fontSize.${sizeName}: expected string or array, got ${typeof sizeValue}`);
            }
        }
    }
    /**
     * Validate spacing section
     * @param spacing - Spacing configuration
     */
    validateSpacing(spacing) {
        for (const [spaceName, spaceValue] of Object.entries(spacing)) {
            if (typeof spaceValue !== 'string') {
                console.warn(`⚠️  Invalid spacing value at spacing.${spaceName}: expected string, got ${typeof spaceValue}`);
            }
        }
    }
    /**
     * Validate borderRadius section
     * @param borderRadius - Border radius configuration
     */
    validateBorderRadius(borderRadius) {
        for (const [radiusName, radiusValue] of Object.entries(borderRadius)) {
            if (typeof radiusValue !== 'string') {
                console.warn(`⚠️  Invalid borderRadius value at borderRadius.${radiusName}: expected string, got ${typeof radiusValue}`);
            }
        }
    }
    /**
     * Validate fontWeight section
     * @param fontWeight - Font weight configuration
     */
    validateFontWeight(fontWeight) {
        for (const [weightName, weightValue] of Object.entries(fontWeight)) {
            if (typeof weightValue !== 'string' && typeof weightValue !== 'number') {
                console.warn(`⚠️  Invalid fontWeight value at fontWeight.${weightName}: expected string or number, got ${typeof weightValue}`);
            }
        }
    }
    /**
     * Print theme summary for verbose mode
     * @param themeResult - Loaded theme result
     */
    printThemeSummary(themeResult) {
        console.log(`📄 Loaded theme from: ${themeResult.path}`);
        const theme = themeResult.theme;
        const sections = Object.keys(theme).filter(key => theme[key] && Object.keys(theme[key]).length > 0);
        if (sections.length === 0) {
            console.log('  📝 No theme overrides found');
            return;
        }
        console.log('  📝 Theme overrides:');
        for (const section of sections) {
            const sectionData = theme[section];
            const keyCount = Object.keys(sectionData).length;
            console.log(`    • ${section}: ${keyCount} ${keyCount === 1 ? 'override' : 'overrides'}`);
        }
    }
}
// Export a default instance for convenience
export const themeLoader = new ThemeLoader();
//# sourceMappingURL=themeLoader.js.map