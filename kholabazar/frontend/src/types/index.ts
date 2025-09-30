export interface Product {
  id?: string;
  name: string;
  price: number;
  image: string;
  description: string;
  category: string;
}

export interface ProductFormData {
  name: string;
  price: string;
  image: string;
  description: string;
  category: string;
}
