import { QwikTailwindConverter } from '../src/lib/converter.js';
import { themeLoader } from '../src/utils/themeLoader.js';
import fs from 'fs/promises';
import path from 'path';
async function testThemeIntegration() {
    console.log('🧪 Testing Theme Integration...\n');
    const testDir = './test-integration';
    await fs.mkdir(testDir, { recursive: true });
    try {
        // Create a test component
        const testComponent = `
import { component$ } from '@builder.io/qwik';

export default component$(() => {
  return (
    <div class="bg-red-500 text-blue-500 text-2xl p-4">
      <span class="font-bold">Theme Test</span>
    </div>
  );
});
    `.trim();
        const componentPath = path.join(testDir, 'test.tsx');
        await fs.writeFile(componentPath, testComponent);
        // Test 1: Conversion without theme (default behavior)
        console.log('📝 Test 1: Conversion without theme...');
        try {
            const converter1 = new QwikTailwindConverter();
            const result1 = await converter1.convertFile(componentPath, `${testDir}/output1`);
            const css1 = await fs.readFile(result1.cssModulesPath, 'utf-8');
            // Check for default Tailwind colors
            const hasDefaultRed = css1.includes('rgb(239 68 68'); // default red-500
            const hasDefaultBlue = css1.includes('rgb(59 130 246'); // default blue-500
            const hasDefault2xl = css1.includes('font-size:1.5rem'); // default 2xl
            if (hasDefaultRed && hasDefaultBlue && hasDefault2xl) {
                console.log('✅ Default theme working correctly');
            }
            else {
                console.error('❌ Default theme not working as expected');
                console.log('  CSS content:', css1);
            }
        }
        catch (error) {
            console.error('❌ Default theme test failed:', error);
        }
        // Test 2: Conversion with custom theme
        console.log('\n📝 Test 2: Conversion with custom theme...');
        try {
            const themeResult = await themeLoader.loadTheme('./examples/themes/minimal.json');
            const converter2 = new QwikTailwindConverter(themeResult.theme);
            const result2 = await converter2.convertFile(componentPath, `${testDir}/output2`);
            const css2 = await fs.readFile(result2.cssModulesPath, 'utf-8');
            // Check for custom theme colors  
            const hasCustomRed = css2.includes('rgb(255 85 85'); // custom #ff5555
            const hasCustomBlue = css2.includes('rgb(85 85 255'); // custom #5555ff
            const hasCustom2xl = css2.includes('font-size:2rem'); // custom 2rem
            if (hasCustomRed && hasCustomBlue && hasCustom2xl) {
                console.log('✅ Custom theme working correctly');
                console.log('  🎨 Custom red-500: #ff5555 detected');
                console.log('  🎨 Custom blue-500: #5555ff detected');
                console.log('  🔤 Custom text-2xl: 2rem detected');
            }
            else {
                console.error('❌ Custom theme not working as expected');
                console.log('  CSS content:', css2);
            }
        }
        catch (error) {
            console.error('❌ Custom theme test failed:', error);
        }
        // Test 3: Theme with more complex configuration
        console.log('\n📝 Test 3: Testing complex theme configuration...');
        try {
            const themeResult = await themeLoader.loadTheme('./examples/themes/custom.ts');
            const converter3 = new QwikTailwindConverter(themeResult.theme);
            // Create a component that uses more theme properties
            const complexComponent = `
import { component$ } from '@builder.io/qwik';

export default component$(() => {
  return <div class="bg-brand-500 text-success p-6 rounded-xl font-semibold">Complex Theme</div>;
});
      `.trim();
            const complexPath = path.join(testDir, 'complex.tsx');
            await fs.writeFile(complexPath, complexComponent);
            const result3 = await converter3.convertFile(complexPath, `${testDir}/output3`);
            const css3 = await fs.readFile(result3.cssModulesPath, 'utf-8');
            // Check for complex theme values
            const hasBrandColor = css3.includes('rgb(14 165 233'); // brand.500 = #0ea5e9
            const hasSuccessColor = css3.includes('rgb(16 185 129'); // success = #10b981
            const hasCustomSpacing = css3.includes('padding:1.5rem'); // p-6 = 1.5rem
            const hasCustomRadius = css3.includes('border-radius:0.75rem'); // rounded-xl = 0.75rem
            const hasCustomWeight = css3.includes('font-weight:600'); // font-semibold = 600
            if (hasBrandColor && hasSuccessColor && hasCustomSpacing && hasCustomRadius && hasCustomWeight) {
                console.log('✅ Complex theme configuration working correctly');
                console.log('  🎨 Brand color detected');
                console.log('  🎨 Success color detected');
                console.log('  📏 Custom spacing detected');
                console.log('  🔄 Custom border radius detected');
                console.log('  📝 Custom font weight detected');
            }
            else {
                console.error('❌ Complex theme configuration not working as expected');
                console.log('  CSS content:', css3);
            }
        }
        catch (error) {
            console.error('❌ Complex theme test failed:', error);
        }
        console.log('\n🎉 Theme integration tests completed!');
    }
    finally {
        // Cleanup
        await fs.rm(testDir, { recursive: true, force: true });
    }
}
// Run the test
testThemeIntegration().catch(console.error);
//# sourceMappingURL=theme-integration.test.js.map