import React from 'react';

const Banner: React.FC = () => {
  return (
    <div className="bg-gradient-to-r from-primary-500 to-primary-600 text-white">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-16">
        <div className="text-center">
          <h1 className="text-4xl md:text-6xl font-bold mb-4">
            Welcome to kholabazar
          </h1>
          <p className="text-xl md:text-2xl mb-8 text-primary-100">
            Your one-stop shop for everything you need
          </p>
          <div className="flex justify-center">
            <button className="bg-white text-primary-600 hover:bg-primary-50 px-8 py-3 rounded-lg font-semibold text-lg transition-colors duration-200">
              Shop Now
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Banner;
