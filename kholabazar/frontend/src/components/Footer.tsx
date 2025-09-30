import React from 'react';

const Footer: React.FC = () => {
  return (
    <footer className="bg-primary-600 text-white">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
          <div>
            <h3 className="text-2xl font-bold text-primary-400 mb-4">kholabazar</h3>
            <p className="text-gray-300">
              Your trusted online marketplace for quality products at great prices.
            </p>
          </div>
          
          <div>
            <h4 className="text-lg font-semibold mb-4">Quick Links</h4>
            <ul className="space-y-2 text-gray-300">
              <li><a href="#" className="hover:text-primary-400 transition-colors duration-200">About Us</a></li>
              <li><a href="#" className="hover:text-primary-400 transition-colors duration-200">Contact</a></li>
              <li><a href="#" className="hover:text-primary-400 transition-colors duration-200">Privacy Policy</a></li>
              <li><a href="#" className="hover:text-primary-400 transition-colors duration-200">Terms of Service</a></li>
            </ul>
          </div>
          
          <div>
            <h4 className="text-lg font-semibold mb-4">Contact Info</h4>
            <div className="space-y-2 text-gray-300">
              <p>Email: info@kholabazar.com</p>
              <p>Phone: +880 123 456 789</p>
              <p>Address: Dhaka, Bangladesh</p>
            </div>
          </div>
        </div>
        
        <div className="border-t border-gray-700 mt-8 pt-8 text-center text-gray-400">
          <p>&copy; 2024 kholabazar. All rights reserved.</p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
