import useCreateProduct from '@/hooks/useCreateProduct'
import { Products } from '@/types/products'
import { FormEvent, useState } from 'react'
import { toast, ToastContainer } from 'react-toastify'

const CreateProductForm = ({ setProducts }: { setProducts: React.Dispatch<React.SetStateAction<Products[]>> }) => {
    const [namaObat, setNamaObat] = useState("")
    const [stok, setStok] = useState(0)
    const [harga, setHarga] = useState(0)
    const { createProduct, loadingCreateProduct, errorCreateProduct } = useCreateProduct({ setProducts })

    const submitCreateProduct = async (e: FormEvent) => {
        e.preventDefault()

        if (namaObat == "") {
            toast.error("tolong masukkan nama obat", {
                theme: "dark",
                position: "bottom-right"
            })
            return
        }

        if (stok == 0) {
            toast.error("tolong masukkan stok", {
                theme: "dark",
                position: "bottom-right"
            })
            return
        }

        if (harga == 0) {
            toast.error("tolong masukkan harga", {
                theme: "dark",
                position: "bottom-right"
            })
            return
        }

        const res = await createProduct(namaObat, stok, harga)
        if (res) {
            toast.success("Successfully create a product")
        } else {
            toast.error(errorCreateProduct, {
                theme: "dark",
                position: "bottom-right"
            })
        }


    }

    return (
        <form onSubmit={submitCreateProduct} className="w-full lg:w-[40%] border border-black/80 rounded-lg p-4 space-y-2">
            <div className="mb-6 font-bold text-2xl border-b border-black/50 pb-3">
                Create product
            </div>
            <div className='flex flex-col gap-4'>
                <div>
                    Nama Obat:
                    <input placeholder="Nama obat" className="w-full border px-2 py-1 rounded-lg" onChange={(e) => setNamaObat(e.target.value)} />
                </div>
                <div>
                    Stok:
                    <input placeholder="Stok" type="number" min={1} className="w-full border px-2 py-1 rounded-lg" onChange={(e) => setStok(Math.max(1, Number(e.target.value)))} />
                </div>
                <div>
                    Harga:
                    <div className="flex items-center">
                        Rp.<input placeholder="Harga" type="number" min={1} className="w-full border px-2 py-1 rounded-lg" onChange={(e) => setHarga(Math.max(1, Number(e.target.value)))} />.00
                    </div>
                </div>
                <button type='submit' className="bg-black text-white px-2 py-1 rounded-lg hover:bg-black/80 cursor-pointer">
                    {loadingCreateProduct ? "Creating product" : "Submit"}
                </button>
            </div>
            <ToastContainer />
        </form>
    )
}

export default CreateProductForm