import React, { useState, useEffect } from 'react';
import Navbar from './Navbar';
import {fetchCategories} from "../services/user";

const Home = () => {
    const [categories, setCategories] = useState([]);

    useEffect(() => {
        fetchCategories().then(r => console.log(r));
    }, []);

    return (
        <div>
            <Navbar />
            <div className="categories">
                <h2>Product Categories</h2>
                <ul>
                    {categories.map((category, index) => (
                        <li key={index}>{category}</li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default Home;