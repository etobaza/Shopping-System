import React, { useState, useEffect } from 'react';
import { fetchItems, parseCategories } from '../services/item';
import { Link } from 'react-router-dom';

const CustomerShop = () => {
    const [categories, setCategories] = useState([]);

    useEffect(() => {
        fetchItems().then((items) => {
            const categories = parseCategories(items);
            setCategories(categories);
        });
    }, []);

    return (
        <div>
            <div className="categories">
                <h2>Product Categories</h2>
                <ul>
                    {categories.map((category, index) => (
                        <li key={index}>
                            <Link to={`/shop/c/${category}`}>{category}</Link>
                        </li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default CustomerShop;
