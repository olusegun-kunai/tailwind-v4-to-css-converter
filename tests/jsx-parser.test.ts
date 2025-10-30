import { JSXParser } from '../src/lib/jsxParser.js';
import fs from 'fs/promises';

async function testJSXParser() {
  console.log('🧪 Testing JSX Parser...\n');
  
  const parser = new JSXParser();
  
  // Read the example checkbox component
  try {
    const checkboxContent = await fs.readFile('./examples/checkbox.tsx', 'utf-8');
    console.log('📄 Parsing checkbox example...\n');
    
    // Test multi-line parsing
    const nodes = parser.parseMultiLineClasses(checkboxContent);
    
    console.log(`🔍 Found ${nodes.length} JSX nodes with classes:\n`);
    
    for (const node of nodes) {
      console.log(`📝 Node: ${node.tagName}`);
      console.log(`   Type: ${node.type}`);
      console.log(`   Semantic Name: ${node.semanticName}`);
      console.log(`   Classes: "${node.className}"`);
      console.log(`   Line: ${node.line}`);
      console.log('---');
    }
    
    // Test individual class extraction
    console.log('\n🎯 Testing class name cleaning...');
    const testClasses = [
      'flex items-center gap-2',
      'size-[25px] rounded-lg relative bg-gray-500',
      '  data-[checked]:flex   justify-center items-center  '
    ];
    
    for (const classes of testClasses) {
      const cleaned = parser.cleanClassName(classes);
      console.log(`Original: "${classes}"`);
      console.log(`Cleaned:  "${cleaned}"`);
      console.log('---');
    }
    
  } catch (error) {
    console.error('❌ Error reading or parsing checkbox file:', error);
  }
}

testJSXParser().catch(console.error); 