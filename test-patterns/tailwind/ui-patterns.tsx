import { component$ } from '@builder.io/qwik';

export const UIPatterns = component$(() => {
  return (
    <div>
      {/* Hero Section Pattern */}
      <div className="min-h-screen flex items-center justify-center bg-gradient-to-r from-blue-500 to-purple-600">
        <h1 className="text-4xl font-bold text-white">Hero Section</h1>
      </div>

      {/* Card Pattern */}
      <div className="bg-white rounded-lg shadow-lg p-6 max-w-sm">
        <h3 className="text-lg font-semibold">Card Title</h3>
        <p className="text-gray-600">Card content</p>
      </div>

      {/* Modal Overlay Pattern */}
      <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
        <div className="bg-white rounded-lg p-6">Modal Content</div>
      </div>

      {/* Dropdown Menu Pattern */}
      <div className="absolute top-10 right-0 bg-white shadow-lg rounded-md py-2">
        <a href="#" className="block px-4 py-2 hover:bg-gray-100">Menu Item</a>
      </div>

      {/* Badge Pattern */}
      <span className="px-2 py-1 text-xs bg-blue-100 text-blue-800 rounded-full">Badge</span>

      {/* Grid Pattern */}
      <div className="grid grid-cols-3 gap-4">
        <div>Item 1</div>
        <div>Item 2</div>
        <div>Item 3</div>
      </div>

      {/* Accordion Trigger Pattern */}
      <button className="w-full text-left px-4 py-2 hover:bg-gray-50 focus:outline-none">
        Accordion Button
      </button>

      {/* Navigation Pattern */}
      <nav className="flex space-x-4">
        <a href="#" className="text-blue-600 hover:text-blue-800">Home</a>
        <a href="#" className="text-blue-600 hover:text-blue-800">About</a>
      </nav>

      {/* Toast Notification Pattern */}
      <div className="fixed top-4 right-4 bg-green-500 text-white px-4 py-2 rounded-md">
        Success message
      </div>
    </div>
  );
}); 