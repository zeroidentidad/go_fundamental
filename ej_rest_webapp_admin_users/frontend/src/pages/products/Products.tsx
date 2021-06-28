/* eslint-disable jsx-a11y/alt-text */
/* eslint-disable jsx-a11y/anchor-is-valid */
import axios from 'axios';
import {useEffect, useState} from 'react';
import Wrapper from "../../components/Wrapper";
import {Link} from "react-router-dom";
import Paginator from "../../components/Paginator";
import {Product} from "../../models/product";

const Products = () => {
    const [products, setProducts] = useState([]);
    const [page, setPage] = useState(1);
    const [lastPage, setLastPage] = useState(0);

    useEffect(() => {
        (
            async () => {
                const {data} = await axios.get(`products?page=${page}`);

                setProducts(data.data);
                setLastPage(data.meta.last_page);
            }
        )()
    }, [page]);

    const del = async (id: number) => {
        if (window.confirm('You want to delete this record?')) {
            await axios.delete(`delete-product/${id}`);

            setProducts(products.filter((p: Product) => p.product_id !== id));
        }
    }

    return (
        <Wrapper>
            <div className="pt-3 pb-2 mb-3 border-bottom">
                <Link to="/products/create" className="btn btn-sm btn-outline-secondary">Create</Link>
            </div>
            <div className="table-responsive">
                <table className="table table-striped table-sm">
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Image</th>
                        <th>Title</th>
                        <th>Description</th>
                        <th>Price</th>
                        <th>Action</th>
                    </tr>
                    </thead>
                    <tbody>
                    {products.map((p: Product) => {
                        return (
                            <tr key={p.product_id}>
                                <td>{p.product_id}</td>
                                <td><img src={p.image} width="50"/></td>
                                <td>{p.title}</td>
                                <td>{p.description}</td>
                                <td>{p.price}</td>
                                <td>
                                    <div className="btn-group mr-2">
                                        <Link 
                                        to={`/products/${p.product_id}/edit`}
                                        className="btn btn-sm btn-outline-secondary"
                                        >Edit</Link>
                                    </div>
                                    &nbsp;
                                    <div className="btn-group mr-2">          
                                        <a href="#" 
                                        className="btn btn-sm btn-outline-secondary"
                                        onClick={() => del(p.product_id)}
                                        >Delete</a>
                                    </div>
                                </td>
                            </tr>
                        )
                    })}
                    </tbody>
                </table>
            </div>

            <Paginator page={page} lastPage={lastPage} pageChanged={setPage}/>
        </Wrapper>
    );
};

export default Products;