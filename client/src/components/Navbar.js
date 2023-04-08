import React from 'react';
import { useNavigate } from 'react-router-dom';

const Navbar = ({ username, balance, onLogout }) => {
    const navigate = useNavigate();
    const handleLogout = async () => {
        try {
            const success = await onLogout();
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
            <div className="user-info">
                <span>{username}</span>
                <span>Balance: {balance}</span>
            </div>
            <button onClick={handleLogout}>Logout</button>
        </nav>
    );
};

export default Navbar;
