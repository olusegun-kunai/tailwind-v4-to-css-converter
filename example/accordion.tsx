import { component$ } from '@builder.io/qwik';

export const Accordion = component$(() => {
  return (
    <div className="w-full max-w-2xl mx-auto bg-white rounded-lg shadow-lg">
      <div className="border-b border-gray-200">
        <button 
          className="w-full px-6 py-4 text-left text-lg font-semibold text-gray-800 hover:bg-gray-50 focus:outline-none focus:bg-gray-50 transition-colors duration-200"
        >
          <div className="flex items-center justify-between">
            <span>What is Qwik?</span>
            <svg className="w-5 h-5 transform transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
            </svg>
          </div>
        </button>
        <div className="px-6 pb-4 text-gray-600">
          <p className="leading-relaxed">
            Qwik is a new kind of web framework that can deliver instant loading web applications at any size or complexity.
          </p>
        </div>
      </div>
      
      <div className="border-b border-gray-200">
        <button 
          className="w-full px-6 py-4 text-left text-lg font-semibold text-gray-800 hover:bg-gray-50 focus:outline-none focus:bg-gray-50 transition-colors duration-200"
        >
          <div className="flex items-center justify-between">
            <span>How does resumability work?</span>
            <svg className="w-5 h-5 transform transition-transform duration-200" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
            </svg>
          </div>
        </button>
      </div>
    </div>
  );
}); 