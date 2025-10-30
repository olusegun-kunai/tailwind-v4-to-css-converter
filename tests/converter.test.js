import { QwikTailwindConverter } from '../src/lib/converter.js';
async function testConverter() {
    console.log('🧪 Testing Full Conversion Pipeline...\n');
    const converter = new QwikTailwindConverter();
    try {
        const result = await converter.convertFile('./examples/checkbox.tsx', './test-output');
        console.log('✅ Conversion completed!');
        console.log(`📁 CSS Modules: ${result.cssModulesPath}`);
        console.log(`📄 Component: ${result.componentPath}`);
        console.log('\n📋 Generated CSS Modules:');
        console.log('---');
        console.log(result.cssContent);
        console.log('---');
        console.log('\n📋 Updated Component:');
        console.log('---');
        console.log(result.updatedComponent);
        console.log('---');
    }
    catch (error) {
        console.error('❌ Conversion failed:', error);
    }
}
testConverter().catch(console.error);
//# sourceMappingURL=converter.test.js.map