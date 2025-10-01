import React, { useState } from 'react';
import Banner from './Banner';
import ProductCard from './ProductCard';
import Footer from './Footer';
import Modal from './Modal';
import { type Product, type ProductFormData } from '../types';
import { toast } from 'sonner';

interface HomeProps {
  products: Product[];
  setProducts: React.Dispatch<React.SetStateAction<Product[]>>;
  isModalOpen: boolean;
  setIsModalOpen: React.Dispatch<React.SetStateAction<boolean>>;
}

const Home: React.FC<HomeProps> = ({ products, setProducts, isModalOpen, setIsModalOpen }) => {
  const [editingProduct, setEditingProduct] = useState<Product | null>(null);

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
        const res = await fetch("http://localhost:8080/products", {
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
    <>
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
    </>
  );
};

export default Home;
