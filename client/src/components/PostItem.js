import React, { useState } from 'react';
import { createItem } from '../services/item';
import { useNavigate } from 'react-router-dom';

const PostItem = () => {
    const [name, setName] = useState('');
    const [category, setCategory] = useState('');
    const [description, setDescription] = useState('');
    const [price, setPrice] = useState('');
    const [quantity, setQuantity] = useState('');

    const categories = [
        'Electronics',
        'Fashion',
        'Accessories',
        'Home and Garden',
        'Books and Media',
        'Sporting Goods',
        'Toys and Games',
        'Beauty and Personal Care',
        'Food and Beverages',
    ];

    const navigate = useNavigate();

    const handleSubmit = async (event) => {
        event.preventDefault();

        const itemData = {
            name,
            category,
            description,
            price: parseInt(price),
            quantity: parseInt(quantity),
        };

        const response = await createItem(itemData);

        if (response) {
            navigate('/shop');
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <label>
                Name:
                <input type="text" value={name} onChange={(event) => setName(event.target.value)} />
            </label>
            <label>
                Category:
                <select value={category} onChange={(event) => setCategory(event.target.value)}>
                    <option value="">Select a category</option>
                    {categories.map((category) => (
                        <option key={category} value={category}>
                            {category}
                        </option>
                    ))}
                </select>
            </label>
            <label>
                Description:
                <textarea value={description} onChange={(event) => setDescription(event.target.value)} />
            </label>
            <label>
                Price:
                <input type="number" value={price} onChange={(event) => setPrice(event.target.value)} />
            </label>
            <label>
                Quantity:
                <input type="number" value={quantity} onChange={(event) => setQuantity(event.target.value)} />
            </label>
            <button type="submit">Post Item</button>
        </form>
    );
};

export default PostItem;
