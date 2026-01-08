"use client"

import CreateProductForm from "@/components/CreateProductForm";
import OrderForm from "@/components/OrderForm";
import ProductsTable from "@/components/ProductsTable";
import useProducts from "@/hooks/useProducts";

export default function Home() {
  const { products, loading, refetch, setProducts } = useProducts()


  return (
    <div className="py-10 sm:w-[50%] w-full mx-auto">
      <div className="text-5xl font-bold text-center">Mini-Indobat Inventory System</div>
      <div className="mt-12 space-y-4">
        <div className="space-y-2 w-full">
          <div className="text-4xl font-bold">Dashboard</div>
          <button className="border px-3 py-1 rounded-lg cursor-pointer hover:bg-black/80 mt-4 bg-black text-white" onClick={refetch}>Refresh</button>


          <div className="flex items-start flex-wrap lg:flex-nowrap gap-2">
            <ProductsTable products={products} />
            <CreateProductForm setProducts={setProducts} />
          </div>
        </div>
      </div>
      <OrderForm products={products} />
    </div>
  );
}
