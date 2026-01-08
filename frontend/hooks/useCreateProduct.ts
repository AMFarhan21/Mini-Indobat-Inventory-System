import { Products } from '@/types/products'
import React, { useState } from 'react'

const useCreateProduct = ({ setProducts }: { setProducts: React.Dispatch<React.SetStateAction<Products[]>> }) => {
    const [loadingCreateProduct, setLoading] = useState(false)
    const [errorCreateProduct, setError] = useState("")

    const createProduct = async (nama_obat: string, stok: number, harga: number) => {
        try {
            setLoading(true)
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/products`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ nama_obat, stok, harga })
            })
            const result = await res.json()
            if (!res.ok) {
                throw new Error(result.error)
            }

            setProducts((prev) => [...prev, result.data])

            return result.data
        } catch (error) {
            if (error instanceof Error) {
                setError(error.message)
            } else {
                setError("Failed to create product")
            }
        } finally {
            setLoading(false)
        }
    }

    return { createProduct, loadingCreateProduct, errorCreateProduct }
}

export default useCreateProduct