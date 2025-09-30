import { useEffect, useState } from 'react';
import Navbar from './components/Navbar';
import Banner from './components/Banner';
import ProductCard from './components/ProductCard';
import Footer from './components/Footer';
import Modal from './components/Modal';
import { type Product, type ProductFormData } from './types';
import { toast, Toaster } from 'sonner';

// Sample initial products
const initialProducts: Product[] = [
  {
    id: '1',
    name: 'Wireless Headphones',
    price: 99.99,
    image: 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=300&h=200&fit=crop',
    description: 'High-quality wireless headphones with noise cancellation',
    category: 'Electronics'
  },
  {
    id: '2',
    name: 'Cotton T-Shirt',
    price: 19.99,
    image: 'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=300&h=200&fit=crop',
    description: 'Comfortable 100% cotton t-shirt in various colors',
    category: 'Clothing'
  },
  {
    id: '3',
    name: 'Coffee Mug',
    price: 12.99,
    image: 'https://images.unsplash.com/photo-1520031473529-2c06dea61853?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D',
    description: 'Ceramic coffee mug perfect for your morning brew',
    category: 'Home'
  },
  {
    id: '4',
    name: 'Running Shoes',
    price: 79.99,
    image: 'https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=300&h=200&fit=crop',
    description: 'Lightweight running shoes for all terrains',
    category: 'Sports'
  }
];

function App() {
  const [products, setProducts] = useState<Product[]>([]);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingProduct, setEditingProduct] = useState<Product | null>(null);

  useEffect(() => {
    const productsData = async function () {
      const res = await fetch("http://localhost:8080/products")
      const data = await res.json()
      setProducts(data)
    }
    productsData()
  }, []);

  const handleAddProduct = async (productData: ProductFormData) => {
    if (editingProduct) {
      // Update existing product
      const updatedProduct: Product = {
        ...editingProduct,
        name: productData.name,
        price: parseFloat(productData.price),
        image: productData.image,
        description: productData.description,
        category: productData.category,
      };
      setProducts(products.map(p => p.id === editingProduct.id ? updatedProduct : p));
      setEditingProduct(null);
    } else {
      // Add new product
      const newProduct: Product = {
        name: productData.name,
        price: parseFloat(productData.price),
        image: productData.image,
        description: productData.description,
        category: productData.category,
      };
      try {
        const res = await fetch("http://localhost:8080/create-product", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(newProduct)
        });
        const data = await res.json();

        if (data) {
          toast.success("Product added successfully!")
        }
      } catch (error: any) {
        toast.error(error.message)
      }
      setProducts([...products, newProduct]);
    }
  };

  const handleEditProduct = (product: Product) => {
    setEditingProduct(product);
    setIsModalOpen(true);
  };

  const handleDeleteProduct = (id: string) => {
    setProducts(products.filter(p => p.id !== id));
  };

  const handleModalClose = () => {
    setIsModalOpen(false);
    setEditingProduct(null);
  };

  return (
    <div className="min-h-screen bg-gray-50">
      <Navbar onAddProductClick={() => setIsModalOpen(true)} />
      <Banner />

      {/* Products Section */}
      <section className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-12">
        <h2 className="text-3xl font-bold text-gray-900 mb-8 text-center">
          Our Products
        </h2>

        {products.length === 0 ? (
          <div className="text-center py-12">
            <p className="text-gray-500 text-lg">No products available. Add some products to get started!</p>
          </div>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
            {products.map((product) => (
              <ProductCard
                key={product.id}
                product={product}
                onEdit={handleEditProduct}
                onDelete={handleDeleteProduct}
              />
            ))}
          </div>
        )}
      </section>

      <Footer />

      <Modal
        isOpen={isModalOpen}
        onClose={handleModalClose}
        onSubmit={handleAddProduct}
        editingProduct={editingProduct}
      />
      <Toaster />
    </div>
  );
}

export default App;
