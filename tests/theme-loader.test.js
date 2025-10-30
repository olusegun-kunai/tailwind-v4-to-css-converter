import { ThemeLoader } from '../src/utils/themeLoader.js';
import fs from 'fs/promises';
import path from 'path';
async function testThemeLoader() {
    console.log('🧪 Testing Theme Loader...\n');
    const loader = new ThemeLoader();
    const testDir = './test-themes';
    // Create test directory and files
    await setupTestFiles(testDir);
    try {
        // Test 1: Load valid JSON theme
        console.log('📝 Test 1: Loading valid JSON theme...');
        try {
            const result = await loader.loadTheme(`${testDir}/valid.json`);
            console.log('✅ JSON theme loaded successfully');
            console.log('  Path:', result.path);
            console.log('  Theme keys:', Object.keys(result.theme));
            console.log('  Colors:', Object.keys(result.theme.colors || {}));
            console.log('---\n');
        }
        catch (error) {
            console.error('❌ JSON theme test failed:', error);
        }
        // Test 2: Load valid TypeScript theme
        console.log('📝 Test 2: Loading valid TypeScript theme...');
        try {
            const result = await loader.loadTheme(`${testDir}/valid.ts`);
            console.log('✅ TypeScript theme loaded successfully');
            console.log('  Path:', result.path);
            console.log('  Theme keys:', Object.keys(result.theme));
            console.log('---\n');
        }
        catch (error) {
            console.error('❌ TypeScript theme test failed:', error);
        }
        // Test 3: Load minimal theme
        console.log('📝 Test 3: Loading minimal theme...');
        try {
            const result = await loader.loadTheme(`${testDir}/minimal.json`);
            console.log('✅ Minimal theme loaded successfully');
            console.log('  Theme sections:', Object.keys(result.theme));
            console.log('---\n');
        }
        catch (error) {
            console.error('❌ Minimal theme test failed:', error);
        }
        // Test 4: Missing file error
        console.log('📝 Test 4: Testing missing file error...');
        try {
            await loader.loadTheme(`${testDir}/nonexistent.json`);
            console.error('❌ Expected error for missing file, but none was thrown');
        }
        catch (error) {
            console.log('✅ Correctly caught missing file error:', error.message);
            console.log('---\n');
        }
        // Test 5: Invalid JSON error
        console.log('📝 Test 5: Testing invalid JSON error...');
        try {
            await loader.loadTheme(`${testDir}/invalid.json`);
            console.error('❌ Expected error for invalid JSON, but none was thrown');
        }
        catch (error) {
            console.log('✅ Correctly caught invalid JSON error:', error.message);
            console.log('---\n');
        }
        // Test 6: Invalid theme structure
        console.log('📝 Test 6: Testing invalid theme structure...');
        try {
            await loader.loadTheme(`${testDir}/invalid-structure.json`);
            console.error('❌ Expected error for invalid structure, but none was thrown');
        }
        catch (error) {
            console.log('✅ Correctly caught invalid structure error:', error.message);
            console.log('---\n');
        }
        // Test 7: Empty theme file
        console.log('📝 Test 7: Testing empty theme file...');
        try {
            const result = await loader.loadTheme(`${testDir}/empty.json`);
            console.log('✅ Empty theme loaded (should have no sections)');
            console.log('  Theme sections:', Object.keys(result.theme));
            console.log('---\n');
        }
        catch (error) {
            console.error('❌ Empty theme test failed:', error);
        }
        // Test 8: Theme with warnings
        console.log('📝 Test 8: Testing theme with warnings...');
        try {
            const result = await loader.loadTheme(`${testDir}/with-warnings.json`);
            console.log('✅ Theme with warnings loaded');
            console.log('  Valid sections:', Object.keys(result.theme));
            console.log('---\n');
        }
        catch (error) {
            console.error('❌ Theme with warnings test failed:', error);
        }
        // Test 9: Relative path resolution
        console.log('📝 Test 9: Testing relative path resolution...');
        try {
            const result = await loader.loadTheme('./examples/themes/minimal.json');
            console.log('✅ Relative path resolved successfully');
            console.log('  Resolved to:', result.path);
            console.log('---\n');
        }
        catch (error) {
            console.error('❌ Relative path test failed:', error);
        }
        // Test 10: Theme summary printing
        console.log('📝 Test 10: Testing theme summary printing...');
        try {
            const result = await loader.loadTheme('./examples/themes/custom.json');
            console.log('✅ Testing theme summary output:');
            loader.printThemeSummary(result);
            console.log('---\n');
        }
        catch (error) {
            console.error('❌ Theme summary test failed:', error);
        }
    }
    finally {
        // Cleanup test files
        await cleanupTestFiles(testDir);
    }
    console.log('🎉 Theme loader tests completed!');
}
async function setupTestFiles(testDir) {
    await fs.mkdir(testDir, { recursive: true });
    // Valid JSON theme
    await fs.writeFile(path.join(testDir, 'valid.json'), JSON.stringify({
        colors: {
            red: { 500: '#ff0000' },
            blue: { 500: '#0000ff' }
        },
        fontSize: {
            xl: '1.25rem',
            '2xl': '1.5rem'
        },
        spacing: {
            4: '1rem',
            8: '2rem'
        }
    }, null, 2));
    // Valid TypeScript theme
    await fs.writeFile(path.join(testDir, 'valid.ts'), `
export default {
  colors: {
    green: { 500: '#00ff00' }
  },
  fontSize: {
    lg: '1.125rem'
  }
};
  `.trim());
    // Minimal theme
    await fs.writeFile(path.join(testDir, 'minimal.json'), JSON.stringify({
        colors: {
            primary: '#000000'
        }
    }, null, 2));
    // Invalid JSON
    await fs.writeFile(path.join(testDir, 'invalid.json'), '{ invalid json content }');
    // Invalid structure (not an object)
    await fs.writeFile(path.join(testDir, 'invalid-structure.json'), '"not an object"');
    // Empty theme
    await fs.writeFile(path.join(testDir, 'empty.json'), '{}');
    // Theme with warnings (unknown keys and invalid values)
    await fs.writeFile(path.join(testDir, 'with-warnings.json'), JSON.stringify({
        colors: {
            red: { 500: '#ff0000' }
        },
        unknownKey: 'should be ignored',
        fontSize: {
            xl: 123 // invalid type, should warn
        },
        invalidColors: 'not an object'
    }, null, 2));
}
async function cleanupTestFiles(testDir) {
    try {
        await fs.rm(testDir, { recursive: true, force: true });
    }
    catch (error) {
        console.warn('Failed to cleanup test files:', error);
    }
}
// Run the test
testThemeLoader().catch(console.error);
//# sourceMappingURL=theme-loader.test.js.map