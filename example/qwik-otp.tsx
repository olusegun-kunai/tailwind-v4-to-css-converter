import { component$ } from '@builder.io/qwik';

export const OtpInput = component$(() => {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-50">
      <div className="w-full max-w-md p-6 space-y-6 bg-white rounded-lg shadow-md">
        <div className="text-center">
          <h2 className="text-2xl font-bold text-gray-900">Enter OTP</h2>
          <p className="mt-2 text-sm text-gray-600">
            We've sent a verification code to your email
          </p>
        </div>
        
        <div className="flex justify-center space-x-2">
          <input
            type="text"
            maxLength={1}
            className="w-12 h-12 text-center text-lg font-semibold border-2 border-gray-300 rounded-md focus:border-blue-500 focus:outline-none"
          />
          <input
            type="text"
            maxLength={1}
            className="w-12 h-12 text-center text-lg font-semibold border-2 border-gray-300 rounded-md focus:border-blue-500 focus:outline-none"
          />
          <input
            type="text"
            maxLength={1}
            className="w-12 h-12 text-center text-lg font-semibold border-2 border-gray-300 rounded-md focus:border-blue-500 focus:outline-none"
          />
          <input
            type="text"
            maxLength={1}
            className="w-12 h-12 text-center text-lg font-semibold border-2 border-gray-300 rounded-md focus:border-blue-500 focus:outline-none"
          />
        </div>
        
        <button
          type="submit"
          className="w-full px-4 py-2 text-white bg-blue-600 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2"
        >
          Verify OTP
        </button>
        
        <div className="text-center">
          <button
            type="button"
            className="text-sm text-blue-600 hover:text-blue-800 underline"
          >
            Resend code
          </button>
        </div>
      </div>
    </div>
  );
}); 