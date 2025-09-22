import { JSXNode } from '../types/index.js';
export declare class JSXParser {
    private nodeCounter;
    /**
     * Parse JSX string and extract nodes with class attributes
     * @param jsxString - The JSX component as a string
     * @returns Array of JSX nodes with their class information
     */
    parseJSX(jsxString: string): JSXNode[];
    /**
     * Extract JSX nodes from a single line
     * @param line - Line of code to parse
     * @param lineNumber - Line number for reference
     * @returns Array of JSX nodes found in the line
     */
    private extractJSXNodesFromLine;
    /**
     * Determine if a tag is a component or HTML element
     * @param tagName - The JSX tag name
     * @returns 'component' or 'html'
     */
    private determineNodeType;
    /**
     * Generate semantic class name based on tag name
     * @param tagName - The JSX tag name
     * @returns Semantic class name
     */
    private generateSemanticName;
    /**
     * Extract multi-line class attributes (for classes spanning multiple lines)
     * @param jsxString - Full JSX string
     * @returns Array of JSX nodes including multi-line classes
     */
    parseMultiLineClasses(jsxString: string): JSXNode[];
    /**
     * Clean and normalize class string
     * @param classString - Raw class string from JSX
     * @returns Cleaned class string
     */
    cleanClassName(classString: string): string;
}
