import { useEffect, useState } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './components/Home';
import ProductDetails from './components/ProductDetails';
import { type Product } from './types';
import { Toaster } from 'sonner';

function App() {
  const [products, setProducts] = useState<Product[]>([]);
  const [isModalOpen, setIsModalOpen] = useState(false);

  useEffect(() => {
    const productsData = async function () {
      const res = await fetch("http://localhost:8080/products")
      const data = await res.json()
      setProducts(data)
    }
    productsData()
  }, []);

  return (
    <Router>
      <div className="min-h-screen bg-gray-50">
        <Navbar onAddProductClick={() => setIsModalOpen(true)} />
        
        <Routes>
          <Route 
            path="/" 
            element={
              <Home 
                products={products} 
                setProducts={setProducts} 
                isModalOpen={isModalOpen}
                setIsModalOpen={setIsModalOpen}
              />
            } 
          />
          <Route 
            path="/product/:id" 
            element={
              <ProductDetails 
                products={products} 
              />
            } 
          />
        </Routes>
        
        <Toaster />
      </div>
    </Router>
  );
}

export default App;
