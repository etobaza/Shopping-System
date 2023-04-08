import React, { useState, useEffect } from 'react';
import Navbar from './Navbar';
import {useNavigate} from 'react-router-dom';
import jsCookie from "js-cookie";

const Home = () => {
    const [username, setUsername] = useState('');
    const [balance, setBalance] = useState(0);
    const [categories, setCategories] = useState([]);
    const navigate = useNavigate()

    const fetchUserData = async () => {
        const response = await fetch('/user-data');
        const data = await response.json();
        setUsername(data.username);
        setBalance(data.balance);
    };

    const fetchCategories = async () => {
        const response = await fetch('/categories');
        const data = await response.json();
        setCategories(data);
    };

    const handleLogout = async () => {
        try {
            const response = await fetch('/logout', { method: 'POST' });
            if (!response.ok) {
                throw new Error('Failed to logout');
            }
            navigate('/login', { replace: true });
        } catch (error) {
            console.error(error);
        }
    };
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