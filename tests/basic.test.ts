import { UnoGenerator } from '../src/lib/unoGenerator.js';

async function testUnoGenerator() {
  console.log('🧪 Testing UnoCSS Generator...\n');
  
  const generator = new UnoGenerator();
  
  // Test basic Tailwind classes
  const testClasses = [
    'flex justify-center items-center',
    'bg-blue-500 text-white',
    'hover:bg-blue-600 focus:ring-2',
    'p-4 m-2 rounded-lg'
  ];
  
  for (const classes of testClasses) {
    try {
      console.log(`📝 Testing classes: "${classes}"`);
      const result = await generator.generateCSS(classes);
      console.log(`✅ Generated CSS:`);
      console.log(result.css);
      console.log(`📊 Matched classes: ${Array.from(result.matched).join(', ')}`);
      console.log('---\n');
      
    } catch (error) {
      console.error(`❌ Error with classes "${classes}":`, error);
    }
  }
  
  // Test modifier parsing
  console.log('🔍 Testing modifier parsing...');
  try {
    const modifierTest = await generator.generateCSS('bg-red-100 focus:bg-red-500');
    const parsed = generator.parseModifiers(modifierTest.css);
    console.log('Base CSS:', parsed.base);
    console.log('Modifiers:', parsed.modifiers);
  } catch (error) {
    console.error('❌ Modifier test failed:', error);
  }
}

// Run the test
testUnoGenerator().catch(console.error); 