import { JSXNode } from '../types/index.js';

export class JSXParser {
  private nodeCounter = 0;

  /**
   * Parse JSX string and extract nodes with class attributes
   * @param jsxString - The JSX component as a string
   * @returns Array of JSX nodes with their class information
   */
  parseJSX(jsxString: string): JSXNode[] {
    const lines = jsxString.split('\n');
    const nodes: JSXNode[] = [];
    this.nodeCounter = 0;

    for (let i = 0; i < lines.length; i++) {
      const line = lines[i];
      const jsxNodes = this.extractJSXNodesFromLine(line, i + 1);
      nodes.push(...jsxNodes);
    }

    return nodes;
  }

  /**
   * Extract JSX nodes from a single line
   * @param line - Line of code to parse
   * @param lineNumber - Line number for reference
   * @returns Array of JSX nodes found in the line
   */
  private extractJSXNodesFromLine(line: string, lineNumber: number): JSXNode[] {
    const nodes: JSXNode[] = [];
    
    // Regex to match JSX opening tags with class attributes
    // Handles both class="..." and className="..."
    const jsxRegex = /<(\w+(?:\.\w+)*)\s+[^>]*(?:class|className)=["']([^"']+)["'][^>]*>/g;
    
    let match;
    while ((match = jsxRegex.exec(line)) !== null) {
      const [, tagName, className] = match;
      
      const node: JSXNode = {
        type: this.determineNodeType(tagName),
        tagName,
        className: className.trim(),
        line: lineNumber,
        semanticName: this.generateSemanticName(tagName)
      };
      
      nodes.push(node);
    }

    return nodes;
  }

  /**
   * Determine if a tag is a component or HTML element
   * @param tagName - The JSX tag name
   * @returns 'component' or 'html'
   */
  private determineNodeType(tagName: string): 'component' | 'html' {
    // Component if it starts with uppercase or contains dots (e.g., Checkbox.Indicator)
    if (tagName[0] === tagName[0].toUpperCase() || tagName.includes('.')) {
      return 'component';
    }
    return 'html';
  }

  /**
   * Generate semantic class name based on tag name
   * @param tagName - The JSX tag name
   * @returns Semantic class name
   */
  private generateSemanticName(tagName: string): string {
    if (this.determineNodeType(tagName) === 'component') {
      // For components like Checkbox.Indicator -> indicator
      if (tagName.includes('.')) {
        const parts = tagName.split('.');
        return parts[parts.length - 1].toLowerCase();
      }
      // For components like Button -> button
      return tagName.toLowerCase();
    } else {
      // For HTML elements -> node0, node1, etc.
      return `node${this.nodeCounter++}`;
    }
  }

  /**
   * Extract multi-line class attributes (for classes spanning multiple lines)
   * @param jsxString - Full JSX string
   * @returns Array of JSX nodes including multi-line classes
   */
  parseMultiLineClasses(jsxString: string): JSXNode[] {
    const nodes: JSXNode[] = [];
    this.nodeCounter = 0;
    
    // Regex to match multi-line JSX tags with class attributes
    const multiLineRegex = /<(\w+(?:\.\w+)*)\s+[^>]*(?:class|className)=["']([^"']*(?:\s*[^"']*)*?)["'][^>]*>/gs;
    
    let match;
    while ((match = multiLineRegex.exec(jsxString)) !== null) {
      const [fullMatch, tagName, className] = match;
      
      // Find line number by counting newlines before the match
      const beforeMatch = jsxString.substring(0, match.index);
      const lineNumber = (beforeMatch.match(/\n/g) || []).length + 1;
      
      const node: JSXNode = {
        type: this.determineNodeType(tagName),
        tagName,
        className: className.replace(/\s+/g, ' ').trim(), // Normalize whitespace
        line: lineNumber,
        semanticName: this.generateSemanticName(tagName)
      };
      
      nodes.push(node);
    }

    return nodes;
  }

  /**
   * Clean and normalize class string
   * @param classString - Raw class string from JSX
   * @returns Cleaned class string
   */
  cleanClassName(classString: string): string {
    return classString
      .replace(/\s+/g, ' ') // Replace multiple spaces with single space
      .trim();
  }
} 