import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { fetchUserData, handleLogout as logoutUser } from '../services/user';

const Navbar = () => {
    const navigate = useNavigate();
    const [userData, setUserData] = useState(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const data = await fetchUserData();
                setUserData(data);
            } catch (error) {
                console.error(error);
            }
        };

        fetchData().catch((error) => {
            console.error('Error fetching user data:', error);
        });
    }, []);

    const handleLogout = async () => {
        try {
            const success = await logoutUser();
            if (success) {
                navigate('/login', { replace: true });
            }
        } catch (error) {
            console.error(error);
        }
    };

    return (
        <nav>
            <div className="logo">Shoppie</div>
            <input type="text" placeholder="Search..." />
            {userData && (
                <div className="user-info">
                    <span>{userData.username}</span>
                    <span>Balance: {userData.balance}</span>
                </div>
            )}
            <button onClick={handleLogout}>Logout</button>
        </nav>
    );
};

export default Navbar;
