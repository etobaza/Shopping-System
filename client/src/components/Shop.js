import React, { useState, useEffect } from 'react';
import Navbar from './Navbar';
import { fetchItems, fetchItemsSeller, parseCategories } from '../services/item';
import SellerShop from './SellerShop';
import CustomerShop from './CustomerShop';

const Shop = () => {
    const [categories, setCategories] = useState([]);
    const userType = localStorage.getItem('userType');

    useEffect(() => {
        if (userType === 'seller') {
            fetchItemsSeller().then((items) => {
                const categories = parseCategories(items);
                setCategories(categories);
            });
        } else {
            fetchItems().then((items) => {
                const categories = parseCategories(items);
                setCategories(categories);
            });
        }
    }, [userType]);

    return (
        <div>
            <Navbar />
            {userType === 'seller' ? (
                <SellerShop categories={categories} />
            ) : (
                <CustomerShop categories={categories} />
            )}
        </div>
    );
};

export default Shop;
