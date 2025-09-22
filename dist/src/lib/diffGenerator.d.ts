import { JSXNode } from '../types/index.js';
export interface ChangeReport {
    originalFile: string;
    modifiedFile: string;
    cssModulesFile: string;
    changes: Change[];
    summary: ChangeSummary;
}
export interface Change {
    type: 'class-replacement' | 'import-added';
    line: number;
    original: string;
    modified: string;
    cssClass: string;
    semanticName: string;
}
export interface ChangeSummary {
    totalNodes: number;
    classesConverted: number;
    importsAdded: number;
    cssRulesGenerated: number;
}
export declare class DiffGenerator {
    /**
     * Generate a comprehensive change report
     */
    generateChangeReport(originalContent: string, modifiedContent: string, cssContent: string, nodes: JSXNode[], inputPath: string): ChangeReport;
    /**
     * Generate HTML diff view
     */
    generateHTMLDiff(changeReport: ChangeReport): string;
    /**
     * Generate console diff output
     */
    generateConsoleDiff(changeReport: ChangeReport): string;
    private findClassReplacement;
    private countCSSRules;
    private escapeHtml;
}
