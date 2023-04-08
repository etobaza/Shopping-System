import React from 'react';

const Navbar = ({ username, balance, onLogout }) => {
    const handleLogout = () => {
        if (onLogout) {
            onLogout();
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
