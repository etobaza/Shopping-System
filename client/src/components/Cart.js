import { useEffect, useState } from 'react';
import {fetchUser, updateUser} from "../services/user";
import {updateItem} from "../services/item";
import {useNavigate} from "react-router-dom";
const Cart = () => {
    const [items, setItems] = useState([]);
    const [total, setTotal] = useState(0);
    const [balance, setBalance] = useState(null);
    const navigate = useNavigate();
    useEffect(() => {
        const storedCartItems = JSON.parse(localStorage.getItem('cartItems')) || [];
        setItems(storedCartItems);
        setTotal(storedCartItems.reduce((total, item) => total + item.price, 0));

        const userId = localStorage.getItem('userId');
        if (userId) {
            fetchUser(userId).then((data) => {
                setBalance(data.balance);
            });
        }
    }, []);

    const handleCheckoutClick = async () => {
        const userId = localStorage.getItem('userId');
        const newBalance = balance - total;

        if (newBalance < 0) {
            alert('Insufficient balance');
            return;
        }

        // Update user balance
        const { success: balanceUpdateSuccess } = await updateUser(userId, { balance: newBalance });

        if (!balanceUpdateSuccess) {
            alert('Failed to update user balance');
            return;
        }

        // Update item quantities
        const promises = items.map((item) => {
            const { id, quantity } = item;
            const newQuantity = quantity - 1;
            return updateItem(id, { quantity: newQuantity });
        });

        const results = await Promise.all(promises);
        const itemUpdateSuccess = results.every((result) => result.success);

        if (!itemUpdateSuccess) {
            alert('Failed to update item quantities');
            return;
        }

        // Clear cart and navigate to checkout page
        setItems([]);
        setTotal(0);
        localStorage.removeItem('cartItems');
        alert('Checkout successful!')
        navigate('/shop')
    };

    const handleClearCartClick = () => {
        setItems([]);
        setTotal(0);
        localStorage.removeItem('cartItems');
    };

    const removeItemFromCart = (itemId) => {
        const updatedCartItems = items.filter((item) => item.id !== itemId);
        setItems(updatedCartItems);
        setTotal(updatedCartItems.reduce((total, item) => total + item.price, 0));
        localStorage.setItem('cartItems', JSON.stringify(updatedCartItems));
    };

    return (
        <div>
            {items.length > 0 ? (
                <div>
                    <ul>
                        {items.map((item) => (
                            <li key={item.id}>
                                {item.name} - ${item.price.toFixed(2)}
                                <button onClick={() => removeItemFromCart(item.id)}>Remove</button>
                            </li>
                        ))}
                    </ul>
                    <p>Total: ${total.toFixed(2)}</p>
                    <p>Balance: {balance}</p>
                    <button onClick={handleCheckoutClick}>Checkout</button>
                    <button onClick={handleClearCartClick}>Clear Cart</button>
                </div>
            ) : (
                <p>Your cart is empty.</p>
            )}
        </div>
    );
};

export default Cart;
