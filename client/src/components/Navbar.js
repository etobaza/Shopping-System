import { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { fetchUser, handleLogout as logoutUser } from "../services/user";
import Cart from "./Cart";

function LoggedInNavbar({ firstname, balance, onLogout }) {
    const navigate = useNavigate();

    return (
        <nav>
            <div className="logo">Shoppe</div>
            <div>
                Hello, {firstname}! Your balance is ${balance.toFixed(2)}
            </div>
            <button onClick={onLogout}>Logout</button>
            <button onClick={() => navigate('/cart')}>
                Cart: <Cart />
            </button>
        </nav>
    );
}

function UnloggedInNavbar() {
    const navigate = useNavigate();
    const handleLoginClick = () => {
        navigate('/login');
    };

    return (
        <nav>
            <div className="logo">Shoppe</div>
            <button onClick={handleLoginClick}>Login</button>
        </nav>
    );
}
function Navbar() {
    const [firstname, setFirstName] = useState(null);
    const [balance, setBalance] = useState(null);
    const navigate = useNavigate();
    useEffect(() => {
        const userId = localStorage.getItem('userId');
        if (userId) {
            fetchUser(userId).then((data) => {
                setFirstName(data.firstname);
                setBalance(data.balance);
            });
        }
    }, []);

    const handleLogout = async () => {
        try {
            const success = await logoutUser();
            if (success) {
                localStorage.clear();
                navigate('/login', { replace: true });
            }
        } catch (error) {
            console.error(error);
        }
    };

    return (
        <>
            {firstname ? (
                <LoggedInNavbar
                    firstname={firstname}
                    balance={balance}
                    onLogout={handleLogout}
                />
            ) : (
                <UnloggedInNavbar />
            )}
        </>
    );
}

export default Navbar;