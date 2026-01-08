import useCreateOrder from '@/hooks/useCreateOrder'
import { Products } from '@/types/products'
import React, { FormEvent, useState } from 'react'
import { toast, ToastContainer } from 'react-toastify'

const OrderForm = ({ products }: { products: Products[] }) => {
    const { createOrder, loadingCreateOrder, errorCreateOrder } = useCreateOrder()

    const [productsID, setProductsID] = useState(0)
    const [quantity, setQuantity] = useState(1)
    const [discount, setDiscount] = useState(0)

    const submitOrder = async (e: FormEvent) => {
        e.preventDefault()

        if (productsID == 0) {
            toast.error("Please select a product", {
                theme: 'dark',
                position: 'bottom-right',
            })
            return
        }

        if (quantity == 0) {
            toast.error("Please enter a quantity for this product", {
                theme: 'dark',
                position: 'bottom-right',
            })
            return
        }

        const res = await createOrder(productsID, quantity, discount)
        if (res) {
            toast.success("Order created successfully", {
                theme: "dark"
            })
        } else {
            toast.error(errorCreateOrder, {
                theme: 'dark',
                position: 'bottom-right',
            })
        }
    }

    const selectedProduct = products.find((product) => product.id == productsID)
    let subtotal
    let discountPrice
    let total
    if (selectedProduct) {
        subtotal = selectedProduct?.harga * quantity
        discountPrice = subtotal * (discount / 100)
        total = subtotal - discountPrice
    }
    return (
        <div className="space-y-2 mt-4 flex flex-wrap">
            <div className="text-4xl font-bold">Order form</div>
            <div className="w-full flex flex-wrap lg:flex-nowrap gap-2">
                <form onSubmit={submitOrder} className="w-full lg:w-[60%] border border-black/80 rounded-xl p-2 space-y-3">
                    <div className='flex flex-col gap-4'>
                        <select className="w-full border p-2 rounded-lg" onChange={(e) => setProductsID(Number(e.target.value))}>
                            <option value={0}>Pick products</option>
                            {
                                products.map((product) => (
                                    <option key={product.id} value={product.id}> {product.nama_obat} (Stock: {product.stok}) </option>
                                ))
                            }
                        </select>
                        <div className="flex items-center gap-2">
                            <div className="w-30">Qty</div>
                            <input className="w-20 px-2 py-1 border rounded-lg" placeholder="Quantity" type="number" value={quantity} onChange={(e) => setQuantity(Math.max(1, Number(e.target.value)))} />
                        </div>
                        <div className="flex items-center gap-2">
                            <div className="w-30">Discount</div>
                            <input className="w-20 px-2 py-1 border rounded-lg" min={0} max={100} placeholder="Discount" type="number" value={discount} onChange={(e) => setDiscount(Math.min(Number(e.target.value), 100))} />%</div>
                        <button type="submit" className="border px-3 py-1 rounded-lg bg-black text-white hover:bg-black/80 cursor-pointer">
                            {loadingCreateOrder ? "Processing Order" : "Submit Order"}
                        </button>
                    </div>
                </form>
                <div className="w-full lg:w-[40%] border border-black/80 rounded-lg p-4 space-y-2">
                    <div className="mb-6 font-bold text-3xl border-b border-black/50 pb-3">
                        Estimated price
                    </div>
                    <div>Subtotal: Rp.{subtotal}.00</div>
                    <div className="flex gap-10">
                        <div>Discount: Rp.{discountPrice}.00</div>
                        <div className="text-gray-500">{discount}%</div>
                    </div>

                    <div className="font-bold text-xl">Total: Rp.{total}.00</div>
                </div>
            </div>
            <ToastContainer />
        </div>
    )
}

export default OrderForm