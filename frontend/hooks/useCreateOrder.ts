"use client"

import { Products } from "@/types/products"
import { useState } from "react"

const useCreateOrder = () => {
    const initialProducts: Products = {
        id: 0,
        nama_obat: "",
        harga: 0,
        stok: 0
    }
    const [loadingCreateOrder, setLoading] = useState(false)
    const [errorCreateOrder, setError] = useState("")

    const createOrder = async (product_id: number, quantity: number, discount_percent: number) => {
        try {
            setLoading(true)
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/order`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ product_id, quantity, discount_percent })
            })
            const result = await res.json()
            if (!res.ok) {
                throw new Error(result.error)
            }

            return result.data
        } catch (error) {
            if (error instanceof Error) {
                setError(error.message)
            } else {
                setError("Error on ordering product")
            }
        } finally {
            setLoading(false)
        }
    }

    return { createOrder, loadingCreateOrder, errorCreateOrder }
}

export default useCreateOrder