"use client"

import { Products } from "@/types/products"
import { useEffect, useState } from "react"


const useProducts = () => {
    const [products, setProducts] = useState<Products[]>([])
    const [loading, setLoading] = useState(false)
    const [errorFetchProducts, setError] = useState("")


    const fetchProducts = async () => {
        try {
            setLoading(true)
            const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/products`)
            const result = await res.json()
            if (!res.ok) {
                throw new Error(result.error)
            }

            setProducts(result.data)
        } catch (error) {
            if (error instanceof Error) {
                setError(error.message)
            }
            console.log(error)
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        fetchProducts()
    }, [])

    return { products, loading, errorFetchProducts, refetch: fetchProducts, setProducts }
}

export default useProducts