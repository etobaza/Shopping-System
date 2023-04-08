import React, { useState, useEffect } from 'react';
import Navbar from './Navbar';
import {useNavigate} from 'react-router-dom';
import jsCookie from "js-cookie";
import {fetchCategories, fetchUserData, handleLogout} from "../services/user";

const Home = () => {
    const [username, setUsername] = useState('');
    const [balance, setBalance] = useState(0);
    const [categories, setCategories] = useState([]);

    useEffect(() => {
        fetchUserData().then(r => console.log(r));
        fetchCategories().then(r => console.log(r));
    }, []);

    return (
        <div>
            <Navbar username={username} balance={balance} onLogout={handleLogout} />
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