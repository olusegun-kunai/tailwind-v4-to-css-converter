export class DiffGenerator {
    /**
     * Generate a comprehensive change report
     */
    generateChangeReport(originalContent, modifiedContent, cssContent, nodes, inputPath) {
        const changes = [];
        const originalLines = originalContent.split('\n');
        const modifiedLines = modifiedContent.split('\n');
        // Track import addition
        const hasNewImport = modifiedContent.includes('import styles from');
        if (hasNewImport) {
            const importLine = modifiedLines.findIndex(line => line.includes('import styles from'));
            changes.push({
                type: 'import-added',
                line: importLine + 1,
                original: '',
                modified: modifiedLines[importLine].trim(),
                cssClass: '',
                semanticName: 'styles-import'
            });
        }
        // Track class replacements
        for (const node of nodes) {
            const classReplacement = this.findClassReplacement(originalLines, modifiedLines, node);
            if (classReplacement) {
                changes.push(classReplacement);
            }
        }
        const summary = {
            totalNodes: nodes.length,
            classesConverted: nodes.length,
            importsAdded: hasNewImport ? 1 : 0,
            cssRulesGenerated: this.countCSSRules(cssContent)
        };
        return {
            originalFile: inputPath,
            modifiedFile: inputPath.replace('.tsx', '.tsx'), // Same file, modified
            cssModulesFile: inputPath.replace('.tsx', '.module.css'),
            changes,
            summary
        };
    }
    /**
     * Generate HTML diff view
     */
    generateHTMLDiff(changeReport) {
        const { originalFile, modifiedFile, cssModulesFile, changes, summary } = changeReport;
        return `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tailwind to CSS Modules Conversion Report</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; margin: 0; padding: 20px; background: #f8fafc; }
        .header { background: white; padding: 20px; border-radius: 8px; margin-bottom: 20px; box-shadow: 0 1px 3px rgba(0,0,0,0.1); }
        .summary { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 16px; margin: 20px 0; }
        .stat { background: #f1f5f9; padding: 16px; border-radius: 8px; text-align: center; }
        .stat-value { font-size: 24px; font-weight: bold; color: #0f172a; }
        .stat-label { font-size: 14px; color: #64748b; margin-top: 4px; }
        .diff-container { background: white; border-radius: 8px; overflow: hidden; box-shadow: 0 1px 3px rgba(0,0,0,0.1); margin-bottom: 20px; }
        .diff-header { background: #f8fafc; padding: 12px 16px; border-bottom: 1px solid #e2e8f0; font-weight: 600; }
        .diff-content { display: grid; grid-template-columns: 1fr 1fr; }
        .diff-side { padding: 16px; }
        .diff-side.original { border-right: 1px solid #e2e8f0; background: #fef2f2; }
        .diff-side.modified { background: #f0fdf4; }
        .change-item { margin-bottom: 16px; padding: 12px; border-radius: 6px; border: 1px solid #e2e8f0; }
        .change-type { font-size: 12px; font-weight: 600; text-transform: uppercase; margin-bottom: 8px; }
        .change-type.class-replacement { color: #0369a1; }
        .change-type.import-added { color: #059669; }
        .code { font-family: 'SF Mono', Monaco, monospace; font-size: 14px; background: #f8fafc; padding: 8px 12px; border-radius: 4px; overflow-x: auto; }
        .code.removed { background: #fee2e2; color: #991b1b; }
        .code.added { background: #dcfce7; color: #166534; }
        .line-number { color: #64748b; font-size: 12px; margin-bottom: 4px; }
        h1 { margin: 0; color: #0f172a; }
        h2 { color: #374151; margin-top: 0; }
        .success { color: #059669; }
        .file-path { font-family: monospace; color: #64748b; font-size: 14px; }
    </style>
</head>
<body>
    <div class="header">
        <h1>🔄 Tailwind → CSS Modules Conversion</h1>
        <p class="file-path">${originalFile}</p>
    </div>

    <div class="summary">
        <div class="stat">
            <div class="stat-value">${summary.totalNodes}</div>
            <div class="stat-label">Components Processed</div>
        </div>
        <div class="stat">
            <div class="stat-value">${summary.classesConverted}</div>
            <div class="stat-label">Classes Converted</div>
        </div>
        <div class="stat">
            <div class="stat-value">${summary.cssRulesGenerated}</div>
            <div class="stat-label">CSS Rules Generated</div>
        </div>
        <div class="stat">
            <div class="stat-value success">✅</div>
            <div class="stat-label">Conversion Status</div>
        </div>
    </div>

    <div class="diff-container">
        <div class="diff-header">📝 Changes Made</div>
        <div style="padding: 16px;">
            ${changes.map(change => `
                <div class="change-item">
                    <div class="change-type ${change.type}">${change.type.replace('-', ' ')}</div>
                    <div class="line-number">Line ${change.line}</div>
                    ${change.original ? `
                        <div class="code removed">- ${this.escapeHtml(change.original)}</div>
                    ` : ''}
                    <div class="code added">+ ${this.escapeHtml(change.modified)}</div>
                    ${change.semanticName !== 'styles-import' ? `
                        <div style="margin-top: 8px; font-size: 12px; color: #64748b;">
                            → CSS class: <code>.${change.semanticName}</code>
                        </div>
                    ` : ''}
                </div>
            `).join('')}
        </div>
    </div>

    <div class="diff-container">
        <div class="diff-header">📄 File Comparison</div>
        <div class="diff-content">
            <div class="diff-side original">
                <h2>Original (Tailwind)</h2>
                <div class="file-path">${originalFile}</div>
            </div>
            <div class="diff-side modified">
                <h2>Modified (CSS Modules)</h2>
                <div class="file-path">${modifiedFile} + ${cssModulesFile}</div>
            </div>
        </div>
    </div>
</body>
</html>`;
    }
    /**
     * Generate console diff output
     */
    generateConsoleDiff(changeReport) {
        const { changes, summary } = changeReport;
        let output = '\n🔄 Conversion Summary:\n';
        output += `   Components: ${summary.totalNodes}\n`;
        output += `   Classes converted: ${summary.classesConverted}\n`;
        output += `   CSS rules: ${summary.cssRulesGenerated}\n\n`;
        output += '📝 Changes made:\n';
        for (const change of changes) {
            output += `\n   Line ${change.line}: ${change.type}\n`;
            if (change.original) {
                output += `   - ${change.original}\n`;
            }
            output += `   + ${change.modified}\n`;
            if (change.semanticName !== 'styles-import') {
                output += `     → .${change.semanticName}\n`;
            }
        }
        return output;
    }
    findClassReplacement(originalLines, modifiedLines, node) {
        // Find the line where this node's class was changed
        const targetLine = node.line - 1; // Convert to 0-based index
        if (targetLine < originalLines.length && targetLine < modifiedLines.length) {
            const original = originalLines[targetLine];
            const modified = modifiedLines[targetLine];
            if (original !== modified && modified.includes(`styles.${node.semanticName}`)) {
                return {
                    type: 'class-replacement',
                    line: node.line,
                    original: original.trim(),
                    modified: modified.trim(),
                    cssClass: node.className,
                    semanticName: node.semanticName
                };
            }
        }
        return null;
    }
    countCSSRules(cssContent) {
        // Count CSS rules by counting opening braces
        return (cssContent.match(/\{/g) || []).length;
    }
    escapeHtml(text) {
        return text
            .replace(/&/g, '&amp;')
            .replace(/</g, '&lt;')
            .replace(/>/g, '&gt;')
            .replace(/"/g, '&quot;')
            .replace(/'/g, '&#39;');
    }
}
//# sourceMappingURL=diffGenerator.js.map