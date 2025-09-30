import React from 'react';

interface NavbarProps {
  onAddProductClick: () => void;
}

const Navbar: React.FC<NavbarProps> = ({ onAddProductClick }) => {
  return (
    <nav className="bg-white shadow-lg border-b border-primary-100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="flex items-center">
            <h1 className="text-2xl font-bold text-primary-700">kholabazar</h1>
          </div>
          
          <div className="flex items-center space-x-4">
            <button
              onClick={onAddProductClick}
              className="bg-primary-600 hover:bg-primary-700 text-white px-4 py-2 rounded-lg font-medium transition-colors duration-200 flex items-center space-x-2"
            >
              <svg className="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
              </svg>
              <span>Add Product</span>
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
