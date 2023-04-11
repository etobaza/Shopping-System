import React, { useState, useEffect } from 'react';
import { fetchItemsSeller, parseCategories } from '../services/item';
import { Link } from 'react-router-dom';

const SellerShop = () => {
    const [categories, setCategories] = useState([]);

    useEffect(() => {
        fetchItemsSeller().then((items) => {
            const categories = parseCategories(items);
            setCategories(categories);
        });
    }, []);

    return (
        <div>
            <div className="categories">
                <h2>Your Products in Categories</h2>
                <ul>
                    {categories.map((category, index) => (
                        <li key={index}>
                            <Link to={`/shop/c/${category}`}>{category}</Link>
                        </li>
                    ))}
                </ul>
            </div>
            <Link to='/shop/items/new'><button>Post New Item</button></Link>
        </div>
    );
};

export default SellerShop;
