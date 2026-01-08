import { Products } from '@/types/products'

const ProductsTable = ({ products }: { products: Products[] }) => {
    return (
        <div className="w-full lg:w-[60%] border border-black/80 rounded-xl p-2">
            <table className="w-full border-collapse">
                <thead className="border-b border-black/80">
                    <tr className="text-left">
                        <th className="text-center">Id</th>
                        <th className="p-2">Nama Obat</th>
                        <th className="p-2">Stok</th>
                        <th className="p-2">Harga</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        products.map((product, idx) => (
                            <tr key={product.id} className="text-left">
                                <td className="text-center">{idx + 1}</td>
                                <td className="p-2">{product.nama_obat}</td>
                                <td className="p-2">{product.stok}</td>
                                <td className="p-2">Rp. {product.harga}.00</td>
                            </tr>
                        ))
                    }
                </tbody>
            </table>
        </div>
    )
}

export default ProductsTable